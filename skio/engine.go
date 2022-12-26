package skio

import (
	"fmt"

	"sync"

	goservice "food_delivery/plugin/go-sdk"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
	Run(engine *gin.Engine) error
}

type rtEngine struct {
	server  *socketio.Server
	storage map[int][]AppSocket
	locker  *sync.RWMutex
}

func NewEngine() *rtEngine {
	return &rtEngine{
		storage: make(map[int][]AppSocket),
		locker:  new(sync.RWMutex),
	}
}
func (engine *rtEngine) saveAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	if v, ok := engine.storage[userId]; ok {
		engine.storage[userId] = append(v, appSck)
	} else {
		engine.storage[userId] = []AppSocket{appSck}
	}

	engine.locker.Unlock()
}

func (engine *rtEngine) getAppSocket(userId int) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()

	return engine.storage[userId]
}

func (engine *rtEngine) removeAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	defer engine.locker.RUnlock()

	if v, ok := engine.storage[userId]; ok {
		for i := range v {
			if v[i] == appSck {
				engine.storage[userId] = append(v[:i], v[i+1:]...)
				break
			}
		}
	}
}

func (engine *rtEngine) UserSockets(userId int) []AppSocket {
	var sockets []AppSocket
	if scks, ok := engine.storage[userId]; ok {
		return scks
	}
	return sockets
}
func (engine *rtEngine) EmitToRoom(room string, key string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, key, data)
	return nil
}
func (engine *rtEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := engine.getAppSocket(userId)

	for _, s := range sockets {
		s.Emit(key, data)
	}

	return nil
}
func (engine *rtEngine) Run(r *gin.Engine, sc goservice.ServiceContext) error {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})
	// if err != nil {
	// 	return err
	// }
	engine.server = server

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr(), s.ID())
		return nil
	})

	server.OnError("/", func(c socketio.Conn, err error) {
		fmt.Println("meet error:", err)
	})

	server.OnDisconnect("/", func(c socketio.Conn, reason string) {
		fmt.Println("Closed", reason)
	})

	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		// db := appCtx.GetMainDBConnection()
		// store := userstore.NewSQLStore(db)

		// tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		// payload, err := tokenProvider.Validate(token)
		// if err != nil {
		// 	s.Emit("authentication_failed", err.Error())
		// 	s.Close()
		// 	return
		// }
		// user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UID})

		// if err != nil {
		// 	s.Emit("authentication_failed", err.Error())
		// 	s.Close()
		// 	return
		// }

		// if user.Status == 0 {
		// 	s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
		// 	s.Close()
		// 	return
		// }

		// user.Mask(false)

		// // Important: New AppSocket
		// appSck := NewAppSocket(s, user)
		// engine.saveAppSocket(user.Id, appSck)

		// s.Emit("authenticated", user)

		// // appSck.Join(user.GetRole()) // The same
		// // if user.GetRole() == "admin" {
		// // 	appSck.Join("admin")
		// // }
		// server.OnEvent("/", "UserUpdateLocation", skuser.OnUserUpdateLocation(appCtx, user))
	})
	go server.Serve()
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))

	return nil
}
