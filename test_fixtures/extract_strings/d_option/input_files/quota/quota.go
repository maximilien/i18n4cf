package quota

import (
	"errors"
	"fmt"
)

type showQuota struct {
}

func (command *showQuota) Metadata() CommandMetadata {
	return CommandMetadata{
		Name:        "quota",
		Usage:       "CF_NAME quota QUOTA",
		Description: "Show quota info",
	}
}

func (cmd *showQuota) GetRequirements() error {
	fmt.Printf("quotas")
}

func (cmd *showQuota) Run(context *cli.Context) {
	fmt.Printf("Getting quota %s info as %s...", quotaName, username)
	table := []string{"", ""}
	table.Add([]string{"Memory", formatters.ByteSize(quota.MemoryLimit * formatters.MEGABYTE)})
	table.Add([]string{"Routes", fmt.Sprintf("%d", quota.RoutesLimit)})
	table.Add([]string{"Services", fmt.Sprintf("%d", quota.ServicesLimit)})
	table.Add([]string{"Paid service plans", formatters.Allowed(quota.NonBasicServicesAllowed)})
}
