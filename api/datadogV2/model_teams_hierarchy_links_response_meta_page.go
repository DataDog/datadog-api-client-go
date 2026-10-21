// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsHierarchyLinksResponseMetaPage Metadata related to paging information that is included in the response when querying the team hierarchy links
type TeamsHierarchyLinksResponseMetaPage struct {
	// First page number.
	FirstNumber *int64 `json:"first_number,omitempty"`
	// First offset.
	FirstOffset *int64 `json:"first_offset,omitempty"`
	// Last page number.
	LastNumber *int64 `json:"last_number,omitempty"`
	// Last offset.
	LastOffset datadog.NullableInt64 `json:"last_offset,omitempty"`
	// Page limit, when offset pagination is used.
	Limit *int64 `json:"limit,omitempty"`
	// Next page number.
	NextNumber datadog.NullableInt64 `json:"next_number,omitempty"`
	// Next offset.
	NextOffset datadog.NullableInt64 `json:"next_offset,omitempty"`
	// Page number.
	Number *int64 `json:"number,omitempty"`
	// Current offset, when offset pagination is used.
	Offset *int64 `json:"offset,omitempty"`
	// Previous page number.
	PrevNumber datadog.NullableInt64 `json:"prev_number,omitempty"`
	// Previous offset.
	PrevOffset datadog.NullableInt64 `json:"prev_offset,omitempty"`
	// Page size.
	Size *int64 `json:"size,omitempty"`
	// Total number of results.
	Total *int64 `json:"total,omitempty"`
	// Pagination type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsHierarchyLinksResponseMetaPage instantiates a new TeamsHierarchyLinksResponseMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsHierarchyLinksResponseMetaPage() *TeamsHierarchyLinksResponseMetaPage {
	this := TeamsHierarchyLinksResponseMetaPage{}
	return &this
}

// NewTeamsHierarchyLinksResponseMetaPageWithDefaults instantiates a new TeamsHierarchyLinksResponseMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsHierarchyLinksResponseMetaPageWithDefaults() *TeamsHierarchyLinksResponseMetaPage {
	this := TeamsHierarchyLinksResponseMetaPage{}
	return &this
}

// GetFirstNumber returns the FirstNumber field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetFirstNumber() int64 {
	if o == nil || o.FirstNumber == nil {
		var ret int64
		return ret
	}
	return *o.FirstNumber
}

// GetFirstNumberOk returns a tuple with the FirstNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetFirstNumberOk() (*int64, bool) {
	if o == nil || o.FirstNumber == nil {
		return nil, false
	}
	return o.FirstNumber, true
}

// HasFirstNumber returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasFirstNumber() bool {
	return o != nil && o.FirstNumber != nil
}

// SetFirstNumber gets a reference to the given int64 and assigns it to the FirstNumber field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetFirstNumber(v int64) {
	o.FirstNumber = &v
}

// GetFirstOffset returns the FirstOffset field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetFirstOffset() int64 {
	if o == nil || o.FirstOffset == nil {
		var ret int64
		return ret
	}
	return *o.FirstOffset
}

// GetFirstOffsetOk returns a tuple with the FirstOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetFirstOffsetOk() (*int64, bool) {
	if o == nil || o.FirstOffset == nil {
		return nil, false
	}
	return o.FirstOffset, true
}

// HasFirstOffset returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasFirstOffset() bool {
	return o != nil && o.FirstOffset != nil
}

// SetFirstOffset gets a reference to the given int64 and assigns it to the FirstOffset field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetFirstOffset(v int64) {
	o.FirstOffset = &v
}

// GetLastNumber returns the LastNumber field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetLastNumber() int64 {
	if o == nil || o.LastNumber == nil {
		var ret int64
		return ret
	}
	return *o.LastNumber
}

// GetLastNumberOk returns a tuple with the LastNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetLastNumberOk() (*int64, bool) {
	if o == nil || o.LastNumber == nil {
		return nil, false
	}
	return o.LastNumber, true
}

// HasLastNumber returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasLastNumber() bool {
	return o != nil && o.LastNumber != nil
}

// SetLastNumber gets a reference to the given int64 and assigns it to the LastNumber field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetLastNumber(v int64) {
	o.LastNumber = &v
}

// GetLastOffset returns the LastOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsHierarchyLinksResponseMetaPage) GetLastOffset() int64 {
	if o == nil || o.LastOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastOffset.Get()
}

// GetLastOffsetOk returns a tuple with the LastOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamsHierarchyLinksResponseMetaPage) GetLastOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastOffset.Get(), o.LastOffset.IsSet()
}

// HasLastOffset returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasLastOffset() bool {
	return o != nil && o.LastOffset.IsSet()
}

// SetLastOffset gets a reference to the given datadog.NullableInt64 and assigns it to the LastOffset field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetLastOffset(v int64) {
	o.LastOffset.Set(&v)
}

// SetLastOffsetNil sets the value for LastOffset to be an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) SetLastOffsetNil() {
	o.LastOffset.Set(nil)
}

// UnsetLastOffset ensures that no value is present for LastOffset, not even an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) UnsetLastOffset() {
	o.LastOffset.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetLimit(v int64) {
	o.Limit = &v
}

// GetNextNumber returns the NextNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsHierarchyLinksResponseMetaPage) GetNextNumber() int64 {
	if o == nil || o.NextNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextNumber.Get()
}

// GetNextNumberOk returns a tuple with the NextNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamsHierarchyLinksResponseMetaPage) GetNextNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextNumber.Get(), o.NextNumber.IsSet()
}

// HasNextNumber returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasNextNumber() bool {
	return o != nil && o.NextNumber.IsSet()
}

// SetNextNumber gets a reference to the given datadog.NullableInt64 and assigns it to the NextNumber field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetNextNumber(v int64) {
	o.NextNumber.Set(&v)
}

// SetNextNumberNil sets the value for NextNumber to be an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) SetNextNumberNil() {
	o.NextNumber.Set(nil)
}

// UnsetNextNumber ensures that no value is present for NextNumber, not even an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) UnsetNextNumber() {
	o.NextNumber.Unset()
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsHierarchyLinksResponseMetaPage) GetNextOffset() int64 {
	if o == nil || o.NextOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset.Get()
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamsHierarchyLinksResponseMetaPage) GetNextOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextOffset.Get(), o.NextOffset.IsSet()
}

// HasNextOffset returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasNextOffset() bool {
	return o != nil && o.NextOffset.IsSet()
}

// SetNextOffset gets a reference to the given datadog.NullableInt64 and assigns it to the NextOffset field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetNextOffset(v int64) {
	o.NextOffset.Set(&v)
}

// SetNextOffsetNil sets the value for NextOffset to be an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) SetNextOffsetNil() {
	o.NextOffset.Set(nil)
}

// UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) UnsetNextOffset() {
	o.NextOffset.Unset()
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetNumber() int64 {
	if o == nil || o.Number == nil {
		var ret int64
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetNumberOk() (*int64, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasNumber() bool {
	return o != nil && o.Number != nil
}

// SetNumber gets a reference to the given int64 and assigns it to the Number field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetNumber(v int64) {
	o.Number = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetOffset(v int64) {
	o.Offset = &v
}

// GetPrevNumber returns the PrevNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsHierarchyLinksResponseMetaPage) GetPrevNumber() int64 {
	if o == nil || o.PrevNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrevNumber.Get()
}

// GetPrevNumberOk returns a tuple with the PrevNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamsHierarchyLinksResponseMetaPage) GetPrevNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevNumber.Get(), o.PrevNumber.IsSet()
}

// HasPrevNumber returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasPrevNumber() bool {
	return o != nil && o.PrevNumber.IsSet()
}

// SetPrevNumber gets a reference to the given datadog.NullableInt64 and assigns it to the PrevNumber field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetPrevNumber(v int64) {
	o.PrevNumber.Set(&v)
}

// SetPrevNumberNil sets the value for PrevNumber to be an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) SetPrevNumberNil() {
	o.PrevNumber.Set(nil)
}

// UnsetPrevNumber ensures that no value is present for PrevNumber, not even an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) UnsetPrevNumber() {
	o.PrevNumber.Unset()
}

// GetPrevOffset returns the PrevOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamsHierarchyLinksResponseMetaPage) GetPrevOffset() int64 {
	if o == nil || o.PrevOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrevOffset.Get()
}

// GetPrevOffsetOk returns a tuple with the PrevOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamsHierarchyLinksResponseMetaPage) GetPrevOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevOffset.Get(), o.PrevOffset.IsSet()
}

// HasPrevOffset returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasPrevOffset() bool {
	return o != nil && o.PrevOffset.IsSet()
}

// SetPrevOffset gets a reference to the given datadog.NullableInt64 and assigns it to the PrevOffset field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetPrevOffset(v int64) {
	o.PrevOffset.Set(&v)
}

// SetPrevOffsetNil sets the value for PrevOffset to be an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) SetPrevOffsetNil() {
	o.PrevOffset.Set(nil)
}

// UnsetPrevOffset ensures that no value is present for PrevOffset, not even an explicit nil.
func (o *TeamsHierarchyLinksResponseMetaPage) UnsetPrevOffset() {
	o.PrevOffset.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetSize(v int64) {
	o.Size = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetTotal(v int64) {
	o.Total = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TeamsHierarchyLinksResponseMetaPage) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TeamsHierarchyLinksResponseMetaPage) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TeamsHierarchyLinksResponseMetaPage) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsHierarchyLinksResponseMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FirstNumber != nil {
		toSerialize["first_number"] = o.FirstNumber
	}
	if o.FirstOffset != nil {
		toSerialize["first_offset"] = o.FirstOffset
	}
	if o.LastNumber != nil {
		toSerialize["last_number"] = o.LastNumber
	}
	if o.LastOffset.IsSet() {
		toSerialize["last_offset"] = o.LastOffset.Get()
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextNumber.IsSet() {
		toSerialize["next_number"] = o.NextNumber.Get()
	}
	if o.NextOffset.IsSet() {
		toSerialize["next_offset"] = o.NextOffset.Get()
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.PrevNumber.IsSet() {
		toSerialize["prev_number"] = o.PrevNumber.Get()
	}
	if o.PrevOffset.IsSet() {
		toSerialize["prev_offset"] = o.PrevOffset.Get()
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
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
func (o *TeamsHierarchyLinksResponseMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstNumber *int64                `json:"first_number,omitempty"`
		FirstOffset *int64                `json:"first_offset,omitempty"`
		LastNumber  *int64                `json:"last_number,omitempty"`
		LastOffset  datadog.NullableInt64 `json:"last_offset,omitempty"`
		Limit       *int64                `json:"limit,omitempty"`
		NextNumber  datadog.NullableInt64 `json:"next_number,omitempty"`
		NextOffset  datadog.NullableInt64 `json:"next_offset,omitempty"`
		Number      *int64                `json:"number,omitempty"`
		Offset      *int64                `json:"offset,omitempty"`
		PrevNumber  datadog.NullableInt64 `json:"prev_number,omitempty"`
		PrevOffset  datadog.NullableInt64 `json:"prev_offset,omitempty"`
		Size        *int64                `json:"size,omitempty"`
		Total       *int64                `json:"total,omitempty"`
		Type        *string               `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first_number", "first_offset", "last_number", "last_offset", "limit", "next_number", "next_offset", "number", "offset", "prev_number", "prev_offset", "size", "total", "type"})
	} else {
		return err
	}
	o.FirstNumber = all.FirstNumber
	o.FirstOffset = all.FirstOffset
	o.LastNumber = all.LastNumber
	o.LastOffset = all.LastOffset
	o.Limit = all.Limit
	o.NextNumber = all.NextNumber
	o.NextOffset = all.NextOffset
	o.Number = all.Number
	o.Offset = all.Offset
	o.PrevNumber = all.PrevNumber
	o.PrevOffset = all.PrevOffset
	o.Size = all.Size
	o.Total = all.Total
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
