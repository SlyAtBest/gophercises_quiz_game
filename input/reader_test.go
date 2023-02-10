package input_test

import (
	"testing"

	"quiz/input"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	t.Parallel()

	t.Run("ReadData throws error if file doesn't exist", func(t *testing.T) {
		t.Parallel()
		reader := input.NewReader("../data/non-existent.csv")
		_, got := reader.ReadData()
		assert.Error(t, got)
	})

	t.Run("ReadData does not throw error if file does exist", func(t *testing.T) {
		t.Parallel()
		reader := input.NewReader("../data/problems.csv")
		_, got := reader.ReadData()
		assert.NoError(t, got)
	})

	t.Run("ReadData successfully reads file if it exists", func(t *testing.T) {
		t.Parallel()
		reader := input.NewReader("../data/problems.csv")
		got, _ := reader.ReadData()
		assert.NotEmpty(t, got)
	})
}
