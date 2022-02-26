package anycast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToIntP(t *testing.T) {
	t.Run("正常系 整数", func(t *testing.T) {
		want := 1
		have := ToIntP(1)
		assert.Equal(t, &want, have)
	})

	t.Run("正常系 負数", func(t *testing.T) {
		want := -1
		have := ToIntP(-1)
		assert.Equal(t, &want, have)
	})

	t.Run("異常系 数値以外の変換 数値になり得るstring", func(t *testing.T) {
		have := ToIntP("1")
		assert.Equal(t, (*int)(nil), have)
	})

	t.Run("異常系 数値以外の変換 string", func(t *testing.T) {
		have := ToIntP("string")
		assert.Equal(t, (*int)(nil), have)
	})

	t.Run("異常系 数値以外の変換 bool", func(t *testing.T) {
		have := ToIntP(true)
		assert.Equal(t, (*int)(nil), have)
	})

	t.Run("異常系 数値以外の変換 nil", func(t *testing.T) {
		have := ToIntP(nil)
		assert.Equal(t, (*int)(nil), have)
	})
}
