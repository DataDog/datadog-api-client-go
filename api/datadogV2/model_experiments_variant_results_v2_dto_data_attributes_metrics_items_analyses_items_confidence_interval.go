// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval Lower and upper bounds of the reported statistical interval.
type ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval struct {
	// Lower bound of the reported interval.
	Lower *float64 `json:"lower,omitempty"`
	// Upper bound of the reported interval.
	Upper *float64 `json:"upper,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval instantiates a new ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval() *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval {
	this := ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval{}
	return &this
}

// NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceIntervalWithDefaults instantiates a new ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceIntervalWithDefaults() *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval {
	this := ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval{}
	return &this
}

// GetLower returns the Lower field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval) GetLower() float64 {
	if o == nil || o.Lower == nil {
		var ret float64
		return ret
	}
	return *o.Lower
}

// GetLowerOk returns a tuple with the Lower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval) GetLowerOk() (*float64, bool) {
	if o == nil || o.Lower == nil {
		return nil, false
	}
	return o.Lower, true
}

// HasLower returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval) HasLower() bool {
	return o != nil && o.Lower != nil
}

// SetLower gets a reference to the given float64 and assigns it to the Lower field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval) SetLower(v float64) {
	o.Lower = &v
}

// GetUpper returns the Upper field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval) GetUpper() float64 {
	if o == nil || o.Upper == nil {
		var ret float64
		return ret
	}
	return *o.Upper
}

// GetUpperOk returns a tuple with the Upper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval) GetUpperOk() (*float64, bool) {
	if o == nil || o.Upper == nil {
		return nil, false
	}
	return o.Upper, true
}

// HasUpper returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval) HasUpper() bool {
	return o != nil && o.Upper != nil
}

// SetUpper gets a reference to the given float64 and assigns it to the Upper field.
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval) SetUpper(v float64) {
	o.Upper = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsVariantResultsV2DTODataAttributesMetricsItemsAnalysesItemsConfidenceInterval) UnmarshalJSON(bytes []byte) (err error) {
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
