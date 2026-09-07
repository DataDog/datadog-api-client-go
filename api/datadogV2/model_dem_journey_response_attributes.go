// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemJourneyResponseAttributes Attributes returned in a DEM journey response.
type DemJourneyResponseAttributes struct {
	// The timestamp when the journey was created.
	CreatedAt time.Time `json:"created_at"`
	// A Datadog user associated with a DEM operation.
	CreatedBy DemUser `json:"created_by"`
	// An optional human-readable description of the journey.
	Description datadog.NullableString `json:"description"`
	// The RUM definition for a DEM journey.
	JourneyRum DemJourneyRum `json:"journey_rum"`
	// The name of the DEM journey.
	Name string `json:"name"`
	// The organization ID that owns this journey.
	OrgId int64 `json:"org_id"`
	// List of tags associated with a DEM resource.
	Tags []string `json:"tags"`
	// A test suite associated with a DEM resource.
	TestSuite DemTestSuiteNested `json:"test_suite"`
	// The timestamp when the journey was last updated.
	UpdatedAt datadog.NullableTime `json:"updated_at"`
	// A Datadog user associated with a DEM operation.
	UpdatedBy DemUser `json:"updated_by"`
	// List of variants associated with a DEM journey.
	Variants []DemVariant `json:"variants"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemJourneyResponseAttributes instantiates a new DemJourneyResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemJourneyResponseAttributes(createdAt time.Time, createdBy DemUser, description datadog.NullableString, journeyRum DemJourneyRum, name string, orgId int64, tags []string, testSuite DemTestSuiteNested, updatedAt datadog.NullableTime, updatedBy DemUser, variants []DemVariant) *DemJourneyResponseAttributes {
	this := DemJourneyResponseAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Description = description
	this.JourneyRum = journeyRum
	this.Name = name
	this.OrgId = orgId
	this.Tags = tags
	this.TestSuite = testSuite
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	this.Variants = variants
	return &this
}

// NewDemJourneyResponseAttributesWithDefaults instantiates a new DemJourneyResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemJourneyResponseAttributesWithDefaults() *DemJourneyResponseAttributes {
	this := DemJourneyResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DemJourneyResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DemJourneyResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *DemJourneyResponseAttributes) GetCreatedBy() DemUser {
	if o == nil {
		var ret DemUser
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseAttributes) GetCreatedByOk() (*DemUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *DemJourneyResponseAttributes) SetCreatedBy(v DemUser) {
	o.CreatedBy = v
}

// GetDescription returns the Description field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *DemJourneyResponseAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DemJourneyResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value.
func (o *DemJourneyResponseAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetJourneyRum returns the JourneyRum field value.
func (o *DemJourneyResponseAttributes) GetJourneyRum() DemJourneyRum {
	if o == nil {
		var ret DemJourneyRum
		return ret
	}
	return o.JourneyRum
}

// GetJourneyRumOk returns a tuple with the JourneyRum field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseAttributes) GetJourneyRumOk() (*DemJourneyRum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JourneyRum, true
}

// SetJourneyRum sets field value.
func (o *DemJourneyResponseAttributes) SetJourneyRum(v DemJourneyRum) {
	o.JourneyRum = v
}

// GetName returns the Name field value.
func (o *DemJourneyResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DemJourneyResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *DemJourneyResponseAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *DemJourneyResponseAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetTags returns the Tags field value.
func (o *DemJourneyResponseAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *DemJourneyResponseAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTestSuite returns the TestSuite field value.
func (o *DemJourneyResponseAttributes) GetTestSuite() DemTestSuiteNested {
	if o == nil {
		var ret DemTestSuiteNested
		return ret
	}
	return o.TestSuite
}

// GetTestSuiteOk returns a tuple with the TestSuite field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseAttributes) GetTestSuiteOk() (*DemTestSuiteNested, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestSuite, true
}

// SetTestSuite sets field value.
func (o *DemJourneyResponseAttributes) SetTestSuite(v DemTestSuiteNested) {
	o.TestSuite = v
}

// GetUpdatedAt returns the UpdatedAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *DemJourneyResponseAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DemJourneyResponseAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value.
func (o *DemJourneyResponseAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// GetUpdatedBy returns the UpdatedBy field value.
func (o *DemJourneyResponseAttributes) GetUpdatedBy() DemUser {
	if o == nil {
		var ret DemUser
		return ret
	}
	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseAttributes) GetUpdatedByOk() (*DemUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value.
func (o *DemJourneyResponseAttributes) SetUpdatedBy(v DemUser) {
	o.UpdatedBy = v
}

// GetVariants returns the Variants field value.
func (o *DemJourneyResponseAttributes) GetVariants() []DemVariant {
	if o == nil {
		var ret []DemVariant
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *DemJourneyResponseAttributes) GetVariantsOk() (*[]DemVariant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variants, true
}

// SetVariants sets field value.
func (o *DemJourneyResponseAttributes) SetVariants(v []DemVariant) {
	o.Variants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DemJourneyResponseAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["description"] = o.Description.Get()
	toSerialize["journey_rum"] = o.JourneyRum
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	toSerialize["tags"] = o.Tags
	toSerialize["test_suite"] = o.TestSuite
	toSerialize["updated_at"] = o.UpdatedAt.Get()
	toSerialize["updated_by"] = o.UpdatedBy
	toSerialize["variants"] = o.Variants

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemJourneyResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time             `json:"created_at"`
		CreatedBy   *DemUser               `json:"created_by"`
		Description datadog.NullableString `json:"description"`
		JourneyRum  *DemJourneyRum         `json:"journey_rum"`
		Name        *string                `json:"name"`
		OrgId       *int64                 `json:"org_id"`
		Tags        *[]string              `json:"tags"`
		TestSuite   *DemTestSuiteNested    `json:"test_suite"`
		UpdatedAt   datadog.NullableTime   `json:"updated_at"`
		UpdatedBy   *DemUser               `json:"updated_by"`
		Variants    *[]DemVariant          `json:"variants"`
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
	if !all.Description.IsSet() {
		return fmt.Errorf("required field description missing")
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
	if all.TestSuite == nil {
		return fmt.Errorf("required field test_suite missing")
	}
	if !all.UpdatedAt.IsSet() {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.UpdatedBy == nil {
		return fmt.Errorf("required field updated_by missing")
	}
	if all.Variants == nil {
		return fmt.Errorf("required field variants missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "description", "journey_rum", "name", "org_id", "tags", "test_suite", "updated_at", "updated_by", "variants"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = *all.CreatedBy
	o.Description = all.Description
	if all.JourneyRum.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JourneyRum = *all.JourneyRum
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	o.Tags = *all.Tags
	if all.TestSuite.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TestSuite = *all.TestSuite
	o.UpdatedAt = all.UpdatedAt
	if all.UpdatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UpdatedBy = *all.UpdatedBy
	o.Variants = *all.Variants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
