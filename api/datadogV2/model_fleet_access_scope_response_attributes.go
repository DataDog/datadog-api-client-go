// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAccessScopeResponseAttributes Fleet access scope attributes and audit information.
type FleetAccessScopeResponseAttributes struct {
	// RFC3339 timestamp when the scope was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Handle of the user who created the scope.
	Creator *string `json:"creator,omitempty"`
	// Human-readable description of the access scope. Omission clears the description on update.
	Description *string `json:"description,omitempty"`
	// Fleet actions authorized on agents matching the access scope query.
	Features *FleetAccessScopeFeatures `json:"features,omitempty"`
	// Agent query selecting hosts covered by the scope. An empty query matches all hosts. Omission clears the query on update.
	Query *string `json:"query,omitempty"`
	// RFC3339 timestamp when the scope was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// Handle of the user who last updated the scope.
	Updater *string `json:"updater,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAccessScopeResponseAttributes instantiates a new FleetAccessScopeResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAccessScopeResponseAttributes() *FleetAccessScopeResponseAttributes {
	this := FleetAccessScopeResponseAttributes{}
	var description string = ""
	this.Description = &description
	var query string = ""
	this.Query = &query
	return &this
}

// NewFleetAccessScopeResponseAttributesWithDefaults instantiates a new FleetAccessScopeResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAccessScopeResponseAttributesWithDefaults() *FleetAccessScopeResponseAttributes {
	this := FleetAccessScopeResponseAttributes{}
	var description string = ""
	this.Description = &description
	var query string = ""
	this.Query = &query
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FleetAccessScopeResponseAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeResponseAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FleetAccessScopeResponseAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *FleetAccessScopeResponseAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *FleetAccessScopeResponseAttributes) GetCreator() string {
	if o == nil || o.Creator == nil {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeResponseAttributes) GetCreatorOk() (*string, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *FleetAccessScopeResponseAttributes) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *FleetAccessScopeResponseAttributes) SetCreator(v string) {
	o.Creator = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FleetAccessScopeResponseAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FleetAccessScopeResponseAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FleetAccessScopeResponseAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *FleetAccessScopeResponseAttributes) GetFeatures() FleetAccessScopeFeatures {
	if o == nil || o.Features == nil {
		var ret FleetAccessScopeFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeResponseAttributes) GetFeaturesOk() (*FleetAccessScopeFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *FleetAccessScopeResponseAttributes) HasFeatures() bool {
	return o != nil && o.Features != nil
}

// SetFeatures gets a reference to the given FleetAccessScopeFeatures and assigns it to the Features field.
func (o *FleetAccessScopeResponseAttributes) SetFeatures(v FleetAccessScopeFeatures) {
	o.Features = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *FleetAccessScopeResponseAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeResponseAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *FleetAccessScopeResponseAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *FleetAccessScopeResponseAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FleetAccessScopeResponseAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeResponseAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FleetAccessScopeResponseAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *FleetAccessScopeResponseAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdater returns the Updater field value if set, zero value otherwise.
func (o *FleetAccessScopeResponseAttributes) GetUpdater() string {
	if o == nil || o.Updater == nil {
		var ret string
		return ret
	}
	return *o.Updater
}

// GetUpdaterOk returns a tuple with the Updater field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeResponseAttributes) GetUpdaterOk() (*string, bool) {
	if o == nil || o.Updater == nil {
		return nil, false
	}
	return o.Updater, true
}

// HasUpdater returns a boolean if a field has been set.
func (o *FleetAccessScopeResponseAttributes) HasUpdater() bool {
	return o != nil && o.Updater != nil
}

// SetUpdater gets a reference to the given string and assigns it to the Updater field.
func (o *FleetAccessScopeResponseAttributes) SetUpdater(v string) {
	o.Updater = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAccessScopeResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Updater != nil {
		toSerialize["updater"] = o.Updater
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetAccessScopeResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *string                   `json:"created_at,omitempty"`
		Creator     *string                   `json:"creator,omitempty"`
		Description *string                   `json:"description,omitempty"`
		Features    *FleetAccessScopeFeatures `json:"features,omitempty"`
		Query       *string                   `json:"query,omitempty"`
		UpdatedAt   *string                   `json:"updated_at,omitempty"`
		Updater     *string                   `json:"updater,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "creator", "description", "features", "query", "updated_at", "updater"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.Creator = all.Creator
	o.Description = all.Description
	if all.Features != nil && all.Features.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Features = all.Features
	o.Query = all.Query
	o.UpdatedAt = all.UpdatedAt
	o.Updater = all.Updater

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
