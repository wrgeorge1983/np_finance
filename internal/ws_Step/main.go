package ws_Step

import "np_finance/internal/config"

type Step interface {
	Execute(config *config.WorksheetConfig, workSheet *WorkSheet)
	Display() string
}

type WorkSheet struct {
	Schedule config.Schedule
	Steps    []Step
}
