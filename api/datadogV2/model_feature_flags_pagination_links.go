// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagsPaginationLinks Pagination links. Follow `next` with the same authentication to retrieve the next page. Links use `page[limit]` and `page[offset]` and preserve the filters. Navigation links are omitted when not applicable.
type FeatureFlagsPaginationLinks struct {
	// Link to the first page.
	First *string `json:"first,omitempty"`
	// Link to the last page.
	Last *string `json:"last,omitempty"`
	// Link to the next page, when available.
	Next *string `json:"next,omitempty"`
	// Link to the previous page, when available.
	Prev *string `json:"prev,omitempty"`
	// Link to the current page.
	Self *string `json:"self,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagsPaginationLinks instantiates a new FeatureFlagsPaginationLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagsPaginationLinks() *FeatureFlagsPaginationLinks {
	this := FeatureFlagsPaginationLinks{}
	return &this
}

// NewFeatureFlagsPaginationLinksWithDefaults instantiates a new FeatureFlagsPaginationLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagsPaginationLinksWithDefaults() *FeatureFlagsPaginationLinks {
	this := FeatureFlagsPaginationLinks{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *FeatureFlagsPaginationLinks) GetFirst() string {
	if o == nil || o.First == nil {
		var ret string
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagsPaginationLinks) GetFirstOk() (*string, bool) {
	if o == nil || o.First == nil {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationLinks) HasFirst() bool {
	return o != nil && o.First != nil
}

// SetFirst gets a reference to the given string and assigns it to the First field.
func (o *FeatureFlagsPaginationLinks) SetFirst(v string) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *FeatureFlagsPaginationLinks) GetLast() string {
	if o == nil || o.Last == nil {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagsPaginationLinks) GetLastOk() (*string, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationLinks) HasLast() bool {
	return o != nil && o.Last != nil
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *FeatureFlagsPaginationLinks) SetLast(v string) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *FeatureFlagsPaginationLinks) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagsPaginationLinks) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationLinks) HasNext() bool {
	return o != nil && o.Next != nil
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *FeatureFlagsPaginationLinks) SetNext(v string) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *FeatureFlagsPaginationLinks) GetPrev() string {
	if o == nil || o.Prev == nil {
		var ret string
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagsPaginationLinks) GetPrevOk() (*string, bool) {
	if o == nil || o.Prev == nil {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationLinks) HasPrev() bool {
	return o != nil && o.Prev != nil
}

// SetPrev gets a reference to the given string and assigns it to the Prev field.
func (o *FeatureFlagsPaginationLinks) SetPrev(v string) {
	o.Prev = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *FeatureFlagsPaginationLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagsPaginationLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *FeatureFlagsPaginationLinks) HasSelf() bool {
	return o != nil && o.Self != nil
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *FeatureFlagsPaginationLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagsPaginationLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.First != nil {
		toSerialize["first"] = o.First
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Prev != nil {
		toSerialize["prev"] = o.Prev
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagsPaginationLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		First *string `json:"first,omitempty"`
		Last  *string `json:"last,omitempty"`
		Next  *string `json:"next,omitempty"`
		Prev  *string `json:"prev,omitempty"`
		Self  *string `json:"self,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first", "last", "next", "prev", "self"})
	} else {
		return err
	}
	o.First = all.First
	o.Last = all.Last
	o.Next = all.Next
	o.Prev = all.Prev
	o.Self = all.Self

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
