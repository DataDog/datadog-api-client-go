// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemJourneyTestSuiteResponseAttributes Attributes of a DEM journey test suite response.
type DemJourneyTestSuiteResponseAttributes struct {
	// The timestamp when the test suite was created.
	CreatedAt time.Time `json:"created_at"`
	// Test IDs omitted because the caller lacks read access.
	DroppedTestIds []string `json:"dropped_test_ids,omitempty"`
	// The name of the test suite.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemJourneyTestSuiteResponseAttributes instantiates a new DemJourneyTestSuiteResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemJourneyTestSuiteResponseAttributes(createdAt time.Time, name string) *DemJourneyTestSuiteResponseAttributes {
	this := DemJourneyTestSuiteResponseAttributes{}
	this.CreatedAt = createdAt
	this.Name = name
	return &this
}

// NewDemJourneyTestSuiteResponseAttributesWithDefaults instantiates a new DemJourneyTestSuiteResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemJourneyTestSuiteResponseAttributesWithDefaults() *DemJourneyTestSuiteResponseAttributes {
	this := DemJourneyTestSuiteResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DemJourneyTestSuiteResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DemJourneyTestSuiteResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DemJourneyTestSuiteResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDroppedTestIds returns the DroppedTestIds field value if set, zero value otherwise.
func (o *DemJourneyTestSuiteResponseAttributes) GetDroppedTestIds() []string {
	if o == nil || o.DroppedTestIds == nil {
		var ret []string
		return ret
	}
	return o.DroppedTestIds
}

// GetDroppedTestIdsOk returns a tuple with the DroppedTestIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemJourneyTestSuiteResponseAttributes) GetDroppedTestIdsOk() (*[]string, bool) {
	if o == nil || o.DroppedTestIds == nil {
		return nil, false
	}
	return &o.DroppedTestIds, true
}

// HasDroppedTestIds returns a boolean if a field has been set.
func (o *DemJourneyTestSuiteResponseAttributes) HasDroppedTestIds() bool {
	return o != nil && o.DroppedTestIds != nil
}

// SetDroppedTestIds gets a reference to the given []string and assigns it to the DroppedTestIds field.
func (o *DemJourneyTestSuiteResponseAttributes) SetDroppedTestIds(v []string) {
	o.DroppedTestIds = v
}

// GetName returns the Name field value.
func (o *DemJourneyTestSuiteResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DemJourneyTestSuiteResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DemJourneyTestSuiteResponseAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DemJourneyTestSuiteResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DroppedTestIds != nil {
		toSerialize["dropped_test_ids"] = o.DroppedTestIds
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemJourneyTestSuiteResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      *time.Time `json:"created_at"`
		DroppedTestIds []string   `json:"dropped_test_ids,omitempty"`
		Name           *string    `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "dropped_test_ids", "name"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.DroppedTestIds = all.DroppedTestIds
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
