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
	for k, v := range res {
		Repos = append(Repos, csv.Repo{k, v.Name, v.Description, v.URL, v.Stars, v.CurrentPeriodStars})
		con := users.GetContributorsInstance()
		if err := con.ListContributorsForRepo(v.Author, v.Name); err != nil {
			log.Fatalf("[watcher] failed to list contribotutors (error: %v)", err)
		}
		log.Printf("CCCCCCCCC")
		users, _ := con.GetCompanyOfUserFromList()
		company_list := con.SummaryForEachCompany(users)
		for _, c := range company_list {
			Companies = append(Companies, csv.Company{k, v.Name, c.Company, c.TotalCommits})
		}
	}
	csv.WriteReposToCSV(repo_csv, Repos)
	csv.WriteCompanyToCSV(company_csv, Companies)
	return nil
}
