// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentEventPlatformBlock Selects the users or accounts that sent matching RUM or Product Analytics events in a time range.
// `from` and `to` are epoch milliseconds, or negative offsets in milliseconds relative to now. `0` means now.
type SegmentEventPlatformBlock struct {
	// Facet that identifies the entity, such as `@usr.id`.
	Facet string `json:"facet"`
	// Start of the time range.
	From *string `json:"from,omitempty"`
	// Name of the block, used in `combination`.
	Name *string `json:"name,omitempty"`
	// Event search query.
	Query *string `json:"query,omitempty"`
	// End of the time range.
	To *string `json:"to,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentEventPlatformBlock instantiates a new SegmentEventPlatformBlock object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentEventPlatformBlock(facet string) *SegmentEventPlatformBlock {
	this := SegmentEventPlatformBlock{}
	this.Facet = facet
	return &this
}

// NewSegmentEventPlatformBlockWithDefaults instantiates a new SegmentEventPlatformBlock object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentEventPlatformBlockWithDefaults() *SegmentEventPlatformBlock {
	this := SegmentEventPlatformBlock{}
	return &this
}

// GetFacet returns the Facet field value.
func (o *SegmentEventPlatformBlock) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *SegmentEventPlatformBlock) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *SegmentEventPlatformBlock) SetFacet(v string) {
	o.Facet = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *SegmentEventPlatformBlock) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentEventPlatformBlock) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *SegmentEventPlatformBlock) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *SegmentEventPlatformBlock) SetFrom(v string) {
	o.From = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SegmentEventPlatformBlock) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentEventPlatformBlock) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SegmentEventPlatformBlock) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SegmentEventPlatformBlock) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SegmentEventPlatformBlock) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentEventPlatformBlock) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SegmentEventPlatformBlock) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SegmentEventPlatformBlock) SetQuery(v string) {
	o.Query = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *SegmentEventPlatformBlock) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentEventPlatformBlock) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *SegmentEventPlatformBlock) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *SegmentEventPlatformBlock) SetTo(v string) {
	o.To = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentEventPlatformBlock) MarshalJSON() ([]byte, error) {
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
func (o *SegmentEventPlatformBlock) UnmarshalJSON(bytes []byte) (err error) {
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
