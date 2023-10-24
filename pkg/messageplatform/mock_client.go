package messageplatform

import (
	"github.com/stretchr/testify/mock"
)

// MockMessageClient mocks the stats repository
type MockMessageClient struct {
	mock.Mock
}

// SendPromotionMessage method mock
func (m *MockMessageClient) SendPromotionMessage(msg PromotionMessage) error {
	args := m.Called(msg)
	return args.Error(0)
}
