package worksheetB

import (
	"fmt"
	"math/big"

	"np_finance/internal/config"
	"np_finance/internal/ws_Step"
)

type Step12 struct {
	// Insurance Expenses
	OutputMother, OutputFather, OutputCombined big.Rat
}

func (s *Step12) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	mother, err := config.GetNamedInput("motherInsurance")
	if err != nil {
		panic(err)
	}
	father, err := config.GetNamedInput("fatherInsurance")
	if err != nil {
		panic(err)
	}
	s.OutputMother.SetString(mother)
	s.OutputFather.SetString(father)
	s.OutputCombined.Add(&s.OutputMother, &s.OutputFather)
}

func (s *Step12) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	combined, _ := s.OutputCombined.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f Combined: $%.2f", mother, father, combined)
}

type Step13 struct {
	// Work-related child care expenses
	OutputMother, OutputFather, OutputCombined big.Rat
}

func (s *Step13) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	mother, err := config.GetNamedInput("motherChildcare")
	if err != nil {
		panic(err)
	}
	father, err := config.GetNamedInput("fatherChildcare")
	if err != nil {
		panic(err)
	}
	s.OutputMother.SetString(mother)
	s.OutputFather.SetString(father)
	s.OutputCombined.Add(&s.OutputMother, &s.OutputFather)
}

func (s *Step13) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	combined, _ := s.OutputCombined.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f Combined: $%.2f", mother, father, combined)

}

type Step14 struct {
	// Additional expenses
	OutputMother, OutputFather, OutputCombined big.Rat
}

func (s *Step14) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	mother, err := config.GetNamedInput("motherOtherExpense")
	if err != nil {
		panic(err)
	}
	father, err := config.GetNamedInput("fatherOtherExpense")
	if err != nil {
		panic(err)
	}
	s.OutputMother.SetString(mother)
	s.OutputFather.SetString(father)
	s.OutputCombined.Add(&s.OutputMother, &s.OutputFather)
}

func (s *Step14) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	combined, _ := s.OutputCombined.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f Combined: $%.2f", mother, father, combined)

}

type Step15 struct {
	// Total additional payments
	OutputMother, OutputFather, OutputCombined big.Rat
}

func (s *Step15) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step12 := (*steps)[11].(*Step12)
	step13 := (*steps)[12].(*Step13)
	step14 := (*steps)[13].(*Step14)
	s.OutputMother.Add(&step12.OutputMother, &step13.OutputMother)
	s.OutputMother.Add(&s.OutputMother, &step14.OutputMother)
	s.OutputFather.Add(&step12.OutputFather, &step13.OutputFather)
	s.OutputFather.Add(&s.OutputFather, &step14.OutputFather)
	s.OutputCombined.Add(&s.OutputMother, &s.OutputFather)
}

func (s *Step15) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	combined, _ := s.OutputCombined.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f Combined: $%.2f", mother, father, combined)

}

type Step16 struct {
	// Each parent's Obligation
	OutputMother, OutputFather big.Rat
}

func (s *Step16) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step15 := (*steps)[14].(*Step15)
	step2 := (*steps)[1].(*Step2)
	s.OutputMother.Mul(&step15.OutputCombined, &step2.OutputMother)
	s.OutputFather.Mul(&step15.OutputCombined, &step2.OutputFather)
}

func (s *Step16) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f", mother, father)
}

type Step17 struct {
	// Amount transferred
	// Parent with the "minus" figure pays that amount to the other parent
	OutputMother, OutputFather big.Rat
	Payer                      string
}

func (s *Step17) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step16 := (*steps)[15].(*Step16)
	step15 := (*steps)[14].(*Step15)
	s.OutputMother.Sub(&step15.OutputMother, &step16.OutputMother)
	s.OutputFather.Sub(&step15.OutputFather, &step16.OutputFather)
	if s.OutputMother.Sign() == -1 {
		s.Payer = "mother"
	} else {
		s.Payer = "father"
	}
}

func (s *Step17) Display() string {
	mother, _ := s.OutputMother.Float64()
	father, _ := s.OutputFather.Float64()
	return fmt.Sprintf("Mother: $%.2f Father: $%.2f Payer: %s", mother, father, s.Payer)
}

type Step18 struct {
	// Net Transferred
	// combine lines 11 and 17 by addition if same parent pays, by subtraction if different parent pays
	Output big.Rat
	Payer  string
}

func (s *Step18) Execute(config *config.WorksheetConfig, worksheet *ws_Step.WorkSheet) {
	steps := &worksheet.Steps
	step11 := (*steps)[10].(*Step11)
	step17 := (*steps)[16].(*Step17)
	var step17Output big.Rat
	if step17.Payer == "mother" {
		step17Output = step17.OutputMother
	} else {
		step17Output = step17.OutputFather
	}
	step17Output.Neg(&step17Output)
	if step11.Payer == step17.Payer {
		s.Output.Add(&step11.Output, &step17Output)
	} else {
		s.Output.Sub(&step11.Output, &step17Output)
	}
}

func (s *Step18) Display() string {
	output, _ := s.Output.Float64()
	return fmt.Sprintf("Net Transferred: $%.2f, payer: %s", output, s.Payer)
}
