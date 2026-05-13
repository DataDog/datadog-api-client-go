// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListPROutputsAttributes Attributes of a PR outputs response.
type ListPROutputsAttributes struct {
	// The list of pull requests created by the workflow execution.
	PrOutputs []PROutput `json:"pr_outputs"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListPROutputsAttributes instantiates a new ListPROutputsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListPROutputsAttributes(prOutputs []PROutput) *ListPROutputsAttributes {
	this := ListPROutputsAttributes{}
	this.PrOutputs = prOutputs
	return &this
}

// NewListPROutputsAttributesWithDefaults instantiates a new ListPROutputsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListPROutputsAttributesWithDefaults() *ListPROutputsAttributes {
	this := ListPROutputsAttributes{}
	return &this
}

// GetPrOutputs returns the PrOutputs field value.
func (o *ListPROutputsAttributes) GetPrOutputs() []PROutput {
	if o == nil {
		var ret []PROutput
		return ret
	}
	return o.PrOutputs
}

// GetPrOutputsOk returns a tuple with the PrOutputs field value
// and a boolean to check if the value has been set.
func (o *ListPROutputsAttributes) GetPrOutputsOk() (*[]PROutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrOutputs, true
}

// SetPrOutputs sets field value.
func (o *ListPROutputsAttributes) SetPrOutputs(v []PROutput) {
	o.PrOutputs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListPROutputsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["pr_outputs"] = o.PrOutputs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListPROutputsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PrOutputs *[]PROutput `json:"pr_outputs"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PrOutputs == nil {
		return fmt.Errorf("required field pr_outputs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"pr_outputs"})
	} else {
		return err
	}
	o.PrOutputs = *all.PrOutputs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
