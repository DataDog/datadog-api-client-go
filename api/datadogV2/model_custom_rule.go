// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRule A custom static analysis rule within a ruleset, as supplied in a create or update
// request. Nested rules are sent flat, without a `data`/`type`/`attributes` envelope.
// `id` and `name` are client-supplied and must match each other. The remaining members
// are server-assigned and read-only; they are declared so that a ruleset previously
// read back can be supplied unchanged.
type CustomRule struct {
	// Creation timestamp
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Creator identifier
	CreatedBy *string `json:"created_by,omitempty"`
	// Rule identifier, which is the same as the rule name.
	Id string `json:"id"`
	// A revision of a custom static analysis rule as embedded in a rule supplied by a create
	// or update request. Nested revisions are sent flat, without a `data`/`type`/`attributes`
	// envelope. `id`, `version_id`, `checksum`, `created_at` and `created_by` are server-assigned
	// and read-only; they are declared so that a ruleset previously read back can be supplied
	// unchanged.
	LastRevision *CustomRuleRevisionInput `json:"last_revision,omitempty"`
	// Rule name
	Name string `json:"name"`
	// Revision history of the rule.
	Revisions []CustomRuleRevisionInput `json:"revisions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomRule instantiates a new CustomRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomRule(id string, name string) *CustomRule {
	this := CustomRule{}
	this.Id = id
	this.Name = name
	return &this
}

// NewCustomRuleWithDefaults instantiates a new CustomRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomRuleWithDefaults() *CustomRule {
	this := CustomRule{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomRule) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomRule) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *CustomRule) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRule) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CustomRule) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *CustomRule) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetId returns the Id field value.
func (o *CustomRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CustomRule) SetId(v string) {
	o.Id = v
}

// GetLastRevision returns the LastRevision field value if set, zero value otherwise.
func (o *CustomRule) GetLastRevision() CustomRuleRevisionInput {
	if o == nil || o.LastRevision == nil {
		var ret CustomRuleRevisionInput
		return ret
	}
	return *o.LastRevision
}

// GetLastRevisionOk returns a tuple with the LastRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRule) GetLastRevisionOk() (*CustomRuleRevisionInput, bool) {
	if o == nil || o.LastRevision == nil {
		return nil, false
	}
	return o.LastRevision, true
}

// HasLastRevision returns a boolean if a field has been set.
func (o *CustomRule) HasLastRevision() bool {
	return o != nil && o.LastRevision != nil
}

// SetLastRevision gets a reference to the given CustomRuleRevisionInput and assigns it to the LastRevision field.
func (o *CustomRule) SetLastRevision(v CustomRuleRevisionInput) {
	o.LastRevision = &v
}

// GetName returns the Name field value.
func (o *CustomRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CustomRule) SetName(v string) {
	o.Name = v
}

// GetRevisions returns the Revisions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomRule) GetRevisions() []CustomRuleRevisionInput {
	if o == nil {
		var ret []CustomRuleRevisionInput
		return ret
	}
	return o.Revisions
}

// GetRevisionsOk returns a tuple with the Revisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRule) GetRevisionsOk() (*[]CustomRuleRevisionInput, bool) {
	if o == nil || o.Revisions == nil {
		return nil, false
	}
	return &o.Revisions, true
}

// HasRevisions returns a boolean if a field has been set.
func (o *CustomRule) HasRevisions() bool {
	return o != nil && o.Revisions != nil
}

// SetRevisions gets a reference to the given []CustomRuleRevisionInput and assigns it to the Revisions field.
func (o *CustomRule) SetRevisions(v []CustomRuleRevisionInput) {
	o.Revisions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	toSerialize["id"] = o.Id
	if o.LastRevision != nil {
		toSerialize["last_revision"] = o.LastRevision
	}
	toSerialize["name"] = o.Name
	if o.Revisions != nil {
		toSerialize["revisions"] = o.Revisions
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt    *time.Time                `json:"created_at,omitempty"`
		CreatedBy    *string                   `json:"created_by,omitempty"`
		Id           *string                   `json:"id"`
		LastRevision *CustomRuleRevisionInput  `json:"last_revision,omitempty"`
		Name         *string                   `json:"name"`
		Revisions    []CustomRuleRevisionInput `json:"revisions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Id = *all.Id
	if all.LastRevision != nil && all.LastRevision.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastRevision = all.LastRevision
	o.Name = *all.Name
	o.Revisions = all.Revisions

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
