package authz

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/pkg/errors"
)

// EntityJSON is a JSON representation of entities.
// It matches the Cedar Rust SDK JSON format.
type EntityJSON struct {
	UID     UID            `json:"uid"`
	Attrs   map[string]any `json:"attrs"`
	Parents []UID          `json:"parents"`
}

type BatchPutEntityJSONInput struct {
	Entities []EntityJSON
}

func (c *Client) BatchPutEntityJSON(ctx context.Context, input BatchPutEntityJSONInput) (*authzv1alpha1.BatchPutEntityResponse, error) {
	var req = &authzv1alpha1.BatchPutEntityRequest{
		Universe: "default",
		Entities: []*authzv1alpha1.Entity{},
	}

	for _, e := range input.Entities {
		parsed, err := transformJSONToEntity(e)
		if err != nil {
			return nil, err
		}

		req.Entities = append(req.Entities, parsed)
	}

	res, err := c.raw.BatchPutEntity(ctx, connect.NewRequest(req))
	if err != nil {
		return nil, err
	}

	return res.Msg, nil
}

func transformJSONToEntity(e EntityJSON) (*authzv1alpha1.Entity, error) {
	attrs, err := transformAttrs(e.Attrs)
	if err != nil {
		return nil, err
	}

	res := authzv1alpha1.Entity{
		Uid:        e.UID.ToAPI(),
		Attributes: attrs,
		Parents:    []*authzv1alpha1.UID{},
	}

	for _, v := range e.Parents {
		res.Parents = append(res.Parents, v.ToAPI())
	}

	return &res, nil
}

func transformAttrs(attrs map[string]any) ([]*authzv1alpha1.Attribute, error) {
	res := []*authzv1alpha1.Attribute{}

	for k, v := range attrs {
		attr, err := transformAttr(v)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing %s", k)
		}

		res = append(res, &authzv1alpha1.Attribute{
			Key:   k,
			Value: attr,
		})
	}

	return res, nil
}

func transformAttr(v any) (*authzv1alpha1.Value, error) {
	switch val := v.(type) {
	case string:
		return &authzv1alpha1.Value{
			Value: &authzv1alpha1.Value_Str{
				Str: val,
			},
		}, nil

	case int:
		return &authzv1alpha1.Value{
			Value: &authzv1alpha1.Value_Long{
				Long: int64(val),
			},
		}, nil

	case float64:
		return &authzv1alpha1.Value{
			Value: &authzv1alpha1.Value_Long{
				Long: int64(val),
			},
		}, nil

	case []any:
		setValue := authzv1alpha1.Value_Set{
			Set: &authzv1alpha1.Set{
				Values: []*authzv1alpha1.Value{},
			},
		}

		for _, nested := range val {
			attr, err := transformAttr(nested)
			if err != nil {
				return nil, err
			}
			setValue.Set.Values = append(setValue.Set.Values, attr)
		}

		return &authzv1alpha1.Value{Value: &setValue}, nil

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

				return &authzv1alpha1.Value{
					Value: &authzv1alpha1.Value_Entity{
						Entity: &authzv1alpha1.UID{
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

			return &authzv1alpha1.Value{
				Value: &authzv1alpha1.Value_Entity{
					Entity: &authzv1alpha1.UID{
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

			return &authzv1alpha1.Value{
				Value: &authzv1alpha1.Value_Record{
					Record: &authzv1alpha1.Record{
						Attributes: record,
					},
				},
			}, nil
		}

	default:
		return nil, fmt.Errorf("unhandled attribute type %T", v)
	}
}
