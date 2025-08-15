// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAppsDatastoreFromImportResponseDataAttributes The definition of `CreateAppsDatastoreFromImportResponseDataAttributes` object.
type CreateAppsDatastoreFromImportResponseDataAttributes struct {
	// The `attributes` `item_count`.
	ItemCount *int64 `json:"item_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateAppsDatastoreFromImportResponseDataAttributes instantiates a new CreateAppsDatastoreFromImportResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateAppsDatastoreFromImportResponseDataAttributes() *CreateAppsDatastoreFromImportResponseDataAttributes {
	this := CreateAppsDatastoreFromImportResponseDataAttributes{}
	return &this
}

// NewCreateAppsDatastoreFromImportResponseDataAttributesWithDefaults instantiates a new CreateAppsDatastoreFromImportResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateAppsDatastoreFromImportResponseDataAttributesWithDefaults() *CreateAppsDatastoreFromImportResponseDataAttributes {
	this := CreateAppsDatastoreFromImportResponseDataAttributes{}
	return &this
}

// GetItemCount returns the ItemCount field value if set, zero value otherwise.
func (o *CreateAppsDatastoreFromImportResponseDataAttributes) GetItemCount() int64 {
	if o == nil || o.ItemCount == nil {
		var ret int64
		return ret
	}
	return *o.ItemCount
}

// GetItemCountOk returns a tuple with the ItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAppsDatastoreFromImportResponseDataAttributes) GetItemCountOk() (*int64, bool) {
	if o == nil || o.ItemCount == nil {
		return nil, false
	}
	return o.ItemCount, true
}

// HasItemCount returns a boolean if a field has been set.
func (o *CreateAppsDatastoreFromImportResponseDataAttributes) HasItemCount() bool {
	return o != nil && o.ItemCount != nil
}

// SetItemCount gets a reference to the given int64 and assigns it to the ItemCount field.
func (o *CreateAppsDatastoreFromImportResponseDataAttributes) SetItemCount(v int64) {
	o.ItemCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateAppsDatastoreFromImportResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ItemCount != nil {
		toSerialize["item_count"] = o.ItemCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateAppsDatastoreFromImportResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ItemCount *int64 `json:"item_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"item_count"})
	} else {
		return err
	}
	o.ItemCount = all.ItemCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
