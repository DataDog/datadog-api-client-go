// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyUsageAttributionSupportedMetrics Supported metrics for monthly usage attribution requests.
type MonthlyUsageAttributionSupportedMetrics string

// List of MonthlyUsageAttributionSupportedMetrics.
const (
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_API_USAGE                                    MonthlyUsageAttributionSupportedMetrics = "api_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_API_PERCENTAGE                               MonthlyUsageAttributionSupportedMetrics = "api_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_FARGATE_USAGE                            MonthlyUsageAttributionSupportedMetrics = "apm_fargate_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_FARGATE_PERCENTAGE                       MonthlyUsageAttributionSupportedMetrics = "apm_fargate_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APPSEC_FARGATE_USAGE                         MonthlyUsageAttributionSupportedMetrics = "appsec_fargate_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APPSEC_FARGATE_PERCENTAGE                    MonthlyUsageAttributionSupportedMetrics = "appsec_fargate_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_HOST_USAGE                               MonthlyUsageAttributionSupportedMetrics = "apm_host_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_HOST_PERCENTAGE                          MonthlyUsageAttributionSupportedMetrics = "apm_host_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_USM_USAGE                                MonthlyUsageAttributionSupportedMetrics = "apm_usm_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_USM_PERCENTAGE                           MonthlyUsageAttributionSupportedMetrics = "apm_usm_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APPSEC_USAGE                                 MonthlyUsageAttributionSupportedMetrics = "appsec_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APPSEC_PERCENTAGE                            MonthlyUsageAttributionSupportedMetrics = "appsec_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ASM_SERVERLESS_TRACED_INVOCATIONS_USAGE      MonthlyUsageAttributionSupportedMetrics = "asm_serverless_traced_invocations_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ASM_SERVERLESS_TRACED_INVOCATIONS_PERCENTAGE MonthlyUsageAttributionSupportedMetrics = "asm_serverless_traced_invocations_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_BROWSER_USAGE                                MonthlyUsageAttributionSupportedMetrics = "browser_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_BROWSER_PERCENTAGE                           MonthlyUsageAttributionSupportedMetrics = "browser_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_VISIBILITY_ITR_USAGE                      MonthlyUsageAttributionSupportedMetrics = "ci_visibility_itr_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_VISIBILITY_ITR_PERCENTAGE                 MonthlyUsageAttributionSupportedMetrics = "ci_visibility_itr_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CLOUD_SIEM_USAGE                             MonthlyUsageAttributionSupportedMetrics = "cloud_siem_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CLOUD_SIEM_PERCENTAGE                        MonthlyUsageAttributionSupportedMetrics = "cloud_siem_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CONTAINER_EXCL_AGENT_USAGE                   MonthlyUsageAttributionSupportedMetrics = "container_excl_agent_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CONTAINER_EXCL_AGENT_PERCENTAGE              MonthlyUsageAttributionSupportedMetrics = "container_excl_agent_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CONTAINER_USAGE                              MonthlyUsageAttributionSupportedMetrics = "container_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CONTAINER_PERCENTAGE                         MonthlyUsageAttributionSupportedMetrics = "container_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CSPM_CONTAINERS_PERCENTAGE                   MonthlyUsageAttributionSupportedMetrics = "cspm_containers_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CSPM_CONTAINERS_USAGE                        MonthlyUsageAttributionSupportedMetrics = "cspm_containers_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CSPM_HOSTS_PERCENTAGE                        MonthlyUsageAttributionSupportedMetrics = "cspm_hosts_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CSPM_HOSTS_USAGE                             MonthlyUsageAttributionSupportedMetrics = "cspm_hosts_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_TIMESERIES_USAGE                      MonthlyUsageAttributionSupportedMetrics = "custom_timeseries_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_TIMESERIES_PERCENTAGE                 MonthlyUsageAttributionSupportedMetrics = "custom_timeseries_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_INGESTED_TIMESERIES_USAGE             MonthlyUsageAttributionSupportedMetrics = "custom_ingested_timeseries_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_INGESTED_TIMESERIES_PERCENTAGE        MonthlyUsageAttributionSupportedMetrics = "custom_ingested_timeseries_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CWS_CONTAINERS_PERCENTAGE                    MonthlyUsageAttributionSupportedMetrics = "cws_containers_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CWS_CONTAINERS_USAGE                         MonthlyUsageAttributionSupportedMetrics = "cws_containers_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CWS_HOSTS_PERCENTAGE                         MonthlyUsageAttributionSupportedMetrics = "cws_hosts_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CWS_HOSTS_USAGE                              MonthlyUsageAttributionSupportedMetrics = "cws_hosts_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_DBM_HOSTS_PERCENTAGE                         MonthlyUsageAttributionSupportedMetrics = "dbm_hosts_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_DBM_HOSTS_USAGE                              MonthlyUsageAttributionSupportedMetrics = "dbm_hosts_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_DBM_QUERIES_PERCENTAGE                       MonthlyUsageAttributionSupportedMetrics = "dbm_queries_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_DBM_QUERIES_USAGE                            MonthlyUsageAttributionSupportedMetrics = "dbm_queries_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ERROR_TRACKING_USAGE                         MonthlyUsageAttributionSupportedMetrics = "error_tracking_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ERROR_TRACKING_PERCENTAGE                    MonthlyUsageAttributionSupportedMetrics = "error_tracking_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INDEXED_LOGS_USAGE                 MonthlyUsageAttributionSupportedMetrics = "estimated_indexed_logs_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INDEXED_LOGS_PERCENTAGE            MonthlyUsageAttributionSupportedMetrics = "estimated_indexed_logs_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INGESTED_LOGS_USAGE                MonthlyUsageAttributionSupportedMetrics = "estimated_ingested_logs_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INGESTED_LOGS_PERCENTAGE           MonthlyUsageAttributionSupportedMetrics = "estimated_ingested_logs_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INDEXED_SPANS_USAGE                MonthlyUsageAttributionSupportedMetrics = "estimated_indexed_spans_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INDEXED_SPANS_PERCENTAGE           MonthlyUsageAttributionSupportedMetrics = "estimated_indexed_spans_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INGESTED_SPANS_USAGE               MonthlyUsageAttributionSupportedMetrics = "estimated_ingested_spans_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INGESTED_SPANS_PERCENTAGE          MonthlyUsageAttributionSupportedMetrics = "estimated_ingested_spans_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_FARGATE_USAGE                                MonthlyUsageAttributionSupportedMetrics = "fargate_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_FARGATE_PERCENTAGE                           MonthlyUsageAttributionSupportedMetrics = "fargate_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_FUNCTIONS_USAGE                              MonthlyUsageAttributionSupportedMetrics = "functions_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_FUNCTIONS_PERCENTAGE                         MonthlyUsageAttributionSupportedMetrics = "functions_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INFRA_HOST_USAGE                             MonthlyUsageAttributionSupportedMetrics = "infra_host_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INFRA_HOST_PERCENTAGE                        MonthlyUsageAttributionSupportedMetrics = "infra_host_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INVOCATIONS_USAGE                            MonthlyUsageAttributionSupportedMetrics = "invocations_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INVOCATIONS_PERCENTAGE                       MonthlyUsageAttributionSupportedMetrics = "invocations_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LAMBDA_TRACED_INVOCATIONS_USAGE              MonthlyUsageAttributionSupportedMetrics = "lambda_traced_invocations_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LAMBDA_TRACED_INVOCATIONS_PERCENTAGE         MonthlyUsageAttributionSupportedMetrics = "lambda_traced_invocations_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_MOBILE_APP_TESTING_USAGE                     MonthlyUsageAttributionSupportedMetrics = "mobile_app_testing_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_MOBILE_APP_TESTING_PERCENTAGE                MonthlyUsageAttributionSupportedMetrics = "mobile_app_testing_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_NDM_NETFLOW_USAGE                            MonthlyUsageAttributionSupportedMetrics = "ndm_netflow_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_NDM_NETFLOW_PERCENTAGE                       MonthlyUsageAttributionSupportedMetrics = "ndm_netflow_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_NPM_HOST_USAGE                               MonthlyUsageAttributionSupportedMetrics = "npm_host_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_NPM_HOST_PERCENTAGE                          MonthlyUsageAttributionSupportedMetrics = "npm_host_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_OBS_PIPELINE_BYTES_USAGE                     MonthlyUsageAttributionSupportedMetrics = "obs_pipeline_bytes_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_OBS_PIPELINE_BYTES_PERCENTAGE                MonthlyUsageAttributionSupportedMetrics = "obs_pipeline_bytes_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_CONTAINER_USAGE                     MonthlyUsageAttributionSupportedMetrics = "profiled_container_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_CONTAINER_PERCENTAGE                MonthlyUsageAttributionSupportedMetrics = "profiled_container_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_FARGATE_USAGE                       MonthlyUsageAttributionSupportedMetrics = "profiled_fargate_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_FARGATE_PERCENTAGE                  MonthlyUsageAttributionSupportedMetrics = "profiled_fargate_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_HOST_USAGE                          MonthlyUsageAttributionSupportedMetrics = "profiled_host_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_HOST_PERCENTAGE                     MonthlyUsageAttributionSupportedMetrics = "profiled_host_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SERVERLESS_APPS_USAGE                        MonthlyUsageAttributionSupportedMetrics = "serverless_apps_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SERVERLESS_APPS_PERCENTAGE                   MonthlyUsageAttributionSupportedMetrics = "serverless_apps_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SNMP_USAGE                                   MonthlyUsageAttributionSupportedMetrics = "snmp_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SNMP_PERCENTAGE                              MonthlyUsageAttributionSupportedMetrics = "snmp_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_RUM_SESSIONS_USAGE                 MonthlyUsageAttributionSupportedMetrics = "estimated_rum_sessions_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_RUM_SESSIONS_PERCENTAGE            MonthlyUsageAttributionSupportedMetrics = "estimated_rum_sessions_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_UNIVERSAL_SERVICE_MONITORING_USAGE           MonthlyUsageAttributionSupportedMetrics = "universal_service_monitoring_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_UNIVERSAL_SERVICE_MONITORING_PERCENTAGE      MonthlyUsageAttributionSupportedMetrics = "universal_service_monitoring_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_VULN_MANAGEMENT_HOSTS_USAGE                  MonthlyUsageAttributionSupportedMetrics = "vuln_management_hosts_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_VULN_MANAGEMENT_HOSTS_PERCENTAGE             MonthlyUsageAttributionSupportedMetrics = "vuln_management_hosts_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SDS_SCANNED_BYTES_USAGE                      MonthlyUsageAttributionSupportedMetrics = "sds_scanned_bytes_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SDS_SCANNED_BYTES_PERCENTAGE                 MonthlyUsageAttributionSupportedMetrics = "sds_scanned_bytes_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_TEST_INDEXED_SPANS_USAGE                  MonthlyUsageAttributionSupportedMetrics = "ci_test_indexed_spans_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_TEST_INDEXED_SPANS_PERCENTAGE             MonthlyUsageAttributionSupportedMetrics = "ci_test_indexed_spans_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INGESTED_LOGS_BYTES_USAGE                    MonthlyUsageAttributionSupportedMetrics = "ingested_logs_bytes_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INGESTED_LOGS_BYTES_PERCENTAGE               MonthlyUsageAttributionSupportedMetrics = "ingested_logs_bytes_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_PIPELINE_INDEXED_SPANS_USAGE              MonthlyUsageAttributionSupportedMetrics = "ci_pipeline_indexed_spans_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_PIPELINE_INDEXED_SPANS_PERCENTAGE         MonthlyUsageAttributionSupportedMetrics = "ci_pipeline_indexed_spans_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INDEXED_SPANS_USAGE                          MonthlyUsageAttributionSupportedMetrics = "indexed_spans_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INDEXED_SPANS_PERCENTAGE                     MonthlyUsageAttributionSupportedMetrics = "indexed_spans_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_EVENT_USAGE                           MonthlyUsageAttributionSupportedMetrics = "custom_event_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_EVENT_PERCENTAGE                      MonthlyUsageAttributionSupportedMetrics = "custom_event_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_CUSTOM_RETENTION_USAGE          MonthlyUsageAttributionSupportedMetrics = "logs_indexed_custom_retention_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_CUSTOM_RETENTION_PERCENTAGE     MonthlyUsageAttributionSupportedMetrics = "logs_indexed_custom_retention_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_360DAY_USAGE                    MonthlyUsageAttributionSupportedMetrics = "logs_indexed_360day_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_360DAY_PERCENTAGE               MonthlyUsageAttributionSupportedMetrics = "logs_indexed_360day_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_180DAY_USAGE                    MonthlyUsageAttributionSupportedMetrics = "logs_indexed_180day_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_180DAY_PERCENTAGE               MonthlyUsageAttributionSupportedMetrics = "logs_indexed_180day_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_90DAY_USAGE                     MonthlyUsageAttributionSupportedMetrics = "logs_indexed_90day_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_90DAY_PERCENTAGE                MonthlyUsageAttributionSupportedMetrics = "logs_indexed_90day_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_60DAY_USAGE                     MonthlyUsageAttributionSupportedMetrics = "logs_indexed_60day_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_60DAY_PERCENTAGE                MonthlyUsageAttributionSupportedMetrics = "logs_indexed_60day_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_45DAY_USAGE                     MonthlyUsageAttributionSupportedMetrics = "logs_indexed_45day_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_45DAY_PERCENTAGE                MonthlyUsageAttributionSupportedMetrics = "logs_indexed_45day_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_30DAY_USAGE                     MonthlyUsageAttributionSupportedMetrics = "logs_indexed_30day_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_30DAY_PERCENTAGE                MonthlyUsageAttributionSupportedMetrics = "logs_indexed_30day_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_15DAY_USAGE                     MonthlyUsageAttributionSupportedMetrics = "logs_indexed_15day_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_15DAY_PERCENTAGE                MonthlyUsageAttributionSupportedMetrics = "logs_indexed_15day_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_7DAY_USAGE                      MonthlyUsageAttributionSupportedMetrics = "logs_indexed_7day_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_7DAY_PERCENTAGE                 MonthlyUsageAttributionSupportedMetrics = "logs_indexed_7day_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_3DAY_USAGE                      MonthlyUsageAttributionSupportedMetrics = "logs_indexed_3day_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_3DAY_PERCENTAGE                 MonthlyUsageAttributionSupportedMetrics = "logs_indexed_3day_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_RUM_REPLAY_SESSIONS_USAGE                    MonthlyUsageAttributionSupportedMetrics = "rum_replay_sessions_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_RUM_REPLAY_SESSIONS_PERCENTAGE               MonthlyUsageAttributionSupportedMetrics = "rum_replay_sessions_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_RUM_BROWSER_MOBILE_SESSIONS_USAGE            MonthlyUsageAttributionSupportedMetrics = "rum_browser_mobile_sessions_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_RUM_BROWSER_MOBILE_SESSIONS_PERCENTAGE       MonthlyUsageAttributionSupportedMetrics = "rum_browser_mobile_sessions_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INGESTED_SPANS_BYTES_USAGE                   MonthlyUsageAttributionSupportedMetrics = "ingested_spans_bytes_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INGESTED_SPANS_BYTES_PERCENTAGE              MonthlyUsageAttributionSupportedMetrics = "ingested_spans_bytes_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SIEM_INGESTED_BYTES_USAGE                    MonthlyUsageAttributionSupportedMetrics = "siem_ingested_bytes_usage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SIEM_INGESTED_BYTES_PERCENTAGE               MonthlyUsageAttributionSupportedMetrics = "siem_ingested_bytes_percentage"
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ALL                                          MonthlyUsageAttributionSupportedMetrics = "*"
)

var allowedMonthlyUsageAttributionSupportedMetricsEnumValues = []MonthlyUsageAttributionSupportedMetrics{
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_API_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_API_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_FARGATE_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_FARGATE_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APPSEC_FARGATE_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APPSEC_FARGATE_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_HOST_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_HOST_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_USM_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APM_USM_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APPSEC_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_APPSEC_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ASM_SERVERLESS_TRACED_INVOCATIONS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ASM_SERVERLESS_TRACED_INVOCATIONS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_BROWSER_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_BROWSER_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_VISIBILITY_ITR_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_VISIBILITY_ITR_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CLOUD_SIEM_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CLOUD_SIEM_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CONTAINER_EXCL_AGENT_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CONTAINER_EXCL_AGENT_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CONTAINER_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CONTAINER_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CSPM_CONTAINERS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CSPM_CONTAINERS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CSPM_HOSTS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CSPM_HOSTS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_TIMESERIES_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_TIMESERIES_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_INGESTED_TIMESERIES_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_INGESTED_TIMESERIES_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CWS_CONTAINERS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CWS_CONTAINERS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CWS_HOSTS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CWS_HOSTS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_DBM_HOSTS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_DBM_HOSTS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_DBM_QUERIES_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_DBM_QUERIES_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ERROR_TRACKING_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ERROR_TRACKING_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INDEXED_LOGS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INDEXED_LOGS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INGESTED_LOGS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INGESTED_LOGS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INDEXED_SPANS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INDEXED_SPANS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INGESTED_SPANS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_INGESTED_SPANS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_FARGATE_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_FARGATE_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_FUNCTIONS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_FUNCTIONS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INFRA_HOST_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INFRA_HOST_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INVOCATIONS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INVOCATIONS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LAMBDA_TRACED_INVOCATIONS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LAMBDA_TRACED_INVOCATIONS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_MOBILE_APP_TESTING_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_MOBILE_APP_TESTING_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_NDM_NETFLOW_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_NDM_NETFLOW_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_NPM_HOST_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_NPM_HOST_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_OBS_PIPELINE_BYTES_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_OBS_PIPELINE_BYTES_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_CONTAINER_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_CONTAINER_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_FARGATE_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_FARGATE_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_HOST_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_PROFILED_HOST_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SERVERLESS_APPS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SERVERLESS_APPS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SNMP_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SNMP_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_RUM_SESSIONS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ESTIMATED_RUM_SESSIONS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_UNIVERSAL_SERVICE_MONITORING_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_UNIVERSAL_SERVICE_MONITORING_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_VULN_MANAGEMENT_HOSTS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_VULN_MANAGEMENT_HOSTS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SDS_SCANNED_BYTES_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SDS_SCANNED_BYTES_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_TEST_INDEXED_SPANS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_TEST_INDEXED_SPANS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INGESTED_LOGS_BYTES_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INGESTED_LOGS_BYTES_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_PIPELINE_INDEXED_SPANS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CI_PIPELINE_INDEXED_SPANS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INDEXED_SPANS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INDEXED_SPANS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_EVENT_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_CUSTOM_EVENT_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_CUSTOM_RETENTION_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_CUSTOM_RETENTION_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_360DAY_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_360DAY_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_180DAY_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_180DAY_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_90DAY_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_90DAY_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_60DAY_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_60DAY_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_45DAY_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_45DAY_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_30DAY_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_30DAY_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_15DAY_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_15DAY_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_7DAY_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_7DAY_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_3DAY_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_LOGS_INDEXED_3DAY_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_RUM_REPLAY_SESSIONS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_RUM_REPLAY_SESSIONS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_RUM_BROWSER_MOBILE_SESSIONS_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_RUM_BROWSER_MOBILE_SESSIONS_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INGESTED_SPANS_BYTES_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_INGESTED_SPANS_BYTES_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SIEM_INGESTED_BYTES_USAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_SIEM_INGESTED_BYTES_PERCENTAGE,
	MONTHLYUSAGEATTRIBUTIONSUPPORTEDMETRICS_ALL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MonthlyUsageAttributionSupportedMetrics) GetAllowedValues() []MonthlyUsageAttributionSupportedMetrics {
	return allowedMonthlyUsageAttributionSupportedMetricsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MonthlyUsageAttributionSupportedMetrics) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MonthlyUsageAttributionSupportedMetrics(value)
	return nil
}

// NewMonthlyUsageAttributionSupportedMetricsFromValue returns a pointer to a valid MonthlyUsageAttributionSupportedMetrics
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMonthlyUsageAttributionSupportedMetricsFromValue(v string) (*MonthlyUsageAttributionSupportedMetrics, error) {
	ev := MonthlyUsageAttributionSupportedMetrics(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MonthlyUsageAttributionSupportedMetrics: valid values are %v", v, allowedMonthlyUsageAttributionSupportedMetricsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MonthlyUsageAttributionSupportedMetrics) IsValid() bool {
	for _, existing := range allowedMonthlyUsageAttributionSupportedMetricsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonthlyUsageAttributionSupportedMetrics value.
func (v MonthlyUsageAttributionSupportedMetrics) Ptr() *MonthlyUsageAttributionSupportedMetrics {
	return &v
}
