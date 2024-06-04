package config

import "fmt"

type Schedule struct {
	RawData []string
	Data    []struct {
		Income  int
		Support []int
	}
}

func (s *Schedule) GetSupport(income, nChildren int) (int, error) {
	income = (income + 25) / 50 * 50 // round to nearest 50
	if income < s.Data[0].Income {
		return s.Data[0].Support[nChildren-1], nil
	}
	for _, d := range s.Data {
		if d.Income == income {
			return d.Support[nChildren-1], nil
		}
	}
	return 0, fmt.Errorf("income %d not found", income)
}
