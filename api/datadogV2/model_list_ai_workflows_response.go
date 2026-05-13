// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAIWorkflowsResponse Response containing a list of AI workflows.
type ListAIWorkflowsResponse struct {
	// List of AI workflow resources.
	Data []AIWorkflowData `json:"data"`
	// Metadata for the list AI workflows response.
	Meta AIWorkflowListMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListAIWorkflowsResponse instantiates a new ListAIWorkflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListAIWorkflowsResponse(data []AIWorkflowData, meta AIWorkflowListMeta) *ListAIWorkflowsResponse {
	this := ListAIWorkflowsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewListAIWorkflowsResponseWithDefaults instantiates a new ListAIWorkflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListAIWorkflowsResponseWithDefaults() *ListAIWorkflowsResponse {
	this := ListAIWorkflowsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ListAIWorkflowsResponse) GetData() []AIWorkflowData {
	if o == nil {
		var ret []AIWorkflowData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListAIWorkflowsResponse) GetDataOk() (*[]AIWorkflowData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ListAIWorkflowsResponse) SetData(v []AIWorkflowData) {
	o.Data = v
}

// GetMeta returns the Meta field value.
func (o *ListAIWorkflowsResponse) GetMeta() AIWorkflowListMeta {
	if o == nil {
		var ret AIWorkflowListMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ListAIWorkflowsResponse) GetMetaOk() (*AIWorkflowListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *ListAIWorkflowsResponse) SetMeta(v AIWorkflowListMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListAIWorkflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListAIWorkflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]AIWorkflowData   `json:"data"`
		Meta *AIWorkflowListMeta `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
