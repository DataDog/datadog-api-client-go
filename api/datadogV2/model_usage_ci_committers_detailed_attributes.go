// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageCICommittersDetailedAttributes The response containing CI Committers Detailed attributes for specified organization.
type UsageCICommittersDetailedAttributes struct {
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// The region of the organization.
	Region *string `json:"region,omitempty"`
	// Shows the hour of the usage.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The usage type associated with the commit.
	UsageType *string `json:"usage_type,omitempty"`
	// The user email of CI committer.
	UserEmail *string `json:"user_email,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageCICommittersDetailedAttributes instantiates a new UsageCICommittersDetailedAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageCICommittersDetailedAttributes() *UsageCICommittersDetailedAttributes {
	this := UsageCICommittersDetailedAttributes{}
	return &this
}

// NewUsageCICommittersDetailedAttributesWithDefaults instantiates a new UsageCICommittersDetailedAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageCICommittersDetailedAttributesWithDefaults() *UsageCICommittersDetailedAttributes {
	this := UsageCICommittersDetailedAttributes{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageCICommittersDetailedAttributes) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCICommittersDetailedAttributes) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageCICommittersDetailedAttributes) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageCICommittersDetailedAttributes) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageCICommittersDetailedAttributes) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCICommittersDetailedAttributes) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageCICommittersDetailedAttributes) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageCICommittersDetailedAttributes) SetPublicId(v string) {
	o.PublicId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *UsageCICommittersDetailedAttributes) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCICommittersDetailedAttributes) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *UsageCICommittersDetailedAttributes) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *UsageCICommittersDetailedAttributes) SetRegion(v string) {
	o.Region = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *UsageCICommittersDetailedAttributes) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCICommittersDetailedAttributes) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *UsageCICommittersDetailedAttributes) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *UsageCICommittersDetailedAttributes) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetUsageType returns the UsageType field value if set, zero value otherwise.
func (o *UsageCICommittersDetailedAttributes) GetUsageType() string {
	if o == nil || o.UsageType == nil {
		var ret string
		return ret
	}
	return *o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCICommittersDetailedAttributes) GetUsageTypeOk() (*string, bool) {
	if o == nil || o.UsageType == nil {
		return nil, false
	}
	return o.UsageType, true
}

// HasUsageType returns a boolean if a field has been set.
func (o *UsageCICommittersDetailedAttributes) HasUsageType() bool {
	return o != nil && o.UsageType != nil
}

// SetUsageType gets a reference to the given string and assigns it to the UsageType field.
func (o *UsageCICommittersDetailedAttributes) SetUsageType(v string) {
	o.UsageType = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *UsageCICommittersDetailedAttributes) GetUserEmail() string {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCICommittersDetailedAttributes) GetUserEmailOk() (*string, bool) {
	if o == nil || o.UserEmail == nil {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *UsageCICommittersDetailedAttributes) HasUserEmail() bool {
	return o != nil && o.UserEmail != nil
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *UsageCICommittersDetailedAttributes) SetUserEmail(v string) {
	o.UserEmail = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageCICommittersDetailedAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
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
	if o.Timestamp != nil {
		if o.Timestamp.Nanosecond() == 0 {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UsageType != nil {
		toSerialize["usage_type"] = o.UsageType
	}
	if o.UserEmail != nil {
		toSerialize["user_email"] = o.UserEmail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageCICommittersDetailedAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName   *string    `json:"org_name,omitempty"`
		PublicId  *string    `json:"public_id,omitempty"`
		Region    *string    `json:"region,omitempty"`
		Timestamp *time.Time `json:"timestamp,omitempty"`
		UsageType *string    `json:"usage_type,omitempty"`
		UserEmail *string    `json:"user_email,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"org_name", "public_id", "region", "timestamp", "usage_type", "user_email"})
	} else {
		return err
	}
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	o.Region = all.Region
	o.Timestamp = all.Timestamp
	o.UsageType = all.UsageType
	o.UserEmail = all.UserEmail

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
