// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsSavingsTimeseriesResponseAttributes Attributes of a timeseries savings response.
type CommitmentsSavingsTimeseriesResponseAttributes struct {
	// A timeseries metric containing timestamps, series values, and optional unit metadata.
	ActualCost CommitmentsTimeseriesMetric `json:"actual_cost"`
	// A timeseries metric containing timestamps, series values, and optional unit metadata.
	EffectiveSavingsRate CommitmentsTimeseriesMetric `json:"effective_savings_rate"`
	// Metadata reporting that the cost data lacks the unused costs of some commitments, as with Azure exports
	// scoped below the billing profile. Utilization is then overstated, and realized savings and effective
	// savings rate are upper bounds. Omitted when the unused costs are available.
	Meta *CommitmentsUnusedCostMeta `json:"meta,omitempty"`
	// A timeseries metric containing timestamps, series values, and optional unit metadata.
	OnDemandEquivalentCost CommitmentsTimeseriesMetric `json:"on_demand_equivalent_cost"`
	// The query that produces the response values, in the same format as dashboard and notebook widget
	// requests, so a client can graph the values in a widget. Each formula's `id` names the response value
	// that it produces.
	QueryDefinition *CommitmentsQueryDefinition `json:"query_definition,omitempty"`
	// A timeseries metric containing timestamps, series values, and optional unit metadata.
	RealizedSavings CommitmentsTimeseriesMetric `json:"realized_savings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsSavingsTimeseriesResponseAttributes instantiates a new CommitmentsSavingsTimeseriesResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsSavingsTimeseriesResponseAttributes(actualCost CommitmentsTimeseriesMetric, effectiveSavingsRate CommitmentsTimeseriesMetric, onDemandEquivalentCost CommitmentsTimeseriesMetric, realizedSavings CommitmentsTimeseriesMetric) *CommitmentsSavingsTimeseriesResponseAttributes {
	this := CommitmentsSavingsTimeseriesResponseAttributes{}
	this.ActualCost = actualCost
	this.EffectiveSavingsRate = effectiveSavingsRate
	this.OnDemandEquivalentCost = onDemandEquivalentCost
	this.RealizedSavings = realizedSavings
	return &this
}

// NewCommitmentsSavingsTimeseriesResponseAttributesWithDefaults instantiates a new CommitmentsSavingsTimeseriesResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsSavingsTimeseriesResponseAttributesWithDefaults() *CommitmentsSavingsTimeseriesResponseAttributes {
	this := CommitmentsSavingsTimeseriesResponseAttributes{}
	return &this
}

// GetActualCost returns the ActualCost field value.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetActualCost() CommitmentsTimeseriesMetric {
	if o == nil {
		var ret CommitmentsTimeseriesMetric
		return ret
	}
	return o.ActualCost
}

// GetActualCostOk returns a tuple with the ActualCost field value
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetActualCostOk() (*CommitmentsTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActualCost, true
}

// SetActualCost sets field value.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) SetActualCost(v CommitmentsTimeseriesMetric) {
	o.ActualCost = v
}

// GetEffectiveSavingsRate returns the EffectiveSavingsRate field value.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetEffectiveSavingsRate() CommitmentsTimeseriesMetric {
	if o == nil {
		var ret CommitmentsTimeseriesMetric
		return ret
	}
	return o.EffectiveSavingsRate
}

// GetEffectiveSavingsRateOk returns a tuple with the EffectiveSavingsRate field value
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetEffectiveSavingsRateOk() (*CommitmentsTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveSavingsRate, true
}

// SetEffectiveSavingsRate sets field value.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) SetEffectiveSavingsRate(v CommitmentsTimeseriesMetric) {
	o.EffectiveSavingsRate = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetMeta() CommitmentsUnusedCostMeta {
	if o == nil || o.Meta == nil {
		var ret CommitmentsUnusedCostMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetMetaOk() (*CommitmentsUnusedCostMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given CommitmentsUnusedCostMeta and assigns it to the Meta field.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) SetMeta(v CommitmentsUnusedCostMeta) {
	o.Meta = &v
}

// GetOnDemandEquivalentCost returns the OnDemandEquivalentCost field value.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetOnDemandEquivalentCost() CommitmentsTimeseriesMetric {
	if o == nil {
		var ret CommitmentsTimeseriesMetric
		return ret
	}
	return o.OnDemandEquivalentCost
}

// GetOnDemandEquivalentCostOk returns a tuple with the OnDemandEquivalentCost field value
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetOnDemandEquivalentCostOk() (*CommitmentsTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnDemandEquivalentCost, true
}

// SetOnDemandEquivalentCost sets field value.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) SetOnDemandEquivalentCost(v CommitmentsTimeseriesMetric) {
	o.OnDemandEquivalentCost = v
}

// GetQueryDefinition returns the QueryDefinition field value if set, zero value otherwise.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetQueryDefinition() CommitmentsQueryDefinition {
	if o == nil || o.QueryDefinition == nil {
		var ret CommitmentsQueryDefinition
		return ret
	}
	return *o.QueryDefinition
}

// GetQueryDefinitionOk returns a tuple with the QueryDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetQueryDefinitionOk() (*CommitmentsQueryDefinition, bool) {
	if o == nil || o.QueryDefinition == nil {
		return nil, false
	}
	return o.QueryDefinition, true
}

// HasQueryDefinition returns a boolean if a field has been set.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) HasQueryDefinition() bool {
	return o != nil && o.QueryDefinition != nil
}

// SetQueryDefinition gets a reference to the given CommitmentsQueryDefinition and assigns it to the QueryDefinition field.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) SetQueryDefinition(v CommitmentsQueryDefinition) {
	o.QueryDefinition = &v
}

// GetRealizedSavings returns the RealizedSavings field value.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetRealizedSavings() CommitmentsTimeseriesMetric {
	if o == nil {
		var ret CommitmentsTimeseriesMetric
		return ret
	}
	return o.RealizedSavings
}

// GetRealizedSavingsOk returns a tuple with the RealizedSavings field value
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) GetRealizedSavingsOk() (*CommitmentsTimeseriesMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RealizedSavings, true
}

// SetRealizedSavings sets field value.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) SetRealizedSavings(v CommitmentsTimeseriesMetric) {
	o.RealizedSavings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsSavingsTimeseriesResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["actual_cost"] = o.ActualCost
	toSerialize["effective_savings_rate"] = o.EffectiveSavingsRate
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["on_demand_equivalent_cost"] = o.OnDemandEquivalentCost
	if o.QueryDefinition != nil {
		toSerialize["query_definition"] = o.QueryDefinition
	}
	toSerialize["realized_savings"] = o.RealizedSavings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsSavingsTimeseriesResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActualCost             *CommitmentsTimeseriesMetric `json:"actual_cost"`
		EffectiveSavingsRate   *CommitmentsTimeseriesMetric `json:"effective_savings_rate"`
		Meta                   *CommitmentsUnusedCostMeta   `json:"meta,omitempty"`
		OnDemandEquivalentCost *CommitmentsTimeseriesMetric `json:"on_demand_equivalent_cost"`
		QueryDefinition        *CommitmentsQueryDefinition  `json:"query_definition,omitempty"`
		RealizedSavings        *CommitmentsTimeseriesMetric `json:"realized_savings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActualCost == nil {
		return fmt.Errorf("required field actual_cost missing")
	}
	if all.EffectiveSavingsRate == nil {
		return fmt.Errorf("required field effective_savings_rate missing")
	}
	if all.OnDemandEquivalentCost == nil {
		return fmt.Errorf("required field on_demand_equivalent_cost missing")
	}
	if all.RealizedSavings == nil {
		return fmt.Errorf("required field realized_savings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actual_cost", "effective_savings_rate", "meta", "on_demand_equivalent_cost", "query_definition", "realized_savings"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ActualCost.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ActualCost = *all.ActualCost
	if all.EffectiveSavingsRate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EffectiveSavingsRate = *all.EffectiveSavingsRate
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	if all.OnDemandEquivalentCost.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OnDemandEquivalentCost = *all.OnDemandEquivalentCost
	if all.QueryDefinition != nil && all.QueryDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.QueryDefinition = all.QueryDefinition
	if all.RealizedSavings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RealizedSavings = *all.RealizedSavings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
