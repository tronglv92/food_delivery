package imgprocessing

import "food_delivery/plugin/go-sdk/sdkcm"

type Response struct {
	sdkcm.AppError
	Data *sdkcm.Image `json:"data"`
}
