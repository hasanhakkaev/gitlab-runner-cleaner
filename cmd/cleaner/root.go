package cmd

import (
	"fmt"
	"os"

	config "github.com/hasanhakkaev/gitlab-runner-cleaner/internal/config"
	"github.com/hasanhakkaev/gitlab-runner-cleaner/internal/gitlab"
	"github.com/olekukonko/tablewriter"
)

// Execute is the entry point for the cleaner command
func Execute() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}
	runners, err := gitlab.GetRunners(cfg.RunnerStatus, 1)
	if err != nil {
		fmt.Println(err)
	}
	if runners == nil {
		fmt.Println("No runners found")
		os.Exit(0)
	}
	fmt.Printf("Runners found:%d\n", len(runners))
	fmt.Printf("RunnerStatus: %s\n", cfg.RunnerStatus)
	fmt.Printf("DryRun: %v\n", cfg.DryRun)

	if cfg.DryRun {
		fmt.Println("Dry run enabled.. Printing runners to be deleted")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Status"})

		for _, runner := range runners {
			runnerInfo := []string{
				fmt.Sprintf("%d", runner.ID), runner.Description, runner.Status}
			table.Append(runnerInfo)
		}

		table.Render()

	} else {
		fmt.Println("Dry run disabled.. Deleting runners..")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Status"})

		for _, runner := range runners {
			runnerInfo := []string{
				fmt.Sprintf("%d", runner.ID), runner.Description, runner.Status}
			table.Append(runnerInfo)
		}

		table.Render()

		for _, runner := range runners {
			err := gitlab.DeleteRunner(runner.ID)
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("Runners have been deleted!")
	}

}
