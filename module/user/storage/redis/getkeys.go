package userstore

import (
	"context"
	"fmt"
)

const cachePrefixAT = "user:%d:*"

func (c *authUserCached) GetKeys(ctx context.Context, conditions map[string]interface{}) ([]string, error) {

	var results []string
	userId := conditions["id"].(int)

	key := fmt.Sprintf(cachePrefixAT, userId)

	for {

		scanCmd := c.cacheStore.Scan(ctx, 0, key, 100)
		keys, cursor, err := scanCmd.Result()
		if err != nil {
			return nil, err
		}
		results = append(results, keys...)
		fmt.Printf("cursor ", cursor)
		if cursor == 0 {
			break
		}

	}

	return results, nil
}
