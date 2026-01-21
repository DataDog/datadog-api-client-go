// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesPagination Offset-based pagination schema.
type StatusPagesPagination struct {
	// Integer representing the offset to fetch the first page of results.
	FirstOffset *int64 `json:"first_offset,omitempty"`
	// Integer representing the offset to fetch the last page of results.
	LastOffset datadog.NullableInt64 `json:"last_offset,omitempty"`
	// Integer representing the number of elements to returned in the results.
	Limit *int64 `json:"limit,omitempty"`
	// Integer representing the index of the first element in the next page of results. Equal to page size added to the current offset.
	NextOffset datadog.NullableInt64 `json:"next_offset,omitempty"`
	// Integer representing the index of the first element in the results.
	Offset *int64 `json:"offset,omitempty"`
	// Integer representing the index of the first element in the previous page of results.
	PrevOffset datadog.NullableInt64 `json:"prev_offset,omitempty"`
	// Integer representing the total number of elements available.
	Total datadog.NullableInt64 `json:"total,omitempty"`
	//
	Type *StatusPagesPaginationType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPagesPagination instantiates a new StatusPagesPagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPagesPagination() *StatusPagesPagination {
	this := StatusPagesPagination{}
	var typeVar StatusPagesPaginationType = STATUSPAGESPAGINATIONTYPE_OFFSET_LIMIT
	this.Type = &typeVar
	return &this
}

// NewStatusPagesPaginationWithDefaults instantiates a new StatusPagesPagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPagesPaginationWithDefaults() *StatusPagesPagination {
	this := StatusPagesPagination{}
	var typeVar StatusPagesPaginationType = STATUSPAGESPAGINATIONTYPE_OFFSET_LIMIT
	this.Type = &typeVar
	return &this
}

// GetFirstOffset returns the FirstOffset field value if set, zero value otherwise.
func (o *StatusPagesPagination) GetFirstOffset() int64 {
	if o == nil || o.FirstOffset == nil {
		var ret int64
		return ret
	}
	return *o.FirstOffset
}

// GetFirstOffsetOk returns a tuple with the FirstOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesPagination) GetFirstOffsetOk() (*int64, bool) {
	if o == nil || o.FirstOffset == nil {
		return nil, false
	}
	return o.FirstOffset, true
}

// HasFirstOffset returns a boolean if a field has been set.
func (o *StatusPagesPagination) HasFirstOffset() bool {
	return o != nil && o.FirstOffset != nil
}

// SetFirstOffset gets a reference to the given int64 and assigns it to the FirstOffset field.
func (o *StatusPagesPagination) SetFirstOffset(v int64) {
	o.FirstOffset = &v
}

// GetLastOffset returns the LastOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusPagesPagination) GetLastOffset() int64 {
	if o == nil || o.LastOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastOffset.Get()
}

// GetLastOffsetOk returns a tuple with the LastOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StatusPagesPagination) GetLastOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastOffset.Get(), o.LastOffset.IsSet()
}

// HasLastOffset returns a boolean if a field has been set.
func (o *StatusPagesPagination) HasLastOffset() bool {
	return o != nil && o.LastOffset.IsSet()
}

// SetLastOffset gets a reference to the given datadog.NullableInt64 and assigns it to the LastOffset field.
func (o *StatusPagesPagination) SetLastOffset(v int64) {
	o.LastOffset.Set(&v)
}

// SetLastOffsetNil sets the value for LastOffset to be an explicit nil.
func (o *StatusPagesPagination) SetLastOffsetNil() {
	o.LastOffset.Set(nil)
}

// UnsetLastOffset ensures that no value is present for LastOffset, not even an explicit nil.
func (o *StatusPagesPagination) UnsetLastOffset() {
	o.LastOffset.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *StatusPagesPagination) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesPagination) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *StatusPagesPagination) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *StatusPagesPagination) SetLimit(v int64) {
	o.Limit = &v
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusPagesPagination) GetNextOffset() int64 {
	if o == nil || o.NextOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset.Get()
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StatusPagesPagination) GetNextOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextOffset.Get(), o.NextOffset.IsSet()
}

// HasNextOffset returns a boolean if a field has been set.
func (o *StatusPagesPagination) HasNextOffset() bool {
	return o != nil && o.NextOffset.IsSet()
}

// SetNextOffset gets a reference to the given datadog.NullableInt64 and assigns it to the NextOffset field.
func (o *StatusPagesPagination) SetNextOffset(v int64) {
	o.NextOffset.Set(&v)
}

// SetNextOffsetNil sets the value for NextOffset to be an explicit nil.
func (o *StatusPagesPagination) SetNextOffsetNil() {
	o.NextOffset.Set(nil)
}

// UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil.
func (o *StatusPagesPagination) UnsetNextOffset() {
	o.NextOffset.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *StatusPagesPagination) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesPagination) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *StatusPagesPagination) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *StatusPagesPagination) SetOffset(v int64) {
	o.Offset = &v
}

// GetPrevOffset returns the PrevOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusPagesPagination) GetPrevOffset() int64 {
	if o == nil || o.PrevOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrevOffset.Get()
}

// GetPrevOffsetOk returns a tuple with the PrevOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StatusPagesPagination) GetPrevOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevOffset.Get(), o.PrevOffset.IsSet()
}

// HasPrevOffset returns a boolean if a field has been set.
func (o *StatusPagesPagination) HasPrevOffset() bool {
	return o != nil && o.PrevOffset.IsSet()
}

// SetPrevOffset gets a reference to the given datadog.NullableInt64 and assigns it to the PrevOffset field.
func (o *StatusPagesPagination) SetPrevOffset(v int64) {
	o.PrevOffset.Set(&v)
}

// SetPrevOffsetNil sets the value for PrevOffset to be an explicit nil.
func (o *StatusPagesPagination) SetPrevOffsetNil() {
	o.PrevOffset.Set(nil)
}

// UnsetPrevOffset ensures that no value is present for PrevOffset, not even an explicit nil.
func (o *StatusPagesPagination) UnsetPrevOffset() {
	o.PrevOffset.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusPagesPagination) GetTotal() int64 {
	if o == nil || o.Total.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StatusPagesPagination) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *StatusPagesPagination) HasTotal() bool {
	return o != nil && o.Total.IsSet()
}

// SetTotal gets a reference to the given datadog.NullableInt64 and assigns it to the Total field.
func (o *StatusPagesPagination) SetTotal(v int64) {
	o.Total.Set(&v)
}

// SetTotalNil sets the value for Total to be an explicit nil.
func (o *StatusPagesPagination) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil.
func (o *StatusPagesPagination) UnsetTotal() {
	o.Total.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StatusPagesPagination) GetType() StatusPagesPaginationType {
	if o == nil || o.Type == nil {
		var ret StatusPagesPaginationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesPagination) GetTypeOk() (*StatusPagesPaginationType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StatusPagesPagination) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given StatusPagesPaginationType and assigns it to the Type field.
func (o *StatusPagesPagination) SetType(v StatusPagesPaginationType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPagesPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FirstOffset != nil {
		toSerialize["first_offset"] = o.FirstOffset
	}
	if o.LastOffset.IsSet() {
		toSerialize["last_offset"] = o.LastOffset.Get()
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextOffset.IsSet() {
		toSerialize["next_offset"] = o.NextOffset.Get()
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.PrevOffset.IsSet() {
		toSerialize["prev_offset"] = o.PrevOffset.Get()
	}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
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
func (o *StatusPagesPagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstOffset *int64                     `json:"first_offset,omitempty"`
		LastOffset  datadog.NullableInt64      `json:"last_offset,omitempty"`
		Limit       *int64                     `json:"limit,omitempty"`
		NextOffset  datadog.NullableInt64      `json:"next_offset,omitempty"`
		Offset      *int64                     `json:"offset,omitempty"`
		PrevOffset  datadog.NullableInt64      `json:"prev_offset,omitempty"`
		Total       datadog.NullableInt64      `json:"total,omitempty"`
		Type        *StatusPagesPaginationType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first_offset", "last_offset", "limit", "next_offset", "offset", "prev_offset", "total", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FirstOffset = all.FirstOffset
	o.LastOffset = all.LastOffset
	o.Limit = all.Limit
	o.NextOffset = all.NextOffset
	o.Offset = all.Offset
	o.PrevOffset = all.PrevOffset
	o.Total = all.Total
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
