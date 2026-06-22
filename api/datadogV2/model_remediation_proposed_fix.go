// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationProposedFix A recommendation-oriented summary of a candidate remediation.
type RemediationProposedFix struct {
	// Explanation of the proposed change and why it resolves the root cause. Treat as user-provided content and escape before display.
	Description *string `json:"description,omitempty"`
	// Whether the proposed fix can be reversed after execution.
	Reversible *bool `json:"reversible,omitempty"`
	// Estimated risk of a remediation step or proposed fix.
	Risk *RemediationRiskLevel `json:"risk,omitempty"`
	// Short title for the proposed fix.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationProposedFix instantiates a new RemediationProposedFix object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationProposedFix() *RemediationProposedFix {
	this := RemediationProposedFix{}
	return &this
}

// NewRemediationProposedFixWithDefaults instantiates a new RemediationProposedFix object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationProposedFixWithDefaults() *RemediationProposedFix {
	this := RemediationProposedFix{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RemediationProposedFix) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProposedFix) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RemediationProposedFix) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RemediationProposedFix) SetDescription(v string) {
	o.Description = &v
}

// GetReversible returns the Reversible field value if set, zero value otherwise.
func (o *RemediationProposedFix) GetReversible() bool {
	if o == nil || o.Reversible == nil {
		var ret bool
		return ret
	}
	return *o.Reversible
}

// GetReversibleOk returns a tuple with the Reversible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProposedFix) GetReversibleOk() (*bool, bool) {
	if o == nil || o.Reversible == nil {
		return nil, false
	}
	return o.Reversible, true
}

// HasReversible returns a boolean if a field has been set.
func (o *RemediationProposedFix) HasReversible() bool {
	return o != nil && o.Reversible != nil
}

// SetReversible gets a reference to the given bool and assigns it to the Reversible field.
func (o *RemediationProposedFix) SetReversible(v bool) {
	o.Reversible = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *RemediationProposedFix) GetRisk() RemediationRiskLevel {
	if o == nil || o.Risk == nil {
		var ret RemediationRiskLevel
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProposedFix) GetRiskOk() (*RemediationRiskLevel, bool) {
	if o == nil || o.Risk == nil {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *RemediationProposedFix) HasRisk() bool {
	return o != nil && o.Risk != nil
}

// SetRisk gets a reference to the given RemediationRiskLevel and assigns it to the Risk field.
func (o *RemediationProposedFix) SetRisk(v RemediationRiskLevel) {
	o.Risk = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RemediationProposedFix) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProposedFix) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RemediationProposedFix) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *RemediationProposedFix) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationProposedFix) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Reversible != nil {
		toSerialize["reversible"] = o.Reversible
	}
	if o.Risk != nil {
		toSerialize["risk"] = o.Risk
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationProposedFix) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string               `json:"description,omitempty"`
		Reversible  *bool                 `json:"reversible,omitempty"`
		Risk        *RemediationRiskLevel `json:"risk,omitempty"`
		Title       *string               `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "reversible", "risk", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Reversible = all.Reversible
	if all.Risk != nil && !all.Risk.IsValid() {
		hasInvalidField = true
	} else {
		o.Risk = all.Risk
	}
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
