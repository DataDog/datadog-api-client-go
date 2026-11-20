// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMArchiveConfigListResponse JSON:API document.
type RUMArchiveConfigListResponse struct {
	// Configurations in page order.
	Data []RUMArchiveConfig `json:"data"`
	// Pagination metadata.
	Meta RUMArchiveConfigMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMArchiveConfigListResponse instantiates a new RUMArchiveConfigListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMArchiveConfigListResponse(data []RUMArchiveConfig, meta RUMArchiveConfigMeta) *RUMArchiveConfigListResponse {
	this := RUMArchiveConfigListResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewRUMArchiveConfigListResponseWithDefaults instantiates a new RUMArchiveConfigListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMArchiveConfigListResponseWithDefaults() *RUMArchiveConfigListResponse {
	this := RUMArchiveConfigListResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *RUMArchiveConfigListResponse) GetData() []RUMArchiveConfig {
	if o == nil {
		var ret []RUMArchiveConfig
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RUMArchiveConfigListResponse) GetDataOk() (*[]RUMArchiveConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *RUMArchiveConfigListResponse) SetData(v []RUMArchiveConfig) {
	o.Data = v
}

// GetMeta returns the Meta field value.
func (o *RUMArchiveConfigListResponse) GetMeta() RUMArchiveConfigMeta {
	if o == nil {
		var ret RUMArchiveConfigMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *RUMArchiveConfigListResponse) GetMetaOk() (*RUMArchiveConfigMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *RUMArchiveConfigListResponse) SetMeta(v RUMArchiveConfigMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMArchiveConfigListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMArchiveConfigListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]RUMArchiveConfig   `json:"data"`
		Meta *RUMArchiveConfigMeta `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
