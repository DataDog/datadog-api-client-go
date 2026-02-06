// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssueFinding
type JiraIssueFinding struct {
	// Description of the finding.
	Description string `json:"description"`
	//
	Ids []JiraIssueFindingId `json:"ids"`
	// Number of impacted resources.
	Impacted *int64 `json:"impacted,omitempty"`
	// References for the finding.
	References string `json:"references"`
	// Remediation instructions for the finding.
	Remediation string `json:"remediation"`
	// The status of the finding.
	Severity FindingStatus `json:"severity"`
	// Title of the finding.
	Title string `json:"title"`
	// Type of the finding.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssueFinding instantiates a new JiraIssueFinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssueFinding(description string, ids []JiraIssueFindingId, references string, remediation string, severity FindingStatus, title string, typeVar string) *JiraIssueFinding {
	this := JiraIssueFinding{}
	this.Description = description
	this.Ids = ids
	this.References = references
	this.Remediation = remediation
	this.Severity = severity
	this.Title = title
	this.Type = typeVar
	return &this
}

// NewJiraIssueFindingWithDefaults instantiates a new JiraIssueFinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssueFindingWithDefaults() *JiraIssueFinding {
	this := JiraIssueFinding{}
	return &this
}

// GetDescription returns the Description field value.
func (o *JiraIssueFinding) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *JiraIssueFinding) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *JiraIssueFinding) SetDescription(v string) {
	o.Description = v
}

// GetIds returns the Ids field value.
func (o *JiraIssueFinding) GetIds() []JiraIssueFindingId {
	if o == nil {
		var ret []JiraIssueFindingId
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *JiraIssueFinding) GetIdsOk() (*[]JiraIssueFindingId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ids, true
}

// SetIds sets field value.
func (o *JiraIssueFinding) SetIds(v []JiraIssueFindingId) {
	o.Ids = v
}

// GetImpacted returns the Impacted field value if set, zero value otherwise.
func (o *JiraIssueFinding) GetImpacted() int64 {
	if o == nil || o.Impacted == nil {
		var ret int64
		return ret
	}
	return *o.Impacted
}

// GetImpactedOk returns a tuple with the Impacted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssueFinding) GetImpactedOk() (*int64, bool) {
	if o == nil || o.Impacted == nil {
		return nil, false
	}
	return o.Impacted, true
}

// HasImpacted returns a boolean if a field has been set.
func (o *JiraIssueFinding) HasImpacted() bool {
	return o != nil && o.Impacted != nil
}

// SetImpacted gets a reference to the given int64 and assigns it to the Impacted field.
func (o *JiraIssueFinding) SetImpacted(v int64) {
	o.Impacted = &v
}

// GetReferences returns the References field value.
func (o *JiraIssueFinding) GetReferences() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value
// and a boolean to check if the value has been set.
func (o *JiraIssueFinding) GetReferencesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.References, true
}

// SetReferences sets field value.
func (o *JiraIssueFinding) SetReferences(v string) {
	o.References = v
}

// GetRemediation returns the Remediation field value.
func (o *JiraIssueFinding) GetRemediation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Remediation
}

// GetRemediationOk returns a tuple with the Remediation field value
// and a boolean to check if the value has been set.
func (o *JiraIssueFinding) GetRemediationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remediation, true
}

// SetRemediation sets field value.
func (o *JiraIssueFinding) SetRemediation(v string) {
	o.Remediation = v
}

// GetSeverity returns the Severity field value.
func (o *JiraIssueFinding) GetSeverity() FindingStatus {
	if o == nil {
		var ret FindingStatus
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *JiraIssueFinding) GetSeverityOk() (*FindingStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *JiraIssueFinding) SetSeverity(v FindingStatus) {
	o.Severity = v
}

// GetTitle returns the Title field value.
func (o *JiraIssueFinding) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *JiraIssueFinding) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *JiraIssueFinding) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value.
func (o *JiraIssueFinding) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JiraIssueFinding) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *JiraIssueFinding) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssueFinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["ids"] = o.Ids
	if o.Impacted != nil {
		toSerialize["impacted"] = o.Impacted
	}
	toSerialize["references"] = o.References
	toSerialize["remediation"] = o.Remediation
	toSerialize["severity"] = o.Severity
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssueFinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string               `json:"description"`
		Ids         *[]JiraIssueFindingId `json:"ids"`
		Impacted    *int64                `json:"impacted,omitempty"`
		References  *string               `json:"references"`
		Remediation *string               `json:"remediation"`
		Severity    *FindingStatus        `json:"severity"`
		Title       *string               `json:"title"`
		Type        *string               `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Ids == nil {
		return fmt.Errorf("required field ids missing")
	}
	if all.References == nil {
		return fmt.Errorf("required field references missing")
	}
	if all.Remediation == nil {
		return fmt.Errorf("required field remediation missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "ids", "impacted", "references", "remediation", "severity", "title", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = *all.Description
	o.Ids = *all.Ids
	o.Impacted = all.Impacted
	o.References = *all.References
	o.Remediation = *all.Remediation
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.Title = *all.Title
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
