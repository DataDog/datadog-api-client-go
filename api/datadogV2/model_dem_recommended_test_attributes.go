// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemRecommendedTestAttributes Attributes of an AI-recommended synthetic test for a DEM journey.
type DemRecommendedTestAttributes struct {
	// The browser test configuration that can be used to create the recommended test.
	Config map[string]interface{} `json:"config"`
	// The time when the recommendation was generated.
	CreatedAt time.Time `json:"created_at"`
	// The display name of the recommended test.
	Name string `json:"name"`
	// The identifier of the validating sample run, when available.
	ResultId *string `json:"result_id,omitempty"`
	// The RUM session identifier for the validating sample run, when available.
	SessionId *string `json:"session_id,omitempty"`
	// The pipeline that produced the recommendation.
	Source string `json:"source"`
	// The type of synthetic test.
	Type string `json:"type"`
	// The variant associated with the recommendation, when applicable.
	VariantId *string `json:"variant_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemRecommendedTestAttributes instantiates a new DemRecommendedTestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemRecommendedTestAttributes(config map[string]interface{}, createdAt time.Time, name string, source string, typeVar string) *DemRecommendedTestAttributes {
	this := DemRecommendedTestAttributes{}
	this.Config = config
	this.CreatedAt = createdAt
	this.Name = name
	this.Source = source
	this.Type = typeVar
	return &this
}

// NewDemRecommendedTestAttributesWithDefaults instantiates a new DemRecommendedTestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemRecommendedTestAttributesWithDefaults() *DemRecommendedTestAttributes {
	this := DemRecommendedTestAttributes{}
	return &this
}

// GetConfig returns the Config field value.
func (o *DemRecommendedTestAttributes) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *DemRecommendedTestAttributes) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value.
func (o *DemRecommendedTestAttributes) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DemRecommendedTestAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DemRecommendedTestAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DemRecommendedTestAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetName returns the Name field value.
func (o *DemRecommendedTestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DemRecommendedTestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DemRecommendedTestAttributes) SetName(v string) {
	o.Name = v
}

// GetResultId returns the ResultId field value if set, zero value otherwise.
func (o *DemRecommendedTestAttributes) GetResultId() string {
	if o == nil || o.ResultId == nil {
		var ret string
		return ret
	}
	return *o.ResultId
}

// GetResultIdOk returns a tuple with the ResultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemRecommendedTestAttributes) GetResultIdOk() (*string, bool) {
	if o == nil || o.ResultId == nil {
		return nil, false
	}
	return o.ResultId, true
}

// HasResultId returns a boolean if a field has been set.
func (o *DemRecommendedTestAttributes) HasResultId() bool {
	return o != nil && o.ResultId != nil
}

// SetResultId gets a reference to the given string and assigns it to the ResultId field.
func (o *DemRecommendedTestAttributes) SetResultId(v string) {
	o.ResultId = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *DemRecommendedTestAttributes) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemRecommendedTestAttributes) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *DemRecommendedTestAttributes) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *DemRecommendedTestAttributes) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSource returns the Source field value.
func (o *DemRecommendedTestAttributes) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DemRecommendedTestAttributes) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *DemRecommendedTestAttributes) SetSource(v string) {
	o.Source = v
}

// GetType returns the Type field value.
func (o *DemRecommendedTestAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DemRecommendedTestAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DemRecommendedTestAttributes) SetType(v string) {
	o.Type = v
}

// GetVariantId returns the VariantId field value if set, zero value otherwise.
func (o *DemRecommendedTestAttributes) GetVariantId() string {
	if o == nil || o.VariantId == nil {
		var ret string
		return ret
	}
	return *o.VariantId
}

// GetVariantIdOk returns a tuple with the VariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemRecommendedTestAttributes) GetVariantIdOk() (*string, bool) {
	if o == nil || o.VariantId == nil {
		return nil, false
	}
	return o.VariantId, true
}

// HasVariantId returns a boolean if a field has been set.
func (o *DemRecommendedTestAttributes) HasVariantId() bool {
	return o != nil && o.VariantId != nil
}

// SetVariantId gets a reference to the given string and assigns it to the VariantId field.
func (o *DemRecommendedTestAttributes) SetVariantId(v string) {
	o.VariantId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DemRecommendedTestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["config"] = o.Config
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	if o.ResultId != nil {
		toSerialize["result_id"] = o.ResultId
	}
	if o.SessionId != nil {
		toSerialize["session_id"] = o.SessionId
	}
	toSerialize["source"] = o.Source
	toSerialize["type"] = o.Type
	if o.VariantId != nil {
		toSerialize["variant_id"] = o.VariantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemRecommendedTestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config    *map[string]interface{} `json:"config"`
		CreatedAt *time.Time              `json:"created_at"`
		Name      *string                 `json:"name"`
		ResultId  *string                 `json:"result_id,omitempty"`
		SessionId *string                 `json:"session_id,omitempty"`
		Source    *string                 `json:"source"`
		Type      *string                 `json:"type"`
		VariantId *string                 `json:"variant_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Config == nil {
		return fmt.Errorf("required field config missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config", "created_at", "name", "result_id", "session_id", "source", "type", "variant_id"})
	} else {
		return err
	}
	o.Config = *all.Config
	o.CreatedAt = *all.CreatedAt
	o.Name = *all.Name
	o.ResultId = all.ResultId
	o.SessionId = all.SessionId
	o.Source = *all.Source
	o.Type = *all.Type
	o.VariantId = all.VariantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
