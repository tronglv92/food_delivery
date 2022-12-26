package ginsendjob

import (
	"food_delivery/common"
	tasksModel "food_delivery/module/sendtask/model"
	"log"

	"net/http"

	"food_delivery/plugin/asynqclient"
	goservice "food_delivery/plugin/go-sdk"

	sendemailbiz "food_delivery/module/sendtask/biz"

	"github.com/gin-gonic/gin"
)


func SendEmail(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var data tasksModel.EmailTaskRequest

		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		asynqClient := sc.MustGet(common.PluginAsynqClient).(*asynqclient.AsynqClient)

		biz := sendemailbiz.NewSendEmailBiz(asynqClient)
		info, err := biz.SendEmailTask(ctx.Request.Context(), data)
		if err != nil {
			panic(common.ErrInternal(err))
		}
		log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse("Send email success"))
	}
}
