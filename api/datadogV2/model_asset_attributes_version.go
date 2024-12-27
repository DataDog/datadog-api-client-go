// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssetAttributesVersion Asset version.
type AssetAttributesVersion struct {
	// Asset first version.
	First *string `json:"first,omitempty"`
	// Asset last version.
	Last *string `json:"last,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAssetAttributesVersion instantiates a new AssetAttributesVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssetAttributesVersion() *AssetAttributesVersion {
	this := AssetAttributesVersion{}
	return &this
}

// NewAssetAttributesVersionWithDefaults instantiates a new AssetAttributesVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssetAttributesVersionWithDefaults() *AssetAttributesVersion {
	this := AssetAttributesVersion{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *AssetAttributesVersion) GetFirst() string {
	if o == nil || o.First == nil {
		var ret string
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAttributesVersion) GetFirstOk() (*string, bool) {
	if o == nil || o.First == nil {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *AssetAttributesVersion) HasFirst() bool {
	return o != nil && o.First != nil
}

// SetFirst gets a reference to the given string and assigns it to the First field.
func (o *AssetAttributesVersion) SetFirst(v string) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *AssetAttributesVersion) GetLast() string {
	if o == nil || o.Last == nil {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAttributesVersion) GetLastOk() (*string, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *AssetAttributesVersion) HasLast() bool {
	return o != nil && o.Last != nil
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *AssetAttributesVersion) SetLast(v string) {
	o.Last = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssetAttributesVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.First != nil {
		toSerialize["first"] = o.First
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AssetAttributesVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		First *string `json:"first,omitempty"`
		Last  *string `json:"last,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first", "last"})
	} else {
		return err
	}
	o.First = all.First
	o.Last = all.Last

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
