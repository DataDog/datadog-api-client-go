// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse Settings of the Data Observability dataflow.
type SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse struct {
	// Cron expression setting how often Datadog crawls your Snowflake table metadata.
	DoTableCrawlerCron *string `json:"do_table_crawler_cron,omitempty"`
	// Whether metadata from the Snowflake `SNOWFLAKE` system database is included in Data Observability alongside your own databases.
	SyncSnowflakeSystemDatabase *bool `json:"sync_snowflake_system_database,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse instantiates a new SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse() *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse {
	this := SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse{}
	return &this
}

// NewSnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponseWithDefaults instantiates a new SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponseWithDefaults() *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse {
	this := SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse{}
	return &this
}

// GetDoTableCrawlerCron returns the DoTableCrawlerCron field value if set, zero value otherwise.
func (o *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) GetDoTableCrawlerCron() string {
	if o == nil || o.DoTableCrawlerCron == nil {
		var ret string
		return ret
	}
	return *o.DoTableCrawlerCron
}

// GetDoTableCrawlerCronOk returns a tuple with the DoTableCrawlerCron field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) GetDoTableCrawlerCronOk() (*string, bool) {
	if o == nil || o.DoTableCrawlerCron == nil {
		return nil, false
	}
	return o.DoTableCrawlerCron, true
}

// HasDoTableCrawlerCron returns a boolean if a field has been set.
func (o *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) HasDoTableCrawlerCron() bool {
	return o != nil && o.DoTableCrawlerCron != nil
}

// SetDoTableCrawlerCron gets a reference to the given string and assigns it to the DoTableCrawlerCron field.
func (o *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) SetDoTableCrawlerCron(v string) {
	o.DoTableCrawlerCron = &v
}

// GetSyncSnowflakeSystemDatabase returns the SyncSnowflakeSystemDatabase field value if set, zero value otherwise.
func (o *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) GetSyncSnowflakeSystemDatabase() bool {
	if o == nil || o.SyncSnowflakeSystemDatabase == nil {
		var ret bool
		return ret
	}
	return *o.SyncSnowflakeSystemDatabase
}

// GetSyncSnowflakeSystemDatabaseOk returns a tuple with the SyncSnowflakeSystemDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) GetSyncSnowflakeSystemDatabaseOk() (*bool, bool) {
	if o == nil || o.SyncSnowflakeSystemDatabase == nil {
		return nil, false
	}
	return o.SyncSnowflakeSystemDatabase, true
}

// HasSyncSnowflakeSystemDatabase returns a boolean if a field has been set.
func (o *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) HasSyncSnowflakeSystemDatabase() bool {
	return o != nil && o.SyncSnowflakeSystemDatabase != nil
}

// SetSyncSnowflakeSystemDatabase gets a reference to the given bool and assigns it to the SyncSnowflakeSystemDatabase field.
func (o *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) SetSyncSnowflakeSystemDatabase(v bool) {
	o.SyncSnowflakeSystemDatabase = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DoTableCrawlerCron != nil {
		toSerialize["do_table_crawler_cron"] = o.DoTableCrawlerCron
	}
	if o.SyncSnowflakeSystemDatabase != nil {
		toSerialize["sync_snowflake_system_database"] = o.SyncSnowflakeSystemDatabase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DoTableCrawlerCron          *string `json:"do_table_crawler_cron,omitempty"`
		SyncSnowflakeSystemDatabase *bool   `json:"sync_snowflake_system_database,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"do_table_crawler_cron", "sync_snowflake_system_database"})
	} else {
		return err
	}
	o.DoTableCrawlerCron = all.DoTableCrawlerCron
	o.SyncSnowflakeSystemDatabase = all.SyncSnowflakeSystemDatabase

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
