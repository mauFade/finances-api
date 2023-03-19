package repository

import (
	"database/sql"

	"github.com/mauFade/finances-api/internal/entity"
)

type PlanRepository struct {
	DB *sql.DB
}

func NewPlanRepository(database *sql.DB) *PlanRepository {
	return &PlanRepository{DB: database}
}

func (repository *PlanRepository) Create(plan *entity.Plan) error {
	_, err := repository.DB.Exec("INSERT INTO plans (id, exchange_type, amount, name) VALUES ($1, $2, $3, $3)",
		plan.ID, plan.ExchangeType, plan.Amount, plan.Name,
	)

	if err != nil {
		return err
	}

	return nil
}
