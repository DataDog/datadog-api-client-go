// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsFunnelSuggestionsFeature A feature detected in the application, with the journeys users take to use it.
type ProductAnalyticsFunnelSuggestionsFeature struct {
	// The suggested journey that users follow to use the feature. Contains one journey.
	CoreUserJourneys []ProductAnalyticsCoreUserJourney `json:"core_user_journeys,omitempty"`
	// Description of the feature.
	Description *string `json:"description,omitempty"`
	// ID of the feature.
	Id *string `json:"id,omitempty"`
	// Name of the feature.
	Name *string `json:"name,omitempty"`
	// Relevance score of the feature. Higher is more relevant.
	Score *int64 `json:"score,omitempty"`
	// Number of sessions that used the feature.
	Support *int64 `json:"support,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsFunnelSuggestionsFeature instantiates a new ProductAnalyticsFunnelSuggestionsFeature object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsFunnelSuggestionsFeature() *ProductAnalyticsFunnelSuggestionsFeature {
	this := ProductAnalyticsFunnelSuggestionsFeature{}
	return &this
}

// NewProductAnalyticsFunnelSuggestionsFeatureWithDefaults instantiates a new ProductAnalyticsFunnelSuggestionsFeature object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsFunnelSuggestionsFeatureWithDefaults() *ProductAnalyticsFunnelSuggestionsFeature {
	this := ProductAnalyticsFunnelSuggestionsFeature{}
	return &this
}

// GetCoreUserJourneys returns the CoreUserJourneys field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetCoreUserJourneys() []ProductAnalyticsCoreUserJourney {
	if o == nil || o.CoreUserJourneys == nil {
		var ret []ProductAnalyticsCoreUserJourney
		return ret
	}
	return o.CoreUserJourneys
}

// GetCoreUserJourneysOk returns a tuple with the CoreUserJourneys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetCoreUserJourneysOk() (*[]ProductAnalyticsCoreUserJourney, bool) {
	if o == nil || o.CoreUserJourneys == nil {
		return nil, false
	}
	return &o.CoreUserJourneys, true
}

// HasCoreUserJourneys returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) HasCoreUserJourneys() bool {
	return o != nil && o.CoreUserJourneys != nil
}

// SetCoreUserJourneys gets a reference to the given []ProductAnalyticsCoreUserJourney and assigns it to the CoreUserJourneys field.
func (o *ProductAnalyticsFunnelSuggestionsFeature) SetCoreUserJourneys(v []ProductAnalyticsCoreUserJourney) {
	o.CoreUserJourneys = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductAnalyticsFunnelSuggestionsFeature) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductAnalyticsFunnelSuggestionsFeature) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAnalyticsFunnelSuggestionsFeature) SetName(v string) {
	o.Name = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetScore() int64 {
	if o == nil || o.Score == nil {
		var ret int64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetScoreOk() (*int64, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) HasScore() bool {
	return o != nil && o.Score != nil
}

// SetScore gets a reference to the given int64 and assigns it to the Score field.
func (o *ProductAnalyticsFunnelSuggestionsFeature) SetScore(v int64) {
	o.Score = &v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetSupport() int64 {
	if o == nil || o.Support == nil {
		var ret int64
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) GetSupportOk() (*int64, bool) {
	if o == nil || o.Support == nil {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *ProductAnalyticsFunnelSuggestionsFeature) HasSupport() bool {
	return o != nil && o.Support != nil
}

// SetSupport gets a reference to the given int64 and assigns it to the Support field.
func (o *ProductAnalyticsFunnelSuggestionsFeature) SetSupport(v int64) {
	o.Support = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsFunnelSuggestionsFeature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CoreUserJourneys != nil {
		toSerialize["core_user_journeys"] = o.CoreUserJourneys
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.Support != nil {
		toSerialize["support"] = o.Support
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsFunnelSuggestionsFeature) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CoreUserJourneys []ProductAnalyticsCoreUserJourney `json:"core_user_journeys,omitempty"`
		Description      *string                           `json:"description,omitempty"`
		Id               *string                           `json:"id,omitempty"`
		Name             *string                           `json:"name,omitempty"`
		Score            *int64                            `json:"score,omitempty"`
		Support          *int64                            `json:"support,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"core_user_journeys", "description", "id", "name", "score", "support"})
	} else {
		return err
	}
	o.CoreUserJourneys = all.CoreUserJourneys
	o.Description = all.Description
	o.Id = all.Id
	o.Name = all.Name
	o.Score = all.Score
	o.Support = all.Support

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
