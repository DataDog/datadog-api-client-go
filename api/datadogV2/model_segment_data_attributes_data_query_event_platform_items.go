// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentDataAttributesDataQueryEventPlatformItems
type SegmentDataAttributesDataQueryEventPlatformItems struct {
	//
	Facet string `json:"facet"`
	//
	From *string `json:"from,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Query *string `json:"query,omitempty"`
	//
	To *string `json:"to,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentDataAttributesDataQueryEventPlatformItems instantiates a new SegmentDataAttributesDataQueryEventPlatformItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentDataAttributesDataQueryEventPlatformItems(facet string) *SegmentDataAttributesDataQueryEventPlatformItems {
	this := SegmentDataAttributesDataQueryEventPlatformItems{}
	this.Facet = facet
	return &this
}

// NewSegmentDataAttributesDataQueryEventPlatformItemsWithDefaults instantiates a new SegmentDataAttributesDataQueryEventPlatformItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentDataAttributesDataQueryEventPlatformItemsWithDefaults() *SegmentDataAttributesDataQueryEventPlatformItems {
	this := SegmentDataAttributesDataQueryEventPlatformItems{}
	return &this
}

// GetFacet returns the Facet field value.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) SetFacet(v string) {
	o.Facet = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) SetFrom(v string) {
	o.From = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) SetQuery(v string) {
	o.Query = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) SetTo(v string) {
	o.To = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentDataAttributesDataQueryEventPlatformItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["facet"] = o.Facet
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentDataAttributesDataQueryEventPlatformItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet *string `json:"facet"`
		From  *string `json:"from,omitempty"`
		Name  *string `json:"name,omitempty"`
		Query *string `json:"query,omitempty"`
		To    *string `json:"to,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet", "from", "name", "query", "to"})
	} else {
		return err
	}
	o.Facet = *all.Facet
	o.From = all.From
	o.Name = all.Name
	o.Query = all.Query
	o.To = all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
