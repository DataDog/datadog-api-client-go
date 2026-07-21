// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptVersionDataAttributes Attributes of a specific version of an LLM Observability prompt.
type LLMObsPromptVersionDataAttributes struct {
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
	// A text template or a list of chat messages.
	Template LLMObsPromptTemplate `json:"template"`
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

// NewLLMObsPromptVersionDataAttributes instantiates a new LLMObsPromptVersionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPromptVersionDataAttributes(promptId string, promptUuid string, template LLMObsPromptTemplate, version int64) *LLMObsPromptVersionDataAttributes {
	this := LLMObsPromptVersionDataAttributes{}
	this.PromptId = promptId
	this.PromptUuid = promptUuid
	this.Template = template
	this.Version = version
	return &this
}

// NewLLMObsPromptVersionDataAttributesWithDefaults instantiates a new LLMObsPromptVersionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPromptVersionDataAttributesWithDefaults() *LLMObsPromptVersionDataAttributes {
	this := LLMObsPromptVersionDataAttributes{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *LLMObsPromptVersionDataAttributes) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *LLMObsPromptVersionDataAttributes) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *LLMObsPromptVersionDataAttributes) SetAuthor(v string) {
	o.Author = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LLMObsPromptVersionDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LLMObsPromptVersionDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LLMObsPromptVersionDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDatasets returns the Datasets field value if set, zero value otherwise.
func (o *LLMObsPromptVersionDataAttributes) GetDatasets() []LLMObsPromptDataset {
	if o == nil || o.Datasets == nil {
		var ret []LLMObsPromptDataset
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetDatasetsOk() (*[]LLMObsPromptDataset, bool) {
	if o == nil || o.Datasets == nil {
		return nil, false
	}
	return &o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *LLMObsPromptVersionDataAttributes) HasDatasets() bool {
	return o != nil && o.Datasets != nil
}

// SetDatasets gets a reference to the given []LLMObsPromptDataset and assigns it to the Datasets field.
func (o *LLMObsPromptVersionDataAttributes) SetDatasets(v []LLMObsPromptDataset) {
	o.Datasets = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsPromptVersionDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsPromptVersionDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsPromptVersionDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
// Deprecated
func (o *LLMObsPromptVersionDataAttributes) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LLMObsPromptVersionDataAttributes) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *LLMObsPromptVersionDataAttributes) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
// Deprecated
func (o *LLMObsPromptVersionDataAttributes) SetLabels(v []string) {
	o.Labels = v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *LLMObsPromptVersionDataAttributes) GetLastSeenAt() time.Time {
	if o == nil || o.LastSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *LLMObsPromptVersionDataAttributes) HasLastSeenAt() bool {
	return o != nil && o.LastSeenAt != nil
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *LLMObsPromptVersionDataAttributes) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

// GetMlApp returns the MlApp field value if set, zero value otherwise.
func (o *LLMObsPromptVersionDataAttributes) GetMlApp() string {
	if o == nil || o.MlApp == nil {
		var ret string
		return ret
	}
	return *o.MlApp
}

// GetMlAppOk returns a tuple with the MlApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetMlAppOk() (*string, bool) {
	if o == nil || o.MlApp == nil {
		return nil, false
	}
	return o.MlApp, true
}

// HasMlApp returns a boolean if a field has been set.
func (o *LLMObsPromptVersionDataAttributes) HasMlApp() bool {
	return o != nil && o.MlApp != nil
}

// SetMlApp gets a reference to the given string and assigns it to the MlApp field.
func (o *LLMObsPromptVersionDataAttributes) SetMlApp(v string) {
	o.MlApp = &v
}

// GetMlApps returns the MlApps field value if set, zero value otherwise.
func (o *LLMObsPromptVersionDataAttributes) GetMlApps() []string {
	if o == nil || o.MlApps == nil {
		var ret []string
		return ret
	}
	return o.MlApps
}

// GetMlAppsOk returns a tuple with the MlApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetMlAppsOk() (*[]string, bool) {
	if o == nil || o.MlApps == nil {
		return nil, false
	}
	return &o.MlApps, true
}

// HasMlApps returns a boolean if a field has been set.
func (o *LLMObsPromptVersionDataAttributes) HasMlApps() bool {
	return o != nil && o.MlApps != nil
}

// SetMlApps gets a reference to the given []string and assigns it to the MlApps field.
func (o *LLMObsPromptVersionDataAttributes) SetMlApps(v []string) {
	o.MlApps = v
}

// GetPromptId returns the PromptId field value.
func (o *LLMObsPromptVersionDataAttributes) GetPromptId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PromptId
}

// GetPromptIdOk returns a tuple with the PromptId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetPromptIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptId, true
}

// SetPromptId sets field value.
func (o *LLMObsPromptVersionDataAttributes) SetPromptId(v string) {
	o.PromptId = v
}

// GetPromptUuid returns the PromptUuid field value.
func (o *LLMObsPromptVersionDataAttributes) GetPromptUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PromptUuid
}

// GetPromptUuidOk returns a tuple with the PromptUuid field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetPromptUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptUuid, true
}

// SetPromptUuid sets field value.
func (o *LLMObsPromptVersionDataAttributes) SetPromptUuid(v string) {
	o.PromptUuid = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsPromptVersionDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsPromptVersionDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LLMObsPromptVersionDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTemplate returns the Template field value.
func (o *LLMObsPromptVersionDataAttributes) GetTemplate() LLMObsPromptTemplate {
	if o == nil {
		var ret LLMObsPromptTemplate
		return ret
	}
	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetTemplateOk() (*LLMObsPromptTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value.
func (o *LLMObsPromptVersionDataAttributes) SetTemplate(v LLMObsPromptTemplate) {
	o.Template = v
}

// GetUserVersion returns the UserVersion field value if set, zero value otherwise.
func (o *LLMObsPromptVersionDataAttributes) GetUserVersion() string {
	if o == nil || o.UserVersion == nil {
		var ret string
		return ret
	}
	return *o.UserVersion
}

// GetUserVersionOk returns a tuple with the UserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetUserVersionOk() (*string, bool) {
	if o == nil || o.UserVersion == nil {
		return nil, false
	}
	return o.UserVersion, true
}

// HasUserVersion returns a boolean if a field has been set.
func (o *LLMObsPromptVersionDataAttributes) HasUserVersion() bool {
	return o != nil && o.UserVersion != nil
}

// SetUserVersion gets a reference to the given string and assigns it to the UserVersion field.
func (o *LLMObsPromptVersionDataAttributes) SetUserVersion(v string) {
	o.UserVersion = &v
}

// GetVersion returns the Version field value.
func (o *LLMObsPromptVersionDataAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *LLMObsPromptVersionDataAttributes) SetVersion(v int64) {
	o.Version = v
}

// GetVersionCreatedAt returns the VersionCreatedAt field value if set, zero value otherwise.
func (o *LLMObsPromptVersionDataAttributes) GetVersionCreatedAt() time.Time {
	if o == nil || o.VersionCreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.VersionCreatedAt
}

// GetVersionCreatedAtOk returns a tuple with the VersionCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptVersionDataAttributes) GetVersionCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.VersionCreatedAt == nil {
		return nil, false
	}
	return o.VersionCreatedAt, true
}

// HasVersionCreatedAt returns a boolean if a field has been set.
func (o *LLMObsPromptVersionDataAttributes) HasVersionCreatedAt() bool {
	return o != nil && o.VersionCreatedAt != nil
}

// SetVersionCreatedAt gets a reference to the given time.Time and assigns it to the VersionCreatedAt field.
func (o *LLMObsPromptVersionDataAttributes) SetVersionCreatedAt(v time.Time) {
	o.VersionCreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPromptVersionDataAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["template"] = o.Template
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
func (o *LLMObsPromptVersionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
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
		Template         *LLMObsPromptTemplate `json:"template"`
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
	if all.Template == nil {
		return fmt.Errorf("required field template missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "created_at", "datasets", "description", "labels", "last_seen_at", "ml_app", "ml_apps", "prompt_id", "prompt_uuid", "tags", "template", "user_version", "version", "version_created_at"})
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
	o.Template = *all.Template
	o.UserVersion = all.UserVersion
	o.Version = *all.Version
	o.VersionCreatedAt = all.VersionCreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
