// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptVersionListDataAttributes Attributes of a prompt version returned in a list, excluding its template.
type LLMObsPromptVersionListDataAttributes struct {
	// UUID of the user who authored this version.
	Author *string `json:"author,omitempty"`
	// Timestamp stored on this prompt version.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Datasets observed in runs associated with this prompt version.
	Datasets []LLMObsPromptDataset `json:"datasets,omitempty"`
	// Description of this version.
	Description *string `json:"description,omitempty"`
	// Labels attached to this version (for example `development`, `staging`, `production`).
	// Deprecated
	Labels []string `json:"labels,omitempty"`
	// Timestamp of the most recent observed run of this prompt version.
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`
	// The ML application this prompt is associated with.
	MlApp *string `json:"ml_app,omitempty"`
	// ML applications observed running this prompt version.
	MlApps []string `json:"ml_apps,omitempty"`
	// Customer-provided identifier of the parent prompt.
	PromptId string `json:"prompt_id"`
	// Unique identifier of the parent prompt.
	PromptUuid string `json:"prompt_uuid"`
	// Tags observed on runs of this prompt version.
	Tags []string `json:"tags,omitempty"`
	// User-supplied identifier for this version.
	UserVersion *string `json:"user_version,omitempty"`
	// Sequential version number.
	Version int64 `json:"version"`
	// Timestamp when this version was created.
	VersionCreatedAt *time.Time `json:"version_created_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPromptVersionListDataAttributes instantiates a new LLMObsPromptVersionListDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPromptVersionListDataAttributes(promptId string, promptUuid string, version int64) *LLMObsPromptVersionListDataAttributes {
	this := LLMObsPromptVersionListDataAttributes{}
	this.PromptId = promptId
	this.PromptUuid = promptUuid
	this.Version = version
	return &this
}

// NewLLMObsPromptVersionListDataAttributesWithDefaults instantiates a new LLMObsPromptVersionListDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPromptVersionListDataAttributesWithDefaults() *LLMObsPromptVersionListDataAttributes {
	this := LLMObsPromptVersionListDataAttributes{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *LLMObsPromptVersionListDataAttributes) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *LLMObsPromptVersionListDataAttributes) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *LLMObsPromptVersionListDataAttributes) SetAuthor(v string) {
	o.Author = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LLMObsPromptVersionListDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LLMObsPromptVersionListDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LLMObsPromptVersionListDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDatasets returns the Datasets field value if set, zero value otherwise.
func (o *LLMObsPromptVersionListDataAttributes) GetDatasets() []LLMObsPromptDataset {
	if o == nil || o.Datasets == nil {
		var ret []LLMObsPromptDataset
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetDatasetsOk() (*[]LLMObsPromptDataset, bool) {
	if o == nil || o.Datasets == nil {
		return nil, false
	}
	return &o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *LLMObsPromptVersionListDataAttributes) HasDatasets() bool {
	return o != nil && o.Datasets != nil
}

// SetDatasets gets a reference to the given []LLMObsPromptDataset and assigns it to the Datasets field.
func (o *LLMObsPromptVersionListDataAttributes) SetDatasets(v []LLMObsPromptDataset) {
	o.Datasets = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsPromptVersionListDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsPromptVersionListDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsPromptVersionListDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
// Deprecated
func (o *LLMObsPromptVersionListDataAttributes) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LLMObsPromptVersionListDataAttributes) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LLMObsPromptVersionListDataAttributes) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
// Deprecated
func (o *LLMObsPromptVersionListDataAttributes) SetLabels(v []string) {
	o.Labels = v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *LLMObsPromptVersionListDataAttributes) GetLastSeenAt() time.Time {
	if o == nil || o.LastSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *LLMObsPromptVersionListDataAttributes) HasLastSeenAt() bool {
	return o != nil && o.LastSeenAt != nil
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *LLMObsPromptVersionListDataAttributes) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

// GetMlApp returns the MlApp field value if set, zero value otherwise.
func (o *LLMObsPromptVersionListDataAttributes) GetMlApp() string {
	if o == nil || o.MlApp == nil {
		var ret string
		return ret
	}
	return *o.MlApp
}

// GetMlAppOk returns a tuple with the MlApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetMlAppOk() (*string, bool) {
	if o == nil || o.MlApp == nil {
		return nil, false
	}
	return o.MlApp, true
}

// HasMlApp returns a boolean if a field has been set.
func (o *LLMObsPromptVersionListDataAttributes) HasMlApp() bool {
	return o != nil && o.MlApp != nil
}

// SetMlApp gets a reference to the given string and assigns it to the MlApp field.
func (o *LLMObsPromptVersionListDataAttributes) SetMlApp(v string) {
	o.MlApp = &v
}

// GetMlApps returns the MlApps field value if set, zero value otherwise.
func (o *LLMObsPromptVersionListDataAttributes) GetMlApps() []string {
	if o == nil || o.MlApps == nil {
		var ret []string
		return ret
	}
	return o.MlApps
}

// GetMlAppsOk returns a tuple with the MlApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetMlAppsOk() (*[]string, bool) {
	if o == nil || o.MlApps == nil {
		return nil, false
	}
	return &o.MlApps, true
}

// HasMlApps returns a boolean if a field has been set.
func (o *LLMObsPromptVersionListDataAttributes) HasMlApps() bool {
	return o != nil && o.MlApps != nil
}

// SetMlApps gets a reference to the given []string and assigns it to the MlApps field.
func (o *LLMObsPromptVersionListDataAttributes) SetMlApps(v []string) {
	o.MlApps = v
}

// GetPromptId returns the PromptId field value.
func (o *LLMObsPromptVersionListDataAttributes) GetPromptId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PromptId
}

// GetPromptIdOk returns a tuple with the PromptId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetPromptIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptId, true
}

// SetPromptId sets field value.
func (o *LLMObsPromptVersionListDataAttributes) SetPromptId(v string) {
	o.PromptId = v
}

// GetPromptUuid returns the PromptUuid field value.
func (o *LLMObsPromptVersionListDataAttributes) GetPromptUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PromptUuid
}

// GetPromptUuidOk returns a tuple with the PromptUuid field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetPromptUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptUuid, true
}

// SetPromptUuid sets field value.
func (o *LLMObsPromptVersionListDataAttributes) SetPromptUuid(v string) {
	o.PromptUuid = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsPromptVersionListDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsPromptVersionListDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LLMObsPromptVersionListDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetUserVersion returns the UserVersion field value if set, zero value otherwise.
func (o *LLMObsPromptVersionListDataAttributes) GetUserVersion() string {
	if o == nil || o.UserVersion == nil {
		var ret string
		return ret
	}
	return *o.UserVersion
}

// GetUserVersionOk returns a tuple with the UserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetUserVersionOk() (*string, bool) {
	if o == nil || o.UserVersion == nil {
		return nil, false
	}
	return o.UserVersion, true
}

// HasUserVersion returns a boolean if a field has been set.
func (o *LLMObsPromptVersionListDataAttributes) HasUserVersion() bool {
	return o != nil && o.UserVersion != nil
}

// SetUserVersion gets a reference to the given string and assigns it to the UserVersion field.
func (o *LLMObsPromptVersionListDataAttributes) SetUserVersion(v string) {
	o.UserVersion = &v
}

// GetVersion returns the Version field value.
func (o *LLMObsPromptVersionListDataAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *LLMObsPromptVersionListDataAttributes) SetVersion(v int64) {
	o.Version = v
}

// GetVersionCreatedAt returns the VersionCreatedAt field value if set, zero value otherwise.
func (o *LLMObsPromptVersionListDataAttributes) GetVersionCreatedAt() time.Time {
	if o == nil || o.VersionCreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.VersionCreatedAt
}

// GetVersionCreatedAtOk returns a tuple with the VersionCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionListDataAttributes) GetVersionCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.VersionCreatedAt == nil {
		return nil, false
	}
	return o.VersionCreatedAt, true
}

// HasVersionCreatedAt returns a boolean if a field has been set.
func (o *LLMObsPromptVersionListDataAttributes) HasVersionCreatedAt() bool {
	return o != nil && o.VersionCreatedAt != nil
}

// SetVersionCreatedAt gets a reference to the given time.Time and assigns it to the VersionCreatedAt field.
func (o *LLMObsPromptVersionListDataAttributes) SetVersionCreatedAt(v time.Time) {
	o.VersionCreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPromptVersionListDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Datasets != nil {
		toSerialize["datasets"] = o.Datasets
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.LastSeenAt != nil {
		if o.LastSeenAt.Nanosecond() == 0 {
			toSerialize["last_seen_at"] = o.LastSeenAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_seen_at"] = o.LastSeenAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.MlApp != nil {
		toSerialize["ml_app"] = o.MlApp
	}
	if o.MlApps != nil {
		toSerialize["ml_apps"] = o.MlApps
	}
	toSerialize["prompt_id"] = o.PromptId
	toSerialize["prompt_uuid"] = o.PromptUuid
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UserVersion != nil {
		toSerialize["user_version"] = o.UserVersion
	}
	toSerialize["version"] = o.Version
	if o.VersionCreatedAt != nil {
		if o.VersionCreatedAt.Nanosecond() == 0 {
			toSerialize["version_created_at"] = o.VersionCreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["version_created_at"] = o.VersionCreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPromptVersionListDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author           *string               `json:"author,omitempty"`
		CreatedAt        *time.Time            `json:"created_at,omitempty"`
		Datasets         []LLMObsPromptDataset `json:"datasets,omitempty"`
		Description      *string               `json:"description,omitempty"`
		Labels           []string              `json:"labels,omitempty"`
		LastSeenAt       *time.Time            `json:"last_seen_at,omitempty"`
		MlApp            *string               `json:"ml_app,omitempty"`
		MlApps           []string              `json:"ml_apps,omitempty"`
		PromptId         *string               `json:"prompt_id"`
		PromptUuid       *string               `json:"prompt_uuid"`
		Tags             []string              `json:"tags,omitempty"`
		UserVersion      *string               `json:"user_version,omitempty"`
		Version          *int64                `json:"version"`
		VersionCreatedAt *time.Time            `json:"version_created_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PromptId == nil {
		return fmt.Errorf("required field prompt_id missing")
	}
	if all.PromptUuid == nil {
		return fmt.Errorf("required field prompt_uuid missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "created_at", "datasets", "description", "labels", "last_seen_at", "ml_app", "ml_apps", "prompt_id", "prompt_uuid", "tags", "user_version", "version", "version_created_at"})
	} else {
		return err
	}
	o.Author = all.Author
	o.CreatedAt = all.CreatedAt
	o.Datasets = all.Datasets
	o.Description = all.Description
	o.Labels = all.Labels
	o.LastSeenAt = all.LastSeenAt
	o.MlApp = all.MlApp
	o.MlApps = all.MlApps
	o.PromptId = *all.PromptId
	o.PromptUuid = *all.PromptUuid
	o.Tags = all.Tags
	o.UserVersion = all.UserVersion
	o.Version = *all.Version
	o.VersionCreatedAt = all.VersionCreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
