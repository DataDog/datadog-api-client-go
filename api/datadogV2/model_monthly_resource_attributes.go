// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyResourceAttributes Resource usage and organization details for the requested month.
type MonthlyResourceAttributes struct {
	// Date of the source usage data. May be empty when unavailable.
	DataDate *string `json:"data_date,omitempty"`
	// Name of the organization that owns the resource.
	OrgName *string `json:"org_name,omitempty"`
	// Public ID of the organization that owns the resource.
	PublicId *string `json:"public_id,omitempty"`
	// Name of the resource, such as a host name, metric name, or log index.
	ResourceName *string `json:"resource_name,omitempty"`
	// Resource tags when available. Empty for resource types without tag data.
	Tags map[string]string `json:"tags,omitempty"`
	// Selected billing hour for percentile usage, or the first day of the month for monthly usage. May be empty when unavailable.
	Timestamp *string `json:"timestamp,omitempty"`
	// Usage type of the resource. Values depend on the billing dimension and are not limited to host types.
	UsageType *string `json:"usage_type,omitempty"`
	// Usage attributed to the resource in the units of the requested billing dimension. Values can be fractional.
	UsageValue datadog.NullableFloat64 `json:"usage_value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonthlyResourceAttributes instantiates a new MonthlyResourceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonthlyResourceAttributes() *MonthlyResourceAttributes {
	this := MonthlyResourceAttributes{}
	return &this
}

// NewMonthlyResourceAttributesWithDefaults instantiates a new MonthlyResourceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonthlyResourceAttributesWithDefaults() *MonthlyResourceAttributes {
	this := MonthlyResourceAttributes{}
	return &this
}

// GetDataDate returns the DataDate field value if set, zero value otherwise.
func (o *MonthlyResourceAttributes) GetDataDate() string {
	if o == nil || o.DataDate == nil {
		var ret string
		return ret
	}
	return *o.DataDate
}

// GetDataDateOk returns a tuple with the DataDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyResourceAttributes) GetDataDateOk() (*string, bool) {
	if o == nil || o.DataDate == nil {
		return nil, false
	}
	return o.DataDate, true
}

// HasDataDate returns a boolean if a field has been set.
func (o *MonthlyResourceAttributes) HasDataDate() bool {
	return o != nil && o.DataDate != nil
}

// SetDataDate gets a reference to the given string and assigns it to the DataDate field.
func (o *MonthlyResourceAttributes) SetDataDate(v string) {
	o.DataDate = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *MonthlyResourceAttributes) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyResourceAttributes) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *MonthlyResourceAttributes) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *MonthlyResourceAttributes) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *MonthlyResourceAttributes) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyResourceAttributes) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *MonthlyResourceAttributes) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *MonthlyResourceAttributes) SetPublicId(v string) {
	o.PublicId = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *MonthlyResourceAttributes) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyResourceAttributes) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *MonthlyResourceAttributes) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *MonthlyResourceAttributes) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MonthlyResourceAttributes) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyResourceAttributes) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MonthlyResourceAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *MonthlyResourceAttributes) SetTags(v map[string]string) {
	o.Tags = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MonthlyResourceAttributes) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyResourceAttributes) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MonthlyResourceAttributes) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *MonthlyResourceAttributes) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetUsageType returns the UsageType field value if set, zero value otherwise.
func (o *MonthlyResourceAttributes) GetUsageType() string {
	if o == nil || o.UsageType == nil {
		var ret string
		return ret
	}
	return *o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyResourceAttributes) GetUsageTypeOk() (*string, bool) {
	if o == nil || o.UsageType == nil {
		return nil, false
	}
	return o.UsageType, true
}

// HasUsageType returns a boolean if a field has been set.
func (o *MonthlyResourceAttributes) HasUsageType() bool {
	return o != nil && o.UsageType != nil
}

// SetUsageType gets a reference to the given string and assigns it to the UsageType field.
func (o *MonthlyResourceAttributes) SetUsageType(v string) {
	o.UsageType = &v
}

// GetUsageValue returns the UsageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyResourceAttributes) GetUsageValue() float64 {
	if o == nil || o.UsageValue.Get() == nil {
		var ret float64
		return ret
	}
	return *o.UsageValue.Get()
}

// GetUsageValueOk returns a tuple with the UsageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyResourceAttributes) GetUsageValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageValue.Get(), o.UsageValue.IsSet()
}

// HasUsageValue returns a boolean if a field has been set.
func (o *MonthlyResourceAttributes) HasUsageValue() bool {
	return o != nil && o.UsageValue.IsSet()
}

// SetUsageValue gets a reference to the given datadog.NullableFloat64 and assigns it to the UsageValue field.
func (o *MonthlyResourceAttributes) SetUsageValue(v float64) {
	o.UsageValue.Set(&v)
}

// SetUsageValueNil sets the value for UsageValue to be an explicit nil.
func (o *MonthlyResourceAttributes) SetUsageValueNil() {
	o.UsageValue.Set(nil)
}

// UnsetUsageValue ensures that no value is present for UsageValue, not even an explicit nil.
func (o *MonthlyResourceAttributes) UnsetUsageValue() {
	o.UsageValue.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o MonthlyResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DataDate != nil {
		toSerialize["data_date"] = o.DataDate
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.UsageType != nil {
		toSerialize["usage_type"] = o.UsageType
	}
	if o.UsageValue.IsSet() {
		toSerialize["usage_value"] = o.UsageValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonthlyResourceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataDate     *string                 `json:"data_date,omitempty"`
		OrgName      *string                 `json:"org_name,omitempty"`
		PublicId     *string                 `json:"public_id,omitempty"`
		ResourceName *string                 `json:"resource_name,omitempty"`
		Tags         map[string]string       `json:"tags,omitempty"`
		Timestamp    *string                 `json:"timestamp,omitempty"`
		UsageType    *string                 `json:"usage_type,omitempty"`
		UsageValue   datadog.NullableFloat64 `json:"usage_value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_date", "org_name", "public_id", "resource_name", "tags", "timestamp", "usage_type", "usage_value"})
	} else {
		return err
	}
	o.DataDate = all.DataDate
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	o.ResourceName = all.ResourceName
	o.Tags = all.Tags
	o.Timestamp = all.Timestamp
	o.UsageType = all.UsageType
	o.UsageValue = all.UsageValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
