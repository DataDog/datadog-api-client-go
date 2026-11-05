// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationKeyEvidence A piece of evidence the verdict rests on.
type SecurityAgentInvestigationKeyEvidence struct {
	// Category of the evidence.
	Category string `json:"category"`
	// What the investigation found.
	Finding string `json:"finding"`
	// Names of the steps the evidence comes from.
	SourceSteps []string `json:"source_steps,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationKeyEvidence instantiates a new SecurityAgentInvestigationKeyEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationKeyEvidence(category string, finding string) *SecurityAgentInvestigationKeyEvidence {
	this := SecurityAgentInvestigationKeyEvidence{}
	this.Category = category
	this.Finding = finding
	return &this
}

// NewSecurityAgentInvestigationKeyEvidenceWithDefaults instantiates a new SecurityAgentInvestigationKeyEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationKeyEvidenceWithDefaults() *SecurityAgentInvestigationKeyEvidence {
	this := SecurityAgentInvestigationKeyEvidence{}
	return &this
}

// GetCategory returns the Category field value.
func (o *SecurityAgentInvestigationKeyEvidence) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationKeyEvidence) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *SecurityAgentInvestigationKeyEvidence) SetCategory(v string) {
	o.Category = v
}

// GetFinding returns the Finding field value.
func (o *SecurityAgentInvestigationKeyEvidence) GetFinding() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Finding
}

// GetFindingOk returns a tuple with the Finding field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationKeyEvidence) GetFindingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Finding, true
}

// SetFinding sets field value.
func (o *SecurityAgentInvestigationKeyEvidence) SetFinding(v string) {
	o.Finding = v
}

// GetSourceSteps returns the SourceSteps field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationKeyEvidence) GetSourceSteps() []string {
	if o == nil || o.SourceSteps == nil {
		var ret []string
		return ret
	}
	return o.SourceSteps
}

// GetSourceStepsOk returns a tuple with the SourceSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationKeyEvidence) GetSourceStepsOk() (*[]string, bool) {
	if o == nil || o.SourceSteps == nil {
		return nil, false
	}
	return &o.SourceSteps, true
}

// HasSourceSteps returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationKeyEvidence) HasSourceSteps() bool {
	return o != nil && o.SourceSteps != nil
}

// SetSourceSteps gets a reference to the given []string and assigns it to the SourceSteps field.
func (o *SecurityAgentInvestigationKeyEvidence) SetSourceSteps(v []string) {
	o.SourceSteps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationKeyEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["finding"] = o.Finding
	if o.SourceSteps != nil {
		toSerialize["source_steps"] = o.SourceSteps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationKeyEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category    *string  `json:"category"`
		Finding     *string  `json:"finding"`
		SourceSteps []string `json:"source_steps,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Finding == nil {
		return fmt.Errorf("required field finding missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "finding", "source_steps"})
	} else {
		return err
	}
	o.Category = *all.Category
	o.Finding = *all.Finding
	o.SourceSteps = all.SourceSteps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
