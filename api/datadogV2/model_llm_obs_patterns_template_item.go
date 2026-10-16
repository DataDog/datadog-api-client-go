// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsTemplateItem A single Agent Observability patterns template in a list response.
type LLMObsPatternsTemplateItem struct {
	// Instructions guiding how topics are clustered for configurations linked to this template.
	ClusteringInstructions string `json:"clustering_instructions"`
	// Timestamp when the template was created.
	CreatedAt time.Time `json:"created_at"`
	// Human-readable description of what the template is for.
	Description string `json:"description"`
	// Prefix merged into the query that selects spans for configurations linked to this template.
	EvpQueryPrefix string `json:"evp_query_prefix"`
	// Unique identifier of the template.
	Id string `json:"id"`
	// Name of the template.
	Name string `json:"name"`
	// Scope the template applies to.
	Scope string `json:"scope"`
	// Template used to guide topic generation for configurations linked to this template.
	Template string `json:"template"`
	// Default lookback window, in milliseconds, for configurations linked to this template.
	TimeFrameDurationMs int64 `json:"time_frame_duration_ms"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsTemplateItem instantiates a new LLMObsPatternsTemplateItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsTemplateItem(clusteringInstructions string, createdAt time.Time, description string, evpQueryPrefix string, id string, name string, scope string, template string, timeFrameDurationMs int64) *LLMObsPatternsTemplateItem {
	this := LLMObsPatternsTemplateItem{}
	this.ClusteringInstructions = clusteringInstructions
	this.CreatedAt = createdAt
	this.Description = description
	this.EvpQueryPrefix = evpQueryPrefix
	this.Id = id
	this.Name = name
	this.Scope = scope
	this.Template = template
	this.TimeFrameDurationMs = timeFrameDurationMs
	return &this
}

// NewLLMObsPatternsTemplateItemWithDefaults instantiates a new LLMObsPatternsTemplateItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsTemplateItemWithDefaults() *LLMObsPatternsTemplateItem {
	this := LLMObsPatternsTemplateItem{}
	return &this
}

// GetClusteringInstructions returns the ClusteringInstructions field value.
func (o *LLMObsPatternsTemplateItem) GetClusteringInstructions() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusteringInstructions
}

// GetClusteringInstructionsOk returns a tuple with the ClusteringInstructions field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTemplateItem) GetClusteringInstructionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusteringInstructions, true
}

// SetClusteringInstructions sets field value.
func (o *LLMObsPatternsTemplateItem) SetClusteringInstructions(v string) {
	o.ClusteringInstructions = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsPatternsTemplateItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTemplateItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsPatternsTemplateItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value.
func (o *LLMObsPatternsTemplateItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTemplateItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *LLMObsPatternsTemplateItem) SetDescription(v string) {
	o.Description = v
}

// GetEvpQueryPrefix returns the EvpQueryPrefix field value.
func (o *LLMObsPatternsTemplateItem) GetEvpQueryPrefix() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvpQueryPrefix
}

// GetEvpQueryPrefixOk returns a tuple with the EvpQueryPrefix field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTemplateItem) GetEvpQueryPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvpQueryPrefix, true
}

// SetEvpQueryPrefix sets field value.
func (o *LLMObsPatternsTemplateItem) SetEvpQueryPrefix(v string) {
	o.EvpQueryPrefix = v
}

// GetId returns the Id field value.
func (o *LLMObsPatternsTemplateItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTemplateItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsPatternsTemplateItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *LLMObsPatternsTemplateItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTemplateItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsPatternsTemplateItem) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value.
func (o *LLMObsPatternsTemplateItem) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTemplateItem) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *LLMObsPatternsTemplateItem) SetScope(v string) {
	o.Scope = v
}

// GetTemplate returns the Template field value.
func (o *LLMObsPatternsTemplateItem) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTemplateItem) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value.
func (o *LLMObsPatternsTemplateItem) SetTemplate(v string) {
	o.Template = v
}

// GetTimeFrameDurationMs returns the TimeFrameDurationMs field value.
func (o *LLMObsPatternsTemplateItem) GetTimeFrameDurationMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TimeFrameDurationMs
}

// GetTimeFrameDurationMsOk returns a tuple with the TimeFrameDurationMs field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTemplateItem) GetTimeFrameDurationMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeFrameDurationMs, true
}

// SetTimeFrameDurationMs sets field value.
func (o *LLMObsPatternsTemplateItem) SetTimeFrameDurationMs(v int64) {
	o.TimeFrameDurationMs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsTemplateItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["clustering_instructions"] = o.ClusteringInstructions
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["description"] = o.Description
	toSerialize["evp_query_prefix"] = o.EvpQueryPrefix
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["scope"] = o.Scope
	toSerialize["template"] = o.Template
	toSerialize["time_frame_duration_ms"] = o.TimeFrameDurationMs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsTemplateItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusteringInstructions *string    `json:"clustering_instructions"`
		CreatedAt              *time.Time `json:"created_at"`
		Description            *string    `json:"description"`
		EvpQueryPrefix         *string    `json:"evp_query_prefix"`
		Id                     *string    `json:"id"`
		Name                   *string    `json:"name"`
		Scope                  *string    `json:"scope"`
		Template               *string    `json:"template"`
		TimeFrameDurationMs    *int64     `json:"time_frame_duration_ms"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClusteringInstructions == nil {
		return fmt.Errorf("required field clustering_instructions missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.EvpQueryPrefix == nil {
		return fmt.Errorf("required field evp_query_prefix missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.Template == nil {
		return fmt.Errorf("required field template missing")
	}
	if all.TimeFrameDurationMs == nil {
		return fmt.Errorf("required field time_frame_duration_ms missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"clustering_instructions", "created_at", "description", "evp_query_prefix", "id", "name", "scope", "template", "time_frame_duration_ms"})
	} else {
		return err
	}
	o.ClusteringInstructions = *all.ClusteringInstructions
	o.CreatedAt = *all.CreatedAt
	o.Description = *all.Description
	o.EvpQueryPrefix = *all.EvpQueryPrefix
	o.Id = *all.Id
	o.Name = *all.Name
	o.Scope = *all.Scope
	o.Template = *all.Template
	o.TimeFrameDurationMs = *all.TimeFrameDurationMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
