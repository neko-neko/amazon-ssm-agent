// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Amazon Software License (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/asl/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package processor contains the methods for update ssm agent.
// It also provides methods for sendReply and updateInstanceInfo
package processor

import (
	"testing"
	"time"

	"github.com/aws/amazon-ssm-agent/agent/appconfig"
	"github.com/aws/amazon-ssm-agent/agent/log"
	messageService "github.com/aws/amazon-ssm-agent/agent/message/service"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ssmmds"
	"github.com/stretchr/testify/assert"
)

// stubSdkService is the stub for sdkService
type stubSdkService struct{}

func (s *stubSdkService) GetMessages(log log.T, instanceID string) (messages *ssmmds.GetMessagesOutput, err error) {
	return &ssmmds.GetMessagesOutput{}, nil
}

func (s *stubSdkService) AcknowledgeMessage(log log.T, messageID string) error {
	return nil
}

func (s *stubSdkService) SendReply(log log.T, messageID string, payload string) error {
	return nil
}

func (s *stubSdkService) FailMessage(log log.T, messageID string, failureType messageService.FailureType) error {
	return nil
}

func (s *stubSdkService) DeleteMessage(log log.T, messageID string) error {
	return nil
}

func (s *stubSdkService) Stop() {}

func stubNewMsgSvc(region string, endpoint string, creds *credentials.Credentials, connectionTimeout time.Duration) messageService.Service {
	return &stubSdkService{}
}

func TestSendReply(t *testing.T) {
	context := createUpdateContext(Installed)
	service := svcManager{}
	// setup
	getAppConfig = func(bool) (appconfig.T, error) {
		config := appconfig.T{}
		return config, nil
	}

	newMsgSvc = stubNewMsgSvc

	// action
	err := service.SendReply(logger, context.Current)

	// assert
	assert.NoError(t, err)
}

func TestSendReplyDeleteMessage(t *testing.T) {
	context := createUpdateContext(Installed)
	service := svcManager{}
	// setup
	getAppConfig = func(bool) (appconfig.T, error) {
		config := appconfig.T{}
		return config, nil
	}
	newMsgSvc = stubNewMsgSvc

	// action
	err := service.DeleteMessage(logger, context.Current)

	// assert
	assert.NoError(t, err)
}