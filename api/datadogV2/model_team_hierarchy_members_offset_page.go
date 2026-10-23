// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyMembersOffsetPage Pagination metadata using offset and limit.
type TeamHierarchyMembersOffsetPage struct {
	// First offset.
	FirstOffset int64 `json:"first_offset"`
	// Last offset. Null when unavailable.
	LastOffset datadog.NullableInt64 `json:"last_offset"`
	// Maximum number of members per page.
	Limit int64 `json:"limit"`
	// Next offset. Null when unavailable.
	NextOffset datadog.NullableInt64 `json:"next_offset"`
	// Current offset.
	Offset int64 `json:"offset"`
	// Previous offset. Null when unavailable.
	PrevOffset datadog.NullableInt64 `json:"prev_offset"`
	// Total number of members matching the filters.
	Total int64 `json:"total"`
	// Pagination mode.
	Type TeamHierarchyMembersOffsetPageType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyMembersOffsetPage instantiates a new TeamHierarchyMembersOffsetPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyMembersOffsetPage(firstOffset int64, lastOffset datadog.NullableInt64, limit int64, nextOffset datadog.NullableInt64, offset int64, prevOffset datadog.NullableInt64, total int64, typeVar TeamHierarchyMembersOffsetPageType) *TeamHierarchyMembersOffsetPage {
	this := TeamHierarchyMembersOffsetPage{}
	this.FirstOffset = firstOffset
	this.LastOffset = lastOffset
	this.Limit = limit
	this.NextOffset = nextOffset
	this.Offset = offset
	this.PrevOffset = prevOffset
	this.Total = total
	this.Type = typeVar
	return &this
}

// NewTeamHierarchyMembersOffsetPageWithDefaults instantiates a new TeamHierarchyMembersOffsetPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyMembersOffsetPageWithDefaults() *TeamHierarchyMembersOffsetPage {
	this := TeamHierarchyMembersOffsetPage{}
	var typeVar TeamHierarchyMembersOffsetPageType = TEAMHIERARCHYMEMBERSOFFSETPAGETYPE_OFFSET_LIMIT
	this.Type = typeVar
	return &this
}

// GetFirstOffset returns the FirstOffset field value.
func (o *TeamHierarchyMembersOffsetPage) GetFirstOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FirstOffset
}

// GetFirstOffsetOk returns a tuple with the FirstOffset field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersOffsetPage) GetFirstOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstOffset, true
}

// SetFirstOffset sets field value.
func (o *TeamHierarchyMembersOffsetPage) SetFirstOffset(v int64) {
	o.FirstOffset = v
}

// GetLastOffset returns the LastOffset field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *TeamHierarchyMembersOffsetPage) GetLastOffset() int64 {
	if o == nil || o.LastOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastOffset.Get()
}

// GetLastOffsetOk returns a tuple with the LastOffset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamHierarchyMembersOffsetPage) GetLastOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastOffset.Get(), o.LastOffset.IsSet()
}

// SetLastOffset sets field value.
func (o *TeamHierarchyMembersOffsetPage) SetLastOffset(v int64) {
	o.LastOffset.Set(&v)
}

// GetLimit returns the Limit field value.
func (o *TeamHierarchyMembersOffsetPage) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersOffsetPage) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *TeamHierarchyMembersOffsetPage) SetLimit(v int64) {
	o.Limit = v
}

// GetNextOffset returns the NextOffset field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *TeamHierarchyMembersOffsetPage) GetNextOffset() int64 {
	if o == nil || o.NextOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset.Get()
}

// GetNextOffsetOk returns a tuple with the NextOffset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamHierarchyMembersOffsetPage) GetNextOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextOffset.Get(), o.NextOffset.IsSet()
}

// SetNextOffset sets field value.
func (o *TeamHierarchyMembersOffsetPage) SetNextOffset(v int64) {
	o.NextOffset.Set(&v)
}

// GetOffset returns the Offset field value.
func (o *TeamHierarchyMembersOffsetPage) GetOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersOffsetPage) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value.
func (o *TeamHierarchyMembersOffsetPage) SetOffset(v int64) {
	o.Offset = v
}

// GetPrevOffset returns the PrevOffset field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *TeamHierarchyMembersOffsetPage) GetPrevOffset() int64 {
	if o == nil || o.PrevOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrevOffset.Get()
}

// GetPrevOffsetOk returns a tuple with the PrevOffset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamHierarchyMembersOffsetPage) GetPrevOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevOffset.Get(), o.PrevOffset.IsSet()
}

// SetPrevOffset sets field value.
func (o *TeamHierarchyMembersOffsetPage) SetPrevOffset(v int64) {
	o.PrevOffset.Set(&v)
}

// GetTotal returns the Total field value.
func (o *TeamHierarchyMembersOffsetPage) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersOffsetPage) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *TeamHierarchyMembersOffsetPage) SetTotal(v int64) {
	o.Total = v
}

// GetType returns the Type field value.
func (o *TeamHierarchyMembersOffsetPage) GetType() TeamHierarchyMembersOffsetPageType {
	if o == nil {
		var ret TeamHierarchyMembersOffsetPageType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersOffsetPage) GetTypeOk() (*TeamHierarchyMembersOffsetPageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamHierarchyMembersOffsetPage) SetType(v TeamHierarchyMembersOffsetPageType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyMembersOffsetPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["first_offset"] = o.FirstOffset
	toSerialize["last_offset"] = o.LastOffset.Get()
	toSerialize["limit"] = o.Limit
	toSerialize["next_offset"] = o.NextOffset.Get()
	toSerialize["offset"] = o.Offset
	toSerialize["prev_offset"] = o.PrevOffset.Get()
	toSerialize["total"] = o.Total
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamHierarchyMembersOffsetPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstOffset *int64                              `json:"first_offset"`
		LastOffset  datadog.NullableInt64               `json:"last_offset"`
		Limit       *int64                              `json:"limit"`
		NextOffset  datadog.NullableInt64               `json:"next_offset"`
		Offset      *int64                              `json:"offset"`
		PrevOffset  datadog.NullableInt64               `json:"prev_offset"`
		Total       *int64                              `json:"total"`
		Type        *TeamHierarchyMembersOffsetPageType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FirstOffset == nil {
		return fmt.Errorf("required field first_offset missing")
	}
	if !all.LastOffset.IsSet() {
		return fmt.Errorf("required field last_offset missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if !all.NextOffset.IsSet() {
		return fmt.Errorf("required field next_offset missing")
	}
	if all.Offset == nil {
		return fmt.Errorf("required field offset missing")
	}
	if !all.PrevOffset.IsSet() {
		return fmt.Errorf("required field prev_offset missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first_offset", "last_offset", "limit", "next_offset", "offset", "prev_offset", "total", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FirstOffset = *all.FirstOffset
	o.LastOffset = all.LastOffset
	o.Limit = *all.Limit
	o.NextOffset = all.NextOffset
	o.Offset = *all.Offset
	o.PrevOffset = all.PrevOffset
	o.Total = *all.Total
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
