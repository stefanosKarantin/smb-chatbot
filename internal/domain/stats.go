package domain

// Stats represents the statistics of the chatbot's interactions with the users.
type Stats struct {
	FirstMsgResponses int // Number of times a user responded to the first message.
	SecondMsgReads    int // Number of times a user read the second message.
	RevealClicks      int // Number of times a user clicked the reveal button.
	YesClicks         int // Number of times a user clicked the yes button.
	NoClicks          int // Number of times a user clicked the no button.
}

// StatsRepo represents the repository for the chatbot's statistics.
type StatsRepo interface {
	GetStats() (Stats, error)
	IncreaseFirstMsgResponses() 
	IncreaseSecondMsgReads()
	IncreaseRevealClicks()
	IncreaseYesClicks()
	IncreaseNoClicks()
}