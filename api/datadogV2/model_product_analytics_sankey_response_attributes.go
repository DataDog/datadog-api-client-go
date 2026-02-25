// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSankeyResponseAttributes
type ProductAnalyticsSankeyResponseAttributes struct {
	// The links (flows) between nodes.
	Links []ProductAnalyticsSankeyLink `json:"links,omitempty"`
	// The nodes (pages) in the Sankey diagram.
	Nodes []ProductAnalyticsSankeyNode `json:"nodes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsSankeyResponseAttributes instantiates a new ProductAnalyticsSankeyResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsSankeyResponseAttributes() *ProductAnalyticsSankeyResponseAttributes {
	this := ProductAnalyticsSankeyResponseAttributes{}
	return &this
}

// NewProductAnalyticsSankeyResponseAttributesWithDefaults instantiates a new ProductAnalyticsSankeyResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsSankeyResponseAttributesWithDefaults() *ProductAnalyticsSankeyResponseAttributes {
	this := ProductAnalyticsSankeyResponseAttributes{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyResponseAttributes) GetLinks() []ProductAnalyticsSankeyLink {
	if o == nil || o.Links == nil {
		var ret []ProductAnalyticsSankeyLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyResponseAttributes) GetLinksOk() (*[]ProductAnalyticsSankeyLink, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyResponseAttributes) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given []ProductAnalyticsSankeyLink and assigns it to the Links field.
func (o *ProductAnalyticsSankeyResponseAttributes) SetLinks(v []ProductAnalyticsSankeyLink) {
	o.Links = v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *ProductAnalyticsSankeyResponseAttributes) GetNodes() []ProductAnalyticsSankeyNode {
	if o == nil || o.Nodes == nil {
		var ret []ProductAnalyticsSankeyNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyResponseAttributes) GetNodesOk() (*[]ProductAnalyticsSankeyNode, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *ProductAnalyticsSankeyResponseAttributes) HasNodes() bool {
	return o != nil && o.Nodes != nil
}

// SetNodes gets a reference to the given []ProductAnalyticsSankeyNode and assigns it to the Nodes field.
func (o *ProductAnalyticsSankeyResponseAttributes) SetNodes(v []ProductAnalyticsSankeyNode) {
	o.Nodes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsSankeyResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsSankeyResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Links []ProductAnalyticsSankeyLink `json:"links,omitempty"`
		Nodes []ProductAnalyticsSankeyNode `json:"nodes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"links", "nodes"})
	} else {
		return err
	}
	o.Links = all.Links
	o.Nodes = all.Nodes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
