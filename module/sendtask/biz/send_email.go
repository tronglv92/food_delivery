package sendemailbiz

import (
	"context"
	tasksModel "food_delivery/module/sendtask/model"
	"food_delivery/plugin/asynqclient"

	"log"

	"github.com/hibiken/asynq"
)

type sendEmailBiz struct {
	clientJob *asynqclient.AsynqClient
}

func NewSendEmailBiz(
	clientJob *asynqclient.AsynqClient,

) *sendEmailBiz {
	return &sendEmailBiz{
		clientJob: clientJob,
	}
}
func (biz *sendEmailBiz) SendEmailTask(ctx context.Context, data tasksModel.EmailTaskRequest) (*asynq.TaskInfo, error) {
	task, err := tasksModel.NewEmailDeliveryTask(data.UserID, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	result, err := biz.clientJob.Enqueue(task)
	if err != nil {
		return nil, err
	}
	return result, nil

}
