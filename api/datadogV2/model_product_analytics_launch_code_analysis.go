// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchCodeAnalysis Latest inspection of the repository for the tracking plan instrumentation. A new analysis replaces the stored
// one.
// Events and properties that it does not list are not verified.
type ProductAnalyticsLaunchCodeAnalysis struct {
	// Time of the analysis. Set by the server.
	AnalyzedAt *time.Time `json:"analyzed_at,omitempty"`
	// Results for the events of the tracking plan. List each event once.
	Events []ProductAnalyticsLaunchEventFinding `json:"events,omitempty"`
	// Code revision that was inspected, such as a commit SHA.
	Revision *string `json:"revision,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchCodeAnalysis instantiates a new ProductAnalyticsLaunchCodeAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchCodeAnalysis() *ProductAnalyticsLaunchCodeAnalysis {
	this := ProductAnalyticsLaunchCodeAnalysis{}
	return &this
}

// NewProductAnalyticsLaunchCodeAnalysisWithDefaults instantiates a new ProductAnalyticsLaunchCodeAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchCodeAnalysisWithDefaults() *ProductAnalyticsLaunchCodeAnalysis {
	this := ProductAnalyticsLaunchCodeAnalysis{}
	return &this
}

// GetAnalyzedAt returns the AnalyzedAt field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchCodeAnalysis) GetAnalyzedAt() time.Time {
	if o == nil || o.AnalyzedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AnalyzedAt
}

// GetAnalyzedAtOk returns a tuple with the AnalyzedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchCodeAnalysis) GetAnalyzedAtOk() (*time.Time, bool) {
	if o == nil || o.AnalyzedAt == nil {
		return nil, false
	}
	return o.AnalyzedAt, true
}

// HasAnalyzedAt returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchCodeAnalysis) HasAnalyzedAt() bool {
	return o != nil && o.AnalyzedAt != nil
}

// SetAnalyzedAt gets a reference to the given time.Time and assigns it to the AnalyzedAt field.
func (o *ProductAnalyticsLaunchCodeAnalysis) SetAnalyzedAt(v time.Time) {
	o.AnalyzedAt = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchCodeAnalysis) GetEvents() []ProductAnalyticsLaunchEventFinding {
	if o == nil || o.Events == nil {
		var ret []ProductAnalyticsLaunchEventFinding
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchCodeAnalysis) GetEventsOk() (*[]ProductAnalyticsLaunchEventFinding, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchCodeAnalysis) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []ProductAnalyticsLaunchEventFinding and assigns it to the Events field.
func (o *ProductAnalyticsLaunchCodeAnalysis) SetEvents(v []ProductAnalyticsLaunchEventFinding) {
	o.Events = v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchCodeAnalysis) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchCodeAnalysis) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchCodeAnalysis) HasRevision() bool {
	return o != nil && o.Revision != nil
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *ProductAnalyticsLaunchCodeAnalysis) SetRevision(v string) {
	o.Revision = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchCodeAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnalyzedAt != nil {
		if o.AnalyzedAt.Nanosecond() == 0 {
			toSerialize["analyzed_at"] = o.AnalyzedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["analyzed_at"] = o.AnalyzedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchCodeAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnalyzedAt *time.Time                           `json:"analyzed_at,omitempty"`
		Events     []ProductAnalyticsLaunchEventFinding `json:"events,omitempty"`
		Revision   *string                              `json:"revision,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"analyzed_at", "events", "revision"})
	} else {
		return err
	}
	o.AnalyzedAt = all.AnalyzedAt
	o.Events = all.Events
	o.Revision = all.Revision

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
