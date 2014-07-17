package organization

import (
	"fmt"
)

type ListOrgs struct {
}

func (command ListOrgs) Metadata() CommandMetadata {
	return CommandMetadata{
		Name:        "orgs",
		ShortName:   "o",
		Description: "List all orgs",
		Usage:       "CF_NAME orgs",
	}
}

func (cmd ListOrgs) Run() {
	var username string
	fmt.Printf("Getting orgs as %s...\n", username)

	table := []string{"name"}

	fmt.Printf("Failed fetching orgs.\n%s", table)
	fmt.Printf("No orgs found")
}
