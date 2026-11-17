// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoreWeaveCloudIntegrationDataflowsResponse Data Datadog collects from CoreWeave, keyed by dataflow id.
type CoreWeaveCloudIntegrationDataflowsResponse struct {
	// Service logs for CoreWeave-managed components in each CoreWeave region: the CKS control plane (such as `kube-apiserver` and `etcd`), CoreWeave add-on services in the `cw-*` and `kube-system` namespaces, SUNK, and node system logs such as `kubelet` and `containerd`. Logs from your own workloads and Kubernetes audit logs are not collected. Logs are sent with `source:coreweave`, with CoreWeave log labels such as `namespace`, `app`, and `systemd_unit` added as tags.
	CoreweaveCloudLogs *CoreWeaveCloudLogsIntegrationDataflowResponse `json:"coreweave-cloud-logs,omitempty"`
	// Metrics from your CoreWeave organization: CoreWeave service metrics collected through its PromQL-compatible Prometheus endpoint, plus the `coreweave.*` namespaces from CoreWeave products such as Kubernetes, Database, Distributed File Storage, Inference, and Object Storage. Compute metrics (`coreweave.bare_metal`, `coreweave.gpu`, `coreweave.nccl`, and `coreweave.node`) are in Preview. Container metrics are not collected; install the Datadog Agent on your clusters for those.
	CoreweaveCloudMetrics *CoreWeaveCloudMetricsIntegrationDataflowResponse `json:"coreweave-cloud-metrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCoreWeaveCloudIntegrationDataflowsResponse instantiates a new CoreWeaveCloudIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCoreWeaveCloudIntegrationDataflowsResponse() *CoreWeaveCloudIntegrationDataflowsResponse {
	this := CoreWeaveCloudIntegrationDataflowsResponse{}
	return &this
}

// NewCoreWeaveCloudIntegrationDataflowsResponseWithDefaults instantiates a new CoreWeaveCloudIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCoreWeaveCloudIntegrationDataflowsResponseWithDefaults() *CoreWeaveCloudIntegrationDataflowsResponse {
	this := CoreWeaveCloudIntegrationDataflowsResponse{}
	return &this
}

// GetCoreweaveCloudLogs returns the CoreweaveCloudLogs field value if set, zero value otherwise.
func (o *CoreWeaveCloudIntegrationDataflowsResponse) GetCoreweaveCloudLogs() CoreWeaveCloudLogsIntegrationDataflowResponse {
	if o == nil || o.CoreweaveCloudLogs == nil {
		var ret CoreWeaveCloudLogsIntegrationDataflowResponse
		return ret
	}
	return *o.CoreweaveCloudLogs
}

// GetCoreweaveCloudLogsOk returns a tuple with the CoreweaveCloudLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationDataflowsResponse) GetCoreweaveCloudLogsOk() (*CoreWeaveCloudLogsIntegrationDataflowResponse, bool) {
	if o == nil || o.CoreweaveCloudLogs == nil {
		return nil, false
	}
	return o.CoreweaveCloudLogs, true
}

// HasCoreweaveCloudLogs returns a boolean if a field has been set.
func (o *CoreWeaveCloudIntegrationDataflowsResponse) HasCoreweaveCloudLogs() bool {
	return o != nil && o.CoreweaveCloudLogs != nil
}

// SetCoreweaveCloudLogs gets a reference to the given CoreWeaveCloudLogsIntegrationDataflowResponse and assigns it to the CoreweaveCloudLogs field.
func (o *CoreWeaveCloudIntegrationDataflowsResponse) SetCoreweaveCloudLogs(v CoreWeaveCloudLogsIntegrationDataflowResponse) {
	o.CoreweaveCloudLogs = &v
}

// GetCoreweaveCloudMetrics returns the CoreweaveCloudMetrics field value if set, zero value otherwise.
func (o *CoreWeaveCloudIntegrationDataflowsResponse) GetCoreweaveCloudMetrics() CoreWeaveCloudMetricsIntegrationDataflowResponse {
	if o == nil || o.CoreweaveCloudMetrics == nil {
		var ret CoreWeaveCloudMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.CoreweaveCloudMetrics
}

// GetCoreweaveCloudMetricsOk returns a tuple with the CoreweaveCloudMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreWeaveCloudIntegrationDataflowsResponse) GetCoreweaveCloudMetricsOk() (*CoreWeaveCloudMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.CoreweaveCloudMetrics == nil {
		return nil, false
	}
	return o.CoreweaveCloudMetrics, true
}

// HasCoreweaveCloudMetrics returns a boolean if a field has been set.
func (o *CoreWeaveCloudIntegrationDataflowsResponse) HasCoreweaveCloudMetrics() bool {
	return o != nil && o.CoreweaveCloudMetrics != nil
}

// SetCoreweaveCloudMetrics gets a reference to the given CoreWeaveCloudMetricsIntegrationDataflowResponse and assigns it to the CoreweaveCloudMetrics field.
func (o *CoreWeaveCloudIntegrationDataflowsResponse) SetCoreweaveCloudMetrics(v CoreWeaveCloudMetricsIntegrationDataflowResponse) {
	o.CoreweaveCloudMetrics = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CoreWeaveCloudIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CoreweaveCloudLogs != nil {
		toSerialize["coreweave-cloud-logs"] = o.CoreweaveCloudLogs
	}
	if o.CoreweaveCloudMetrics != nil {
		toSerialize["coreweave-cloud-metrics"] = o.CoreweaveCloudMetrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CoreWeaveCloudIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CoreweaveCloudLogs    *CoreWeaveCloudLogsIntegrationDataflowResponse    `json:"coreweave-cloud-logs,omitempty"`
		CoreweaveCloudMetrics *CoreWeaveCloudMetricsIntegrationDataflowResponse `json:"coreweave-cloud-metrics,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"coreweave-cloud-logs", "coreweave-cloud-metrics"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CoreweaveCloudLogs != nil && all.CoreweaveCloudLogs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CoreweaveCloudLogs = all.CoreweaveCloudLogs
	if all.CoreweaveCloudMetrics != nil && all.CoreweaveCloudMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CoreweaveCloudMetrics = all.CoreweaveCloudMetrics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
