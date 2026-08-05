// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceInsightAttributes The attributes of a governance insight. Exactly one of `metric_query`, `event_query`,
// `usage_query`, `audit_query`, or `percentage_query` is populated, depending on the data
// source the insight is computed from; the rest are `null`.
type GovernanceInsightAttributes struct {
	// An audit log query used to compute an insight value.
	AuditQuery *GovernanceInsightAuditQuery `json:"audit_query,omitempty"`
	// A human-readable description of what the insight measures.
	Description string `json:"description"`
	// Human-readable name of the insight.
	DisplayName string `json:"display_name"`
	// An event query used to compute an insight value.
	EventQuery *GovernanceInsightEventQuery `json:"event_query,omitempty"`
	// A metric query used to compute an insight value.
	MetricQuery *GovernanceInsightMetricQuery `json:"metric_query,omitempty"`
	// A percentage query that computes an insight value as a ratio of two metric queries.
	PercentageQuery *GovernanceInsightPercentageQuery `json:"percentage_query,omitempty"`
	// The product the insight belongs to.
	Product string `json:"product"`
	// Query execution context for running insight queries directly.
	QueryConfig *GovernanceInsightQueryConfig `json:"query_config,omitempty"`
	// The sub-product the insight belongs to, if any.
	SubProduct string `json:"sub_product"`
	// The time range the insight value is computed over, if applicable.
	TimeRange string `json:"time_range"`
	// The unit that the insight's value is measured in.
	UnitName string `json:"unit_name"`
	// A usage query used to compute an insight value.
	UsageQuery *GovernanceInsightUsageQuery `json:"usage_query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceInsightAttributes instantiates a new GovernanceInsightAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceInsightAttributes(description string, displayName string, product string, subProduct string, timeRange string, unitName string) *GovernanceInsightAttributes {
	this := GovernanceInsightAttributes{}
	this.Description = description
	this.DisplayName = displayName
	this.Product = product
	this.SubProduct = subProduct
	this.TimeRange = timeRange
	this.UnitName = unitName
	return &this
}

// NewGovernanceInsightAttributesWithDefaults instantiates a new GovernanceInsightAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceInsightAttributesWithDefaults() *GovernanceInsightAttributes {
	this := GovernanceInsightAttributes{}
	return &this
}

// GetAuditQuery returns the AuditQuery field value if set, zero value otherwise.
func (o *GovernanceInsightAttributes) GetAuditQuery() GovernanceInsightAuditQuery {
	if o == nil || o.AuditQuery == nil {
		var ret GovernanceInsightAuditQuery
		return ret
	}
	return *o.AuditQuery
}

// GetAuditQueryOk returns a tuple with the AuditQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetAuditQueryOk() (*GovernanceInsightAuditQuery, bool) {
	if o == nil || o.AuditQuery == nil {
		return nil, false
	}
	return o.AuditQuery, true
}

// HasAuditQuery returns a boolean if a field has been set.
func (o *GovernanceInsightAttributes) HasAuditQuery() bool {
	return o != nil && o.AuditQuery != nil
}

// SetAuditQuery gets a reference to the given GovernanceInsightAuditQuery and assigns it to the AuditQuery field.
func (o *GovernanceInsightAttributes) SetAuditQuery(v GovernanceInsightAuditQuery) {
	o.AuditQuery = &v
}

// GetDescription returns the Description field value.
func (o *GovernanceInsightAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GovernanceInsightAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value.
func (o *GovernanceInsightAttributes) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *GovernanceInsightAttributes) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEventQuery returns the EventQuery field value if set, zero value otherwise.
func (o *GovernanceInsightAttributes) GetEventQuery() GovernanceInsightEventQuery {
	if o == nil || o.EventQuery == nil {
		var ret GovernanceInsightEventQuery
		return ret
	}
	return *o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetEventQueryOk() (*GovernanceInsightEventQuery, bool) {
	if o == nil || o.EventQuery == nil {
		return nil, false
	}
	return o.EventQuery, true
}

// HasEventQuery returns a boolean if a field has been set.
func (o *GovernanceInsightAttributes) HasEventQuery() bool {
	return o != nil && o.EventQuery != nil
}

// SetEventQuery gets a reference to the given GovernanceInsightEventQuery and assigns it to the EventQuery field.
func (o *GovernanceInsightAttributes) SetEventQuery(v GovernanceInsightEventQuery) {
	o.EventQuery = &v
}

// GetMetricQuery returns the MetricQuery field value if set, zero value otherwise.
func (o *GovernanceInsightAttributes) GetMetricQuery() GovernanceInsightMetricQuery {
	if o == nil || o.MetricQuery == nil {
		var ret GovernanceInsightMetricQuery
		return ret
	}
	return *o.MetricQuery
}

// GetMetricQueryOk returns a tuple with the MetricQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetMetricQueryOk() (*GovernanceInsightMetricQuery, bool) {
	if o == nil || o.MetricQuery == nil {
		return nil, false
	}
	return o.MetricQuery, true
}

// HasMetricQuery returns a boolean if a field has been set.
func (o *GovernanceInsightAttributes) HasMetricQuery() bool {
	return o != nil && o.MetricQuery != nil
}

// SetMetricQuery gets a reference to the given GovernanceInsightMetricQuery and assigns it to the MetricQuery field.
func (o *GovernanceInsightAttributes) SetMetricQuery(v GovernanceInsightMetricQuery) {
	o.MetricQuery = &v
}

// GetPercentageQuery returns the PercentageQuery field value if set, zero value otherwise.
func (o *GovernanceInsightAttributes) GetPercentageQuery() GovernanceInsightPercentageQuery {
	if o == nil || o.PercentageQuery == nil {
		var ret GovernanceInsightPercentageQuery
		return ret
	}
	return *o.PercentageQuery
}

// GetPercentageQueryOk returns a tuple with the PercentageQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetPercentageQueryOk() (*GovernanceInsightPercentageQuery, bool) {
	if o == nil || o.PercentageQuery == nil {
		return nil, false
	}
	return o.PercentageQuery, true
}

// HasPercentageQuery returns a boolean if a field has been set.
func (o *GovernanceInsightAttributes) HasPercentageQuery() bool {
	return o != nil && o.PercentageQuery != nil
}

// SetPercentageQuery gets a reference to the given GovernanceInsightPercentageQuery and assigns it to the PercentageQuery field.
func (o *GovernanceInsightAttributes) SetPercentageQuery(v GovernanceInsightPercentageQuery) {
	o.PercentageQuery = &v
}

// GetProduct returns the Product field value.
func (o *GovernanceInsightAttributes) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value.
func (o *GovernanceInsightAttributes) SetProduct(v string) {
	o.Product = v
}

// GetQueryConfig returns the QueryConfig field value if set, zero value otherwise.
func (o *GovernanceInsightAttributes) GetQueryConfig() GovernanceInsightQueryConfig {
	if o == nil || o.QueryConfig == nil {
		var ret GovernanceInsightQueryConfig
		return ret
	}
	return *o.QueryConfig
}

// GetQueryConfigOk returns a tuple with the QueryConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetQueryConfigOk() (*GovernanceInsightQueryConfig, bool) {
	if o == nil || o.QueryConfig == nil {
		return nil, false
	}
	return o.QueryConfig, true
}

// HasQueryConfig returns a boolean if a field has been set.
func (o *GovernanceInsightAttributes) HasQueryConfig() bool {
	return o != nil && o.QueryConfig != nil
}

// SetQueryConfig gets a reference to the given GovernanceInsightQueryConfig and assigns it to the QueryConfig field.
func (o *GovernanceInsightAttributes) SetQueryConfig(v GovernanceInsightQueryConfig) {
	o.QueryConfig = &v
}

// GetSubProduct returns the SubProduct field value.
func (o *GovernanceInsightAttributes) GetSubProduct() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SubProduct
}

// GetSubProductOk returns a tuple with the SubProduct field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetSubProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubProduct, true
}

// SetSubProduct sets field value.
func (o *GovernanceInsightAttributes) SetSubProduct(v string) {
	o.SubProduct = v
}

// GetTimeRange returns the TimeRange field value.
func (o *GovernanceInsightAttributes) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value.
func (o *GovernanceInsightAttributes) SetTimeRange(v string) {
	o.TimeRange = v
}

// GetUnitName returns the UnitName field value.
func (o *GovernanceInsightAttributes) GetUnitName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetUnitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitName, true
}

// SetUnitName sets field value.
func (o *GovernanceInsightAttributes) SetUnitName(v string) {
	o.UnitName = v
}

// GetUsageQuery returns the UsageQuery field value if set, zero value otherwise.
func (o *GovernanceInsightAttributes) GetUsageQuery() GovernanceInsightUsageQuery {
	if o == nil || o.UsageQuery == nil {
		var ret GovernanceInsightUsageQuery
		return ret
	}
	return *o.UsageQuery
}

// GetUsageQueryOk returns a tuple with the UsageQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceInsightAttributes) GetUsageQueryOk() (*GovernanceInsightUsageQuery, bool) {
	if o == nil || o.UsageQuery == nil {
		return nil, false
	}
	return o.UsageQuery, true
}

// HasUsageQuery returns a boolean if a field has been set.
func (o *GovernanceInsightAttributes) HasUsageQuery() bool {
	return o != nil && o.UsageQuery != nil
}

// SetUsageQuery gets a reference to the given GovernanceInsightUsageQuery and assigns it to the UsageQuery field.
func (o *GovernanceInsightAttributes) SetUsageQuery(v GovernanceInsightUsageQuery) {
	o.UsageQuery = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceInsightAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuditQuery != nil {
		toSerialize["audit_query"] = o.AuditQuery
	}
	toSerialize["description"] = o.Description
	toSerialize["display_name"] = o.DisplayName
	if o.EventQuery != nil {
		toSerialize["event_query"] = o.EventQuery
	}
	if o.MetricQuery != nil {
		toSerialize["metric_query"] = o.MetricQuery
	}
	if o.PercentageQuery != nil {
		toSerialize["percentage_query"] = o.PercentageQuery
	}
	toSerialize["product"] = o.Product
	if o.QueryConfig != nil {
		toSerialize["query_config"] = o.QueryConfig
	}
	toSerialize["sub_product"] = o.SubProduct
	toSerialize["time_range"] = o.TimeRange
	toSerialize["unit_name"] = o.UnitName
	if o.UsageQuery != nil {
		toSerialize["usage_query"] = o.UsageQuery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceInsightAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuditQuery      *GovernanceInsightAuditQuery      `json:"audit_query,omitempty"`
		Description     *string                           `json:"description"`
		DisplayName     *string                           `json:"display_name"`
		EventQuery      *GovernanceInsightEventQuery      `json:"event_query,omitempty"`
		MetricQuery     *GovernanceInsightMetricQuery     `json:"metric_query,omitempty"`
		PercentageQuery *GovernanceInsightPercentageQuery `json:"percentage_query,omitempty"`
		Product         *string                           `json:"product"`
		QueryConfig     *GovernanceInsightQueryConfig     `json:"query_config,omitempty"`
		SubProduct      *string                           `json:"sub_product"`
		TimeRange       *string                           `json:"time_range"`
		UnitName        *string                           `json:"unit_name"`
		UsageQuery      *GovernanceInsightUsageQuery      `json:"usage_query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field display_name missing")
	}
	if all.Product == nil {
		return fmt.Errorf("required field product missing")
	}
	if all.SubProduct == nil {
		return fmt.Errorf("required field sub_product missing")
	}
	if all.TimeRange == nil {
		return fmt.Errorf("required field time_range missing")
	}
	if all.UnitName == nil {
		return fmt.Errorf("required field unit_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audit_query", "description", "display_name", "event_query", "metric_query", "percentage_query", "product", "query_config", "sub_product", "time_range", "unit_name", "usage_query"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AuditQuery != nil && all.AuditQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AuditQuery = all.AuditQuery
	o.Description = *all.Description
	o.DisplayName = *all.DisplayName
	if all.EventQuery != nil && all.EventQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EventQuery = all.EventQuery
	if all.MetricQuery != nil && all.MetricQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricQuery = all.MetricQuery
	if all.PercentageQuery != nil && all.PercentageQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PercentageQuery = all.PercentageQuery
	o.Product = *all.Product
	if all.QueryConfig != nil && all.QueryConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.QueryConfig = all.QueryConfig
	o.SubProduct = *all.SubProduct
	o.TimeRange = *all.TimeRange
	o.UnitName = *all.UnitName
	if all.UsageQuery != nil && all.UsageQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UsageQuery = all.UsageQuery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
