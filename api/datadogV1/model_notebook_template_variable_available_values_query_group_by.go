// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotebookTemplateVariableAvailableValuesQueryGroupBy A group-by facet for an available values query.
type NotebookTemplateVariableAvailableValuesQueryGroupBy struct {
	// The facet name to group by.
	Facet string `json:"facet"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNotebookTemplateVariableAvailableValuesQueryGroupBy instantiates a new NotebookTemplateVariableAvailableValuesQueryGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotebookTemplateVariableAvailableValuesQueryGroupBy(facet string) *NotebookTemplateVariableAvailableValuesQueryGroupBy {
	this := NotebookTemplateVariableAvailableValuesQueryGroupBy{}
	this.Facet = facet
	return &this
}

// NewNotebookTemplateVariableAvailableValuesQueryGroupByWithDefaults instantiates a new NotebookTemplateVariableAvailableValuesQueryGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotebookTemplateVariableAvailableValuesQueryGroupByWithDefaults() *NotebookTemplateVariableAvailableValuesQueryGroupBy {
	this := NotebookTemplateVariableAvailableValuesQueryGroupBy{}
	return &this
}

// GetFacet returns the Facet field value.
func (o *NotebookTemplateVariableAvailableValuesQueryGroupBy) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *NotebookTemplateVariableAvailableValuesQueryGroupBy) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *NotebookTemplateVariableAvailableValuesQueryGroupBy) SetFacet(v string) {
	o.Facet = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotebookTemplateVariableAvailableValuesQueryGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["facet"] = o.Facet
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotebookTemplateVariableAvailableValuesQueryGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet *string `json:"facet"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	o.Facet = *all.Facet

	return nil
}
