// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRoleAssignmentDataAttributesResponse Attributes of an incident role assignment.
type IncidentRoleAssignmentDataAttributesResponse struct {
	// Timestamp when the role assignment was created.
	Created time.Time `json:"created"`
	// Timestamp when the role assignment was last modified.
	Modified time.Time `json:"modified"`
	// The type of the responder.
	ResponderType string `json:"responder_type"`
	// The name of the assigned role.
	Role string `json:"role"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRoleAssignmentDataAttributesResponse instantiates a new IncidentRoleAssignmentDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRoleAssignmentDataAttributesResponse(created time.Time, modified time.Time, responderType string, role string) *IncidentRoleAssignmentDataAttributesResponse {
	this := IncidentRoleAssignmentDataAttributesResponse{}
	this.Created = created
	this.Modified = modified
	this.ResponderType = responderType
	this.Role = role
	return &this
}

// NewIncidentRoleAssignmentDataAttributesResponseWithDefaults instantiates a new IncidentRoleAssignmentDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRoleAssignmentDataAttributesResponseWithDefaults() *IncidentRoleAssignmentDataAttributesResponse {
	this := IncidentRoleAssignmentDataAttributesResponse{}
	return &this
}

// GetCreated returns the Created field value.
func (o *IncidentRoleAssignmentDataAttributesResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentRoleAssignmentDataAttributesResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentRoleAssignmentDataAttributesResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value.
func (o *IncidentRoleAssignmentDataAttributesResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentRoleAssignmentDataAttributesResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentRoleAssignmentDataAttributesResponse) SetModified(v time.Time) {
	o.Modified = v
}

// GetResponderType returns the ResponderType field value.
func (o *IncidentRoleAssignmentDataAttributesResponse) GetResponderType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResponderType
}

// GetResponderTypeOk returns a tuple with the ResponderType field value
// and a boolean to check if the value has been set.
func (o *IncidentRoleAssignmentDataAttributesResponse) GetResponderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponderType, true
}

// SetResponderType sets field value.
func (o *IncidentRoleAssignmentDataAttributesResponse) SetResponderType(v string) {
	o.ResponderType = v
}

// GetRole returns the Role field value.
func (o *IncidentRoleAssignmentDataAttributesResponse) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *IncidentRoleAssignmentDataAttributesResponse) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *IncidentRoleAssignmentDataAttributesResponse) SetRole(v string) {
	o.Role = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRoleAssignmentDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["responder_type"] = o.ResponderType
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRoleAssignmentDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created       *time.Time `json:"created"`
		Modified      *time.Time `json:"modified"`
		ResponderType *string    `json:"responder_type"`
		Role          *string    `json:"role"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	if all.ResponderType == nil {
		return fmt.Errorf("required field responder_type missing")
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "modified", "responder_type", "role"})
	} else {
		return err
	}
	o.Created = *all.Created
	o.Modified = *all.Modified
	o.ResponderType = *all.ResponderType
	o.Role = *all.Role

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
