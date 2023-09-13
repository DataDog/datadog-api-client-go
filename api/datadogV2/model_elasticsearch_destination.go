// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticsearchDestination The Elasticsearch destination.
type ElasticsearchDestination struct {
	// The HTTP destination basic username and password auth.
	Auth *HttpDestinationBasicAuth `json:"auth,omitempty"`
	// The intake URL to forward events to.
	Endpoint string `json:"endpoint"`
	// The Elasticsearch index name.
	IndexName string `json:"indexName"`
	// The pattern to append to the index name for rotation in Elasticsearch.
	IndexRotation string `json:"indexRotation"`
	// The Elasticsearch destination type.
	Type ElasticsearchDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewElasticsearchDestination instantiates a new ElasticsearchDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchDestination(endpoint string, indexName string, indexRotation string, typeVar ElasticsearchDestinationType) *ElasticsearchDestination {
	this := ElasticsearchDestination{}
	this.Endpoint = endpoint
	this.IndexName = indexName
	this.IndexRotation = indexRotation
	this.Type = typeVar
	return &this
}

// NewElasticsearchDestinationWithDefaults instantiates a new ElasticsearchDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchDestinationWithDefaults() *ElasticsearchDestination {
	this := ElasticsearchDestination{}
	var typeVar ElasticsearchDestinationType = ELASTICSEARCHDESTINATIONTYPE_ELASTICSEARCH
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ElasticsearchDestination) GetAuth() HttpDestinationBasicAuth {
	if o == nil || o.Auth == nil {
		var ret HttpDestinationBasicAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchDestination) GetAuthOk() (*HttpDestinationBasicAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ElasticsearchDestination) HasAuth() bool {
	return o != nil && o.Auth != nil
}

// SetAuth gets a reference to the given HttpDestinationBasicAuth and assigns it to the Auth field.
func (o *ElasticsearchDestination) SetAuth(v HttpDestinationBasicAuth) {
	o.Auth = &v
}

// GetEndpoint returns the Endpoint field value.
func (o *ElasticsearchDestination) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchDestination) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value.
func (o *ElasticsearchDestination) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetIndexName returns the IndexName field value.
func (o *ElasticsearchDestination) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchDestination) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value.
func (o *ElasticsearchDestination) SetIndexName(v string) {
	o.IndexName = v
}

// GetIndexRotation returns the IndexRotation field value.
func (o *ElasticsearchDestination) GetIndexRotation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IndexRotation
}

// GetIndexRotationOk returns a tuple with the IndexRotation field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchDestination) GetIndexRotationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexRotation, true
}

// SetIndexRotation sets field value.
func (o *ElasticsearchDestination) SetIndexRotation(v string) {
	o.IndexRotation = v
}

// GetType returns the Type field value.
func (o *ElasticsearchDestination) GetType() ElasticsearchDestinationType {
	if o == nil {
		var ret ElasticsearchDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchDestination) GetTypeOk() (*ElasticsearchDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ElasticsearchDestination) SetType(v ElasticsearchDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["indexName"] = o.IndexName
	toSerialize["indexRotation"] = o.IndexRotation
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth          *HttpDestinationBasicAuth     `json:"auth,omitempty"`
		Endpoint      *string                       `json:"endpoint"`
		IndexName     *string                       `json:"indexName"`
		IndexRotation *string                       `json:"indexRotation"`
		Type          *ElasticsearchDestinationType `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Endpoint == nil {
		return fmt.Errorf("required field endpoint missing")
	}
	if all.IndexName == nil {
		return fmt.Errorf("required field indexName missing")
	}
	if all.IndexRotation == nil {
		return fmt.Errorf("required field indexRotation missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "endpoint", "indexName", "indexRotation", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth != nil && all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = all.Auth
	o.Endpoint = *all.Endpoint
	o.IndexName = *all.IndexName
	o.IndexRotation = *all.IndexRotation
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
