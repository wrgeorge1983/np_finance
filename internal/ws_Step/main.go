package ws_Step

import "np_finance/internal/config"

type Step interface {
	Execute(config *config.WorksheetConfig, form *Form)
	Display() string
}

type Form struct {
	Schedule config.Schedule
	Steps    []Step
}
