// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorAttributes Attributes of a Product Analytics monitor.
type ProductAnalyticsMonitorAttributes struct {
	// Time the monitor was created.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// The user who created the monitor.
	Creator *ProductAnalyticsMonitorCreator `json:"creator,omitempty"`
	// Notification message sent when the monitor alerts.
	Message string `json:"message"`
	// Time the monitor was last modified.
	ModifiedTime *time.Time `json:"modified_time,omitempty"`
	// Name of the monitor.
	Name string `json:"name"`
	// Options of a Product Analytics monitor.
	Options ProductAnalyticsMonitorOptions `json:"options"`
	// ID of the organization that owns the monitor.
	OrgId int64 `json:"org_id"`
	// Priority of the monitor, from 1 (highest) to 5 (lowest).
	Priority *int32 `json:"priority,omitempty"`
	// Monitor query.
	Query string `json:"query"`
	// Current alert status of the monitor. `No Data` is returned when the monitor has not been evaluated yet.
	Status ProductAnalyticsMonitorStatus `json:"status"`
	// Monitor tags, as `key:value` strings.
	Tags datadog.NullableList[string] `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsMonitorAttributes instantiates a new ProductAnalyticsMonitorAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsMonitorAttributes(message string, name string, options ProductAnalyticsMonitorOptions, orgId int64, query string, status ProductAnalyticsMonitorStatus) *ProductAnalyticsMonitorAttributes {
	this := ProductAnalyticsMonitorAttributes{}
	this.Message = message
	this.Name = name
	this.Options = options
	this.OrgId = orgId
	this.Query = query
	this.Status = status
	return &this
}

// NewProductAnalyticsMonitorAttributesWithDefaults instantiates a new ProductAnalyticsMonitorAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsMonitorAttributesWithDefaults() *ProductAnalyticsMonitorAttributes {
	this := ProductAnalyticsMonitorAttributes{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorAttributes) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorAttributes) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorAttributes) HasCreatedTime() bool {
	return o != nil && o.CreatedTime != nil
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *ProductAnalyticsMonitorAttributes) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorAttributes) GetCreator() ProductAnalyticsMonitorCreator {
	if o == nil || o.Creator == nil {
		var ret ProductAnalyticsMonitorCreator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorAttributes) GetCreatorOk() (*ProductAnalyticsMonitorCreator, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorAttributes) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given ProductAnalyticsMonitorCreator and assigns it to the Creator field.
func (o *ProductAnalyticsMonitorAttributes) SetCreator(v ProductAnalyticsMonitorCreator) {
	o.Creator = &v
}

// GetMessage returns the Message field value.
func (o *ProductAnalyticsMonitorAttributes) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorAttributes) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *ProductAnalyticsMonitorAttributes) SetMessage(v string) {
	o.Message = v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorAttributes) GetModifiedTime() time.Time {
	if o == nil || o.ModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorAttributes) GetModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.ModifiedTime == nil {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorAttributes) HasModifiedTime() bool {
	return o != nil && o.ModifiedTime != nil
}

// SetModifiedTime gets a reference to the given time.Time and assigns it to the ModifiedTime field.
func (o *ProductAnalyticsMonitorAttributes) SetModifiedTime(v time.Time) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value.
func (o *ProductAnalyticsMonitorAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProductAnalyticsMonitorAttributes) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value.
func (o *ProductAnalyticsMonitorAttributes) GetOptions() ProductAnalyticsMonitorOptions {
	if o == nil {
		var ret ProductAnalyticsMonitorOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorAttributes) GetOptionsOk() (*ProductAnalyticsMonitorOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *ProductAnalyticsMonitorAttributes) SetOptions(v ProductAnalyticsMonitorOptions) {
	o.Options = v
}

// GetOrgId returns the OrgId field value.
func (o *ProductAnalyticsMonitorAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *ProductAnalyticsMonitorAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorAttributes) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorAttributes) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ProductAnalyticsMonitorAttributes) SetPriority(v int32) {
	o.Priority = &v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsMonitorAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsMonitorAttributes) SetQuery(v string) {
	o.Query = v
}

// GetStatus returns the Status field value.
func (o *ProductAnalyticsMonitorAttributes) GetStatus() ProductAnalyticsMonitorStatus {
	if o == nil {
		var ret ProductAnalyticsMonitorStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorAttributes) GetStatusOk() (*ProductAnalyticsMonitorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ProductAnalyticsMonitorAttributes) SetStatus(v ProductAnalyticsMonitorStatus) {
	o.Status = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsMonitorAttributes) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsMonitorAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorAttributes) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *ProductAnalyticsMonitorAttributes) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *ProductAnalyticsMonitorAttributes) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *ProductAnalyticsMonitorAttributes) UnsetTags() {
	o.Tags.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsMonitorAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedTime != nil {
		if o.CreatedTime.Nanosecond() == 0 {
			toSerialize["created_time"] = o.CreatedTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_time"] = o.CreatedTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	toSerialize["message"] = o.Message
	if o.ModifiedTime != nil {
		if o.ModifiedTime.Nanosecond() == 0 {
			toSerialize["modified_time"] = o.ModifiedTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_time"] = o.ModifiedTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	toSerialize["org_id"] = o.OrgId
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["query"] = o.Query
	toSerialize["status"] = o.Status
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsMonitorAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedTime  *time.Time                      `json:"created_time,omitempty"`
		Creator      *ProductAnalyticsMonitorCreator `json:"creator,omitempty"`
		Message      *string                         `json:"message"`
		ModifiedTime *time.Time                      `json:"modified_time,omitempty"`
		Name         *string                         `json:"name"`
		Options      *ProductAnalyticsMonitorOptions `json:"options"`
		OrgId        *int64                          `json:"org_id"`
		Priority     *int32                          `json:"priority,omitempty"`
		Query        *string                         `json:"query"`
		Status       *ProductAnalyticsMonitorStatus  `json:"status"`
		Tags         datadog.NullableList[string]    `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Options == nil {
		return fmt.Errorf("required field options missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_time", "creator", "message", "modified_time", "name", "options", "org_id", "priority", "query", "status", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedTime = all.CreatedTime
	if all.Creator != nil && all.Creator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Creator = all.Creator
	o.Message = *all.Message
	o.ModifiedTime = all.ModifiedTime
	o.Name = *all.Name
	if all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = *all.Options
	o.OrgId = *all.OrgId
	o.Priority = all.Priority
	o.Query = *all.Query
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
