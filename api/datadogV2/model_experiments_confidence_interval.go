// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsConfidenceInterval Lower and upper bounds of the reported statistical interval.
type ExperimentsConfidenceInterval struct {
	// Lower bound of the reported interval.
	Lower *float64 `json:"lower,omitempty"`
	// Upper bound of the reported interval.
	Upper *float64 `json:"upper,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsConfidenceInterval instantiates a new ExperimentsConfidenceInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsConfidenceInterval() *ExperimentsConfidenceInterval {
	this := ExperimentsConfidenceInterval{}
	return &this
}

// NewExperimentsConfidenceIntervalWithDefaults instantiates a new ExperimentsConfidenceInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsConfidenceIntervalWithDefaults() *ExperimentsConfidenceInterval {
	this := ExperimentsConfidenceInterval{}
	return &this
}

// GetLower returns the Lower field value if set, zero value otherwise.
func (o *ExperimentsConfidenceInterval) GetLower() float64 {
	if o == nil || o.Lower == nil {
		var ret float64
		return ret
	}
	return *o.Lower
}

// GetLowerOk returns a tuple with the Lower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsConfidenceInterval) GetLowerOk() (*float64, bool) {
	if o == nil || o.Lower == nil {
		return nil, false
	}
	return o.Lower, true
}

// HasLower returns a boolean if a field has been set.
func (o *ExperimentsConfidenceInterval) HasLower() bool {
	return o != nil && o.Lower != nil
}

// SetLower gets a reference to the given float64 and assigns it to the Lower field.
func (o *ExperimentsConfidenceInterval) SetLower(v float64) {
	o.Lower = &v
}

// GetUpper returns the Upper field value if set, zero value otherwise.
func (o *ExperimentsConfidenceInterval) GetUpper() float64 {
	if o == nil || o.Upper == nil {
		var ret float64
		return ret
	}
	return *o.Upper
}

// GetUpperOk returns a tuple with the Upper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsConfidenceInterval) GetUpperOk() (*float64, bool) {
	if o == nil || o.Upper == nil {
		return nil, false
	}
	return o.Upper, true
}

// HasUpper returns a boolean if a field has been set.
func (o *ExperimentsConfidenceInterval) HasUpper() bool {
	return o != nil && o.Upper != nil
}

// SetUpper gets a reference to the given float64 and assigns it to the Upper field.
func (o *ExperimentsConfidenceInterval) SetUpper(v float64) {
	o.Upper = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsConfidenceInterval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Lower != nil {
		toSerialize["lower"] = o.Lower
	}
	if o.Upper != nil {
		toSerialize["upper"] = o.Upper
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsConfidenceInterval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Lower *float64 `json:"lower,omitempty"`
		Upper *float64 `json:"upper,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"lower", "upper"})
	} else {
		return err
	}
	o.Lower = all.Lower
	o.Upper = all.Upper

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
