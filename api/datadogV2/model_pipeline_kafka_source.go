// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineKafkaSource The `kafka` source ingests data from Apache Kafka topics.
type PipelineKafkaSource struct {
	// Consumer group ID used by the Kafka client.
	GroupId string `json:"group_id"`
	// The unique identifier for this component. Used to reference this component in other parts of the pipeline (e.g., as input to downstream components).
	Id string `json:"id"`
	// Optional list of advanced Kafka client configuration options, defined as key-value pairs.
	LibrdkafkaOptions []PipelineKafkaSourceLibrdkafkaOption `json:"librdkafka_options,omitempty"`
	// Specifies the SASL mechanism for authenticating with a Kafka cluster.
	Sasl *PipelineKafkaSourceSasl `json:"sasl,omitempty"`
	// Configuration for enabling TLS encryption.
	Tls *PipelineTls `json:"tls,omitempty"`
	// A list of Kafka topic names to subscribe to. The source ingests messages from each topic specified.
	Topics []string `json:"topics"`
	// The source type. The value should always be `kafka`.
	Type PipelineKafkaSourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPipelineKafkaSource instantiates a new PipelineKafkaSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPipelineKafkaSource(groupId string, id string, topics []string, typeVar PipelineKafkaSourceType) *PipelineKafkaSource {
	this := PipelineKafkaSource{}
	this.GroupId = groupId
	this.Id = id
	this.Topics = topics
	this.Type = typeVar
	return &this
}

// NewPipelineKafkaSourceWithDefaults instantiates a new PipelineKafkaSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPipelineKafkaSourceWithDefaults() *PipelineKafkaSource {
	this := PipelineKafkaSource{}
	return &this
}

// GetGroupId returns the GroupId field value.
func (o *PipelineKafkaSource) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *PipelineKafkaSource) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value.
func (o *PipelineKafkaSource) SetGroupId(v string) {
	o.GroupId = v
}

// GetId returns the Id field value.
func (o *PipelineKafkaSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PipelineKafkaSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *PipelineKafkaSource) SetId(v string) {
	o.Id = v
}

// GetLibrdkafkaOptions returns the LibrdkafkaOptions field value if set, zero value otherwise.
func (o *PipelineKafkaSource) GetLibrdkafkaOptions() []PipelineKafkaSourceLibrdkafkaOption {
	if o == nil || o.LibrdkafkaOptions == nil {
		var ret []PipelineKafkaSourceLibrdkafkaOption
		return ret
	}
	return o.LibrdkafkaOptions
}

// GetLibrdkafkaOptionsOk returns a tuple with the LibrdkafkaOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineKafkaSource) GetLibrdkafkaOptionsOk() (*[]PipelineKafkaSourceLibrdkafkaOption, bool) {
	if o == nil || o.LibrdkafkaOptions == nil {
		return nil, false
	}
	return &o.LibrdkafkaOptions, true
}

// HasLibrdkafkaOptions returns a boolean if a field has been set.
func (o *PipelineKafkaSource) HasLibrdkafkaOptions() bool {
	return o != nil && o.LibrdkafkaOptions != nil
}

// SetLibrdkafkaOptions gets a reference to the given []PipelineKafkaSourceLibrdkafkaOption and assigns it to the LibrdkafkaOptions field.
func (o *PipelineKafkaSource) SetLibrdkafkaOptions(v []PipelineKafkaSourceLibrdkafkaOption) {
	o.LibrdkafkaOptions = v
}

// GetSasl returns the Sasl field value if set, zero value otherwise.
func (o *PipelineKafkaSource) GetSasl() PipelineKafkaSourceSasl {
	if o == nil || o.Sasl == nil {
		var ret PipelineKafkaSourceSasl
		return ret
	}
	return *o.Sasl
}

// GetSaslOk returns a tuple with the Sasl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineKafkaSource) GetSaslOk() (*PipelineKafkaSourceSasl, bool) {
	if o == nil || o.Sasl == nil {
		return nil, false
	}
	return o.Sasl, true
}

// HasSasl returns a boolean if a field has been set.
func (o *PipelineKafkaSource) HasSasl() bool {
	return o != nil && o.Sasl != nil
}

// SetSasl gets a reference to the given PipelineKafkaSourceSasl and assigns it to the Sasl field.
func (o *PipelineKafkaSource) SetSasl(v PipelineKafkaSourceSasl) {
	o.Sasl = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *PipelineKafkaSource) GetTls() PipelineTls {
	if o == nil || o.Tls == nil {
		var ret PipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineKafkaSource) GetTlsOk() (*PipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *PipelineKafkaSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given PipelineTls and assigns it to the Tls field.
func (o *PipelineKafkaSource) SetTls(v PipelineTls) {
	o.Tls = &v
}

// GetTopics returns the Topics field value.
func (o *PipelineKafkaSource) GetTopics() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value
// and a boolean to check if the value has been set.
func (o *PipelineKafkaSource) GetTopicsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topics, true
}

// SetTopics sets field value.
func (o *PipelineKafkaSource) SetTopics(v []string) {
	o.Topics = v
}

// GetType returns the Type field value.
func (o *PipelineKafkaSource) GetType() PipelineKafkaSourceType {
	if o == nil {
		var ret PipelineKafkaSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PipelineKafkaSource) GetTypeOk() (*PipelineKafkaSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *PipelineKafkaSource) SetType(v PipelineKafkaSourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PipelineKafkaSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["group_id"] = o.GroupId
	toSerialize["id"] = o.Id
	if o.LibrdkafkaOptions != nil {
		toSerialize["librdkafka_options"] = o.LibrdkafkaOptions
	}
	if o.Sasl != nil {
		toSerialize["sasl"] = o.Sasl
	}
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["topics"] = o.Topics
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PipelineKafkaSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupId           *string                               `json:"group_id"`
		Id                *string                               `json:"id"`
		LibrdkafkaOptions []PipelineKafkaSourceLibrdkafkaOption `json:"librdkafka_options,omitempty"`
		Sasl              *PipelineKafkaSourceSasl              `json:"sasl,omitempty"`
		Tls               *PipelineTls                          `json:"tls,omitempty"`
		Topics            *[]string                             `json:"topics"`
		Type              *PipelineKafkaSourceType              `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GroupId == nil {
		return fmt.Errorf("required field group_id missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Topics == nil {
		return fmt.Errorf("required field topics missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group_id", "id", "librdkafka_options", "sasl", "tls", "topics", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.GroupId = *all.GroupId
	o.Id = *all.Id
	o.LibrdkafkaOptions = all.LibrdkafkaOptions
	if all.Sasl != nil && all.Sasl.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sasl = all.Sasl
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
	o.Topics = *all.Topics
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
