// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentPatchByVersionRemediationByVersion Remediation details identified by the version of the remediation deployment, matched against the same service and environment as the failed deployment.
type DORADeploymentPatchByVersionRemediationByVersion struct {
	// The type of remediation action taken. Required when the failed deployment must be linked to a remediation deployment.
	Type DORADeploymentPatchRemediationType `json:"type"`
	// The version of the remediation deployment.
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDORADeploymentPatchByVersionRemediationByVersion instantiates a new DORADeploymentPatchByVersionRemediationByVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORADeploymentPatchByVersionRemediationByVersion(typeVar DORADeploymentPatchRemediationType, version string) *DORADeploymentPatchByVersionRemediationByVersion {
	this := DORADeploymentPatchByVersionRemediationByVersion{}
	this.Type = typeVar
	this.Version = version
	return &this
}

// NewDORADeploymentPatchByVersionRemediationByVersionWithDefaults instantiates a new DORADeploymentPatchByVersionRemediationByVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORADeploymentPatchByVersionRemediationByVersionWithDefaults() *DORADeploymentPatchByVersionRemediationByVersion {
	this := DORADeploymentPatchByVersionRemediationByVersion{}
	return &this
}

// GetType returns the Type field value.
func (o *DORADeploymentPatchByVersionRemediationByVersion) GetType() DORADeploymentPatchRemediationType {
	if o == nil {
		var ret DORADeploymentPatchRemediationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentPatchByVersionRemediationByVersion) GetTypeOk() (*DORADeploymentPatchRemediationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DORADeploymentPatchByVersionRemediationByVersion) SetType(v DORADeploymentPatchRemediationType) {
	o.Type = v
}

// GetVersion returns the Version field value.
func (o *DORADeploymentPatchByVersionRemediationByVersion) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentPatchByVersionRemediationByVersion) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *DORADeploymentPatchByVersionRemediationByVersion) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORADeploymentPatchByVersionRemediationByVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORADeploymentPatchByVersionRemediationByVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type    *DORADeploymentPatchRemediationType `json:"type"`
		Version *string                             `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Version = *all.Version

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
