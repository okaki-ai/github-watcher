package repositories

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

type Repositories struct {
	query    string
	keyword  string
	language string
	since    string
	client   *github.Client
}

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

func (r *Repositories) SearchTrendingReposity(a Repositories) error {

}
