// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataTransformationDescriptionResponse
type DataTransformationDescriptionResponse struct {
	// Detailed description of the transformation.
	Details string `json:"details"`
	// The generation ID.
	Id string `json:"id"`
	// A brief summary of the transformation.
	Summary string `json:"summary"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataTransformationDescriptionResponse instantiates a new DataTransformationDescriptionResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataTransformationDescriptionResponse(details string, id string, summary string) *DataTransformationDescriptionResponse {
	this := DataTransformationDescriptionResponse{}
	this.Details = details
	this.Id = id
	this.Summary = summary
	return &this
}

// NewDataTransformationDescriptionResponseWithDefaults instantiates a new DataTransformationDescriptionResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataTransformationDescriptionResponseWithDefaults() *DataTransformationDescriptionResponse {
	this := DataTransformationDescriptionResponse{}
	return &this
}

// GetDetails returns the Details field value.
func (o *DataTransformationDescriptionResponse) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *DataTransformationDescriptionResponse) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value.
func (o *DataTransformationDescriptionResponse) SetDetails(v string) {
	o.Details = v
}

// GetId returns the Id field value.
func (o *DataTransformationDescriptionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DataTransformationDescriptionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DataTransformationDescriptionResponse) SetId(v string) {
	o.Id = v
}

// GetSummary returns the Summary field value.
func (o *DataTransformationDescriptionResponse) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DataTransformationDescriptionResponse) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DataTransformationDescriptionResponse) SetSummary(v string) {
	o.Summary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataTransformationDescriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["details"] = o.Details
	toSerialize["id"] = o.Id
	toSerialize["summary"] = o.Summary

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataTransformationDescriptionResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Details *string `json:"details"`
		Id      *string `json:"id"`
		Summary *string `json:"summary"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Details == nil {
		return fmt.Errorf("required field details missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"details", "id", "summary"})
	} else {
		return err
	}
	o.Details = *all.Details
	o.Id = *all.Id
	o.Summary = *all.Summary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
