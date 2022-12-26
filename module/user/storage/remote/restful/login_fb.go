package userstore

import (
	"context"
	"errors"
	"fmt"
	"food_delivery/common"
	usermodel "food_delivery/module/user/model"
	"food_delivery/plugin/go-sdk/logger"
)

func (s *userRestfulStore) LoginFacebook(ctx context.Context, accessToken string) (*usermodel.LoginFacebookResponse, error) {
	logger := logger.GetCurrent().GetLogger("user.storage.login_fb.go")
	// type requestUserParam struct {
	// 	Ids []int `json:"ids"`
	// }

	var result usermodel.LoginFacebookResponse

	url := fmt.Sprintf("%s%s", common.OauthFacebookUrlAPI, accessToken)
	logger.Debugf("url: %v", url)
	resp, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		// SetBody(requestUserParam{Ids: ids}).
		SetResult(&result).
		Get(url)

	if err != nil {
		logger.Debugf("err: %v", err)

		return nil, err
	}
	logger.Debugln(resp.RawResponse)
	if !resp.IsSuccess() {

		return nil, errors.New("cannot call api get users from google")
	}
	logger.Debugln(result)

	// for i := range result.Data {
	// 	result.Data[i].GetRealId()
	// }

	return &result, nil
}
