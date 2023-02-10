package input

import (
	"encoding/csv"
	"errors"
	"strings"
)

//go:generate go run github.com/golang/mock/mockgen -package=mocks -destination=mocks/mock_reader.go . Reader
type (
	Parser struct {
		reader Reader
	}

	Reader interface {
		ReadData() (string, error)
	}

	Problem struct {
		Question string
		Answer   string
	}
)

func (p Parser) ParseData() ([]Problem, error) {
	data, err := p.reader.ReadData()
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(strings.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var problems []Problem
	for r := range records {
		if len(records[r]) != 2 {
			return nil, errors.New("data not in the expected format of question,answer")
		}
		problems = append(problems, Problem{Question: records[r][0], Answer: records[r][1]})
	}

	return problems, nil
}

func NewParser(reader Reader) *Parser {
	return &Parser{
		reader: reader,
	}
}
