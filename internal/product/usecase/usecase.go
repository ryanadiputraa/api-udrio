package usecase

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/url"

	"github.com/google/uuid"
	"github.com/ryanadiputraa/api-udrio/config"
	"github.com/ryanadiputraa/api-udrio/internal/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/pagination"
)

type usecase struct {
	conf       config.Config
	repository domain.ProductRepository
}

func NewProductUsecase(repository domain.ProductRepository) domain.ProductUsecase {
	return &usecase{repository: repository}
}

func (u *usecase) GetProductList(ctx context.Context, size int, page int, category int, query string) (products []domain.Product, meta pagination.Page, err error) {
	if size <= 0 {
		size = 20
	}
	if page <= 0 {
		page = 1
	}

	decodedQuery, err := url.QueryUnescape(query)
	if err != nil {
		return nil, pagination.Page{}, err
	}

	offset := pagination.Offset(size, page)
	products, count, err := u.repository.Fetch(ctx, size, offset, category, decodedQuery)
	if err != nil {
		return nil, pagination.Page{}, err
	}

	pages := *pagination.NewPagination(size, page, int(count))
	meta = pagination.Page{
		CurrentPage:  pages.Page,
		TotalPage:    pages.TotalPage,
		TotalData:    pages.TotalData,
		NextPage:     pages.NextPage(),
		PreviousPage: pages.PreviousPage(),
	}

	return products, meta, nil
}

func (u *usecase) GetProductDetail(ctx context.Context, productID string) (domain.Product, error) {
	product, err := u.repository.FindByID(ctx, productID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (u *usecase) GetProductCategoryList(ctx context.Context) ([]domain.ProductCategory, error) {
	categories, err := u.repository.FetchCategory(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (u *usecase) UploadProductImage(ctx context.Context, productID string, file multipart.File) (err error) {
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, file); err != nil {
		return
	}

	imageID := uuid.NewString()
	bucketName := u.conf.Firebase.Bucket
	storageToken := u.conf.Firebase.StorageToken
	url := "https://firebasestorage.googleapis.com/v0/b/" + bucketName + "/o/" + imageID + "?alt=media&token=" + storageToken

	image := domain.ProductImage{
		ID:        imageID,
		Image:     url,
		ProductID: productID,
	}

	err = u.repository.SaveImage(ctx, buf.Bytes(), image)
	return
}
