package userstore

import "context"

func (c *authUserCached) DelKeys(ctx context.Context, keys []string) error {
	for _, key := range keys {
		_ = c.cacheStore.Delete(ctx, key)
	}
	return nil
}
