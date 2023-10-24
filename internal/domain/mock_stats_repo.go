package domain

import (
	"github.com/stretchr/testify/mock"
)

//MockStatsRepo mocks the stats repository
type MockStatsRepo struct {
	mock.Mock
}
//GetStats method mock
func (m *MockStatsRepo) GetStats() (Stats, error) {
	args := m.Called()
	return args.Get(0).(Stats), args.Error(1)
}
//IncreaseFirstMsgResponses method mock
func (m *MockStatsRepo) IncreaseFirstMsgResponses(){
	m.Called()
}
//IncreaseSecondMsgReads method mock
func (m *MockStatsRepo) IncreaseSecondMsgReads() {
	m.Called()
}
//IncreaseRevealClicks method mock
func (m *MockStatsRepo) IncreaseRevealClicks() {
	m.Called()
}
//IncreaseYesClicks method mock
func (m *MockStatsRepo) IncreaseYesClicks() {
	m.Called()
}
//IncreaseNoClicks method mock
func (m *MockStatsRepo) IncreaseNoClicks() {
	m.Called()
}