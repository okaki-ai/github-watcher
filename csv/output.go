package csv

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Repo struct {
	Id                   int    `csv:"ID"`
	Repository           string `csv:"Repository"`
	Tag                  string `csv:"Tag"`
	Description          string `csv:"Description"`
	Url                  string `csv:"URL"`
	Stars                int    `csv:"Stars"`
	Current_period_stars int    `csv:"Current_period_stars"`
}

type Company struct {
	Id         int    `csv:"ID"`
	Repository string `csv:"Repository"`
	Company    string `csv:"Company"`
	Commits    int    `csv:"Commits"`
}

func WriteReposToCSV(filename string, data []Repo) error {
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err := gocsv.MarshalFile(data, file); err != nil {
		return fmt.Errorf("[csv] failed to output csv")
	}
	return nil
}

func WriteCompanyToCSV(filename string, data []Company) error {
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err := gocsv.MarshalFile(data, file); err != nil {
		return fmt.Errorf("[csv] failed to output csv")
	}
	return nil
}
