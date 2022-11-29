package userstore

import (
	"context"
	"fmt"
)



func (c *authUserCached) GetKeys(ctx context.Context, conditions map[string]interface{}) ([]string, error) {

	var results []string
	userId := conditions["id"].(int)

	key := fmt.Sprintf(cachePrefixAT, userId)

	var cursor uint64
	for {

		scanCmd := c.cacheStore.Scan(ctx, cursor, key, 2)
		keys, cs, err := scanCmd.Result()
		if err != nil {
			return nil, err
		}
		cursor = cs
		results = append(results, keys...)
		// fmt.Printf("cursor %v", cursor)
		if cursor == 0 {
			break
		}

	}

	return results, nil
}
