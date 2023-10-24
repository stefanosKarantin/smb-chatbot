package promotion

import (
	"errors"
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
		coupon            int
		expectedPromotion d.Promotion
		expectedErr       error
	}{
		{
			name:         "success",
			customerID:   123,
			customerName: "Giorgos",
			telephone:    "5555555555",
			image:        "image.jpg",
			coupon:       4444,
			expectedPromotion: d.Promotion{
				ID:             1,
				CustomerID:     123,
				CustomerName:   "Giorgos",
				Telephone:      "5555555555",
				CurrentStep:    1,
				DeliveryStatus: "pending",
				Finished:       false,
				Steps: func() map[int]d.Step {
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

func TestGetPromotionByID(t *testing.T) {
	storage := NewPromotionStorage()
	expectedPromotion, err := storage.CreatePromotion(123, "Giorgos", "11111111", "image.jpg", 4444)
	assert.NoError(t, err)

	//Test success
	result, err := storage.GetPromotionByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedPromotion, result)

	// Test error
	expectedError := errors.New("promotion not found")
	_, err = storage.GetPromotionByID(2)
	assert.Equal(t, expectedError, err)
}

func TestUpdateDeliveryStatus(t *testing.T) {
	storage := NewPromotionStorage()
	_, err := storage.CreatePromotion(123, "Giorgos", "11111111", "image.jpg", 4444)
	assert.NoError(t, err)

	//Test success
	err = storage.UpdateDeliveryStatus(1, "delivered")
	assert.NoError(t, err)

	// Test error
	expectedError := errors.New("promotion not found")
	err = storage.UpdateDeliveryStatus(2, "delivered")
	assert.Equal(t, expectedError, err)
}

func TestUpdateCurrentStep(t *testing.T) {
	storage := NewPromotionStorage()
	_, err := storage.CreatePromotion(123, "Giorgos", "11111111", "image.jpg", 4444)
	assert.NoError(t, err)

	//Test success
	err = storage.UpdateCurrentStep(1, 2)
	assert.NoError(t, err)

	// Test error
	expectedError := errors.New("promotion not found")
	err = storage.UpdateCurrentStep(2, 2)
	assert.Equal(t, expectedError, err)
}

func TestUpdateFinished(t *testing.T) {
	storage := NewPromotionStorage()
	_, err := storage.CreatePromotion(123, "Giorgos", "11111111", "image.jpg", 4444)
	assert.NoError(t, err)

	//Test success
	err = storage.UpdateFinished(1)
	assert.NoError(t, err)

	// Test error
	expectedError := errors.New("promotion not found")
	err = storage.UpdateFinished(2)
	assert.Equal(t, expectedError, err)
}

func TestGetLastPromotionID(t *testing.T) {
	storage := NewPromotionStorage()
	_, err := storage.CreatePromotion(123, "Giorgos", "11111111", "image.jpg", 4444)
	assert.NoError(t, err)

	//Test success 1
	expectedID := 1
	ID := storage.GetLastPromotionID()
	assert.Equal(t, expectedID, ID)

	//Test success 2
	_, err = storage.CreatePromotion(345, "Giorgos", "11411111", "image.jpg", 4444)
	assert.NoError(t, err)
	
	expectedID = 2
	ID = storage.GetLastPromotionID()
	assert.Equal(t, expectedID, ID)
}
