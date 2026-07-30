// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationStrongLinksListResponseMeta Metadata for a list of RUM operation strong links.
type RUMOperationStrongLinksListResponseMeta struct {
	// The pagination limit.
	Limit *int64 `json:"limit,omitempty"`
	// The current offset.
	Offset *int64 `json:"offset,omitempty"`
	// The total number of strong links matching the request.
	Total *int64 `json:"total,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMOperationStrongLinksListResponseMeta instantiates a new RUMOperationStrongLinksListResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMOperationStrongLinksListResponseMeta() *RUMOperationStrongLinksListResponseMeta {
	this := RUMOperationStrongLinksListResponseMeta{}
	return &this
}

// NewRUMOperationStrongLinksListResponseMetaWithDefaults instantiates a new RUMOperationStrongLinksListResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMOperationStrongLinksListResponseMetaWithDefaults() *RUMOperationStrongLinksListResponseMeta {
	this := RUMOperationStrongLinksListResponseMeta{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *RUMOperationStrongLinksListResponseMeta) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinksListResponseMeta) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *RUMOperationStrongLinksListResponseMeta) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *RUMOperationStrongLinksListResponseMeta) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *RUMOperationStrongLinksListResponseMeta) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinksListResponseMeta) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *RUMOperationStrongLinksListResponseMeta) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *RUMOperationStrongLinksListResponseMeta) SetOffset(v int64) {
	o.Offset = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *RUMOperationStrongLinksListResponseMeta) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinksListResponseMeta) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *RUMOperationStrongLinksListResponseMeta) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *RUMOperationStrongLinksListResponseMeta) SetTotal(v int64) {
	o.Total = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMOperationStrongLinksListResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMOperationStrongLinksListResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Limit  *int64 `json:"limit,omitempty"`
		Offset *int64 `json:"offset,omitempty"`
		Total  *int64 `json:"total,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"limit", "offset", "total"})
	} else {
		return err
	}
	o.Limit = all.Limit
	o.Offset = all.Offset
	o.Total = all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
