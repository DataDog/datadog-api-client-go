// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsCoverageTimeseriesResponseAttributes Attributes of a timeseries coverage response.
type CommitmentsCoverageTimeseriesResponseAttributes struct {
	// A coverage timeseries metric. For AWS and Azure, `series` holds the `on_demand_only`, `spot_only`, `ri`,
	// and `sp` buckets, plus `ppp` for Azure pre-purchase plans when enabled. For GCP, `series` holds the
	// `on_demand_only` and `flex_cud` buckets.
	Cost CommitmentsCoverageTimeseriesMetric `json:"cost"`
	// A coverage timeseries metric. For AWS and Azure, `series` holds the `on_demand_only`, `spot_only`, `ri`,
	// and `sp` buckets, plus `ppp` for Azure pre-purchase plans when enabled. For GCP, `series` holds the
	// `on_demand_only` and `flex_cud` buckets.
	Hours CommitmentsCoverageTimeseriesMetric `json:"hours"`
	// The query that produces the response values, in the same format as dashboard and notebook widget
	// requests, so a client can graph the values in a widget. Each formula's `id` names the response value
	// that it produces.
	QueryDefinition *CommitmentsQueryDefinition `json:"query_definition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsCoverageTimeseriesResponseAttributes instantiates a new CommitmentsCoverageTimeseriesResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsCoverageTimeseriesResponseAttributes(cost CommitmentsCoverageTimeseriesMetric, hours CommitmentsCoverageTimeseriesMetric) *CommitmentsCoverageTimeseriesResponseAttributes {
	this := CommitmentsCoverageTimeseriesResponseAttributes{}
	this.Cost = cost
	this.Hours = hours
	return &this
}

// NewCommitmentsCoverageTimeseriesResponseAttributesWithDefaults instantiates a new CommitmentsCoverageTimeseriesResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsCoverageTimeseriesResponseAttributesWithDefaults() *CommitmentsCoverageTimeseriesResponseAttributes {
	this := CommitmentsCoverageTimeseriesResponseAttributes{}
	return &this
}

// GetCost returns the Cost field value.
func (o *CommitmentsCoverageTimeseriesResponseAttributes) GetCost() CommitmentsCoverageTimeseriesMetric {
	if o == nil {
		var ret CommitmentsCoverageTimeseriesMetric
		return ret
	}
	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageTimeseriesResponseAttributes) GetCostOk() (*CommitmentsCoverageTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value.
func (o *CommitmentsCoverageTimeseriesResponseAttributes) SetCost(v CommitmentsCoverageTimeseriesMetric) {
	o.Cost = v
}

// GetHours returns the Hours field value.
func (o *CommitmentsCoverageTimeseriesResponseAttributes) GetHours() CommitmentsCoverageTimeseriesMetric {
	if o == nil {
		var ret CommitmentsCoverageTimeseriesMetric
		return ret
	}
	return o.Hours
}

// GetHoursOk returns a tuple with the Hours field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageTimeseriesResponseAttributes) GetHoursOk() (*CommitmentsCoverageTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hours, true
}

// SetHours sets field value.
func (o *CommitmentsCoverageTimeseriesResponseAttributes) SetHours(v CommitmentsCoverageTimeseriesMetric) {
	o.Hours = v
}

// GetQueryDefinition returns the QueryDefinition field value if set, zero value otherwise.
func (o *CommitmentsCoverageTimeseriesResponseAttributes) GetQueryDefinition() CommitmentsQueryDefinition {
	if o == nil || o.QueryDefinition == nil {
		var ret CommitmentsQueryDefinition
		return ret
	}
	return *o.QueryDefinition
}

// GetQueryDefinitionOk returns a tuple with the QueryDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageTimeseriesResponseAttributes) GetQueryDefinitionOk() (*CommitmentsQueryDefinition, bool) {
	if o == nil || o.QueryDefinition == nil {
		return nil, false
	}
	return o.QueryDefinition, true
}

// HasQueryDefinition returns a boolean if a field has been set.
func (o *CommitmentsCoverageTimeseriesResponseAttributes) HasQueryDefinition() bool {
	return o != nil && o.QueryDefinition != nil
}

// SetQueryDefinition gets a reference to the given CommitmentsQueryDefinition and assigns it to the QueryDefinition field.
func (o *CommitmentsCoverageTimeseriesResponseAttributes) SetQueryDefinition(v CommitmentsQueryDefinition) {
	o.QueryDefinition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsCoverageTimeseriesResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cost"] = o.Cost
	toSerialize["hours"] = o.Hours
	if o.QueryDefinition != nil {
		toSerialize["query_definition"] = o.QueryDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsCoverageTimeseriesResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cost            *CommitmentsCoverageTimeseriesMetric `json:"cost"`
		Hours           *CommitmentsCoverageTimeseriesMetric `json:"hours"`
		QueryDefinition *CommitmentsQueryDefinition          `json:"query_definition,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Cost == nil {
		return fmt.Errorf("required field cost missing")
	}
	if all.Hours == nil {
		return fmt.Errorf("required field hours missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cost", "hours", "query_definition"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Cost.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cost = *all.Cost
	if all.Hours.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Hours = *all.Hours
	if all.QueryDefinition != nil && all.QueryDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.QueryDefinition = all.QueryDefinition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
