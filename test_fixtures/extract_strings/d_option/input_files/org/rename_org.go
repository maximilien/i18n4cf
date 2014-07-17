package organization

import (
	"fmt"
)

type RenameOrg struct {
}

func (command *RenameOrg) Metadata() CommandMetadata {
	return CommandMetadata{
		Name:        "rename-org",
		Description: "Rename an org",
		Usage:       "CF_NAME rename-org ORG NEW_ORG",
	}
}

func (cmd *RenameOrg) GetRequirements(err error) {
	err = errors.New("Incorrect Usage")
}

func (cmd *RenameOrg) Run() {
	var orgName, newName, username string
	cmd.ui.Say("Renaming org %s to %s as %s...", orgName, newName, username)
}
