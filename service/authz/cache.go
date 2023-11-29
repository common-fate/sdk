package authz

import (
	"context"
	"errors"
	"fmt"

	"github.com/common-fate/apikit/logger"
	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	"github.com/common-fate/sdk/service/authz/uid"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
)

type AttrCache struct {
	client *Client
}

// AttrCache returns a cach for looking up entity attributes.
// This cache should ONLY be used for attributes not related to authorization
// decisions, such as displaying names, email addresses and labels.
func (c *Client) AttrCache() *AttrCache {
	return &AttrCache{client: c}
}

// String obtains a cached string attribute for an entity. Attributes like names do not change that often, so they are cached in memory.
// If the attribute does not exist for an entity or there is an error, an empty string is returned
//
// An error is logged if there was an issue reaching the Authz service.
func (c *AttrCache) String(ctx context.Context, uid uid.UID, attr string) string {
	log := logger.Get(ctx).Named("namecache").With("uid", uid, "attr", attr)

	cached, ok := c.client.cache.Get(uid.String())
	if ok {
		// cache hit
		res, err := extractStringAttr(cached, attr)
		if err != nil {
			log.Errorw("error extracting attribute from cached entity", zap.Error(err))
			return ""
		}

		// renew the cached duration
		c.client.cache.Set(uid.String(), cached, cache.DefaultExpiration)

		return res
	}

	log.Debugw("cache miss")

	res, err := c.client.GetEntity(ctx, GetEntityInput{
		UID: uid,
	})
	if err != nil {
		log.Errorw("error loading entity from authz", zap.Error(err))
		return ""
	}
	if res == nil {
		log.Errorw("entity did not exist", zap.Error(err))
		return ""
	}

	for _, a := range res.Entity.Attributes {
		if a.Key == attr {
			return a.Value.GetStr()
		}
	}

	err = fmt.Errorf("entity did not have attribute %s", attr)

	log.Errorw("error extracting attribute from loaded entity", zap.Error(err))

	return ""
}

func extractStringAttr(cached any, attr string) (string, error) {
	entity, ok := cached.(*authzv1alpha1.Entity)
	if !ok {
		return "", errors.New("could not cast to *authzv1alpha1.Entity")
	}
	for _, a := range entity.Attributes {
		if a.Key == attr {
			return a.Value.GetStr(), nil
		}
	}

	return "", fmt.Errorf("entity did not have attribute %s", attr)
}
