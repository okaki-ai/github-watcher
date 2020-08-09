package contributors

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

type Contributors struct {
	Contributors []*github.ContributorStats
	Client       *github.Client
}

type User struct {
	user_name     string
	company       string
	total_commits int
}

type CompanyCommits struct {
	company       string
	total_commits int
}

func (c *Contributors) ListContributorsForRepo(owner string, repository string) (*Contributors, error) {
	ctx := context.Background()
	contributors, _, err := c.Client.Repositories.ListContributorsStats(ctx, owner, repository)
	if err != nil {
		return nil, fmt.Errorf("[users] failed to get contributors")
	}
	b := &Contributors{
		Contributors: contributors,
	}
	return b, nil
}

func (c *Contributors) GetCompanyOfUserFromList() ([]User, error) {
	var users []User
	for i, v := range c.Contributors {
		user := *(v.Author.Login)
		users[i].user_name = *(v.Author.Login)
		company, _ := c.GetCompanyOfUser(user)
		users[i].company = *company
		users[i].total_commits = v.GetTotal()
	}
	return users, nil
}

func (c *Contributors) SummaryForEachCompany(users []User) []CompanyCommits {
	var res []CompanyCommits
	for _, t := range users {
		res = c.AddComitsForEachCompany(res, t)
	}
	return res
}

func (c *Contributors) AddComitsForEachCompany(targets []CompanyCommits, user User) []CompanyCommits {
	var exist = false
	for _, t := range targets {
		if t.company == user.company {
			t.total_commits = t.total_commits + user.total_commits
			exist = true
			break
		}
	}
	if exist == false {
		n := CompanyCommits{
			company:       user.company,
			total_commits: user.total_commits,
		}
		targets = append(targets, n)
	}
	return targets
}

func (c *Contributors) GetCompanyOfUser(user string) (*string, error) {
	ctx := context.Background()
	res, _, err := c.Client.Users.Get(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("[users] failed to get company info.")
	}
	return res.Company, nil
}
