package worksheetB

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
	mother, err := config.GetNamedInput("motherSalary")
	if err != nil {
		panic(err)
	}
	father, err := config.GetNamedInput("fatherSalary")
	if err != nil {
		panic(err)
	}

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
	children, err := config.GetNamedInput("children")
	if err != nil {
		panic(err)
	}
	s.OutputChildren.SetString(children)
}

func (s *Step3) Display() string {
	return fmt.Sprintf("Number of Children: %v", s.OutputChildren.Num())
}

type Step4 struct {
	// get basic child support from schedule
	OutputBasicChildSupport big.Rat
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
	s.OutputBasicChildSupport.SetInt64(int64(support))
}

func (s *Step4) Display() string {
	output, _ := s.OutputBasicChildSupport.Float64()
	return fmt.Sprintf("Basic Child Support: $%.2f", output)
}

type Step5 struct {
	// shared responsibility basic obligation
	OutputSharedResponsibilityBasicObligation big.Rat
}

func (s *Step5) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step4 := (*steps)[3].(*Step4)
	// needs to be 1.5
	multiplier := new(big.Rat)
	multiplier.SetString("1.5")
	s.OutputSharedResponsibilityBasicObligation.Mul(&step4.OutputBasicChildSupport, multiplier)
}

func (s *Step5) Display() string {
	output, _ := s.OutputSharedResponsibilityBasicObligation.Float64()
	return fmt.Sprintf("Shared Responsibility Basic Obligation: $%.2f", output)
}

type Step6 struct {
	// each parent's share
	OutputMother, OutputFather big.Rat
}

func (s *Step6) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step2 := (*steps)[1].(*Step2)
	step5 := (*steps)[4].(*Step5)
	s.OutputMother.Mul(&step5.OutputSharedResponsibilityBasicObligation, &step2.OutputMother)
	s.OutputFather.Mul(&step5.OutputSharedResponsibilityBasicObligation, &step2.OutputFather)
}

func (s *Step6) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f", mother, father)
}

type Step7 struct {
	// number of 24 hour days with each parent
	OutputMother, OutputFather big.Rat
}

func (s *Step7) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	mother, err := config.GetNamedInput("motherDays")
	if err != nil {
		panic(err)
	}
	father, err := config.GetNamedInput("fatherDays")
	if err != nil {
		panic(err)
	}
	s.OutputMother.SetString(mother)
	s.OutputFather.SetString(father)
}

func (s *Step7) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	return fmt.Sprintf("Mother: %v Father: %v", mother, father)
}

type Step8 struct {
	// percentage of time with each parent
	OutputMother, OutputFather big.Rat
}

func (s *Step8) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step7 := (*steps)[6].(*Step7)
	total := new(big.Rat)
	total.Add(&step7.OutputMother, &step7.OutputFather)
	//year := new(big.Rat)
	//year.SetInt64(365)
	// some math
	s.OutputMother.Quo(&step7.OutputMother, total)
	s.OutputFather.Quo(&step7.OutputFather, total)
}

func (s *Step8) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	return fmt.Sprintf("Mother: %.2f%% Father: %.2f%%", mother*100, father*100)
}

type Step9 struct {
	// Amount retained
	OutputMother, OutputFather big.Rat
}

func (s *Step9) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step6 := (*steps)[5].(*Step6)
	step8 := (*steps)[7].(*Step8)
	s.OutputMother.Mul(&step6.OutputMother, &step8.OutputMother)
	s.OutputFather.Mul(&step6.OutputFather, &step8.OutputFather)
}

func (s *Step9) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f", mother, father)
}

type Step10 struct {
	// Each parent's obligation
	OutputMother, OutputFather big.Rat
}

func (s *Step10) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step6 := (*steps)[5].(*Step6)
	step9 := (*steps)[8].(*Step9)
	s.OutputMother.Sub(&step6.OutputMother, &step9.OutputMother)
	s.OutputFather.Sub(&step6.OutputFather, &step9.OutputFather)
}

func (s *Step10) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f", mother, father)
}

type Step11 struct {
	// Amount transferred
	Output big.Rat
	Payer  string
}

func (s *Step11) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step10 := (*steps)[9].(*Step10)

	mother := step10.OutputMother
	father := step10.OutputFather
	switch mother.Cmp(&father) {
	case 1:
		s.Output.Sub(&step10.OutputMother, &step10.OutputFather)
		s.Payer = "mother"
	default:
		s.Output.Sub(&step10.OutputFather, &step10.OutputMother)
		s.Payer = "father"
	}
}

func (s *Step11) Display() string {
	output, _ := s.Output.Float64()
	return fmt.Sprintf("Amount Transferred: $%.2f Payer: %s", output, s.Payer)

}
