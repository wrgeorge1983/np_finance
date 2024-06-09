package run_handle

import (
	"github.com/labstack/echo/v4"
)

type RunHandlers struct {
	// SetupRoutes(app *echo.Echo) error
}

type ScheduleType struct {
	Value       string
	Description string
}

type InputFormData struct {
	Title            string `form:"title"`
	ScheduleVersion  string `form:"scheduleVersion"`
	MSalary          string `form:"motherSalary"`
	FSalary          string `form:"fatherSalary"`
	MDays            int    `form:"motherDays"`
	FDays            int    `form:"fatherDays"`
	NumberOfChildren int    `form:"children"`
	MInsuranceExp    string `form:"motherInsuranceExpense"`
	FInsuranceExp    string `form:"fatherInsuranceExpense"`
	MChildcareExp    string `form:"motherChildcareExpense"`
	FChildcareExp    string `form:"fatherChildcareExpense"`
	MOtherExpense    string `form:"motherOtherExpense"`
	FOtherExpense    string `form:"fatherOtherExpense"`

	ScheduleTypes []ScheduleType
}

func DefaultFormData() InputFormData {
	return InputFormData{
		Title:            "This Month",
		ScheduleVersion:  "basic-schedule",
		MSalary:          "1000",
		FSalary:          "1000",
		MDays:            5,
		FDays:            2,
		NumberOfChildren: 1,
		MInsuranceExp:    "100",
		FInsuranceExp:    "0",
		MChildcareExp:    "0",
		FChildcareExp:    "100",
		MOtherExpense:    "0",
		FOtherExpense:    "100",
		ScheduleTypes:    ScheduleTypes,
	}
}

var ScheduleTypes = []ScheduleType{
	{Value: "basic-schedule", Description: "basic-schedule.txt"},
	{Value: "2024-basic-schedule", Description: "2024-basic-schedule.txt"},
}

func (h *RunHandlers) SetupRoutes(app *echo.Echo) error {
	fd := DefaultFormData()
	app.GET("/run", func(c echo.Context) error {
		return c.Render(200, "runIndex", fd)
	})

	app.GET("/run/start", func(c echo.Context) error {
		return c.String(200, "<hr>Step 1")
	})
	return nil
}
