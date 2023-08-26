package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ryanadiputraa/api-udrio/internal/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/mail"
	"github.com/ryanadiputraa/api-udrio/pkg/pagination"
)

type orderUsecase struct {
	repository domain.OrderRepository
}

func NewOrderUsecase(repository domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{repository: repository}
}

func (u *orderUsecase) GetUserOrders(ctx context.Context, userID string, size int, page int) (order []domain.OrderDTO, meta pagination.Page, err error) {
	if size <= 0 {
		size = 20
	}
	if page <= 0 {
		page = 1
	}

	offset := pagination.Offset(size, page)
	order, count, err := u.repository.FetchOrdersByUserID(ctx, userID, size, offset)
	if order == nil {
		order = []domain.OrderDTO{}
	}
	if err != nil {
		if err.Error() == "record not found" {
			return []domain.OrderDTO{}, pagination.Page{}, nil
		}
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

func (u *orderUsecase) CreateOrder(ctx context.Context, userID string, payload domain.OrderPayload) (err error) {
	var productIDs []string

	for _, v := range payload.Orders {
		if len(v.ProductID) == 0 {
			return errors.New("invalid param: missing product id")
		}
		if v.Quantity < 1 {
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

	user, err := u.repository.SaveOrder(ctx, order, payload.Orders, productIDs)
	if err != nil {
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
	}

	return
}
