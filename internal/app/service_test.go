package app_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stefanosKarantin/smb-chatbot/internal/app"
	d "github.com/stefanosKarantin/smb-chatbot/internal/domain"
	m "github.com/stefanosKarantin/smb-chatbot/pkg/messageplatform"
)

func TestService_StartPromotion(t *testing.T) {
	testCases := []struct {
		name          string
		customerID    int
		customerName  string
		telephone     string
		image         string
		promoRepo     d.PromotionRepo
		messageClient m.MessageClient
		expectedErr   error
	}{
		{
			name:         "success",
			customerID:   1,
			customerName: "John",
			telephone:    "1234567890",
			image:        "image.png",
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    1,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps: map[int]d.Step{
						1: {
							ID:        1,
							Text:      "hello",
							Responses: []d.Response{},
						},
					},
				}
				mp.On(
					"CreatePromotion",
					promotion.CustomerID,
					promotion.CustomerName,
					promotion.Telephone,
					"image.png",
				).Return(promotion, nil)
				mp.On("UpdateDeliveryStatus", promotion.ID, "sent").Return(nil)
				return &mp
			}(),
			messageClient: func() *m.MockMessageClient {
				mc := m.MockMessageClient{}
				promoMsg := m.PromotionMessage{
					ID:           2,
					CustomerID:   1,
					CustomerName: "John",
					Telephone:    "1234567890",
					Message: m.Message{
						Text: "hello",
					},
					Responses: []m.Response{},
				}
				mc.On("SendPromotionMessage", promoMsg).Return(nil)
				return &mc
			}(),
			expectedErr: nil,
		},
		{
			name:         "failed to create promotion",
			customerID:   1,
			customerName: "John",
			telephone:    "1234567890",
			image:        "image.png",
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    1,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps: map[int]d.Step{
						1: {
							ID:        1,
							Text:      "hello",
							Responses: []d.Response{},
						},
					},
				}
				mp.On(
					"CreatePromotion",
					promotion.CustomerID,
					promotion.CustomerName,
					promotion.Telephone,
					"image.png",
				).Return(promotion, errors.New("failed to create promotion"))
				return &mp
			}(),
			messageClient: new(m.MockMessageClient),
			expectedErr:   errors.New("failed to create promotion"),
		},
		{
			name:         "failed to send messsage",
			customerID:   1,
			customerName: "John",
			telephone:    "1234567890",
			image:        "image.png",
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    1,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps: map[int]d.Step{
						1: {
							ID:        1,
							Text:      "hello",
							Responses: []d.Response{},
						},
					},
				}
				mp.On(
					"CreatePromotion",
					promotion.CustomerID,
					promotion.CustomerName,
					promotion.Telephone,
					"image.png",
				).Return(promotion, nil)
				return &mp
			}(),
			messageClient: func() *m.MockMessageClient {
				mc := m.MockMessageClient{}
				promoMsg := m.PromotionMessage{
					ID:           2,
					CustomerID:   1,
					CustomerName: "John",
					Telephone:    "1234567890",
					Message: m.Message{
						Text: "hello",
					},
					Responses: []m.Response{},
				}
				mc.On("SendPromotionMessage", promoMsg).Return(errors.New("failed to send messsage"))
				return &mc
			}(),
			expectedErr: errors.New("failed to send messsage"),
		},
		{
			name:         "failed to update delivery status",
			customerID:   1,
			customerName: "John",
			telephone:    "1234567890",
			image:        "image.png",
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    1,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps: map[int]d.Step{
						1: {
							ID:        1,
							Text:      "hello",
							Responses: []d.Response{},
						},
					},
				}
				mp.On(
					"CreatePromotion",
					promotion.CustomerID,
					promotion.CustomerName,
					promotion.Telephone,
					"image.png",
				).Return(promotion, nil)
				mp.On("UpdateDeliveryStatus", promotion.ID, "sent").Return(errors.New("failed to update delivery status"))
				return &mp
			}(),
			messageClient: func() *m.MockMessageClient {
				mc := m.MockMessageClient{}
				promoMsg := m.PromotionMessage{
					ID:           2,
					CustomerID:   1,
					CustomerName: "John",
					Telephone:    "1234567890",
					Message: m.Message{
						Text: "hello",
					},
					Responses: []m.Response{},
				}
				mc.On("SendPromotionMessage", promoMsg).Return(nil)
				return &mc
			}(),
			expectedErr: errors.New("failed to update delivery status"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockPromotionRepo := tc.promoRepo

			mockStatsRepo := &d.MockStatsRepo{}

			mockMessageClient := tc.messageClient

			service := app.Service{
				PromotionRepo: mockPromotionRepo,
				StatsRepo:     mockStatsRepo,
				MessageClient: mockMessageClient,
			}

			err := service.StartPromotion(tc.customerID, tc.customerName, tc.telephone, tc.image)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestService_HandleResponse(t *testing.T) {
	testCases := []struct {
		name          string
		customerID    int
		promotionID   int
		responseID    int
		nextStep      int
		promoRepo     d.PromotionRepo
		statsRepo     d.StatsRepo
		messageClient m.MessageClient
		expectedErr   error
	}{
		{
			name:        "success with second next step",
			customerID:  1,
			promotionID: 2,
			nextStep:    2,
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    1,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps: map[int]d.Step{
						2: {
							ID:        2,
							Text:      "hello",
							Responses: []d.Response{},
						},
					},
				}
				mp.On("GetPromotionByID", promotion.ID).Return(promotion, nil)
				mp.On("UpdateCurrentStep", promotion.ID, 2).Return(nil)
				mp.On("UpdateDeliveryStatus", promotion.ID, "sent").Return(nil)
				return &mp
			}(),
			statsRepo: func() *d.MockStatsRepo {
				ms := d.MockStatsRepo{}
				ms.On("IncreaseFirstMsgResponses").Return(nil)
				ms.On("IncreaseYesClicks").Return(nil)
				return &ms
			}(),
			messageClient: func() *m.MockMessageClient {
				mc := m.MockMessageClient{}
				promoMsg := m.PromotionMessage{
					ID:           2,
					CustomerID:   1,
					CustomerName: "John",
					Telephone:    "1234567890",
					Message: m.Message{
						Text: "hello",
					},
					Responses: []m.Response{},
				}
				mc.On("SendPromotionMessage", promoMsg).Return(nil)
				return &mc
			}(),
			expectedErr: nil,
		},
		{
			name:        "success with third next step",
			customerID:  1,
			promotionID: 2,
			nextStep:    3,
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    1,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps: map[int]d.Step{
						3: {
							ID:        3,
							Text:      "hello",
							Responses: []d.Response{},
						},
					},
				}
				mp.On("GetPromotionByID", promotion.ID).Return(promotion, nil)
				mp.On("UpdateCurrentStep", promotion.ID, 3).Return(nil)
				mp.On("UpdateDeliveryStatus", promotion.ID, "sent").Return(nil)
				return &mp
			}(),
			statsRepo: func() *d.MockStatsRepo {
				ms := d.MockStatsRepo{}
				ms.On("IncreaseFirstMsgResponses").Return(nil)
				ms.On("IncreaseNoClicks").Return(nil)
				return &ms
			}(),
			messageClient: func() *m.MockMessageClient {
				mc := m.MockMessageClient{}
				promoMsg := m.PromotionMessage{
					ID:           2,
					CustomerID:   1,
					CustomerName: "John",
					Telephone:    "1234567890",
					Message: m.Message{
						Text: "hello",
					},
					Responses: []m.Response{},
				}
				mc.On("SendPromotionMessage", promoMsg).Return(nil)
				return &mc
			}(),
			expectedErr: nil,
		},
		{
			name:        "success with forth next step",
			customerID:  1,
			promotionID: 2,
			nextStep:    4,
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    2,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps: map[int]d.Step{
						4: {
							ID:        4,
							Text:      "hello",
							Responses: []d.Response{},
						},
					},
				}
				mp.On("GetPromotionByID", promotion.ID).Return(promotion, nil)
				mp.On("UpdateCurrentStep", promotion.ID, 4).Return(nil)
				mp.On("UpdateDeliveryStatus", promotion.ID, "sent").Return(nil)
				return &mp
			}(),
			statsRepo: func() *d.MockStatsRepo {
				ms := d.MockStatsRepo{}
				ms.On("IncreaseRevealClicks").Return(nil)
				return &ms
			}(),
			messageClient: func() *m.MockMessageClient {
				mc := m.MockMessageClient{}
				promoMsg := m.PromotionMessage{
					ID:           2,
					CustomerID:   1,
					CustomerName: "John",
					Telephone:    "1234567890",
					Message: m.Message{
						Text: "hello",
					},
					Responses: []m.Response{},
				}
				mc.On("SendPromotionMessage", promoMsg).Return(nil)
				return &mc
			}(),
			expectedErr: nil,
		},
		{
			name:        "success with finish",
			customerID:  1,
			promotionID: 2,
			nextStep:    0,
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    3,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps:          map[int]d.Step{},
				}
				mp.On("GetPromotionByID", promotion.ID).Return(promotion, nil)
				mp.On("UpdateCurrentStep", promotion.ID, 0).Return(nil)
				mp.On("UpdateFinished", promotion.ID).Return(nil)
				return &mp
			}(),
			statsRepo:     &d.MockStatsRepo{},
			messageClient: &m.MockMessageClient{},
			expectedErr:   nil,
		},
		{
			name:        "failed to fetch promotion",
			customerID:  1,
			promotionID: 2,
			nextStep:    0,
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				mp.On("GetPromotionByID", 2).Return(d.Promotion{}, errors.New("not found"))
				return &mp
			}(),
			statsRepo:     &d.MockStatsRepo{},
			messageClient: &m.MockMessageClient{},
			expectedErr:   errors.New("not found"),
		},
		{
			name:        "failed to update current step",
			customerID:  1,
			promotionID: 2,
			nextStep:    2,
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    3,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps:          map[int]d.Step{},
				}
				mp.On("GetPromotionByID", 2).Return(promotion, nil)
				mp.On("UpdateCurrentStep", 2, 2).Return(errors.New("error"))
				return &mp
			}(),
			statsRepo:     &d.MockStatsRepo{},
			messageClient: &m.MockMessageClient{},
			expectedErr:   errors.New("error"),
		},
		{
			name:        "failed to update finished",
			customerID:  1,
			promotionID: 2,
			nextStep:    0,
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    3,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps:          map[int]d.Step{},
				}
				mp.On("GetPromotionByID", 2).Return(promotion, nil)
				mp.On("UpdateCurrentStep", 2, 0).Return(nil)
				mp.On("UpdateFinished", 2).Return(errors.New("error"))
				return &mp
			}(),
			statsRepo: func() *d.MockStatsRepo {
				ms := d.MockStatsRepo{}
				ms.On("IncreaseFirstMsgResponses").Return(nil)
				ms.On("IncreaseNoClicks").Return(nil)
				return &ms
			}(),
			messageClient: &m.MockMessageClient{},
			expectedErr:   errors.New("error"),
		},
		{
			name:        "failed to find step",
			customerID:  1,
			promotionID: 2,
			nextStep:    4,
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    2,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps: map[int]d.Step{
						3: {
							ID:        1,
							Text:      "hello",
							Responses: []d.Response{},
						},
					},
				}
				mp.On("GetPromotionByID", 2).Return(promotion, nil)
				mp.On("UpdateCurrentStep", 2, 4).Return(nil)
				return &mp
			}(),
			statsRepo: func() *d.MockStatsRepo {
				ms := d.MockStatsRepo{}
				ms.On("IncreaseRevealClicks").Return(nil)
				return &ms
			}(),
			messageClient: &m.MockMessageClient{},
			expectedErr:   errors.New("step 4 not found"),
		},
		{
			name:        "failed to update delivery status",
			customerID:  1,
			promotionID: 2,
			nextStep:    2,
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    1,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps: map[int]d.Step{
						2: {
							ID:        1,
							Text:      "hello",
							Responses: []d.Response{},
						},
					},
				}
				mp.On("GetPromotionByID", 2).Return(promotion, nil)
				mp.On("UpdateCurrentStep", 2, 2).Return(nil)
				mp.On("UpdateDeliveryStatus", 2, "sent").Return(errors.New("failed to update delivery status"))
				return &mp
			}(),
			statsRepo: func() *d.MockStatsRepo {
				ms := d.MockStatsRepo{}
				ms.On("IncreaseFirstMsgResponses").Return(nil)
				ms.On("IncreaseYesClicks").Return(nil)
				return &ms
			}(),
			messageClient: func() *m.MockMessageClient {
				mc := m.MockMessageClient{}
				promoMsg := m.PromotionMessage{
					ID:           2,
					CustomerID:   1,
					CustomerName: "John",
					Telephone:    "1234567890",
					Message: m.Message{
						Text: "hello",
					},
					Responses: []m.Response{},
				}
				mc.On("SendPromotionMessage", promoMsg).Return(nil)
				return &mc
			}(),
			expectedErr: errors.New("failed to update delivery status"),
		},
		{
			name:        "failed to send message",
			customerID:  1,
			promotionID: 2,
			nextStep:    2,
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             2,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    1,
					DeliveryStatus: "pending",
					Finished:       false,
					Steps: map[int]d.Step{
						2: {
							ID:        1,
							Text:      "hello",
							Responses: []d.Response{},
						},
					},
				}
				mp.On("GetPromotionByID", 2).Return(promotion, nil)
				mp.On("UpdateCurrentStep", 2, 2).Return(nil)
				return &mp
			}(),
			statsRepo: func() *d.MockStatsRepo {
				ms := d.MockStatsRepo{}
				ms.On("IncreaseFirstMsgResponses").Return(nil)
				ms.On("IncreaseYesClicks").Return(nil)
				return &ms
			}(),
			messageClient: func() *m.MockMessageClient {
				mc := m.MockMessageClient{}
				promoMsg := m.PromotionMessage{
					ID:           2,
					CustomerID:   1,
					CustomerName: "John",
					Telephone:    "1234567890",
					Message: m.Message{
						Text: "hello",
					},
					Responses: []m.Response{},
				}
				mc.On("SendPromotionMessage", promoMsg).Return(errors.New("failed to send messsage"))
				return &mc
			}(),
			expectedErr: errors.New("failed to send messsage"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			service := app.Service{
				PromotionRepo: tc.promoRepo,
				StatsRepo:     tc.statsRepo,
				MessageClient: tc.messageClient,
			}

			err := service.HandleResponse(tc.customerID, tc.promotionID, tc.nextStep)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestService_HandleDeliveryStatus(t *testing.T) {
	testCases := []struct {
		name           string
		promotionID    int
		step           int
		deliveryStatus string
		promoRepo      d.PromotionRepo
		statsRepo      d.StatsRepo
		expectedErr    error
	}{
		{
			name:           "success",
			promotionID:    1,
			step:           3,
			deliveryStatus: "read",
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             1,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    3,
					DeliveryStatus: "delivered",
					Finished:       false,
					Steps:          map[int]d.Step{},
				}
				mp.On("GetPromotionByID", 1).Return(promotion, nil)
				mp.On("UpdateDeliveryStatus", 1, "read").Return(nil)
				return &mp
			}(),
			statsRepo:   new(d.MockStatsRepo),
			expectedErr: nil,
		},
		{
			name:           "success with read stat update",
			promotionID:    1,
			step:           2,
			deliveryStatus: "read",
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             1,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    2,
					DeliveryStatus: "delivered",
					Finished:       false,
					Steps:          map[int]d.Step{},
				}
				mp.On("GetPromotionByID", 1).Return(promotion, nil)
				mp.On("UpdateDeliveryStatus", 1, "read").Return(nil)
				return &mp
			}(),
			statsRepo: func() *d.MockStatsRepo {
				ms := d.MockStatsRepo{}
				ms.On("IncreaseSecondMsgReads").Return(nil)
				return &ms
			}(),
			expectedErr: nil,
		},
		{
			name:           "failed to fetch promotion",
			promotionID:    1,
			step:           2,
			deliveryStatus: "read",
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             1,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    2,
					DeliveryStatus: "delivered",
					Finished:       false,
					Steps:          map[int]d.Step{},
				}
				mp.On("GetPromotionByID", 1).Return(promotion, errors.New("failed to fetch promotion"))
				return &mp
			}(),
			statsRepo:   new(d.MockStatsRepo),
			expectedErr: errors.New("failed to fetch promotion"),
		},
		{
			name:           "failed on step mismatch",
			promotionID:    1,
			step:           2,
			deliveryStatus: "read",
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             1,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    3,
					DeliveryStatus: "delivered",
					Finished:       false,
					Steps:          map[int]d.Step{},
				}
				mp.On("GetPromotionByID", 1).Return(promotion, nil)
				return &mp
			}(),
			statsRepo:   new(d.MockStatsRepo),
			expectedErr: errors.New("step 2 is not the current step"),
		},
		{
			name:           "failed to update delivery status",
			promotionID:    1,
			step:           2,
			deliveryStatus: "read",
			promoRepo: func() *d.MockPromotionRepo {
				mp := d.MockPromotionRepo{}
				promotion := d.Promotion{
					ID:             1,
					CustomerID:     1,
					CustomerName:   "John",
					Telephone:      "1234567890",
					CurrentStep:    2,
					DeliveryStatus: "delivered",
					Finished:       false,
					Steps:          map[int]d.Step{},
				}
				mp.On("GetPromotionByID", 1).Return(promotion, nil)
				mp.On("UpdateDeliveryStatus", 1, "read").Return(errors.New("failed to update delivery status"))
				return &mp
			}(),
			statsRepo: func() *d.MockStatsRepo {
				ms := d.MockStatsRepo{}
				ms.On("IncreaseSecondMsgReads").Return(nil)
				return &ms
			}(),
			expectedErr: errors.New("failed to update delivery status"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockPromoRepo := tc.promoRepo

			mockStatsRepo := tc.statsRepo

			mockMessageClient := &m.MockMessageClient{}

			service := app.Service{
				PromotionRepo: mockPromoRepo,
				StatsRepo:     mockStatsRepo,
				MessageClient: mockMessageClient,
			}

			err := service.HandleDeliveryStatus(tc.promotionID, tc.step, tc.deliveryStatus)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestService_GetStats(t *testing.T) {
	testCases := []struct {
		name          string
		statsRepo     d.StatsRepo
		expectedStats d.Stats
		expectedErr   error
	}{
		{
			name: "success",
			statsRepo: func() *d.MockStatsRepo {
				ms := d.MockStatsRepo{}
				stats := d.Stats{
					FirstMsgResponses: 10,
					SecondMsgReads:    15,
					RevealClicks:      24,
					YesClicks:         22,
					NoClicks:          64,
				}
				ms.On("GetStats").Return(stats, nil)
				return &ms
			}(),
			expectedStats: d.Stats{
				FirstMsgResponses: 10,
				SecondMsgReads:    15,
				RevealClicks:      24,
				YesClicks:         22,
				NoClicks:          64,
			},
			expectedErr: nil,
		},
		{
			name: "get stats error",
			statsRepo: func() *d.MockStatsRepo {
				ms := d.MockStatsRepo{}
				ms.On("GetStats").Return(d.Stats{}, errors.New("failed"))
				return &ms
			}(),
			expectedStats: d.Stats{},
			expectedErr:   errors.New("failed"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockPromoRepo := &d.MockPromotionRepo{}

			mockMessageClient := &m.MockMessageClient{}

			service := app.Service{
				PromotionRepo: mockPromoRepo,
				StatsRepo:     tc.statsRepo,
				MessageClient: mockMessageClient,
			}

			stats, err := service.GetStats()

			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedStats, stats)
		})
	}
}
