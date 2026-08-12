// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentPatchByVersionRemediationByID Remediation details identified by the ID of the remediation deployment.
type DORADeploymentPatchByVersionRemediationByID struct {
	// The ID of the remediation deployment.
	Id string `json:"id"`
	// The type of remediation action taken. Required when the failed deployment must be linked to a remediation deployment.
	Type DORADeploymentPatchRemediationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDORADeploymentPatchByVersionRemediationByID instantiates a new DORADeploymentPatchByVersionRemediationByID object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORADeploymentPatchByVersionRemediationByID(id string, typeVar DORADeploymentPatchRemediationType) *DORADeploymentPatchByVersionRemediationByID {
	this := DORADeploymentPatchByVersionRemediationByID{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewDORADeploymentPatchByVersionRemediationByIDWithDefaults instantiates a new DORADeploymentPatchByVersionRemediationByID object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORADeploymentPatchByVersionRemediationByIDWithDefaults() *DORADeploymentPatchByVersionRemediationByID {
	this := DORADeploymentPatchByVersionRemediationByID{}
	return &this
}

// GetId returns the Id field value.
func (o *DORADeploymentPatchByVersionRemediationByID) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentPatchByVersionRemediationByID) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DORADeploymentPatchByVersionRemediationByID) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *DORADeploymentPatchByVersionRemediationByID) GetType() DORADeploymentPatchRemediationType {
	if o == nil {
		var ret DORADeploymentPatchRemediationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DORADeploymentPatchByVersionRemediationByID) GetTypeOk() (*DORADeploymentPatchRemediationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DORADeploymentPatchByVersionRemediationByID) SetType(v DORADeploymentPatchRemediationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORADeploymentPatchByVersionRemediationByID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORADeploymentPatchByVersionRemediationByID) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                             `json:"id"`
		Type *DORADeploymentPatchRemediationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
