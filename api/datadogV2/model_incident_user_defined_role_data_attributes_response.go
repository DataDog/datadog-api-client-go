// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedRoleDataAttributesResponse Attributes of an incident user-defined role.
type IncidentUserDefinedRoleDataAttributesResponse struct {
	// Timestamp when the role was created.
	Created time.Time `json:"created"`
	// A description of the user-defined role.
	Description datadog.NullableString `json:"description,omitempty"`
	// Timestamp when the role was last modified.
	Modified time.Time `json:"modified"`
	// The name of the user-defined role.
	Name string `json:"name"`
	// Policy configuration for a user-defined role.
	Policy IncidentUserDefinedRolePolicy `json:"policy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedRoleDataAttributesResponse instantiates a new IncidentUserDefinedRoleDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedRoleDataAttributesResponse(created time.Time, modified time.Time, name string, policy IncidentUserDefinedRolePolicy) *IncidentUserDefinedRoleDataAttributesResponse {
	this := IncidentUserDefinedRoleDataAttributesResponse{}
	this.Created = created
	this.Modified = modified
	this.Name = name
	this.Policy = policy
	return &this
}

// NewIncidentUserDefinedRoleDataAttributesResponseWithDefaults instantiates a new IncidentUserDefinedRoleDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedRoleDataAttributesResponseWithDefaults() *IncidentUserDefinedRoleDataAttributesResponse {
	this := IncidentUserDefinedRoleDataAttributesResponse{}
	return &this
}

// GetCreated returns the Created field value.
func (o *IncidentUserDefinedRoleDataAttributesResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataAttributesResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentUserDefinedRoleDataAttributesResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentUserDefinedRoleDataAttributesResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentUserDefinedRoleDataAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IncidentUserDefinedRoleDataAttributesResponse) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *IncidentUserDefinedRoleDataAttributesResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *IncidentUserDefinedRoleDataAttributesResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *IncidentUserDefinedRoleDataAttributesResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetModified returns the Modified field value.
func (o *IncidentUserDefinedRoleDataAttributesResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataAttributesResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentUserDefinedRoleDataAttributesResponse) SetModified(v time.Time) {
	o.Modified = v
}

// GetName returns the Name field value.
func (o *IncidentUserDefinedRoleDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentUserDefinedRoleDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetPolicy returns the Policy field value.
func (o *IncidentUserDefinedRoleDataAttributesResponse) GetPolicy() IncidentUserDefinedRolePolicy {
	if o == nil {
		var ret IncidentUserDefinedRolePolicy
		return ret
	}
	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRoleDataAttributesResponse) GetPolicyOk() (*IncidentUserDefinedRolePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value.
func (o *IncidentUserDefinedRoleDataAttributesResponse) SetPolicy(v IncidentUserDefinedRolePolicy) {
	o.Policy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedRoleDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["policy"] = o.Policy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedRoleDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created     *time.Time                     `json:"created"`
		Description datadog.NullableString         `json:"description,omitempty"`
		Modified    *time.Time                     `json:"modified"`
		Name        *string                        `json:"name"`
		Policy      *IncidentUserDefinedRolePolicy `json:"policy"`
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
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Policy == nil {
		return fmt.Errorf("required field policy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "description", "modified", "name", "policy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Created = *all.Created
	o.Description = all.Description
	o.Modified = *all.Modified
	o.Name = *all.Name
	if all.Policy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Policy = *all.Policy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
