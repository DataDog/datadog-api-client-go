// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyMembersNumberPage Pagination metadata using number and size.
type TeamHierarchyMembersNumberPage struct {
	// First page number.
	FirstNumber int64 `json:"first_number"`
	// Last page number. Null when unavailable.
	LastNumber datadog.NullableInt64 `json:"last_number"`
	// Next page number. Null when unavailable.
	NextNumber datadog.NullableInt64 `json:"next_number"`
	// Current zero-based page number.
	Number int64 `json:"number"`
	// Previous page number. Null when unavailable.
	PrevNumber datadog.NullableInt64 `json:"prev_number"`
	// Maximum number of members per page.
	Size int64 `json:"size"`
	// Total number of members matching the filters.
	Total int64 `json:"total"`
	// Pagination mode.
	Type TeamHierarchyMembersNumberPageType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyMembersNumberPage instantiates a new TeamHierarchyMembersNumberPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyMembersNumberPage(firstNumber int64, lastNumber datadog.NullableInt64, nextNumber datadog.NullableInt64, number int64, prevNumber datadog.NullableInt64, size int64, total int64, typeVar TeamHierarchyMembersNumberPageType) *TeamHierarchyMembersNumberPage {
	this := TeamHierarchyMembersNumberPage{}
	this.FirstNumber = firstNumber
	this.LastNumber = lastNumber
	this.NextNumber = nextNumber
	this.Number = number
	this.PrevNumber = prevNumber
	this.Size = size
	this.Total = total
	this.Type = typeVar
	return &this
}

// NewTeamHierarchyMembersNumberPageWithDefaults instantiates a new TeamHierarchyMembersNumberPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyMembersNumberPageWithDefaults() *TeamHierarchyMembersNumberPage {
	this := TeamHierarchyMembersNumberPage{}
	var typeVar TeamHierarchyMembersNumberPageType = TEAMHIERARCHYMEMBERSNUMBERPAGETYPE_NUMBER_SIZE
	this.Type = typeVar
	return &this
}

// GetFirstNumber returns the FirstNumber field value.
func (o *TeamHierarchyMembersNumberPage) GetFirstNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FirstNumber
}

// GetFirstNumberOk returns a tuple with the FirstNumber field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersNumberPage) GetFirstNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstNumber, true
}

// SetFirstNumber sets field value.
func (o *TeamHierarchyMembersNumberPage) SetFirstNumber(v int64) {
	o.FirstNumber = v
}

// GetLastNumber returns the LastNumber field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *TeamHierarchyMembersNumberPage) GetLastNumber() int64 {
	if o == nil || o.LastNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastNumber.Get()
}

// GetLastNumberOk returns a tuple with the LastNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamHierarchyMembersNumberPage) GetLastNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastNumber.Get(), o.LastNumber.IsSet()
}

// SetLastNumber sets field value.
func (o *TeamHierarchyMembersNumberPage) SetLastNumber(v int64) {
	o.LastNumber.Set(&v)
}

// GetNextNumber returns the NextNumber field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *TeamHierarchyMembersNumberPage) GetNextNumber() int64 {
	if o == nil || o.NextNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextNumber.Get()
}

// GetNextNumberOk returns a tuple with the NextNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamHierarchyMembersNumberPage) GetNextNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextNumber.Get(), o.NextNumber.IsSet()
}

// SetNextNumber sets field value.
func (o *TeamHierarchyMembersNumberPage) SetNextNumber(v int64) {
	o.NextNumber.Set(&v)
}

// GetNumber returns the Number field value.
func (o *TeamHierarchyMembersNumberPage) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersNumberPage) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value.
func (o *TeamHierarchyMembersNumberPage) SetNumber(v int64) {
	o.Number = v
}

// GetPrevNumber returns the PrevNumber field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *TeamHierarchyMembersNumberPage) GetPrevNumber() int64 {
	if o == nil || o.PrevNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PrevNumber.Get()
}

// GetPrevNumberOk returns a tuple with the PrevNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamHierarchyMembersNumberPage) GetPrevNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrevNumber.Get(), o.PrevNumber.IsSet()
}

// SetPrevNumber sets field value.
func (o *TeamHierarchyMembersNumberPage) SetPrevNumber(v int64) {
	o.PrevNumber.Set(&v)
}

// GetSize returns the Size field value.
func (o *TeamHierarchyMembersNumberPage) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersNumberPage) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *TeamHierarchyMembersNumberPage) SetSize(v int64) {
	o.Size = v
}

// GetTotal returns the Total field value.
func (o *TeamHierarchyMembersNumberPage) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersNumberPage) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *TeamHierarchyMembersNumberPage) SetTotal(v int64) {
	o.Total = v
}

// GetType returns the Type field value.
func (o *TeamHierarchyMembersNumberPage) GetType() TeamHierarchyMembersNumberPageType {
	if o == nil {
		var ret TeamHierarchyMembersNumberPageType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMembersNumberPage) GetTypeOk() (*TeamHierarchyMembersNumberPageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamHierarchyMembersNumberPage) SetType(v TeamHierarchyMembersNumberPageType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyMembersNumberPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["first_number"] = o.FirstNumber
	toSerialize["last_number"] = o.LastNumber.Get()
	toSerialize["next_number"] = o.NextNumber.Get()
	toSerialize["number"] = o.Number
	toSerialize["prev_number"] = o.PrevNumber.Get()
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamHierarchyMembersNumberPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FirstNumber *int64                              `json:"first_number"`
		LastNumber  datadog.NullableInt64               `json:"last_number"`
		NextNumber  datadog.NullableInt64               `json:"next_number"`
		Number      *int64                              `json:"number"`
		PrevNumber  datadog.NullableInt64               `json:"prev_number"`
		Size        *int64                              `json:"size"`
		Total       *int64                              `json:"total"`
		Type        *TeamHierarchyMembersNumberPageType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FirstNumber == nil {
		return fmt.Errorf("required field first_number missing")
	}
	if !all.LastNumber.IsSet() {
		return fmt.Errorf("required field last_number missing")
	}
	if !all.NextNumber.IsSet() {
		return fmt.Errorf("required field next_number missing")
	}
	if all.Number == nil {
		return fmt.Errorf("required field number missing")
	}
	if !all.PrevNumber.IsSet() {
		return fmt.Errorf("required field prev_number missing")
	}
	if all.Size == nil {
		return fmt.Errorf("required field size missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first_number", "last_number", "next_number", "number", "prev_number", "size", "total", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FirstNumber = *all.FirstNumber
	o.LastNumber = all.LastNumber
	o.NextNumber = all.NextNumber
	o.Number = *all.Number
	o.PrevNumber = all.PrevNumber
	o.Size = *all.Size
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
