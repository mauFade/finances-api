package usecase

import (
	"fmt"

	"github.com/mauFade/finances-api/internal/entity"
)

type CreatePlanInputDTO struct {
	ExchangeType string  `json:"exchange_type"`
	Amount       float64 `json:"amount"`
	Name         string  `json:"name"`
}

type CreatePlanOutputDTO struct {
	ID           string
	ExchangeType string
	Amount       float64
	Name         string
}

type CreatePlanUseCase struct {
	PlanRepository entity.PlanRepository
}

func NewCreatePlanUseCase(repo entity.PlanRepository) *CreatePlanUseCase {
	return &CreatePlanUseCase{
		PlanRepository: repo,
	}
}

func (uc *CreatePlanUseCase) Execute(data CreatePlanInputDTO) (*CreatePlanOutputDTO, error) {
	plan := entity.NewPLan(data.ExchangeType, data.Amount, data.Name)

	err := plan.Validate()

	if err != nil {
		fmt.Printf("Entity error: \n%f", err)
	}

	err = uc.PlanRepository.Create(plan)

	if err != nil {
		return nil, err
	}

	return &CreatePlanOutputDTO{
		ID:           plan.ID,
		ExchangeType: plan.ExchangeType,
		Amount:       plan.Amount,
		Name:         plan.Name,
	}, nil
}
