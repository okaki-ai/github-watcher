package users

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
	Name         string
	Company      string
	TotalCommits int
}

type CompanyCommits struct {
	Company      string
	TotalCommits int
}

var sharedContributorsInstance *Contributors

func GetContributorsInstance() *Contributors {
	if sharedContributorsInstance == nil {
		sharedContributorsInstance = &Contributors{}
	}
	return sharedContributorsInstance
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
		if t.Company == user.Company {
			t.TotalCommits = t.TotalCommits + user.TotalCommits
			exist = true
			break
		}
	}
	if exist == false {
		n := CompanyCommits{
			Company:      user.Company,
			TotalCommits: user.TotalCommits,
		}
		targets = append(targets, n)
	}
	return targets
}

// based on contributors struct make []User struct
func (c *Contributors) GetCompanyOfUserFromList() ([]User, error) {
	var users []User
	for i, v := range c.Contributors {
		user := *(v.Author.Login)
		users[i].Name = *(v.Author.Login)
		company, _ := c.GetCompanyOfUser(user)
		users[i].Company = *company
		users[i].TotalCommits = v.GetTotal()
	}
	return users, nil
}

func (c *Contributors) GetCompanyOfUser(user string) (*string, error) {
	ctx := context.Background()
	res, _, err := c.Client.Users.Get(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("[users] failed to get company info.")
	}
	return res.Company, nil
}
