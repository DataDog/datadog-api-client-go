// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRulesetRuleEmbedded A custom static analysis rule as embedded in the rules list of a ruleset response.
type CustomRulesetRuleEmbedded struct {
	// Creation timestamp
	CreatedAt time.Time `json:"created_at"`
	// Creator identifier
	CreatedBy string `json:"created_by"`
	// Rule identifier, which is the same as the rule name.
	Id string `json:"id"`
	// A revision of a custom static analysis rule as embedded in a rule or ruleset response.
	LastRevision CustomRuleRevisionEmbedded `json:"last_revision"`
	// Rule name
	Name string `json:"name"`
	// Revision history of the rule.
	Revisions datadog.NullableList[CustomRuleRevisionEmbedded] `json:"revisions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomRulesetRuleEmbedded instantiates a new CustomRulesetRuleEmbedded object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomRulesetRuleEmbedded(createdAt time.Time, createdBy string, id string, lastRevision CustomRuleRevisionEmbedded, name string, revisions datadog.NullableList[CustomRuleRevisionEmbedded]) *CustomRulesetRuleEmbedded {
	this := CustomRulesetRuleEmbedded{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Id = id
	this.LastRevision = lastRevision
	this.Name = name
	this.Revisions = revisions
	return &this
}

// NewCustomRulesetRuleEmbeddedWithDefaults instantiates a new CustomRulesetRuleEmbedded object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomRulesetRuleEmbeddedWithDefaults() *CustomRulesetRuleEmbedded {
	this := CustomRulesetRuleEmbedded{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CustomRulesetRuleEmbedded) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomRulesetRuleEmbedded) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CustomRulesetRuleEmbedded) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *CustomRulesetRuleEmbedded) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CustomRulesetRuleEmbedded) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *CustomRulesetRuleEmbedded) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetId returns the Id field value.
func (o *CustomRulesetRuleEmbedded) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomRulesetRuleEmbedded) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CustomRulesetRuleEmbedded) SetId(v string) {
	o.Id = v
}

// GetLastRevision returns the LastRevision field value.
func (o *CustomRulesetRuleEmbedded) GetLastRevision() CustomRuleRevisionEmbedded {
	if o == nil {
		var ret CustomRuleRevisionEmbedded
		return ret
	}
	return o.LastRevision
}

// GetLastRevisionOk returns a tuple with the LastRevision field value
// and a boolean to check if the value has been set.
func (o *CustomRulesetRuleEmbedded) GetLastRevisionOk() (*CustomRuleRevisionEmbedded, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastRevision, true
}

// SetLastRevision sets field value.
func (o *CustomRulesetRuleEmbedded) SetLastRevision(v CustomRuleRevisionEmbedded) {
	o.LastRevision = v
}

// GetName returns the Name field value.
func (o *CustomRulesetRuleEmbedded) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomRulesetRuleEmbedded) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CustomRulesetRuleEmbedded) SetName(v string) {
	o.Name = v
}

// GetRevisions returns the Revisions field value.
// If the value is explicit nil, the zero value for []CustomRuleRevisionEmbedded will be returned.
func (o *CustomRulesetRuleEmbedded) GetRevisions() []CustomRuleRevisionEmbedded {
	if o == nil {
		var ret []CustomRuleRevisionEmbedded
		return ret
	}
	return *o.Revisions.Get()
}

// GetRevisionsOk returns a tuple with the Revisions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRulesetRuleEmbedded) GetRevisionsOk() (*[]CustomRuleRevisionEmbedded, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revisions.Get(), o.Revisions.IsSet()
}

// SetRevisions sets field value.
func (o *CustomRulesetRuleEmbedded) SetRevisions(v []CustomRuleRevisionEmbedded) {
	o.Revisions.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomRulesetRuleEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["id"] = o.Id
	toSerialize["last_revision"] = o.LastRevision
	toSerialize["name"] = o.Name
	toSerialize["revisions"] = o.Revisions.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomRulesetRuleEmbedded) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt    *time.Time                                       `json:"created_at"`
		CreatedBy    *string                                          `json:"created_by"`
		Id           *string                                          `json:"id"`
		LastRevision *CustomRuleRevisionEmbedded                      `json:"last_revision"`
		Name         *string                                          `json:"name"`
		Revisions    datadog.NullableList[CustomRuleRevisionEmbedded] `json:"revisions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.LastRevision == nil {
		return fmt.Errorf("required field last_revision missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if !all.Revisions.IsSet() {
		return fmt.Errorf("required field revisions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "id", "last_revision", "name", "revisions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Id = *all.Id
	if all.LastRevision.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastRevision = *all.LastRevision
	o.Name = *all.Name
	o.Revisions = all.Revisions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
