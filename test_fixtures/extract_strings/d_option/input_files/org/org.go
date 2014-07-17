package organization

import (
	"errors"
	"fmt"
)

type ShowOrg struct {
}

func (command *ShowOrg) Metadata() CommandMetadata {
	return CommandMetadata{
		Name:        "org",
		Description: "Show org info",
		Usage:       "CF_NAME org ORG",
	}
}

func (cmd *ShowOrg) GetRequirements(err error) {
	err = errors.New("Incorrect Usage")
}

func (cmd *ShowOrg) Run() {
	var orgName, username, domains, orgQuota, spaces string
	fmt.Printf("Getting info for org %s as %s...", orName, username)
	orgQuota := fmt.Sprintf("%s (%dM memory limit, %d routes, %d services, paid services %s)",
		fmt.Printf("  domains: %s", domains),
		fmt.Printf("  quota:   %s", orgQuota),
		fmt.Printf("  spaces:  %s", spaces, ", "))
}
