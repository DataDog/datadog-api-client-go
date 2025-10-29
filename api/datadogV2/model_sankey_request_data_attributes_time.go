// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyRequestDataAttributesTime
type SankeyRequestDataAttributesTime struct {
	//
	From *int64 `json:"from,omitempty"`
	//
	To *int64 `json:"to,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSankeyRequestDataAttributesTime instantiates a new SankeyRequestDataAttributesTime object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyRequestDataAttributesTime() *SankeyRequestDataAttributesTime {
	this := SankeyRequestDataAttributesTime{}
	return &this
}

// NewSankeyRequestDataAttributesTimeWithDefaults instantiates a new SankeyRequestDataAttributesTime object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyRequestDataAttributesTimeWithDefaults() *SankeyRequestDataAttributesTime {
	this := SankeyRequestDataAttributesTime{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributesTime) GetFrom() int64 {
	if o == nil || o.From == nil {
		var ret int64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributesTime) GetFromOk() (*int64, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributesTime) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given int64 and assigns it to the From field.
func (o *SankeyRequestDataAttributesTime) SetFrom(v int64) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributesTime) GetTo() int64 {
	if o == nil || o.To == nil {
		var ret int64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributesTime) GetToOk() (*int64, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributesTime) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given int64 and assigns it to the To field.
func (o *SankeyRequestDataAttributesTime) SetTo(v int64) {
	o.To = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyRequestDataAttributesTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SankeyRequestDataAttributesTime) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From *int64 `json:"from,omitempty"`
		To   *int64 `json:"to,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "to"})
	} else {
		return err
	}
	o.From = all.From
	o.To = all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
