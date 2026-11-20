// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AudienceFiltersSegmentsItems Test-only generated schema for the dashboard versioning proof of concept.
type AudienceFiltersSegmentsItems struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Name string `json:"name"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SegmentId string `json:"segment_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAudienceFiltersSegmentsItems instantiates a new AudienceFiltersSegmentsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAudienceFiltersSegmentsItems(name string, segmentId string) *AudienceFiltersSegmentsItems {
	this := AudienceFiltersSegmentsItems{}
	this.Name = name
	this.SegmentId = segmentId
	return &this
}

// NewAudienceFiltersSegmentsItemsWithDefaults instantiates a new AudienceFiltersSegmentsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAudienceFiltersSegmentsItemsWithDefaults() *AudienceFiltersSegmentsItems {
	this := AudienceFiltersSegmentsItems{}
	return &this
}

// GetName returns the Name field value.
func (o *AudienceFiltersSegmentsItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AudienceFiltersSegmentsItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AudienceFiltersSegmentsItems) SetName(v string) {
	o.Name = v
}

// GetSegmentId returns the SegmentId field value.
func (o *AudienceFiltersSegmentsItems) GetSegmentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value
// and a boolean to check if the value has been set.
func (o *AudienceFiltersSegmentsItems) GetSegmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentId, true
}

// SetSegmentId sets field value.
func (o *AudienceFiltersSegmentsItems) SetSegmentId(v string) {
	o.SegmentId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AudienceFiltersSegmentsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["segment_id"] = o.SegmentId
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AudienceFiltersSegmentsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string `json:"name"`
		SegmentId *string `json:"segment_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SegmentId == nil {
		return fmt.Errorf("required field segment_id missing")
	}
	o.Name = *all.Name
	o.SegmentId = *all.SegmentId

	return nil
}
