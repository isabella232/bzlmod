// Copyright 2020 Google LLC
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
	"os"

	"github.com/bazelbuild/bzlmod/resolve"
	"github.com/spf13/cobra"
)

// resolveCmd represents the resolve command
var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolves dependencies and outputs a WORKSPACE file for Bazel",
	Long: `Sets up the current Bazel workspace by reading the MODULE.bazel file,
resolving transitive dependencies, and outputting a WORKSPACE file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("resolve called")
		if err := resolve.Resolve(".", nil); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(resolveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resolveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resolveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
