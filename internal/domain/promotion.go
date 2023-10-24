package domain

type Promotion struct {
	ID            int
	CustomerID    int
	CustomerName  string
	DeliveryStatus string
	Telephone     string
	Finished      bool
	CurrentStep int
	Steps       map[int]Step
}

type Step struct {
	ID         int
	StepNumber int
	Text       string
	Responses  []Response
	Image      string
}

type Response struct {
	ID       int
	Text     string
	NextStep int
}

type PromotionRepo interface {
	CreatePromotion(customerID int, customerName string, telephone string, image string) (Promotion, error)
	GetPromotionByID(ID int) (Promotion, error)
	UpdateDeliveryStatus(ID int, deliveryStatus string) error
	UpdateCurrentStep(ID int, currentStep int) error
	UpdateFinished(ID int) error
	GetLastPromotionID() int
}

