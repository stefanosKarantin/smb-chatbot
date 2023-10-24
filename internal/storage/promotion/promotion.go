package promotion

import (
	"errors"
	"log"

	d "github.com/stefanosKarantin/smb-chatbot/internal/domain"
)

type PromotionStorage struct {
	Promotions map[int]d.Promotion
}

func NewPromotionStorage() PromotionStorage {
	return PromotionStorage{
		Promotions: make(map[int]d.Promotion),
	}
}

func (s *PromotionStorage) CreatePromotion(
	customerID int,
	customerName string,
	telephone string,
	image string,
) (d.Promotion, error) {
	ID := len(s.Promotions) + 1
	promotionSteps := make(map[int]d.Step)
	for _, step := range NewPromotionFlowSteps(customerName, image) {
		promotionSteps[step.ID] = step
	}
	promotion := d.Promotion{
		ID:             customerID,
		CustomerID:     customerID,
		CustomerName:   customerName,
		Telephone:      telephone,
		CurrentStep:    1,
		DeliveryStatus: "pending",
		Finished:       false,
		Steps:          promotionSteps,
	}
	s.Promotions[ID] = promotion
	log.Println("Created new promotion with ID: ", ID)

	return promotion, nil
}

func (s *PromotionStorage) GetPromotionByID(ID int) (d.Promotion, error) {
	if promotion, ok := s.Promotions[ID]; ok {
		return promotion, nil
	}
	return d.Promotion{}, errors.New("promotion not found")
}

func (s *PromotionStorage) UpdateDeliveryStatus(ID int, deliveryStatus string) error {
	if promotion, ok := s.Promotions[ID]; ok {
		promotion.DeliveryStatus = deliveryStatus
		log.Printf("Updated delivery status to %s for promotion with ID: %d", deliveryStatus, ID)
		return nil
	}
	return errors.New("promotion not found")
}

func (s *PromotionStorage) UpdateCurrentStep(ID int, currentStep int) error {
	if promotion, ok := s.Promotions[ID]; ok {
		promotion.CurrentStep = currentStep
		log.Printf("Updated current step to %d for promotion with ID: %d", currentStep, ID)
		return nil
	}
	return errors.New("promotion not found")
}

func (s *PromotionStorage) UpdateFinished(ID int) error {
	if promotion, ok := s.Promotions[ID]; ok {
		promotion.Finished = true
		log.Printf("Updated finished to true for promotion with ID: %d", ID)
		return nil
	}
	return errors.New("promotion not found")
}

func (s *PromotionStorage) GetLastPromotionID() int {
	return len(s.Promotions)
}
