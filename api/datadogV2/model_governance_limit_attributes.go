// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceLimitAttributes The attributes of a governance limit.
type GovernanceLimitAttributes struct {
	// A description of what the limit measures.
	Description string `json:"description"`
	// The human-readable name of the limit.
	DisplayName string `json:"display_name"`
	// The type of limit, such as a rate limit or a usage limit.
	LimitType string `json:"limit_type"`
	// The Datadog product the limit belongs to.
	Product string `json:"product"`
	// A metric query used to compute usage against a limit.
	Query GovernanceLimitQuery `json:"query"`
	// The query execution context used to visualize a limit and its usage.
	QueryConfig GovernanceLimitQueryConfig `json:"query_config"`
	// The time range over which usage against the limit is measured.
	TimeRange string `json:"time_range"`
	// The number of times usage has reached the limit within the measured time range.
	TimesHitLimit float64 `json:"times_hit_limit"`
	// The unit in which the limit and its usage are measured.
	UnitName string `json:"unit_name"`
	// The current usage status of the limit relative to its threshold.
	UsageStatus string `json:"usage_status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceLimitAttributes instantiates a new GovernanceLimitAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceLimitAttributes(description string, displayName string, limitType string, product string, query GovernanceLimitQuery, queryConfig GovernanceLimitQueryConfig, timeRange string, timesHitLimit float64, unitName string, usageStatus string) *GovernanceLimitAttributes {
	this := GovernanceLimitAttributes{}
	this.Description = description
	this.DisplayName = displayName
	this.LimitType = limitType
	this.Product = product
	this.Query = query
	this.QueryConfig = queryConfig
	this.TimeRange = timeRange
	this.TimesHitLimit = timesHitLimit
	this.UnitName = unitName
	this.UsageStatus = usageStatus
	return &this
}

// NewGovernanceLimitAttributesWithDefaults instantiates a new GovernanceLimitAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceLimitAttributesWithDefaults() *GovernanceLimitAttributes {
	this := GovernanceLimitAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *GovernanceLimitAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GovernanceLimitAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value.
func (o *GovernanceLimitAttributes) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *GovernanceLimitAttributes) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetLimitType returns the LimitType field value.
func (o *GovernanceLimitAttributes) GetLimitType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LimitType
}

// GetLimitTypeOk returns a tuple with the LimitType field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitAttributes) GetLimitTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitType, true
}

// SetLimitType sets field value.
func (o *GovernanceLimitAttributes) SetLimitType(v string) {
	o.LimitType = v
}

// GetProduct returns the Product field value.
func (o *GovernanceLimitAttributes) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitAttributes) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value.
func (o *GovernanceLimitAttributes) SetProduct(v string) {
	o.Product = v
}

// GetQuery returns the Query field value.
func (o *GovernanceLimitAttributes) GetQuery() GovernanceLimitQuery {
	if o == nil {
		var ret GovernanceLimitQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitAttributes) GetQueryOk() (*GovernanceLimitQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *GovernanceLimitAttributes) SetQuery(v GovernanceLimitQuery) {
	o.Query = v
}

// GetQueryConfig returns the QueryConfig field value.
func (o *GovernanceLimitAttributes) GetQueryConfig() GovernanceLimitQueryConfig {
	if o == nil {
		var ret GovernanceLimitQueryConfig
		return ret
	}
	return o.QueryConfig
}

// GetQueryConfigOk returns a tuple with the QueryConfig field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitAttributes) GetQueryConfigOk() (*GovernanceLimitQueryConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryConfig, true
}

// SetQueryConfig sets field value.
func (o *GovernanceLimitAttributes) SetQueryConfig(v GovernanceLimitQueryConfig) {
	o.QueryConfig = v
}

// GetTimeRange returns the TimeRange field value.
func (o *GovernanceLimitAttributes) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitAttributes) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value.
func (o *GovernanceLimitAttributes) SetTimeRange(v string) {
	o.TimeRange = v
}

// GetTimesHitLimit returns the TimesHitLimit field value.
func (o *GovernanceLimitAttributes) GetTimesHitLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.TimesHitLimit
}

// GetTimesHitLimitOk returns a tuple with the TimesHitLimit field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitAttributes) GetTimesHitLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimesHitLimit, true
}

// SetTimesHitLimit sets field value.
func (o *GovernanceLimitAttributes) SetTimesHitLimit(v float64) {
	o.TimesHitLimit = v
}

// GetUnitName returns the UnitName field value.
func (o *GovernanceLimitAttributes) GetUnitName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitAttributes) GetUnitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitName, true
}

// SetUnitName sets field value.
func (o *GovernanceLimitAttributes) SetUnitName(v string) {
	o.UnitName = v
}

// GetUsageStatus returns the UsageStatus field value.
func (o *GovernanceLimitAttributes) GetUsageStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UsageStatus
}

// GetUsageStatusOk returns a tuple with the UsageStatus field value
// and a boolean to check if the value has been set.
func (o *GovernanceLimitAttributes) GetUsageStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageStatus, true
}

// SetUsageStatus sets field value.
func (o *GovernanceLimitAttributes) SetUsageStatus(v string) {
	o.UsageStatus = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceLimitAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["display_name"] = o.DisplayName
	toSerialize["limit_type"] = o.LimitType
	toSerialize["product"] = o.Product
	toSerialize["query"] = o.Query
	toSerialize["query_config"] = o.QueryConfig
	toSerialize["time_range"] = o.TimeRange
	toSerialize["times_hit_limit"] = o.TimesHitLimit
	toSerialize["unit_name"] = o.UnitName
	toSerialize["usage_status"] = o.UsageStatus

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceLimitAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description   *string                     `json:"description"`
		DisplayName   *string                     `json:"display_name"`
		LimitType     *string                     `json:"limit_type"`
		Product       *string                     `json:"product"`
		Query         *GovernanceLimitQuery       `json:"query"`
		QueryConfig   *GovernanceLimitQueryConfig `json:"query_config"`
		TimeRange     *string                     `json:"time_range"`
		TimesHitLimit *float64                    `json:"times_hit_limit"`
		UnitName      *string                     `json:"unit_name"`
		UsageStatus   *string                     `json:"usage_status"`
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
	if all.LimitType == nil {
		return fmt.Errorf("required field limit_type missing")
	}
	if all.Product == nil {
		return fmt.Errorf("required field product missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.QueryConfig == nil {
		return fmt.Errorf("required field query_config missing")
	}
	if all.TimeRange == nil {
		return fmt.Errorf("required field time_range missing")
	}
	if all.TimesHitLimit == nil {
		return fmt.Errorf("required field times_hit_limit missing")
	}
	if all.UnitName == nil {
		return fmt.Errorf("required field unit_name missing")
	}
	if all.UsageStatus == nil {
		return fmt.Errorf("required field usage_status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "display_name", "limit_type", "product", "query", "query_config", "time_range", "times_hit_limit", "unit_name", "usage_status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = *all.Description
	o.DisplayName = *all.DisplayName
	o.LimitType = *all.LimitType
	o.Product = *all.Product
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
	if all.QueryConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.QueryConfig = *all.QueryConfig
	o.TimeRange = *all.TimeRange
	o.TimesHitLimit = *all.TimesHitLimit
	o.UnitName = *all.UnitName
	o.UsageStatus = *all.UsageStatus

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
