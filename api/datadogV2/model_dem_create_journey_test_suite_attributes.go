// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemCreateJourneyTestSuiteAttributes Attributes for creating a test suite for a DEM journey.
type DemCreateJourneyTestSuiteAttributes struct {
	// Whether to populate the test suite based on journey coverage data.
	IncludeTestsFromJourneyCoverage datadog.NullableBool `json:"include_tests_from_journey_coverage,omitempty"`
	// An optional custom name for the auto-created test suite.
	TestSuiteName datadog.NullableString `json:"test_suite_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemCreateJourneyTestSuiteAttributes instantiates a new DemCreateJourneyTestSuiteAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemCreateJourneyTestSuiteAttributes() *DemCreateJourneyTestSuiteAttributes {
	this := DemCreateJourneyTestSuiteAttributes{}
	return &this
}

// NewDemCreateJourneyTestSuiteAttributesWithDefaults instantiates a new DemCreateJourneyTestSuiteAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemCreateJourneyTestSuiteAttributesWithDefaults() *DemCreateJourneyTestSuiteAttributes {
	this := DemCreateJourneyTestSuiteAttributes{}
	return &this
}

// GetIncludeTestsFromJourneyCoverage returns the IncludeTestsFromJourneyCoverage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DemCreateJourneyTestSuiteAttributes) GetIncludeTestsFromJourneyCoverage() bool {
	if o == nil || o.IncludeTestsFromJourneyCoverage.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTestsFromJourneyCoverage.Get()
}

// GetIncludeTestsFromJourneyCoverageOk returns a tuple with the IncludeTestsFromJourneyCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DemCreateJourneyTestSuiteAttributes) GetIncludeTestsFromJourneyCoverageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeTestsFromJourneyCoverage.Get(), o.IncludeTestsFromJourneyCoverage.IsSet()
}

// HasIncludeTestsFromJourneyCoverage returns a boolean if a field has been set.
func (o *DemCreateJourneyTestSuiteAttributes) HasIncludeTestsFromJourneyCoverage() bool {
	return o != nil && o.IncludeTestsFromJourneyCoverage.IsSet()
}

// SetIncludeTestsFromJourneyCoverage gets a reference to the given datadog.NullableBool and assigns it to the IncludeTestsFromJourneyCoverage field.
func (o *DemCreateJourneyTestSuiteAttributes) SetIncludeTestsFromJourneyCoverage(v bool) {
	o.IncludeTestsFromJourneyCoverage.Set(&v)
}

// SetIncludeTestsFromJourneyCoverageNil sets the value for IncludeTestsFromJourneyCoverage to be an explicit nil.
func (o *DemCreateJourneyTestSuiteAttributes) SetIncludeTestsFromJourneyCoverageNil() {
	o.IncludeTestsFromJourneyCoverage.Set(nil)
}

// UnsetIncludeTestsFromJourneyCoverage ensures that no value is present for IncludeTestsFromJourneyCoverage, not even an explicit nil.
func (o *DemCreateJourneyTestSuiteAttributes) UnsetIncludeTestsFromJourneyCoverage() {
	o.IncludeTestsFromJourneyCoverage.Unset()
}

// GetTestSuiteName returns the TestSuiteName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DemCreateJourneyTestSuiteAttributes) GetTestSuiteName() string {
	if o == nil || o.TestSuiteName.Get() == nil {
		var ret string
		return ret
	}
	return *o.TestSuiteName.Get()
}

// GetTestSuiteNameOk returns a tuple with the TestSuiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DemCreateJourneyTestSuiteAttributes) GetTestSuiteNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestSuiteName.Get(), o.TestSuiteName.IsSet()
}

// HasTestSuiteName returns a boolean if a field has been set.
func (o *DemCreateJourneyTestSuiteAttributes) HasTestSuiteName() bool {
	return o != nil && o.TestSuiteName.IsSet()
}

// SetTestSuiteName gets a reference to the given datadog.NullableString and assigns it to the TestSuiteName field.
func (o *DemCreateJourneyTestSuiteAttributes) SetTestSuiteName(v string) {
	o.TestSuiteName.Set(&v)
}

// SetTestSuiteNameNil sets the value for TestSuiteName to be an explicit nil.
func (o *DemCreateJourneyTestSuiteAttributes) SetTestSuiteNameNil() {
	o.TestSuiteName.Set(nil)
}

// UnsetTestSuiteName ensures that no value is present for TestSuiteName, not even an explicit nil.
func (o *DemCreateJourneyTestSuiteAttributes) UnsetTestSuiteName() {
	o.TestSuiteName.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DemCreateJourneyTestSuiteAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IncludeTestsFromJourneyCoverage.IsSet() {
		toSerialize["include_tests_from_journey_coverage"] = o.IncludeTestsFromJourneyCoverage.Get()
	}
	if o.TestSuiteName.IsSet() {
		toSerialize["test_suite_name"] = o.TestSuiteName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemCreateJourneyTestSuiteAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeTestsFromJourneyCoverage datadog.NullableBool   `json:"include_tests_from_journey_coverage,omitempty"`
		TestSuiteName                   datadog.NullableString `json:"test_suite_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_tests_from_journey_coverage", "test_suite_name"})
	} else {
		return err
	}
	o.IncludeTestsFromJourneyCoverage = all.IncludeTestsFromJourneyCoverage
	o.TestSuiteName = all.TestSuiteName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
