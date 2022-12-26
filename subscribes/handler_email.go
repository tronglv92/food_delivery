package subscribes

import (
	"context"
	"encoding/json"
	"fmt"
	tasksModel "food_delivery/module/sendtask/model"
	"log"

	"github.com/hibiken/asynq"
)

func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p tasksModel.EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Email to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID)
	// Email delivery code ...
	return nil
}
