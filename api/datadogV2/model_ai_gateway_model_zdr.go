// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelZdr The Datadog side of a model's zero data retention eligibility. The customer side is the provider's
// `zdr_eligible` setting.
type AiGatewayModelZdr struct {
	// Datadog's coverage of the model's operations under zero data retention. A model is eligible when
	// coverage is `full` and its provider's `zdr_eligible` is `true`. `partial` means only the
	// `covered_operations` can be served: requests on other operations are not dispatched to it.
	Coverage AiGatewayModelZdrCoverage `json:"coverage"`
	// Operations covered by zero data retention.
	CoveredOperations []string `json:"covered_operations,omitempty"`
	// Operations not covered by zero data retention.
	UncoveredOperations []string `json:"uncovered_operations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelZdr instantiates a new AiGatewayModelZdr object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelZdr(coverage AiGatewayModelZdrCoverage) *AiGatewayModelZdr {
	this := AiGatewayModelZdr{}
	this.Coverage = coverage
	return &this
}

// NewAiGatewayModelZdrWithDefaults instantiates a new AiGatewayModelZdr object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelZdrWithDefaults() *AiGatewayModelZdr {
	this := AiGatewayModelZdr{}
	return &this
}

// GetCoverage returns the Coverage field value.
func (o *AiGatewayModelZdr) GetCoverage() AiGatewayModelZdrCoverage {
	if o == nil {
		var ret AiGatewayModelZdrCoverage
		return ret
	}
	return o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelZdr) GetCoverageOk() (*AiGatewayModelZdrCoverage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coverage, true
}

// SetCoverage sets field value.
func (o *AiGatewayModelZdr) SetCoverage(v AiGatewayModelZdrCoverage) {
	o.Coverage = v
}

// GetCoveredOperations returns the CoveredOperations field value if set, zero value otherwise.
func (o *AiGatewayModelZdr) GetCoveredOperations() []string {
	if o == nil || o.CoveredOperations == nil {
		var ret []string
		return ret
	}
	return o.CoveredOperations
}

// GetCoveredOperationsOk returns a tuple with the CoveredOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelZdr) GetCoveredOperationsOk() (*[]string, bool) {
	if o == nil || o.CoveredOperations == nil {
		return nil, false
	}
	return &o.CoveredOperations, true
}

// HasCoveredOperations returns a boolean if a field has been set.
func (o *AiGatewayModelZdr) HasCoveredOperations() bool {
	return o != nil && o.CoveredOperations != nil
}

// SetCoveredOperations gets a reference to the given []string and assigns it to the CoveredOperations field.
func (o *AiGatewayModelZdr) SetCoveredOperations(v []string) {
	o.CoveredOperations = v
}

// GetUncoveredOperations returns the UncoveredOperations field value if set, zero value otherwise.
func (o *AiGatewayModelZdr) GetUncoveredOperations() []string {
	if o == nil || o.UncoveredOperations == nil {
		var ret []string
		return ret
	}
	return o.UncoveredOperations
}

// GetUncoveredOperationsOk returns a tuple with the UncoveredOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelZdr) GetUncoveredOperationsOk() (*[]string, bool) {
	if o == nil || o.UncoveredOperations == nil {
		return nil, false
	}
	return &o.UncoveredOperations, true
}

// HasUncoveredOperations returns a boolean if a field has been set.
func (o *AiGatewayModelZdr) HasUncoveredOperations() bool {
	return o != nil && o.UncoveredOperations != nil
}

// SetUncoveredOperations gets a reference to the given []string and assigns it to the UncoveredOperations field.
func (o *AiGatewayModelZdr) SetUncoveredOperations(v []string) {
	o.UncoveredOperations = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelZdr) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["coverage"] = o.Coverage
	if o.CoveredOperations != nil {
		toSerialize["covered_operations"] = o.CoveredOperations
	}
	if o.UncoveredOperations != nil {
		toSerialize["uncovered_operations"] = o.UncoveredOperations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelZdr) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Coverage            *AiGatewayModelZdrCoverage `json:"coverage"`
		CoveredOperations   []string                   `json:"covered_operations,omitempty"`
		UncoveredOperations []string                   `json:"uncovered_operations,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Coverage == nil {
		return fmt.Errorf("required field coverage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"coverage", "covered_operations", "uncovered_operations"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Coverage.IsValid() {
		hasInvalidField = true
	} else {
		o.Coverage = *all.Coverage
	}
	o.CoveredOperations = all.CoveredOperations
	o.UncoveredOperations = all.UncoveredOperations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
