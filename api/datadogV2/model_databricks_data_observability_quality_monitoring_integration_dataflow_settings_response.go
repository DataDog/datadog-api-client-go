// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse Settings of the Data Observability dataflow.
type DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse struct {
	// Cron expression setting how often Datadog connects to your Databricks warehouse to collect metadata. Currently, only hourly (`0 * * * *`) and daily (`0 0 * * *`) are supported.
	DoCrawlersCron *string `json:"do_crawlers_cron,omitempty"`
	// Whether metadata from the Databricks `system` catalog is included in Data Observability alongside your data catalogs.
	SyncSystemCatalog *bool `json:"sync_system_catalog,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse instantiates a new DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse() *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse {
	this := DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse{}
	return &this
}

// NewDatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponseWithDefaults instantiates a new DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponseWithDefaults() *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse {
	this := DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse{}
	return &this
}

// GetDoCrawlersCron returns the DoCrawlersCron field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) GetDoCrawlersCron() string {
	if o == nil || o.DoCrawlersCron == nil {
		var ret string
		return ret
	}
	return *o.DoCrawlersCron
}

// GetDoCrawlersCronOk returns a tuple with the DoCrawlersCron field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) GetDoCrawlersCronOk() (*string, bool) {
	if o == nil || o.DoCrawlersCron == nil {
		return nil, false
	}
	return o.DoCrawlersCron, true
}

// HasDoCrawlersCron returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) HasDoCrawlersCron() bool {
	return o != nil && o.DoCrawlersCron != nil
}

// SetDoCrawlersCron gets a reference to the given string and assigns it to the DoCrawlersCron field.
func (o *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) SetDoCrawlersCron(v string) {
	o.DoCrawlersCron = &v
}

// GetSyncSystemCatalog returns the SyncSystemCatalog field value if set, zero value otherwise.
func (o *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) GetSyncSystemCatalog() bool {
	if o == nil || o.SyncSystemCatalog == nil {
		var ret bool
		return ret
	}
	return *o.SyncSystemCatalog
}

// GetSyncSystemCatalogOk returns a tuple with the SyncSystemCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) GetSyncSystemCatalogOk() (*bool, bool) {
	if o == nil || o.SyncSystemCatalog == nil {
		return nil, false
	}
	return o.SyncSystemCatalog, true
}

// HasSyncSystemCatalog returns a boolean if a field has been set.
func (o *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) HasSyncSystemCatalog() bool {
	return o != nil && o.SyncSystemCatalog != nil
}

// SetSyncSystemCatalog gets a reference to the given bool and assigns it to the SyncSystemCatalog field.
func (o *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) SetSyncSystemCatalog(v bool) {
	o.SyncSystemCatalog = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DoCrawlersCron != nil {
		toSerialize["do_crawlers_cron"] = o.DoCrawlersCron
	}
	if o.SyncSystemCatalog != nil {
		toSerialize["sync_system_catalog"] = o.SyncSystemCatalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DoCrawlersCron    *string `json:"do_crawlers_cron,omitempty"`
		SyncSystemCatalog *bool   `json:"sync_system_catalog,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"do_crawlers_cron", "sync_system_catalog"})
	} else {
		return err
	}
	o.DoCrawlersCron = all.DoCrawlersCron
	o.SyncSystemCatalog = all.SyncSystemCatalog

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
