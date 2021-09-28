// Copyright © 2021 Alibaba Group Holding Ltd.
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

package cmd

import (
	"fmt"

	"github.com/alibaba/sealer/check/service"

	"github.com/spf13/cobra"
)

type CheckArgs struct {
	Pre  bool
	Post bool
}

var checkArgs *CheckArgs

// pushCmd represents the push command
var checkCmd = &cobra.Command{
	Use:     "check",
	Short:   "check the state of cluster ",
	Example: `sealer check --pre or sealer check --post`,
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var checker service.CheckerService
		if checkArgs.Pre && checkArgs.Post {
			return fmt.Errorf("don't allow to set tow flags --pre and --post")
		}

		if checkArgs.Pre {
			checker = service.NewPreCheckerService()
		} else if checkArgs.Post {
			checker = service.NewPostCheckerService()
		} else {
			checker = service.NewDefaultCheckerService()
		}
		return checker.Run()
	},
}

func init() {
	checkArgs = &CheckArgs{}
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().BoolVar(&checkArgs.Pre, "pre", false, "Check dependencies before cluster creation")
	checkCmd.Flags().BoolVar(&checkArgs.Post, "post", false, "Check the status of the cluster after it is created")
}
