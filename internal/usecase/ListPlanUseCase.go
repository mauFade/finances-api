package usecase

import "github.com/mauFade/finances-api/internal/entity"

type ListPlanUseCase struct {
	PlanRepository entity.PlanRepository
}

func NewListPlanUseCase(repo entity.PlanRepository) *ListPlanUseCase {
	return &ListPlanUseCase{
		PlanRepository: repo,
	}
}

func (useCase *ListPlanUseCase) Execute() ([]*entity.Plan, error) {
	query, err := useCase.PlanRepository.Find()

	if err != nil {
		return nil, err
	}

	return query, nil
}
