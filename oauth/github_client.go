package oauth

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubAuth struct {
	AccessToken string
	Client      *github.Client
}

func (g *GithubAuth) NewGithubClient(access_token string) (*GithubAuth, error) {
	g.AccessToken = access_token
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: g.AccessToken},
	)
	if ts == nil {
		return nil, fmt.Errorf("[oauth] failed to authenticate")
	}
	tc := oauth2.NewClient(ctx, ts)
	g.Client = github.NewClient(tc)
	return g, nil
}
