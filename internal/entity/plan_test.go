package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfIdIsBlank(t *testing.T) {
	plan := &Plan{}

	assert.Error(t, plan.Validate(), "Plan id is required")
}

func TestIfExchangeTypeIsBlank(t *testing.T) {
	plan := &Plan{}

	plan.ID = "ID"

	assert.Error(t, plan.Validate(), "Plan exchange type is required")
}

func TestIfAmountIsBlank(t *testing.T) {
	plan := &Plan{}

	plan.ID = "ID"
	plan.ExchangeType = "Expend"
	plan.Amount = 0

	assert.Error(t, plan.Validate(), "Plan amount is required")
}

func TestIfNametIsBlank(t *testing.T) {
	plan := &Plan{}

	plan.ID = "ID"
	plan.ExchangeType = "Expend"
	plan.Amount = 10.0

	assert.Error(t, plan.Validate(), "Plan name is required")
}
