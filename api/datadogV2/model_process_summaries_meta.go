// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.


package datadogV2

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"

)


// ProcessSummariesMeta Response metadata object.
type ProcessSummariesMeta struct {
	// Paging attributes.
	Page *ProcessSummariesMetaPage `json:"page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewProcessSummariesMeta instantiates a new ProcessSummariesMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProcessSummariesMeta() *ProcessSummariesMeta {
	this := ProcessSummariesMeta{}
	return &this
}

// NewProcessSummariesMetaWithDefaults instantiates a new ProcessSummariesMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProcessSummariesMetaWithDefaults() *ProcessSummariesMeta {
	this := ProcessSummariesMeta{}
	return &this
}
// GetPage returns the Page field value if set, zero value otherwise.
func (o *ProcessSummariesMeta) GetPage() ProcessSummariesMetaPage {
	if o == nil || o.Page == nil {
		var ret ProcessSummariesMetaPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessSummariesMeta) GetPageOk() (*ProcessSummariesMetaPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ProcessSummariesMeta) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given ProcessSummariesMetaPage and assigns it to the Page field.
func (o *ProcessSummariesMeta) SetPage(v ProcessSummariesMetaPage) {
	o.Page = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o ProcessSummariesMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProcessSummariesMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page *ProcessSummariesMetaPage `json:"page,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{ "page",  })
	} else {
		return err
	}

	hasInvalidField := false
	if  all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
