// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsTrafficSummaryV2DTODataAttributes Details of the experiment traffic summary.
type ExperimentsTrafficSummaryV2DTODataAttributes struct {
	// Whether the observed variant traffic is imbalanced.
	IsTrafficImbalanced *bool `json:"is_traffic_imbalanced,omitempty"`
	// Total number of subjects included in the traffic summary.
	TotalSubjects *int64 `json:"total_subjects,omitempty"`
	// Exposure counts for each experiment variant.
	Variants []ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems `json:"variants,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsTrafficSummaryV2DTODataAttributes instantiates a new ExperimentsTrafficSummaryV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsTrafficSummaryV2DTODataAttributes() *ExperimentsTrafficSummaryV2DTODataAttributes {
	this := ExperimentsTrafficSummaryV2DTODataAttributes{}
	return &this
}

// NewExperimentsTrafficSummaryV2DTODataAttributesWithDefaults instantiates a new ExperimentsTrafficSummaryV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsTrafficSummaryV2DTODataAttributesWithDefaults() *ExperimentsTrafficSummaryV2DTODataAttributes {
	this := ExperimentsTrafficSummaryV2DTODataAttributes{}
	return &this
}

// GetIsTrafficImbalanced returns the IsTrafficImbalanced field value if set, zero value otherwise.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) GetIsTrafficImbalanced() bool {
	if o == nil || o.IsTrafficImbalanced == nil {
		var ret bool
		return ret
	}
	return *o.IsTrafficImbalanced
}

// GetIsTrafficImbalancedOk returns a tuple with the IsTrafficImbalanced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) GetIsTrafficImbalancedOk() (*bool, bool) {
	if o == nil || o.IsTrafficImbalanced == nil {
		return nil, false
	}
	return o.IsTrafficImbalanced, true
}

// HasIsTrafficImbalanced returns a boolean if a field has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) HasIsTrafficImbalanced() bool {
	return o != nil && o.IsTrafficImbalanced != nil
}

// SetIsTrafficImbalanced gets a reference to the given bool and assigns it to the IsTrafficImbalanced field.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) SetIsTrafficImbalanced(v bool) {
	o.IsTrafficImbalanced = &v
}

// GetTotalSubjects returns the TotalSubjects field value if set, zero value otherwise.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) GetTotalSubjects() int64 {
	if o == nil || o.TotalSubjects == nil {
		var ret int64
		return ret
	}
	return *o.TotalSubjects
}

// GetTotalSubjectsOk returns a tuple with the TotalSubjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) GetTotalSubjectsOk() (*int64, bool) {
	if o == nil || o.TotalSubjects == nil {
		return nil, false
	}
	return o.TotalSubjects, true
}

// HasTotalSubjects returns a boolean if a field has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) HasTotalSubjects() bool {
	return o != nil && o.TotalSubjects != nil
}

// SetTotalSubjects gets a reference to the given int64 and assigns it to the TotalSubjects field.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) SetTotalSubjects(v int64) {
	o.TotalSubjects = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) GetVariants() []ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems {
	if o == nil || o.Variants == nil {
		var ret []ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) GetVariantsOk() (*[]ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems, bool) {
	if o == nil || o.Variants == nil {
		return nil, false
	}
	return &o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) HasVariants() bool {
	return o != nil && o.Variants != nil
}

// SetVariants gets a reference to the given []ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems and assigns it to the Variants field.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) SetVariants(v []ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) {
	o.Variants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsTrafficSummaryV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsTrafficImbalanced != nil {
		toSerialize["is_traffic_imbalanced"] = o.IsTrafficImbalanced
	}
	if o.TotalSubjects != nil {
		toSerialize["total_subjects"] = o.TotalSubjects
	}
	if o.Variants != nil {
		toSerialize["variants"] = o.Variants
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsTrafficSummaryV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsTrafficImbalanced *bool                                                       `json:"is_traffic_imbalanced,omitempty"`
		TotalSubjects       *int64                                                      `json:"total_subjects,omitempty"`
		Variants            []ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems `json:"variants,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_traffic_imbalanced", "total_subjects", "variants"})
	} else {
		return err
	}
	o.IsTrafficImbalanced = all.IsTrafficImbalanced
	o.TotalSubjects = all.TotalSubjects
	o.Variants = all.Variants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
