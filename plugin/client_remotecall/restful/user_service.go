package remoterestful

import (
	"context"
	"errors"
	"fmt"
	"food_delivery/common"
	"log"
)

// func NewUserService() *userService { return &userService{} }

// type userService struct {
// 	client     *resty.Client
// 	serviceURL string
// 	logger     logger.Logger
// }

// func (*userService) GetPrefix() string {
// 	return common.PluginUserService
// }

// func (s *userService) Get() interface{} {
// 	return s
// }

// func (userService) Name() string {
// 	return common.PluginUserService
// }

// func (s *userService) InitFlags() {
// 	flag.StringVar(&s.serviceURL, s.GetPrefix()+"-url", "", "URL of user service (Ex: http://user-service:8080)")
// }

// func (s *userService) Configure() error {
// 	s.client = resty.New()
// 	s.logger = logger.GetCurrent().GetLogger(s.GetPrefix())

// 	if s.serviceURL == "" {
// 		s.logger.Errorln("Missing service URL")
// 		return errors.New("missing service URL")
// 	}

// 	return nil
// }

// func (s *userService) Run() error {
// 	return s.Configure()
// }

// func (s *userService) Stop() <-chan bool {
// 	c := make(chan bool)

// 	go func() {
// 		c <- true
// 		s.logger.Infoln("Stopped")
// 	}()
// 	return c
// }

func (s *userRestfulStore) GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error) {
	type requestUserParam struct {
		Ids []int `json:"ids"`
	}

	type responseUser struct {
		Data []common.SimpleUser `json:"data"`
	}

	var result responseUser

	resp, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestUserParam{Ids: ids}).
		SetResult(&result).
		Post(fmt.Sprintf("%s/%s", s.serviceURL, "internal/get-users-by-ids"))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if !resp.IsSuccess() {
		log.Println(resp.RawResponse)
		return nil, errors.New("cannot call api get users")
	}

	for i := range result.Data {
		result.Data[i].GetRealId()
	}

	return result.Data, nil
}
