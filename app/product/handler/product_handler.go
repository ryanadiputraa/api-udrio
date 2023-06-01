package handler

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/url"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/pagination"
	log "github.com/sirupsen/logrus"
)

type productHandler struct {
	productRepository domain.IProductRepository
}

func NewProductHandler(repository domain.IProductRepository) domain.IProductHandler {
	return &productHandler{productRepository: repository}
}

func (h *productHandler) GetProductList(ctx context.Context, size int, page int, category int, query string) (products []domain.Product, meta pagination.Page, err error) {
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
	products, count, err := h.productRepository.Fetch(ctx, size, offset, category, decodedQuery)
	if err != nil {
		log.Error("fail to fetch products: ", err.Error())
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

func (h *productHandler) GetProductDetail(ctx context.Context, productID string) (domain.Product, error) {
	product, err := h.productRepository.FindByID(ctx, productID)
	if err != nil {
		log.Error("fail to find product: ", err.Error())
		return product, err
	}
	return product, nil
}

func (h *productHandler) GetProductCategoryList(ctx context.Context) ([]domain.ProductCategory, error) {
	categories, err := h.productRepository.FetchCategory(ctx)
	if err != nil {
		log.Error("fail to fetch categories: ", err.Error())
		return nil, err
	}
	return categories, nil
}

func (h *productHandler) UploadProductImage(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (err error) {
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, file); err != nil {
		log.Error("fail to copy image buffer:", err.Error())
		return
	}

	url, err := h.productRepository.UploadImage(ctx, buf.Bytes(), fileHeader.Filename)
	if err != nil {
		log.Error("fail to save image to firebase: ", err.Error())
	}
	log.Error(url)
	return
}
