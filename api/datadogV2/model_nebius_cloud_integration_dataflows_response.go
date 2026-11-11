// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudIntegrationDataflowsResponse Data Datadog collects from Nebius Cloud, keyed by dataflow id.
type NebiusCloudIntegrationDataflowsResponse struct {
	// Cloud cost data read from a Nebius Object Storage bucket export and used by [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). Requires a bucket holding your cost export and a Nebius access key pair that can read it; without both this dataflow collects no data.
	NebiusCloudCostCollection *NebiusCloudCostCollectionIntegrationDataflowResponse `json:"nebius-cloud-cost-collection,omitempty"`
	// Logs from your Nebius tenant: Kubernetes control plane and audit logs, tenant-level audit events, compute serial console logs, cloud apps and AI inference logs, PostgreSQL logs, and MLflow logs, collected from the services named in `settings.enabled_services`.
	NebiusCloudLogs *NebiusCloudLogsIntegrationDataflowResponse `json:"nebius-cloud-logs,omitempty"`
	// Metadata about resources in your Nebius tenant, such as compute instances and Kubernetes clusters.
	NebiusCloudResources *NebiusCloudResourcesIntegrationDataflowResponse `json:"nebius-cloud-resources,omitempty"`
	// Usage metrics tracked against your Nebius service quotas, such as `nebius_cloud.servicequota.percentage_used` and `nebius_cloud.servicequota.limit`.
	NebiusCloudServiceQuotaCollection *NebiusCloudServiceQuotaCollectionIntegrationDataflowResponse `json:"nebius-cloud-service-quota-collection,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNebiusCloudIntegrationDataflowsResponse instantiates a new NebiusCloudIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudIntegrationDataflowsResponse() *NebiusCloudIntegrationDataflowsResponse {
	this := NebiusCloudIntegrationDataflowsResponse{}
	return &this
}

// NewNebiusCloudIntegrationDataflowsResponseWithDefaults instantiates a new NebiusCloudIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudIntegrationDataflowsResponseWithDefaults() *NebiusCloudIntegrationDataflowsResponse {
	this := NebiusCloudIntegrationDataflowsResponse{}
	return &this
}

// GetNebiusCloudCostCollection returns the NebiusCloudCostCollection field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationDataflowsResponse) GetNebiusCloudCostCollection() NebiusCloudCostCollectionIntegrationDataflowResponse {
	if o == nil || o.NebiusCloudCostCollection == nil {
		var ret NebiusCloudCostCollectionIntegrationDataflowResponse
		return ret
	}
	return *o.NebiusCloudCostCollection
}

// GetNebiusCloudCostCollectionOk returns a tuple with the NebiusCloudCostCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationDataflowsResponse) GetNebiusCloudCostCollectionOk() (*NebiusCloudCostCollectionIntegrationDataflowResponse, bool) {
	if o == nil || o.NebiusCloudCostCollection == nil {
		return nil, false
	}
	return o.NebiusCloudCostCollection, true
}

// HasNebiusCloudCostCollection returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationDataflowsResponse) HasNebiusCloudCostCollection() bool {
	return o != nil && o.NebiusCloudCostCollection != nil
}

// SetNebiusCloudCostCollection gets a reference to the given NebiusCloudCostCollectionIntegrationDataflowResponse and assigns it to the NebiusCloudCostCollection field.
func (o *NebiusCloudIntegrationDataflowsResponse) SetNebiusCloudCostCollection(v NebiusCloudCostCollectionIntegrationDataflowResponse) {
	o.NebiusCloudCostCollection = &v
}

// GetNebiusCloudLogs returns the NebiusCloudLogs field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationDataflowsResponse) GetNebiusCloudLogs() NebiusCloudLogsIntegrationDataflowResponse {
	if o == nil || o.NebiusCloudLogs == nil {
		var ret NebiusCloudLogsIntegrationDataflowResponse
		return ret
	}
	return *o.NebiusCloudLogs
}

// GetNebiusCloudLogsOk returns a tuple with the NebiusCloudLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationDataflowsResponse) GetNebiusCloudLogsOk() (*NebiusCloudLogsIntegrationDataflowResponse, bool) {
	if o == nil || o.NebiusCloudLogs == nil {
		return nil, false
	}
	return o.NebiusCloudLogs, true
}

// HasNebiusCloudLogs returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationDataflowsResponse) HasNebiusCloudLogs() bool {
	return o != nil && o.NebiusCloudLogs != nil
}

// SetNebiusCloudLogs gets a reference to the given NebiusCloudLogsIntegrationDataflowResponse and assigns it to the NebiusCloudLogs field.
func (o *NebiusCloudIntegrationDataflowsResponse) SetNebiusCloudLogs(v NebiusCloudLogsIntegrationDataflowResponse) {
	o.NebiusCloudLogs = &v
}

// GetNebiusCloudResources returns the NebiusCloudResources field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationDataflowsResponse) GetNebiusCloudResources() NebiusCloudResourcesIntegrationDataflowResponse {
	if o == nil || o.NebiusCloudResources == nil {
		var ret NebiusCloudResourcesIntegrationDataflowResponse
		return ret
	}
	return *o.NebiusCloudResources
}

// GetNebiusCloudResourcesOk returns a tuple with the NebiusCloudResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationDataflowsResponse) GetNebiusCloudResourcesOk() (*NebiusCloudResourcesIntegrationDataflowResponse, bool) {
	if o == nil || o.NebiusCloudResources == nil {
		return nil, false
	}
	return o.NebiusCloudResources, true
}

// HasNebiusCloudResources returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationDataflowsResponse) HasNebiusCloudResources() bool {
	return o != nil && o.NebiusCloudResources != nil
}

// SetNebiusCloudResources gets a reference to the given NebiusCloudResourcesIntegrationDataflowResponse and assigns it to the NebiusCloudResources field.
func (o *NebiusCloudIntegrationDataflowsResponse) SetNebiusCloudResources(v NebiusCloudResourcesIntegrationDataflowResponse) {
	o.NebiusCloudResources = &v
}

// GetNebiusCloudServiceQuotaCollection returns the NebiusCloudServiceQuotaCollection field value if set, zero value otherwise.
func (o *NebiusCloudIntegrationDataflowsResponse) GetNebiusCloudServiceQuotaCollection() NebiusCloudServiceQuotaCollectionIntegrationDataflowResponse {
	if o == nil || o.NebiusCloudServiceQuotaCollection == nil {
		var ret NebiusCloudServiceQuotaCollectionIntegrationDataflowResponse
		return ret
	}
	return *o.NebiusCloudServiceQuotaCollection
}

// GetNebiusCloudServiceQuotaCollectionOk returns a tuple with the NebiusCloudServiceQuotaCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudIntegrationDataflowsResponse) GetNebiusCloudServiceQuotaCollectionOk() (*NebiusCloudServiceQuotaCollectionIntegrationDataflowResponse, bool) {
	if o == nil || o.NebiusCloudServiceQuotaCollection == nil {
		return nil, false
	}
	return o.NebiusCloudServiceQuotaCollection, true
}

// HasNebiusCloudServiceQuotaCollection returns a boolean if a field has been set.
func (o *NebiusCloudIntegrationDataflowsResponse) HasNebiusCloudServiceQuotaCollection() bool {
	return o != nil && o.NebiusCloudServiceQuotaCollection != nil
}

// SetNebiusCloudServiceQuotaCollection gets a reference to the given NebiusCloudServiceQuotaCollectionIntegrationDataflowResponse and assigns it to the NebiusCloudServiceQuotaCollection field.
func (o *NebiusCloudIntegrationDataflowsResponse) SetNebiusCloudServiceQuotaCollection(v NebiusCloudServiceQuotaCollectionIntegrationDataflowResponse) {
	o.NebiusCloudServiceQuotaCollection = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NebiusCloudCostCollection         *NebiusCloudCostCollectionIntegrationDataflowResponse         `json:"nebius-cloud-cost-collection,omitempty"`
		NebiusCloudLogs                   *NebiusCloudLogsIntegrationDataflowResponse                   `json:"nebius-cloud-logs,omitempty"`
		NebiusCloudResources              *NebiusCloudResourcesIntegrationDataflowResponse              `json:"nebius-cloud-resources,omitempty"`
		NebiusCloudServiceQuotaCollection *NebiusCloudServiceQuotaCollectionIntegrationDataflowResponse `json:"nebius-cloud-service-quota-collection,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"nebius-cloud-cost-collection", "nebius-cloud-logs", "nebius-cloud-resources", "nebius-cloud-service-quota-collection"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
