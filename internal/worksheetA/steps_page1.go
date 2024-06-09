package worksheetA

import (
	"fmt"
	"math/big"

	"np_finance/internal/config"
	"np_finance/internal/ws_Step"
)

type Step1 struct {
	OutputMother, OutputFather, OutputCombined big.Rat
}

func (s *Step1) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	// Get the mother and father gross monthly income
	mother, _ := config.Gsi(1, "mother")
	father, _ := config.Gsi(1, "father")

	s.OutputMother.SetString(mother)
	s.OutputFather.SetString(father)
	s.OutputCombined.Add(&s.OutputMother, &s.OutputFather)
}

func (s *Step1) Display() string {

	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	combined, _ := s.OutputCombined.Float64()

	return fmt.Sprintf("Mother: $%.2f Father: $%.2f Combined: $%.2f", mother, father, combined)
}

type Step2 struct {
	// compute each parent's income divided by combined income
	OutputMother, OutputFather big.Rat
}

func (s *Step2) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step1 := (*steps)[0].(*Step1)
	s.OutputMother.Quo(&step1.OutputMother, &step1.OutputCombined)
	s.OutputFather.Quo(&step1.OutputFather, &step1.OutputCombined)
}

func (s *Step2) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	return fmt.Sprintf("Mother: %.2f%% Father: %.2f%%", mother*100, father*100)
}

type Step3 struct {
	// get the total number of children
	OutputChildren big.Rat
}

func (s *Step3) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	children, _ := config.Gsi(3, "number")
	s.OutputChildren.SetString(children)
}

func (s *Step3) Display() string {
	return fmt.Sprintf("Number of Children: %v", s.OutputChildren.Num())
}

type Step4 struct {
	// get basic child support from schedule
	Output big.Rat
}

func (s *Step4) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step1 := (*steps)[0].(*Step1)
	step3 := (*steps)[2].(*Step3)
	f, _ := step1.OutputCombined.Float64()
	income := int(f)
	f, _ = step3.OutputChildren.Float64()
	children := int(f)
	support, err := worksheet.Schedule.GetSupport(income, children)
	if err != nil {
		panic(err)
	}
	// some math
	s.Output.SetInt64(int64(support))
}

func (s *Step4) Display() string {
	result, _ := s.Output.Float64()
	return fmt.Sprintf("Basic Child Support: $%.2f", result)
}
