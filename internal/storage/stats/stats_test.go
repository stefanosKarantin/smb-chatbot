package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"

	d "github.com/stefanosKarantin/smb-chatbot/internal/domain"
)

func TestStatsRepo(t *testing.T) {
	statsStorage := NewStatsStorage()

	// Test GetStats
	stats, err := statsStorage.GetStats()
	assert.NoError(t, err)
	assert.Equal(t, d.Stats{}, stats)

	// Test IncreaseFirstMsgResponses
	statsStorage.IncreaseFirstMsgResponses()
	stats, err = statsStorage.GetStats()
	assert.NoError(t, err)
	assert.Equal(t, d.Stats{FirstMsgResponses: 1}, stats)

	// Test IncreaseSecondMsgReads
	statsStorage.IncreaseSecondMsgReads()
	stats, err = statsStorage.GetStats()
	assert.NoError(t, err)
	assert.Equal(t, d.Stats{FirstMsgResponses: 1, SecondMsgReads: 1}, stats)

	// Test IncreaseRevealClicks
	statsStorage.IncreaseRevealClicks()
	stats, err = statsStorage.GetStats()
	assert.NoError(t, err)
	assert.Equal(t, d.Stats{FirstMsgResponses: 1, SecondMsgReads: 1, RevealClicks: 1}, stats)

	// Test IncreaseYesClicks
	statsStorage.IncreaseYesClicks()
	stats, err = statsStorage.GetStats()
	assert.NoError(t, err)
	assert.Equal(t, d.Stats{FirstMsgResponses: 1, SecondMsgReads: 1, RevealClicks: 1, YesClicks: 1}, stats)

	// Test IncreaseNoClicks
	statsStorage.IncreaseNoClicks()
	stats, err = statsStorage.GetStats()
	assert.NoError(t, err)
	assert.Equal(t, d.Stats{FirstMsgResponses: 1, SecondMsgReads: 1, RevealClicks: 1, YesClicks: 1, NoClicks: 1}, stats)
}
