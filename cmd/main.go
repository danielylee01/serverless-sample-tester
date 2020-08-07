// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/GoogleCloudPlatform/serverless-sample-tester/internal/cmd"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "sst [sample-dir]",
		Short: "An end-to-end tester for GCP samples",
		Args:  cobra.ExactArgs(1),
		RunE:  cmd.Root,
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error ocurred in the execution of this program: %v\n", err)
	}
}