package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Plan struct {
	ID           string
	ExchangeType string
	Amount       float64
	Name         string
}

func NewPLan(exchangeType string, amount float64, name string) *Plan {
	return &Plan{
		ID:           uuid.NewString(),
		ExchangeType: exchangeType,
		Amount:       amount,
		Name:         name,
	}
}

func (plan *Plan) Validate() error {
	if plan.ID == "" {
		return errors.New("Plan id is required")
	}

	if plan.ExchangeType == "" {
		return errors.New(("Plan exchange type is required"))
	}

	if plan.Amount == 0 {
		return errors.New(("Plan amount is required"))
	}

	if plan.Name == "" {
		return errors.New(("Plan name is required"))
	}

	return nil
}
