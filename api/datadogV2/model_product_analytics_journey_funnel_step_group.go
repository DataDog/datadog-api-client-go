// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyFunnelStepGroup Breakdown of a funnel step for one combination of group-by values.
type ProductAnalyticsJourneyFunnelStepGroup struct {
	// Number of entities in this group that reached the next step.
	ConversionCount int64 `json:"conversion_count"`
	// Elapsed time statistics (min/max/avg in milliseconds).
	ElapsedTimeToNextStep ProductAnalyticsElapsedTime `json:"elapsed_time_to_next_step"`
	// Group-by values identifying this cohort.
	GroupTags []string `json:"group_tags"`
	// Value of the computed metric for this group at this step.
	Value float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyFunnelStepGroup instantiates a new ProductAnalyticsJourneyFunnelStepGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyFunnelStepGroup(conversionCount int64, elapsedTimeToNextStep ProductAnalyticsElapsedTime, groupTags []string, value float64) *ProductAnalyticsJourneyFunnelStepGroup {
	this := ProductAnalyticsJourneyFunnelStepGroup{}
	this.ConversionCount = conversionCount
	this.ElapsedTimeToNextStep = elapsedTimeToNextStep
	this.GroupTags = groupTags
	this.Value = value
	return &this
}

// NewProductAnalyticsJourneyFunnelStepGroupWithDefaults instantiates a new ProductAnalyticsJourneyFunnelStepGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyFunnelStepGroupWithDefaults() *ProductAnalyticsJourneyFunnelStepGroup {
	this := ProductAnalyticsJourneyFunnelStepGroup{}
	return &this
}

// GetConversionCount returns the ConversionCount field value.
func (o *ProductAnalyticsJourneyFunnelStepGroup) GetConversionCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ConversionCount
}

// GetConversionCountOk returns a tuple with the ConversionCount field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelStepGroup) GetConversionCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversionCount, true
}

// SetConversionCount sets field value.
func (o *ProductAnalyticsJourneyFunnelStepGroup) SetConversionCount(v int64) {
	o.ConversionCount = v
}

// GetElapsedTimeToNextStep returns the ElapsedTimeToNextStep field value.
func (o *ProductAnalyticsJourneyFunnelStepGroup) GetElapsedTimeToNextStep() ProductAnalyticsElapsedTime {
	if o == nil {
		var ret ProductAnalyticsElapsedTime
		return ret
	}
	return o.ElapsedTimeToNextStep
}

// GetElapsedTimeToNextStepOk returns a tuple with the ElapsedTimeToNextStep field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelStepGroup) GetElapsedTimeToNextStepOk() (*ProductAnalyticsElapsedTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ElapsedTimeToNextStep, true
}

// SetElapsedTimeToNextStep sets field value.
func (o *ProductAnalyticsJourneyFunnelStepGroup) SetElapsedTimeToNextStep(v ProductAnalyticsElapsedTime) {
	o.ElapsedTimeToNextStep = v
}

// GetGroupTags returns the GroupTags field value.
func (o *ProductAnalyticsJourneyFunnelStepGroup) GetGroupTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GroupTags
}

// GetGroupTagsOk returns a tuple with the GroupTags field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelStepGroup) GetGroupTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupTags, true
}

// SetGroupTags sets field value.
func (o *ProductAnalyticsJourneyFunnelStepGroup) SetGroupTags(v []string) {
	o.GroupTags = v
}

// GetValue returns the Value field value.
func (o *ProductAnalyticsJourneyFunnelStepGroup) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyFunnelStepGroup) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ProductAnalyticsJourneyFunnelStepGroup) SetValue(v float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyFunnelStepGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["conversion_count"] = o.ConversionCount
	toSerialize["elapsed_time_to_next_step"] = o.ElapsedTimeToNextStep
	toSerialize["group_tags"] = o.GroupTags
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyFunnelStepGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConversionCount       *int64                       `json:"conversion_count"`
		ElapsedTimeToNextStep *ProductAnalyticsElapsedTime `json:"elapsed_time_to_next_step"`
		GroupTags             *[]string                    `json:"group_tags"`
		Value                 *float64                     `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConversionCount == nil {
		return fmt.Errorf("required field conversion_count missing")
	}
	if all.ElapsedTimeToNextStep == nil {
		return fmt.Errorf("required field elapsed_time_to_next_step missing")
	}
	if all.GroupTags == nil {
		return fmt.Errorf("required field group_tags missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conversion_count", "elapsed_time_to_next_step", "group_tags", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConversionCount = *all.ConversionCount
	if all.ElapsedTimeToNextStep.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ElapsedTimeToNextStep = *all.ElapsedTimeToNextStep
	o.GroupTags = *all.GroupTags
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
