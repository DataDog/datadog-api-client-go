// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationExecuteResponse Response payload for executing a remediation.
type RemediationExecuteResponse struct {
	// Echoes the investigation ID that the remediation runs against.
	InvestigationId string `json:"investigation_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationExecuteResponse instantiates a new RemediationExecuteResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationExecuteResponse(investigationId string) *RemediationExecuteResponse {
	this := RemediationExecuteResponse{}
	this.InvestigationId = investigationId
	return &this
}

// NewRemediationExecuteResponseWithDefaults instantiates a new RemediationExecuteResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationExecuteResponseWithDefaults() *RemediationExecuteResponse {
	this := RemediationExecuteResponse{}
	return &this
}

// GetInvestigationId returns the InvestigationId field value.
func (o *RemediationExecuteResponse) GetInvestigationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InvestigationId
}

// GetInvestigationIdOk returns a tuple with the InvestigationId field value
// and a boolean to check if the value has been set.
func (o *RemediationExecuteResponse) GetInvestigationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvestigationId, true
}

// SetInvestigationId sets field value.
func (o *RemediationExecuteResponse) SetInvestigationId(v string) {
	o.InvestigationId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationExecuteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["investigation_id"] = o.InvestigationId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationExecuteResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InvestigationId *string `json:"investigation_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InvestigationId == nil {
		return fmt.Errorf("required field investigation_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"investigation_id"})
	} else {
		return err
	}
	o.InvestigationId = *all.InvestigationId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
