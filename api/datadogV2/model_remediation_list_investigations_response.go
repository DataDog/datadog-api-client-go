// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationListInvestigationsResponse Response payload for listing investigations.
type RemediationListInvestigationsResponse struct {
	// The matching investigations.
	Investigations []RemediationInvestigation `json:"investigations,omitempty"`
	// Token to pass as page_token on the next call. Empty when there are no further pages.
	NextPageToken *string `json:"next_page_token,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationListInvestigationsResponse instantiates a new RemediationListInvestigationsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationListInvestigationsResponse() *RemediationListInvestigationsResponse {
	this := RemediationListInvestigationsResponse{}
	return &this
}

// NewRemediationListInvestigationsResponseWithDefaults instantiates a new RemediationListInvestigationsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationListInvestigationsResponseWithDefaults() *RemediationListInvestigationsResponse {
	this := RemediationListInvestigationsResponse{}
	return &this
}

// GetInvestigations returns the Investigations field value if set, zero value otherwise.
func (o *RemediationListInvestigationsResponse) GetInvestigations() []RemediationInvestigation {
	if o == nil || o.Investigations == nil {
		var ret []RemediationInvestigation
		return ret
	}
	return o.Investigations
}

// GetInvestigationsOk returns a tuple with the Investigations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationListInvestigationsResponse) GetInvestigationsOk() (*[]RemediationInvestigation, bool) {
	if o == nil || o.Investigations == nil {
		return nil, false
	}
	return &o.Investigations, true
}

// HasInvestigations returns a boolean if a field has been set.
func (o *RemediationListInvestigationsResponse) HasInvestigations() bool {
	return o != nil && o.Investigations != nil
}

// SetInvestigations gets a reference to the given []RemediationInvestigation and assigns it to the Investigations field.
func (o *RemediationListInvestigationsResponse) SetInvestigations(v []RemediationInvestigation) {
	o.Investigations = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *RemediationListInvestigationsResponse) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationListInvestigationsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *RemediationListInvestigationsResponse) HasNextPageToken() bool {
	return o != nil && o.NextPageToken != nil
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *RemediationListInvestigationsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationListInvestigationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Investigations != nil {
		toSerialize["investigations"] = o.Investigations
	}
	if o.NextPageToken != nil {
		toSerialize["next_page_token"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationListInvestigationsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Investigations []RemediationInvestigation `json:"investigations,omitempty"`
		NextPageToken  *string                    `json:"next_page_token,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"investigations", "next_page_token"})
	} else {
		return err
	}
	o.Investigations = all.Investigations
	o.NextPageToken = all.NextPageToken

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
