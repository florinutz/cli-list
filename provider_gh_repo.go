package list

import (
	"context"
	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
)

type RepoUsersProvider struct {
	client *github.Client

	owner string
	repo  string

	data [][]string

	logger *log.Entry
}

// NewRepoUsersProvider constructor
func NewRepoUsersProvider(client *github.Client, owner, repo string, logger *log.Entry) *RepoUsersProvider {
	ctx := context.Background()
	p := &RepoUsersProvider{client: client, owner: owner, repo: repo, logger: logger}
	p.FetchRepo(ctx)

	return p
}

// FetchData fetches data
func (p *RepoUsersProvider) FetchData(ctx context.Context, filters []CanFilter, sorters []CanSort) error {
	p.data = [][]string{}
	return nil
}

// FetchRepo fetches the repository corresponding to this provider
func (p *RepoUsersProvider) FetchRepo(ctx context.Context) {
	r, _, err := p.client.Repositories.Get(ctx, p.owner, p.repo)
	if err != nil {
		p.logger.WithError(err).Fatalln("problem fetching repo")
		return
	}
	p.logger.WithField("repo", r).Debug("found repo info")
}

// GetData implements the provider interface
func (p *RepoUsersProvider) GetData() (data [][]string) {
	if p == nil {
		panic("no data")
	}
	return p.data
}
