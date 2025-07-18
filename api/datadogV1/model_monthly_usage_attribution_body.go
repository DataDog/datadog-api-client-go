// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyUsageAttributionBody Usage Summary by tag for a given organization.
type MonthlyUsageAttributionBody struct {
	// Datetime in ISO-8601 format, UTC, precise to month: [YYYY-MM].
	Month *time.Time `json:"month,omitempty"`
	// The name of the organization.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// The region of the Datadog instance that the organization belongs to.
	Region *string `json:"region,omitempty"`
	// The source of the usage attribution tag configuration and the selected tags in the format `<source_org_name>:::<selected tag 1>///<selected tag 2>///<selected tag 3>`.
	TagConfigSource *string `json:"tag_config_source,omitempty"`
	// Tag keys and values.
	//
	// A `null` value here means that the requested tag breakdown cannot be applied because it does not match the [tags
	// configured for usage attribution](https://docs.datadoghq.com/account_management/billing/usage_attribution/#getting-started).
	// In this scenario the API returns the total usage, not broken down by tags.
	Tags map[string][]string `json:"tags,omitempty"`
	// Datetime of the most recent update to the usage values.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Fields in Usage Summary by tag(s).
	// The following values have been **deprecated**: `estimated_indexed_spans_usage`, `estimated_indexed_spans_percentage`, `estimated_ingested_spans_usage`, `estimated_ingested_spans_percentage`.
	Values *MonthlyUsageAttributionValues `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonthlyUsageAttributionBody instantiates a new MonthlyUsageAttributionBody object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonthlyUsageAttributionBody() *MonthlyUsageAttributionBody {
	this := MonthlyUsageAttributionBody{}
	return &this
}

// NewMonthlyUsageAttributionBodyWithDefaults instantiates a new MonthlyUsageAttributionBody object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonthlyUsageAttributionBodyWithDefaults() *MonthlyUsageAttributionBody {
	this := MonthlyUsageAttributionBody{}
	return &this
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionBody) GetMonth() time.Time {
	if o == nil || o.Month == nil {
		var ret time.Time
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionBody) GetMonthOk() (*time.Time, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionBody) HasMonth() bool {
	return o != nil && o.Month != nil
}

// SetMonth gets a reference to the given time.Time and assigns it to the Month field.
func (o *MonthlyUsageAttributionBody) SetMonth(v time.Time) {
	o.Month = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionBody) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionBody) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionBody) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *MonthlyUsageAttributionBody) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionBody) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionBody) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionBody) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *MonthlyUsageAttributionBody) SetPublicId(v string) {
	o.PublicId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionBody) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionBody) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionBody) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *MonthlyUsageAttributionBody) SetRegion(v string) {
	o.Region = &v
}

// GetTagConfigSource returns the TagConfigSource field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionBody) GetTagConfigSource() string {
	if o == nil || o.TagConfigSource == nil {
		var ret string
		return ret
	}
	return *o.TagConfigSource
}

// GetTagConfigSourceOk returns a tuple with the TagConfigSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionBody) GetTagConfigSourceOk() (*string, bool) {
	if o == nil || o.TagConfigSource == nil {
		return nil, false
	}
	return o.TagConfigSource, true
}

// HasTagConfigSource returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionBody) HasTagConfigSource() bool {
	return o != nil && o.TagConfigSource != nil
}

// SetTagConfigSource gets a reference to the given string and assigns it to the TagConfigSource field.
func (o *MonthlyUsageAttributionBody) SetTagConfigSource(v string) {
	o.TagConfigSource = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyUsageAttributionBody) GetTags() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyUsageAttributionBody) GetTagsOk() (*map[string][]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionBody) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string][]string and assigns it to the Tags field.
func (o *MonthlyUsageAttributionBody) SetTags(v map[string][]string) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionBody) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionBody) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionBody) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *MonthlyUsageAttributionBody) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionBody) GetValues() MonthlyUsageAttributionValues {
	if o == nil || o.Values == nil {
		var ret MonthlyUsageAttributionValues
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionBody) GetValuesOk() (*MonthlyUsageAttributionValues, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionBody) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given MonthlyUsageAttributionValues and assigns it to the Values field.
func (o *MonthlyUsageAttributionBody) SetValues(v MonthlyUsageAttributionValues) {
	o.Values = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonthlyUsageAttributionBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Month != nil {
		if o.Month.Nanosecond() == 0 {
			toSerialize["month"] = o.Month.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["month"] = o.Month.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.TagConfigSource != nil {
		toSerialize["tag_config_source"] = o.TagConfigSource
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonthlyUsageAttributionBody) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Month           *time.Time                     `json:"month,omitempty"`
		OrgName         *string                        `json:"org_name,omitempty"`
		PublicId        *string                        `json:"public_id,omitempty"`
		Region          *string                        `json:"region,omitempty"`
		TagConfigSource *string                        `json:"tag_config_source,omitempty"`
		Tags            map[string][]string            `json:"tags,omitempty"`
		UpdatedAt       *time.Time                     `json:"updated_at,omitempty"`
		Values          *MonthlyUsageAttributionValues `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"month", "org_name", "public_id", "region", "tag_config_source", "tags", "updated_at", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Month = all.Month
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	o.Region = all.Region
	o.TagConfigSource = all.TagConfigSource
	o.Tags = all.Tags
	o.UpdatedAt = all.UpdatedAt
	if all.Values != nil && all.Values.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
