// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FullAPIKeyAttributes Attributes of a full API key.
type FullAPIKeyAttributes struct {
	// The category of the API key.
	Category *string `json:"category,omitempty"`
	// Creation date of the API key.
	CreatedAt *string `json:"created_at,omitempty"`
	// The date and time the API key was last used.
	DateLastUsed datadog.NullableString `json:"date_last_used,omitempty"`
	// The API key.
	Key *string `json:"key,omitempty"`
	// The last four characters of the API key.
	Last4 *string `json:"last4,omitempty"`
	// Attributes for the last time the specific API key was used.
	LastUsedDate *FullAPIKeyLastUsedDate `json:"last_used_date,omitempty"`
	// Date the API key was last modified.
	ModifiedAt *string `json:"modified_at,omitempty"`
	// Name of the API key.
	Name *string `json:"name,omitempty"`
	// The remote config read enabled status.
	RemoteConfigReadEnabled *bool `json:"remote_config_read_enabled,omitempty"`
	// If the API key was used within the last 24 hours.
	UsedInLast24Hours *bool `json:"used_in_last_24_hours,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewFullAPIKeyAttributes instantiates a new FullAPIKeyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFullAPIKeyAttributes() *FullAPIKeyAttributes {
	this := FullAPIKeyAttributes{}
	return &this
}

// NewFullAPIKeyAttributesWithDefaults instantiates a new FullAPIKeyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFullAPIKeyAttributesWithDefaults() *FullAPIKeyAttributes {
	this := FullAPIKeyAttributes{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *FullAPIKeyAttributes) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKeyAttributes) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *FullAPIKeyAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *FullAPIKeyAttributes) SetCategory(v string) {
	o.Category = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FullAPIKeyAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKeyAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FullAPIKeyAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *FullAPIKeyAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDateLastUsed returns the DateLastUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FullAPIKeyAttributes) GetDateLastUsed() string {
	if o == nil || o.DateLastUsed.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateLastUsed.Get()
}

// GetDateLastUsedOk returns a tuple with the DateLastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FullAPIKeyAttributes) GetDateLastUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateLastUsed.Get(), o.DateLastUsed.IsSet()
}

// HasDateLastUsed returns a boolean if a field has been set.
func (o *FullAPIKeyAttributes) HasDateLastUsed() bool {
	return o != nil && o.DateLastUsed.IsSet()
}

// SetDateLastUsed gets a reference to the given datadog.NullableString and assigns it to the DateLastUsed field.
func (o *FullAPIKeyAttributes) SetDateLastUsed(v string) {
	o.DateLastUsed.Set(&v)
}

// SetDateLastUsedNil sets the value for DateLastUsed to be an explicit nil.
func (o *FullAPIKeyAttributes) SetDateLastUsedNil() {
	o.DateLastUsed.Set(nil)
}

// UnsetDateLastUsed ensures that no value is present for DateLastUsed, not even an explicit nil.
func (o *FullAPIKeyAttributes) UnsetDateLastUsed() {
	o.DateLastUsed.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FullAPIKeyAttributes) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKeyAttributes) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FullAPIKeyAttributes) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FullAPIKeyAttributes) SetKey(v string) {
	o.Key = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *FullAPIKeyAttributes) GetLast4() string {
	if o == nil || o.Last4 == nil {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKeyAttributes) GetLast4Ok() (*string, bool) {
	if o == nil || o.Last4 == nil {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *FullAPIKeyAttributes) HasLast4() bool {
	return o != nil && o.Last4 != nil
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *FullAPIKeyAttributes) SetLast4(v string) {
	o.Last4 = &v
}

// GetLastUsedDate returns the LastUsedDate field value if set, zero value otherwise.
func (o *FullAPIKeyAttributes) GetLastUsedDate() FullAPIKeyLastUsedDate {
	if o == nil || o.LastUsedDate == nil {
		var ret FullAPIKeyLastUsedDate
		return ret
	}
	return *o.LastUsedDate
}

// GetLastUsedDateOk returns a tuple with the LastUsedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKeyAttributes) GetLastUsedDateOk() (*FullAPIKeyLastUsedDate, bool) {
	if o == nil || o.LastUsedDate == nil {
		return nil, false
	}
	return o.LastUsedDate, true
}

// HasLastUsedDate returns a boolean if a field has been set.
func (o *FullAPIKeyAttributes) HasLastUsedDate() bool {
	return o != nil && o.LastUsedDate != nil
}

// SetLastUsedDate gets a reference to the given FullAPIKeyLastUsedDate and assigns it to the LastUsedDate field.
func (o *FullAPIKeyAttributes) SetLastUsedDate(v FullAPIKeyLastUsedDate) {
	o.LastUsedDate = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *FullAPIKeyAttributes) GetModifiedAt() string {
	if o == nil || o.ModifiedAt == nil {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKeyAttributes) GetModifiedAtOk() (*string, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *FullAPIKeyAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *FullAPIKeyAttributes) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FullAPIKeyAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKeyAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FullAPIKeyAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FullAPIKeyAttributes) SetName(v string) {
	o.Name = &v
}

// GetRemoteConfigReadEnabled returns the RemoteConfigReadEnabled field value if set, zero value otherwise.
func (o *FullAPIKeyAttributes) GetRemoteConfigReadEnabled() bool {
	if o == nil || o.RemoteConfigReadEnabled == nil {
		var ret bool
		return ret
	}
	return *o.RemoteConfigReadEnabled
}

// GetRemoteConfigReadEnabledOk returns a tuple with the RemoteConfigReadEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKeyAttributes) GetRemoteConfigReadEnabledOk() (*bool, bool) {
	if o == nil || o.RemoteConfigReadEnabled == nil {
		return nil, false
	}
	return o.RemoteConfigReadEnabled, true
}

// HasRemoteConfigReadEnabled returns a boolean if a field has been set.
func (o *FullAPIKeyAttributes) HasRemoteConfigReadEnabled() bool {
	return o != nil && o.RemoteConfigReadEnabled != nil
}

// SetRemoteConfigReadEnabled gets a reference to the given bool and assigns it to the RemoteConfigReadEnabled field.
func (o *FullAPIKeyAttributes) SetRemoteConfigReadEnabled(v bool) {
	o.RemoteConfigReadEnabled = &v
}

// GetUsedInLast24Hours returns the UsedInLast24Hours field value if set, zero value otherwise.
func (o *FullAPIKeyAttributes) GetUsedInLast24Hours() bool {
	if o == nil || o.UsedInLast24Hours == nil {
		var ret bool
		return ret
	}
	return *o.UsedInLast24Hours
}

// GetUsedInLast24HoursOk returns a tuple with the UsedInLast24Hours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullAPIKeyAttributes) GetUsedInLast24HoursOk() (*bool, bool) {
	if o == nil || o.UsedInLast24Hours == nil {
		return nil, false
	}
	return o.UsedInLast24Hours, true
}

// HasUsedInLast24Hours returns a boolean if a field has been set.
func (o *FullAPIKeyAttributes) HasUsedInLast24Hours() bool {
	return o != nil && o.UsedInLast24Hours != nil
}

// SetUsedInLast24Hours gets a reference to the given bool and assigns it to the UsedInLast24Hours field.
func (o *FullAPIKeyAttributes) SetUsedInLast24Hours(v bool) {
	o.UsedInLast24Hours = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FullAPIKeyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DateLastUsed.IsSet() {
		toSerialize["date_last_used"] = o.DateLastUsed.Get()
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Last4 != nil {
		toSerialize["last4"] = o.Last4
	}
	if o.LastUsedDate != nil {
		toSerialize["last_used_date"] = o.LastUsedDate
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RemoteConfigReadEnabled != nil {
		toSerialize["remote_config_read_enabled"] = o.RemoteConfigReadEnabled
	}
	if o.UsedInLast24Hours != nil {
		toSerialize["used_in_last_24_hours"] = o.UsedInLast24Hours
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FullAPIKeyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category                *string                 `json:"category,omitempty"`
		CreatedAt               *string                 `json:"created_at,omitempty"`
		DateLastUsed            datadog.NullableString  `json:"date_last_used,omitempty"`
		Key                     *string                 `json:"key,omitempty"`
		Last4                   *string                 `json:"last4,omitempty"`
		LastUsedDate            *FullAPIKeyLastUsedDate `json:"last_used_date,omitempty"`
		ModifiedAt              *string                 `json:"modified_at,omitempty"`
		Name                    *string                 `json:"name,omitempty"`
		RemoteConfigReadEnabled *bool                   `json:"remote_config_read_enabled,omitempty"`
		UsedInLast24Hours       *bool                   `json:"used_in_last_24_hours,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "created_at", "date_last_used", "key", "last4", "last_used_date", "modified_at", "name", "remote_config_read_enabled", "used_in_last_24_hours"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Category = all.Category
	o.CreatedAt = all.CreatedAt
	o.DateLastUsed = all.DateLastUsed
	o.Key = all.Key
	o.Last4 = all.Last4
	if all.LastUsedDate != nil && all.LastUsedDate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastUsedDate = all.LastUsedDate
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.RemoteConfigReadEnabled = all.RemoteConfigReadEnabled
	o.UsedInLast24Hours = all.UsedInLast24Hours

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
