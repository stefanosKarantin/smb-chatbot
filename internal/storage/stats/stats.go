package stats

import (
	"log"

	d "github.com/stefanosKarantin/smb-chatbot/internal/domain"
)

type StatsStorage struct {
	Stats d.Stats
}

func NewStatsStorage() StatsStorage {
	return StatsStorage{
		Stats: d.Stats{},
	}
}

func (s *StatsStorage) GetStats() (d.Stats, error) {
	return s.Stats, nil
}

func (s *StatsStorage) IncreaseFirstMsgResponses() {
	s.Stats.FirstMsgResponses++
	log.Println("Increased first message responses")
}

func (s *StatsStorage) IncreaseSecondMsgReads() {
	s.Stats.SecondMsgReads++
	log.Println("Increased second message reads")
}

func (s *StatsStorage) IncreaseRevealClicks() {
	s.Stats.RevealClicks++
	log.Println("Increased reveal clicks")
}

func (s *StatsStorage) IncreaseYesClicks() {
	s.Stats.YesClicks++
	log.Println("Increased yes clicks")
}

func (s *StatsStorage) IncreaseNoClicks() {
	s.Stats.NoClicks++
	log.Println("Increased no clicks")
}
