package worksheetA

import (
	"np_finance/internal/config"
	"np_finance/internal/ws_Step"
)

func NewForm(schedule *config.Schedule) *ws_Step.WorkSheet {
	return &ws_Step.WorkSheet{
		Schedule: *schedule,
		Steps: []ws_Step.Step{
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
			&Step5{},
			&Step6{},
			&Step7{},
			&Step8{},
			&Step9{},
			&Step10{},
			&Step11{},
		},
	}
}
