package promotion

import (
	// "errors"
	"testing"

	"github.com/stretchr/testify/assert"

	d "github.com/stefanosKarantin/smb-chatbot/internal/domain"
)

func TestCreatePromotion(t *testing.T) {
	storage := NewPromotionStorage()
	testCases := []struct {
		name              string
		customerID        int
		customerName      string
		telephone         string
		image             string
		coupon int
		expectedPromotion d.Promotion
		expectedErr       error
	}{
		{
			name:         "success",
			customerID:   123,
			customerName: "Giorgos",
			telephone:    "5555555555",
			image:        "image.jpg",
			coupon: 4444,
			expectedPromotion: d.Promotion{
				ID:             1,
				CustomerID:     123,
				CustomerName:   "Giorgos",
				Telephone:      "5555555555",
				CurrentStep:    1,
				DeliveryStatus: "pending",
				Finished:       false,
				Steps:          func() map[int]d.Step {
					promotionSteps := make(map[int]d.Step)
					for _, step := range NewPromotionFlowSteps("Giorgos", "image.jpg", 4444) {
						promotionSteps[step.ID] = step
					}
					return promotionSteps
				}(),
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			promotion, err := storage.CreatePromotion(tc.customerID, tc.customerName, tc.telephone, tc.image, tc.coupon)

			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedPromotion, promotion)
		})
	}
}
