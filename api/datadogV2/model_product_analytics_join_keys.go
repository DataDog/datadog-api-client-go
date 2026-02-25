// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJoinKeys Join key configuration for correlating events.
type ProductAnalyticsJoinKeys struct {
	// The primary join key facet.
	Primary *string `json:"primary,omitempty"`
	// Secondary join key facets.
	Secondary []string `json:"secondary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJoinKeys instantiates a new ProductAnalyticsJoinKeys object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJoinKeys() *ProductAnalyticsJoinKeys {
	this := ProductAnalyticsJoinKeys{}
	return &this
}

// NewProductAnalyticsJoinKeysWithDefaults instantiates a new ProductAnalyticsJoinKeys object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJoinKeysWithDefaults() *ProductAnalyticsJoinKeys {
	this := ProductAnalyticsJoinKeys{}
	return &this
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *ProductAnalyticsJoinKeys) GetPrimary() string {
	if o == nil || o.Primary == nil {
		var ret string
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJoinKeys) GetPrimaryOk() (*string, bool) {
	if o == nil || o.Primary == nil {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *ProductAnalyticsJoinKeys) HasPrimary() bool {
	return o != nil && o.Primary != nil
}

// SetPrimary gets a reference to the given string and assigns it to the Primary field.
func (o *ProductAnalyticsJoinKeys) SetPrimary(v string) {
	o.Primary = &v
}

// GetSecondary returns the Secondary field value if set, zero value otherwise.
func (o *ProductAnalyticsJoinKeys) GetSecondary() []string {
	if o == nil || o.Secondary == nil {
		var ret []string
		return ret
	}
	return o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJoinKeys) GetSecondaryOk() (*[]string, bool) {
	if o == nil || o.Secondary == nil {
		return nil, false
	}
	return &o.Secondary, true
}

// HasSecondary returns a boolean if a field has been set.
func (o *ProductAnalyticsJoinKeys) HasSecondary() bool {
	return o != nil && o.Secondary != nil
}

// SetSecondary gets a reference to the given []string and assigns it to the Secondary field.
func (o *ProductAnalyticsJoinKeys) SetSecondary(v []string) {
	o.Secondary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJoinKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Primary != nil {
		toSerialize["primary"] = o.Primary
	}
	if o.Secondary != nil {
		toSerialize["secondary"] = o.Secondary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJoinKeys) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Primary   *string  `json:"primary,omitempty"`
		Secondary []string `json:"secondary,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"primary", "secondary"})
	} else {
		return err
	}
	o.Primary = all.Primary
	o.Secondary = all.Secondary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
