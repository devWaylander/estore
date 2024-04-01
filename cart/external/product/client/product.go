package product

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	internalErrors "route256.ozon.ru/project/cart/internal/errors"
	"route256.ozon.ru/project/cart/internal/model"
)

const (
	productUrl    = "http://route256.pavl.uk:8080"
	tokenProductS = "testtoken"
)

type service struct {
}

func New() *service {
	return &service{}
}

func (s *service) GetProduct(ctx context.Context, SKU uint32) (*model.Good, error) {
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

	resp, err := http.Post(productUrl+"/get_product", "application/json", reader)
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

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
		Count: 0,
	}

	return &good, err
}
