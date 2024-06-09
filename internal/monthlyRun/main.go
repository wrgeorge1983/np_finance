package monthlyRun

import (
	"fmt"
	"math/big"
)

type Input struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type MonthlyRun struct {
	Title   string   `yaml:"title"`
	Inputs  []Input  `yaml:"inputs"`
	Outputs []string `yaml:"outputs"`
}

func (mr *MonthlyRun) GetInputString(name string) (string, error) {
	for _, s := range mr.Inputs {
		if s.Name == name {
			return s.Value, nil
		}
	}
	return "", fmt.Errorf("input %s not found", name)
}

func (mr *MonthlyRun) GetInputRat(name string) (big.Rat, error) {
	s, err := mr.GetInputString(name)
	if err != nil {
		return big.Rat{}, err
	}
	r := big.Rat{}
	r.SetString(s)
	return r, nil
}

func (mr *MonthlyRun) SetInputString(name, value string) string {
	for i, s := range mr.Inputs {
		if s.Name == name {
			mr.Inputs[i].Value = value
			return value
		}
	}
	mr.Inputs = append(mr.Inputs, Input{Name: name, Value: value})
	return value
}

func (mr *MonthlyRun) SetInputRat(name string, value big.Rat) string {
	return mr.SetInputString(name, value.String())
}

func (mr *MonthlyRun) GetOutputString(name string) (string, error) {
	for _, s := range mr.Outputs {
		if s == name {
			return s, nil
		}
	}
	return "", fmt.Errorf("output %s not found", name)
}

func (mr *MonthlyRun) GetOutputRat(name string) (big.Rat, error) {
	s, err := mr.GetOutputString(name)
	if err != nil {
		return big.Rat{}, err
	}
	r := big.Rat{}
	r.SetString(s)
	return r, nil
}
