// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRespondersResponseMeta Pagination metadata for the responder list.
type IncidentRespondersResponseMeta struct {
	// Total count for the responder list.
	Page *IncidentRespondersResponseMetaPage `json:"page,omitempty"`
	// Offset-based pagination for incident responders.
	Pagination *IncidentRespondersPagination `json:"pagination,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRespondersResponseMeta instantiates a new IncidentRespondersResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRespondersResponseMeta() *IncidentRespondersResponseMeta {
	this := IncidentRespondersResponseMeta{}
	return &this
}

// NewIncidentRespondersResponseMetaWithDefaults instantiates a new IncidentRespondersResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRespondersResponseMetaWithDefaults() *IncidentRespondersResponseMeta {
	this := IncidentRespondersResponseMeta{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *IncidentRespondersResponseMeta) GetPage() IncidentRespondersResponseMetaPage {
	if o == nil || o.Page == nil {
		var ret IncidentRespondersResponseMetaPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRespondersResponseMeta) GetPageOk() (*IncidentRespondersResponseMetaPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *IncidentRespondersResponseMeta) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given IncidentRespondersResponseMetaPage and assigns it to the Page field.
func (o *IncidentRespondersResponseMeta) SetPage(v IncidentRespondersResponseMetaPage) {
	o.Page = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *IncidentRespondersResponseMeta) GetPagination() IncidentRespondersPagination {
	if o == nil || o.Pagination == nil {
		var ret IncidentRespondersPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRespondersResponseMeta) GetPaginationOk() (*IncidentRespondersPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *IncidentRespondersResponseMeta) HasPagination() bool {
	return o != nil && o.Pagination != nil
}

// SetPagination gets a reference to the given IncidentRespondersPagination and assigns it to the Pagination field.
func (o *IncidentRespondersResponseMeta) SetPagination(v IncidentRespondersPagination) {
	o.Pagination = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRespondersResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRespondersResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page       *IncidentRespondersResponseMetaPage `json:"page,omitempty"`
		Pagination *IncidentRespondersPagination       `json:"pagination,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"page", "pagination"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page
	if all.Pagination != nil && all.Pagination.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagination = all.Pagination

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
