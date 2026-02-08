// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardSearchMetadata Metadata about the dashboard.
type DashboardSearchMetadata struct {
	// User information.
	Author DashboardSearchUser `json:"author"`
	// Time at which the dashboard was created.
	CreatedAt time.Time `json:"created_at"`
	// Time at which the dashboard was deleted, or null if not deleted.
	DeletedAt datadog.NullableTime `json:"deleted_at"`
	// List of domains the dashboard is allowed to be embedded in.
	EmbeddableDomains datadog.NullableList[string] `json:"embeddable_domains"`
	// Dashboard experience type.
	ExperienceType string `json:"experience_type"`
	// When the public dashboard link will expire.
	Expiration datadog.NullableTime `json:"expiration"`
	// Whether the dashboard has monitors.
	HasMonitors datadog.NullableBool `json:"has_monitors"`
	// Whether the dashboard is favorited by the user.
	IsFavorited bool `json:"is_favorited"`
	// Whether the public dashboard owner is deactivated.
	IsPublicDashboardOwnerless bool `json:"is_public_dashboard_ownerless"`
	// Whether the dashboard is shared publicly.
	IsShared bool `json:"is_shared"`
	// Last time the dashboard was accessed.
	LastAccessed datadog.NullableTime `json:"last_accessed"`
	// Time at which the dashboard was last updated.
	ModifiedAt time.Time `json:"modified_at"`
	// Relative measure of dashboard popularity.
	Popularity float64 `json:"popularity"`
	// Published title of the public dashboard.
	PublicTitle string `json:"public_title"`
	// Quality score of the dashboard.
	QualityScore datadog.NullableFloat64 `json:"quality_score"`
	// List of email addresses for invite-only public dashboards.
	ShareList datadog.NullableList[string] `json:"share_list"`
	// Share type of the public dashboard.
	ShareType string `json:"share_type"`
	// User information.
	SharedBy DashboardSearchUser `json:"shared_by"`
	// Status of the public dashboard.
	Status string `json:"status"`
	// Unique public dashboard token.
	Token string `json:"token"`
	// Dashboard type.
	Type string `json:"type"`
	// URL path to the dashboard.
	Url string `json:"url"`
	// Number of widgets in the dashboard.
	WidgetCount int64 `json:"widget_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardSearchMetadata instantiates a new DashboardSearchMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardSearchMetadata(author DashboardSearchUser, createdAt time.Time, deletedAt datadog.NullableTime, embeddableDomains datadog.NullableList[string], experienceType string, expiration datadog.NullableTime, hasMonitors datadog.NullableBool, isFavorited bool, isPublicDashboardOwnerless bool, isShared bool, lastAccessed datadog.NullableTime, modifiedAt time.Time, popularity float64, publicTitle string, qualityScore datadog.NullableFloat64, shareList datadog.NullableList[string], shareType string, sharedBy DashboardSearchUser, status string, token string, typeVar string, url string, widgetCount int64) *DashboardSearchMetadata {
	this := DashboardSearchMetadata{}
	this.Author = author
	this.CreatedAt = createdAt
	this.DeletedAt = deletedAt
	this.EmbeddableDomains = embeddableDomains
	this.ExperienceType = experienceType
	this.Expiration = expiration
	this.HasMonitors = hasMonitors
	this.IsFavorited = isFavorited
	this.IsPublicDashboardOwnerless = isPublicDashboardOwnerless
	this.IsShared = isShared
	this.LastAccessed = lastAccessed
	this.ModifiedAt = modifiedAt
	this.Popularity = popularity
	this.PublicTitle = publicTitle
	this.QualityScore = qualityScore
	this.ShareList = shareList
	this.ShareType = shareType
	this.SharedBy = sharedBy
	this.Status = status
	this.Token = token
	this.Type = typeVar
	this.Url = url
	this.WidgetCount = widgetCount
	return &this
}

// NewDashboardSearchMetadataWithDefaults instantiates a new DashboardSearchMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardSearchMetadataWithDefaults() *DashboardSearchMetadata {
	this := DashboardSearchMetadata{}
	return &this
}

// GetAuthor returns the Author field value.
func (o *DashboardSearchMetadata) GetAuthor() DashboardSearchUser {
	if o == nil {
		var ret DashboardSearchUser
		return ret
	}
	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetAuthorOk() (*DashboardSearchUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value.
func (o *DashboardSearchMetadata) SetAuthor(v DashboardSearchUser) {
	o.Author = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DashboardSearchMetadata) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DashboardSearchMetadata) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeletedAt returns the DeletedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *DashboardSearchMetadata) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSearchMetadata) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// SetDeletedAt sets field value.
func (o *DashboardSearchMetadata) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// GetEmbeddableDomains returns the EmbeddableDomains field value.
// If the value is explicit nil, the zero value for []string will be returned.
func (o *DashboardSearchMetadata) GetEmbeddableDomains() []string {
	if o == nil || o.EmbeddableDomains.Get() == nil {
		var ret []string
		return ret
	}
	return *o.EmbeddableDomains.Get()
}

// GetEmbeddableDomainsOk returns a tuple with the EmbeddableDomains field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSearchMetadata) GetEmbeddableDomainsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmbeddableDomains.Get(), o.EmbeddableDomains.IsSet()
}

// SetEmbeddableDomains sets field value.
func (o *DashboardSearchMetadata) SetEmbeddableDomains(v []string) {
	o.EmbeddableDomains.Set(&v)
}

// GetExperienceType returns the ExperienceType field value.
func (o *DashboardSearchMetadata) GetExperienceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExperienceType
}

// GetExperienceTypeOk returns a tuple with the ExperienceType field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetExperienceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExperienceType, true
}

// SetExperienceType sets field value.
func (o *DashboardSearchMetadata) SetExperienceType(v string) {
	o.ExperienceType = v
}

// GetExpiration returns the Expiration field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *DashboardSearchMetadata) GetExpiration() time.Time {
	if o == nil || o.Expiration.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSearchMetadata) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// SetExpiration sets field value.
func (o *DashboardSearchMetadata) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}

// GetHasMonitors returns the HasMonitors field value.
// If the value is explicit nil, the zero value for bool will be returned.
func (o *DashboardSearchMetadata) GetHasMonitors() bool {
	if o == nil || o.HasMonitors.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasMonitors.Get()
}

// GetHasMonitorsOk returns a tuple with the HasMonitors field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSearchMetadata) GetHasMonitorsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasMonitors.Get(), o.HasMonitors.IsSet()
}

// SetHasMonitors sets field value.
func (o *DashboardSearchMetadata) SetHasMonitors(v bool) {
	o.HasMonitors.Set(&v)
}

// GetIsFavorited returns the IsFavorited field value.
func (o *DashboardSearchMetadata) GetIsFavorited() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsFavorited
}

// GetIsFavoritedOk returns a tuple with the IsFavorited field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetIsFavoritedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFavorited, true
}

// SetIsFavorited sets field value.
func (o *DashboardSearchMetadata) SetIsFavorited(v bool) {
	o.IsFavorited = v
}

// GetIsPublicDashboardOwnerless returns the IsPublicDashboardOwnerless field value.
func (o *DashboardSearchMetadata) GetIsPublicDashboardOwnerless() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPublicDashboardOwnerless
}

// GetIsPublicDashboardOwnerlessOk returns a tuple with the IsPublicDashboardOwnerless field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetIsPublicDashboardOwnerlessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublicDashboardOwnerless, true
}

// SetIsPublicDashboardOwnerless sets field value.
func (o *DashboardSearchMetadata) SetIsPublicDashboardOwnerless(v bool) {
	o.IsPublicDashboardOwnerless = v
}

// GetIsShared returns the IsShared field value.
func (o *DashboardSearchMetadata) GetIsShared() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetIsSharedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsShared, true
}

// SetIsShared sets field value.
func (o *DashboardSearchMetadata) SetIsShared(v bool) {
	o.IsShared = v
}

// GetLastAccessed returns the LastAccessed field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *DashboardSearchMetadata) GetLastAccessed() time.Time {
	if o == nil || o.LastAccessed.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessed.Get()
}

// GetLastAccessedOk returns a tuple with the LastAccessed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSearchMetadata) GetLastAccessedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastAccessed.Get(), o.LastAccessed.IsSet()
}

// SetLastAccessed sets field value.
func (o *DashboardSearchMetadata) SetLastAccessed(v time.Time) {
	o.LastAccessed.Set(&v)
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *DashboardSearchMetadata) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *DashboardSearchMetadata) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetPopularity returns the Popularity field value.
func (o *DashboardSearchMetadata) GetPopularity() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetPopularityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Popularity, true
}

// SetPopularity sets field value.
func (o *DashboardSearchMetadata) SetPopularity(v float64) {
	o.Popularity = v
}

// GetPublicTitle returns the PublicTitle field value.
func (o *DashboardSearchMetadata) GetPublicTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PublicTitle
}

// GetPublicTitleOk returns a tuple with the PublicTitle field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetPublicTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicTitle, true
}

// SetPublicTitle sets field value.
func (o *DashboardSearchMetadata) SetPublicTitle(v string) {
	o.PublicTitle = v
}

// GetQualityScore returns the QualityScore field value.
// If the value is explicit nil, the zero value for float64 will be returned.
func (o *DashboardSearchMetadata) GetQualityScore() float64 {
	if o == nil || o.QualityScore.Get() == nil {
		var ret float64
		return ret
	}
	return *o.QualityScore.Get()
}

// GetQualityScoreOk returns a tuple with the QualityScore field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSearchMetadata) GetQualityScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.QualityScore.Get(), o.QualityScore.IsSet()
}

// SetQualityScore sets field value.
func (o *DashboardSearchMetadata) SetQualityScore(v float64) {
	o.QualityScore.Set(&v)
}

// GetShareList returns the ShareList field value.
// If the value is explicit nil, the zero value for []string will be returned.
func (o *DashboardSearchMetadata) GetShareList() []string {
	if o == nil || o.ShareList.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ShareList.Get()
}

// GetShareListOk returns a tuple with the ShareList field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DashboardSearchMetadata) GetShareListOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareList.Get(), o.ShareList.IsSet()
}

// SetShareList sets field value.
func (o *DashboardSearchMetadata) SetShareList(v []string) {
	o.ShareList.Set(&v)
}

// GetShareType returns the ShareType field value.
func (o *DashboardSearchMetadata) GetShareType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShareType
}

// GetShareTypeOk returns a tuple with the ShareType field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetShareTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareType, true
}

// SetShareType sets field value.
func (o *DashboardSearchMetadata) SetShareType(v string) {
	o.ShareType = v
}

// GetSharedBy returns the SharedBy field value.
func (o *DashboardSearchMetadata) GetSharedBy() DashboardSearchUser {
	if o == nil {
		var ret DashboardSearchUser
		return ret
	}
	return o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetSharedByOk() (*DashboardSearchUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SharedBy, true
}

// SetSharedBy sets field value.
func (o *DashboardSearchMetadata) SetSharedBy(v DashboardSearchUser) {
	o.SharedBy = v
}

// GetStatus returns the Status field value.
func (o *DashboardSearchMetadata) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DashboardSearchMetadata) SetStatus(v string) {
	o.Status = v
}

// GetToken returns the Token field value.
func (o *DashboardSearchMetadata) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value.
func (o *DashboardSearchMetadata) SetToken(v string) {
	o.Token = v
}

// GetType returns the Type field value.
func (o *DashboardSearchMetadata) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DashboardSearchMetadata) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value.
func (o *DashboardSearchMetadata) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *DashboardSearchMetadata) SetUrl(v string) {
	o.Url = v
}

// GetWidgetCount returns the WidgetCount field value.
func (o *DashboardSearchMetadata) GetWidgetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.WidgetCount
}

// GetWidgetCountOk returns a tuple with the WidgetCount field value
// and a boolean to check if the value has been set.
func (o *DashboardSearchMetadata) GetWidgetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WidgetCount, true
}

// SetWidgetCount sets field value.
func (o *DashboardSearchMetadata) SetWidgetCount(v int64) {
	o.WidgetCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardSearchMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["author"] = o.Author
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["deleted_at"] = o.DeletedAt.Get()
	toSerialize["embeddable_domains"] = o.EmbeddableDomains.Get()
	toSerialize["experience_type"] = o.ExperienceType
	toSerialize["expiration"] = o.Expiration.Get()
	toSerialize["has_monitors"] = o.HasMonitors.Get()
	toSerialize["is_favorited"] = o.IsFavorited
	toSerialize["is_public_dashboard_ownerless"] = o.IsPublicDashboardOwnerless
	toSerialize["is_shared"] = o.IsShared
	toSerialize["last_accessed"] = o.LastAccessed.Get()
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["popularity"] = o.Popularity
	toSerialize["public_title"] = o.PublicTitle
	toSerialize["quality_score"] = o.QualityScore.Get()
	toSerialize["share_list"] = o.ShareList.Get()
	toSerialize["share_type"] = o.ShareType
	toSerialize["shared_by"] = o.SharedBy
	toSerialize["status"] = o.Status
	toSerialize["token"] = o.Token
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	toSerialize["widget_count"] = o.WidgetCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardSearchMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author                     *DashboardSearchUser         `json:"author"`
		CreatedAt                  *time.Time                   `json:"created_at"`
		DeletedAt                  datadog.NullableTime         `json:"deleted_at"`
		EmbeddableDomains          datadog.NullableList[string] `json:"embeddable_domains"`
		ExperienceType             *string                      `json:"experience_type"`
		Expiration                 datadog.NullableTime         `json:"expiration"`
		HasMonitors                datadog.NullableBool         `json:"has_monitors"`
		IsFavorited                *bool                        `json:"is_favorited"`
		IsPublicDashboardOwnerless *bool                        `json:"is_public_dashboard_ownerless"`
		IsShared                   *bool                        `json:"is_shared"`
		LastAccessed               datadog.NullableTime         `json:"last_accessed"`
		ModifiedAt                 *time.Time                   `json:"modified_at"`
		Popularity                 *float64                     `json:"popularity"`
		PublicTitle                *string                      `json:"public_title"`
		QualityScore               datadog.NullableFloat64      `json:"quality_score"`
		ShareList                  datadog.NullableList[string] `json:"share_list"`
		ShareType                  *string                      `json:"share_type"`
		SharedBy                   *DashboardSearchUser         `json:"shared_by"`
		Status                     *string                      `json:"status"`
		Token                      *string                      `json:"token"`
		Type                       *string                      `json:"type"`
		Url                        *string                      `json:"url"`
		WidgetCount                *int64                       `json:"widget_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Author == nil {
		return fmt.Errorf("required field author missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if !all.DeletedAt.IsSet() {
		return fmt.Errorf("required field deleted_at missing")
	}
	if !all.EmbeddableDomains.IsSet() {
		return fmt.Errorf("required field embeddable_domains missing")
	}
	if all.ExperienceType == nil {
		return fmt.Errorf("required field experience_type missing")
	}
	if !all.Expiration.IsSet() {
		return fmt.Errorf("required field expiration missing")
	}
	if !all.HasMonitors.IsSet() {
		return fmt.Errorf("required field has_monitors missing")
	}
	if all.IsFavorited == nil {
		return fmt.Errorf("required field is_favorited missing")
	}
	if all.IsPublicDashboardOwnerless == nil {
		return fmt.Errorf("required field is_public_dashboard_ownerless missing")
	}
	if all.IsShared == nil {
		return fmt.Errorf("required field is_shared missing")
	}
	if !all.LastAccessed.IsSet() {
		return fmt.Errorf("required field last_accessed missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.Popularity == nil {
		return fmt.Errorf("required field popularity missing")
	}
	if all.PublicTitle == nil {
		return fmt.Errorf("required field public_title missing")
	}
	if !all.QualityScore.IsSet() {
		return fmt.Errorf("required field quality_score missing")
	}
	if !all.ShareList.IsSet() {
		return fmt.Errorf("required field share_list missing")
	}
	if all.ShareType == nil {
		return fmt.Errorf("required field share_type missing")
	}
	if all.SharedBy == nil {
		return fmt.Errorf("required field shared_by missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Token == nil {
		return fmt.Errorf("required field token missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	if all.WidgetCount == nil {
		return fmt.Errorf("required field widget_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "created_at", "deleted_at", "embeddable_domains", "experience_type", "expiration", "has_monitors", "is_favorited", "is_public_dashboard_ownerless", "is_shared", "last_accessed", "modified_at", "popularity", "public_title", "quality_score", "share_list", "share_type", "shared_by", "status", "token", "type", "url", "widget_count"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = *all.Author
	o.CreatedAt = *all.CreatedAt
	o.DeletedAt = all.DeletedAt
	o.EmbeddableDomains = all.EmbeddableDomains
	o.ExperienceType = *all.ExperienceType
	o.Expiration = all.Expiration
	o.HasMonitors = all.HasMonitors
	o.IsFavorited = *all.IsFavorited
	o.IsPublicDashboardOwnerless = *all.IsPublicDashboardOwnerless
	o.IsShared = *all.IsShared
	o.LastAccessed = all.LastAccessed
	o.ModifiedAt = *all.ModifiedAt
	o.Popularity = *all.Popularity
	o.PublicTitle = *all.PublicTitle
	o.QualityScore = all.QualityScore
	o.ShareList = all.ShareList
	o.ShareType = *all.ShareType
	if all.SharedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SharedBy = *all.SharedBy
	o.Status = *all.Status
	o.Token = *all.Token
	o.Type = *all.Type
	o.Url = *all.Url
	o.WidgetCount = *all.WidgetCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
