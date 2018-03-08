// Copyright 2018 The droller Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/repejota/droller"
)

var (
	verboseFlag bool
	versionFlag bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "droller",
	Short: "Roll docker containers",
	Long:  `Droller rolls docker container images on the system`,
	Run: func(cmd *cobra.Command, args []string) {

		if versionFlag {
			showVersion()
			os.Exit(2)
		}

		log.SetLevel(log.FatalLevel)
		if verboseFlag {
			log.SetLevel(log.DebugLevel)
		}

		droller.Main()
	},
}

// Execute adds all child commands to the root command and sets flags
// appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().BoolVarP(&versionFlag, "version", "V", false, "show version number")
	RootCmd.Flags().BoolVarP(&verboseFlag, "verbose", "v", false, "enable verbose mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Unimplemented
}

// snowVersion shows the program build and version information.
func showVersion() {
	Version := "0.0.0"
	Build := "buildid"
	versionInformation := fmt.Sprintf("droller v.%s-%s", Version, Build)
	fmt.Println(versionInformation)
}
