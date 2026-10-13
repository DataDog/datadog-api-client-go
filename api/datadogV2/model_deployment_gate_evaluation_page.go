// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGateEvaluationPage Cursor pagination information.
type DeploymentGateEvaluationPage struct {
	// Opaque cursor for the next page. Absent on the final page.
	NextCursor *string `json:"next_cursor,omitempty"`
	// Requested maximum number of resources in this page.
	Size int64 `json:"size"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGateEvaluationPage instantiates a new DeploymentGateEvaluationPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGateEvaluationPage(size int64) *DeploymentGateEvaluationPage {
	this := DeploymentGateEvaluationPage{}
	this.Size = size
	return &this
}

// NewDeploymentGateEvaluationPageWithDefaults instantiates a new DeploymentGateEvaluationPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGateEvaluationPageWithDefaults() *DeploymentGateEvaluationPage {
	this := DeploymentGateEvaluationPage{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *DeploymentGateEvaluationPage) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationPage) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *DeploymentGateEvaluationPage) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *DeploymentGateEvaluationPage) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetSize returns the Size field value.
func (o *DeploymentGateEvaluationPage) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *DeploymentGateEvaluationPage) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *DeploymentGateEvaluationPage) SetSize(v int64) {
	o.Size = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGateEvaluationPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	toSerialize["size"] = o.Size

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGateEvaluationPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextCursor *string `json:"next_cursor,omitempty"`
		Size       *int64  `json:"size"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Size == nil {
		return fmt.Errorf("required field size missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next_cursor", "size"})
	} else {
		return err
	}
	o.NextCursor = all.NextCursor
	o.Size = *all.Size

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
