// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyResourcesResponse Monthly usage for host and non-host resources.
type MonthlyResourcesResponse struct {
	// Resource usage records. Empty when no matching resources are available.
	Data []MonthlyResource `json:"data,omitempty"`
	// Metadata for monthly resource usage.
	Meta *MonthlyResourcesMetadata `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonthlyResourcesResponse instantiates a new MonthlyResourcesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonthlyResourcesResponse() *MonthlyResourcesResponse {
	this := MonthlyResourcesResponse{}
	return &this
}

// NewMonthlyResourcesResponseWithDefaults instantiates a new MonthlyResourcesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonthlyResourcesResponseWithDefaults() *MonthlyResourcesResponse {
	this := MonthlyResourcesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MonthlyResourcesResponse) GetData() []MonthlyResource {
	if o == nil || o.Data == nil {
		var ret []MonthlyResource
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyResourcesResponse) GetDataOk() (*[]MonthlyResource, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MonthlyResourcesResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []MonthlyResource and assigns it to the Data field.
func (o *MonthlyResourcesResponse) SetData(v []MonthlyResource) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MonthlyResourcesResponse) GetMeta() MonthlyResourcesMetadata {
	if o == nil || o.Meta == nil {
		var ret MonthlyResourcesMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyResourcesResponse) GetMetaOk() (*MonthlyResourcesMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MonthlyResourcesResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given MonthlyResourcesMetadata and assigns it to the Meta field.
func (o *MonthlyResourcesResponse) SetMeta(v MonthlyResourcesMetadata) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonthlyResourcesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonthlyResourcesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []MonthlyResource         `json:"data,omitempty"`
		Meta *MonthlyResourcesMetadata `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
