package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	internalErrors "route256.ozon.ru/project/cart/internal/errors"
	"route256.ozon.ru/project/cart/internal/model"
)

const (
	userIDPath = "user_id"
	skuIDPath  = "sku_id"
)

type Service interface {
	AddToCart(ctx context.Context, userID, skuID int64, count uint16) error
	RemoveFromCart(ctx context.Context, userID, skuID int64) error
	CleanUpCart(ctx context.Context, userID int64) error
	GetCart(ctx context.Context, userID int64) (*model.Cart, error)
	Checkout(ctx context.Context, userID int64) (model.Order, error)
}

func Configure(ctx context.Context, mux *http.ServeMux, service Service) {
	// Добавить товар в корзину: POST /user/<user_id>/cart/<sku_id>
	mux.HandleFunc(fmt.Sprintf("POST /user/{%s}/cart/{%s}", userIDPath, skuIDPath), func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(r.PathValue(userIDPath), 10, 64)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrMissingUserID, http.StatusBadRequest)
			return
		}
		skuID, err := strconv.ParseInt(r.PathValue(skuIDPath), 10, 64)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrMissingSKU, http.StatusBadRequest)
			return
		}

		var body = AddGoodBody{}
		err = json.NewDecoder(r.Body).Decode(&body)
		if err != nil || body.Count <= 0 {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrCountIsMissingOrNotValid, http.StatusBadRequest)
			return
		}

		if err := service.AddToCart(ctx, int64(userID), int64(skuID), body.Count); err != nil {
			log.Println(err)
			if err.Error() == internalErrors.ErrSKUNotFound || err.Error() == internalErrors.ErrBadSKU {
				return
			}
			if err.Error() == internalErrors.ErrStockCount {
				http.Error(w, internalErrors.ErrStockCount, http.StatusPreconditionFailed)
				return
			}
			http.Error(w, internalErrors.ErrFailedAddAGood, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	// Удалить товар из корзины: DELETE /user/<user_id>/cart/<sku_id>
	mux.HandleFunc(fmt.Sprintf("DELETE /user/{%s}/cart/{%s}", userIDPath, skuIDPath), func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(r.PathValue(userIDPath), 10, 64)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrMissingUserID, http.StatusBadRequest)
			return
		}
		skuID, err := strconv.ParseInt(r.PathValue(skuIDPath), 10, 64)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrMissingSKU, http.StatusBadRequest)
			return
		}

		if err := service.RemoveFromCart(ctx, int64(userID), int64(skuID)); err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrFailedDeleteAGood, http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusNoContent)
	})

	// Зачистить корзину: DELETE /user/<user_id>/cart
	mux.HandleFunc(fmt.Sprintf("DELETE /user/{%s}/cart", userIDPath), func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(r.PathValue(userIDPath), 10, 64)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrMissingUserID, http.StatusBadRequest)
			return
		}

		if err := service.CleanUpCart(ctx, int64(userID)); err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrFailedCleanUpACart, http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusNoContent)
	})

	// Получить корзину: GET /user/<user_id>/cart
	mux.HandleFunc(fmt.Sprintf("GET /user/{%s}/cart", userIDPath), func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(r.PathValue(userIDPath), 10, 64)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrMissingUserID, http.StatusBadRequest)
			return
		}

		cart, err := service.GetCart(ctx, int64(userID))
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrFailedGetCart, http.StatusBadRequest)
			return
		}
		if cart == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resp := cart.ToCartRespV1()
		data, err := json.Marshal(resp)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrMarshalResponse, http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(data); err != nil {
			log.Println(err)
		}
	})

	// Создать заказ из корзины: POST /user/<user_id>/cart/checkout
	mux.HandleFunc(fmt.Sprintf("POST /user/{%s}/cart/checkout", userIDPath), func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(r.PathValue(userIDPath), 10, 64)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrMissingUserID, http.StatusBadRequest)
			return
		}

		order, err := service.Checkout(ctx, int64(userID))
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrFailedGetCart, http.StatusInternalServerError)
			return
		}

		resp := order.ToOrderV1()
		data, err := json.Marshal(resp)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrMarshalResponse, http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(data); err != nil {
			log.Println(err)
		}
	})
}
