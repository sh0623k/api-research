package errors

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCombineErrors(t *testing.T) {
	t.Run(
		"Return wrapped error when original error and second error are not nil",
		func(t *testing.T) {
			originalErr := errors.New("original error")
			secondErr := errors.New("second error")

			actual := CombineErrors(originalErr, secondErr).Error()

			expected := "second error: original error"
			assert.Equal(t, actual, expected)
		},
	)

	t.Run(
		"Return second error when original error is nil and second error is not nil",
		func(t *testing.T) {
			secondErr := errors.New("second error")

			actual := CombineErrors(nil, secondErr).Error()

			expected := "second error"
			assert.Equal(t, expected, actual)
		},
	)

	t.Run(
		"Return original error when original error is not nil and second error is nil",
		func(t *testing.T) {
			originalErr := errors.New("original error")

			actual := CombineErrors(nil, originalErr).Error()

			expected := "original error"
			assert.Equal(t, expected, actual)
		},
	)

	t.Run(
		"Return nil when both arguments are nil",
		func(t *testing.T) {
			actual := CombineErrors(nil, nil)

			assert.Nil(t, actual)
		},
	)
}
