// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExternalSecretsManager - External secrets manager configuration.
type ExternalSecretsManager struct {
	ExternalSecretsManagerOneOf *ExternalSecretsManagerOneOf

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ExternalSecretsManagerOneOfAsExternalSecretsManager is a convenience function that returns ExternalSecretsManagerOneOf wrapped in ExternalSecretsManager.
func ExternalSecretsManagerOneOfAsExternalSecretsManager(v *ExternalSecretsManagerOneOf) ExternalSecretsManager {
	return ExternalSecretsManager{ExternalSecretsManagerOneOf: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ExternalSecretsManager) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ExternalSecretsManagerOneOf
	err = datadog.Unmarshal(data, &obj.ExternalSecretsManagerOneOf)
	if err == nil {
		if obj.ExternalSecretsManagerOneOf != nil && obj.ExternalSecretsManagerOneOf.UnparsedObject == nil {
			jsonExternalSecretsManagerOneOf, _ := datadog.Marshal(obj.ExternalSecretsManagerOneOf)
			if string(jsonExternalSecretsManagerOneOf) == "{}" { // empty struct
				obj.ExternalSecretsManagerOneOf = nil
			} else {
				match++
			}
		} else {
			obj.ExternalSecretsManagerOneOf = nil
		}
	} else {
		obj.ExternalSecretsManagerOneOf = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ExternalSecretsManagerOneOf = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ExternalSecretsManager) MarshalJSON() ([]byte, error) {
	if obj.ExternalSecretsManagerOneOf != nil {
		return datadog.Marshal(&obj.ExternalSecretsManagerOneOf)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ExternalSecretsManager) GetActualInstance() interface{} {
	if obj.ExternalSecretsManagerOneOf != nil {
		return obj.ExternalSecretsManagerOneOf
	}

	// all schemas are nil
	return nil
}
