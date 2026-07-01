// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BatchRowsQueryResponse Response object for a batch rows query against a reference table.
type BatchRowsQueryResponse struct {
	// Data object for a batch rows query response.
	Data *BatchRowsQueryResponseData `json:"data,omitempty"`
	// Full row resources matching the query, included alongside the relationship references in `data`.
	Included []TableRowResourceData `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBatchRowsQueryResponse instantiates a new BatchRowsQueryResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBatchRowsQueryResponse() *BatchRowsQueryResponse {
	this := BatchRowsQueryResponse{}
	return &this
}

// NewBatchRowsQueryResponseWithDefaults instantiates a new BatchRowsQueryResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBatchRowsQueryResponseWithDefaults() *BatchRowsQueryResponse {
	this := BatchRowsQueryResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BatchRowsQueryResponse) GetData() BatchRowsQueryResponseData {
	if o == nil || o.Data == nil {
		var ret BatchRowsQueryResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRowsQueryResponse) GetDataOk() (*BatchRowsQueryResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BatchRowsQueryResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given BatchRowsQueryResponseData and assigns it to the Data field.
func (o *BatchRowsQueryResponse) SetData(v BatchRowsQueryResponseData) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BatchRowsQueryResponse) GetIncluded() []TableRowResourceData {
	if o == nil || o.Included == nil {
		var ret []TableRowResourceData
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRowsQueryResponse) GetIncludedOk() (*[]TableRowResourceData, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BatchRowsQueryResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []TableRowResourceData and assigns it to the Included field.
func (o *BatchRowsQueryResponse) SetIncluded(v []TableRowResourceData) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BatchRowsQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BatchRowsQueryResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *BatchRowsQueryResponseData `json:"data,omitempty"`
		Included []TableRowResourceData      `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
