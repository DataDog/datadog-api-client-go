// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationMitigationAction An action that applies a mitigation.
type SecurityAgentInvestigationMitigationAction struct {
	// Whether the action has run.
	Applied bool `json:"applied"`
	// When the action ran.
	AppliedAt *time.Time `json:"applied_at,omitempty"`
	// Who ran the action.
	AppliedBy *string `json:"applied_by,omitempty"`
	// ID of the action in the Action Catalog.
	Fqn *string `json:"fqn,omitempty"`
	// ID of the action.
	Id string `json:"id"`
	// Name of the action in the Action Catalog.
	Name *string `json:"name,omitempty"`
	// The resource the action changes.
	Resource *string `json:"resource,omitempty"`
	// Integration that runs the action, such as `dd`, `aws`, or `okta`.
	Source string `json:"source"`
	// Whether the organization has a connection that can run the action.
	SourceConnected bool `json:"source_connected"`
	// What the action targets.
	Target *string `json:"target,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationMitigationAction instantiates a new SecurityAgentInvestigationMitigationAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationMitigationAction(applied bool, id string, source string, sourceConnected bool) *SecurityAgentInvestigationMitigationAction {
	this := SecurityAgentInvestigationMitigationAction{}
	this.Applied = applied
	this.Id = id
	this.Source = source
	this.SourceConnected = sourceConnected
	return &this
}

// NewSecurityAgentInvestigationMitigationActionWithDefaults instantiates a new SecurityAgentInvestigationMitigationAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationMitigationActionWithDefaults() *SecurityAgentInvestigationMitigationAction {
	this := SecurityAgentInvestigationMitigationAction{}
	return &this
}

// GetApplied returns the Applied field value.
func (o *SecurityAgentInvestigationMitigationAction) GetApplied() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Applied
}

// GetAppliedOk returns a tuple with the Applied field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationAction) GetAppliedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Applied, true
}

// SetApplied sets field value.
func (o *SecurityAgentInvestigationMitigationAction) SetApplied(v bool) {
	o.Applied = v
}

// GetAppliedAt returns the AppliedAt field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationMitigationAction) GetAppliedAt() time.Time {
	if o == nil || o.AppliedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AppliedAt
}

// GetAppliedAtOk returns a tuple with the AppliedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationAction) GetAppliedAtOk() (*time.Time, bool) {
	if o == nil || o.AppliedAt == nil {
		return nil, false
	}
	return o.AppliedAt, true
}

// HasAppliedAt returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationMitigationAction) HasAppliedAt() bool {
	return o != nil && o.AppliedAt != nil
}

// SetAppliedAt gets a reference to the given time.Time and assigns it to the AppliedAt field.
func (o *SecurityAgentInvestigationMitigationAction) SetAppliedAt(v time.Time) {
	o.AppliedAt = &v
}

// GetAppliedBy returns the AppliedBy field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationMitigationAction) GetAppliedBy() string {
	if o == nil || o.AppliedBy == nil {
		var ret string
		return ret
	}
	return *o.AppliedBy
}

// GetAppliedByOk returns a tuple with the AppliedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationAction) GetAppliedByOk() (*string, bool) {
	if o == nil || o.AppliedBy == nil {
		return nil, false
	}
	return o.AppliedBy, true
}

// HasAppliedBy returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationMitigationAction) HasAppliedBy() bool {
	return o != nil && o.AppliedBy != nil
}

// SetAppliedBy gets a reference to the given string and assigns it to the AppliedBy field.
func (o *SecurityAgentInvestigationMitigationAction) SetAppliedBy(v string) {
	o.AppliedBy = &v
}

// GetFqn returns the Fqn field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationMitigationAction) GetFqn() string {
	if o == nil || o.Fqn == nil {
		var ret string
		return ret
	}
	return *o.Fqn
}

// GetFqnOk returns a tuple with the Fqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationAction) GetFqnOk() (*string, bool) {
	if o == nil || o.Fqn == nil {
		return nil, false
	}
	return o.Fqn, true
}

// HasFqn returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationMitigationAction) HasFqn() bool {
	return o != nil && o.Fqn != nil
}

// SetFqn gets a reference to the given string and assigns it to the Fqn field.
func (o *SecurityAgentInvestigationMitigationAction) SetFqn(v string) {
	o.Fqn = &v
}

// GetId returns the Id field value.
func (o *SecurityAgentInvestigationMitigationAction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationAction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SecurityAgentInvestigationMitigationAction) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationMitigationAction) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationAction) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationMitigationAction) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityAgentInvestigationMitigationAction) SetName(v string) {
	o.Name = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationMitigationAction) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationAction) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationMitigationAction) HasResource() bool {
	return o != nil && o.Resource != nil
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *SecurityAgentInvestigationMitigationAction) SetResource(v string) {
	o.Resource = &v
}

// GetSource returns the Source field value.
func (o *SecurityAgentInvestigationMitigationAction) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationAction) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *SecurityAgentInvestigationMitigationAction) SetSource(v string) {
	o.Source = v
}

// GetSourceConnected returns the SourceConnected field value.
func (o *SecurityAgentInvestigationMitigationAction) GetSourceConnected() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SourceConnected
}

// GetSourceConnectedOk returns a tuple with the SourceConnected field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationAction) GetSourceConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceConnected, true
}

// SetSourceConnected sets field value.
func (o *SecurityAgentInvestigationMitigationAction) SetSourceConnected(v bool) {
	o.SourceConnected = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationMitigationAction) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationMitigationAction) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationMitigationAction) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *SecurityAgentInvestigationMitigationAction) SetTarget(v string) {
	o.Target = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationMitigationAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["applied"] = o.Applied
	if o.AppliedAt != nil {
		if o.AppliedAt.Nanosecond() == 0 {
			toSerialize["applied_at"] = o.AppliedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["applied_at"] = o.AppliedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.AppliedBy != nil {
		toSerialize["applied_by"] = o.AppliedBy
	}
	if o.Fqn != nil {
		toSerialize["fqn"] = o.Fqn
	}
	toSerialize["id"] = o.Id
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	toSerialize["source"] = o.Source
	toSerialize["source_connected"] = o.SourceConnected
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationMitigationAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Applied         *bool      `json:"applied"`
		AppliedAt       *time.Time `json:"applied_at,omitempty"`
		AppliedBy       *string    `json:"applied_by,omitempty"`
		Fqn             *string    `json:"fqn,omitempty"`
		Id              *string    `json:"id"`
		Name            *string    `json:"name,omitempty"`
		Resource        *string    `json:"resource,omitempty"`
		Source          *string    `json:"source"`
		SourceConnected *bool      `json:"source_connected"`
		Target          *string    `json:"target,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Applied == nil {
		return fmt.Errorf("required field applied missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.SourceConnected == nil {
		return fmt.Errorf("required field source_connected missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"applied", "applied_at", "applied_by", "fqn", "id", "name", "resource", "source", "source_connected", "target"})
	} else {
		return err
	}
	o.Applied = *all.Applied
	o.AppliedAt = all.AppliedAt
	o.AppliedBy = all.AppliedBy
	o.Fqn = all.Fqn
	o.Id = *all.Id
	o.Name = all.Name
	o.Resource = all.Resource
	o.Source = *all.Source
	o.SourceConnected = *all.SourceConnected
	o.Target = all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
