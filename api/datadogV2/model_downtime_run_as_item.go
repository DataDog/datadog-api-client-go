// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeRunAsItem A set of principals allowed to act on behalf of the downtime.
type DowntimeRunAsItem struct {
	// List of principals allowed to act on behalf of the downtime.
	Principals []DowntimeRunAsPrincipal `json:"principals,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeRunAsItem instantiates a new DowntimeRunAsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeRunAsItem() *DowntimeRunAsItem {
	this := DowntimeRunAsItem{}
	return &this
}

// NewDowntimeRunAsItemWithDefaults instantiates a new DowntimeRunAsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeRunAsItemWithDefaults() *DowntimeRunAsItem {
	this := DowntimeRunAsItem{}
	return &this
}

// GetPrincipals returns the Principals field value if set, zero value otherwise.
func (o *DowntimeRunAsItem) GetPrincipals() []DowntimeRunAsPrincipal {
	if o == nil || o.Principals == nil {
		var ret []DowntimeRunAsPrincipal
		return ret
	}
	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeRunAsItem) GetPrincipalsOk() (*[]DowntimeRunAsPrincipal, bool) {
	if o == nil || o.Principals == nil {
		return nil, false
	}
	return &o.Principals, true
}

// HasPrincipals returns a boolean if a field has been set.
func (o *DowntimeRunAsItem) HasPrincipals() bool {
	return o != nil && o.Principals != nil
}

// SetPrincipals gets a reference to the given []DowntimeRunAsPrincipal and assigns it to the Principals field.
func (o *DowntimeRunAsItem) SetPrincipals(v []DowntimeRunAsPrincipal) {
	o.Principals = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeRunAsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Principals != nil {
		toSerialize["principals"] = o.Principals
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeRunAsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Principals []DowntimeRunAsPrincipal `json:"principals,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"principals"})
	} else {
		return err
	}
	o.Principals = all.Principals

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
