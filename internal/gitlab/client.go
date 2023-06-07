package gitlab

import (
	"fmt"

	config "github.com/hasanhakkaev/gitlab-runner-cleaner/internal/config"
	gitlab "github.com/xanzy/go-gitlab"
)

// New returns a new gitlab client
func New() (*gitlab.Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	client, err := gitlab.NewClient(cfg.Token, gitlab.WithBaseURL(cfg.BaseURL+"/api/v4"))
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetRunners returns a list of runners
func GetRunners(filter string, page int) ([]*gitlab.Runner, error) {
	totalPages := 1

	client, err := New()
	if err != nil {
		return nil, err
	}

	var allRunners []*gitlab.Runner

	for page <= totalPages {
		runners, resp, err := client.Runners.ListAllRunners(&gitlab.ListRunnersOptions{
			Status: gitlab.String(filter),
			ListOptions: gitlab.ListOptions{
				Page:    page,
				PerPage: 100,
			},
		})
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		allRunners = append(allRunners, runners...)
		totalPages = resp.TotalPages
		page++

	}

	return allRunners, nil
}

// DeleteRunner deletes a runner
func DeleteRunner(id int) error {
	client, err := New()
	if err != nil {
		return err
	}

	fmt.Printf("Deleting runner with id: %d\n", id)
	_, err = client.Runners.DeleteRegisteredRunnerByID(id)
	if err != nil {
		return err
	}

	return nil
}
