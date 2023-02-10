package input_test

import (
	"errors"
	"testing"

	"quiz/input"
	"quiz/input/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCsvParser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	reader := mocks.NewMockReader(ctrl)

	t.Run("ParseData returns error if data is empty", func(t *testing.T) {
		t.Parallel()
		reader.EXPECT().ReadData().Return("", errors.New("failed"))
		parser := input.NewParser(reader)
		_, err := parser.ParseData()
		assert.Error(t, err)
	})

	t.Run("ParseData returns no error if there is data", func(t *testing.T) {
		t.Parallel()
		reader.EXPECT().ReadData().Return("1+1,2\n2+2,4", nil)
		parser := input.NewParser(reader)
		_, err := parser.ParseData()
		assert.NoError(t, err)
	})

	t.Run("ParseData returns expected data when there is valid input", func(t *testing.T) {
		t.Parallel()
		want := []input.Problem{{Question: "1+1", Answer: "2"}, {Question: "2+2", Answer: "4"}}
		reader.EXPECT().ReadData().Return("1+1,2\n2+2,4", nil)
		parser := input.NewParser(reader)
		got, _ := parser.ParseData()
		assert.Equal(t, want, got)
	})
}
