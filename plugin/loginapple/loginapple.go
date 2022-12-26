/*----------------------------------------------------------------*\
 * @author          Ly Nam <lyquocnam@live.com>
 * @copyright       2019 Viet Tran <viettranx@gmail.com>
 * @license         Apache-2.0
 * @description		Plugin to upload image to AWS S3
 *----------------------------------------------------------------*/
package loginapple

import (
	"context"
	"flag"
	"fmt"
	usermodel "food_delivery/module/user/model"
	"food_delivery/plugin/go-sdk/logger"
	"food_delivery/plugin/go-sdk/sdkcm"

	"github.com/Timothylock/go-signin-with-apple/apple"
)

var (
	ErrAppleClientIDMissing  = sdkcm.CustomError("ErrAppleClientIDMissing", "Apple Client ID is missing")
	ErrAppleTeamIDMissing    = sdkcm.CustomError("ErrAppleTeamIDMissing", "Apple TeamID is missing")
	ErrAppleKeyIDMissing     = sdkcm.CustomError("ErrAppleKeyIDMissing", "Apple Key ID is missing")
	ErrAppleSecretKeyMissing = sdkcm.CustomError("ErrAppleSecretKeyMissing", "Apple Secret Key is missing")
)

type LoginApple interface {
	LoginAppleByAccessToken(ctx context.Context, accessToken string, name string) (*usermodel.LoginAppleResponse, error)
}
type loginApple struct {
	name   string
	prefix string
	logger logger.Logger

	cfg appleConfig

	client *apple.Client
}

type appleConfig struct {
	clientId  string
	teamId    string
	keyId     string
	secretKey string
}

func NewLoginApple(prefix ...string) *loginApple {
	pre := "apple"

	if len(prefix) > 0 {
		pre = prefix[0]
	}

	return &loginApple{
		name:   "apple",
		prefix: pre,
	}
}

func (s *loginApple) Get() interface{} {
	return s
}

func (s *loginApple) Name() string {
	return s.name
}

func (s *loginApple) InitFlags() {

	flag.StringVar(&s.cfg.clientId, fmt.Sprintf("%s-%s", s.prefix, "client-id"), "", "Apple Client ID")
	flag.StringVar(&s.cfg.teamId, fmt.Sprintf("%s-%s", s.GetPrefix(), "team-id"), "", "Apple team id")
	flag.StringVar(&s.cfg.keyId, fmt.Sprintf("%s-%s", s.GetPrefix(), "key-id"), "", "Apple key id")
	flag.StringVar(&s.cfg.secretKey, fmt.Sprintf("%s-%s", s.GetPrefix(), "secret-key"), "", "Apple secret key")

}

func (s *loginApple) Configure() error {
	s.logger = logger.GetCurrent().GetLogger(s.Name())

	if err := s.cfg.check(); err != nil {
		s.logger.Errorln(err)
		return err
	}

	client := apple.New()

	s.client = client

	s.logger.Infoln("Connected Login Apple")

	return nil
}

func (s *loginApple) GetPrefix() string {
	return s.prefix
}

func (s *loginApple) Run() error {
	return s.Configure()
}

func (s *loginApple) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
func (s *loginApple) LoginAppleByAccessToken(ctx context.Context, accessToken string, name string) (*usermodel.LoginAppleResponse, error) {

	s.logger.Debugln("GenerateClientSecret secretKey: ", s.cfg.secretKey)
	s.logger.Debugln("GenerateClientSecret s.cfg.teamId: ", s.cfg.teamId)
	s.logger.Debugln("GenerateClientSecret s.cfg.clientId: ", s.cfg.clientId)
	s.logger.Debugln("GenerateClientSecret s.cfg.keyId: ", s.cfg.keyId)
	secret, err := apple.GenerateClientSecret(s.cfg.secretKey, s.cfg.teamId, s.cfg.clientId, s.cfg.keyId)

	if err != nil {
		s.logger.Debugln("GenerateClientSecret err: ", err)
		return nil, err
	}
	req := apple.AppValidationTokenRequest{
		ClientID:     s.cfg.clientId,
		ClientSecret: secret,
		Code:         accessToken,
	}

	var resp apple.ValidationResponse

	// Do the verification
	err = s.client.VerifyAppToken(context.Background(), req, &resp)
	if err != nil {

		return nil, err
	}

	if resp.Error != "" {
		fmt.Printf("apple returned an error: %s - %s\n", resp.Error, resp.ErrorDescription)
		if err != nil {
			return nil, err
		}
	}

	// Get the unique user ID
	userId, err := apple.GetUniqueID(resp.IDToken)
	if err != nil {

		return nil, err
	}

	// Get the email

	claim, err := apple.GetClaims(resp.IDToken)
	if err != nil {

		return nil, err
	}
	email := (*claim)["email"].(string)
	reponse := &usermodel.LoginAppleResponse{
		Id:    userId,
		Email: email,
		Name:  name,
	}
	return reponse, nil
}

func (cfg *appleConfig) check() error {
	if len(cfg.clientId) < 1 {
		return ErrAppleClientIDMissing
	}
	if len(cfg.teamId) < 1 {
		return ErrAppleTeamIDMissing
	}
	if len(cfg.keyId) < 1 {
		return ErrAppleKeyIDMissing
	}
	if len(cfg.secretKey) < 1 {
		return ErrAppleSecretKeyMissing
	}
	return nil
}
