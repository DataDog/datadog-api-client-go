// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayBudgetInstanceDimensionAttributes Values of one budget split dimension.
type AiGatewayBudgetInstanceDimensionAttributes struct {
	// Budget identifier.
	BudgetId uuid.UUID `json:"budget_id"`
	// Split dimension name, as used in `split_dims` and the `split_dim` filter.
	Dim string `json:"dim"`
	// Whether identifying instance details were withheld. Aggregate totals still include restricted instances.
	RestrictedDetails bool `json:"restricted_details"`
	// Whether more values matched than were returned. Narrow the results with `query` or `value` to see them.
	Truncated bool `json:"truncated"`
	// Values of this dimension across matching instance counters, highest spend first and at most `limit` of them. Values are counted under every `split_dim` filter except the one for this dimension.
	Values []AiGatewayBudgetInstanceDimensionValue `json:"values"`
	// Start of the accounting window.
	WindowStart time.Time `json:"window_start"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayBudgetInstanceDimensionAttributes instantiates a new AiGatewayBudgetInstanceDimensionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayBudgetInstanceDimensionAttributes(budgetId uuid.UUID, dim string, restrictedDetails bool, truncated bool, values []AiGatewayBudgetInstanceDimensionValue, windowStart time.Time) *AiGatewayBudgetInstanceDimensionAttributes {
	this := AiGatewayBudgetInstanceDimensionAttributes{}
	this.BudgetId = budgetId
	this.Dim = dim
	this.RestrictedDetails = restrictedDetails
	this.Truncated = truncated
	this.Values = values
	this.WindowStart = windowStart
	return &this
}

// NewAiGatewayBudgetInstanceDimensionAttributesWithDefaults instantiates a new AiGatewayBudgetInstanceDimensionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayBudgetInstanceDimensionAttributesWithDefaults() *AiGatewayBudgetInstanceDimensionAttributes {
	this := AiGatewayBudgetInstanceDimensionAttributes{}
	return &this
}

// GetBudgetId returns the BudgetId field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetBudgetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.BudgetId
}

// GetBudgetIdOk returns a tuple with the BudgetId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetBudgetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetId, true
}

// SetBudgetId sets field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) SetBudgetId(v uuid.UUID) {
	o.BudgetId = v
}

// GetDim returns the Dim field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetDim() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Dim
}

// GetDimOk returns a tuple with the Dim field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetDimOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dim, true
}

// SetDim sets field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) SetDim(v string) {
	o.Dim = v
}

// GetRestrictedDetails returns the RestrictedDetails field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetRestrictedDetails() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.RestrictedDetails
}

// GetRestrictedDetailsOk returns a tuple with the RestrictedDetails field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetRestrictedDetailsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestrictedDetails, true
}

// SetRestrictedDetails sets field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) SetRestrictedDetails(v bool) {
	o.RestrictedDetails = v
}

// GetTruncated returns the Truncated field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) SetTruncated(v bool) {
	o.Truncated = v
}

// GetValues returns the Values field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetValues() []AiGatewayBudgetInstanceDimensionValue {
	if o == nil {
		var ret []AiGatewayBudgetInstanceDimensionValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetValuesOk() (*[]AiGatewayBudgetInstanceDimensionValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) SetValues(v []AiGatewayBudgetInstanceDimensionValue) {
	o.Values = v
}

// GetWindowStart returns the WindowStart field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetWindowStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.WindowStart
}

// GetWindowStartOk returns a tuple with the WindowStart field value
// and a boolean to check if the value has been set.
func (o *AiGatewayBudgetInstanceDimensionAttributes) GetWindowStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WindowStart, true
}

// SetWindowStart sets field value.
func (o *AiGatewayBudgetInstanceDimensionAttributes) SetWindowStart(v time.Time) {
	o.WindowStart = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayBudgetInstanceDimensionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["budget_id"] = o.BudgetId
	toSerialize["dim"] = o.Dim
	toSerialize["restricted_details"] = o.RestrictedDetails
	toSerialize["truncated"] = o.Truncated
	toSerialize["values"] = o.Values
	if o.WindowStart.Nanosecond() == 0 {
		toSerialize["window_start"] = o.WindowStart.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["window_start"] = o.WindowStart.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayBudgetInstanceDimensionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BudgetId          *uuid.UUID                               `json:"budget_id"`
		Dim               *string                                  `json:"dim"`
		RestrictedDetails *bool                                    `json:"restricted_details"`
		Truncated         *bool                                    `json:"truncated"`
		Values            *[]AiGatewayBudgetInstanceDimensionValue `json:"values"`
		WindowStart       *time.Time                               `json:"window_start"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BudgetId == nil {
		return fmt.Errorf("required field budget_id missing")
	}
	if all.Dim == nil {
		return fmt.Errorf("required field dim missing")
	}
	if all.RestrictedDetails == nil {
		return fmt.Errorf("required field restricted_details missing")
	}
	if all.Truncated == nil {
		return fmt.Errorf("required field truncated missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	if all.WindowStart == nil {
		return fmt.Errorf("required field window_start missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"budget_id", "dim", "restricted_details", "truncated", "values", "window_start"})
	} else {
		return err
	}
	o.BudgetId = *all.BudgetId
	o.Dim = *all.Dim
	o.RestrictedDetails = *all.RestrictedDetails
	o.Truncated = *all.Truncated
	o.Values = *all.Values
	o.WindowStart = *all.WindowStart

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
