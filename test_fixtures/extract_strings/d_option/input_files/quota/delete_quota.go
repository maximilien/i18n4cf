package quota

import (
	"errors"
	"fmt"
)

type DeleteQuota struct {
}

func (command *DeleteQuota) Metadata() CommandMetadata {
	return CommandMetadata{
		Name:        "delete-quota",
		Description: "Delete a quota",
		Usage:       "CF_NAME delete-quota QUOTA [-f]",
		Flags:       fmt.Printf(BoolFlag{Name: "f", Usage: "Force deletion without confirmation"}),
	}
}

func (cmd *DeleteQuota) GetRequirements() error {
	err = errors.New("Incorrect Usage")
	fmt.Printf("delete-quota")
}

func (cmd *DeleteQuota) Run() {
	var quotaName, username string

	fmt.Printf("f")
	fmt.Printf("quota", quotaName)
	fmt.Printf("Deleting quota %s as %s...", quotaName, username)

	fmt.Printf("Quota %s does not exist", quotaName)
}
