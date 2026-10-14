// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagsPaginationMetaPage Pagination metadata for feature flags list responses.
type FeatureFlagsPaginationMetaPage struct {
	// Offset of the first page.
	FirstOffset *int64 `json:"first_offset,omitempty"`
	// Offset of the last page.
	LastOffset datadog.NullableInt64 `json:"last_offset,omitempty"`
	// Maximum number of flags returned on this page.
	Limit *int64 `json:"limit,omitempty"`
	// Offset of the next page, or null when there is no next page.
	NextOffset datadog.NullableInt64 `json:"next_offset,omitempty"`
	// Number of matching flags skipped before this page.
	Offset *int64 `json:"offset,omitempty"`
	// Offset of the previous page, or null when there is no previous page.
	PrevOffset datadog.NullableInt64 `json:"prev_offset,omitempty"`
	// Total number of flags matching the filters, independent of the page size.
	Total *int64 `json:"total,omitempty"`
	// Legacy count field. Use `total` for the number of matching flags.
	// Deprecated
	TotalCount *int64 `json:"total_count,omitempty"`
	// Legacy filtered count field. Use `total` for the number of matching flags.
	// Deprecated
	TotalFilteredCount *int64 `json:"total_filtered_count,omitempty"`
	// The pagination method. Always `offset_limit`.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagsPaginationMetaPage instantiates a new FeatureFlagsPaginationMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagsPaginationMetaPage() *FeatureFlagsPaginationMetaPage {
	this := FeatureFlagsPaginationMetaPage{}
	return &this
}

// NewFeatureFlagsPaginationMetaPageWithDefaults instantiates a new FeatureFlagsPaginationMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagsPaginationMetaPageWithDefaults() *FeatureFlagsPaginationMetaPage {
	this := FeatureFlagsPaginationMetaPage{}
	return &this
}

// GetFirstOffset returns the FirstOffset field value if set, zero value otherwise.
func (o *FeatureFlagsPaginationMetaPage) GetFirstOffset() int64 {
	if o == nil || o.FirstOffset == nil {
		var ret int64
		return ret
	}
	return *o.FirstOffset
}

// GetFirstOffsetOk returns a tuple with the FirstOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagsPaginationMetaPage) GetFirstOffsetOk() (*int64, bool) {
	if o == nil || o.FirstOffset == nil {
		return nil, false
	}
	return o.FirstOffset, true
}

// HasFirstOffset returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationMetaPage) HasFirstOffset() bool {
	return o != nil && o.FirstOffset != nil
}

// SetFirstOffset gets a reference to the given int64 and assigns it to the FirstOffset field.
func (o *FeatureFlagsPaginationMetaPage) SetFirstOffset(v int64) {
	o.FirstOffset = &v
}

// GetLastOffset returns the LastOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagsPaginationMetaPage) GetLastOffset() int64 {
	if o == nil || o.LastOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastOffset.Get()
}

// GetLastOffsetOk returns a tuple with the LastOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagsPaginationMetaPage) GetLastOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastOffset.Get(), o.LastOffset.IsSet()
}

// HasLastOffset returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationMetaPage) HasLastOffset() bool {
	return o != nil && o.LastOffset.IsSet()
}

// SetLastOffset gets a reference to the given datadog.NullableInt64 and assigns it to the LastOffset field.
func (o *FeatureFlagsPaginationMetaPage) SetLastOffset(v int64) {
	o.LastOffset.Set(&v)
}

// SetLastOffsetNil sets the value for LastOffset to be an explicit nil.
func (o *FeatureFlagsPaginationMetaPage) SetLastOffsetNil() {
	o.LastOffset.Set(nil)
}

// UnsetLastOffset ensures that no value is present for LastOffset, not even an explicit nil.
func (o *FeatureFlagsPaginationMetaPage) UnsetLastOffset() {
	o.LastOffset.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *FeatureFlagsPaginationMetaPage) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagsPaginationMetaPage) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationMetaPage) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *FeatureFlagsPaginationMetaPage) SetLimit(v int64) {
	o.Limit = &v
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagsPaginationMetaPage) GetNextOffset() int64 {
	if o == nil || o.NextOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset.Get()
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagsPaginationMetaPage) GetNextOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextOffset.Get(), o.NextOffset.IsSet()
}

// HasNextOffset returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationMetaPage) HasNextOffset() bool {
	return o != nil && o.NextOffset.IsSet()
}

// SetNextOffset gets a reference to the given datadog.NullableInt64 and assigns it to the NextOffset field.
func (o *FeatureFlagsPaginationMetaPage) SetNextOffset(v int64) {
	o.NextOffset.Set(&v)
}

// SetNextOffsetNil sets the value for NextOffset to be an explicit nil.
func (o *FeatureFlagsPaginationMetaPage) SetNextOffsetNil() {
	o.NextOffset.Set(nil)
}

// UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil.
func (o *FeatureFlagsPaginationMetaPage) UnsetNextOffset() {
	o.NextOffset.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *FeatureFlagsPaginationMetaPage) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagsPaginationMetaPage) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationMetaPage) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *FeatureFlagsPaginationMetaPage) SetOffset(v int64) {
	o.Offset = &v
}

// GetPrevOffset returns the PrevOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureFlagsPaginationMetaPage) GetPrevOffset() int64 {
	if o == nil || o.PrevOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrevOffset.Get()
}

// GetPrevOffsetOk returns a tuple with the PrevOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FeatureFlagsPaginationMetaPage) GetPrevOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevOffset.Get(), o.PrevOffset.IsSet()
}

// HasPrevOffset returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationMetaPage) HasPrevOffset() bool {
	return o != nil && o.PrevOffset.IsSet()
}

// SetPrevOffset gets a reference to the given datadog.NullableInt64 and assigns it to the PrevOffset field.
func (o *FeatureFlagsPaginationMetaPage) SetPrevOffset(v int64) {
	o.PrevOffset.Set(&v)
}

// SetPrevOffsetNil sets the value for PrevOffset to be an explicit nil.
func (o *FeatureFlagsPaginationMetaPage) SetPrevOffsetNil() {
	o.PrevOffset.Set(nil)
}

// UnsetPrevOffset ensures that no value is present for PrevOffset, not even an explicit nil.
func (o *FeatureFlagsPaginationMetaPage) UnsetPrevOffset() {
	o.PrevOffset.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *FeatureFlagsPaginationMetaPage) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagsPaginationMetaPage) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationMetaPage) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *FeatureFlagsPaginationMetaPage) SetTotal(v int64) {
	o.Total = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
// Deprecated
func (o *FeatureFlagsPaginationMetaPage) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *FeatureFlagsPaginationMetaPage) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationMetaPage) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
// Deprecated
func (o *FeatureFlagsPaginationMetaPage) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetTotalFilteredCount returns the TotalFilteredCount field value if set, zero value otherwise.
// Deprecated
func (o *FeatureFlagsPaginationMetaPage) GetTotalFilteredCount() int64 {
	if o == nil || o.TotalFilteredCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalFilteredCount
}

// GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *FeatureFlagsPaginationMetaPage) GetTotalFilteredCountOk() (*int64, bool) {
	if o == nil || o.TotalFilteredCount == nil {
		return nil, false
	}
	return o.TotalFilteredCount, true
}

// HasTotalFilteredCount returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationMetaPage) HasTotalFilteredCount() bool {
	return o != nil && o.TotalFilteredCount != nil
}

// SetTotalFilteredCount gets a reference to the given int64 and assigns it to the TotalFilteredCount field.
// Deprecated
func (o *FeatureFlagsPaginationMetaPage) SetTotalFilteredCount(v int64) {
	o.TotalFilteredCount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FeatureFlagsPaginationMetaPage) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagsPaginationMetaPage) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationMetaPage) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FeatureFlagsPaginationMetaPage) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagsPaginationMetaPage) MarshalJSON() ([]byte, error) {
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
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	if o.TotalFilteredCount != nil {
		toSerialize["total_filtered_count"] = o.TotalFilteredCount
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
func (o *FeatureFlagsPaginationMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstOffset        *int64                `json:"first_offset,omitempty"`
		LastOffset         datadog.NullableInt64 `json:"last_offset,omitempty"`
		Limit              *int64                `json:"limit,omitempty"`
		NextOffset         datadog.NullableInt64 `json:"next_offset,omitempty"`
		Offset             *int64                `json:"offset,omitempty"`
		PrevOffset         datadog.NullableInt64 `json:"prev_offset,omitempty"`
		Total              *int64                `json:"total,omitempty"`
		TotalCount         *int64                `json:"total_count,omitempty"`
		TotalFilteredCount *int64                `json:"total_filtered_count,omitempty"`
		Type               *string               `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first_offset", "last_offset", "limit", "next_offset", "offset", "prev_offset", "total", "total_count", "total_filtered_count", "type"})
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
	o.TotalCount = all.TotalCount
	o.TotalFilteredCount = all.TotalFilteredCount
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
