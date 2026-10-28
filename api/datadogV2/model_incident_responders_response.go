// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRespondersResponse Response with a list of incident responders.
type IncidentRespondersResponse struct {
	// List of incident responders.
	Data []IncidentResponderDataResponse `json:"data"`
	// Included related resources.
	Included []IncidentResponderIncludedItem `json:"included,omitempty"`
	// Links to pages of incident responders.
	Links map[string]string `json:"links,omitempty"`
	// Pagination metadata for the responder list.
	Meta *IncidentRespondersResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRespondersResponse instantiates a new IncidentRespondersResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRespondersResponse(data []IncidentResponderDataResponse) *IncidentRespondersResponse {
	this := IncidentRespondersResponse{}
	this.Data = data
	return &this
}

// NewIncidentRespondersResponseWithDefaults instantiates a new IncidentRespondersResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRespondersResponseWithDefaults() *IncidentRespondersResponse {
	this := IncidentRespondersResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentRespondersResponse) GetData() []IncidentResponderDataResponse {
	if o == nil {
		var ret []IncidentResponderDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentRespondersResponse) GetDataOk() (*[]IncidentResponderDataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentRespondersResponse) SetData(v []IncidentResponderDataResponse) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentRespondersResponse) GetIncluded() []IncidentResponderIncludedItem {
	if o == nil || o.Included == nil {
		var ret []IncidentResponderIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRespondersResponse) GetIncludedOk() (*[]IncidentResponderIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentRespondersResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []IncidentResponderIncludedItem and assigns it to the Included field.
func (o *IncidentRespondersResponse) SetIncluded(v []IncidentResponderIncludedItem) {
	o.Included = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IncidentRespondersResponse) GetLinks() map[string]string {
	if o == nil || o.Links == nil {
		var ret map[string]string
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRespondersResponse) GetLinksOk() (*map[string]string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IncidentRespondersResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *IncidentRespondersResponse) SetLinks(v map[string]string) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentRespondersResponse) GetMeta() IncidentRespondersResponseMeta {
	if o == nil || o.Meta == nil {
		var ret IncidentRespondersResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRespondersResponse) GetMetaOk() (*IncidentRespondersResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentRespondersResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given IncidentRespondersResponseMeta and assigns it to the Meta field.
func (o *IncidentRespondersResponse) SetMeta(v IncidentRespondersResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRespondersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRespondersResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]IncidentResponderDataResponse `json:"data"`
		Included []IncidentResponderIncludedItem  `json:"included,omitempty"`
		Links    map[string]string                `json:"links,omitempty"`
		Meta     *IncidentRespondersResponseMeta  `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	o.Included = all.Included
	o.Links = all.Links
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
