//
// Author:: Salim Afiune Maya (<afiune@lacework.net>)
// Copyright:: Copyright 2020, Lacework Inc.
// License:: Apache License, Version 2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	// lqlCreateCmd represents the lql create command
	lqlCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "create an LQL query",
		Long:  `Create an LQL query.`,
		Args:  cobra.NoArgs,
		RunE:  createQuery,
	}
)

func init() {
	// add sub-commands to the lql command
	lqlCmd.AddCommand(lqlCreateCmd)

	setQuerySourceFlags(lqlCreateCmd)
}

func createQuery(cmd *cobra.Command, args []string) error {
	query, err := inputQuery(cmd, args)
	if err != nil {
		return errors.Wrap(err, "unable to create LQL query")
	}

	cli.Log.Debugw("creating LQL query", "query", query)
	create, err := cli.LwApi.LQL.CreateQuery(query)

	if err != nil {
		err = queryErrorCrumbs(query, err)
		return errors.Wrap(err, "unable to create LQL query")
	}
	if cli.JSONOutput() {
		return cli.OutputJSON(create.Data)
	}
	queryID := "unknown"
	if len(create.Data) > 0 {
		queryID = create.Data[0].ID
	}
	cli.OutputHuman(fmt.Sprintf("LQL query (%s) created successfully.\n", queryID))
	return nil
}