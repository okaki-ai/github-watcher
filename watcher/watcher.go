package watcher

import (
	"log"

	"github.com/okaki-ai/github-watcher/csv"
	"github.com/okaki-ai/github-watcher/repositories"
	"github.com/okaki-ai/github-watcher/users"
)

const (
	language        = "go"
	since           = "weekly"
	spoken_language = ""
	repo_csv        = "./repo.csv"
	company_csv     = "./company.csv"
)

func GetTrendingRepoController() error {
	t := repositories.GetTrendingRepositriesInstance()
	t.SetParmsTrendingRepository(language, since, spoken_language)
	res, err := t.SearchTrendingRepository()
	if err != nil {
		log.Printf("[watcher] failed to search trending repo (error: %v)", err)
		return err
	}
	var Repos []csv.Repo
	var Companies []csv.Company
	for k, v := range res.TrendingResults {
		Repos[k].Id = k
		Repos[k].Repository = res.Name
		Repos[k].Description = res.Description
		Repos[k].Current_period_stars = res.CurrentPeriodStars
		Repos[k].Stars = res.Stars
		Repos[k].Url = res.URL
		con := users.GetContributorsInstance()
		contributors, _ := con.ListContributorsForRepo(v.Author, v.Name)
		users, _ := contributors.GetCompanyOfUserFromList()
		company_list := contributors.SummaryForEachCompany(users)
		for _, c := range company_list {
			Companies[k].Id = k
			Companies[k].Repository = res.Name
			Companies[k].Company = c.Company
			Companies[k].Commits = c.TotalCommits
		}
	}
	csv.WriteReposToCSV(repo_csv, Repos)
	csv.WriteCompanyToCSV(company_csv, Companies)
	return nil
}
