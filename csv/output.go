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
	Current_period_stars string `csv:"Current_period_stars"`
	total_commits        int    `csv:"total_commits"`
}

type Company struct {
	id         int    `csv:"ID"`
	repository string `csv:"Repository"`
	company    string `csv:"Company"`
	commits    int    `csv:"Commits"`
}

func (r *Repo) WriteReposToCSV(data []Repo) error {
	file, _ := os.OpenFile("repository.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err := gocsv.MarshalFile(data, file); err != nil {
		return fmt.Errorf("[csv] failed to output csv")
	}
	return nil
}

func (r *Repo) WriteCompanyToCSV(data []Company) error {
	file, _ := os.OpenFile("company.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err := gocsv.MarshalFile(data, file); err != nil {
		return fmt.Errorf("[csv] failed to output csv")
	}
	return nil
}
