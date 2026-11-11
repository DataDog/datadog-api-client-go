// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagHistoryPaginationMetaPage Pagination metadata for feature flag change history responses.
type FeatureFlagHistoryPaginationMetaPage struct {
	// Offset of the first page.
	FirstOffset *int64 `json:"first_offset,omitempty"`
	// Offset of the last page.
	LastOffset datadog.NullableInt64 `json:"last_offset,omitempty"`
	// Maximum number of events returned on this page.
	Limit *int64 `json:"limit,omitempty"`
	// Offset of the next page, or null when there is no next page.
	NextOffset datadog.NullableInt64 `json:"next_offset,omitempty"`
	// Number of matching events skipped before this page.
	Offset *int64 `json:"offset,omitempty"`
	// Offset of the previous page, or null when there is no previous page.
	PrevOffset datadog.NullableInt64 `json:"prev_offset,omitempty"`
	// Total number of events matching the filters, independent of the page size.
	// It counts experiment proposals that you cannot read, which are left out of the pages, so a page can hold fewer events than `limit`.
	Total *int64 `json:"total,omitempty"`
	// The pagination method. Always `offset_limit`.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagHistoryPaginationMetaPage instantiates a new FeatureFlagHistoryPaginationMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagHistoryPaginationMetaPage() *FeatureFlagHistoryPaginationMetaPage {
	this := FeatureFlagHistoryPaginationMetaPage{}
	return &this
}

// NewFeatureFlagHistoryPaginationMetaPageWithDefaults instantiates a new FeatureFlagHistoryPaginationMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagHistoryPaginationMetaPageWithDefaults() *FeatureFlagHistoryPaginationMetaPage {
	this := FeatureFlagHistoryPaginationMetaPage{}
	return &this
}

// GetFirstOffset returns the FirstOffset field value if set, zero value otherwise.
func (o *FeatureFlagHistoryPaginationMetaPage) GetFirstOffset() int64 {
	if o == nil || o.FirstOffset == nil {
		var ret int64
		return ret
	}
	return *o.FirstOffset
}

// GetFirstOffsetOk returns a tuple with the FirstOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) GetFirstOffsetOk() (*int64, bool) {
	if o == nil || o.FirstOffset == nil {
		return nil, false
	}
	return o.FirstOffset, true
}

// HasFirstOffset returns a boolean if a field has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) HasFirstOffset() bool {
	return o != nil && o.FirstOffset != nil
}

// SetFirstOffset gets a reference to the given int64 and assigns it to the FirstOffset field.
func (o *FeatureFlagHistoryPaginationMetaPage) SetFirstOffset(v int64) {
	o.FirstOffset = &v
}

// GetLastOffset returns the LastOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagHistoryPaginationMetaPage) GetLastOffset() int64 {
	if o == nil || o.LastOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastOffset.Get()
}

// GetLastOffsetOk returns a tuple with the LastOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagHistoryPaginationMetaPage) GetLastOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastOffset.Get(), o.LastOffset.IsSet()
}

// HasLastOffset returns a boolean if a field has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) HasLastOffset() bool {
	return o != nil && o.LastOffset.IsSet()
}

// SetLastOffset gets a reference to the given datadog.NullableInt64 and assigns it to the LastOffset field.
func (o *FeatureFlagHistoryPaginationMetaPage) SetLastOffset(v int64) {
	o.LastOffset.Set(&v)
}

// SetLastOffsetNil sets the value for LastOffset to be an explicit nil.
func (o *FeatureFlagHistoryPaginationMetaPage) SetLastOffsetNil() {
	o.LastOffset.Set(nil)
}

// UnsetLastOffset ensures that no value is present for LastOffset, not even an explicit nil.
func (o *FeatureFlagHistoryPaginationMetaPage) UnsetLastOffset() {
	o.LastOffset.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *FeatureFlagHistoryPaginationMetaPage) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *FeatureFlagHistoryPaginationMetaPage) SetLimit(v int64) {
	o.Limit = &v
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagHistoryPaginationMetaPage) GetNextOffset() int64 {
	if o == nil || o.NextOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset.Get()
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagHistoryPaginationMetaPage) GetNextOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextOffset.Get(), o.NextOffset.IsSet()
}

// HasNextOffset returns a boolean if a field has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) HasNextOffset() bool {
	return o != nil && o.NextOffset.IsSet()
}

// SetNextOffset gets a reference to the given datadog.NullableInt64 and assigns it to the NextOffset field.
func (o *FeatureFlagHistoryPaginationMetaPage) SetNextOffset(v int64) {
	o.NextOffset.Set(&v)
}

// SetNextOffsetNil sets the value for NextOffset to be an explicit nil.
func (o *FeatureFlagHistoryPaginationMetaPage) SetNextOffsetNil() {
	o.NextOffset.Set(nil)
}

// UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil.
func (o *FeatureFlagHistoryPaginationMetaPage) UnsetNextOffset() {
	o.NextOffset.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *FeatureFlagHistoryPaginationMetaPage) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *FeatureFlagHistoryPaginationMetaPage) SetOffset(v int64) {
	o.Offset = &v
}

// GetPrevOffset returns the PrevOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagHistoryPaginationMetaPage) GetPrevOffset() int64 {
	if o == nil || o.PrevOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrevOffset.Get()
}

// GetPrevOffsetOk returns a tuple with the PrevOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagHistoryPaginationMetaPage) GetPrevOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevOffset.Get(), o.PrevOffset.IsSet()
}

// HasPrevOffset returns a boolean if a field has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) HasPrevOffset() bool {
	return o != nil && o.PrevOffset.IsSet()
}

// SetPrevOffset gets a reference to the given datadog.NullableInt64 and assigns it to the PrevOffset field.
func (o *FeatureFlagHistoryPaginationMetaPage) SetPrevOffset(v int64) {
	o.PrevOffset.Set(&v)
}

// SetPrevOffsetNil sets the value for PrevOffset to be an explicit nil.
func (o *FeatureFlagHistoryPaginationMetaPage) SetPrevOffsetNil() {
	o.PrevOffset.Set(nil)
}

// UnsetPrevOffset ensures that no value is present for PrevOffset, not even an explicit nil.
func (o *FeatureFlagHistoryPaginationMetaPage) UnsetPrevOffset() {
	o.PrevOffset.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *FeatureFlagHistoryPaginationMetaPage) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *FeatureFlagHistoryPaginationMetaPage) SetTotal(v int64) {
	o.Total = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FeatureFlagHistoryPaginationMetaPage) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FeatureFlagHistoryPaginationMetaPage) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FeatureFlagHistoryPaginationMetaPage) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagHistoryPaginationMetaPage) MarshalJSON() ([]byte, error) {
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
	if o.Total != nil {
		toSerialize["total"] = o.Total
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
func (o *FeatureFlagHistoryPaginationMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstOffset *int64                `json:"first_offset,omitempty"`
		LastOffset  datadog.NullableInt64 `json:"last_offset,omitempty"`
		Limit       *int64                `json:"limit,omitempty"`
		NextOffset  datadog.NullableInt64 `json:"next_offset,omitempty"`
		Offset      *int64                `json:"offset,omitempty"`
		PrevOffset  datadog.NullableInt64 `json:"prev_offset,omitempty"`
		Total       *int64                `json:"total,omitempty"`
		Type        *string               `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first_offset", "last_offset", "limit", "next_offset", "offset", "prev_offset", "total", "type"})
	} else {
		return err
	}
	o.FirstOffset = all.FirstOffset
	o.LastOffset = all.LastOffset
	o.Limit = all.Limit
	o.NextOffset = all.NextOffset
	o.Offset = all.Offset
	o.PrevOffset = all.PrevOffset
	o.Total = all.Total
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
