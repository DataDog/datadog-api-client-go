// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

// DashboardSummaryDefinition Dashboard definition.
type DashboardSummaryDefinition struct {
	// Object describing the creator of the shared element.
	Author *datadogV1.Creator `json:"author,omitempty"`
	// Date of creation of the dashboard.
	Created datadog.NullableTime `json:"created,omitempty"`
	// URL to the icon of the dashboard.
	Icon datadog.NullableString `json:"icon,omitempty"`
	// ID of the dashboard.
	Id *DashboardSummaryID `json:"id,omitempty"`
	// The short name of the integration.
	IntegrationId datadog.NullableString `json:"integration_id,omitempty"`
	// Whether the dashboard is in the favorites.
	IsFavorite *bool `json:"is_favorite,omitempty"`
	// Whether the dashboard is read only.
	IsReadOnly *bool `json:"is_read_only,omitempty"`
	// Whether the dashboard is publicly shared.
	IsShared *bool `json:"is_shared,omitempty"`
	// Date when the dashboard was last viewed.
	LastViewDate datadog.NullableString `json:"last_view_date,omitempty"`
	// Date of last edition of the dashboard.
	Modified datadog.NullableTime `json:"modified,omitempty"`
	// Popularity of the dashboard.
	Popularity *int32 `json:"popularity,omitempty"`
	// List of team names representing ownership of the dashboard.
	Tags datadog.NullableList[string] `json:"tags,omitempty"`
	// Title of the dashboard.
	Title *string `json:"title,omitempty"`
	// The type of the dashboard.
	Type *string `json:"type,omitempty"`
	// URL path to the dashboard.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardSummaryDefinition instantiates a new DashboardSummaryDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardSummaryDefinition() *DashboardSummaryDefinition {
	this := DashboardSummaryDefinition{}
	return &this
}

// NewDashboardSummaryDefinitionWithDefaults instantiates a new DashboardSummaryDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardSummaryDefinitionWithDefaults() *DashboardSummaryDefinition {
	this := DashboardSummaryDefinition{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *DashboardSummaryDefinition) GetAuthor() datadogV1.Creator {
	if o == nil || o.Author == nil {
		var ret datadogV1.Creator
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDefinition) GetAuthorOk() (*datadogV1.Creator, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given datadogV1.Creator and assigns it to the Author field.
func (o *DashboardSummaryDefinition) SetAuthor(v datadogV1.Creator) {
	o.Author = &v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardSummaryDefinition) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSummaryDefinition) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasCreated() bool {
	return o != nil && o.Created.IsSet()
}

// SetCreated gets a reference to the given datadog.NullableTime and assigns it to the Created field.
func (o *DashboardSummaryDefinition) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// SetCreatedNil sets the value for Created to be an explicit nil.
func (o *DashboardSummaryDefinition) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil.
func (o *DashboardSummaryDefinition) UnsetCreated() {
	o.Created.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardSummaryDefinition) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSummaryDefinition) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasIcon() bool {
	return o != nil && o.Icon.IsSet()
}

// SetIcon gets a reference to the given datadog.NullableString and assigns it to the Icon field.
func (o *DashboardSummaryDefinition) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil.
func (o *DashboardSummaryDefinition) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil.
func (o *DashboardSummaryDefinition) UnsetIcon() {
	o.Icon.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DashboardSummaryDefinition) GetId() DashboardSummaryID {
	if o == nil || o.Id == nil {
		var ret DashboardSummaryID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDefinition) GetIdOk() (*DashboardSummaryID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given DashboardSummaryID and assigns it to the Id field.
func (o *DashboardSummaryDefinition) SetId(v DashboardSummaryID) {
	o.Id = &v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardSummaryDefinition) GetIntegrationId() string {
	if o == nil || o.IntegrationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.IntegrationId.Get()
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSummaryDefinition) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntegrationId.Get(), o.IntegrationId.IsSet()
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasIntegrationId() bool {
	return o != nil && o.IntegrationId.IsSet()
}

// SetIntegrationId gets a reference to the given datadog.NullableString and assigns it to the IntegrationId field.
func (o *DashboardSummaryDefinition) SetIntegrationId(v string) {
	o.IntegrationId.Set(&v)
}

// SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil.
func (o *DashboardSummaryDefinition) SetIntegrationIdNil() {
	o.IntegrationId.Set(nil)
}

// UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil.
func (o *DashboardSummaryDefinition) UnsetIntegrationId() {
	o.IntegrationId.Unset()
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *DashboardSummaryDefinition) GetIsFavorite() bool {
	if o == nil || o.IsFavorite == nil {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDefinition) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || o.IsFavorite == nil {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasIsFavorite() bool {
	return o != nil && o.IsFavorite != nil
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *DashboardSummaryDefinition) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *DashboardSummaryDefinition) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDefinition) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasIsReadOnly() bool {
	return o != nil && o.IsReadOnly != nil
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *DashboardSummaryDefinition) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *DashboardSummaryDefinition) GetIsShared() bool {
	if o == nil || o.IsShared == nil {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDefinition) GetIsSharedOk() (*bool, bool) {
	if o == nil || o.IsShared == nil {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasIsShared() bool {
	return o != nil && o.IsShared != nil
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *DashboardSummaryDefinition) SetIsShared(v bool) {
	o.IsShared = &v
}

// GetLastViewDate returns the LastViewDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardSummaryDefinition) GetLastViewDate() string {
	if o == nil || o.LastViewDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastViewDate.Get()
}

// GetLastViewDateOk returns a tuple with the LastViewDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSummaryDefinition) GetLastViewDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastViewDate.Get(), o.LastViewDate.IsSet()
}

// HasLastViewDate returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasLastViewDate() bool {
	return o != nil && o.LastViewDate.IsSet()
}

// SetLastViewDate gets a reference to the given datadog.NullableString and assigns it to the LastViewDate field.
func (o *DashboardSummaryDefinition) SetLastViewDate(v string) {
	o.LastViewDate.Set(&v)
}

// SetLastViewDateNil sets the value for LastViewDate to be an explicit nil.
func (o *DashboardSummaryDefinition) SetLastViewDateNil() {
	o.LastViewDate.Set(nil)
}

// UnsetLastViewDate ensures that no value is present for LastViewDate, not even an explicit nil.
func (o *DashboardSummaryDefinition) UnsetLastViewDate() {
	o.LastViewDate.Unset()
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardSummaryDefinition) GetModified() time.Time {
	if o == nil || o.Modified.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSummaryDefinition) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasModified() bool {
	return o != nil && o.Modified.IsSet()
}

// SetModified gets a reference to the given datadog.NullableTime and assigns it to the Modified field.
func (o *DashboardSummaryDefinition) SetModified(v time.Time) {
	o.Modified.Set(&v)
}

// SetModifiedNil sets the value for Modified to be an explicit nil.
func (o *DashboardSummaryDefinition) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil.
func (o *DashboardSummaryDefinition) UnsetModified() {
	o.Modified.Unset()
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *DashboardSummaryDefinition) GetPopularity() int32 {
	if o == nil || o.Popularity == nil {
		var ret int32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDefinition) GetPopularityOk() (*int32, bool) {
	if o == nil || o.Popularity == nil {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasPopularity() bool {
	return o != nil && o.Popularity != nil
}

// SetPopularity gets a reference to the given int32 and assigns it to the Popularity field.
func (o *DashboardSummaryDefinition) SetPopularity(v int32) {
	o.Popularity = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DashboardSummaryDefinition) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSummaryDefinition) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *DashboardSummaryDefinition) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *DashboardSummaryDefinition) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *DashboardSummaryDefinition) UnsetTags() {
	o.Tags.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DashboardSummaryDefinition) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDefinition) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DashboardSummaryDefinition) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DashboardSummaryDefinition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDefinition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DashboardSummaryDefinition) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DashboardSummaryDefinition) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSummaryDefinition) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DashboardSummaryDefinition) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DashboardSummaryDefinition) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardSummaryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IntegrationId.IsSet() {
		toSerialize["integration_id"] = o.IntegrationId.Get()
	}
	if o.IsFavorite != nil {
		toSerialize["is_favorite"] = o.IsFavorite
	}
	if o.IsReadOnly != nil {
		toSerialize["is_read_only"] = o.IsReadOnly
	}
	if o.IsShared != nil {
		toSerialize["is_shared"] = o.IsShared
	}
	if o.LastViewDate.IsSet() {
		toSerialize["last_view_date"] = o.LastViewDate.Get()
	}
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if o.Popularity != nil {
		toSerialize["popularity"] = o.Popularity
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardSummaryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author        *datadogV1.Creator           `json:"author,omitempty"`
		Created       datadog.NullableTime         `json:"created,omitempty"`
		Icon          datadog.NullableString       `json:"icon,omitempty"`
		Id            *DashboardSummaryID          `json:"id,omitempty"`
		IntegrationId datadog.NullableString       `json:"integration_id,omitempty"`
		IsFavorite    *bool                        `json:"is_favorite,omitempty"`
		IsReadOnly    *bool                        `json:"is_read_only,omitempty"`
		IsShared      *bool                        `json:"is_shared,omitempty"`
		LastViewDate  datadog.NullableString       `json:"last_view_date,omitempty"`
		Modified      datadog.NullableTime         `json:"modified,omitempty"`
		Popularity    *int32                       `json:"popularity,omitempty"`
		Tags          datadog.NullableList[string] `json:"tags,omitempty"`
		Title         *string                      `json:"title,omitempty"`
		Type          *string                      `json:"type,omitempty"`
		Url           *string                      `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "created", "icon", "id", "integration_id", "is_favorite", "is_read_only", "is_shared", "last_view_date", "modified", "popularity", "tags", "title", "type", "url"})
	} else {
		return err
	}
	o.Author = all.Author
	o.Created = all.Created
	o.Icon = all.Icon
	o.Id = all.Id
	o.IntegrationId = all.IntegrationId
	o.IsFavorite = all.IsFavorite
	o.IsReadOnly = all.IsReadOnly
	o.IsShared = all.IsShared
	o.LastViewDate = all.LastViewDate
	o.Modified = all.Modified
	o.Popularity = all.Popularity
	o.Tags = all.Tags
	o.Title = all.Title
	o.Type = all.Type
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
