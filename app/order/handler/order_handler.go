package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/mail"
	"github.com/ryanadiputraa/api-udrio/pkg/pagination"
	log "github.com/sirupsen/logrus"
)

type orderHandler struct {
	repository domain.IOrderRepository
}

func NewOrderHandler(repository domain.IOrderRepository) domain.IOrderHandler {
	return &orderHandler{repository: repository}
}

func (h *orderHandler) GetUserOrders(ctx context.Context, userID string, size int, page int) (order []domain.OrderDTO, meta pagination.Page, err error) {
	if size <= 0 {
		size = 20
	}
	if page <= 0 {
		page = 1
	}

	order, count, err := h.repository.FetchOrdersByUserID(ctx, userID, size, page)
	if order == nil {
		order = []domain.OrderDTO{}
	}
	if err != nil {
		if err.Error() == "record not found" {
			return []domain.OrderDTO{}, pagination.Page{}, nil
		}
		log.Error("fail to fetch user orders: ", err.Error())
	}
	pages := *pagination.NewPagination(size, page, int(count))
	meta = pagination.Page{
		CurrentPage:  pages.Page,
		TotalPage:    pages.TotalPage,
		TotalData:    pages.TotalData,
		NextPage:     pages.NextPage(),
		PreviousPage: pages.PreviousPage(),
	}

	return
}

func (h *orderHandler) CreateOrder(ctx context.Context, userID string, payload domain.OrderPayload) (err error) {
	var productIDs []string

	for _, v := range payload.Orders {
		if len(v.ProductID) == 0 {
			log.Error("invalid param: missing product id")
			return errors.New("invalid param: missing product id")
		}
		if v.Quantity < 1 {
			log.Error("invalid param: missing quantity or must be greater than 0 ")
			return errors.New("invalid param: missing quantity or must be greater than 0 ")
		}
		productIDs = append(productIDs, v.ProductID)
	}

	order := domain.Order{
		UserID:    userID,
		SubTotal:  0,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	user, err := h.repository.SaveOrder(ctx, order, payload.Orders, productIDs)
	if err != nil {
		log.Error("fail to create order: ", err.Error())
		return
	}

	// send notification mail
	mailBody := fmt.Sprintf(
		"Hi, %s!\nTerima kasih telah membuat pesanan, namun mohon maaf website ini masih dalam status prototype.",
		user.FirstName,
	)
	err = mail.SendMail("Pesanan UD Rio Digital Printing", mailBody, []string{user.Email})
	// mail error can be ignored
	if err != nil {
		log.Error("fail to send notification mail: ", err.Error())
	}

	return
}
