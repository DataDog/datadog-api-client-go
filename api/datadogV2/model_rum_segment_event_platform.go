// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentEventPlatform An event platform query block within a segment data query.
type RumSegmentEventPlatform struct {
	// The facet to extract user identifiers from.
	Facet string `json:"facet"`
	// The start of the time range in milliseconds since epoch.
	From *int64 `json:"from,omitempty"`
	// The name of this query block.
	Name string `json:"name"`
	// The search query for filtering events.
	Query string `json:"query"`
	// The end of the time range in milliseconds since epoch.
	To *int64 `json:"to,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentEventPlatform instantiates a new RumSegmentEventPlatform object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentEventPlatform(facet string, name string, query string) *RumSegmentEventPlatform {
	this := RumSegmentEventPlatform{}
	this.Facet = facet
	this.Name = name
	this.Query = query
	return &this
}

// NewRumSegmentEventPlatformWithDefaults instantiates a new RumSegmentEventPlatform object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentEventPlatformWithDefaults() *RumSegmentEventPlatform {
	this := RumSegmentEventPlatform{}
	return &this
}

// GetFacet returns the Facet field value.
func (o *RumSegmentEventPlatform) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *RumSegmentEventPlatform) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *RumSegmentEventPlatform) SetFacet(v string) {
	o.Facet = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *RumSegmentEventPlatform) GetFrom() int64 {
	if o == nil || o.From == nil {
		var ret int64
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentEventPlatform) GetFromOk() (*int64, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *RumSegmentEventPlatform) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given int64 and assigns it to the From field.
func (o *RumSegmentEventPlatform) SetFrom(v int64) {
	o.From = &v
}

// GetName returns the Name field value.
func (o *RumSegmentEventPlatform) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RumSegmentEventPlatform) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RumSegmentEventPlatform) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *RumSegmentEventPlatform) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *RumSegmentEventPlatform) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *RumSegmentEventPlatform) SetQuery(v string) {
	o.Query = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *RumSegmentEventPlatform) GetTo() int64 {
	if o == nil || o.To == nil {
		var ret int64
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentEventPlatform) GetToOk() (*int64, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *RumSegmentEventPlatform) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given int64 and assigns it to the To field.
func (o *RumSegmentEventPlatform) SetTo(v int64) {
	o.To = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentEventPlatform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["facet"] = o.Facet
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	if o.To != nil {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentEventPlatform) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet *string `json:"facet"`
		From  *int64  `json:"from,omitempty"`
		Name  *string `json:"name"`
		Query *string `json:"query"`
		To    *int64  `json:"to,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet", "from", "name", "query", "to"})
	} else {
		return err
	}
	o.Facet = *all.Facet
	o.From = all.From
	o.Name = *all.Name
	o.Query = *all.Query
	o.To = all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
