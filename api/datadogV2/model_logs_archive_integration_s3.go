// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveIntegrationS3 - The S3 Archive's integration destination. You must provide one of the following: `access_key_id` alone, or both `account_id` and `role_name` together.
type LogsArchiveIntegrationS3 struct {
	LogsArchiveIntegrationS3AccessKey *LogsArchiveIntegrationS3AccessKey
	LogsArchiveIntegrationS3Role      *LogsArchiveIntegrationS3Role

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// LogsArchiveIntegrationS3AccessKeyAsLogsArchiveIntegrationS3 is a convenience function that returns LogsArchiveIntegrationS3AccessKey wrapped in LogsArchiveIntegrationS3.
func LogsArchiveIntegrationS3AccessKeyAsLogsArchiveIntegrationS3(v *LogsArchiveIntegrationS3AccessKey) LogsArchiveIntegrationS3 {
	return LogsArchiveIntegrationS3{LogsArchiveIntegrationS3AccessKey: v}
}

// LogsArchiveIntegrationS3RoleAsLogsArchiveIntegrationS3 is a convenience function that returns LogsArchiveIntegrationS3Role wrapped in LogsArchiveIntegrationS3.
func LogsArchiveIntegrationS3RoleAsLogsArchiveIntegrationS3(v *LogsArchiveIntegrationS3Role) LogsArchiveIntegrationS3 {
	return LogsArchiveIntegrationS3{LogsArchiveIntegrationS3Role: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *LogsArchiveIntegrationS3) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogsArchiveIntegrationS3AccessKey
	err = datadog.Unmarshal(data, &obj.LogsArchiveIntegrationS3AccessKey)
	if err == nil {
		if obj.LogsArchiveIntegrationS3AccessKey != nil && obj.LogsArchiveIntegrationS3AccessKey.UnparsedObject == nil {
			jsonLogsArchiveIntegrationS3AccessKey, _ := datadog.Marshal(obj.LogsArchiveIntegrationS3AccessKey)
			if string(jsonLogsArchiveIntegrationS3AccessKey) == "{}" { // empty struct
				obj.LogsArchiveIntegrationS3AccessKey = nil
			} else {
				match++
			}
		} else {
			obj.LogsArchiveIntegrationS3AccessKey = nil
		}
	} else {
		obj.LogsArchiveIntegrationS3AccessKey = nil
	}

	// try to unmarshal data into LogsArchiveIntegrationS3Role
	err = datadog.Unmarshal(data, &obj.LogsArchiveIntegrationS3Role)
	if err == nil {
		if obj.LogsArchiveIntegrationS3Role != nil && obj.LogsArchiveIntegrationS3Role.UnparsedObject == nil {
			jsonLogsArchiveIntegrationS3Role, _ := datadog.Marshal(obj.LogsArchiveIntegrationS3Role)
			if string(jsonLogsArchiveIntegrationS3Role) == "{}" { // empty struct
				obj.LogsArchiveIntegrationS3Role = nil
			} else {
				match++
			}
		} else {
			obj.LogsArchiveIntegrationS3Role = nil
		}
	} else {
		obj.LogsArchiveIntegrationS3Role = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.LogsArchiveIntegrationS3AccessKey = nil
		obj.LogsArchiveIntegrationS3Role = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj LogsArchiveIntegrationS3) MarshalJSON() ([]byte, error) {
	if obj.LogsArchiveIntegrationS3AccessKey != nil {
		return datadog.Marshal(&obj.LogsArchiveIntegrationS3AccessKey)
	}

	if obj.LogsArchiveIntegrationS3Role != nil {
		return datadog.Marshal(&obj.LogsArchiveIntegrationS3Role)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *LogsArchiveIntegrationS3) GetActualInstance() interface{} {
	if obj.LogsArchiveIntegrationS3AccessKey != nil {
		return obj.LogsArchiveIntegrationS3AccessKey
	}

	if obj.LogsArchiveIntegrationS3Role != nil {
		return obj.LogsArchiveIntegrationS3Role
	}

	// all schemas are nil
	return nil
}
