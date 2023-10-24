package domain

import (
	"github.com/stretchr/testify/mock"
)

//MockPromotionRepo mocks the promotion repository
type MockPromotionRepo struct {
	mock.Mock
}
//CreatePromotion method mock
func (m *MockPromotionRepo) CreatePromotion(customerID int, customerName string, telephone string, image string, coupon int) (Promotion, error) {
	args := m.Called(customerID, customerName, telephone, image, coupon)
	return args.Get(0).(Promotion), args.Error(1)
}
//GetPromotionByID method mock
func (m *MockPromotionRepo) GetPromotionByID(id int) (Promotion, error) {
	args := m.Called(id)
	return args.Get(0).(Promotion), args.Error(1)
}
//UpdateDeliveryStatus method mock
func (m *MockPromotionRepo) UpdateDeliveryStatus(ID int, deliveryStatus string) error {
	args := m.Called(ID, deliveryStatus)
	return args.Error(0)
}
//UpdateCurrentStep method mock
func (m *MockPromotionRepo) UpdateCurrentStep(ID int, currentStep int) error {
	args := m.Called(ID, currentStep)
	return args.Error(0)
}
//UpdateFinished method mock
func (m *MockPromotionRepo) UpdateFinished(ID int) error {
	args := m.Called(ID)
	return args.Error(0)
}
//GetLastPromotionID method mock
func (m *MockPromotionRepo) GetLastPromotionID() int {
	args := m.Called()
	return args.Get(0).(int)
}
