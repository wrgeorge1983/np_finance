package worksheetB

import (
	"np_finance/internal/config"
)

type Step interface {
	Execute(config *config.WorksheetConfig, form *Form)
	Display() string
}

type Form struct {
	Schedule config.Schedule
	Steps    []Step
}

func NewForm(schedule *config.Schedule) *Form {
	return &Form{
		Schedule: *schedule,
		Steps: []Step{
			&Step1{},
			&Step2{},
			&Step3{},
			&Step4{},
			&Step5{},
			&Step6{},
			&Step7{},
			&Step8{},
			&Step9{},
			&Step10{},
			&Step11{},
			&Step12{},
			&Step13{},
			&Step14{},
			&Step15{},
			&Step16{},
			&Step17{},
			&Step18{},
		},
	}
}
