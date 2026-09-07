// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemInferredJourneyIgnoredAttributes Attributes of an ignored inferred journey.
type DemInferredJourneyIgnoredAttributes struct {
	// Timestamp when the inferred journey was first observed.
	CreatedAt time.Time `json:"created_at"`
	// An optional description of the inferred journey.
	Description datadog.NullableString `json:"description,omitempty"`
	// Timestamp when the inferred journey was ignored.
	IgnoredAt time.Time `json:"ignored_at"`
	// The RUM definition for a DEM journey.
	JourneyRum DemJourneyRum `json:"journey_rum"`
	// The name of the inferred journey.
	Name string `json:"name"`
	// The organization ID that owns this inferred journey.
	OrgId int64 `json:"org_id"`
	// List of tags associated with a DEM resource.
	Tags []string `json:"tags"`
	// A test suite associated with a DEM resource.
	TestSuite *DemTestSuiteNested `json:"test_suite,omitempty"`
	// List of variants associated with a DEM journey.
	Variants []DemVariant `json:"variants"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemInferredJourneyIgnoredAttributes instantiates a new DemInferredJourneyIgnoredAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemInferredJourneyIgnoredAttributes(createdAt time.Time, ignoredAt time.Time, journeyRum DemJourneyRum, name string, orgId int64, tags []string, variants []DemVariant) *DemInferredJourneyIgnoredAttributes {
	this := DemInferredJourneyIgnoredAttributes{}
	this.CreatedAt = createdAt
	this.IgnoredAt = ignoredAt
	this.JourneyRum = journeyRum
	this.Name = name
	this.OrgId = orgId
	this.Tags = tags
	this.Variants = variants
	return &this
}

// NewDemInferredJourneyIgnoredAttributesWithDefaults instantiates a new DemInferredJourneyIgnoredAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemInferredJourneyIgnoredAttributesWithDefaults() *DemInferredJourneyIgnoredAttributes {
	this := DemInferredJourneyIgnoredAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DemInferredJourneyIgnoredAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DemInferredJourneyIgnoredAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DemInferredJourneyIgnoredAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DemInferredJourneyIgnoredAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DemInferredJourneyIgnoredAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DemInferredJourneyIgnoredAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *DemInferredJourneyIgnoredAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *DemInferredJourneyIgnoredAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *DemInferredJourneyIgnoredAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetIgnoredAt returns the IgnoredAt field value.
func (o *DemInferredJourneyIgnoredAttributes) GetIgnoredAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.IgnoredAt
}

// GetIgnoredAtOk returns a tuple with the IgnoredAt field value
// and a boolean to check if the value has been set.
func (o *DemInferredJourneyIgnoredAttributes) GetIgnoredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgnoredAt, true
}

// SetIgnoredAt sets field value.
func (o *DemInferredJourneyIgnoredAttributes) SetIgnoredAt(v time.Time) {
	o.IgnoredAt = v
}

// GetJourneyRum returns the JourneyRum field value.
func (o *DemInferredJourneyIgnoredAttributes) GetJourneyRum() DemJourneyRum {
	if o == nil {
		var ret DemJourneyRum
		return ret
	}
	return o.JourneyRum
}

// GetJourneyRumOk returns a tuple with the JourneyRum field value
// and a boolean to check if the value has been set.
func (o *DemInferredJourneyIgnoredAttributes) GetJourneyRumOk() (*DemJourneyRum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JourneyRum, true
}

// SetJourneyRum sets field value.
func (o *DemInferredJourneyIgnoredAttributes) SetJourneyRum(v DemJourneyRum) {
	o.JourneyRum = v
}

// GetName returns the Name field value.
func (o *DemInferredJourneyIgnoredAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DemInferredJourneyIgnoredAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DemInferredJourneyIgnoredAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *DemInferredJourneyIgnoredAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *DemInferredJourneyIgnoredAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *DemInferredJourneyIgnoredAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetTags returns the Tags field value.
func (o *DemInferredJourneyIgnoredAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *DemInferredJourneyIgnoredAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *DemInferredJourneyIgnoredAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTestSuite returns the TestSuite field value if set, zero value otherwise.
func (o *DemInferredJourneyIgnoredAttributes) GetTestSuite() DemTestSuiteNested {
	if o == nil || o.TestSuite == nil {
		var ret DemTestSuiteNested
		return ret
	}
	return *o.TestSuite
}

// GetTestSuiteOk returns a tuple with the TestSuite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemInferredJourneyIgnoredAttributes) GetTestSuiteOk() (*DemTestSuiteNested, bool) {
	if o == nil || o.TestSuite == nil {
		return nil, false
	}
	return o.TestSuite, true
}

// HasTestSuite returns a boolean if a field has been set.
func (o *DemInferredJourneyIgnoredAttributes) HasTestSuite() bool {
	return o != nil && o.TestSuite != nil
}

// SetTestSuite gets a reference to the given DemTestSuiteNested and assigns it to the TestSuite field.
func (o *DemInferredJourneyIgnoredAttributes) SetTestSuite(v DemTestSuiteNested) {
	o.TestSuite = &v
}

// GetVariants returns the Variants field value.
func (o *DemInferredJourneyIgnoredAttributes) GetVariants() []DemVariant {
	if o == nil {
		var ret []DemVariant
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *DemInferredJourneyIgnoredAttributes) GetVariantsOk() (*[]DemVariant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variants, true
}

// SetVariants sets field value.
func (o *DemInferredJourneyIgnoredAttributes) SetVariants(v []DemVariant) {
	o.Variants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DemInferredJourneyIgnoredAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.IgnoredAt.Nanosecond() == 0 {
		toSerialize["ignored_at"] = o.IgnoredAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["ignored_at"] = o.IgnoredAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["journey_rum"] = o.JourneyRum
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	toSerialize["tags"] = o.Tags
	if o.TestSuite != nil {
		toSerialize["test_suite"] = o.TestSuite
	}
	toSerialize["variants"] = o.Variants

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemInferredJourneyIgnoredAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time             `json:"created_at"`
		Description datadog.NullableString `json:"description,omitempty"`
		IgnoredAt   *time.Time             `json:"ignored_at"`
		JourneyRum  *DemJourneyRum         `json:"journey_rum"`
		Name        *string                `json:"name"`
		OrgId       *int64                 `json:"org_id"`
		Tags        *[]string              `json:"tags"`
		TestSuite   *DemTestSuiteNested    `json:"test_suite,omitempty"`
		Variants    *[]DemVariant          `json:"variants"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.IgnoredAt == nil {
		return fmt.Errorf("required field ignored_at missing")
	}
	if all.JourneyRum == nil {
		return fmt.Errorf("required field journey_rum missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	if all.Variants == nil {
		return fmt.Errorf("required field variants missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "description", "ignored_at", "journey_rum", "name", "org_id", "tags", "test_suite", "variants"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.Description = all.Description
	o.IgnoredAt = *all.IgnoredAt
	if all.JourneyRum.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JourneyRum = *all.JourneyRum
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	o.Tags = *all.Tags
	if all.TestSuite != nil && all.TestSuite.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TestSuite = all.TestSuite
	o.Variants = *all.Variants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
