package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-github/github"
)

type Repositories struct {
	query           string
	keyword         string
	language        string
	spoken_language string
	since           string
	client          *github.Client
}

type TrendingResult []struct {
	Author             string `json:"author"`
	Name               string `json:"name"`
	Avatar             string `json:"avatar"`
	URL                string `json:"url"`
	Description        string `json:"description"`
	Language           string `json:"language"`
	LanguageColor      string `json:"languageColor"`
	Stars              int    `json:"stars"`
	Forks              int    `json:"forks"`
	CurrentPeriodStars int    `json:"currentPeriodStars"`
	BuiltBy            []struct {
		Href     string `json:"href"`
		Avatar   string `json:"avatar"`
		Username string `json:"username"`
	} `json:"builtBy"`
}

// Search All Repository
func (r *Repositories) SearchAllRepository(a Repositories) (*github.RepositoriesSearchResult, error) {
	ctx := context.Background()
	opts := &github.SearchOptions{
		Sort:  "stars",
		Order: "desc",
	}
	query := a.keyword + " in:name,description,readme language:" + a.language
	res, _, err := a.client.Search.Repositories(ctx, query, opts)
	if err != nil {
		return nil, fmt.Errorf("[repositoris] failed to search repository")
	}
	return res, nil
}

// Search Trending Repository
func (r *Repositories) SearchTrendingRepository(a Repositories) (*TrendingResult, error) {
	request_url := "https://ghapi.huchen.dev/repositories?" + "language=" + a.language + "&since=" + a.since + "&spoken_language_code=" + a.spoken_language
	resp, _ := http.Get(request_url)
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("[oauth]: bad response status code %d", resp.StatusCode)
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	data := new(TrendingResult)
	if err := json.Unmarshal(byteArray, data); err != nil {
		return nil, err
	}
	return data, nil
}
