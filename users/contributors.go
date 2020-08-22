package users

import (
	"context"
	"fmt"
	"log"

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
		sharedContributorsInstance.Client = github.NewClient(nil)
	}
	return sharedContributorsInstance
}
func (c *Contributors) ListContributorsForRepo(owner string, repository string) error {
	ctx := context.Background()
	contributors, resp, err := c.Client.Repositories.ListContributorsStats(ctx, owner, repository)
	if err != nil {
		return fmt.Errorf("[users] failed to get contributors, %v", err)
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("[contributors]: bad response status code %d", resp.StatusCode)
	}
	c.Contributors = contributors
	return nil
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
	for _, v := range c.Contributors {
		user := *(v.Author.Login)
		log.Printf("[CONTRIBUTORS] user %v", user)
		company, _ := c.GetCompanyOfUser(user)
		log.Printf("[COMPANY] company : %v", company)
		users = append(users, User{user, company, v.GetTotal()})
	}
	return users, nil
}

func (c *Contributors) GetCompanyOfUser(user string) (string, error) {
	ctx := context.Background()
	res, resp, err := c.Client.Users.Get(ctx, user)
	if err != nil {
		return "", fmt.Errorf("[users] failed to get company info.")
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("[contributors]: bad response status code %d", resp.StatusCode)
	}
	log.Printf("[contributors] email %v", res.GetCompany)
	return res.GetCompany(), nil
}
