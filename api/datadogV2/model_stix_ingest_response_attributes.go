// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// STIXIngestResponseAttributes Counters describing the result of the STIX ingestion request.
type STIXIngestResponseAttributes struct {
	// The number of supported indicators added.
	Added int64 `json:"added"`
	// The number of indicators with patterns that could not be parsed.
	Invalid int64 `json:"invalid"`
	// The number of unsupported objects or patterns.
	Unsupported int64 `json:"unsupported"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSTIXIngestResponseAttributes instantiates a new STIXIngestResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSTIXIngestResponseAttributes(added int64, invalid int64, unsupported int64) *STIXIngestResponseAttributes {
	this := STIXIngestResponseAttributes{}
	this.Added = added
	this.Invalid = invalid
	this.Unsupported = unsupported
	return &this
}

// NewSTIXIngestResponseAttributesWithDefaults instantiates a new STIXIngestResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSTIXIngestResponseAttributesWithDefaults() *STIXIngestResponseAttributes {
	this := STIXIngestResponseAttributes{}
	return &this
}

// GetAdded returns the Added field value.
func (o *STIXIngestResponseAttributes) GetAdded() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Added
}

// GetAddedOk returns a tuple with the Added field value
// and a boolean to check if the value has been set.
func (o *STIXIngestResponseAttributes) GetAddedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Added, true
}

// SetAdded sets field value.
func (o *STIXIngestResponseAttributes) SetAdded(v int64) {
	o.Added = v
}

// GetInvalid returns the Invalid field value.
func (o *STIXIngestResponseAttributes) GetInvalid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Invalid
}

// GetInvalidOk returns a tuple with the Invalid field value
// and a boolean to check if the value has been set.
func (o *STIXIngestResponseAttributes) GetInvalidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invalid, true
}

// SetInvalid sets field value.
func (o *STIXIngestResponseAttributes) SetInvalid(v int64) {
	o.Invalid = v
}

// GetUnsupported returns the Unsupported field value.
func (o *STIXIngestResponseAttributes) GetUnsupported() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Unsupported
}

// GetUnsupportedOk returns a tuple with the Unsupported field value
// and a boolean to check if the value has been set.
func (o *STIXIngestResponseAttributes) GetUnsupportedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unsupported, true
}

// SetUnsupported sets field value.
func (o *STIXIngestResponseAttributes) SetUnsupported(v int64) {
	o.Unsupported = v
}

// MarshalJSON serializes the struct using spec logic.
func (o STIXIngestResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["added"] = o.Added
	toSerialize["invalid"] = o.Invalid
	toSerialize["unsupported"] = o.Unsupported

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *STIXIngestResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Added       *int64 `json:"added"`
		Invalid     *int64 `json:"invalid"`
		Unsupported *int64 `json:"unsupported"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Added == nil {
		return fmt.Errorf("required field added missing")
	}
	if all.Invalid == nil {
		return fmt.Errorf("required field invalid missing")
	}
	if all.Unsupported == nil {
		return fmt.Errorf("required field unsupported missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"added", "invalid", "unsupported"})
	} else {
		return err
	}
	o.Added = *all.Added
	o.Invalid = *all.Invalid
	o.Unsupported = *all.Unsupported

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
