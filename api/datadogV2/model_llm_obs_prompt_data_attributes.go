// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPromptDataAttributes Attributes of an LLM Observability prompt registry entry.
type LLMObsPromptDataAttributes struct {
	// UUID of the user who authored the prompt.
	Author *string `json:"author,omitempty"`
	// Timestamp when the prompt was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Source that created the prompt, such as `ui-registry`, `sdk-registry`, or `sdk-instrumentation`.
	CreatedFrom string `json:"created_from"`
	// Datasets observed in runs associated with this prompt.
	Datasets []LLMObsPromptDataset `json:"datasets,omitempty"`
	// Description of the prompt.
	Description *string `json:"description,omitempty"`
	// Source prompt from which this prompt was extracted, when applicable.
	ExtractedFrom *string `json:"extracted_from,omitempty"`
	// Whether the prompt is a registry entry (as opposed to a code-discovered prompt).
	InRegistry bool `json:"in_registry"`
	// Timestamp of the most recent observed run of this prompt.
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`
	// Timestamp when the most recent version of the prompt was created.
	LastVersionCreatedAt *time.Time `json:"last_version_created_at,omitempty"`
	// The ML application this prompt is associated with.
	MlApp *string `json:"ml_app,omitempty"`
	// ML applications observed running this prompt.
	MlApps []string `json:"ml_apps,omitempty"`
	// Number of versions of the prompt.
	NumVersions int64 `json:"num_versions"`
	// Customer-provided identifier of the prompt.
	PromptId string `json:"prompt_id"`
	// Whether the prompt was created from the registry or discovered from observed LLM calls.
	Source LLMObsPromptResponseSource `json:"source"`
	// Tags observed on runs of this prompt.
	Tags []string `json:"tags,omitempty"`
	// Title of the prompt.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPromptDataAttributes instantiates a new LLMObsPromptDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPromptDataAttributes(createdFrom string, inRegistry bool, numVersions int64, promptId string, source LLMObsPromptResponseSource) *LLMObsPromptDataAttributes {
	this := LLMObsPromptDataAttributes{}
	this.CreatedFrom = createdFrom
	this.InRegistry = inRegistry
	this.NumVersions = numVersions
	this.PromptId = promptId
	this.Source = source
	return &this
}

// NewLLMObsPromptDataAttributesWithDefaults instantiates a new LLMObsPromptDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPromptDataAttributesWithDefaults() *LLMObsPromptDataAttributes {
	this := LLMObsPromptDataAttributes{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *LLMObsPromptDataAttributes) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *LLMObsPromptDataAttributes) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *LLMObsPromptDataAttributes) SetAuthor(v string) {
	o.Author = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LLMObsPromptDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LLMObsPromptDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LLMObsPromptDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedFrom returns the CreatedFrom field value.
func (o *LLMObsPromptDataAttributes) GetCreatedFrom() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedFrom
}

// GetCreatedFromOk returns a tuple with the CreatedFrom field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetCreatedFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedFrom, true
}

// SetCreatedFrom sets field value.
func (o *LLMObsPromptDataAttributes) SetCreatedFrom(v string) {
	o.CreatedFrom = v
}

// GetDatasets returns the Datasets field value if set, zero value otherwise.
func (o *LLMObsPromptDataAttributes) GetDatasets() []LLMObsPromptDataset {
	if o == nil || o.Datasets == nil {
		var ret []LLMObsPromptDataset
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetDatasetsOk() (*[]LLMObsPromptDataset, bool) {
	if o == nil || o.Datasets == nil {
		return nil, false
	}
	return &o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *LLMObsPromptDataAttributes) HasDatasets() bool {
	return o != nil && o.Datasets != nil
}

// SetDatasets gets a reference to the given []LLMObsPromptDataset and assigns it to the Datasets field.
func (o *LLMObsPromptDataAttributes) SetDatasets(v []LLMObsPromptDataset) {
	o.Datasets = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsPromptDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsPromptDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsPromptDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetExtractedFrom returns the ExtractedFrom field value if set, zero value otherwise.
func (o *LLMObsPromptDataAttributes) GetExtractedFrom() string {
	if o == nil || o.ExtractedFrom == nil {
		var ret string
		return ret
	}
	return *o.ExtractedFrom
}

// GetExtractedFromOk returns a tuple with the ExtractedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetExtractedFromOk() (*string, bool) {
	if o == nil || o.ExtractedFrom == nil {
		return nil, false
	}
	return o.ExtractedFrom, true
}

// HasExtractedFrom returns a boolean if a field has been set.
func (o *LLMObsPromptDataAttributes) HasExtractedFrom() bool {
	return o != nil && o.ExtractedFrom != nil
}

// SetExtractedFrom gets a reference to the given string and assigns it to the ExtractedFrom field.
func (o *LLMObsPromptDataAttributes) SetExtractedFrom(v string) {
	o.ExtractedFrom = &v
}

// GetInRegistry returns the InRegistry field value.
func (o *LLMObsPromptDataAttributes) GetInRegistry() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.InRegistry
}

// GetInRegistryOk returns a tuple with the InRegistry field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetInRegistryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InRegistry, true
}

// SetInRegistry sets field value.
func (o *LLMObsPromptDataAttributes) SetInRegistry(v bool) {
	o.InRegistry = v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *LLMObsPromptDataAttributes) GetLastSeenAt() time.Time {
	if o == nil || o.LastSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *LLMObsPromptDataAttributes) HasLastSeenAt() bool {
	return o != nil && o.LastSeenAt != nil
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *LLMObsPromptDataAttributes) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

// GetLastVersionCreatedAt returns the LastVersionCreatedAt field value if set, zero value otherwise.
func (o *LLMObsPromptDataAttributes) GetLastVersionCreatedAt() time.Time {
	if o == nil || o.LastVersionCreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastVersionCreatedAt
}

// GetLastVersionCreatedAtOk returns a tuple with the LastVersionCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetLastVersionCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.LastVersionCreatedAt == nil {
		return nil, false
	}
	return o.LastVersionCreatedAt, true
}

// HasLastVersionCreatedAt returns a boolean if a field has been set.
func (o *LLMObsPromptDataAttributes) HasLastVersionCreatedAt() bool {
	return o != nil && o.LastVersionCreatedAt != nil
}

// SetLastVersionCreatedAt gets a reference to the given time.Time and assigns it to the LastVersionCreatedAt field.
func (o *LLMObsPromptDataAttributes) SetLastVersionCreatedAt(v time.Time) {
	o.LastVersionCreatedAt = &v
}

// GetMlApp returns the MlApp field value if set, zero value otherwise.
func (o *LLMObsPromptDataAttributes) GetMlApp() string {
	if o == nil || o.MlApp == nil {
		var ret string
		return ret
	}
	return *o.MlApp
}

// GetMlAppOk returns a tuple with the MlApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetMlAppOk() (*string, bool) {
	if o == nil || o.MlApp == nil {
		return nil, false
	}
	return o.MlApp, true
}

// HasMlApp returns a boolean if a field has been set.
func (o *LLMObsPromptDataAttributes) HasMlApp() bool {
	return o != nil && o.MlApp != nil
}

// SetMlApp gets a reference to the given string and assigns it to the MlApp field.
func (o *LLMObsPromptDataAttributes) SetMlApp(v string) {
	o.MlApp = &v
}

// GetMlApps returns the MlApps field value if set, zero value otherwise.
func (o *LLMObsPromptDataAttributes) GetMlApps() []string {
	if o == nil || o.MlApps == nil {
		var ret []string
		return ret
	}
	return o.MlApps
}

// GetMlAppsOk returns a tuple with the MlApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetMlAppsOk() (*[]string, bool) {
	if o == nil || o.MlApps == nil {
		return nil, false
	}
	return &o.MlApps, true
}

// HasMlApps returns a boolean if a field has been set.
func (o *LLMObsPromptDataAttributes) HasMlApps() bool {
	return o != nil && o.MlApps != nil
}

// SetMlApps gets a reference to the given []string and assigns it to the MlApps field.
func (o *LLMObsPromptDataAttributes) SetMlApps(v []string) {
	o.MlApps = v
}

// GetNumVersions returns the NumVersions field value.
func (o *LLMObsPromptDataAttributes) GetNumVersions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.NumVersions
}

// GetNumVersionsOk returns a tuple with the NumVersions field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetNumVersionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumVersions, true
}

// SetNumVersions sets field value.
func (o *LLMObsPromptDataAttributes) SetNumVersions(v int64) {
	o.NumVersions = v
}

// GetPromptId returns the PromptId field value.
func (o *LLMObsPromptDataAttributes) GetPromptId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PromptId
}

// GetPromptIdOk returns a tuple with the PromptId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetPromptIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptId, true
}

// SetPromptId sets field value.
func (o *LLMObsPromptDataAttributes) SetPromptId(v string) {
	o.PromptId = v
}

// GetSource returns the Source field value.
func (o *LLMObsPromptDataAttributes) GetSource() LLMObsPromptResponseSource {
	if o == nil {
		var ret LLMObsPromptResponseSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetSourceOk() (*LLMObsPromptResponseSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *LLMObsPromptDataAttributes) SetSource(v LLMObsPromptResponseSource) {
	o.Source = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LLMObsPromptDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LLMObsPromptDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LLMObsPromptDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LLMObsPromptDataAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsPromptDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LLMObsPromptDataAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LLMObsPromptDataAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPromptDataAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["created_from"] = o.CreatedFrom
	if o.Datasets != nil {
		toSerialize["datasets"] = o.Datasets
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExtractedFrom != nil {
		toSerialize["extracted_from"] = o.ExtractedFrom
	}
	toSerialize["in_registry"] = o.InRegistry
	if o.LastSeenAt != nil {
		if o.LastSeenAt.Nanosecond() == 0 {
			toSerialize["last_seen_at"] = o.LastSeenAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_seen_at"] = o.LastSeenAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LastVersionCreatedAt != nil {
		if o.LastVersionCreatedAt.Nanosecond() == 0 {
			toSerialize["last_version_created_at"] = o.LastVersionCreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_version_created_at"] = o.LastVersionCreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.MlApp != nil {
		toSerialize["ml_app"] = o.MlApp
	}
	if o.MlApps != nil {
		toSerialize["ml_apps"] = o.MlApps
	}
	toSerialize["num_versions"] = o.NumVersions
	toSerialize["prompt_id"] = o.PromptId
	toSerialize["source"] = o.Source
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPromptDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author               *string                     `json:"author,omitempty"`
		CreatedAt            *time.Time                  `json:"created_at,omitempty"`
		CreatedFrom          *string                     `json:"created_from"`
		Datasets             []LLMObsPromptDataset       `json:"datasets,omitempty"`
		Description          *string                     `json:"description,omitempty"`
		ExtractedFrom        *string                     `json:"extracted_from,omitempty"`
		InRegistry           *bool                       `json:"in_registry"`
		LastSeenAt           *time.Time                  `json:"last_seen_at,omitempty"`
		LastVersionCreatedAt *time.Time                  `json:"last_version_created_at,omitempty"`
		MlApp                *string                     `json:"ml_app,omitempty"`
		MlApps               []string                    `json:"ml_apps,omitempty"`
		NumVersions          *int64                      `json:"num_versions"`
		PromptId             *string                     `json:"prompt_id"`
		Source               *LLMObsPromptResponseSource `json:"source"`
		Tags                 []string                    `json:"tags,omitempty"`
		Title                *string                     `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedFrom == nil {
		return fmt.Errorf("required field created_from missing")
	}
	if all.InRegistry == nil {
		return fmt.Errorf("required field in_registry missing")
	}
	if all.NumVersions == nil {
		return fmt.Errorf("required field num_versions missing")
	}
	if all.PromptId == nil {
		return fmt.Errorf("required field prompt_id missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "created_at", "created_from", "datasets", "description", "extracted_from", "in_registry", "last_seen_at", "last_version_created_at", "ml_app", "ml_apps", "num_versions", "prompt_id", "source", "tags", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Author = all.Author
	o.CreatedAt = all.CreatedAt
	o.CreatedFrom = *all.CreatedFrom
	o.Datasets = all.Datasets
	o.Description = all.Description
	o.ExtractedFrom = all.ExtractedFrom
	o.InRegistry = *all.InRegistry
	o.LastSeenAt = all.LastSeenAt
	o.LastVersionCreatedAt = all.LastVersionCreatedAt
	o.MlApp = all.MlApp
	o.MlApps = all.MlApps
	o.NumVersions = *all.NumVersions
	o.PromptId = *all.PromptId
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.Tags = all.Tags
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
