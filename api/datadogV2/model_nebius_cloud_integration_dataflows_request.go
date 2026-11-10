// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationDataflowsRequest Data Datadog collects from Nebius Cloud, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to stop or start it, and use `settings` to configure what it collects. Defaults listed on each dataflow apply when the account is created; on update, omitted fields keep their current values. Every dataflow except `nebius-cloud-cost-collection` runs as the service account in `authentication`, using its authorized key pair and the static API key; enabling one of them without those keys is rejected with a `422` response, and one enabled with keys that later expire is stored but collects no data.
type NebiusCloudIntegrationDataflowsRequest struct {
	// Cloud cost data read from a Nebius Object Storage bucket export and used by [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). Enabling it requires a bucket holding your cost export and a Nebius access key pair that can read it; a request that enables this dataflow without them is rejected with a `422` response.
	NebiusCloudCostCollection *NebiusCloudCostCollectionIntegrationDataflowRequest `json:"nebius-cloud-cost-collection,omitempty"`
	// Logs from your Nebius tenant: Kubernetes control plane and audit logs, tenant-level audit events, compute serial console logs, cloud apps and AI inference logs, PostgreSQL logs, and MLflow logs, collected from the services named in `settings.enabled_services`. `enabled_services` is the sole source of truth for what is collected; `enabled` set to `false` stops log collection, while `true` on its own changes nothing. Defaults to no services.
	NebiusCloudLogs *NebiusCloudLogsIntegrationDataflowRequest `json:"nebius-cloud-logs,omitempty"`
	// Metadata about resources in your Nebius tenant, such as compute instances and Kubernetes clusters.
	NebiusCloudResources *NebiusCloudResourcesIntegrationDataflowRequest `json:"nebius-cloud-resources,omitempty"`
	// Usage metrics tracked against your Nebius service quotas, such as `nebius_cloud.servicequota.percentage_used` and `nebius_cloud.servicequota.limit`.
	NebiusCloudServiceQuotaCollection *NebiusCloudServiceQuotaCollectionIntegrationDataflowRequest `json:"nebius-cloud-service-quota-collection,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationDataflowsRequest instantiates a new NebiusCloudIntegrationDataflowsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationDataflowsRequest() *NebiusCloudIntegrationDataflowsRequest {
	this := NebiusCloudIntegrationDataflowsRequest{}
	return &this
}

// NewNebiusCloudIntegrationDataflowsRequestWithDefaults instantiates a new NebiusCloudIntegrationDataflowsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationDataflowsRequestWithDefaults() *NebiusCloudIntegrationDataflowsRequest {
	this := NebiusCloudIntegrationDataflowsRequest{}
	return &this
}

// GetNebiusCloudCostCollection returns the NebiusCloudCostCollection field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationDataflowsRequest) GetNebiusCloudCostCollection() NebiusCloudCostCollectionIntegrationDataflowRequest {
	if o == nil || o.NebiusCloudCostCollection == nil {
		var ret NebiusCloudCostCollectionIntegrationDataflowRequest
		return ret
	}
	return *o.NebiusCloudCostCollection
}

// GetNebiusCloudCostCollectionOk returns a tuple with the NebiusCloudCostCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationDataflowsRequest) GetNebiusCloudCostCollectionOk() (*NebiusCloudCostCollectionIntegrationDataflowRequest, bool) {
	if o == nil || o.NebiusCloudCostCollection == nil {
		return nil, false
	}
	return o.NebiusCloudCostCollection, true
}

// HasNebiusCloudCostCollection returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationDataflowsRequest) HasNebiusCloudCostCollection() bool {
	return o != nil && o.NebiusCloudCostCollection != nil
}

// SetNebiusCloudCostCollection gets a reference to the given NebiusCloudCostCollectionIntegrationDataflowRequest and assigns it to the NebiusCloudCostCollection field.
func (o *NebiusCloudIntegrationDataflowsRequest) SetNebiusCloudCostCollection(v NebiusCloudCostCollectionIntegrationDataflowRequest) {
	o.NebiusCloudCostCollection = &v
}

// GetNebiusCloudLogs returns the NebiusCloudLogs field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationDataflowsRequest) GetNebiusCloudLogs() NebiusCloudLogsIntegrationDataflowRequest {
	if o == nil || o.NebiusCloudLogs == nil {
		var ret NebiusCloudLogsIntegrationDataflowRequest
		return ret
	}
	return *o.NebiusCloudLogs
}

// GetNebiusCloudLogsOk returns a tuple with the NebiusCloudLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationDataflowsRequest) GetNebiusCloudLogsOk() (*NebiusCloudLogsIntegrationDataflowRequest, bool) {
	if o == nil || o.NebiusCloudLogs == nil {
		return nil, false
	}
	return o.NebiusCloudLogs, true
}

// HasNebiusCloudLogs returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationDataflowsRequest) HasNebiusCloudLogs() bool {
	return o != nil && o.NebiusCloudLogs != nil
}

// SetNebiusCloudLogs gets a reference to the given NebiusCloudLogsIntegrationDataflowRequest and assigns it to the NebiusCloudLogs field.
func (o *NebiusCloudIntegrationDataflowsRequest) SetNebiusCloudLogs(v NebiusCloudLogsIntegrationDataflowRequest) {
	o.NebiusCloudLogs = &v
}

// GetNebiusCloudResources returns the NebiusCloudResources field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationDataflowsRequest) GetNebiusCloudResources() NebiusCloudResourcesIntegrationDataflowRequest {
	if o == nil || o.NebiusCloudResources == nil {
		var ret NebiusCloudResourcesIntegrationDataflowRequest
		return ret
	}
	return *o.NebiusCloudResources
}

// GetNebiusCloudResourcesOk returns a tuple with the NebiusCloudResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationDataflowsRequest) GetNebiusCloudResourcesOk() (*NebiusCloudResourcesIntegrationDataflowRequest, bool) {
	if o == nil || o.NebiusCloudResources == nil {
		return nil, false
	}
	return o.NebiusCloudResources, true
}

// HasNebiusCloudResources returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationDataflowsRequest) HasNebiusCloudResources() bool {
	return o != nil && o.NebiusCloudResources != nil
}

// SetNebiusCloudResources gets a reference to the given NebiusCloudResourcesIntegrationDataflowRequest and assigns it to the NebiusCloudResources field.
func (o *NebiusCloudIntegrationDataflowsRequest) SetNebiusCloudResources(v NebiusCloudResourcesIntegrationDataflowRequest) {
	o.NebiusCloudResources = &v
}

// GetNebiusCloudServiceQuotaCollection returns the NebiusCloudServiceQuotaCollection field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationDataflowsRequest) GetNebiusCloudServiceQuotaCollection() NebiusCloudServiceQuotaCollectionIntegrationDataflowRequest {
	if o == nil || o.NebiusCloudServiceQuotaCollection == nil {
		var ret NebiusCloudServiceQuotaCollectionIntegrationDataflowRequest
		return ret
	}
	return *o.NebiusCloudServiceQuotaCollection
}

// GetNebiusCloudServiceQuotaCollectionOk returns a tuple with the NebiusCloudServiceQuotaCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationDataflowsRequest) GetNebiusCloudServiceQuotaCollectionOk() (*NebiusCloudServiceQuotaCollectionIntegrationDataflowRequest, bool) {
	if o == nil || o.NebiusCloudServiceQuotaCollection == nil {
		return nil, false
	}
	return o.NebiusCloudServiceQuotaCollection, true
}

// HasNebiusCloudServiceQuotaCollection returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationDataflowsRequest) HasNebiusCloudServiceQuotaCollection() bool {
	return o != nil && o.NebiusCloudServiceQuotaCollection != nil
}

// SetNebiusCloudServiceQuotaCollection gets a reference to the given NebiusCloudServiceQuotaCollectionIntegrationDataflowRequest and assigns it to the NebiusCloudServiceQuotaCollection field.
func (o *NebiusCloudIntegrationDataflowsRequest) SetNebiusCloudServiceQuotaCollection(v NebiusCloudServiceQuotaCollectionIntegrationDataflowRequest) {
	o.NebiusCloudServiceQuotaCollection = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationDataflowsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NebiusCloudCostCollection != nil {
		toSerialize["nebius-cloud-cost-collection"] = o.NebiusCloudCostCollection
	}
	if o.NebiusCloudLogs != nil {
		toSerialize["nebius-cloud-logs"] = o.NebiusCloudLogs
	}
	if o.NebiusCloudResources != nil {
		toSerialize["nebius-cloud-resources"] = o.NebiusCloudResources
	}
	if o.NebiusCloudServiceQuotaCollection != nil {
		toSerialize["nebius-cloud-service-quota-collection"] = o.NebiusCloudServiceQuotaCollection
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NebiusCloudCostCollection         *NebiusCloudCostCollectionIntegrationDataflowRequest         `json:"nebius-cloud-cost-collection,omitempty"`
		NebiusCloudLogs                   *NebiusCloudLogsIntegrationDataflowRequest                   `json:"nebius-cloud-logs,omitempty"`
		NebiusCloudResources              *NebiusCloudResourcesIntegrationDataflowRequest              `json:"nebius-cloud-resources,omitempty"`
		NebiusCloudServiceQuotaCollection *NebiusCloudServiceQuotaCollectionIntegrationDataflowRequest `json:"nebius-cloud-service-quota-collection,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.NebiusCloudCostCollection != nil && all.NebiusCloudCostCollection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NebiusCloudCostCollection = all.NebiusCloudCostCollection
	if all.NebiusCloudLogs != nil && all.NebiusCloudLogs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NebiusCloudLogs = all.NebiusCloudLogs
	if all.NebiusCloudResources != nil && all.NebiusCloudResources.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NebiusCloudResources = all.NebiusCloudResources
	if all.NebiusCloudServiceQuotaCollection != nil && all.NebiusCloudServiceQuotaCollection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NebiusCloudServiceQuotaCollection = all.NebiusCloudServiceQuotaCollection

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
