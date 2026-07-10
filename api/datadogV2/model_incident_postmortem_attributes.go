// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPostmortemAttributes The postmortem's attributes.
type IncidentPostmortemAttributes struct {
	// Timestamp when the postmortem was created.
	Created time.Time `json:"created"`
	// The identifier of the postmortem document within its host platform.
	DocumentId string `json:"document_id"`
	// The type of document backing the postmortem (for example, `datadog_notebooks`, `confluence`, or `google_docs`). Can be empty if the document type is unknown.
	DocumentType string `json:"document_type"`
	// The URL of the postmortem document.
	DocumentUrl string `json:"document_url"`
	// Timestamp when the postmortem was last modified.
	Modified time.Time `json:"modified"`
	// The status of the postmortem.
	Status PostmortemStatus `json:"status"`
	// The title of the postmortem.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentPostmortemAttributes instantiates a new IncidentPostmortemAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentPostmortemAttributes(created time.Time, documentId string, documentType string, documentUrl string, modified time.Time, status PostmortemStatus, title string) *IncidentPostmortemAttributes {
	this := IncidentPostmortemAttributes{}
	this.Created = created
	this.DocumentId = documentId
	this.DocumentType = documentType
	this.DocumentUrl = documentUrl
	this.Modified = modified
	this.Status = status
	this.Title = title
	return &this
}

// NewIncidentPostmortemAttributesWithDefaults instantiates a new IncidentPostmortemAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentPostmortemAttributesWithDefaults() *IncidentPostmortemAttributes {
	this := IncidentPostmortemAttributes{}
	return &this
}

// GetCreated returns the Created field value.
func (o *IncidentPostmortemAttributes) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentPostmortemAttributes) SetCreated(v time.Time) {
	o.Created = v
}

// GetDocumentId returns the DocumentId field value.
func (o *IncidentPostmortemAttributes) GetDocumentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemAttributes) GetDocumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentId, true
}

// SetDocumentId sets field value.
func (o *IncidentPostmortemAttributes) SetDocumentId(v string) {
	o.DocumentId = v
}

// GetDocumentType returns the DocumentType field value.
func (o *IncidentPostmortemAttributes) GetDocumentType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemAttributes) GetDocumentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value.
func (o *IncidentPostmortemAttributes) SetDocumentType(v string) {
	o.DocumentType = v
}

// GetDocumentUrl returns the DocumentUrl field value.
func (o *IncidentPostmortemAttributes) GetDocumentUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DocumentUrl
}

// GetDocumentUrlOk returns a tuple with the DocumentUrl field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemAttributes) GetDocumentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentUrl, true
}

// SetDocumentUrl sets field value.
func (o *IncidentPostmortemAttributes) SetDocumentUrl(v string) {
	o.DocumentUrl = v
}

// GetModified returns the Modified field value.
func (o *IncidentPostmortemAttributes) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentPostmortemAttributes) SetModified(v time.Time) {
	o.Modified = v
}

// GetStatus returns the Status field value.
func (o *IncidentPostmortemAttributes) GetStatus() PostmortemStatus {
	if o == nil {
		var ret PostmortemStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemAttributes) GetStatusOk() (*PostmortemStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *IncidentPostmortemAttributes) SetStatus(v PostmortemStatus) {
	o.Status = v
}

// GetTitle returns the Title field value.
func (o *IncidentPostmortemAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *IncidentPostmortemAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentPostmortemAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["document_id"] = o.DocumentId
	toSerialize["document_type"] = o.DocumentType
	toSerialize["document_url"] = o.DocumentUrl
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentPostmortemAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created      *time.Time        `json:"created"`
		DocumentId   *string           `json:"document_id"`
		DocumentType *string           `json:"document_type"`
		DocumentUrl  *string           `json:"document_url"`
		Modified     *time.Time        `json:"modified"`
		Status       *PostmortemStatus `json:"status"`
		Title        *string           `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.DocumentId == nil {
		return fmt.Errorf("required field document_id missing")
	}
	if all.DocumentType == nil {
		return fmt.Errorf("required field document_type missing")
	}
	if all.DocumentUrl == nil {
		return fmt.Errorf("required field document_url missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "document_id", "document_type", "document_url", "modified", "status", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Created = *all.Created
	o.DocumentId = *all.DocumentId
	o.DocumentType = *all.DocumentType
	o.DocumentUrl = *all.DocumentUrl
	o.Modified = *all.Modified
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
