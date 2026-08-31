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
	// The number of supported indicators accepted.
	Accepted int64 `json:"accepted"`
	// The number of indicators with patterns that could not be parsed.
	Invalid int64 `json:"invalid"`
	// The number of indicator objects with an unsupported STIX version or a pattern that produced no supported observable values.
	Unsupported int64 `json:"unsupported"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSTIXIngestResponseAttributes instantiates a new STIXIngestResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSTIXIngestResponseAttributes(accepted int64, invalid int64, unsupported int64) *STIXIngestResponseAttributes {
	this := STIXIngestResponseAttributes{}
	this.Accepted = accepted
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

// GetAccepted returns the Accepted field value.
func (o *STIXIngestResponseAttributes) GetAccepted() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Accepted
}

// GetAcceptedOk returns a tuple with the Accepted field value
// and a boolean to check if the value has been set.
func (o *STIXIngestResponseAttributes) GetAcceptedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accepted, true
}

// SetAccepted sets field value.
func (o *STIXIngestResponseAttributes) SetAccepted(v int64) {
	o.Accepted = v
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
	toSerialize["accepted"] = o.Accepted
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
		Accepted    *int64 `json:"accepted"`
		Invalid     *int64 `json:"invalid"`
		Unsupported *int64 `json:"unsupported"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Accepted == nil {
		return fmt.Errorf("required field accepted missing")
	}
	if all.Invalid == nil {
		return fmt.Errorf("required field invalid missing")
	}
	if all.Unsupported == nil {
		return fmt.Errorf("required field unsupported missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accepted", "invalid", "unsupported"})
	} else {
		return err
	}
	o.Accepted = *all.Accepted
	o.Invalid = *all.Invalid
	o.Unsupported = *all.Unsupported

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
