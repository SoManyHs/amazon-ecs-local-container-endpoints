// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Code generated by [agent/gogenerate/awssdk.go] DO NOT EDIT.

package ecstcs

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
)

type AckPublishHealth struct {
	_ struct{} `type:"structure"`

	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s AckPublishHealth) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AckPublishHealth) GoString() string {
	return s.String()
}

type AckPublishMetric struct {
	_ struct{} `type:"structure"`

	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s AckPublishMetric) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AckPublishMetric) GoString() string {
	return s.String()
}

type BadRequestException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s BadRequestException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BadRequestException) GoString() string {
	return s.String()
}

func newErrorBadRequestException(v protocol.ResponseMetadata) error {
	return &BadRequestException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *BadRequestException) Code() string {
	return "BadRequestException"
}

// Message returns the exception's message.
func (s *BadRequestException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *BadRequestException) OrigErr() error {
	return nil
}

func (s *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *BadRequestException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *BadRequestException) RequestID() string {
	return s.RespMetadata.RequestID
}

type CWStatsSet struct {
	_ struct{} `type:"structure"`

	Max *float64 `locationName:"max" type:"double"`

	Min *float64 `locationName:"min" type:"double"`

	SampleCount *int64 `locationName:"sampleCount" type:"integer"`

	Sum *float64 `locationName:"sum" type:"double"`
}

// String returns the string representation
func (s CWStatsSet) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CWStatsSet) GoString() string {
	return s.String()
}

type ContainerHealth struct {
	_ struct{} `type:"structure"`

	ContainerName *string `locationName:"containerName" type:"string"`

	HealthStatus *string `locationName:"healthStatus" type:"string" enum:"HealthStatus"`

	StatusSince *time.Time `locationName:"statusSince" type:"timestamp"`
}

// String returns the string representation
func (s ContainerHealth) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ContainerHealth) GoString() string {
	return s.String()
}

type ContainerMetric struct {
	_ struct{} `type:"structure"`

	ContainerArn *string `locationName:"containerArn" type:"string"`

	ContainerName *string `locationName:"containerName" type:"string"`

	CpuStatsSet *CWStatsSet `locationName:"cpuStatsSet" type:"structure"`

	MemoryStatsSet *CWStatsSet `locationName:"memoryStatsSet" type:"structure"`

	NetworkStatsSet *NetworkStatsSet `locationName:"networkStatsSet" type:"structure"`

	StorageStatsSet *StorageStatsSet `locationName:"storageStatsSet" type:"structure"`
}

// String returns the string representation
func (s ContainerMetric) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ContainerMetric) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ContainerMetric) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ContainerMetric"}
	if s.NetworkStatsSet != nil {
		if err := s.NetworkStatsSet.Validate(); err != nil {
			invalidParams.AddNested("NetworkStatsSet", err.(request.ErrInvalidParams))
		}
	}
	if s.StorageStatsSet != nil {
		if err := s.StorageStatsSet.Validate(); err != nil {
			invalidParams.AddNested("StorageStatsSet", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type HealthMetadata struct {
	_ struct{} `type:"structure"`

	Cluster *string `locationName:"cluster" type:"string"`

	ContainerInstance *string `locationName:"containerInstance" type:"string"`

	Fin *bool `locationName:"fin" type:"boolean"`

	MessageId *string `locationName:"messageId" type:"string"`
}

// String returns the string representation
func (s HealthMetadata) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HealthMetadata) GoString() string {
	return s.String()
}

type HeartbeatInput struct {
	_ struct{} `type:"structure"`

	Healthy *bool `locationName:"healthy" type:"boolean"`
}

// String returns the string representation
func (s HeartbeatInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HeartbeatInput) GoString() string {
	return s.String()
}

type HeartbeatMessage struct {
	_ struct{} `type:"structure"`

	Healthy *bool `locationName:"healthy" type:"boolean"`
}

// String returns the string representation
func (s HeartbeatMessage) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HeartbeatMessage) GoString() string {
	return s.String()
}

type HeartbeatOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s HeartbeatOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HeartbeatOutput) GoString() string {
	return s.String()
}

type InvalidParameterException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s InvalidParameterException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InvalidParameterException) GoString() string {
	return s.String()
}

func newErrorInvalidParameterException(v protocol.ResponseMetadata) error {
	return &InvalidParameterException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InvalidParameterException) Code() string {
	return "InvalidParameterException"
}

// Message returns the exception's message.
func (s *InvalidParameterException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InvalidParameterException) OrigErr() error {
	return nil
}

func (s *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InvalidParameterException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InvalidParameterException) RequestID() string {
	return s.RespMetadata.RequestID
}

type MetricsMetadata struct {
	_ struct{} `type:"structure"`

	Cluster *string `locationName:"cluster" type:"string"`

	ContainerInstance *string `locationName:"containerInstance" type:"string"`

	Fin *bool `locationName:"fin" type:"boolean"`

	Idle *bool `locationName:"idle" type:"boolean"`

	MessageId *string `locationName:"messageId" type:"string"`
}

// String returns the string representation
func (s MetricsMetadata) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MetricsMetadata) GoString() string {
	return s.String()
}

type NetworkStatsSet struct {
	_ struct{} `type:"structure"`

	RxBytes *ULongStatsSet `locationName:"rxBytes" type:"structure"`

	RxBytesPerSecond *UDoubleCWStatsSet `locationName:"rxBytesPerSecond" type:"structure"`

	RxDropped *ULongStatsSet `locationName:"rxDropped" type:"structure"`

	RxErrors *ULongStatsSet `locationName:"rxErrors" type:"structure"`

	RxPackets *ULongStatsSet `locationName:"rxPackets" type:"structure"`

	TxBytes *ULongStatsSet `locationName:"txBytes" type:"structure"`

	TxBytesPerSecond *UDoubleCWStatsSet `locationName:"txBytesPerSecond" type:"structure"`

	TxDropped *ULongStatsSet `locationName:"txDropped" type:"structure"`

	TxErrors *ULongStatsSet `locationName:"txErrors" type:"structure"`

	TxPackets *ULongStatsSet `locationName:"txPackets" type:"structure"`
}

// String returns the string representation
func (s NetworkStatsSet) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkStatsSet) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NetworkStatsSet) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "NetworkStatsSet"}
	if s.RxBytes != nil {
		if err := s.RxBytes.Validate(); err != nil {
			invalidParams.AddNested("RxBytes", err.(request.ErrInvalidParams))
		}
	}
	if s.RxBytesPerSecond != nil {
		if err := s.RxBytesPerSecond.Validate(); err != nil {
			invalidParams.AddNested("RxBytesPerSecond", err.(request.ErrInvalidParams))
		}
	}
	if s.RxDropped != nil {
		if err := s.RxDropped.Validate(); err != nil {
			invalidParams.AddNested("RxDropped", err.(request.ErrInvalidParams))
		}
	}
	if s.RxErrors != nil {
		if err := s.RxErrors.Validate(); err != nil {
			invalidParams.AddNested("RxErrors", err.(request.ErrInvalidParams))
		}
	}
	if s.RxPackets != nil {
		if err := s.RxPackets.Validate(); err != nil {
			invalidParams.AddNested("RxPackets", err.(request.ErrInvalidParams))
		}
	}
	if s.TxBytes != nil {
		if err := s.TxBytes.Validate(); err != nil {
			invalidParams.AddNested("TxBytes", err.(request.ErrInvalidParams))
		}
	}
	if s.TxBytesPerSecond != nil {
		if err := s.TxBytesPerSecond.Validate(); err != nil {
			invalidParams.AddNested("TxBytesPerSecond", err.(request.ErrInvalidParams))
		}
	}
	if s.TxDropped != nil {
		if err := s.TxDropped.Validate(); err != nil {
			invalidParams.AddNested("TxDropped", err.(request.ErrInvalidParams))
		}
	}
	if s.TxErrors != nil {
		if err := s.TxErrors.Validate(); err != nil {
			invalidParams.AddNested("TxErrors", err.(request.ErrInvalidParams))
		}
	}
	if s.TxPackets != nil {
		if err := s.TxPackets.Validate(); err != nil {
			invalidParams.AddNested("TxPackets", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PublishHealthInput struct {
	_ struct{} `type:"structure"`

	Metadata *HealthMetadata `locationName:"metadata" type:"structure"`

	Tasks []*TaskHealth `locationName:"tasks" type:"list"`

	Timestamp *time.Time `locationName:"timestamp" type:"timestamp"`
}

// String returns the string representation
func (s PublishHealthInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishHealthInput) GoString() string {
	return s.String()
}

type PublishHealthOutput struct {
	_ struct{} `type:"structure"`

	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s PublishHealthOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishHealthOutput) GoString() string {
	return s.String()
}

type PublishHealthRequest struct {
	_ struct{} `type:"structure"`

	Metadata *HealthMetadata `locationName:"metadata" type:"structure"`

	Tasks []*TaskHealth `locationName:"tasks" type:"list"`

	Timestamp *time.Time `locationName:"timestamp" type:"timestamp"`
}

// String returns the string representation
func (s PublishHealthRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishHealthRequest) GoString() string {
	return s.String()
}

type PublishMetricsInput struct {
	_ struct{} `type:"structure"`

	Metadata *MetricsMetadata `locationName:"metadata" type:"structure"`

	TaskMetrics []*TaskMetric `locationName:"taskMetrics" type:"list"`

	Timestamp *time.Time `locationName:"timestamp" type:"timestamp"`
}

// String returns the string representation
func (s PublishMetricsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishMetricsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PublishMetricsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PublishMetricsInput"}
	if s.TaskMetrics != nil {
		for i, v := range s.TaskMetrics {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TaskMetrics", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PublishMetricsOutput struct {
	_ struct{} `type:"structure"`

	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s PublishMetricsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishMetricsOutput) GoString() string {
	return s.String()
}

type PublishMetricsRequest struct {
	_ struct{} `type:"structure"`

	Metadata *MetricsMetadata `locationName:"metadata" type:"structure"`

	TaskMetrics []*TaskMetric `locationName:"taskMetrics" type:"list"`

	Timestamp *time.Time `locationName:"timestamp" type:"timestamp"`
}

// String returns the string representation
func (s PublishMetricsRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishMetricsRequest) GoString() string {
	return s.String()
}

type ResourceValidationException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s ResourceValidationException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceValidationException) GoString() string {
	return s.String()
}

func newErrorResourceValidationException(v protocol.ResponseMetadata) error {
	return &ResourceValidationException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ResourceValidationException) Code() string {
	return "ResourceValidationException"
}

// Message returns the exception's message.
func (s *ResourceValidationException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ResourceValidationException) OrigErr() error {
	return nil
}

func (s *ResourceValidationException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ResourceValidationException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ResourceValidationException) RequestID() string {
	return s.RespMetadata.RequestID
}

type ServerException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s ServerException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerException) GoString() string {
	return s.String()
}

func newErrorServerException(v protocol.ResponseMetadata) error {
	return &ServerException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ServerException) Code() string {
	return "ServerException"
}

// Message returns the exception's message.
func (s *ServerException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ServerException) OrigErr() error {
	return nil
}

func (s *ServerException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ServerException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ServerException) RequestID() string {
	return s.RespMetadata.RequestID
}

type StartTelemetrySessionInput struct {
	_ struct{} `type:"structure"`

	Cluster *string `locationName:"cluster" type:"string"`

	ContainerInstance *string `locationName:"containerInstance" type:"string"`
}

// String returns the string representation
func (s StartTelemetrySessionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartTelemetrySessionInput) GoString() string {
	return s.String()
}

type StartTelemetrySessionOutput struct {
	_ struct{} `type:"structure"`

	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s StartTelemetrySessionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartTelemetrySessionOutput) GoString() string {
	return s.String()
}

type StartTelemetrySessionRequest struct {
	_ struct{} `type:"structure"`

	Cluster *string `locationName:"cluster" type:"string"`

	ContainerInstance *string `locationName:"containerInstance" type:"string"`
}

// String returns the string representation
func (s StartTelemetrySessionRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartTelemetrySessionRequest) GoString() string {
	return s.String()
}

type StopTelemetrySessionMessage struct {
	_ struct{} `type:"structure"`

	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s StopTelemetrySessionMessage) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StopTelemetrySessionMessage) GoString() string {
	return s.String()
}

type StorageStatsSet struct {
	_ struct{} `type:"structure"`

	ReadSizeBytes *ULongStatsSet `locationName:"readSizeBytes" type:"structure"`

	WriteSizeBytes *ULongStatsSet `locationName:"writeSizeBytes" type:"structure"`
}

// String returns the string representation
func (s StorageStatsSet) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StorageStatsSet) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StorageStatsSet) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StorageStatsSet"}
	if s.ReadSizeBytes != nil {
		if err := s.ReadSizeBytes.Validate(); err != nil {
			invalidParams.AddNested("ReadSizeBytes", err.(request.ErrInvalidParams))
		}
	}
	if s.WriteSizeBytes != nil {
		if err := s.WriteSizeBytes.Validate(); err != nil {
			invalidParams.AddNested("WriteSizeBytes", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TaskHealth struct {
	_ struct{} `type:"structure"`

	Containers []*ContainerHealth `locationName:"containers" type:"list"`

	TaskArn *string `locationName:"taskArn" type:"string"`

	TaskDefinitionFamily *string `locationName:"taskDefinitionFamily" type:"string"`

	TaskDefinitionVersion *string `locationName:"taskDefinitionVersion" type:"string"`
}

// String returns the string representation
func (s TaskHealth) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TaskHealth) GoString() string {
	return s.String()
}

type TaskMetric struct {
	_ struct{} `type:"structure"`

	ContainerMetrics []*ContainerMetric `locationName:"containerMetrics" type:"list"`

	TaskArn *string `locationName:"taskArn" type:"string"`

	TaskDefinitionFamily *string `locationName:"taskDefinitionFamily" type:"string"`

	TaskDefinitionVersion *string `locationName:"taskDefinitionVersion" type:"string"`
}

// String returns the string representation
func (s TaskMetric) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TaskMetric) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TaskMetric) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "TaskMetric"}
	if s.ContainerMetrics != nil {
		for i, v := range s.ContainerMetrics {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ContainerMetrics", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UDoubleCWStatsSet struct {
	_ struct{} `type:"structure"`

	// Max is a required field
	Max *float64 `locationName:"max" type:"double" required:"true"`

	// Min is a required field
	Min *float64 `locationName:"min" type:"double" required:"true"`

	// SampleCount is a required field
	SampleCount *int64 `locationName:"sampleCount" type:"integer" required:"true"`

	// Sum is a required field
	Sum *float64 `locationName:"sum" type:"double" required:"true"`
}

// String returns the string representation
func (s UDoubleCWStatsSet) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UDoubleCWStatsSet) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UDoubleCWStatsSet) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UDoubleCWStatsSet"}
	if s.Max == nil {
		invalidParams.Add(request.NewErrParamRequired("Max"))
	}
	if s.Min == nil {
		invalidParams.Add(request.NewErrParamRequired("Min"))
	}
	if s.SampleCount == nil {
		invalidParams.Add(request.NewErrParamRequired("SampleCount"))
	}
	if s.Sum == nil {
		invalidParams.Add(request.NewErrParamRequired("Sum"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ULongStatsSet struct {
	_ struct{} `type:"structure"`

	// Max is a required field
	Max *int64 `locationName:"max" type:"long" required:"true"`

	// Min is a required field
	Min *int64 `locationName:"min" type:"long" required:"true"`

	OverflowMax *int64 `locationName:"overflowMax" type:"long"`

	OverflowMin *int64 `locationName:"overflowMin" type:"long"`

	OverflowSum *int64 `locationName:"overflowSum" type:"long"`

	// SampleCount is a required field
	SampleCount *int64 `locationName:"sampleCount" type:"long" required:"true"`

	// Sum is a required field
	Sum *int64 `locationName:"sum" type:"long" required:"true"`
}

// String returns the string representation
func (s ULongStatsSet) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ULongStatsSet) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ULongStatsSet) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ULongStatsSet"}
	if s.Max == nil {
		invalidParams.Add(request.NewErrParamRequired("Max"))
	}
	if s.Min == nil {
		invalidParams.Add(request.NewErrParamRequired("Min"))
	}
	if s.SampleCount == nil {
		invalidParams.Add(request.NewErrParamRequired("SampleCount"))
	}
	if s.Sum == nil {
		invalidParams.Add(request.NewErrParamRequired("Sum"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
