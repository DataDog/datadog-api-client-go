// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGetInvestigationResponse Response payload for getting a single investigation.
type RemediationGetInvestigationResponse struct {
	// A single ECS remediation investigation: a detected issue together with its proposed plan, history, and ECS workload metadata.
	Investigation RemediationInvestigation `json:"investigation"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationGetInvestigationResponse instantiates a new RemediationGetInvestigationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationGetInvestigationResponse(investigation RemediationInvestigation) *RemediationGetInvestigationResponse {
	this := RemediationGetInvestigationResponse{}
	this.Investigation = investigation
	return &this
}

// NewRemediationGetInvestigationResponseWithDefaults instantiates a new RemediationGetInvestigationResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationGetInvestigationResponseWithDefaults() *RemediationGetInvestigationResponse {
	this := RemediationGetInvestigationResponse{}
	return &this
}

// GetInvestigation returns the Investigation field value.
func (o *RemediationGetInvestigationResponse) GetInvestigation() RemediationInvestigation {
	if o == nil {
		var ret RemediationInvestigation
		return ret
	}
	return o.Investigation
}

// GetInvestigationOk returns a tuple with the Investigation field value
// and a boolean to check if the value has been set.
func (o *RemediationGetInvestigationResponse) GetInvestigationOk() (*RemediationInvestigation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Investigation, true
}

// SetInvestigation sets field value.
func (o *RemediationGetInvestigationResponse) SetInvestigation(v RemediationInvestigation) {
	o.Investigation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationGetInvestigationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["investigation"] = o.Investigation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationGetInvestigationResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Investigation *RemediationInvestigation `json:"investigation"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Investigation == nil {
		return fmt.Errorf("required field investigation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"investigation"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Investigation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Investigation = *all.Investigation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
