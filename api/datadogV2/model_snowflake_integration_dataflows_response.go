// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationDataflowsResponse Data Datadog collects from Snowflake, keyed by dataflow id.
type SnowflakeIntegrationDataflowsResponse struct {
	// Account-level usage metrics read from the Snowflake `ACCOUNT_USAGE` schema, covering storage usage, credit consumption, and query activity.
	SnowflakeAccountUsageMetrics *SnowflakeAccountUsageMetricsIntegrationDataflowResponse `json:"snowflake-account-usage-metrics,omitempty"`
	// Cost data aggregated from the Snowflake `ORGANIZATION_USAGE` schema. Requires [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/) to be set up for your organization, and the ORGANIZATION_BILLING_VIEWER database role on the Snowflake role.
	SnowflakeCloudCostMetrics *SnowflakeCloudCostMetricsIntegrationDataflowResponse `json:"snowflake-cloud-cost-metrics,omitempty"`
	// Data Observability, which collects lineage and data quality information from your Snowflake databases so you can explore how data flows and detect and resolve quality issues.
	SnowflakeDataObservabilityQualityMonitoring *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowResponse `json:"snowflake-data-observability-quality-monitoring,omitempty"`
	// Records from your Snowflake event tables, used to monitor application behavior and identify issues. `enabled` turns the dataflow on and off as a whole, and the per-record-type toggles in `settings` select which kinds of record it collects while it is on. The Snowflake role needs usage granted on the database, the schema, and the event table itself.
	SnowflakeEventTableLogs *SnowflakeEventTableLogsIntegrationDataflowResponse `json:"snowflake-event-table-logs,omitempty"`
	// Organization-level usage metrics read from the Snowflake `ORGANIZATION_USAGE` schema, covering the credit consumption of every account in the organization and the history of data transferred out of Snowflake. Reading that schema requires the ORGADMIN role.
	SnowflakeOrganizationUsageMetrics *SnowflakeOrganizationUsageMetricsIntegrationDataflowResponse `json:"snowflake-organization-usage-metrics,omitempty"`
	// Per-query logs that let you identify long-running, poorly performing, and expensive queries.
	SnowflakeQueryHistoryLogs *SnowflakeQueryHistoryLogsIntegrationDataflowResponse `json:"snowflake-query-history-logs,omitempty"`
	// Security logs from the Snowflake `ACCOUNT_USAGE` schema, for analyzing the security of your Snowflake account and running threat detection with [Cloud SIEM](https://docs.datadoghq.com/security/cloud_siem/).
	SnowflakeSecurityLogs *SnowflakeSecurityLogsIntegrationDataflowResponse `json:"snowflake-security-logs,omitempty"`
	// Execution logs for your scheduled Snowflake tasks, covering start and end time, status, and any error message.
	SnowflakeTaskHistoryLogs *SnowflakeTaskHistoryLogsIntegrationDataflowResponse `json:"snowflake-task-history-logs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSnowflakeIntegrationDataflowsResponse instantiates a new SnowflakeIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeIntegrationDataflowsResponse() *SnowflakeIntegrationDataflowsResponse {
	this := SnowflakeIntegrationDataflowsResponse{}
	return &this
}

// NewSnowflakeIntegrationDataflowsResponseWithDefaults instantiates a new SnowflakeIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeIntegrationDataflowsResponseWithDefaults() *SnowflakeIntegrationDataflowsResponse {
	this := SnowflakeIntegrationDataflowsResponse{}
	return &this
}

// GetSnowflakeAccountUsageMetrics returns the SnowflakeAccountUsageMetrics field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeAccountUsageMetrics() SnowflakeAccountUsageMetricsIntegrationDataflowResponse {
	if o == nil || o.SnowflakeAccountUsageMetrics == nil {
		var ret SnowflakeAccountUsageMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.SnowflakeAccountUsageMetrics
}

// GetSnowflakeAccountUsageMetricsOk returns a tuple with the SnowflakeAccountUsageMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeAccountUsageMetricsOk() (*SnowflakeAccountUsageMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.SnowflakeAccountUsageMetrics == nil {
		return nil, false
	}
	return o.SnowflakeAccountUsageMetrics, true
}

// HasSnowflakeAccountUsageMetrics returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsResponse) HasSnowflakeAccountUsageMetrics() bool {
	return o != nil && o.SnowflakeAccountUsageMetrics != nil
}

// SetSnowflakeAccountUsageMetrics gets a reference to the given SnowflakeAccountUsageMetricsIntegrationDataflowResponse and assigns it to the SnowflakeAccountUsageMetrics field.
func (o *SnowflakeIntegrationDataflowsResponse) SetSnowflakeAccountUsageMetrics(v SnowflakeAccountUsageMetricsIntegrationDataflowResponse) {
	o.SnowflakeAccountUsageMetrics = &v
}

// GetSnowflakeCloudCostMetrics returns the SnowflakeCloudCostMetrics field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeCloudCostMetrics() SnowflakeCloudCostMetricsIntegrationDataflowResponse {
	if o == nil || o.SnowflakeCloudCostMetrics == nil {
		var ret SnowflakeCloudCostMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.SnowflakeCloudCostMetrics
}

// GetSnowflakeCloudCostMetricsOk returns a tuple with the SnowflakeCloudCostMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeCloudCostMetricsOk() (*SnowflakeCloudCostMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.SnowflakeCloudCostMetrics == nil {
		return nil, false
	}
	return o.SnowflakeCloudCostMetrics, true
}

// HasSnowflakeCloudCostMetrics returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsResponse) HasSnowflakeCloudCostMetrics() bool {
	return o != nil && o.SnowflakeCloudCostMetrics != nil
}

// SetSnowflakeCloudCostMetrics gets a reference to the given SnowflakeCloudCostMetricsIntegrationDataflowResponse and assigns it to the SnowflakeCloudCostMetrics field.
func (o *SnowflakeIntegrationDataflowsResponse) SetSnowflakeCloudCostMetrics(v SnowflakeCloudCostMetricsIntegrationDataflowResponse) {
	o.SnowflakeCloudCostMetrics = &v
}

// GetSnowflakeDataObservabilityQualityMonitoring returns the SnowflakeDataObservabilityQualityMonitoring field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeDataObservabilityQualityMonitoring() SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowResponse {
	if o == nil || o.SnowflakeDataObservabilityQualityMonitoring == nil {
		var ret SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowResponse
		return ret
	}
	return *o.SnowflakeDataObservabilityQualityMonitoring
}

// GetSnowflakeDataObservabilityQualityMonitoringOk returns a tuple with the SnowflakeDataObservabilityQualityMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeDataObservabilityQualityMonitoringOk() (*SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowResponse, bool) {
	if o == nil || o.SnowflakeDataObservabilityQualityMonitoring == nil {
		return nil, false
	}
	return o.SnowflakeDataObservabilityQualityMonitoring, true
}

// HasSnowflakeDataObservabilityQualityMonitoring returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsResponse) HasSnowflakeDataObservabilityQualityMonitoring() bool {
	return o != nil && o.SnowflakeDataObservabilityQualityMonitoring != nil
}

// SetSnowflakeDataObservabilityQualityMonitoring gets a reference to the given SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowResponse and assigns it to the SnowflakeDataObservabilityQualityMonitoring field.
func (o *SnowflakeIntegrationDataflowsResponse) SetSnowflakeDataObservabilityQualityMonitoring(v SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowResponse) {
	o.SnowflakeDataObservabilityQualityMonitoring = &v
}

// GetSnowflakeEventTableLogs returns the SnowflakeEventTableLogs field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeEventTableLogs() SnowflakeEventTableLogsIntegrationDataflowResponse {
	if o == nil || o.SnowflakeEventTableLogs == nil {
		var ret SnowflakeEventTableLogsIntegrationDataflowResponse
		return ret
	}
	return *o.SnowflakeEventTableLogs
}

// GetSnowflakeEventTableLogsOk returns a tuple with the SnowflakeEventTableLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeEventTableLogsOk() (*SnowflakeEventTableLogsIntegrationDataflowResponse, bool) {
	if o == nil || o.SnowflakeEventTableLogs == nil {
		return nil, false
	}
	return o.SnowflakeEventTableLogs, true
}

// HasSnowflakeEventTableLogs returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsResponse) HasSnowflakeEventTableLogs() bool {
	return o != nil && o.SnowflakeEventTableLogs != nil
}

// SetSnowflakeEventTableLogs gets a reference to the given SnowflakeEventTableLogsIntegrationDataflowResponse and assigns it to the SnowflakeEventTableLogs field.
func (o *SnowflakeIntegrationDataflowsResponse) SetSnowflakeEventTableLogs(v SnowflakeEventTableLogsIntegrationDataflowResponse) {
	o.SnowflakeEventTableLogs = &v
}

// GetSnowflakeOrganizationUsageMetrics returns the SnowflakeOrganizationUsageMetrics field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeOrganizationUsageMetrics() SnowflakeOrganizationUsageMetricsIntegrationDataflowResponse {
	if o == nil || o.SnowflakeOrganizationUsageMetrics == nil {
		var ret SnowflakeOrganizationUsageMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.SnowflakeOrganizationUsageMetrics
}

// GetSnowflakeOrganizationUsageMetricsOk returns a tuple with the SnowflakeOrganizationUsageMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeOrganizationUsageMetricsOk() (*SnowflakeOrganizationUsageMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.SnowflakeOrganizationUsageMetrics == nil {
		return nil, false
	}
	return o.SnowflakeOrganizationUsageMetrics, true
}

// HasSnowflakeOrganizationUsageMetrics returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsResponse) HasSnowflakeOrganizationUsageMetrics() bool {
	return o != nil && o.SnowflakeOrganizationUsageMetrics != nil
}

// SetSnowflakeOrganizationUsageMetrics gets a reference to the given SnowflakeOrganizationUsageMetricsIntegrationDataflowResponse and assigns it to the SnowflakeOrganizationUsageMetrics field.
func (o *SnowflakeIntegrationDataflowsResponse) SetSnowflakeOrganizationUsageMetrics(v SnowflakeOrganizationUsageMetricsIntegrationDataflowResponse) {
	o.SnowflakeOrganizationUsageMetrics = &v
}

// GetSnowflakeQueryHistoryLogs returns the SnowflakeQueryHistoryLogs field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeQueryHistoryLogs() SnowflakeQueryHistoryLogsIntegrationDataflowResponse {
	if o == nil || o.SnowflakeQueryHistoryLogs == nil {
		var ret SnowflakeQueryHistoryLogsIntegrationDataflowResponse
		return ret
	}
	return *o.SnowflakeQueryHistoryLogs
}

// GetSnowflakeQueryHistoryLogsOk returns a tuple with the SnowflakeQueryHistoryLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeQueryHistoryLogsOk() (*SnowflakeQueryHistoryLogsIntegrationDataflowResponse, bool) {
	if o == nil || o.SnowflakeQueryHistoryLogs == nil {
		return nil, false
	}
	return o.SnowflakeQueryHistoryLogs, true
}

// HasSnowflakeQueryHistoryLogs returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsResponse) HasSnowflakeQueryHistoryLogs() bool {
	return o != nil && o.SnowflakeQueryHistoryLogs != nil
}

// SetSnowflakeQueryHistoryLogs gets a reference to the given SnowflakeQueryHistoryLogsIntegrationDataflowResponse and assigns it to the SnowflakeQueryHistoryLogs field.
func (o *SnowflakeIntegrationDataflowsResponse) SetSnowflakeQueryHistoryLogs(v SnowflakeQueryHistoryLogsIntegrationDataflowResponse) {
	o.SnowflakeQueryHistoryLogs = &v
}

// GetSnowflakeSecurityLogs returns the SnowflakeSecurityLogs field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeSecurityLogs() SnowflakeSecurityLogsIntegrationDataflowResponse {
	if o == nil || o.SnowflakeSecurityLogs == nil {
		var ret SnowflakeSecurityLogsIntegrationDataflowResponse
		return ret
	}
	return *o.SnowflakeSecurityLogs
}

// GetSnowflakeSecurityLogsOk returns a tuple with the SnowflakeSecurityLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeSecurityLogsOk() (*SnowflakeSecurityLogsIntegrationDataflowResponse, bool) {
	if o == nil || o.SnowflakeSecurityLogs == nil {
		return nil, false
	}
	return o.SnowflakeSecurityLogs, true
}

// HasSnowflakeSecurityLogs returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsResponse) HasSnowflakeSecurityLogs() bool {
	return o != nil && o.SnowflakeSecurityLogs != nil
}

// SetSnowflakeSecurityLogs gets a reference to the given SnowflakeSecurityLogsIntegrationDataflowResponse and assigns it to the SnowflakeSecurityLogs field.
func (o *SnowflakeIntegrationDataflowsResponse) SetSnowflakeSecurityLogs(v SnowflakeSecurityLogsIntegrationDataflowResponse) {
	o.SnowflakeSecurityLogs = &v
}

// GetSnowflakeTaskHistoryLogs returns the SnowflakeTaskHistoryLogs field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeTaskHistoryLogs() SnowflakeTaskHistoryLogsIntegrationDataflowResponse {
	if o == nil || o.SnowflakeTaskHistoryLogs == nil {
		var ret SnowflakeTaskHistoryLogsIntegrationDataflowResponse
		return ret
	}
	return *o.SnowflakeTaskHistoryLogs
}

// GetSnowflakeTaskHistoryLogsOk returns a tuple with the SnowflakeTaskHistoryLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsResponse) GetSnowflakeTaskHistoryLogsOk() (*SnowflakeTaskHistoryLogsIntegrationDataflowResponse, bool) {
	if o == nil || o.SnowflakeTaskHistoryLogs == nil {
		return nil, false
	}
	return o.SnowflakeTaskHistoryLogs, true
}

// HasSnowflakeTaskHistoryLogs returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsResponse) HasSnowflakeTaskHistoryLogs() bool {
	return o != nil && o.SnowflakeTaskHistoryLogs != nil
}

// SetSnowflakeTaskHistoryLogs gets a reference to the given SnowflakeTaskHistoryLogsIntegrationDataflowResponse and assigns it to the SnowflakeTaskHistoryLogs field.
func (o *SnowflakeIntegrationDataflowsResponse) SetSnowflakeTaskHistoryLogs(v SnowflakeTaskHistoryLogsIntegrationDataflowResponse) {
	o.SnowflakeTaskHistoryLogs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.SnowflakeAccountUsageMetrics != nil {
		toSerialize["snowflake-account-usage-metrics"] = o.SnowflakeAccountUsageMetrics
	}
	if o.SnowflakeCloudCostMetrics != nil {
		toSerialize["snowflake-cloud-cost-metrics"] = o.SnowflakeCloudCostMetrics
	}
	if o.SnowflakeDataObservabilityQualityMonitoring != nil {
		toSerialize["snowflake-data-observability-quality-monitoring"] = o.SnowflakeDataObservabilityQualityMonitoring
	}
	if o.SnowflakeEventTableLogs != nil {
		toSerialize["snowflake-event-table-logs"] = o.SnowflakeEventTableLogs
	}
	if o.SnowflakeOrganizationUsageMetrics != nil {
		toSerialize["snowflake-organization-usage-metrics"] = o.SnowflakeOrganizationUsageMetrics
	}
	if o.SnowflakeQueryHistoryLogs != nil {
		toSerialize["snowflake-query-history-logs"] = o.SnowflakeQueryHistoryLogs
	}
	if o.SnowflakeSecurityLogs != nil {
		toSerialize["snowflake-security-logs"] = o.SnowflakeSecurityLogs
	}
	if o.SnowflakeTaskHistoryLogs != nil {
		toSerialize["snowflake-task-history-logs"] = o.SnowflakeTaskHistoryLogs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SnowflakeAccountUsageMetrics                *SnowflakeAccountUsageMetricsIntegrationDataflowResponse                `json:"snowflake-account-usage-metrics,omitempty"`
		SnowflakeCloudCostMetrics                   *SnowflakeCloudCostMetricsIntegrationDataflowResponse                   `json:"snowflake-cloud-cost-metrics,omitempty"`
		SnowflakeDataObservabilityQualityMonitoring *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowResponse `json:"snowflake-data-observability-quality-monitoring,omitempty"`
		SnowflakeEventTableLogs                     *SnowflakeEventTableLogsIntegrationDataflowResponse                     `json:"snowflake-event-table-logs,omitempty"`
		SnowflakeOrganizationUsageMetrics           *SnowflakeOrganizationUsageMetricsIntegrationDataflowResponse           `json:"snowflake-organization-usage-metrics,omitempty"`
		SnowflakeQueryHistoryLogs                   *SnowflakeQueryHistoryLogsIntegrationDataflowResponse                   `json:"snowflake-query-history-logs,omitempty"`
		SnowflakeSecurityLogs                       *SnowflakeSecurityLogsIntegrationDataflowResponse                       `json:"snowflake-security-logs,omitempty"`
		SnowflakeTaskHistoryLogs                    *SnowflakeTaskHistoryLogsIntegrationDataflowResponse                    `json:"snowflake-task-history-logs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"snowflake-account-usage-metrics", "snowflake-cloud-cost-metrics", "snowflake-data-observability-quality-monitoring", "snowflake-event-table-logs", "snowflake-organization-usage-metrics", "snowflake-query-history-logs", "snowflake-security-logs", "snowflake-task-history-logs"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.SnowflakeAccountUsageMetrics != nil && all.SnowflakeAccountUsageMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SnowflakeAccountUsageMetrics = all.SnowflakeAccountUsageMetrics
	if all.SnowflakeCloudCostMetrics != nil && all.SnowflakeCloudCostMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SnowflakeCloudCostMetrics = all.SnowflakeCloudCostMetrics
	if all.SnowflakeDataObservabilityQualityMonitoring != nil && all.SnowflakeDataObservabilityQualityMonitoring.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SnowflakeDataObservabilityQualityMonitoring = all.SnowflakeDataObservabilityQualityMonitoring
	if all.SnowflakeEventTableLogs != nil && all.SnowflakeEventTableLogs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SnowflakeEventTableLogs = all.SnowflakeEventTableLogs
	if all.SnowflakeOrganizationUsageMetrics != nil && all.SnowflakeOrganizationUsageMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SnowflakeOrganizationUsageMetrics = all.SnowflakeOrganizationUsageMetrics
	if all.SnowflakeQueryHistoryLogs != nil && all.SnowflakeQueryHistoryLogs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SnowflakeQueryHistoryLogs = all.SnowflakeQueryHistoryLogs
	if all.SnowflakeSecurityLogs != nil && all.SnowflakeSecurityLogs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SnowflakeSecurityLogs = all.SnowflakeSecurityLogs
	if all.SnowflakeTaskHistoryLogs != nil && all.SnowflakeTaskHistoryLogs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SnowflakeTaskHistoryLogs = all.SnowflakeTaskHistoryLogs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
