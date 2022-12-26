package userstore

import (
	"context"
	usermodel "food_delivery/module/user/model"
)

func (s *userRestfulStore) LoginApple(ctx context.Context, accessToken string, name string) (*usermodel.LoginAppleResponse, error) {
	// // logger := logger.GetCurrent().GetLogger("user.storage.login_apple.go")
	// // type requestUserParam struct {
	// // 	Ids []int `json:"ids"`
	// // }

	// // fmt.Printf("token: %v", token)
	// clientId := "com.pharmacity.socialLoginClient"
	// teamId := "4MXTKML533"
	// keyId := "J8AL7T9PT4"
	// // secretKey := "MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgW+ysr9uN6t8BbbeXXpbiFrdI1SNWzdoWAv/E+yxwmPOgCgYIKoZIzj0DAQehRANCAAQvYjVRXzw0V+JiVawovNem8tGEEWs1jklBaWJQZCrlWglWO9bR2wOQ7NhIrhaWj8hjwFKcHi5pG2Q+BDNl3Tan"
	// secretKey := `-----BEGIN PRIVATE KEY-----
	// MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgW+ysr9uN6t8BbbeX
	// XpbiFrdI1SNWzdoWAv/E+yxwmPOgCgYIKoZIzj0DAQehRANCAAQvYjVRXzw0V+Ji
	// VawovNem8tGEEWs1jklBaWJQZCrlWglWO9bR2wOQ7NhIrhaWj8hjwFKcHi5pG2Q+
	// BDNl3Tan
	// MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgW+ysr9uN6t8BbbeXXpbiFrdI1SNWzdoWAv/E+yxwmPOgCgYIKoZIzj0DAQehRANCAAQvYjVRXzw0V+JiVawovNem8tGEEWs1jklBaWJQZCrlWglWO9bR2wOQ7NhIrhaWj8hjwFKcHi5pG2Q+BDNl3Tan
	// -----END PRIVATE KEY-----`

	// // fmt.Printf("secretKey: %v", secretKey)
	// _, span := trace.StartSpan(ctx, "user.GenerateClientSecret")
	// secret, err := apple.GenerateClientSecret(secretKey, teamId, clientId, keyId)
	// span.End()
	// if err != nil {
	// 	panic(err)
	// }

	// client := apple.New()

	// req := apple.AppValidationTokenRequest{
	// 	ClientID:     clientId,
	// 	ClientSecret: secret,
	// 	Code:         accessToken,
	// }

	// var resp apple.ValidationResponse

	// _, span = trace.StartSpan(ctx, "user.VerifyAppToken")
	// // Do the verification
	// err = client.VerifyAppToken(context.Background(), req, &resp)
	// if err != nil {
	// 	span.End()
	// 	panic(err)
	// }
	// span.End()

	// if resp.Error != "" {
	// 	fmt.Printf("apple returned an error: %s - %s\n", resp.Error, resp.ErrorDescription)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// _, span = trace.StartSpan(ctx, "user.GetUniqueID")
	// // Get the unique user ID
	// userId, err := apple.GetUniqueID(resp.IDToken)
	// if err != nil {
	// 	span.End()
	// 	panic(err)
	// }
	// span.End()
	// // Get the email

	// _, span = trace.StartSpan(ctx, "user.GetClaims")
	// claim, err := apple.GetClaims(resp.IDToken)
	// if err != nil {
	// 	span.End()
	// 	panic(err)
	// }
	// span.End()

	// email := (*claim)["email"].(string)

	// reponse := &usermodel.LoginAppleResponse{
	// 	Id:    userId,
	// 	Email: email,
	// 	Name:  name,
	// }

	return nil, nil
}
