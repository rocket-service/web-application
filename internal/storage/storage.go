package storage

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUserExists           = errors.New("user already exists")
	ErrIncorrectPassword    = errors.New("incorrect password")
	ErrSubscriptionNotFound = errors.New("subscription not found")
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidPromoCode     = errors.New("invalid promo code")
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrPromoCodeNotFound    = errors.New("promo code not found")
	ErrPromoCodeAlreadyUsed = errors.New("promo code already used")
	ErrPromoCodeExpired     = errors.New("promo code expired")
)
