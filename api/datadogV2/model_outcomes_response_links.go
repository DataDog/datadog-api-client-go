// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutcomesResponseLinks Pagination links for rule outcomes.
type OutcomesResponseLinks struct {
	// Link for the first set of results.
	First *string `json:"first,omitempty"`
	// Link for the last set of results.
	Last *string `json:"last,omitempty"`
	// Link for the next set of results.
	Next *string `json:"next,omitempty"`
	// Link for the previous set of results.
	Previous *string `json:"previous,omitempty"`
	// Link for the current set of results.
	Self *string `json:"self,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOutcomesResponseLinks instantiates a new OutcomesResponseLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOutcomesResponseLinks() *OutcomesResponseLinks {
	this := OutcomesResponseLinks{}
	return &this
}

// NewOutcomesResponseLinksWithDefaults instantiates a new OutcomesResponseLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOutcomesResponseLinksWithDefaults() *OutcomesResponseLinks {
	this := OutcomesResponseLinks{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *OutcomesResponseLinks) GetFirst() string {
	if o == nil || o.First == nil {
		var ret string
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesResponseLinks) GetFirstOk() (*string, bool) {
	if o == nil || o.First == nil {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *OutcomesResponseLinks) HasFirst() bool {
	return o != nil && o.First != nil
}

// SetFirst gets a reference to the given string and assigns it to the First field.
func (o *OutcomesResponseLinks) SetFirst(v string) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *OutcomesResponseLinks) GetLast() string {
	if o == nil || o.Last == nil {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesResponseLinks) GetLastOk() (*string, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *OutcomesResponseLinks) HasLast() bool {
	return o != nil && o.Last != nil
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *OutcomesResponseLinks) SetLast(v string) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *OutcomesResponseLinks) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesResponseLinks) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *OutcomesResponseLinks) HasNext() bool {
	return o != nil && o.Next != nil
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *OutcomesResponseLinks) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *OutcomesResponseLinks) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesResponseLinks) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *OutcomesResponseLinks) HasPrevious() bool {
	return o != nil && o.Previous != nil
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *OutcomesResponseLinks) SetPrevious(v string) {
	o.Previous = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *OutcomesResponseLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesResponseLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *OutcomesResponseLinks) HasSelf() bool {
	return o != nil && o.Self != nil
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *OutcomesResponseLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OutcomesResponseLinks) MarshalJSON() ([]byte, error) {
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
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
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
func (o *OutcomesResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		First    *string `json:"first,omitempty"`
		Last     *string `json:"last,omitempty"`
		Next     *string `json:"next,omitempty"`
		Previous *string `json:"previous,omitempty"`
		Self     *string `json:"self,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"first", "last", "next", "previous", "self"})
	} else {
		return err
	}
	o.First = all.First
	o.Last = all.Last
	o.Next = all.Next
	o.Previous = all.Previous
	o.Self = all.Self

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
