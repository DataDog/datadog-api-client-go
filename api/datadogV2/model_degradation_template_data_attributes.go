// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DegradationTemplateDataAttributes The attributes of a degradation template.
type DegradationTemplateDataAttributes struct {
	// The components affected by a degradation created from this template.
	ComponentsAffected []DegradationTemplateDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
	// Timestamp of when the degradation template was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The title used for a degradation created from this template.
	DegradationTitle *string `json:"degradation_title,omitempty"`
	// Timestamp of when the degradation template was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The name of the degradation template.
	Name *string `json:"name,omitempty"`
	// The pre-filled updates for a degradation created from this template.
	Updates []DegradationTemplateDataAttributesUpdatesItems `json:"updates,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDegradationTemplateDataAttributes instantiates a new DegradationTemplateDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDegradationTemplateDataAttributes() *DegradationTemplateDataAttributes {
	this := DegradationTemplateDataAttributes{}
	return &this
}

// NewDegradationTemplateDataAttributesWithDefaults instantiates a new DegradationTemplateDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDegradationTemplateDataAttributesWithDefaults() *DegradationTemplateDataAttributes {
	this := DegradationTemplateDataAttributes{}
	return &this
}

// GetComponentsAffected returns the ComponentsAffected field value if set, zero value otherwise.
func (o *DegradationTemplateDataAttributes) GetComponentsAffected() []DegradationTemplateDataAttributesComponentsAffectedItems {
	if o == nil || o.ComponentsAffected == nil {
		var ret []DegradationTemplateDataAttributesComponentsAffectedItems
		return ret
	}
	return o.ComponentsAffected
}

// GetComponentsAffectedOk returns a tuple with the ComponentsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationTemplateDataAttributes) GetComponentsAffectedOk() (*[]DegradationTemplateDataAttributesComponentsAffectedItems, bool) {
	if o == nil || o.ComponentsAffected == nil {
		return nil, false
	}
	return &o.ComponentsAffected, true
}

// HasComponentsAffected returns a boolean if a field has been set.
func (o *DegradationTemplateDataAttributes) HasComponentsAffected() bool {
	return o != nil && o.ComponentsAffected != nil
}

// SetComponentsAffected gets a reference to the given []DegradationTemplateDataAttributesComponentsAffectedItems and assigns it to the ComponentsAffected field.
func (o *DegradationTemplateDataAttributes) SetComponentsAffected(v []DegradationTemplateDataAttributesComponentsAffectedItems) {
	o.ComponentsAffected = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DegradationTemplateDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationTemplateDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DegradationTemplateDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DegradationTemplateDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDegradationTitle returns the DegradationTitle field value if set, zero value otherwise.
func (o *DegradationTemplateDataAttributes) GetDegradationTitle() string {
	if o == nil || o.DegradationTitle == nil {
		var ret string
		return ret
	}
	return *o.DegradationTitle
}

// GetDegradationTitleOk returns a tuple with the DegradationTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationTemplateDataAttributes) GetDegradationTitleOk() (*string, bool) {
	if o == nil || o.DegradationTitle == nil {
		return nil, false
	}
	return o.DegradationTitle, true
}

// HasDegradationTitle returns a boolean if a field has been set.
func (o *DegradationTemplateDataAttributes) HasDegradationTitle() bool {
	return o != nil && o.DegradationTitle != nil
}

// SetDegradationTitle gets a reference to the given string and assigns it to the DegradationTitle field.
func (o *DegradationTemplateDataAttributes) SetDegradationTitle(v string) {
	o.DegradationTitle = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *DegradationTemplateDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationTemplateDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *DegradationTemplateDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *DegradationTemplateDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DegradationTemplateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationTemplateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DegradationTemplateDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DegradationTemplateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetUpdates returns the Updates field value if set, zero value otherwise.
func (o *DegradationTemplateDataAttributes) GetUpdates() []DegradationTemplateDataAttributesUpdatesItems {
	if o == nil || o.Updates == nil {
		var ret []DegradationTemplateDataAttributesUpdatesItems
		return ret
	}
	return o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationTemplateDataAttributes) GetUpdatesOk() (*[]DegradationTemplateDataAttributesUpdatesItems, bool) {
	if o == nil || o.Updates == nil {
		return nil, false
	}
	return &o.Updates, true
}

// HasUpdates returns a boolean if a field has been set.
func (o *DegradationTemplateDataAttributes) HasUpdates() bool {
	return o != nil && o.Updates != nil
}

// SetUpdates gets a reference to the given []DegradationTemplateDataAttributesUpdatesItems and assigns it to the Updates field.
func (o *DegradationTemplateDataAttributes) SetUpdates(v []DegradationTemplateDataAttributesUpdatesItems) {
	o.Updates = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DegradationTemplateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComponentsAffected != nil {
		toSerialize["components_affected"] = o.ComponentsAffected
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DegradationTitle != nil {
		toSerialize["degradation_title"] = o.DegradationTitle
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Updates != nil {
		toSerialize["updates"] = o.Updates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DegradationTemplateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentsAffected []DegradationTemplateDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
		CreatedAt          *time.Time                                                 `json:"created_at,omitempty"`
		DegradationTitle   *string                                                    `json:"degradation_title,omitempty"`
		ModifiedAt         *time.Time                                                 `json:"modified_at,omitempty"`
		Name               *string                                                    `json:"name,omitempty"`
		Updates            []DegradationTemplateDataAttributesUpdatesItems            `json:"updates,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components_affected", "created_at", "degradation_title", "modified_at", "name", "updates"})
	} else {
		return err
	}
	o.ComponentsAffected = all.ComponentsAffected
	o.CreatedAt = all.CreatedAt
	o.DegradationTitle = all.DegradationTitle
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.Updates = all.Updates

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
