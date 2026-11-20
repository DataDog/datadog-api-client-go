// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCrowdStrikeFdrSource The `crowdstrike_fdr` source ingests CrowdStrike Falcon Data Replicator (FDR) data.
// It reads file notifications from the CrowdStrike-managed SQS queue for your FDR feed and downloads the referenced objects from the CrowdStrike-managed S3 bucket.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineCrowdStrikeFdrSource struct {
	// Name of the environment variable or secret that holds the AWS access key ID that CrowdStrike generated for the FDR feed.
	AccessKeyIdKey *string `json:"access_key_id_key,omitempty"`
	// Number of concurrent tasks polling the SQS queue. Defaults to the number of Worker threads.
	ClientConcurrency *int64 `json:"client_concurrency,omitempty"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// Maximum number of SQS messages to receive per poll. Each message can reference multiple files. Defaults to 1.
	MaxNumberOfMessages *int64 `json:"max_number_of_messages,omitempty"`
	// AWS region of the FDR feed's SQS queue and S3 bucket.
	Region string `json:"region"`
	// Name of the environment variable or secret that holds the AWS secret access key that CrowdStrike generated for the FDR feed.
	SecretAccessKeyKey *string `json:"secret_access_key_key,omitempty"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The source type. Always `crowdstrike_fdr`.
	Type ObservabilityPipelineCrowdStrikeFdrSourceType `json:"type"`
	// Name of the environment variable or secret that holds the URL of the FDR feed's SQS queue.
	UrlKey *string `json:"url_key,omitempty"`
	// Number of seconds a received SQS message is hidden from other consumers. Set it higher than the time needed to download and process the largest FDR file. Defaults to 300.
	VisibilityTimeoutSecs *int64 `json:"visibility_timeout_secs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineCrowdStrikeFdrSource instantiates a new ObservabilityPipelineCrowdStrikeFdrSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineCrowdStrikeFdrSource(id string, region string, typeVar ObservabilityPipelineCrowdStrikeFdrSourceType) *ObservabilityPipelineCrowdStrikeFdrSource {
	this := ObservabilityPipelineCrowdStrikeFdrSource{}
	this.Id = id
	this.Region = region
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineCrowdStrikeFdrSourceWithDefaults instantiates a new ObservabilityPipelineCrowdStrikeFdrSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineCrowdStrikeFdrSourceWithDefaults() *ObservabilityPipelineCrowdStrikeFdrSource {
	this := ObservabilityPipelineCrowdStrikeFdrSource{}
	var typeVar ObservabilityPipelineCrowdStrikeFdrSourceType = OBSERVABILITYPIPELINECROWDSTRIKEFDRSOURCETYPE_CROWDSTRIKE_FDR
	this.Type = typeVar
	return &this
}

// GetAccessKeyIdKey returns the AccessKeyIdKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetAccessKeyIdKey() string {
	if o == nil || o.AccessKeyIdKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyIdKey
}

// GetAccessKeyIdKeyOk returns a tuple with the AccessKeyIdKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetAccessKeyIdKeyOk() (*string, bool) {
	if o == nil || o.AccessKeyIdKey == nil {
		return nil, false
	}
	return o.AccessKeyIdKey, true
}

// HasAccessKeyIdKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) HasAccessKeyIdKey() bool {
	return o != nil && o.AccessKeyIdKey != nil
}

// SetAccessKeyIdKey gets a reference to the given string and assigns it to the AccessKeyIdKey field.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) SetAccessKeyIdKey(v string) {
	o.AccessKeyIdKey = &v
}

// GetClientConcurrency returns the ClientConcurrency field value if set, zero value otherwise.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetClientConcurrency() int64 {
	if o == nil || o.ClientConcurrency == nil {
		var ret int64
		return ret
	}
	return *o.ClientConcurrency
}

// GetClientConcurrencyOk returns a tuple with the ClientConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetClientConcurrencyOk() (*int64, bool) {
	if o == nil || o.ClientConcurrency == nil {
		return nil, false
	}
	return o.ClientConcurrency, true
}

// HasClientConcurrency returns a boolean if a field has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) HasClientConcurrency() bool {
	return o != nil && o.ClientConcurrency != nil
}

// SetClientConcurrency gets a reference to the given int64 and assigns it to the ClientConcurrency field.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) SetClientConcurrency(v int64) {
	o.ClientConcurrency = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) SetId(v string) {
	o.Id = v
}

// GetMaxNumberOfMessages returns the MaxNumberOfMessages field value if set, zero value otherwise.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetMaxNumberOfMessages() int64 {
	if o == nil || o.MaxNumberOfMessages == nil {
		var ret int64
		return ret
	}
	return *o.MaxNumberOfMessages
}

// GetMaxNumberOfMessagesOk returns a tuple with the MaxNumberOfMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetMaxNumberOfMessagesOk() (*int64, bool) {
	if o == nil || o.MaxNumberOfMessages == nil {
		return nil, false
	}
	return o.MaxNumberOfMessages, true
}

// HasMaxNumberOfMessages returns a boolean if a field has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) HasMaxNumberOfMessages() bool {
	return o != nil && o.MaxNumberOfMessages != nil
}

// SetMaxNumberOfMessages gets a reference to the given int64 and assigns it to the MaxNumberOfMessages field.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) SetMaxNumberOfMessages(v int64) {
	o.MaxNumberOfMessages = &v
}

// GetRegion returns the Region field value.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) SetRegion(v string) {
	o.Region = v
}

// GetSecretAccessKeyKey returns the SecretAccessKeyKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetSecretAccessKeyKey() string {
	if o == nil || o.SecretAccessKeyKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKeyKey
}

// GetSecretAccessKeyKeyOk returns a tuple with the SecretAccessKeyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetSecretAccessKeyKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKeyKey == nil {
		return nil, false
	}
	return o.SecretAccessKeyKey, true
}

// HasSecretAccessKeyKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) HasSecretAccessKeyKey() bool {
	return o != nil && o.SecretAccessKeyKey != nil
}

// SetSecretAccessKeyKey gets a reference to the given string and assigns it to the SecretAccessKeyKey field.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) SetSecretAccessKeyKey(v string) {
	o.SecretAccessKeyKey = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetType() ObservabilityPipelineCrowdStrikeFdrSourceType {
	if o == nil {
		var ret ObservabilityPipelineCrowdStrikeFdrSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetTypeOk() (*ObservabilityPipelineCrowdStrikeFdrSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) SetType(v ObservabilityPipelineCrowdStrikeFdrSourceType) {
	o.Type = v
}

// GetUrlKey returns the UrlKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetUrlKey() string {
	if o == nil || o.UrlKey == nil {
		var ret string
		return ret
	}
	return *o.UrlKey
}

// GetUrlKeyOk returns a tuple with the UrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetUrlKeyOk() (*string, bool) {
	if o == nil || o.UrlKey == nil {
		return nil, false
	}
	return o.UrlKey, true
}

// HasUrlKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) HasUrlKey() bool {
	return o != nil && o.UrlKey != nil
}

// SetUrlKey gets a reference to the given string and assigns it to the UrlKey field.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) SetUrlKey(v string) {
	o.UrlKey = &v
}

// GetVisibilityTimeoutSecs returns the VisibilityTimeoutSecs field value if set, zero value otherwise.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetVisibilityTimeoutSecs() int64 {
	if o == nil || o.VisibilityTimeoutSecs == nil {
		var ret int64
		return ret
	}
	return *o.VisibilityTimeoutSecs
}

// GetVisibilityTimeoutSecsOk returns a tuple with the VisibilityTimeoutSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) GetVisibilityTimeoutSecsOk() (*int64, bool) {
	if o == nil || o.VisibilityTimeoutSecs == nil {
		return nil, false
	}
	return o.VisibilityTimeoutSecs, true
}

// HasVisibilityTimeoutSecs returns a boolean if a field has been set.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) HasVisibilityTimeoutSecs() bool {
	return o != nil && o.VisibilityTimeoutSecs != nil
}

// SetVisibilityTimeoutSecs gets a reference to the given int64 and assigns it to the VisibilityTimeoutSecs field.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) SetVisibilityTimeoutSecs(v int64) {
	o.VisibilityTimeoutSecs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineCrowdStrikeFdrSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccessKeyIdKey != nil {
		toSerialize["access_key_id_key"] = o.AccessKeyIdKey
	}
	if o.ClientConcurrency != nil {
		toSerialize["client_concurrency"] = o.ClientConcurrency
	}
	toSerialize["id"] = o.Id
	if o.MaxNumberOfMessages != nil {
		toSerialize["max_number_of_messages"] = o.MaxNumberOfMessages
	}
	toSerialize["region"] = o.Region
	if o.SecretAccessKeyKey != nil {
		toSerialize["secret_access_key_key"] = o.SecretAccessKeyKey
	}
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["type"] = o.Type
	if o.UrlKey != nil {
		toSerialize["url_key"] = o.UrlKey
	}
	if o.VisibilityTimeoutSecs != nil {
		toSerialize["visibility_timeout_secs"] = o.VisibilityTimeoutSecs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineCrowdStrikeFdrSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessKeyIdKey        *string                                        `json:"access_key_id_key,omitempty"`
		ClientConcurrency     *int64                                         `json:"client_concurrency,omitempty"`
		Id                    *string                                        `json:"id"`
		MaxNumberOfMessages   *int64                                         `json:"max_number_of_messages,omitempty"`
		Region                *string                                        `json:"region"`
		SecretAccessKeyKey    *string                                        `json:"secret_access_key_key,omitempty"`
		Tls                   *ObservabilityPipelineTls                      `json:"tls,omitempty"`
		Type                  *ObservabilityPipelineCrowdStrikeFdrSourceType `json:"type"`
		UrlKey                *string                                        `json:"url_key,omitempty"`
		VisibilityTimeoutSecs *int64                                         `json:"visibility_timeout_secs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_key_id_key", "client_concurrency", "id", "max_number_of_messages", "region", "secret_access_key_key", "tls", "type", "url_key", "visibility_timeout_secs"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccessKeyIdKey = all.AccessKeyIdKey
	o.ClientConcurrency = all.ClientConcurrency
	o.Id = *all.Id
	o.MaxNumberOfMessages = all.MaxNumberOfMessages
	o.Region = *all.Region
	o.SecretAccessKeyKey = all.SecretAccessKeyKey
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.UrlKey = all.UrlKey
	o.VisibilityTimeoutSecs = all.VisibilityTimeoutSecs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
