package sessionbiz

import (
	"context"
	"food_delivery/common"
	sessionmodel "food_delivery/module/session/model"
)

type GetSessionStore interface {
	GetSession(ctx context.Context, condition map[string]interface{},
		moreKeys ...string) (*sessionmodel.Session, error)
}
type getSessionBiz struct {
	store GetSessionStore
}

func NewGetSessionBiz(store GetSessionStore) *getSessionBiz {
	return &getSessionBiz{store: store}
}
func (biz *getSessionBiz) GetSession(context context.Context, id string) (*sessionmodel.Session, error) {

	session, err := biz.store.GetSession(context, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotCreateEntity(sessionmodel.EntityName, err)
	}
	return session, nil
}
