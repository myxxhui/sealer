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

	"github.com/alibaba/sealer/common"
	"github.com/alibaba/sealer/image"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const (
	imageID           = "IMAGE ID"
	imageName         = "IMAGE NAME"
	imageCreate       = "CREATE"
	imageSize         = "SIZE"
	timeDefaultFormat = "2006-01-02 15:04:05"
)

var listCmd = &cobra.Command{
	Use:     "images",
	Short:   "list all cluster images",
	Args:    cobra.NoArgs,
	Example: `sealer images`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ims, err := image.NewImageMetadataService()
		if err != nil {
			return err
		}

		imageMetadataList, err := ims.List()
		if err != nil {
			return err
		}
		table := tablewriter.NewWriter(common.StdOut)
		table.SetHeader([]string{imageID, imageName, imageCreate, imageSize})
		for _, image := range imageMetadataList {
			create := image.CREATED.Format(timeDefaultFormat)
			size := formatSize(image.SIZE)
			table.Append([]string{image.ID, image.Name, create, size})
		}
		table.Render()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func formatSize(size int64) (Size string) {
	if size < 1024 {
		Size = fmt.Sprintf("%.2fB", float64(size)/float64(1))
	} else if size < (1024 * 1024) {
		Size = fmt.Sprintf("%.2fKB", float64(size)/float64(1024))
	} else if size < (1024 * 1024 * 1024) {
		Size = fmt.Sprintf("%.2fMB", float64(size)/float64(1024*1024))
	} else {
		Size = fmt.Sprintf("%.2fGB", float64(size)/float64(1024*1024*1024))
	}
	return
}
