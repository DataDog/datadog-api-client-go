// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GithubAppUpdateHostname - Hostname of the GitHub Enterprise Server. Omit this field for github.com.
type GithubAppUpdateHostname struct {
	ActionConnectionLiteralValue               *string
	ActionConnectionAWSSecretsManagerReference *ActionConnectionAWSSecretsManagerReference

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ActionConnectionLiteralValueAsGithubAppUpdateHostname is a convenience function that returns string wrapped in GithubAppUpdateHostname.
func ActionConnectionLiteralValueAsGithubAppUpdateHostname(v *string) GithubAppUpdateHostname {
	return GithubAppUpdateHostname{ActionConnectionLiteralValue: v}
}

// ActionConnectionAWSSecretsManagerReferenceAsGithubAppUpdateHostname is a convenience function that returns ActionConnectionAWSSecretsManagerReference wrapped in GithubAppUpdateHostname.
func ActionConnectionAWSSecretsManagerReferenceAsGithubAppUpdateHostname(v *ActionConnectionAWSSecretsManagerReference) GithubAppUpdateHostname {
	return GithubAppUpdateHostname{ActionConnectionAWSSecretsManagerReference: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GithubAppUpdateHostname) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActionConnectionLiteralValue
	err = datadog.Unmarshal(data, &obj.ActionConnectionLiteralValue)
	if err == nil {
		if obj.ActionConnectionLiteralValue != nil {
			jsonActionConnectionLiteralValue, _ := datadog.Marshal(obj.ActionConnectionLiteralValue)
			if string(jsonActionConnectionLiteralValue) == "{}" { // empty struct
				obj.ActionConnectionLiteralValue = nil
			} else {
				match++
			}
		} else {
			obj.ActionConnectionLiteralValue = nil
		}
	} else {
		obj.ActionConnectionLiteralValue = nil
	}

	// try to unmarshal data into ActionConnectionAWSSecretsManagerReference
	err = datadog.Unmarshal(data, &obj.ActionConnectionAWSSecretsManagerReference)
	if err == nil {
		if obj.ActionConnectionAWSSecretsManagerReference != nil && obj.ActionConnectionAWSSecretsManagerReference.UnparsedObject == nil {
			jsonActionConnectionAWSSecretsManagerReference, _ := datadog.Marshal(obj.ActionConnectionAWSSecretsManagerReference)
			if string(jsonActionConnectionAWSSecretsManagerReference) == "{}" { // empty struct
				obj.ActionConnectionAWSSecretsManagerReference = nil
			} else {
				match++
			}
		} else {
			obj.ActionConnectionAWSSecretsManagerReference = nil
		}
	} else {
		obj.ActionConnectionAWSSecretsManagerReference = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ActionConnectionLiteralValue = nil
		obj.ActionConnectionAWSSecretsManagerReference = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GithubAppUpdateHostname) MarshalJSON() ([]byte, error) {
	if obj.ActionConnectionLiteralValue != nil {
		return datadog.Marshal(&obj.ActionConnectionLiteralValue)
	}

	if obj.ActionConnectionAWSSecretsManagerReference != nil {
		return datadog.Marshal(&obj.ActionConnectionAWSSecretsManagerReference)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GithubAppUpdateHostname) GetActualInstance() interface{} {
	if obj.ActionConnectionLiteralValue != nil {
		return obj.ActionConnectionLiteralValue
	}

	if obj.ActionConnectionAWSSecretsManagerReference != nil {
		return obj.ActionConnectionAWSSecretsManagerReference
	}

	// all schemas are nil
	return nil
}
