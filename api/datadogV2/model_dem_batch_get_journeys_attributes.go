// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemBatchGetJourneysAttributes Attributes for a batch get journeys request.
type DemBatchGetJourneysAttributes struct {
	// List of test suite IDs.
	TestSuiteIds []string `json:"test_suite_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemBatchGetJourneysAttributes instantiates a new DemBatchGetJourneysAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemBatchGetJourneysAttributes(testSuiteIds []string) *DemBatchGetJourneysAttributes {
	this := DemBatchGetJourneysAttributes{}
	this.TestSuiteIds = testSuiteIds
	return &this
}

// NewDemBatchGetJourneysAttributesWithDefaults instantiates a new DemBatchGetJourneysAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemBatchGetJourneysAttributesWithDefaults() *DemBatchGetJourneysAttributes {
	this := DemBatchGetJourneysAttributes{}
	return &this
}

// GetTestSuiteIds returns the TestSuiteIds field value.
func (o *DemBatchGetJourneysAttributes) GetTestSuiteIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TestSuiteIds
}

// GetTestSuiteIdsOk returns a tuple with the TestSuiteIds field value
// and a boolean to check if the value has been set.
func (o *DemBatchGetJourneysAttributes) GetTestSuiteIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestSuiteIds, true
}

// SetTestSuiteIds sets field value.
func (o *DemBatchGetJourneysAttributes) SetTestSuiteIds(v []string) {
	o.TestSuiteIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DemBatchGetJourneysAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["test_suite_ids"] = o.TestSuiteIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemBatchGetJourneysAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TestSuiteIds *[]string `json:"test_suite_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TestSuiteIds == nil {
		return fmt.Errorf("required field test_suite_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"test_suite_ids"})
	} else {
		return err
	}
	o.TestSuiteIds = *all.TestSuiteIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
