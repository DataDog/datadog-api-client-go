// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatuspageIncidentEntry A Statuspage incident entry.
type IncidentStatuspageIncidentEntry struct {
	// The Datadog incident identifier.
	IncidentId string `json:"incident_id"`
	// The Statuspage page identifier.
	PageId string `json:"page_id"`
	// The URL of the Statuspage incident.
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentStatuspageIncidentEntry instantiates a new IncidentStatuspageIncidentEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentStatuspageIncidentEntry(incidentId string, pageId string) *IncidentStatuspageIncidentEntry {
	this := IncidentStatuspageIncidentEntry{}
	this.IncidentId = incidentId
	this.PageId = pageId
	return &this
}

// NewIncidentStatuspageIncidentEntryWithDefaults instantiates a new IncidentStatuspageIncidentEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentStatuspageIncidentEntryWithDefaults() *IncidentStatuspageIncidentEntry {
	this := IncidentStatuspageIncidentEntry{}
	return &this
}

// GetIncidentId returns the IncidentId field value.
func (o *IncidentStatuspageIncidentEntry) GetIncidentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IncidentId
}

// GetIncidentIdOk returns a tuple with the IncidentId field value
// and a boolean to check if the value has been set.
func (o *IncidentStatuspageIncidentEntry) GetIncidentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentId, true
}

// SetIncidentId sets field value.
func (o *IncidentStatuspageIncidentEntry) SetIncidentId(v string) {
	o.IncidentId = v
}

// GetPageId returns the PageId field value.
func (o *IncidentStatuspageIncidentEntry) GetPageId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value
// and a boolean to check if the value has been set.
func (o *IncidentStatuspageIncidentEntry) GetPageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageId, true
}

// SetPageId sets field value.
func (o *IncidentStatuspageIncidentEntry) SetPageId(v string) {
	o.PageId = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *IncidentStatuspageIncidentEntry) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentStatuspageIncidentEntry) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *IncidentStatuspageIncidentEntry) HasRedirectUrl() bool {
	return o != nil && o.RedirectUrl != nil
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *IncidentStatuspageIncidentEntry) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentStatuspageIncidentEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["incident_id"] = o.IncidentId
	toSerialize["page_id"] = o.PageId
	if o.RedirectUrl != nil {
		toSerialize["redirect_url"] = o.RedirectUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentStatuspageIncidentEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncidentId  *string `json:"incident_id"`
		PageId      *string `json:"page_id"`
		RedirectUrl *string `json:"redirect_url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncidentId == nil {
		return fmt.Errorf("required field incident_id missing")
	}
	if all.PageId == nil {
		return fmt.Errorf("required field page_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incident_id", "page_id", "redirect_url"})
	} else {
		return err
	}
	o.IncidentId = *all.IncidentId
	o.PageId = *all.PageId
	o.RedirectUrl = all.RedirectUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
