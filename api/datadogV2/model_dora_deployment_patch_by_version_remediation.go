// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentPatchByVersionRemediation - Remediation details for the deployment. Optional, but required to calculate failed deployment recovery time. Specify either `id` or `version` to identify the remediation deployment, but not both.
type DORADeploymentPatchByVersionRemediation struct {
	DORADeploymentPatchByVersionRemediationByID      *DORADeploymentPatchByVersionRemediationByID
	DORADeploymentPatchByVersionRemediationByVersion *DORADeploymentPatchByVersionRemediationByVersion

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// DORADeploymentPatchByVersionRemediationByIDAsDORADeploymentPatchByVersionRemediation is a convenience function that returns DORADeploymentPatchByVersionRemediationByID wrapped in DORADeploymentPatchByVersionRemediation.
func DORADeploymentPatchByVersionRemediationByIDAsDORADeploymentPatchByVersionRemediation(v *DORADeploymentPatchByVersionRemediationByID) DORADeploymentPatchByVersionRemediation {
	return DORADeploymentPatchByVersionRemediation{DORADeploymentPatchByVersionRemediationByID: v}
}

// DORADeploymentPatchByVersionRemediationByVersionAsDORADeploymentPatchByVersionRemediation is a convenience function that returns DORADeploymentPatchByVersionRemediationByVersion wrapped in DORADeploymentPatchByVersionRemediation.
func DORADeploymentPatchByVersionRemediationByVersionAsDORADeploymentPatchByVersionRemediation(v *DORADeploymentPatchByVersionRemediationByVersion) DORADeploymentPatchByVersionRemediation {
	return DORADeploymentPatchByVersionRemediation{DORADeploymentPatchByVersionRemediationByVersion: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DORADeploymentPatchByVersionRemediation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DORADeploymentPatchByVersionRemediationByID
	err = datadog.Unmarshal(data, &obj.DORADeploymentPatchByVersionRemediationByID)
	if err == nil {
		if obj.DORADeploymentPatchByVersionRemediationByID != nil && obj.DORADeploymentPatchByVersionRemediationByID.UnparsedObject == nil {
			jsonDORADeploymentPatchByVersionRemediationByID, _ := datadog.Marshal(obj.DORADeploymentPatchByVersionRemediationByID)
			if string(jsonDORADeploymentPatchByVersionRemediationByID) == "{}" { // empty struct
				obj.DORADeploymentPatchByVersionRemediationByID = nil
			} else {
				match++
			}
		} else {
			obj.DORADeploymentPatchByVersionRemediationByID = nil
		}
	} else {
		obj.DORADeploymentPatchByVersionRemediationByID = nil
	}

	// try to unmarshal data into DORADeploymentPatchByVersionRemediationByVersion
	err = datadog.Unmarshal(data, &obj.DORADeploymentPatchByVersionRemediationByVersion)
	if err == nil {
		if obj.DORADeploymentPatchByVersionRemediationByVersion != nil && obj.DORADeploymentPatchByVersionRemediationByVersion.UnparsedObject == nil {
			jsonDORADeploymentPatchByVersionRemediationByVersion, _ := datadog.Marshal(obj.DORADeploymentPatchByVersionRemediationByVersion)
			if string(jsonDORADeploymentPatchByVersionRemediationByVersion) == "{}" { // empty struct
				obj.DORADeploymentPatchByVersionRemediationByVersion = nil
			} else {
				match++
			}
		} else {
			obj.DORADeploymentPatchByVersionRemediationByVersion = nil
		}
	} else {
		obj.DORADeploymentPatchByVersionRemediationByVersion = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.DORADeploymentPatchByVersionRemediationByID = nil
		obj.DORADeploymentPatchByVersionRemediationByVersion = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DORADeploymentPatchByVersionRemediation) MarshalJSON() ([]byte, error) {
	if obj.DORADeploymentPatchByVersionRemediationByID != nil {
		return datadog.Marshal(&obj.DORADeploymentPatchByVersionRemediationByID)
	}

	if obj.DORADeploymentPatchByVersionRemediationByVersion != nil {
		return datadog.Marshal(&obj.DORADeploymentPatchByVersionRemediationByVersion)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DORADeploymentPatchByVersionRemediation) GetActualInstance() interface{} {
	if obj.DORADeploymentPatchByVersionRemediationByID != nil {
		return obj.DORADeploymentPatchByVersionRemediationByID
	}

	if obj.DORADeploymentPatchByVersionRemediationByVersion != nil {
		return obj.DORADeploymentPatchByVersionRemediationByVersion
	}

	// all schemas are nil
	return nil
}
