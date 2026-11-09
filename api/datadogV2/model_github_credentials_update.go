// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GithubCredentialsUpdate - Supported GitHub connection credentials for an update.
type GithubCredentialsUpdate struct {
	GithubAppUpdate                 *GithubAppUpdate
	GithubPersonalAccessTokenUpdate *GithubPersonalAccessTokenUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// GithubAppUpdateAsGithubCredentialsUpdate is a convenience function that returns GithubAppUpdate wrapped in GithubCredentialsUpdate.
func GithubAppUpdateAsGithubCredentialsUpdate(v *GithubAppUpdate) GithubCredentialsUpdate {
	return GithubCredentialsUpdate{GithubAppUpdate: v}
}

// GithubPersonalAccessTokenUpdateAsGithubCredentialsUpdate is a convenience function that returns GithubPersonalAccessTokenUpdate wrapped in GithubCredentialsUpdate.
func GithubPersonalAccessTokenUpdateAsGithubCredentialsUpdate(v *GithubPersonalAccessTokenUpdate) GithubCredentialsUpdate {
	return GithubCredentialsUpdate{GithubPersonalAccessTokenUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *GithubCredentialsUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GithubAppUpdate
	err = datadog.Unmarshal(data, &obj.GithubAppUpdate)
	if err == nil {
		if obj.GithubAppUpdate != nil && obj.GithubAppUpdate.UnparsedObject == nil {
			jsonGithubAppUpdate, _ := datadog.Marshal(obj.GithubAppUpdate)
			if string(jsonGithubAppUpdate) == "{}" { // empty struct
				obj.GithubAppUpdate = nil
			} else {
				match++
			}
		} else {
			obj.GithubAppUpdate = nil
		}
	} else {
		obj.GithubAppUpdate = nil
	}

	// try to unmarshal data into GithubPersonalAccessTokenUpdate
	err = datadog.Unmarshal(data, &obj.GithubPersonalAccessTokenUpdate)
	if err == nil {
		if obj.GithubPersonalAccessTokenUpdate != nil && obj.GithubPersonalAccessTokenUpdate.UnparsedObject == nil {
			jsonGithubPersonalAccessTokenUpdate, _ := datadog.Marshal(obj.GithubPersonalAccessTokenUpdate)
			if string(jsonGithubPersonalAccessTokenUpdate) == "{}" { // empty struct
				obj.GithubPersonalAccessTokenUpdate = nil
			} else {
				match++
			}
		} else {
			obj.GithubPersonalAccessTokenUpdate = nil
		}
	} else {
		obj.GithubPersonalAccessTokenUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.GithubAppUpdate = nil
		obj.GithubPersonalAccessTokenUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj GithubCredentialsUpdate) MarshalJSON() ([]byte, error) {
	if obj.GithubAppUpdate != nil {
		return datadog.Marshal(&obj.GithubAppUpdate)
	}

	if obj.GithubPersonalAccessTokenUpdate != nil {
		return datadog.Marshal(&obj.GithubPersonalAccessTokenUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *GithubCredentialsUpdate) GetActualInstance() interface{} {
	if obj.GithubAppUpdate != nil {
		return obj.GithubAppUpdate
	}

	if obj.GithubPersonalAccessTokenUpdate != nil {
		return obj.GithubPersonalAccessTokenUpdate
	}

	// all schemas are nil
	return nil
}
