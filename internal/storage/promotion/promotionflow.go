package promotion

import (
	"fmt"

	d "github.com/stefanosKarantin/smb-chatbot/internal/domain"
)

func NewPromotionFlowSteps(customerName string, image string, coupon int) []d.Step {
	steps := []d.Step{
		{
			ID:         1,
			StepNumber: 1,
			Text:       NewWelcomeText(customerName),
			Responses: []d.Response{
				{
					ID:       1,
					Text:     "Yes! Show me the coupon",
					NextStep: 2,
				},
				{
					ID:       2,
					Text:     "No, I'm not interested",
					NextStep: 3,
				},
			},
		},
		{
			ID:         2,
			StepNumber: 2,
			Text:       "Here is our unique promotional coupon! \n10% off. Limit 1 per Customer",
			Responses: []d.Response{
				{
					ID:       3,
					Text:     "Reveal Coupon",
					NextStep: 4,
				},
			},
		},
		{
			ID:         3,
			StepNumber: 2,
			Text:       "No worries! \nHave a nice day!",
			Image:      image,
			Responses:  []d.Response{},
		},
		{
			ID:         4,
			StepNumber: 3,
			Text:       fmt.Sprint(coupon),
			Responses:  []d.Response{},
		},
	}

	return steps
}

func NewWelcomeText(customerName string) string {
	return fmt.Sprintf("Welcome, %s! How can we assist you today?", customerName)
}
