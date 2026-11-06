// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryBatchListResponse JSON:API document.
type RumRecoveryBatchListResponse struct {
	// Batch requests ordered newest first. A filtered page can be empty while meta.next_cursor is present.
	Data []RumRecoveryBatchListData `json:"data"`
	// Batch-list pagination metadata.
	Meta RumRecoveryBatchListMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRecoveryBatchListResponse instantiates a new RumRecoveryBatchListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRecoveryBatchListResponse(data []RumRecoveryBatchListData, meta RumRecoveryBatchListMeta) *RumRecoveryBatchListResponse {
	this := RumRecoveryBatchListResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewRumRecoveryBatchListResponseWithDefaults instantiates a new RumRecoveryBatchListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRecoveryBatchListResponseWithDefaults() *RumRecoveryBatchListResponse {
	this := RumRecoveryBatchListResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *RumRecoveryBatchListResponse) GetData() []RumRecoveryBatchListData {
	if o == nil {
		var ret []RumRecoveryBatchListData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListResponse) GetDataOk() (*[]RumRecoveryBatchListData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *RumRecoveryBatchListResponse) SetData(v []RumRecoveryBatchListData) {
	o.Data = v
}

// GetMeta returns the Meta field value.
func (o *RumRecoveryBatchListResponse) GetMeta() RumRecoveryBatchListMeta {
	if o == nil {
		var ret RumRecoveryBatchListMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchListResponse) GetMetaOk() (*RumRecoveryBatchListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *RumRecoveryBatchListResponse) SetMeta(v RumRecoveryBatchListMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRecoveryBatchListResponse) MarshalJSON() ([]byte, error) {
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
func (o *RumRecoveryBatchListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]RumRecoveryBatchListData `json:"data"`
		Meta *RumRecoveryBatchListMeta   `json:"meta"`
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
