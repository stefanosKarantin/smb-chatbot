package domain

type Stats struct {
	FirstMsgResponses int
	SecondMsgReads    int
	RevealClicks      int
	YesClicks         int
	NoClicks          int
}

type StatsRepo interface {
	GetStats() (Stats, error)
	IncreaseFirstMsgResponses()
	IncreaseSecondMsgReads()
	IncreaseRevealClicks()
	IncreaseYesClicks()
	IncreaseNoClicks()
}