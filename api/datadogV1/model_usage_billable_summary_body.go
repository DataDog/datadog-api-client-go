// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageBillableSummaryBody Response with properties for each aggregated usage type.
type UsageBillableSummaryBody struct {
	// The total account usage.
	AccountBillableUsage datadog.NullableInt64 `json:"account_billable_usage,omitempty"`
	// Elapsed usage hours for some billable product.
	ElapsedUsageHours datadog.NullableInt64 `json:"elapsed_usage_hours,omitempty"`
	// The first billable hour for the org.
	FirstBillableUsageHour *time.Time `json:"first_billable_usage_hour,omitempty"`
	// The last billable hour for the org.
	LastBillableUsageHour *time.Time `json:"last_billable_usage_hour,omitempty"`
	// The number of units used within the billable timeframe.
	OrgBillableUsage datadog.NullableInt64 `json:"org_billable_usage,omitempty"`
	// The percentage of account usage the org represents.
	PercentageInAccount datadog.NullableFloat64 `json:"percentage_in_account,omitempty"`
	// Units pertaining to the usage.
	UsageUnit *string `json:"usage_unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageBillableSummaryBody instantiates a new UsageBillableSummaryBody object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageBillableSummaryBody() *UsageBillableSummaryBody {
	this := UsageBillableSummaryBody{}
	return &this
}

// NewUsageBillableSummaryBodyWithDefaults instantiates a new UsageBillableSummaryBody object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageBillableSummaryBodyWithDefaults() *UsageBillableSummaryBody {
	this := UsageBillableSummaryBody{}
	return &this
}

// GetAccountBillableUsage returns the AccountBillableUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageBillableSummaryBody) GetAccountBillableUsage() int64 {
	if o == nil || o.AccountBillableUsage.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AccountBillableUsage.Get()
}

// GetAccountBillableUsageOk returns a tuple with the AccountBillableUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageBillableSummaryBody) GetAccountBillableUsageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountBillableUsage.Get(), o.AccountBillableUsage.IsSet()
}

// HasAccountBillableUsage returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasAccountBillableUsage() bool {
	return o != nil && o.AccountBillableUsage.IsSet()
}

// SetAccountBillableUsage gets a reference to the given datadog.NullableInt64 and assigns it to the AccountBillableUsage field.
func (o *UsageBillableSummaryBody) SetAccountBillableUsage(v int64) {
	o.AccountBillableUsage.Set(&v)
}

// SetAccountBillableUsageNil sets the value for AccountBillableUsage to be an explicit nil.
func (o *UsageBillableSummaryBody) SetAccountBillableUsageNil() {
	o.AccountBillableUsage.Set(nil)
}

// UnsetAccountBillableUsage ensures that no value is present for AccountBillableUsage, not even an explicit nil.
func (o *UsageBillableSummaryBody) UnsetAccountBillableUsage() {
	o.AccountBillableUsage.Unset()
}

// GetElapsedUsageHours returns the ElapsedUsageHours field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageBillableSummaryBody) GetElapsedUsageHours() int64 {
	if o == nil || o.ElapsedUsageHours.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ElapsedUsageHours.Get()
}

// GetElapsedUsageHoursOk returns a tuple with the ElapsedUsageHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageBillableSummaryBody) GetElapsedUsageHoursOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ElapsedUsageHours.Get(), o.ElapsedUsageHours.IsSet()
}

// HasElapsedUsageHours returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasElapsedUsageHours() bool {
	return o != nil && o.ElapsedUsageHours.IsSet()
}

// SetElapsedUsageHours gets a reference to the given datadog.NullableInt64 and assigns it to the ElapsedUsageHours field.
func (o *UsageBillableSummaryBody) SetElapsedUsageHours(v int64) {
	o.ElapsedUsageHours.Set(&v)
}

// SetElapsedUsageHoursNil sets the value for ElapsedUsageHours to be an explicit nil.
func (o *UsageBillableSummaryBody) SetElapsedUsageHoursNil() {
	o.ElapsedUsageHours.Set(nil)
}

// UnsetElapsedUsageHours ensures that no value is present for ElapsedUsageHours, not even an explicit nil.
func (o *UsageBillableSummaryBody) UnsetElapsedUsageHours() {
	o.ElapsedUsageHours.Unset()
}

// GetFirstBillableUsageHour returns the FirstBillableUsageHour field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetFirstBillableUsageHour() time.Time {
	if o == nil || o.FirstBillableUsageHour == nil {
		var ret time.Time
		return ret
	}
	return *o.FirstBillableUsageHour
}

// GetFirstBillableUsageHourOk returns a tuple with the FirstBillableUsageHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetFirstBillableUsageHourOk() (*time.Time, bool) {
	if o == nil || o.FirstBillableUsageHour == nil {
		return nil, false
	}
	return o.FirstBillableUsageHour, true
}

// HasFirstBillableUsageHour returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasFirstBillableUsageHour() bool {
	return o != nil && o.FirstBillableUsageHour != nil
}

// SetFirstBillableUsageHour gets a reference to the given time.Time and assigns it to the FirstBillableUsageHour field.
func (o *UsageBillableSummaryBody) SetFirstBillableUsageHour(v time.Time) {
	o.FirstBillableUsageHour = &v
}

// GetLastBillableUsageHour returns the LastBillableUsageHour field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetLastBillableUsageHour() time.Time {
	if o == nil || o.LastBillableUsageHour == nil {
		var ret time.Time
		return ret
	}
	return *o.LastBillableUsageHour
}

// GetLastBillableUsageHourOk returns a tuple with the LastBillableUsageHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetLastBillableUsageHourOk() (*time.Time, bool) {
	if o == nil || o.LastBillableUsageHour == nil {
		return nil, false
	}
	return o.LastBillableUsageHour, true
}

// HasLastBillableUsageHour returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasLastBillableUsageHour() bool {
	return o != nil && o.LastBillableUsageHour != nil
}

// SetLastBillableUsageHour gets a reference to the given time.Time and assigns it to the LastBillableUsageHour field.
func (o *UsageBillableSummaryBody) SetLastBillableUsageHour(v time.Time) {
	o.LastBillableUsageHour = &v
}

// GetOrgBillableUsage returns the OrgBillableUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageBillableSummaryBody) GetOrgBillableUsage() int64 {
	if o == nil || o.OrgBillableUsage.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OrgBillableUsage.Get()
}

// GetOrgBillableUsageOk returns a tuple with the OrgBillableUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageBillableSummaryBody) GetOrgBillableUsageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgBillableUsage.Get(), o.OrgBillableUsage.IsSet()
}

// HasOrgBillableUsage returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasOrgBillableUsage() bool {
	return o != nil && o.OrgBillableUsage.IsSet()
}

// SetOrgBillableUsage gets a reference to the given datadog.NullableInt64 and assigns it to the OrgBillableUsage field.
func (o *UsageBillableSummaryBody) SetOrgBillableUsage(v int64) {
	o.OrgBillableUsage.Set(&v)
}

// SetOrgBillableUsageNil sets the value for OrgBillableUsage to be an explicit nil.
func (o *UsageBillableSummaryBody) SetOrgBillableUsageNil() {
	o.OrgBillableUsage.Set(nil)
}

// UnsetOrgBillableUsage ensures that no value is present for OrgBillableUsage, not even an explicit nil.
func (o *UsageBillableSummaryBody) UnsetOrgBillableUsage() {
	o.OrgBillableUsage.Unset()
}

// GetPercentageInAccount returns the PercentageInAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageBillableSummaryBody) GetPercentageInAccount() float64 {
	if o == nil || o.PercentageInAccount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.PercentageInAccount.Get()
}

// GetPercentageInAccountOk returns a tuple with the PercentageInAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageBillableSummaryBody) GetPercentageInAccountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentageInAccount.Get(), o.PercentageInAccount.IsSet()
}

// HasPercentageInAccount returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasPercentageInAccount() bool {
	return o != nil && o.PercentageInAccount.IsSet()
}

// SetPercentageInAccount gets a reference to the given datadog.NullableFloat64 and assigns it to the PercentageInAccount field.
func (o *UsageBillableSummaryBody) SetPercentageInAccount(v float64) {
	o.PercentageInAccount.Set(&v)
}

// SetPercentageInAccountNil sets the value for PercentageInAccount to be an explicit nil.
func (o *UsageBillableSummaryBody) SetPercentageInAccountNil() {
	o.PercentageInAccount.Set(nil)
}

// UnsetPercentageInAccount ensures that no value is present for PercentageInAccount, not even an explicit nil.
func (o *UsageBillableSummaryBody) UnsetPercentageInAccount() {
	o.PercentageInAccount.Unset()
}

// GetUsageUnit returns the UsageUnit field value if set, zero value otherwise.
func (o *UsageBillableSummaryBody) GetUsageUnit() string {
	if o == nil || o.UsageUnit == nil {
		var ret string
		return ret
	}
	return *o.UsageUnit
}

// GetUsageUnitOk returns a tuple with the UsageUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageBillableSummaryBody) GetUsageUnitOk() (*string, bool) {
	if o == nil || o.UsageUnit == nil {
		return nil, false
	}
	return o.UsageUnit, true
}

// HasUsageUnit returns a boolean if a field has been set.
func (o *UsageBillableSummaryBody) HasUsageUnit() bool {
	return o != nil && o.UsageUnit != nil
}

// SetUsageUnit gets a reference to the given string and assigns it to the UsageUnit field.
func (o *UsageBillableSummaryBody) SetUsageUnit(v string) {
	o.UsageUnit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageBillableSummaryBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AccountBillableUsage.IsSet() {
		toSerialize["account_billable_usage"] = o.AccountBillableUsage.Get()
	}
	if o.ElapsedUsageHours.IsSet() {
		toSerialize["elapsed_usage_hours"] = o.ElapsedUsageHours.Get()
	}
	if o.FirstBillableUsageHour != nil {
		if o.FirstBillableUsageHour.Nanosecond() == 0 {
			toSerialize["first_billable_usage_hour"] = o.FirstBillableUsageHour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["first_billable_usage_hour"] = o.FirstBillableUsageHour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LastBillableUsageHour != nil {
		if o.LastBillableUsageHour.Nanosecond() == 0 {
			toSerialize["last_billable_usage_hour"] = o.LastBillableUsageHour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_billable_usage_hour"] = o.LastBillableUsageHour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.OrgBillableUsage.IsSet() {
		toSerialize["org_billable_usage"] = o.OrgBillableUsage.Get()
	}
	if o.PercentageInAccount.IsSet() {
		toSerialize["percentage_in_account"] = o.PercentageInAccount.Get()
	}
	if o.UsageUnit != nil {
		toSerialize["usage_unit"] = o.UsageUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageBillableSummaryBody) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AccountBillableUsage   datadog.NullableInt64   `json:"account_billable_usage,omitempty"`
		ElapsedUsageHours      datadog.NullableInt64   `json:"elapsed_usage_hours,omitempty"`
		FirstBillableUsageHour *time.Time              `json:"first_billable_usage_hour,omitempty"`
		LastBillableUsageHour  *time.Time              `json:"last_billable_usage_hour,omitempty"`
		OrgBillableUsage       datadog.NullableInt64   `json:"org_billable_usage,omitempty"`
		PercentageInAccount    datadog.NullableFloat64 `json:"percentage_in_account,omitempty"`
		UsageUnit              *string                 `json:"usage_unit,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_billable_usage", "elapsed_usage_hours", "first_billable_usage_hour", "last_billable_usage_hour", "org_billable_usage", "percentage_in_account", "usage_unit"})
	} else {
		return err
	}
	o.AccountBillableUsage = all.AccountBillableUsage
	o.ElapsedUsageHours = all.ElapsedUsageHours
	o.FirstBillableUsageHour = all.FirstBillableUsageHour
	o.LastBillableUsageHour = all.LastBillableUsageHour
	o.OrgBillableUsage = all.OrgBillableUsage
	o.PercentageInAccount = all.PercentageInAccount
	o.UsageUnit = all.UsageUnit
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
