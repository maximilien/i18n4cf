package quota

import (
	"errors"
	"fmt"
)

type ListQuotas struct {
}

func (command *ListQuotas) Metadata() CommandMetadata {
	return CommandMetadata{
		Name:        "quotas",
		Description: "List available usage quotas",
		Usage:       "CF_NAME quotas",
	}
}

func (cmd *ListQuotas) Run() {
	fmt.Printf("Getting quotas as %s...", username)
	fmt.Printf("")
	table := []string{"name", "memory limit", "routes", "service instances", "paid service plans"}
	fmt.Printf(table)
}
