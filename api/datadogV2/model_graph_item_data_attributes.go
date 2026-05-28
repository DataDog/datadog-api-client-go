// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GraphItemDataAttributes Attributes of a single grouping in the End User Device Monitoring graph response.
type GraphItemDataAttributes struct {
	// List of per-value counts for the grouping column.
	Counts []GraphItemDataAttributesCountsItems `json:"counts,omitempty"`
	// Identifier of the grouping column (for example, `os` or `type`).
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGraphItemDataAttributes instantiates a new GraphItemDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGraphItemDataAttributes() *GraphItemDataAttributes {
	this := GraphItemDataAttributes{}
	return &this
}

// NewGraphItemDataAttributesWithDefaults instantiates a new GraphItemDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGraphItemDataAttributesWithDefaults() *GraphItemDataAttributes {
	this := GraphItemDataAttributes{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GraphItemDataAttributes) GetCounts() []GraphItemDataAttributesCountsItems {
	if o == nil || o.Counts == nil {
		var ret []GraphItemDataAttributesCountsItems
		return ret
	}
	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphItemDataAttributes) GetCountsOk() (*[]GraphItemDataAttributesCountsItems, bool) {
	if o == nil || o.Counts == nil {
		return nil, false
	}
	return &o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GraphItemDataAttributes) HasCounts() bool {
	return o != nil && o.Counts != nil
}

// SetCounts gets a reference to the given []GraphItemDataAttributesCountsItems and assigns it to the Counts field.
func (o *GraphItemDataAttributes) SetCounts(v []GraphItemDataAttributesCountsItems) {
	o.Counts = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GraphItemDataAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphItemDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GraphItemDataAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GraphItemDataAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GraphItemDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Counts != nil {
		toSerialize["counts"] = o.Counts
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GraphItemDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Counts []GraphItemDataAttributesCountsItems `json:"counts,omitempty"`
		Type   *string                              `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"counts", "type"})
	} else {
		return err
	}
	o.Counts = all.Counts
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
