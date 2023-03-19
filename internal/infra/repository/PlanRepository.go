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

func (repository *PlanRepository) Find() ([]*entity.Plan, error) {
	query, err := repository.DB.Query("SELECT * FROM plans")

	if err != nil {
		return nil, err
	}

	defer query.Close()

	var plans []*entity.Plan

	for query.Next() {
		var plan *entity.Plan

		err = query.Scan(&plan.ID, &plan.ExchangeType, &plan.Amount, &plan.Name)

		if err != nil {
			return nil, err
		}

		plans = append(plans, plan)
	}

	return plans, nil
}
