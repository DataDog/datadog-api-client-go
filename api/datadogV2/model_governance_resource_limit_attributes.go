// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceResourceLimitAttributes The attributes of a governance resource limit.
type GovernanceResourceLimitAttributes struct {
	// The current limit configured for the resource.
	CurrentLimit float64 `json:"current_limit"`
	// The current value of the resource.
	CurrentValue float64 `json:"current_value"`
	// A link to the Datadog page where the resource and its limit can be managed.
	DeepLink string `json:"deep_link"`
	// The default current value used when the resource value cannot be computed from the query.
	DefaultCurrentValue float64 `json:"default_current_value"`
	// The default limit value used when the limit cannot be computed from the query.
	DefaultLimitValue float64 `json:"default_limit_value"`
	// A description of what the resource limit measures.
	Description string `json:"description"`
	// The human-readable name of the resource limit.
	DisplayName string `json:"display_name"`
	// A metric query used to compute usage against a limit.
	LimitQuery GovernanceLimitQuery `json:"limit_query"`
	// The Datadog product the resource limit belongs to.
	Product string `json:"product"`
	// A metric query used to compute usage against a limit.
	Query GovernanceLimitQuery `json:"query"`
	// The query execution context used to visualize a limit and its usage.
	QueryConfig GovernanceLimitQueryConfig `json:"query_config"`
	// The time range over which the resource value is measured.
	TimeRange string `json:"time_range"`
	// The unit in which the resource value and limit are measured.
	UnitName string `json:"unit_name"`
	// The current usage status of the resource relative to its limit.
	UsageStatus string `json:"usage_status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceResourceLimitAttributes instantiates a new GovernanceResourceLimitAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceResourceLimitAttributes(currentLimit float64, currentValue float64, deepLink string, defaultCurrentValue float64, defaultLimitValue float64, description string, displayName string, limitQuery GovernanceLimitQuery, product string, query GovernanceLimitQuery, queryConfig GovernanceLimitQueryConfig, timeRange string, unitName string, usageStatus string) *GovernanceResourceLimitAttributes {
	this := GovernanceResourceLimitAttributes{}
	this.CurrentLimit = currentLimit
	this.CurrentValue = currentValue
	this.DeepLink = deepLink
	this.DefaultCurrentValue = defaultCurrentValue
	this.DefaultLimitValue = defaultLimitValue
	this.Description = description
	this.DisplayName = displayName
	this.LimitQuery = limitQuery
	this.Product = product
	this.Query = query
	this.QueryConfig = queryConfig
	this.TimeRange = timeRange
	this.UnitName = unitName
	this.UsageStatus = usageStatus
	return &this
}

// NewGovernanceResourceLimitAttributesWithDefaults instantiates a new GovernanceResourceLimitAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceResourceLimitAttributesWithDefaults() *GovernanceResourceLimitAttributes {
	this := GovernanceResourceLimitAttributes{}
	return &this
}

// GetCurrentLimit returns the CurrentLimit field value.
func (o *GovernanceResourceLimitAttributes) GetCurrentLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.CurrentLimit
}

// GetCurrentLimitOk returns a tuple with the CurrentLimit field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetCurrentLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentLimit, true
}

// SetCurrentLimit sets field value.
func (o *GovernanceResourceLimitAttributes) SetCurrentLimit(v float64) {
	o.CurrentLimit = v
}

// GetCurrentValue returns the CurrentValue field value.
func (o *GovernanceResourceLimitAttributes) GetCurrentValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetCurrentValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentValue, true
}

// SetCurrentValue sets field value.
func (o *GovernanceResourceLimitAttributes) SetCurrentValue(v float64) {
	o.CurrentValue = v
}

// GetDeepLink returns the DeepLink field value.
func (o *GovernanceResourceLimitAttributes) GetDeepLink() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DeepLink
}

// GetDeepLinkOk returns a tuple with the DeepLink field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetDeepLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeepLink, true
}

// SetDeepLink sets field value.
func (o *GovernanceResourceLimitAttributes) SetDeepLink(v string) {
	o.DeepLink = v
}

// GetDefaultCurrentValue returns the DefaultCurrentValue field value.
func (o *GovernanceResourceLimitAttributes) GetDefaultCurrentValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.DefaultCurrentValue
}

// GetDefaultCurrentValueOk returns a tuple with the DefaultCurrentValue field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetDefaultCurrentValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultCurrentValue, true
}

// SetDefaultCurrentValue sets field value.
func (o *GovernanceResourceLimitAttributes) SetDefaultCurrentValue(v float64) {
	o.DefaultCurrentValue = v
}

// GetDefaultLimitValue returns the DefaultLimitValue field value.
func (o *GovernanceResourceLimitAttributes) GetDefaultLimitValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.DefaultLimitValue
}

// GetDefaultLimitValueOk returns a tuple with the DefaultLimitValue field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetDefaultLimitValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultLimitValue, true
}

// SetDefaultLimitValue sets field value.
func (o *GovernanceResourceLimitAttributes) SetDefaultLimitValue(v float64) {
	o.DefaultLimitValue = v
}

// GetDescription returns the Description field value.
func (o *GovernanceResourceLimitAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GovernanceResourceLimitAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value.
func (o *GovernanceResourceLimitAttributes) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *GovernanceResourceLimitAttributes) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetLimitQuery returns the LimitQuery field value.
func (o *GovernanceResourceLimitAttributes) GetLimitQuery() GovernanceLimitQuery {
	if o == nil {
		var ret GovernanceLimitQuery
		return ret
	}
	return o.LimitQuery
}

// GetLimitQueryOk returns a tuple with the LimitQuery field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetLimitQueryOk() (*GovernanceLimitQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitQuery, true
}

// SetLimitQuery sets field value.
func (o *GovernanceResourceLimitAttributes) SetLimitQuery(v GovernanceLimitQuery) {
	o.LimitQuery = v
}

// GetProduct returns the Product field value.
func (o *GovernanceResourceLimitAttributes) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value.
func (o *GovernanceResourceLimitAttributes) SetProduct(v string) {
	o.Product = v
}

// GetQuery returns the Query field value.
func (o *GovernanceResourceLimitAttributes) GetQuery() GovernanceLimitQuery {
	if o == nil {
		var ret GovernanceLimitQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetQueryOk() (*GovernanceLimitQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *GovernanceResourceLimitAttributes) SetQuery(v GovernanceLimitQuery) {
	o.Query = v
}

// GetQueryConfig returns the QueryConfig field value.
func (o *GovernanceResourceLimitAttributes) GetQueryConfig() GovernanceLimitQueryConfig {
	if o == nil {
		var ret GovernanceLimitQueryConfig
		return ret
	}
	return o.QueryConfig
}

// GetQueryConfigOk returns a tuple with the QueryConfig field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetQueryConfigOk() (*GovernanceLimitQueryConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryConfig, true
}

// SetQueryConfig sets field value.
func (o *GovernanceResourceLimitAttributes) SetQueryConfig(v GovernanceLimitQueryConfig) {
	o.QueryConfig = v
}

// GetTimeRange returns the TimeRange field value.
func (o *GovernanceResourceLimitAttributes) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value.
func (o *GovernanceResourceLimitAttributes) SetTimeRange(v string) {
	o.TimeRange = v
}

// GetUnitName returns the UnitName field value.
func (o *GovernanceResourceLimitAttributes) GetUnitName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetUnitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitName, true
}

// SetUnitName sets field value.
func (o *GovernanceResourceLimitAttributes) SetUnitName(v string) {
	o.UnitName = v
}

// GetUsageStatus returns the UsageStatus field value.
func (o *GovernanceResourceLimitAttributes) GetUsageStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UsageStatus
}

// GetUsageStatusOk returns a tuple with the UsageStatus field value
// and a boolean to check if the value has been set.
func (o *GovernanceResourceLimitAttributes) GetUsageStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageStatus, true
}

// SetUsageStatus sets field value.
func (o *GovernanceResourceLimitAttributes) SetUsageStatus(v string) {
	o.UsageStatus = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceResourceLimitAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["current_limit"] = o.CurrentLimit
	toSerialize["current_value"] = o.CurrentValue
	toSerialize["deep_link"] = o.DeepLink
	toSerialize["default_current_value"] = o.DefaultCurrentValue
	toSerialize["default_limit_value"] = o.DefaultLimitValue
	toSerialize["description"] = o.Description
	toSerialize["display_name"] = o.DisplayName
	toSerialize["limit_query"] = o.LimitQuery
	toSerialize["product"] = o.Product
	toSerialize["query"] = o.Query
	toSerialize["query_config"] = o.QueryConfig
	toSerialize["time_range"] = o.TimeRange
	toSerialize["unit_name"] = o.UnitName
	toSerialize["usage_status"] = o.UsageStatus

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceResourceLimitAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CurrentLimit        *float64                    `json:"current_limit"`
		CurrentValue        *float64                    `json:"current_value"`
		DeepLink            *string                     `json:"deep_link"`
		DefaultCurrentValue *float64                    `json:"default_current_value"`
		DefaultLimitValue   *float64                    `json:"default_limit_value"`
		Description         *string                     `json:"description"`
		DisplayName         *string                     `json:"display_name"`
		LimitQuery          *GovernanceLimitQuery       `json:"limit_query"`
		Product             *string                     `json:"product"`
		Query               *GovernanceLimitQuery       `json:"query"`
		QueryConfig         *GovernanceLimitQueryConfig `json:"query_config"`
		TimeRange           *string                     `json:"time_range"`
		UnitName            *string                     `json:"unit_name"`
		UsageStatus         *string                     `json:"usage_status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CurrentLimit == nil {
		return fmt.Errorf("required field current_limit missing")
	}
	if all.CurrentValue == nil {
		return fmt.Errorf("required field current_value missing")
	}
	if all.DeepLink == nil {
		return fmt.Errorf("required field deep_link missing")
	}
	if all.DefaultCurrentValue == nil {
		return fmt.Errorf("required field default_current_value missing")
	}
	if all.DefaultLimitValue == nil {
		return fmt.Errorf("required field default_limit_value missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field display_name missing")
	}
	if all.LimitQuery == nil {
		return fmt.Errorf("required field limit_query missing")
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
	if all.UnitName == nil {
		return fmt.Errorf("required field unit_name missing")
	}
	if all.UsageStatus == nil {
		return fmt.Errorf("required field usage_status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"current_limit", "current_value", "deep_link", "default_current_value", "default_limit_value", "description", "display_name", "limit_query", "product", "query", "query_config", "time_range", "unit_name", "usage_status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CurrentLimit = *all.CurrentLimit
	o.CurrentValue = *all.CurrentValue
	o.DeepLink = *all.DeepLink
	o.DefaultCurrentValue = *all.DefaultCurrentValue
	o.DefaultLimitValue = *all.DefaultLimitValue
	o.Description = *all.Description
	o.DisplayName = *all.DisplayName
	if all.LimitQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LimitQuery = *all.LimitQuery
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
