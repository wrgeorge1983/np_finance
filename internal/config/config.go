package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type WorksheetConfig struct {
	Title      string `yaml:"title"`
	Worksheet  string `yaml:"worksheet"`
	Schedule   string `yaml:"schedule"`
	StepInputs []struct {
		Step   int    `yaml:"step"`
		Title  string `yaml:"title"`
		Inputs []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
		} `yaml:"inputs"`
	} `yaml:"step_inputs"`
}

func (c *WorksheetConfig) ReadConfig(filename string) *WorksheetConfig {
	buf, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	err = yaml.Unmarshal(buf, c)
	return c
}

func (c *WorksheetConfig) GetStepInput(step int, name string) (string, error) {
	for _, s := range c.StepInputs {
		if s.Step == step {
			for _, i := range s.Inputs {
				if i.Name == name {
					return i.Value, nil
				}
			}
		}
	}
	return "", fmt.Errorf("step %d, input %s not found", step, name)
}

func (c *WorksheetConfig) Gsi(step int, name string) (string, error) {
	return c.GetStepInput(step, name)
}

func (c *WorksheetConfig) ReadSchedule() *Schedule {
	// Read schedule from file
	filename := fmt.Sprintf("assets/%s", c.Schedule)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer file.Close()

	schedule := Schedule{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		schedule.RawData = append(schedule.RawData, scanner.Text())
	}

	for _, line := range schedule.RawData {
		values := strings.Fields(line)
		income, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatalf("err: %v", err)
		}
		data := struct {
			Income  int
			Support []int
		}{
			Income: income,
		}
		for _, v := range values[1:] {
			support, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("err: %v", err)
			}
			data.Support = append(data.Support, support)
		}
		schedule.Data = append(schedule.Data, data)
	}
	return &schedule
}
