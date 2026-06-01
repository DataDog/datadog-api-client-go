// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentResponderDataAttributesResponse Attributes of an incident responder in a response.
type IncidentResponderDataAttributesResponse struct {
	// Timestamp when the responder was created.
	Created time.Time `json:"created"`
	// The external ID of the responder.
	ExternalId datadog.NullableString `json:"external_id,omitempty"`
	// The external source of the responder.
	ExternalSource datadog.NullableString `json:"external_source,omitempty"`
	// Whether this responder counts toward billing.
	IsBillable bool `json:"is_billable"`
	// Timestamp when the responder was last active.
	LastActive datadog.NullableTime `json:"last_active,omitempty"`
	// Additional metadata for the responder.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Timestamp when the responder was last modified.
	Modified time.Time `json:"modified"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentResponderDataAttributesResponse instantiates a new IncidentResponderDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentResponderDataAttributesResponse(created time.Time, isBillable bool, modified time.Time) *IncidentResponderDataAttributesResponse {
	this := IncidentResponderDataAttributesResponse{}
	this.Created = created
	this.IsBillable = isBillable
	this.Modified = modified
	return &this
}

// NewIncidentResponderDataAttributesResponseWithDefaults instantiates a new IncidentResponderDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentResponderDataAttributesResponseWithDefaults() *IncidentResponderDataAttributesResponse {
	this := IncidentResponderDataAttributesResponse{}
	return &this
}

// GetCreated returns the Created field value.
func (o *IncidentResponderDataAttributesResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentResponderDataAttributesResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentResponderDataAttributesResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponderDataAttributesResponse) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponderDataAttributesResponse) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *IncidentResponderDataAttributesResponse) HasExternalId() bool {
	return o != nil && o.ExternalId.IsSet()
}

// SetExternalId gets a reference to the given datadog.NullableString and assigns it to the ExternalId field.
func (o *IncidentResponderDataAttributesResponse) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}

// SetExternalIdNil sets the value for ExternalId to be an explicit nil.
func (o *IncidentResponderDataAttributesResponse) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil.
func (o *IncidentResponderDataAttributesResponse) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetExternalSource returns the ExternalSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponderDataAttributesResponse) GetExternalSource() string {
	if o == nil || o.ExternalSource.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalSource.Get()
}

// GetExternalSourceOk returns a tuple with the ExternalSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponderDataAttributesResponse) GetExternalSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalSource.Get(), o.ExternalSource.IsSet()
}

// HasExternalSource returns a boolean if a field has been set.
func (o *IncidentResponderDataAttributesResponse) HasExternalSource() bool {
	return o != nil && o.ExternalSource.IsSet()
}

// SetExternalSource gets a reference to the given datadog.NullableString and assigns it to the ExternalSource field.
func (o *IncidentResponderDataAttributesResponse) SetExternalSource(v string) {
	o.ExternalSource.Set(&v)
}

// SetExternalSourceNil sets the value for ExternalSource to be an explicit nil.
func (o *IncidentResponderDataAttributesResponse) SetExternalSourceNil() {
	o.ExternalSource.Set(nil)
}

// UnsetExternalSource ensures that no value is present for ExternalSource, not even an explicit nil.
func (o *IncidentResponderDataAttributesResponse) UnsetExternalSource() {
	o.ExternalSource.Unset()
}

// GetIsBillable returns the IsBillable field value.
func (o *IncidentResponderDataAttributesResponse) GetIsBillable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsBillable
}

// GetIsBillableOk returns a tuple with the IsBillable field value
// and a boolean to check if the value has been set.
func (o *IncidentResponderDataAttributesResponse) GetIsBillableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBillable, true
}

// SetIsBillable sets field value.
func (o *IncidentResponderDataAttributesResponse) SetIsBillable(v bool) {
	o.IsBillable = v
}

// GetLastActive returns the LastActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponderDataAttributesResponse) GetLastActive() time.Time {
	if o == nil || o.LastActive.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastActive.Get()
}

// GetLastActiveOk returns a tuple with the LastActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponderDataAttributesResponse) GetLastActiveOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastActive.Get(), o.LastActive.IsSet()
}

// HasLastActive returns a boolean if a field has been set.
func (o *IncidentResponderDataAttributesResponse) HasLastActive() bool {
	return o != nil && o.LastActive.IsSet()
}

// SetLastActive gets a reference to the given datadog.NullableTime and assigns it to the LastActive field.
func (o *IncidentResponderDataAttributesResponse) SetLastActive(v time.Time) {
	o.LastActive.Set(&v)
}

// SetLastActiveNil sets the value for LastActive to be an explicit nil.
func (o *IncidentResponderDataAttributesResponse) SetLastActiveNil() {
	o.LastActive.Set(nil)
}

// UnsetLastActive ensures that no value is present for LastActive, not even an explicit nil.
func (o *IncidentResponderDataAttributesResponse) UnsetLastActive() {
	o.LastActive.Unset()
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentResponderDataAttributesResponse) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentResponderDataAttributesResponse) GetMetaOk() (*map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return &o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentResponderDataAttributesResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *IncidentResponderDataAttributesResponse) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetModified returns the Modified field value.
func (o *IncidentResponderDataAttributesResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentResponderDataAttributesResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentResponderDataAttributesResponse) SetModified(v time.Time) {
	o.Modified = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentResponderDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ExternalId.IsSet() {
		toSerialize["external_id"] = o.ExternalId.Get()
	}
	if o.ExternalSource.IsSet() {
		toSerialize["external_source"] = o.ExternalSource.Get()
	}
	toSerialize["is_billable"] = o.IsBillable
	if o.LastActive.IsSet() {
		toSerialize["last_active"] = o.LastActive.Get()
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentResponderDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created        *time.Time             `json:"created"`
		ExternalId     datadog.NullableString `json:"external_id,omitempty"`
		ExternalSource datadog.NullableString `json:"external_source,omitempty"`
		IsBillable     *bool                  `json:"is_billable"`
		LastActive     datadog.NullableTime   `json:"last_active,omitempty"`
		Meta           map[string]interface{} `json:"meta,omitempty"`
		Modified       *time.Time             `json:"modified"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.IsBillable == nil {
		return fmt.Errorf("required field is_billable missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "external_id", "external_source", "is_billable", "last_active", "meta", "modified"})
	} else {
		return err
	}
	o.Created = *all.Created
	o.ExternalId = all.ExternalId
	o.ExternalSource = all.ExternalSource
	o.IsBillable = *all.IsBillable
	o.LastActive = all.LastActive
	o.Meta = all.Meta
	o.Modified = *all.Modified

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
