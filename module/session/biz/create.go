package sessionbiz

import (
	"context"
	"food_delivery/common"
	sessionmodel "food_delivery/module/session/model"
)

type CreateSessionStore interface {
	CreateSession(ctx context.Context, data *sessionmodel.SessionCreate) error
}
type createSessionBiz struct {
	store CreateSessionStore
}

func NewCreateSessionBiz(store CreateSessionStore) *createSessionBiz {
	return &createSessionBiz{store: store}
}
func (biz *createSessionBiz) CreateSession(context context.Context,
	data *sessionmodel.SessionCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.store.CreateSession(context, data); err != nil {
		return common.ErrCannotCreateEntity(sessionmodel.EntityName, err)
	}
	return nil
}
