package restaurantstorage

import (
	"context"
	"encoding/json"
	"food_delivery/common"
	restaurantmodel "food_delivery/module/restaurant/model"
	"food_delivery/plugin/go-sdk/logger"
	"log"
	"strings"

	"github.com/olivere/elastic/v7"
)

type RestaurantTest struct {
	Name    string `json:"name"`
	Address string `json:"addr"`
}

func (s *esStore) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	logger := logger.GetCurrent().GetLogger("restaurant.storage.elastic.list")
	var results []restaurantmodel.Restaurant

	var empty []restaurantmodel.Restaurant

	q := elastic.NewBoolQuery()
	if f := filter; f != nil {
		if f.OwnerId > 0 {
			q = q.Must(elastic.NewMatchQuery("user_id", f.OwnerId))
		}

		search := strings.TrimSpace(f.Search)
		if len(search) > 0 {
			q = q.Must(elastic.NewFuzzyQuery("name", search))
		}
	}

	// q = q.Must(elastic.NewRangeQuery("shipping_fee_per_km").From(2).To(3))
	src, err := q.Source()
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(src)
	if err != nil {
		log.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	logger.Debugf("got", got)

	offset := (paging.Page - 1) * paging.Limit
	searchResult, err := s.client.Search().
		Index(restaurantmodel.EntityName). // search in index "twitter"
		Query(q).
		// specify the query
		Sort("id", false).                      // sort by "user" field, ascending
		From(offset).Size(paging.Limit).        // take documents 0-9
		Pretty(true).RestTotalHitsAsInt(false). // pretty print request and response JSON
		Do(context.Background())                // execute
	if err != nil {
		// Handle error
		return empty, err
	}

	// TotalHits is another convenience function that works even when something goes wrong.
	logger.Debugf("Found a total of %d restaurants\n", searchResult.TotalHits())
	// Here's how you iterate through results with full control over each step.
	if searchResult.TotalHits() > 0 {

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var rt restaurantmodel.RestaurantES
			err := json.Unmarshal(hit.Source, &rt)
			if err != nil {
				// Deserialization failed
				logger.Debugf("Deserialization failed")
			}
			results = append(results, rt.ToRestaurant())
			// Work with tweet
			logger.Debugf("Restaurant by %s: %s\n", rt.Id, rt.Addr)
		}
	} else {
		// No hits
		logger.Debugf("Found no restaurant\n")
	}

	paging.Total = searchResult.TotalHits()
	return results, nil
}
