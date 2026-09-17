// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRuleAttributes Attributes of a custom static analysis rule, including its most recent revision and revision history.
type CustomRuleAttributes struct {
	// Creation timestamp
	CreatedAt time.Time `json:"created_at"`
	// Creator identifier
	CreatedBy string `json:"created_by"`
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

// NewCustomRuleAttributes instantiates a new CustomRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomRuleAttributes(createdAt time.Time, createdBy string, lastRevision CustomRuleRevisionEmbedded, name string, revisions datadog.NullableList[CustomRuleRevisionEmbedded]) *CustomRuleAttributes {
	this := CustomRuleAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.LastRevision = lastRevision
	this.Name = name
	this.Revisions = revisions
	return &this
}

// NewCustomRuleAttributesWithDefaults instantiates a new CustomRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomRuleAttributesWithDefaults() *CustomRuleAttributes {
	this := CustomRuleAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CustomRuleAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomRuleAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CustomRuleAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *CustomRuleAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CustomRuleAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *CustomRuleAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetLastRevision returns the LastRevision field value.
func (o *CustomRuleAttributes) GetLastRevision() CustomRuleRevisionEmbedded {
	if o == nil {
		var ret CustomRuleRevisionEmbedded
		return ret
	}
	return o.LastRevision
}

// GetLastRevisionOk returns a tuple with the LastRevision field value
// and a boolean to check if the value has been set.
func (o *CustomRuleAttributes) GetLastRevisionOk() (*CustomRuleRevisionEmbedded, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastRevision, true
}

// SetLastRevision sets field value.
func (o *CustomRuleAttributes) SetLastRevision(v CustomRuleRevisionEmbedded) {
	o.LastRevision = v
}

// GetName returns the Name field value.
func (o *CustomRuleAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CustomRuleAttributes) SetName(v string) {
	o.Name = v
}

// GetRevisions returns the Revisions field value.
// If the value is explicit nil, the zero value for []CustomRuleRevisionEmbedded will be returned.
func (o *CustomRuleAttributes) GetRevisions() []CustomRuleRevisionEmbedded {
	if o == nil {
		var ret []CustomRuleRevisionEmbedded
		return ret
	}
	return *o.Revisions.Get()
}

// GetRevisionsOk returns a tuple with the Revisions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleAttributes) GetRevisionsOk() (*[]CustomRuleRevisionEmbedded, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revisions.Get(), o.Revisions.IsSet()
}

// SetRevisions sets field value.
func (o *CustomRuleAttributes) SetRevisions(v []CustomRuleRevisionEmbedded) {
	o.Revisions.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomRuleAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["last_revision"] = o.LastRevision
	toSerialize["name"] = o.Name
	toSerialize["revisions"] = o.Revisions.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt    *time.Time                                       `json:"created_at"`
		CreatedBy    *string                                          `json:"created_by"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "last_revision", "name", "revisions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
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
