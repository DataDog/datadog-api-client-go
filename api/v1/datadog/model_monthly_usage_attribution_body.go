/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */


package datadog

import (
	"encoding/json"
	"fmt"

)


// MonthlyUsageAttributionBody Usage Summary by tag for a given organization.
type MonthlyUsageAttributionBody struct {
	// Datetime in ISO-8601 format, UTC, precise to month: [YYYY-MM].
	Month *time.Time `json:"month,omitempty"`
	// The name of the organization.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// The source of the usage attribution tag configuration and the selected tags in the format `<source_org_name>:<selected tag 1>///<selected tag 2>///<selected tag 3>`.
	TagConfigSource *string `json:"tag_config_source,omitempty"`
	// Usage Summary by tag name.
	Tags map[string][]string `json:"tags,omitempty"`
	// Datetime of the most recent update to the usage values.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Fields in Usage Summary by tag(s).
	Values *MonthlyUsageAttributionValues `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}



// NewMonthlyUsageAttributionBody instantiates a new MonthlyUsageAttributionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonthlyUsageAttributionBody() *MonthlyUsageAttributionBody {
	this := MonthlyUsageAttributionBody{}
	return &this
}

// NewMonthlyUsageAttributionBodyWithDefaults instantiates a new MonthlyUsageAttributionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
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
	if o != nil && o.Month != nil {
		return true
	}

	return false
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
	if o != nil && o.OrgName != nil {
		return true
	}

	return false
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
	if o != nil && o.PublicId != nil {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *MonthlyUsageAttributionBody) SetPublicId(v string) {
	o.PublicId = &v
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
	if o != nil && o.TagConfigSource != nil {
		return true
	}

	return false
}

// SetTagConfigSource gets a reference to the given string and assigns it to the TagConfigSource field.
func (o *MonthlyUsageAttributionBody) SetTagConfigSource(v string) {
	o.TagConfigSource = &v
}


// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionBody) GetTags() map[string][]string {
	if o == nil || o.Tags == nil {
		var ret map[string][]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionBody) GetTagsOk() (*map[string][]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionBody) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
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
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
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
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given MonthlyUsageAttributionValues and assigns it to the Values field.
func (o *MonthlyUsageAttributionBody) SetValues(v MonthlyUsageAttributionValues) {
	o.Values = &v
}



func (o MonthlyUsageAttributionBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.TagConfigSource != nil {
		toSerialize["tag_config_source"] = o.TagConfigSource
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}


func (o *MonthlyUsageAttributionBody) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Month *time.Time `json:"month,omitempty"`
		OrgName *string `json:"org_name,omitempty"`
		PublicId *string `json:"public_id,omitempty"`
		TagConfigSource *string `json:"tag_config_source,omitempty"`
		Tags map[string][]string `json:"tags,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
		Values *MonthlyUsageAttributionValues `json:"values,omitempty"`
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
	o.Month = all.Month
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	o.TagConfigSource = all.TagConfigSource
	o.Tags = all.Tags
	o.UpdatedAt = all.UpdatedAt
	o.Values = all.Values
	return nil
}
