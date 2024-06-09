package main

import (
	"fmt"
	"os"
	"strings"

	"np_finance/internal/config"
	"np_finance/internal/worksheetA"
	"np_finance/internal/worksheetB"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please provide a config file")
		os.Exit(1)
	}
	configFilename := args[0]
	wsCfg := &config.WorksheetConfig{}
	wsCfg = wsCfg.ReadConfig(configFilename)

	//for _, step := range wsCfg.Inputs {
	//	fmt.Println(step)
	//}
	schedule := wsCfg.ReadSchedule()

	if strings.ToLower(wsCfg.Worksheet) == "b" {
		form := *worksheetB.NewForm(schedule)
		for i, step := range form.Steps {
			step.Execute(wsCfg, &form)
			fmt.Printf("%v: %s\n", i+1, step.Display())
		}
	} else {
		form := *worksheetA.NewForm(schedule)
		for i, step := range form.Steps {
			step.Execute(wsCfg, &form)
			fmt.Printf("%v: %s\n", i+1, step.Display())
		}
	}
}
