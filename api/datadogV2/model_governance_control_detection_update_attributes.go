// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlDetectionUpdateAttributes The attributes of a governance control detection that can be updated. Only the attributes present in the request are modified.
type GovernanceControlDetectionUpdateAttributes struct {
	// The handle of the team the detection is assigned to. Set to an empty string to clear the assignment.
	AssignedTeam *string `json:"assigned_team,omitempty"`
	// The UUID of the user the detection is assigned to. Set to an empty string to clear the assignment.
	AssignedTo *string `json:"assigned_to,omitempty"`
	// The timestamp after which the detection becomes eligible for mitigation. Used to defer mitigation to a later time.
	MitigateAfter *time.Time `json:"mitigate_after,omitempty"`
	// The new state to set for the detection. Set to `exception` to acknowledge the detection and exclude it from active counts, or `active` to reopen it.
	State *GovernanceControlDetectionUpdateState `json:"state,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceControlDetectionUpdateAttributes instantiates a new GovernanceControlDetectionUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceControlDetectionUpdateAttributes() *GovernanceControlDetectionUpdateAttributes {
	this := GovernanceControlDetectionUpdateAttributes{}
	return &this
}

// NewGovernanceControlDetectionUpdateAttributesWithDefaults instantiates a new GovernanceControlDetectionUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceControlDetectionUpdateAttributesWithDefaults() *GovernanceControlDetectionUpdateAttributes {
	this := GovernanceControlDetectionUpdateAttributes{}
	return &this
}

// GetAssignedTeam returns the AssignedTeam field value if set, zero value otherwise.
func (o *GovernanceControlDetectionUpdateAttributes) GetAssignedTeam() string {
	if o == nil || o.AssignedTeam == nil {
		var ret string
		return ret
	}
	return *o.AssignedTeam
}

// GetAssignedTeamOk returns a tuple with the AssignedTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionUpdateAttributes) GetAssignedTeamOk() (*string, bool) {
	if o == nil || o.AssignedTeam == nil {
		return nil, false
	}
	return o.AssignedTeam, true
}

// HasAssignedTeam returns a boolean if a field has been set.
func (o *GovernanceControlDetectionUpdateAttributes) HasAssignedTeam() bool {
	return o != nil && o.AssignedTeam != nil
}

// SetAssignedTeam gets a reference to the given string and assigns it to the AssignedTeam field.
func (o *GovernanceControlDetectionUpdateAttributes) SetAssignedTeam(v string) {
	o.AssignedTeam = &v
}

// GetAssignedTo returns the AssignedTo field value if set, zero value otherwise.
func (o *GovernanceControlDetectionUpdateAttributes) GetAssignedTo() string {
	if o == nil || o.AssignedTo == nil {
		var ret string
		return ret
	}
	return *o.AssignedTo
}

// GetAssignedToOk returns a tuple with the AssignedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionUpdateAttributes) GetAssignedToOk() (*string, bool) {
	if o == nil || o.AssignedTo == nil {
		return nil, false
	}
	return o.AssignedTo, true
}

// HasAssignedTo returns a boolean if a field has been set.
func (o *GovernanceControlDetectionUpdateAttributes) HasAssignedTo() bool {
	return o != nil && o.AssignedTo != nil
}

// SetAssignedTo gets a reference to the given string and assigns it to the AssignedTo field.
func (o *GovernanceControlDetectionUpdateAttributes) SetAssignedTo(v string) {
	o.AssignedTo = &v
}

// GetMitigateAfter returns the MitigateAfter field value if set, zero value otherwise.
func (o *GovernanceControlDetectionUpdateAttributes) GetMitigateAfter() time.Time {
	if o == nil || o.MitigateAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.MitigateAfter
}

// GetMitigateAfterOk returns a tuple with the MitigateAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionUpdateAttributes) GetMitigateAfterOk() (*time.Time, bool) {
	if o == nil || o.MitigateAfter == nil {
		return nil, false
	}
	return o.MitigateAfter, true
}

// HasMitigateAfter returns a boolean if a field has been set.
func (o *GovernanceControlDetectionUpdateAttributes) HasMitigateAfter() bool {
	return o != nil && o.MitigateAfter != nil
}

// SetMitigateAfter gets a reference to the given time.Time and assigns it to the MitigateAfter field.
func (o *GovernanceControlDetectionUpdateAttributes) SetMitigateAfter(v time.Time) {
	o.MitigateAfter = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GovernanceControlDetectionUpdateAttributes) GetState() GovernanceControlDetectionUpdateState {
	if o == nil || o.State == nil {
		var ret GovernanceControlDetectionUpdateState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlDetectionUpdateAttributes) GetStateOk() (*GovernanceControlDetectionUpdateState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GovernanceControlDetectionUpdateAttributes) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given GovernanceControlDetectionUpdateState and assigns it to the State field.
func (o *GovernanceControlDetectionUpdateAttributes) SetState(v GovernanceControlDetectionUpdateState) {
	o.State = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceControlDetectionUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssignedTeam != nil {
		toSerialize["assigned_team"] = o.AssignedTeam
	}
	if o.AssignedTo != nil {
		toSerialize["assigned_to"] = o.AssignedTo
	}
	if o.MitigateAfter != nil {
		if o.MitigateAfter.Nanosecond() == 0 {
			toSerialize["mitigate_after"] = o.MitigateAfter.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["mitigate_after"] = o.MitigateAfter.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceControlDetectionUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignedTeam  *string                                `json:"assigned_team,omitempty"`
		AssignedTo    *string                                `json:"assigned_to,omitempty"`
		MitigateAfter *time.Time                             `json:"mitigate_after,omitempty"`
		State         *GovernanceControlDetectionUpdateState `json:"state,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assigned_team", "assigned_to", "mitigate_after", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AssignedTeam = all.AssignedTeam
	o.AssignedTo = all.AssignedTo
	o.MitigateAfter = all.MitigateAfter
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
