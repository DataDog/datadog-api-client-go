// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeIntegrationDataflowsRequest Data Datadog collects from Snowflake, keyed by dataflow id. Each dataflow turns on a distinct kind of collection: set `enabled` to start or stop it, and use `settings` to configure what it collects. Defaults listed on each dataflow apply when the account is created; on update, omitted fields keep their current values. Every dataflow reads from Snowflake as the user in `settings.username`, so that user's role must be granted access to the underlying views; a dataflow enabled without those grants is stored but collects no data.
type SnowflakeIntegrationDataflowsRequest struct {
	// Account-level usage metrics read from the Snowflake `ACCOUNT_USAGE` schema, covering storage usage, credit consumption, and query activity.
	SnowflakeAccountUsageMetrics *SnowflakeAccountUsageMetricsIntegrationDataflowRequest `json:"snowflake-account-usage-metrics,omitempty"`
	// Cost data aggregated from the Snowflake `ORGANIZATION_USAGE` schema. [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/) must be enabled for your organization while this dataflow is enabled. Any request that enables this dataflow without Cloud Cost Management is rejected with a `422` response. The Snowflake role also needs the ORGANIZATION_BILLING_VIEWER database role to read the underlying cost views.
	SnowflakeCloudCostMetrics *SnowflakeCloudCostMetricsIntegrationDataflowRequest `json:"snowflake-cloud-cost-metrics,omitempty"`
	// Data Observability, which collects lineage and data quality information from your Snowflake databases so you can explore how data flows and detect and resolve quality issues.
	SnowflakeDataObservabilityQualityMonitoring *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowRequest `json:"snowflake-data-observability-quality-monitoring,omitempty"`
	// Records from your Snowflake event tables, used to monitor application behavior and identify issues. `enabled` turns the dataflow on and off as a whole, and the per-record-type toggles in `settings` select which kinds of record it collects while it is on. The Snowflake role needs usage granted on the database, the schema, and the event table itself.
	SnowflakeEventTableLogs *SnowflakeEventTableLogsIntegrationDataflowRequest `json:"snowflake-event-table-logs,omitempty"`
	// Organization-level usage metrics read from the Snowflake `ORGANIZATION_USAGE` schema, covering the credit consumption of every account in the organization and the history of data transferred out of Snowflake. Reading that schema requires the ORGADMIN role.
	SnowflakeOrganizationUsageMetrics *SnowflakeOrganizationUsageMetricsIntegrationDataflowRequest `json:"snowflake-organization-usage-metrics,omitempty"`
	// Per-query logs that let you identify long-running, poorly performing, and expensive queries.
	SnowflakeQueryHistoryLogs *SnowflakeQueryHistoryLogsIntegrationDataflowRequest `json:"snowflake-query-history-logs,omitempty"`
	// Security logs from the Snowflake `ACCOUNT_USAGE` schema, for analyzing the security of your Snowflake account and running threat detection with [Cloud SIEM](https://docs.datadoghq.com/security/cloud_siem/).
	SnowflakeSecurityLogs *SnowflakeSecurityLogsIntegrationDataflowRequest `json:"snowflake-security-logs,omitempty"`
	// Execution logs for your scheduled Snowflake tasks, covering start and end time, status, and any error message.
	SnowflakeTaskHistoryLogs *SnowflakeTaskHistoryLogsIntegrationDataflowRequest `json:"snowflake-task-history-logs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeIntegrationDataflowsRequest instantiates a new SnowflakeIntegrationDataflowsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeIntegrationDataflowsRequest() *SnowflakeIntegrationDataflowsRequest {
	this := SnowflakeIntegrationDataflowsRequest{}
	return &this
}

// NewSnowflakeIntegrationDataflowsRequestWithDefaults instantiates a new SnowflakeIntegrationDataflowsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeIntegrationDataflowsRequestWithDefaults() *SnowflakeIntegrationDataflowsRequest {
	this := SnowflakeIntegrationDataflowsRequest{}
	return &this
}

// GetSnowflakeAccountUsageMetrics returns the SnowflakeAccountUsageMetrics field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeAccountUsageMetrics() SnowflakeAccountUsageMetricsIntegrationDataflowRequest {
	if o == nil || o.SnowflakeAccountUsageMetrics == nil {
		var ret SnowflakeAccountUsageMetricsIntegrationDataflowRequest
		return ret
	}
	return *o.SnowflakeAccountUsageMetrics
}

// GetSnowflakeAccountUsageMetricsOk returns a tuple with the SnowflakeAccountUsageMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeAccountUsageMetricsOk() (*SnowflakeAccountUsageMetricsIntegrationDataflowRequest, bool) {
	if o == nil || o.SnowflakeAccountUsageMetrics == nil {
		return nil, false
	}
	return o.SnowflakeAccountUsageMetrics, true
}

// HasSnowflakeAccountUsageMetrics returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsRequest) HasSnowflakeAccountUsageMetrics() bool {
	return o != nil && o.SnowflakeAccountUsageMetrics != nil
}

// SetSnowflakeAccountUsageMetrics gets a reference to the given SnowflakeAccountUsageMetricsIntegrationDataflowRequest and assigns it to the SnowflakeAccountUsageMetrics field.
func (o *SnowflakeIntegrationDataflowsRequest) SetSnowflakeAccountUsageMetrics(v SnowflakeAccountUsageMetricsIntegrationDataflowRequest) {
	o.SnowflakeAccountUsageMetrics = &v
}

// GetSnowflakeCloudCostMetrics returns the SnowflakeCloudCostMetrics field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeCloudCostMetrics() SnowflakeCloudCostMetricsIntegrationDataflowRequest {
	if o == nil || o.SnowflakeCloudCostMetrics == nil {
		var ret SnowflakeCloudCostMetricsIntegrationDataflowRequest
		return ret
	}
	return *o.SnowflakeCloudCostMetrics
}

// GetSnowflakeCloudCostMetricsOk returns a tuple with the SnowflakeCloudCostMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeCloudCostMetricsOk() (*SnowflakeCloudCostMetricsIntegrationDataflowRequest, bool) {
	if o == nil || o.SnowflakeCloudCostMetrics == nil {
		return nil, false
	}
	return o.SnowflakeCloudCostMetrics, true
}

// HasSnowflakeCloudCostMetrics returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsRequest) HasSnowflakeCloudCostMetrics() bool {
	return o != nil && o.SnowflakeCloudCostMetrics != nil
}

// SetSnowflakeCloudCostMetrics gets a reference to the given SnowflakeCloudCostMetricsIntegrationDataflowRequest and assigns it to the SnowflakeCloudCostMetrics field.
func (o *SnowflakeIntegrationDataflowsRequest) SetSnowflakeCloudCostMetrics(v SnowflakeCloudCostMetricsIntegrationDataflowRequest) {
	o.SnowflakeCloudCostMetrics = &v
}

// GetSnowflakeDataObservabilityQualityMonitoring returns the SnowflakeDataObservabilityQualityMonitoring field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeDataObservabilityQualityMonitoring() SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowRequest {
	if o == nil || o.SnowflakeDataObservabilityQualityMonitoring == nil {
		var ret SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowRequest
		return ret
	}
	return *o.SnowflakeDataObservabilityQualityMonitoring
}

// GetSnowflakeDataObservabilityQualityMonitoringOk returns a tuple with the SnowflakeDataObservabilityQualityMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeDataObservabilityQualityMonitoringOk() (*SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowRequest, bool) {
	if o == nil || o.SnowflakeDataObservabilityQualityMonitoring == nil {
		return nil, false
	}
	return o.SnowflakeDataObservabilityQualityMonitoring, true
}

// HasSnowflakeDataObservabilityQualityMonitoring returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsRequest) HasSnowflakeDataObservabilityQualityMonitoring() bool {
	return o != nil && o.SnowflakeDataObservabilityQualityMonitoring != nil
}

// SetSnowflakeDataObservabilityQualityMonitoring gets a reference to the given SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowRequest and assigns it to the SnowflakeDataObservabilityQualityMonitoring field.
func (o *SnowflakeIntegrationDataflowsRequest) SetSnowflakeDataObservabilityQualityMonitoring(v SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowRequest) {
	o.SnowflakeDataObservabilityQualityMonitoring = &v
}

// GetSnowflakeEventTableLogs returns the SnowflakeEventTableLogs field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeEventTableLogs() SnowflakeEventTableLogsIntegrationDataflowRequest {
	if o == nil || o.SnowflakeEventTableLogs == nil {
		var ret SnowflakeEventTableLogsIntegrationDataflowRequest
		return ret
	}
	return *o.SnowflakeEventTableLogs
}

// GetSnowflakeEventTableLogsOk returns a tuple with the SnowflakeEventTableLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeEventTableLogsOk() (*SnowflakeEventTableLogsIntegrationDataflowRequest, bool) {
	if o == nil || o.SnowflakeEventTableLogs == nil {
		return nil, false
	}
	return o.SnowflakeEventTableLogs, true
}

// HasSnowflakeEventTableLogs returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsRequest) HasSnowflakeEventTableLogs() bool {
	return o != nil && o.SnowflakeEventTableLogs != nil
}

// SetSnowflakeEventTableLogs gets a reference to the given SnowflakeEventTableLogsIntegrationDataflowRequest and assigns it to the SnowflakeEventTableLogs field.
func (o *SnowflakeIntegrationDataflowsRequest) SetSnowflakeEventTableLogs(v SnowflakeEventTableLogsIntegrationDataflowRequest) {
	o.SnowflakeEventTableLogs = &v
}

// GetSnowflakeOrganizationUsageMetrics returns the SnowflakeOrganizationUsageMetrics field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeOrganizationUsageMetrics() SnowflakeOrganizationUsageMetricsIntegrationDataflowRequest {
	if o == nil || o.SnowflakeOrganizationUsageMetrics == nil {
		var ret SnowflakeOrganizationUsageMetricsIntegrationDataflowRequest
		return ret
	}
	return *o.SnowflakeOrganizationUsageMetrics
}

// GetSnowflakeOrganizationUsageMetricsOk returns a tuple with the SnowflakeOrganizationUsageMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeOrganizationUsageMetricsOk() (*SnowflakeOrganizationUsageMetricsIntegrationDataflowRequest, bool) {
	if o == nil || o.SnowflakeOrganizationUsageMetrics == nil {
		return nil, false
	}
	return o.SnowflakeOrganizationUsageMetrics, true
}

// HasSnowflakeOrganizationUsageMetrics returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsRequest) HasSnowflakeOrganizationUsageMetrics() bool {
	return o != nil && o.SnowflakeOrganizationUsageMetrics != nil
}

// SetSnowflakeOrganizationUsageMetrics gets a reference to the given SnowflakeOrganizationUsageMetricsIntegrationDataflowRequest and assigns it to the SnowflakeOrganizationUsageMetrics field.
func (o *SnowflakeIntegrationDataflowsRequest) SetSnowflakeOrganizationUsageMetrics(v SnowflakeOrganizationUsageMetricsIntegrationDataflowRequest) {
	o.SnowflakeOrganizationUsageMetrics = &v
}

// GetSnowflakeQueryHistoryLogs returns the SnowflakeQueryHistoryLogs field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeQueryHistoryLogs() SnowflakeQueryHistoryLogsIntegrationDataflowRequest {
	if o == nil || o.SnowflakeQueryHistoryLogs == nil {
		var ret SnowflakeQueryHistoryLogsIntegrationDataflowRequest
		return ret
	}
	return *o.SnowflakeQueryHistoryLogs
}

// GetSnowflakeQueryHistoryLogsOk returns a tuple with the SnowflakeQueryHistoryLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeQueryHistoryLogsOk() (*SnowflakeQueryHistoryLogsIntegrationDataflowRequest, bool) {
	if o == nil || o.SnowflakeQueryHistoryLogs == nil {
		return nil, false
	}
	return o.SnowflakeQueryHistoryLogs, true
}

// HasSnowflakeQueryHistoryLogs returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsRequest) HasSnowflakeQueryHistoryLogs() bool {
	return o != nil && o.SnowflakeQueryHistoryLogs != nil
}

// SetSnowflakeQueryHistoryLogs gets a reference to the given SnowflakeQueryHistoryLogsIntegrationDataflowRequest and assigns it to the SnowflakeQueryHistoryLogs field.
func (o *SnowflakeIntegrationDataflowsRequest) SetSnowflakeQueryHistoryLogs(v SnowflakeQueryHistoryLogsIntegrationDataflowRequest) {
	o.SnowflakeQueryHistoryLogs = &v
}

// GetSnowflakeSecurityLogs returns the SnowflakeSecurityLogs field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeSecurityLogs() SnowflakeSecurityLogsIntegrationDataflowRequest {
	if o == nil || o.SnowflakeSecurityLogs == nil {
		var ret SnowflakeSecurityLogsIntegrationDataflowRequest
		return ret
	}
	return *o.SnowflakeSecurityLogs
}

// GetSnowflakeSecurityLogsOk returns a tuple with the SnowflakeSecurityLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeSecurityLogsOk() (*SnowflakeSecurityLogsIntegrationDataflowRequest, bool) {
	if o == nil || o.SnowflakeSecurityLogs == nil {
		return nil, false
	}
	return o.SnowflakeSecurityLogs, true
}

// HasSnowflakeSecurityLogs returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsRequest) HasSnowflakeSecurityLogs() bool {
	return o != nil && o.SnowflakeSecurityLogs != nil
}

// SetSnowflakeSecurityLogs gets a reference to the given SnowflakeSecurityLogsIntegrationDataflowRequest and assigns it to the SnowflakeSecurityLogs field.
func (o *SnowflakeIntegrationDataflowsRequest) SetSnowflakeSecurityLogs(v SnowflakeSecurityLogsIntegrationDataflowRequest) {
	o.SnowflakeSecurityLogs = &v
}

// GetSnowflakeTaskHistoryLogs returns the SnowflakeTaskHistoryLogs field value if set, zero value otherwise.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeTaskHistoryLogs() SnowflakeTaskHistoryLogsIntegrationDataflowRequest {
	if o == nil || o.SnowflakeTaskHistoryLogs == nil {
		var ret SnowflakeTaskHistoryLogsIntegrationDataflowRequest
		return ret
	}
	return *o.SnowflakeTaskHistoryLogs
}

// GetSnowflakeTaskHistoryLogsOk returns a tuple with the SnowflakeTaskHistoryLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeIntegrationDataflowsRequest) GetSnowflakeTaskHistoryLogsOk() (*SnowflakeTaskHistoryLogsIntegrationDataflowRequest, bool) {
	if o == nil || o.SnowflakeTaskHistoryLogs == nil {
		return nil, false
	}
	return o.SnowflakeTaskHistoryLogs, true
}

// HasSnowflakeTaskHistoryLogs returns a boolean if a field has been set.
func (o *SnowflakeIntegrationDataflowsRequest) HasSnowflakeTaskHistoryLogs() bool {
	return o != nil && o.SnowflakeTaskHistoryLogs != nil
}

// SetSnowflakeTaskHistoryLogs gets a reference to the given SnowflakeTaskHistoryLogsIntegrationDataflowRequest and assigns it to the SnowflakeTaskHistoryLogs field.
func (o *SnowflakeIntegrationDataflowsRequest) SetSnowflakeTaskHistoryLogs(v SnowflakeTaskHistoryLogsIntegrationDataflowRequest) {
	o.SnowflakeTaskHistoryLogs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeIntegrationDataflowsRequest) MarshalJSON() ([]byte, error) {
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeIntegrationDataflowsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SnowflakeAccountUsageMetrics                *SnowflakeAccountUsageMetricsIntegrationDataflowRequest                `json:"snowflake-account-usage-metrics,omitempty"`
		SnowflakeCloudCostMetrics                   *SnowflakeCloudCostMetricsIntegrationDataflowRequest                   `json:"snowflake-cloud-cost-metrics,omitempty"`
		SnowflakeDataObservabilityQualityMonitoring *SnowflakeDataObservabilityQualityMonitoringIntegrationDataflowRequest `json:"snowflake-data-observability-quality-monitoring,omitempty"`
		SnowflakeEventTableLogs                     *SnowflakeEventTableLogsIntegrationDataflowRequest                     `json:"snowflake-event-table-logs,omitempty"`
		SnowflakeOrganizationUsageMetrics           *SnowflakeOrganizationUsageMetricsIntegrationDataflowRequest           `json:"snowflake-organization-usage-metrics,omitempty"`
		SnowflakeQueryHistoryLogs                   *SnowflakeQueryHistoryLogsIntegrationDataflowRequest                   `json:"snowflake-query-history-logs,omitempty"`
		SnowflakeSecurityLogs                       *SnowflakeSecurityLogsIntegrationDataflowRequest                       `json:"snowflake-security-logs,omitempty"`
		SnowflakeTaskHistoryLogs                    *SnowflakeTaskHistoryLogsIntegrationDataflowRequest                    `json:"snowflake-task-history-logs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
