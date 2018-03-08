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

package droller

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

// DockerClient ...
type DockerClient struct {
	cli *client.Client
}

// NewDockerClient ...
func NewDockerClient() *DockerClient {
	log.Info("Using local docker server")
	client := &DockerClient{}
	return client
}

// Connect ...
func (d *DockerClient) Connect() {
	log.Debug("Conecting to server")
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal("Can't connect to docker server", err)
	}
	log.Debug("Conected to local docker server")
	d.cli = cli
}

// DisConnect ...
func (d *DockerClient) DisConnect() {
	log.Debug("Disconecting from local docker server")
	err := d.cli.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("Disconnected")
}

// Images ...
func (d *DockerClient) Images() ([]types.ImageSummary, error) {
	log.Debug("Get images")
	options := types.ImageListOptions{
		All: true,
	}
	images, err := d.cli.ImageList(context.Background(), options)
	if err != nil {
		log.Error(err)
	}
	return images, err
}

// ImageShortID ...
func ImageShortID(imageLongID string) string {
	parts := strings.Split(imageLongID, ":")
	return parts[1][:12]
}

// ImageRepositoryName ...
func ImageRepositoryName(repoTag string) string {
	parts := strings.Split(repoTag, ":")
	repository := string(parts[0])
	return repository
}

// ImageRepositoryTag ...
func ImageRepositoryTag(repoTag string) string {
	parts := strings.Split(repoTag, ":")
	tag := string(parts[1])
	return tag
}
