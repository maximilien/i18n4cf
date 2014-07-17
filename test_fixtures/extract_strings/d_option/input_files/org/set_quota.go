package organization

import (
	"fmt"
)

type SetQuota struct {
}

func (command *SetQuota) Metadata() CommandMetadata {
	return CommandMetadata{
		Name:        "set-quota",
		Description: "Assign a quota to an org",
		Usage: "CF_NAME set-quota ORG QUOTA\n\n" +
			"TIP:\n" +
			"   View allowable quotas with 'CF_NAME quotas'",
	}
}

func (cmd *SetQuota) GetRequirements(err error) {
	err = errors.New("Incorrect Usage")
}

func (cmd *SetQuota) Run() {
	var quotaName, orgName, username string
	cmd.ui.Say("Setting quota %s to org %s as %s...", quotaName, orgName, username)
}
