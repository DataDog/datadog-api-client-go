// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPostmortemCreateAttributes The postmortem's attributes for a creation request.
type IncidentPostmortemCreateAttributes struct {
	// The URL of the postmortem document (for example, a Datadog notebook, Confluence page, or Google Doc).
	DocumentUrl string `json:"document_url"`
	// The title of the postmortem.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentPostmortemCreateAttributes instantiates a new IncidentPostmortemCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentPostmortemCreateAttributes(documentUrl string, title string) *IncidentPostmortemCreateAttributes {
	this := IncidentPostmortemCreateAttributes{}
	this.DocumentUrl = documentUrl
	this.Title = title
	return &this
}

// NewIncidentPostmortemCreateAttributesWithDefaults instantiates a new IncidentPostmortemCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentPostmortemCreateAttributesWithDefaults() *IncidentPostmortemCreateAttributes {
	this := IncidentPostmortemCreateAttributes{}
	return &this
}

// GetDocumentUrl returns the DocumentUrl field value.
func (o *IncidentPostmortemCreateAttributes) GetDocumentUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DocumentUrl
}

// GetDocumentUrlOk returns a tuple with the DocumentUrl field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemCreateAttributes) GetDocumentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentUrl, true
}

// SetDocumentUrl sets field value.
func (o *IncidentPostmortemCreateAttributes) SetDocumentUrl(v string) {
	o.DocumentUrl = v
}

// GetTitle returns the Title field value.
func (o *IncidentPostmortemCreateAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemCreateAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *IncidentPostmortemCreateAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentPostmortemCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["document_url"] = o.DocumentUrl
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentPostmortemCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DocumentUrl *string `json:"document_url"`
		Title       *string `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DocumentUrl == nil {
		return fmt.Errorf("required field document_url missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"document_url", "title"})
	} else {
		return err
	}
	o.DocumentUrl = *all.DocumentUrl
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
