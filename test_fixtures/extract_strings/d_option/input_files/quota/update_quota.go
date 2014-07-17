package quota

import (
"errors"
"fmt")

type updateQuota struct {
}

func (command *updateQuota) Metadata() CommandMetadata {
	return CommandMetadata{
		Name:        "update-quota",
		Description: "Update an existing resource quota",
		Usage:       "CF_NAME update-quota QUOTA [-m MEMORY] [-n NEW_NAME] [-r ROUTES] [-s SERVICE_INSTANCES] [--allow-paid-service-plans | --disallow-paid-service-plans]",
		Flags: []cli.Flag{
			flag_helpers.NewStringFlag("m", "Total amount of memory (e.g. 1024M, 1G, 10G)"),
			flag_helpers.NewStringFlag("n", "New name"),
			flag_helpers.NewIntFlag("r", "Total number of routes"),
			flag_helpers.NewIntFlag("s", "Total number of service instances"),
			cli.BoolFlag{Name: "allow-paid-service-plans", Usage: "Can provision instances of paid service plans"},
			cli.BoolFlag{Name: "disallow-paid-service-plans", Usage: "Cannot provision instances of paid service plans"},
		},
	}
}

func (cmd *updateQuota) GetRequirements() error {
	fmt.Printf("update-quota")
}

func (cmd *updateQuota) Run() {
	fmt.Printf("allow-paid-service-plans")
	fmt.Printf("disallow-paid-service-plans")
	fmt.Printf("Please choose either allow or disallow. Both flags are not permitted to be passed in the same command. ")
	if String("m") != "" {
		fmt.Printf(String("m"))
		fmt.Printf("update-quota")
	}

	if String("n") != "" {
		quota.Name = String("n")
	}

	if IsSet("s") {
		quota.ServicesLimit = Int("s")
	}

	if IsSet("r") {
		quota.RoutesLimit = Int("r")
	}

	var oldQuotaName, username string
	fmt.Printf("Updating quota %s as %s...", oldQuotaName, username)
}
