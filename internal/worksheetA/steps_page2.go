package worksheetA

import (
	"fmt"
	"math/big"

	"np_finance/internal/config"
	"np_finance/internal/ws_Step"
)

type Step5 struct {
	// Insurance Expenses
	OutputMother, OutputFather, OutputCombined big.Rat
}

func (s *Step5) Execute(config *config.WorksheetConfig, form *ws_Step.Form) {
	mother, _ := config.Gsi(5, "mother")
	father, _ := config.Gsi(5, "father")
	s.OutputMother.SetString(mother)
	s.OutputFather.SetString(father)
	s.OutputCombined.Add(&s.OutputMother, &s.OutputFather)
}

func (s *Step5) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	combined, _ := s.OutputCombined.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f Combined: $%.2f", mother, father, combined)
}

type Step6 struct {
	// Work-related child care expenses
	OutputMother, OutputFather, OutputCombined big.Rat
}

func (s *Step6) Execute(config *config.WorksheetConfig, form *ws_Step.Form) {
	mother, _ := config.Gsi(6, "mother")
	father, _ := config.Gsi(6, "father")
	s.OutputMother.SetString(mother)
	s.OutputFather.SetString(father)
	s.OutputCombined.Add(&s.OutputMother, &s.OutputFather)
}

func (s *Step6) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	combined, _ := s.OutputCombined.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f Combined: $%.2f", mother, father, combined)

}

type Step7 struct {
	// Additional expenses
	OutputMother, OutputFather, OutputCombined big.Rat
}

func (s *Step7) Execute(config *config.WorksheetConfig, form *ws_Step.Form) {
	mother, _ := config.Gsi(7, "mother")
	father, _ := config.Gsi(7, "father")
	s.OutputMother.SetString(mother)
	s.OutputFather.SetString(father)
	s.OutputCombined.Add(&s.OutputMother, &s.OutputFather)
}

func (s *Step7) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	combined, _ := s.OutputCombined.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f Combined: $%.2f", mother, father, combined)

}

type Step8 struct {
	// Total additional payments
	OutputMother, OutputFather, OutputCombined big.Rat
}

func (s *Step8) Execute(config *config.WorksheetConfig, form *ws_Step.Form) {
	steps := &form.Steps
	step4 := (*steps)[3].(*Step4)
	step5 := (*steps)[4].(*Step5)
	step6 := (*steps)[5].(*Step6)
	step7 := (*steps)[6].(*Step7)
	s.OutputMother.Add(&step5.OutputMother, &step6.OutputMother)
	s.OutputMother.Add(&s.OutputMother, &step7.OutputMother)
	s.OutputFather.Add(&step5.OutputFather, &step6.OutputFather)
	s.OutputFather.Add(&s.OutputFather, &step7.OutputFather)
	s.OutputCombined.Add(&s.OutputMother, &s.OutputFather)
	s.OutputCombined.Add(&s.OutputCombined, &step4.Output)
}

func (s *Step8) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	combined, _ := s.OutputCombined.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f Combined: $%.2f", mother, father, combined)

}

type Step9 struct {
	// Each parent's Obligation
	OutputMother, OutputFather big.Rat
}

func (s *Step9) Execute(config *config.WorksheetConfig, form *ws_Step.Form) {
	steps := &form.Steps
	step8 := (*steps)[7].(*Step8)
	step2 := (*steps)[1].(*Step2)
	s.OutputMother.Mul(&step8.OutputCombined, &step2.OutputMother)
	s.OutputFather.Mul(&step8.OutputCombined, &step2.OutputFather)
}

func (s *Step9) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f", mother, father)
}

type Step10 struct {
	// repeat line 8
	OutputMother, OutputFather, OutputCombined big.Rat
}

func (s *Step10) Execute(config *config.WorksheetConfig, form *ws_Step.Form) {
	steps := &form.Steps
	step8 := (*steps)[7].(*Step8)
	s.OutputMother = step8.OutputMother
	s.OutputFather = step8.OutputFather
	s.OutputCombined = step8.OutputCombined
}

func (s *Step10) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	combined, _ := s.OutputCombined.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f Combined: $%.2f", mother, father, combined)

}

type Step11 struct {
	// Each parent's net obligation
	OutputMother, OutputFather big.Rat
}

func (s *Step11) Execute(config *config.WorksheetConfig, form *ws_Step.Form) {
	steps := &form.Steps
	step9 := (*steps)[8].(*Step9)
	step10 := (*steps)[9].(*Step10)
	s.OutputMother.Sub(&step9.OutputMother, &step10.OutputMother)
	s.OutputFather.Sub(&step9.OutputFather, &step10.OutputFather)
}

func (s *Step11) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f", mother, father)
}
