package app

import (
	"fmt"
	"net/http"

	d "github.com/stefanosKarantin/smb-chatbot/internal/domain"
	m "github.com/stefanosKarantin/smb-chatbot/pkg/messageplatform"
)

type Service struct {
	PromotionRepo d.PromotionRepo
	StatsRepo     d.StatsRepo
	MessageClient m.MessageClient
}

func NewService(
	promotionRepo d.PromotionRepo,
	statsRepo d.StatsRepo,
	client http.Client,
	host string,
) Service {
	return Service{
		PromotionRepo: promotionRepo,
		StatsRepo:     statsRepo,
		MessageClient: m.NewMessageClient(host, &client),
	}
}

func (s *Service) StartPromotion(
	customerID int,
	customerName string,
	telephone string,
	image string,
) error {
	promotion, err := s.PromotionRepo.CreatePromotion(customerID, customerName, telephone, image)
	if err != nil {
		return err
	}

	responses := []m.Response{}
	for _, rs := range promotion.Steps[1].Responses {
		r := m.Response{
			ID:   rs.ID,
			Text: rs.Text,
			Url:  fmt.Sprintf("/response/%d", rs.NextStep),
		}
		responses = append(responses, r)
	}
	msg := m.PromotionMessage{
		ID:           promotion.ID,
		CustomerID:   customerID,
		CustomerName: customerName,
		Telephone:    telephone,
		Message: m.Message{
			Text: promotion.Steps[1].Text,
		},
		Responses: responses,
	}
	err = s.MessageClient.SendPromotionMessage(msg)
	if err != nil {
		return err
	}
	return s.PromotionRepo.UpdateDeliveryStatus(promotion.ID, "sent")
}

func (s *Service) HandleResponse(customerID int, promotionID int, nextStep int) error {
	promotion, err := s.PromotionRepo.GetPromotionByID(promotionID)
	if err != nil {
		return err
	}

	err = s.PromotionRepo.UpdateCurrentStep(promotionID, nextStep)
	if err != nil {
		return err
	}
	if nextStep == 2 || nextStep == 3 {
		s.StatsRepo.IncreaseFirstMsgResponses()
	}
	if nextStep == 2 {
		s.StatsRepo.IncreaseYesClicks()
	}
	if nextStep == 3 {
		s.StatsRepo.IncreaseNoClicks()
	}
	if nextStep == 4 {
		s.StatsRepo.IncreaseRevealClicks()
	}
	if nextStep == 0  {
		err = s.PromotionRepo.UpdateFinished(promotionID)
		if err != nil {
			return err
		}
		return nil
	}

	if _, ok := promotion.Steps[nextStep]; !ok {
		return fmt.Errorf("step %d not found", nextStep)
	}
	responses := []m.Response{}
	for _, rs := range promotion.Steps[nextStep].Responses {
		r := m.Response{
			ID:   rs.ID,
			Text: rs.Text,
			Url:  fmt.Sprintf("/response/%d", rs.NextStep),
		}
		responses = append(responses, r)
	}

	msg := m.PromotionMessage{
		ID:           promotionID,
		CustomerID:   promotion.CustomerID,
		CustomerName: promotion.CustomerName,
		Telephone:    promotion.Telephone,
		Message: m.Message{
			Text: promotion.Steps[nextStep].Text,
		},
		Responses: responses,
	}
	err = s.MessageClient.SendPromotionMessage(msg)
	if err != nil {
		return err
	}
	return s.PromotionRepo.UpdateDeliveryStatus(promotion.ID, "sent")
}

func (s *Service) HandleDeliveryStatus(promotionID int, step int, deliveryStatus string) error {
	promotion, err := s.PromotionRepo.GetPromotionByID(promotionID)
	if err != nil {
		return err
	}

	if promotion.CurrentStep != step {
		return fmt.Errorf("step %d is not the current step", step)
	}
	
	if deliveryStatus == "read" && step == 2 {
		s.StatsRepo.IncreaseSecondMsgReads()
	}
	return s.PromotionRepo.UpdateDeliveryStatus(promotionID, deliveryStatus)
	
}	
func (s *Service) GetStats() (d.Stats, error) {
	return s.StatsRepo.GetStats()
}
