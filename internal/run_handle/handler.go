package run_handle

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"np_finance/internal/config"
	"np_finance/internal/db"
	"np_finance/internal/log"
	"np_finance/internal/worksheetA"
	"np_finance/internal/worksheetB"
	"np_finance/internal/ws_Step"
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
	MOtherExp        string `form:"motherOtherExpense"`
	FOtherExp        string `form:"fatherOtherExpense"`

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
		MOtherExp:        "0",
		FOtherExp:        "100",
		ScheduleTypes:    ScheduleTypes,
	}
}

var ScheduleTypes = []ScheduleType{
	{Value: "basic-schedule", Description: "basic-schedule.txt"},
	{Value: "2024-basic-schedule", Description: "2024-basic-schedule.txt"},
}

func runIndex(c echo.Context) error {
	fd := DefaultFormData()
	return c.Render(200, "runIndex", fd)
}

func basicVisitation(motherDays, fatherDays int) bool {
	total := motherDays + fatherDays
	fmt.Printf("Mother: %d;  Father: %d\n", motherDays, fatherDays)
	fmt.Printf("Total: %d\n", total)

	smaller := min(motherDays, fatherDays)
	result := float64(smaller)/float64(total) < 0.35
	fmt.Printf("Result: %v\n", result)
	return result
}

func runStart(c echo.Context) error {

	formData := new(InputFormData)
	if err := c.Bind(formData); err != nil {
		return err
	}
	basic := basicVisitation(formData.MDays, formData.FDays)

	var wsType string
	var wsInit func(*config.Schedule) *ws_Step.WorkSheet
	if basic {
		wsType = "WorksheetA"
		wsInit = worksheetA.NewForm
	} else {
		wsType = "WorksheetB"
		wsInit = worksheetB.NewForm
	}

	wsConfig := config.WorksheetConfig{
		Title:     formData.Title,
		Worksheet: wsType,
		Schedule:  formData.ScheduleVersion + ".txt",
		NamedInputs: map[string]string{
			"motherSalary":       formData.MSalary,
			"fatherSalary":       formData.FSalary,
			"motherDays":         fmt.Sprintf("%d", formData.MDays),
			"fatherDays":         fmt.Sprintf("%d", formData.FDays),
			"children":           fmt.Sprintf("%d", formData.NumberOfChildren),
			"motherInsurance":    formData.MInsuranceExp,
			"fatherInsurance":    formData.FInsuranceExp,
			"motherChildcare":    formData.MChildcareExp,
			"fatherChildcare":    formData.FChildcareExp,
			"motherOtherExpense": formData.MOtherExp,
			"fatherOtherExpense": formData.FOtherExp,
		},
	}
	form := *wsInit(wsConfig.ReadSchedule())
	var output string
	fmt.Printf("%+v\n", wsConfig.NamedInputs)
	for i, step := range form.Steps {
		step.Execute(&wsConfig, &form)
		this := fmt.Sprintf("%v: %s<br>", i+1, step.Display())
		output += this
		fmt.Println(this)
	}

	return c.String(200, output)
}

func purgeDB(db *db.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		err := db.Purge()
		if err != nil {
			return c.String(500, "Error purging database")
		}
		return c.String(200, "Purged")
	}
}

func initDB(db *db.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		err := db.Init()
		if err != nil {
			log.Logger.Log().Msgf("Error initializing database: %v", err)
			return c.String(500, "Error initializing database")
		}
		return c.String(200, "Initialized")
	}
}

func showDB(db *db.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		users, _ := db.AllUsers()

		runs, _ := db.AllRuns()
		return c.String(200, fmt.Sprintf("Users: %v\nRuns: %v", users, runs))
	}
}

func (h *RunHandlers) SetupRoutes(app *echo.Echo, dbService *db.DB) error {
	app.GET("/run", runIndex)

	app.POST("/run/start", runStart)
	app.DELETE("/run/db", purgeDB(dbService))
	app.POST("/run/db", initDB(dbService))
	app.GET("/run/db", showDB(dbService))
	return nil
}
