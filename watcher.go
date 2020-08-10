package watcher

import (
	"log"

	"github.com/okaki-ai/github-watcher/contributors"
	"github.com/okaki-ai/github-watcher/repositories"
)

const (
	language        = "go"
	since           = "weekly"
	spoken_language = ""
)

func GetTrendingRepoController() error {
	t := repositories.GetTrendingRepositriesInstance()
	t.SetParmsTrendingRepository(language, since, spoken_language)
	res, err := t.SearchTrendingRepository()
	if err != nil {
		log.Printf("[watcher] failed to search trending repo (error: %v)", err)
		return err
	}
	for _, v := range res.TrendingResults {
		con := contributors.GetContributorsInstance()
		res := con.ListContributorsForRepo(v.Author, v.Name)
		users := res.GetCompanyOfUserFromList()
		company_list := res.SummaryForEachCompany(users)
		for _, c := range company_list {

		}
	}
}
