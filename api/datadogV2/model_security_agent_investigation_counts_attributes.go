// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationCountsAttributes The counts of each value of the field.
type SecurityAgentInvestigationCountsAttributes struct {
	// One bucket per value, sorted by ascending count.
	Counts []SecurityAgentInvestigationCount `json:"counts"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationCountsAttributes instantiates a new SecurityAgentInvestigationCountsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationCountsAttributes(counts []SecurityAgentInvestigationCount) *SecurityAgentInvestigationCountsAttributes {
	this := SecurityAgentInvestigationCountsAttributes{}
	this.Counts = counts
	return &this
}

// NewSecurityAgentInvestigationCountsAttributesWithDefaults instantiates a new SecurityAgentInvestigationCountsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationCountsAttributesWithDefaults() *SecurityAgentInvestigationCountsAttributes {
	this := SecurityAgentInvestigationCountsAttributes{}
	return &this
}

// GetCounts returns the Counts field value.
func (o *SecurityAgentInvestigationCountsAttributes) GetCounts() []SecurityAgentInvestigationCount {
	if o == nil {
		var ret []SecurityAgentInvestigationCount
		return ret
	}
	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationCountsAttributes) GetCountsOk() (*[]SecurityAgentInvestigationCount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Counts, true
}

// SetCounts sets field value.
func (o *SecurityAgentInvestigationCountsAttributes) SetCounts(v []SecurityAgentInvestigationCount) {
	o.Counts = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationCountsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["counts"] = o.Counts

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationCountsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Counts *[]SecurityAgentInvestigationCount `json:"counts"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Counts == nil {
		return fmt.Errorf("required field counts missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"counts"})
	} else {
		return err
	}
	o.Counts = *all.Counts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
