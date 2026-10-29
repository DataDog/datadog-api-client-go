// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JourneyMapExpandedPrefix A path expanded to show the events that follow it.
type JourneyMapExpandedPrefix struct {
	// Events of the expanded path, in order.
	Events []string `json:"events"`
	// Maximum number of paths shown after the expanded path.
	PathLimit *int64 `json:"path_limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewJourneyMapExpandedPrefix instantiates a new JourneyMapExpandedPrefix object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJourneyMapExpandedPrefix(events []string) *JourneyMapExpandedPrefix {
	this := JourneyMapExpandedPrefix{}
	this.Events = events
	return &this
}

// NewJourneyMapExpandedPrefixWithDefaults instantiates a new JourneyMapExpandedPrefix object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJourneyMapExpandedPrefixWithDefaults() *JourneyMapExpandedPrefix {
	this := JourneyMapExpandedPrefix{}
	return &this
}

// GetEvents returns the Events field value.
func (o *JourneyMapExpandedPrefix) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *JourneyMapExpandedPrefix) GetEventsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value.
func (o *JourneyMapExpandedPrefix) SetEvents(v []string) {
	o.Events = v
}

// GetPathLimit returns the PathLimit field value if set, zero value otherwise.
func (o *JourneyMapExpandedPrefix) GetPathLimit() int64 {
	if o == nil || o.PathLimit == nil {
		var ret int64
		return ret
	}
	return *o.PathLimit
}

// GetPathLimitOk returns a tuple with the PathLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapExpandedPrefix) GetPathLimitOk() (*int64, bool) {
	if o == nil || o.PathLimit == nil {
		return nil, false
	}
	return o.PathLimit, true
}

// HasPathLimit returns a boolean if a field has been set.
func (o *JourneyMapExpandedPrefix) HasPathLimit() bool {
	return o != nil && o.PathLimit != nil
}

// SetPathLimit gets a reference to the given int64 and assigns it to the PathLimit field.
func (o *JourneyMapExpandedPrefix) SetPathLimit(v int64) {
	o.PathLimit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o JourneyMapExpandedPrefix) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["events"] = o.Events
	if o.PathLimit != nil {
		toSerialize["path_limit"] = o.PathLimit
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JourneyMapExpandedPrefix) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Events    *[]string `json:"events"`
		PathLimit *int64    `json:"path_limit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Events == nil {
		return fmt.Errorf("required field events missing")
	}
	o.Events = *all.Events
	o.PathLimit = all.PathLimit

	return nil
}
