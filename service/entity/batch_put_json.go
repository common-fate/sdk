package entity

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	entityv1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	"github.com/common-fate/sdk/uid"
	"github.com/pkg/errors"
)

// EntityJSON is a JSON representation of entities.
// It matches the Cedar Rust SDK JSON format.
type EntityJSON struct {
	UID     uid.UID        `json:"uid"`
	Attrs   map[string]any `json:"attrs"`
	Parents []uid.UID      `json:"parents"`
}

type BatchPutJSONInput struct {
	Entities []EntityJSON
}

func (c *Client) BatchPutJSON(ctx context.Context, input BatchPutJSONInput) (*BatchUpdateOutput, error) {
	var req = &entityv1alpha1.BatchUpdateRequest{
		Universe: "default",
		Put:      []*entityv1alpha1.Entity{},
	}

	for _, e := range input.Entities {
		parsed, children, err := transformJSONToEntity(e)
		if err != nil {
			return nil, err
		}

		req.Put = append(req.Put, parsed)
		req.PutChildren = append(req.PutChildren, children...)
	}

	res, err := c.raw.BatchUpdate(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}

func transformJSONToEntity(e EntityJSON) (*entityv1alpha1.Entity, []*entityv1alpha1.ChildRelation, error) {
	attrs, err := transformAttrs(e.Attrs)
	if err != nil {
		return nil, nil, err
	}

	res := entityv1alpha1.Entity{
		Uid:        e.UID.ToAPI(),
		Attributes: attrs,
	}

	var children []*entityv1alpha1.ChildRelation

	for _, v := range e.Parents {
		children = append(children, &entityv1alpha1.ChildRelation{
			Parent: v.ToAPI(),
			Child:  e.UID.ToAPI(),
		})
	}

	return &res, children, nil
}

func transformAttrs(attrs map[string]any) ([]*entityv1alpha1.Attribute, error) {
	res := []*entityv1alpha1.Attribute{}

	for k, v := range attrs {
		attr, err := transformAttr(v)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing %s", k)
		}

		res = append(res, &entityv1alpha1.Attribute{
			Key:   k,
			Value: attr,
		})
	}

	return res, nil
}

func transformAttr(v any) (*entityv1alpha1.Value, error) {
	switch val := v.(type) {
	case string:
		return &entityv1alpha1.Value{
			Value: &entityv1alpha1.Value_Str{
				Str: val,
			},
		}, nil

	case int:
		return &entityv1alpha1.Value{
			Value: &entityv1alpha1.Value_Long{
				Long: int64(val),
			},
		}, nil

	case float64:
		return &entityv1alpha1.Value{
			Value: &entityv1alpha1.Value_Long{
				Long: int64(val),
			},
		}, nil

	case []any:
		setValue := entityv1alpha1.Value_Set{
			Set: &entityv1alpha1.Set{
				Values: []*entityv1alpha1.Value{},
			},
		}

		for _, nested := range val {
			attr, err := transformAttr(nested)
			if err != nil {
				return nil, err
			}
			setValue.Set.Values = append(setValue.Set.Values, attr)
		}

		return &entityv1alpha1.Value{Value: &setValue}, nil

	case map[string]any:
		entityMap, ok := val["__entity"]
		if ok {
			entityMapStr, ok := entityMap.(map[string]string)
			if ok {
				typ := entityMapStr["type"]
				if typ == "" {
					return nil, errors.New("could not parse __entity reference: type was empty")
				}

				id := entityMapStr["id"]
				if typ == "" {
					return nil, errors.New("could not parse __entity reference: id was empty")
				}

				return &entityv1alpha1.Value{
					Value: &entityv1alpha1.Value_Entity{
						Entity: &entityv1alpha1.UID{
							Type: typ,
							Id:   id,
						},
					},
				}, nil
			}

			entity, ok := entityMap.(map[string]any)
			if !ok {
				return nil, errors.New("could not parse __entity reference")
			}
			typ := entity["type"]
			if typ == "" {
				return nil, errors.New("could not parse __entity reference: type was empty")
			}
			typStr, ok := typ.(string)
			if !ok {
				return nil, errors.New("could not parse __entity reference: type was not string")
			}

			id := entity["id"]
			if typ == "" {
				return nil, errors.New("could not parse __entity reference: id was empty")
			}
			idStr, ok := id.(string)
			if !ok {
				return nil, errors.New("could not parse __entity reference: id was not string")
			}

			return &entityv1alpha1.Value{
				Value: &entityv1alpha1.Value_Entity{
					Entity: &entityv1alpha1.UID{
						Type: typStr,
						Id:   idStr,
					},
				},
			}, nil
		} else {
			record, err := transformAttrs(val)
			if err != nil {
				return nil, err
			}

			return &entityv1alpha1.Value{
				Value: &entityv1alpha1.Value_Record{
					Record: &entityv1alpha1.Record{
						Attributes: record,
					},
				},
			}, nil
		}

	default:
		return nil, fmt.Errorf("unhandled attribute type %T", v)
	}
}
