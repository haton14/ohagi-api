package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFood(t *testing.T) {
	food, err := NewFood(1, "ペレット", 1.6, "g")
	assert.NoError(t, err)
	assert.Equal(t, 1, food.ID())
	assert.Equal(t, "ペレット", food.Name())
	assert.Equal(t, 1.6, food.Amount())
	assert.Equal(t, "g", food.Unit())
	assert.Equal(t, false, food.IsNullID())
	assert.Equal(t, false, food.IsNullAmount())

	food, err = NewFood(0, "ペレット", 0, "g")
	assert.NoError(t, err)
	assert.Equal(t, 0, food.ID())
	assert.Equal(t, "ペレット", food.Name())
	assert.Equal(t, float64(0), food.Amount())
	assert.Equal(t, "g", food.Unit())
	assert.Equal(t, true, food.IsNullID())
	assert.Equal(t, true, food.IsNullAmount())
	food.SetID(1234)
	assert.Equal(t, 1234, food.ID())

	_, err = NewFood(-1, "", -1, "")
	assert.Error(t, err)
	assert.Equal(t, errors.New("id isn't negative number"), err)

	_, err = NewFood(1, "", -1, "")
	assert.Error(t, err)
	assert.Equal(t, errors.New("name isn't empty"), err)

	_, err = NewFood(1, "food", -1, "")
	assert.Error(t, err)
	assert.Equal(t, errors.New("amount isn't negative number"), err)

	_, err = NewFood(1, "food", 1, "")
	assert.Error(t, err)
	assert.Equal(t, errors.New("unit isn't empty"), err)
}
