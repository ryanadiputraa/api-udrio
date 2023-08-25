package handler

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrUpdateIfExist(t *testing.T) {
	cases := []struct {
		description       string
		err               error
		mockRepoBehaviour func(mockUserRepo *mocks.IUserRepository, mockCartRepo *mocks.ICartRepository)
	}{
		{
			description: "should create or update user data & user cart if exist",
			err:         nil,
			mockRepoBehaviour: func(mockUserRepo *mocks.IUserRepository, mockCartRepo *mocks.ICartRepository) {
				mockUserRepo.On("SaveOrUpdate", mock.Anything, mock.Anything).Return(nil)
				mockCartRepo.On("CreateOrUpdate", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			description: "should return err when fail to save user data",
			err:         errors.New("fail to save user data"),
			mockRepoBehaviour: func(mockUserRepo *mocks.IUserRepository, mockCartRepo *mocks.ICartRepository) {
				mockUserRepo.On("SaveOrUpdate", mock.Anything, mock.Anything).Return(errors.New("fail to save user data"))
				mockCartRepo.On("CreateOrUpdate", mock.Anything, mock.Anything).Return(nil)
			},
		},
		{
			description: "should return err when fail to create user cart",
			err:         errors.New("fail to create user cart"),
			mockRepoBehaviour: func(mockUserRepo *mocks.IUserRepository, mockCartRepo *mocks.ICartRepository) {
				mockUserRepo.On("SaveOrUpdate", mock.Anything, mock.Anything).Return(errors.New("fail to create user cart"))
				mockCartRepo.On("CreateOrUpdate", mock.Anything, mock.Anything).Return(nil)
			},
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			userRepo := new(mocks.IUserRepository)
			cartRepo := new(mocks.ICartRepository)
			c.mockRepoBehaviour(userRepo, cartRepo)

			currTime := time.Now()
			handler := NewUserUsecase(userRepo, cartRepo)
			err := handler.CreateOrUpdateIfExist(context.TODO(), domain.User{
				ID:        "1",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john@mail.com",
				Picture:   "https://domain.com/image.png",
				Locale:    "id",
				CreatedAt: currTime,
				UpdatedAt: currTime,
			})

			assert.Equal(t, c.err, err)
		})
	}
}

func TestGetUserInfo(t *testing.T) {
	currTime := time.Now()

	cases := []struct {
		description       string
		expected          domain.User
		err               error
		mockRepoBehaviour func(mockUserRepo *mocks.IUserRepository, mockCartRepo *mocks.ICartRepository)
	}{
		{
			description: "should return user data",
			expected: domain.User{
				ID:        "1",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john@mail.com",
				Picture:   "https://domain.com/image.png",
				Locale:    "id",
				CreatedAt: currTime,
				UpdatedAt: currTime,
			},
			err: nil,
			mockRepoBehaviour: func(mockUserRepo *mocks.IUserRepository, mockCartRepo *mocks.ICartRepository) {
				mockUserRepo.On("FindByID", mock.Anything, "1").Return(domain.User{
					ID:        "1",
					FirstName: "John",
					LastName:  "Doe",
					Email:     "john@mail.com",
					Picture:   "https://domain.com/image.png",
					Locale:    "id",
					CreatedAt: currTime,
					UpdatedAt: currTime,
				}, nil)
			},
		},
		{
			description: "should return error when fail to retrieve data",
			expected: domain.User{
				ID:        "",
				FirstName: "",
				LastName:  "",
				Email:     "",
				Picture:   "",
				Locale:    "",
				CreatedAt: currTime,
				UpdatedAt: currTime,
			},
			err: errors.New("no record found"),
			mockRepoBehaviour: func(mockUserRepo *mocks.IUserRepository, mockCartRepo *mocks.ICartRepository) {
				mockUserRepo.On("FindByID", mock.Anything, mock.Anything).Return(domain.User{
					ID:        "",
					FirstName: "",
					LastName:  "",
					Email:     "",
					Picture:   "",
					Locale:    "",
					CreatedAt: currTime,
					UpdatedAt: currTime,
				}, errors.New("no record found"))
			},
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			userRepo := new(mocks.IUserRepository)
			cartRepo := new(mocks.ICartRepository)
			c.mockRepoBehaviour(userRepo, cartRepo)

			handler := NewUserUsecase(userRepo, cartRepo)
			user, err := handler.GetUserInfo(context.TODO(), "1")
			assert.Equal(t, c.expected, user)
			assert.Equal(t, c.err, err)
		})
	}
}
