// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationSignal Summary of the signal the investigation is about.
type SecurityAgentInvestigationSignal struct {
	// Severity of the signal: `info`, `low`, `medium`, `high`, or `critical`.
	Severity *string `json:"severity,omitempty"`
	// Sources of the signal.
	Sources []string `json:"sources,omitempty"`
	// MITRE ATT&CK tactic, from the signal's `tactic` tag.
	Tactic *string `json:"tactic,omitempty"`
	// Tags of the signal, written `key:value`.
	Tags []string `json:"tags,omitempty"`
	// MITRE ATT&CK technique, from the signal's `technique` tag.
	Technique *string `json:"technique,omitempty"`
	// Time the signal was generated.
	Timestamp *string `json:"timestamp,omitempty"`
	// Title of the signal.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationSignal instantiates a new SecurityAgentInvestigationSignal object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationSignal() *SecurityAgentInvestigationSignal {
	this := SecurityAgentInvestigationSignal{}
	return &this
}

// NewSecurityAgentInvestigationSignalWithDefaults instantiates a new SecurityAgentInvestigationSignal object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationSignalWithDefaults() *SecurityAgentInvestigationSignal {
	this := SecurityAgentInvestigationSignal{}
	return &this
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationSignal) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationSignal) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationSignal) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *SecurityAgentInvestigationSignal) SetSeverity(v string) {
	o.Severity = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationSignal) GetSources() []string {
	if o == nil || o.Sources == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationSignal) GetSourcesOk() (*[]string, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return &o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationSignal) HasSources() bool {
	return o != nil && o.Sources != nil
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *SecurityAgentInvestigationSignal) SetSources(v []string) {
	o.Sources = v
}

// GetTactic returns the Tactic field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationSignal) GetTactic() string {
	if o == nil || o.Tactic == nil {
		var ret string
		return ret
	}
	return *o.Tactic
}

// GetTacticOk returns a tuple with the Tactic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationSignal) GetTacticOk() (*string, bool) {
	if o == nil || o.Tactic == nil {
		return nil, false
	}
	return o.Tactic, true
}

// HasTactic returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationSignal) HasTactic() bool {
	return o != nil && o.Tactic != nil
}

// SetTactic gets a reference to the given string and assigns it to the Tactic field.
func (o *SecurityAgentInvestigationSignal) SetTactic(v string) {
	o.Tactic = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationSignal) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationSignal) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationSignal) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *SecurityAgentInvestigationSignal) SetTags(v []string) {
	o.Tags = v
}

// GetTechnique returns the Technique field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationSignal) GetTechnique() string {
	if o == nil || o.Technique == nil {
		var ret string
		return ret
	}
	return *o.Technique
}

// GetTechniqueOk returns a tuple with the Technique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationSignal) GetTechniqueOk() (*string, bool) {
	if o == nil || o.Technique == nil {
		return nil, false
	}
	return o.Technique, true
}

// HasTechnique returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationSignal) HasTechnique() bool {
	return o != nil && o.Technique != nil
}

// SetTechnique gets a reference to the given string and assigns it to the Technique field.
func (o *SecurityAgentInvestigationSignal) SetTechnique(v string) {
	o.Technique = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationSignal) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationSignal) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationSignal) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *SecurityAgentInvestigationSignal) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationSignal) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationSignal) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationSignal) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SecurityAgentInvestigationSignal) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationSignal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.Tactic != nil {
		toSerialize["tactic"] = o.Tactic
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Technique != nil {
		toSerialize["technique"] = o.Technique
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
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
func (o *SecurityAgentInvestigationSignal) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Severity  *string  `json:"severity,omitempty"`
		Sources   []string `json:"sources,omitempty"`
		Tactic    *string  `json:"tactic,omitempty"`
		Tags      []string `json:"tags,omitempty"`
		Technique *string  `json:"technique,omitempty"`
		Timestamp *string  `json:"timestamp,omitempty"`
		Title     *string  `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"severity", "sources", "tactic", "tags", "technique", "timestamp", "title"})
	} else {
		return err
	}
	o.Severity = all.Severity
	o.Sources = all.Sources
	o.Tactic = all.Tactic
	o.Tags = all.Tags
	o.Technique = all.Technique
	o.Timestamp = all.Timestamp
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
