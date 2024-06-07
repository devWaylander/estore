package product

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	internalErrors "route256.ozon.ru/project/cart/internal/errors"
	"route256.ozon.ru/project/cart/internal/model"
)

const (
	getProductUrl = "http://route256.pavl.uk:8080/get_product"
	tokenProductS = "testtoken"
)

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) GetProduct(ctx context.Context, SKU uint32, count uint16) (*model.Good, error) {
	if SKU == 0 {
		return nil, errors.New(internalErrors.ErrBadSKU)
	}

	product := model.Product{
		Token: tokenProductS,
		SKU:   SKU,
	}

	data, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, "POST", getProductUrl, reader)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("url: %s | create request %w", getProductUrl, err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("url: %s | do request: %w", getProductUrl, err)
	}
	defer resp.Body.Close()

	pResp := model.ProductRespV1{}
	if err := json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
		return nil, err
	}

	if pResp.Code == 16 {
		return nil, errors.New(internalErrors.ErrTokenNotValid)
	}
	if pResp.Code == 5 {
		return nil, errors.New(internalErrors.ErrSKUNotFound)
	}

	good := model.Good{
		SkuID: SKU,
		Name:  pResp.Name,
		Price: pResp.Price,
		Count: count,
	}

	return &good, err
}
