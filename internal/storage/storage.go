package storage

import (
	p "github.com/stefanosKarantin/smb-chatbot/internal/storage/promotion"
	s "github.com/stefanosKarantin/smb-chatbot/internal/storage/stats"
)

type Storage struct {
	PromotionStorage p.PromotionStorage
	StatsStorage s.StatsStorage
}

func NewStorage() Storage {
	return Storage{
		PromotionStorage: p.NewPromotionStorage(),
		StatsStorage: s.NewStatsStorage(),
	}
}

