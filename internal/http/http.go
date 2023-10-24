package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/stefanosKarantin/smb-chatbot/internal/app"
)

type handler struct {
	Router chi.Router
	app    app.Service
}

func NewHandler(app app.Service) handler {
	return handler{
		Router: chi.NewRouter(),
		app:    app,
	}
}

func (h *handler) AppendRoutes() {
	h.Router.Post("/start-promotion", h.startPromotion)
	h.Router.Post("/response/{nextStep}", h.Response)
	h.Router.Post("/delivery-status", h.deliveryStatus)
	h.Router.Get("/stats", h.getStats)
}

func (h *handler) startPromotion(w http.ResponseWriter, r *http.Request) {
	var promotion promotion
	err := json.NewDecoder(r.Body).Decode(&promotion)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error decoding request body: %s", err.Error())
		return
	}

	err = h.app.StartPromotion(
		promotion.CustomerID,
		promotion.CustomerName,
		promotion.Telephone,
		promotion.Image,
		promotion.Coupon,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(errorResponse{Error: err.Error()})
		w.Write(res)
		log.Printf("Error starting promotion: %s", err.Error())
		return
	}
	return
}

func (h *handler) Response(w http.ResponseWriter, r *http.Request) {
	nextStep := chi.URLParam(r, "nextStep")
	ns, err := strconv.Atoi(nextStep)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error converting nextStep to int: %s", err.Error())
		return
	}

	var res customerResponse
	err = json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error decoding request body: %s", err.Error())
		return
	}

	err = h.app.HandleResponse(res.CustomerID, res.PromotionID, ns)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(errorResponse{Error: err.Error()})
		w.Write(res)
		log.Printf("Error handling response: %s", err.Error())
		return
	}
	return
}

func (h *handler) deliveryStatus(w http.ResponseWriter, r *http.Request) {
	var deliveryStatus deliveryStatusResponse
	err := json.NewDecoder(r.Body).Decode(&deliveryStatus)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error decoding request body: %s", err.Error())
		return
	}

	err = h.app.HandleDeliveryStatus(
		deliveryStatus.PromotionID,
		deliveryStatus.Step,
		deliveryStatus.DeliveryStatus,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(errorResponse{Error: err.Error()})
		w.Write(res)
		log.Printf("Error handling delivery status: %s", err.Error())
		return
	}
	return

}

func (h *handler) getStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.app.GetStats()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(errorResponse{Error: err.Error()})
		w.Write(res)
		log.Printf("Error getting stats: %s", err.Error())
		return
	}
	res, _ := json.Marshal(stats)
	w.Write(res)
	return
}

type promotion struct {
	CustomerID   int    `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Telephone    string `json:"telephone"`
	Image        string `json:"image"`
	Coupon       int    `json:"coupon"`
}

type customerResponse struct {
	CustomerID  int `json:"customer_id"`
	PromotionID int `json:"promotion_id"`
	ResponseID  int `json:"response_id"`
}

type deliveryStatusResponse struct {
	PromotionID    int    `json:"promotion_id"`
	DeliveryStatus string `json:"delivery_status"`
	Step           int    `json:"step"`
}

type errorResponse struct {
	Error string `json:"error"`
}
