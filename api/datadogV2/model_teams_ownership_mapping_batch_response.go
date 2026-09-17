// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsOwnershipMappingBatchResponse The response body for the bulk create and remove operation. On success, `atomic:results`
// contains one entry per operation. Add results appear before remove results and may not match
// request order. Correlate add results by their `type` and `id` rather than by array position.
// On failure, no operations were applied and `errors` describes what went wrong.
type TeamsOwnershipMappingBatchResponse struct {
	// The result of each operation.
	// Add operations are processed first, then remove operations, so results may not appear
	// in the same order as the request. Present only on success.
	AtomicResults []TeamsOwnershipMappingBatchResult `json:"atomic:results,omitempty"`
	// The validation or processing errors encountered. Present only when the request could not be completed.
	Errors []TeamsOwnershipMappingBatchError `json:"errors,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsOwnershipMappingBatchResponse instantiates a new TeamsOwnershipMappingBatchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsOwnershipMappingBatchResponse() *TeamsOwnershipMappingBatchResponse {
	this := TeamsOwnershipMappingBatchResponse{}
	return &this
}

// NewTeamsOwnershipMappingBatchResponseWithDefaults instantiates a new TeamsOwnershipMappingBatchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsOwnershipMappingBatchResponseWithDefaults() *TeamsOwnershipMappingBatchResponse {
	this := TeamsOwnershipMappingBatchResponse{}
	return &this
}

// GetAtomicResults returns the AtomicResults field value if set, zero value otherwise.
func (o *TeamsOwnershipMappingBatchResponse) GetAtomicResults() []TeamsOwnershipMappingBatchResult {
	if o == nil || o.AtomicResults == nil {
		var ret []TeamsOwnershipMappingBatchResult
		return ret
	}
	return o.AtomicResults
}

// GetAtomicResultsOk returns a tuple with the AtomicResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchResponse) GetAtomicResultsOk() (*[]TeamsOwnershipMappingBatchResult, bool) {
	if o == nil || o.AtomicResults == nil {
		return nil, false
	}
	return &o.AtomicResults, true
}

// HasAtomicResults returns a boolean if a field has been set.
func (o *TeamsOwnershipMappingBatchResponse) HasAtomicResults() bool {
	return o != nil && o.AtomicResults != nil
}

// SetAtomicResults gets a reference to the given []TeamsOwnershipMappingBatchResult and assigns it to the AtomicResults field.
func (o *TeamsOwnershipMappingBatchResponse) SetAtomicResults(v []TeamsOwnershipMappingBatchResult) {
	o.AtomicResults = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *TeamsOwnershipMappingBatchResponse) GetErrors() []TeamsOwnershipMappingBatchError {
	if o == nil || o.Errors == nil {
		var ret []TeamsOwnershipMappingBatchError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsOwnershipMappingBatchResponse) GetErrorsOk() (*[]TeamsOwnershipMappingBatchError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *TeamsOwnershipMappingBatchResponse) HasErrors() bool {
	return o != nil && o.Errors != nil
}

// SetErrors gets a reference to the given []TeamsOwnershipMappingBatchError and assigns it to the Errors field.
func (o *TeamsOwnershipMappingBatchResponse) SetErrors(v []TeamsOwnershipMappingBatchError) {
	o.Errors = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsOwnershipMappingBatchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AtomicResults != nil {
		toSerialize["atomic:results"] = o.AtomicResults
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamsOwnershipMappingBatchResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AtomicResults []TeamsOwnershipMappingBatchResult `json:"atomic:results,omitempty"`
		Errors        []TeamsOwnershipMappingBatchError  `json:"errors,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"atomic:results", "errors"})
	} else {
		return err
	}
	o.AtomicResults = all.AtomicResults
	o.Errors = all.Errors

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
