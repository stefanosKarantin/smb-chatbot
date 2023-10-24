package domain

// Promotion represents the promotion entity with its properties.
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

// Step represents the step entity with its properties.
type Step struct {
	ID         int
	StepNumber int
	Text       string
	Responses  []Response
	Image      string
}

// Response represents a response entity with its properties.
type Response struct {
	ID       int
	Text     string
	NextStep int
}

// PromotionRepo is an interface for interacting with the promotion repository.
type PromotionRepo interface {
	CreatePromotion(customerID int, customerName string, telephone string, image string, coupon int) (Promotion, error)
	GetPromotionByID(ID int) (Promotion, error)
	UpdateDeliveryStatus(ID int, deliveryStatus string) error
	UpdateCurrentStep(ID int, currentStep int) error
	UpdateFinished(ID int) error
	GetLastPromotionID() int
}

