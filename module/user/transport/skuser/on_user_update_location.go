package skuser

import (
	"food_delivery/common"

	"log"

	socketio "github.com/googollee/go-socket.io"
	"gorm.io/gorm"
)

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type SmallAppContext interface {
	GetMainDBConnection() *gorm.DB
}

func OnUserUpdateLocation(appCtx SmallAppContext, requester common.Requester) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {

		// location belong to user ???
		log.Println("User update location: user id is", requester.GetUserId(), "at location", location)
	}
}
