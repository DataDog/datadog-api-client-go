// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationMitigationSuggestion A suggested mitigation.
type SecurityAgentInvestigationMitigationSuggestion struct {
	// Actions that apply the mitigation.
	Actions []SecurityAgentInvestigationMitigationAction `json:"actions,omitempty"`
	// ID of the mitigation.
	Id string `json:"id"`
	// Why the mitigation is suggested.
	Rationale string `json:"rationale"`
	// Whether the mitigation is recommended.
	Recommended bool `json:"recommended"`
	// What the mitigation applies to.
	ScopeNote *string `json:"scope_note,omitempty"`
	// Order in which to apply the mitigation, starting at 1.
	Step int64 `json:"step"`
	// Title of the mitigation.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationMitigationSuggestion instantiates a new SecurityAgentInvestigationMitigationSuggestion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationMitigationSuggestion(id string, rationale string, recommended bool, step int64, title string) *SecurityAgentInvestigationMitigationSuggestion {
	this := SecurityAgentInvestigationMitigationSuggestion{}
	this.Id = id
	this.Rationale = rationale
	this.Recommended = recommended
	this.Step = step
	this.Title = title
	return &this
}

// NewSecurityAgentInvestigationMitigationSuggestionWithDefaults instantiates a new SecurityAgentInvestigationMitigationSuggestion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationMitigationSuggestionWithDefaults() *SecurityAgentInvestigationMitigationSuggestion {
	this := SecurityAgentInvestigationMitigationSuggestion{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetActions() []SecurityAgentInvestigationMitigationAction {
	if o == nil || o.Actions == nil {
		var ret []SecurityAgentInvestigationMitigationAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetActionsOk() (*[]SecurityAgentInvestigationMitigationAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return &o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationMitigationSuggestion) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given []SecurityAgentInvestigationMitigationAction and assigns it to the Actions field.
func (o *SecurityAgentInvestigationMitigationSuggestion) SetActions(v []SecurityAgentInvestigationMitigationAction) {
	o.Actions = v
}

// GetId returns the Id field value.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SecurityAgentInvestigationMitigationSuggestion) SetId(v string) {
	o.Id = v
}

// GetRationale returns the Rationale field value.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetRationale() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rationale
}

// GetRationaleOk returns a tuple with the Rationale field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetRationaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rationale, true
}

// SetRationale sets field value.
func (o *SecurityAgentInvestigationMitigationSuggestion) SetRationale(v string) {
	o.Rationale = v
}

// GetRecommended returns the Recommended field value.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetRecommended() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Recommended
}

// GetRecommendedOk returns a tuple with the Recommended field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetRecommendedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recommended, true
}

// SetRecommended sets field value.
func (o *SecurityAgentInvestigationMitigationSuggestion) SetRecommended(v bool) {
	o.Recommended = v
}

// GetScopeNote returns the ScopeNote field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetScopeNote() string {
	if o == nil || o.ScopeNote == nil {
		var ret string
		return ret
	}
	return *o.ScopeNote
}

// GetScopeNoteOk returns a tuple with the ScopeNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetScopeNoteOk() (*string, bool) {
	if o == nil || o.ScopeNote == nil {
		return nil, false
	}
	return o.ScopeNote, true
}

// HasScopeNote returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationMitigationSuggestion) HasScopeNote() bool {
	return o != nil && o.ScopeNote != nil
}

// SetScopeNote gets a reference to the given string and assigns it to the ScopeNote field.
func (o *SecurityAgentInvestigationMitigationSuggestion) SetScopeNote(v string) {
	o.ScopeNote = &v
}

// GetStep returns the Step field value.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetStep() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Step
}

// GetStepOk returns a tuple with the Step field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetStepOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Step, true
}

// SetStep sets field value.
func (o *SecurityAgentInvestigationMitigationSuggestion) SetStep(v int64) {
	o.Step = v
}

// GetTitle returns the Title field value.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationSuggestion) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *SecurityAgentInvestigationMitigationSuggestion) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationMitigationSuggestion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	toSerialize["id"] = o.Id
	toSerialize["rationale"] = o.Rationale
	toSerialize["recommended"] = o.Recommended
	if o.ScopeNote != nil {
		toSerialize["scope_note"] = o.ScopeNote
	}
	toSerialize["step"] = o.Step
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationMitigationSuggestion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actions     []SecurityAgentInvestigationMitigationAction `json:"actions,omitempty"`
		Id          *string                                      `json:"id"`
		Rationale   *string                                      `json:"rationale"`
		Recommended *bool                                        `json:"recommended"`
		ScopeNote   *string                                      `json:"scope_note,omitempty"`
		Step        *int64                                       `json:"step"`
		Title       *string                                      `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Rationale == nil {
		return fmt.Errorf("required field rationale missing")
	}
	if all.Recommended == nil {
		return fmt.Errorf("required field recommended missing")
	}
	if all.Step == nil {
		return fmt.Errorf("required field step missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actions", "id", "rationale", "recommended", "scope_note", "step", "title"})
	} else {
		return err
	}
	o.Actions = all.Actions
	o.Id = *all.Id
	o.Rationale = *all.Rationale
	o.Recommended = *all.Recommended
	o.ScopeNote = all.ScopeNote
	o.Step = *all.Step
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
