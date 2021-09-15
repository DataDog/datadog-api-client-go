# CHANGELOG

## 1.4.0 / 2021-09-15

* [Added] Added `available_values` property to template variables schema. See [#1089](https://github.com/DataDog/datadog-api-client-go/pull/1089).
* [Added] Add `follow_redirects` options to test request in Synthetics. See [#1096](https://github.com/DataDog/datadog-api-client-go/pull/1096).
* [Added] ApmDependencyStatsQuery for formulas and functions dashboard widgets. See [#1103](https://github.com/DataDog/datadog-api-client-go/pull/1103).
* [Added] Add formula and function APM resource stats query definition for dashboards. See [#1104](https://github.com/DataDog/datadog-api-client-go/pull/1104).
* [Fixed] Fix SLO history error response type for overall errors. See [#1095](https://github.com/DataDog/datadog-api-client-go/pull/1095).
* [Fixed] Mark SLO Correction Type as required. See [#1093](https://github.com/DataDog/datadog-api-client-go/pull/1093).
* [Fixed] Make the `name` property required for APM Dependency Stat Query. See [#1110](https://github.com/DataDog/datadog-api-client-go/pull/1110).
* [Changed] Fix SLO history schema for groups and monitors fields. See [#1099](https://github.com/DataDog/datadog-api-client-go/pull/1099).
* [Changed] Remove metadata from required list for metric SLO history endpoint. See [#1102](https://github.com/DataDog/datadog-api-client-go/pull/1102).

## 1.3.0 / 2021-08-26

* [Added] Add config variables to Synthetics browser test config. See [#1086](https://github.com/DataDog/datadog-api-client-go/pull/1086).
* [Added] Add DBM usage endpoint. See [#1068](https://github.com/DataDog/datadog-api-client-go/pull/1068).
* [Added] Add `audit alert` monitor type. See [#1081](https://github.com/DataDog/datadog-api-client-go/pull/1081).
* [Added] Add `batch_id` to the synthetics trigger endpoint response. See [#1079](https://github.com/DataDog/datadog-api-client-go/pull/1079).
* [Added] Adding support for security monitoring rule `type` property. See [#1065](https://github.com/DataDog/datadog-api-client-go/pull/1065).
* [Added] Add events data source to Dashboard widgets. See [#1067](https://github.com/DataDog/datadog-api-client-go/pull/1067).
* [Added] Add restricted roles for Synthetics global variables. See [#1072](https://github.com/DataDog/datadog-api-client-go/pull/1072).
* [Added] Webhooks integration SDK. See [#1071](https://github.com/DataDog/datadog-api-client-go/pull/1071).
* [Added] Add missing synthetics variable parser type `x_path`. See [#1070](https://github.com/DataDog/datadog-api-client-go/pull/1070).
* [Added] Add `audit_stream` to `ListStreamSource`. See [#1056](https://github.com/DataDog/datadog-api-client-go/pull/1056).
* [Added] Add percentile to dashboard `WidgetAggregator` schema. See [#1051](https://github.com/DataDog/datadog-api-client-go/pull/1051).
* [Added] Add `id_str` property to Event response. See [#1059](https://github.com/DataDog/datadog-api-client-go/pull/1059).
* [Added] Add edge to Synthetics devices. See [#1063](https://github.com/DataDog/datadog-api-client-go/pull/1063).
* [Added] Add endpoints to manage Service Accounts v2. See [#1043](https://github.com/DataDog/datadog-api-client-go/pull/1043).
* [Added] Add `new_group_delay` and deprecate `new_host_delay` monitor properties. See [#1055](https://github.com/DataDog/datadog-api-client-go/pull/1055).
* [Added] Add `include_descendants` param to usage attribution API. See [#1062](https://github.com/DataDog/datadog-api-client-go/pull/1062).
* [Added] Improve resiliency of go SDK when deserializing enums/oneOfs. See [#1028](https://github.com/DataDog/datadog-api-client-go/pull/1028).
* [Added] Add `ContainsUnparsedObject` utility method to check if an object wasn't fully deserialized. See [#1073](https://github.com/DataDog/datadog-api-client-go/pull/1073) and [#1077](https://github.com/DataDog/datadog-api-client-go/pull/1077).
* [Added] Add support for list widget in dashboards. See [#1023](https://github.com/DataDog/datadog-api-client-go/pull/1023).
* [Added] Extend table widget requests to support formulas and functions in dashboards. See [#1046](https://github.com/DataDog/datadog-api-client-go/pull/1046).
* [Added] Add CSPM to usage attribution. See [#1037](https://github.com/DataDog/datadog-api-client-go/pull/1037).
* [Added] Add support for dashboard bulk delete and restore endpoints. See [#1020](https://github.com/DataDog/datadog-api-client-go/pull/1020).
* [Added] Add support for audit logs data source in dashboards. See [#1041](https://github.com/DataDog/datadog-api-client-go/pull/1041).
* [Added] Add `allow_insecure` option for multistep steps in Synthetics. See [#1031](https://github.com/DataDog/datadog-api-client-go/pull/1031).
* [Fixed] Make SLO history metadata unit nullable. See [#1078](https://github.com/DataDog/datadog-api-client-go/pull/1078).
* [Fixed] Minor fixes of the incident schema. See [#1074](https://github.com/DataDog/datadog-api-client-go/pull/1074).
* [Fixed] Fix serialization of query metrics response containing nullable points. See [#1034](https://github.com/DataDog/datadog-api-client-go/pull/1034).
* [Fixed] Fix `status` property name for browser error status in Synthetics. See [#1036](https://github.com/DataDog/datadog-api-client-go/pull/1036).
* [Changed] Add separate schema for deleting AWS account. See [#1030](https://github.com/DataDog/datadog-api-client-go/pull/1030).
* [Removed] Remove deprecated endpoints `/api/v1/usage/traces` and `/api/v1/usage/tracing-without-limits`. See [#1038](https://github.com/DataDog/datadog-api-client-go/pull/1038).

## 1.2.0 / 2021-07-09

* [Added] Add support for `GET /api/v2/application_keys/{app_key_id}`. See [#1021](https://github.com/DataDog/datadog-api-client-go/pull/1021).
* [Added] Add `meta` property with pagination info to SLOCorrectionList endpoint response. See [#1018](https://github.com/DataDog/datadog-api-client-go/pull/1018).
* [Added] Add support for treemap widget. See [#1013](https://github.com/DataDog/datadog-api-client-go/pull/1013).
* [Added] Add missing properties `query_index` and `tag_set` to `MetricsQueryMetadata`. See [#979](https://github.com/DataDog/datadog-api-client-go/pull/979).
* [Fixed] Remove US only constraint for AWS tag filtering. See [#1007](https://github.com/DataDog/datadog-api-client-go/pull/1007).
* [Fixed] Add BDD tests to synthetics. See [#1006](https://github.com/DataDog/datadog-api-client-go/pull/1006).
* [Fixed] Fix response of security filter delete. See [#1002](https://github.com/DataDog/datadog-api-client-go/pull/1002).
* [Fixed] Handle null in query metrics unit. See [#1001](https://github.com/DataDog/datadog-api-client-go/pull/1001).
* [Changed] Remove Synthetics tick interval enum. See [#1005](https://github.com/DataDog/datadog-api-client-go/pull/1005).

## 1.1.0 / 2021-06-16

* [Added] Add missing fields `hasExtendedTitle`, `type`, `version` and `updateAuthorId` for Security Monitoring Rule endpoints. See [#998](https://github.com/DataDog/datadog-api-client-go/pull/998).
* [Added] Dashboard RBAC role support. See [#993](https://github.com/DataDog/datadog-api-client-go/pull/993).
* [Fixed] Fix go JSON struct. See [#992](https://github.com/DataDog/datadog-api-client-go/pull/992).

## 1.0.0 / 2021-06-10

* [Added] Add missing fields in usage billable summary keys. See [#987](https://github.com/DataDog/datadog-api-client-go/pull/987).
* [Added] Add monitor name and priority options. See [#984](https://github.com/DataDog/datadog-api-client-go/pull/984).
* [Added] Add endpoint to list Synthetics global variables. See [#965](https://github.com/DataDog/datadog-api-client-go/pull/965).
* [Added] Add monitors search endpoints. See [#959](https://github.com/DataDog/datadog-api-client-go/pull/959).
* [Added] Add CWS to usage metering endpoint. See [#964](https://github.com/DataDog/datadog-api-client-go/pull/964).
* [Added] Add `tag_config_source` to usage attribution response. See [#952](https://github.com/DataDog/datadog-api-client-go/pull/952).
* [Added] Add audit logs to usage endpoints. See [#978](https://github.com/DataDog/datadog-api-client-go/pull/978).
* [Fixed] Make `assertions` field optional for multistep synthetics tests, and add `global` config variable type. See [#961](https://github.com/DataDog/datadog-api-client-go/pull/961).
* [Fixed] Fix type of day/month response attribute in custom metrics usage. See [#981](https://github.com/DataDog/datadog-api-client-go/pull/981).
* [Fixed] Properly mark monitor required fields. See [#950](https://github.com/DataDog/datadog-api-client-go/pull/950).
* [Changed] Rename `compliance` to `CSPM` in usage endpoint. See [#978](https://github.com/DataDog/datadog-api-client-go/pull/978).
* [Changed] Rename `incident_integration_metadata` to `incident_integrations` to match API. See [#944](https://github.com/DataDog/datadog-api-client-go/pull/944).

## 1.0.0-beta.22 / 2021-05-17

* [Added] Add endpoints to configure Security Filters. See [#938](https://github.com/DataDog/datadog-api-client-go/pull/938).
* [Added] Add `active_child` nested downtime object to `Downtime` component for downtime APIs. See [#930](https://github.com/DataDog/datadog-api-client-go/pull/930).
* [Changed] Change Dashboard WidgetCustomLink properties. See [#937](https://github.com/DataDog/datadog-api-client-go/pull/937).
* [Changed] Make various fixes to synthetics models. See [#935](https://github.com/DataDog/datadog-api-client-go/pull/935).
* [Changed] Update usage attribute endpoint metadata fields. See [#932](https://github.com/DataDog/datadog-api-client-go/pull/932).

## 1.0.0-beta.21 / 2021-05-12

* [Added] Notebooks Public API Documentation. See [#926](https://github.com/DataDog/datadog-api-client-go/pull/926).
* [Added] Add `logs_by_retention` usage property and `GetUsageLogsByRetention` endpoint. See [#915](https://github.com/DataDog/datadog-api-client-go/pull/915).
* [Added] Add anomaly detection method to `SecurityMonitoringRuleDetectionMethod` enum. See [#914](https://github.com/DataDog/datadog-api-client-go/pull/914).
* [Added] Add `with_configured_alert_ids` parameter to get a SLO details endpoint. See [#910](https://github.com/DataDog/datadog-api-client-go/pull/910).
* [Added] Add `setCookie`, `dnsServerPort`,  `allowFailure ` and `isCritical` fields for Synthetics tests. See [#903](https://github.com/DataDog/datadog-api-client-go/pull/903).
* [Added] Add `metadata` property with pagination info to `SLOList` endpoint response. See [#899](https://github.com/DataDog/datadog-api-client-go/pull/899).
* [Added] Add new properties to group widget, note widget and image widget. See [#895](https://github.com/DataDog/datadog-api-client-go/pull/895).
* [Added] Add support for a `rate` metric type in manage metric tags v2 endpoint. See [#892](https://github.com/DataDog/datadog-api-client-go/pull/892).
* [Fixed] Handle typed nils for go client. See [#927](https://github.com/DataDog/datadog-api-client-go/pull/927).
* [Fixed] Remove default value of `is_column_break` layout property of dashboard. See [#925](https://github.com/DataDog/datadog-api-client-go/pull/925).
* [Changed] Enumerate accepted values for fields parameter in usage attr requests. See [#919](https://github.com/DataDog/datadog-api-client-go/pull/919).
* [Changed] Add frequency and remove request as required field from synthetics test. See [#916](https://github.com/DataDog/datadog-api-client-go/pull/916).

## 1.0.0-beta.20 / 2021-04-27

* [Added] Add support for ICMP Synthetics tests. See [#887](https://github.com/DataDog/datadog-api-client-go/pull/887).
* [Added] Add vSphere usage information. See [#880](https://github.com/DataDog/datadog-api-client-go/pull/880).
* [Added] Update properties for dashboard distribution widget. See [#877](https://github.com/DataDog/datadog-api-client-go/pull/877).
* [Added] Mark metric volumes and ingested tags endpoints as stable. See [#872](https://github.com/DataDog/datadog-api-client-go/pull/872).
* [Added] Add `filter[shared]` query parameter for searching dashboards. See [#860](https://github.com/DataDog/datadog-api-client-go/pull/860).
* [Added] Add profiling product fields in usage metering endpoint. See [#859](https://github.com/DataDog/datadog-api-client-go/pull/859).
* [Added] Add `title` and `background_color` properties to dashboard group widget. See [#858](https://github.com/DataDog/datadog-api-client-go/pull/858).
* [Changed] Use new model for Go client API. See [#885](https://github.com/DataDog/datadog-api-client-go/pull/885).
* [Removed] Remove deprecated Synthetics methods `CreateTest` and `UpdateTest`. See [#881](https://github.com/DataDog/datadog-api-client-go/pull/881).

## 1.0.0-beta.19 / 2021-04-14

* [Added] Add `reflow_type` property to dashboard object. See [#841](https://github.com/DataDog/datadog-api-client-go/pull/841).
* [Added] Add security track and formulas and functions support for geomap dashboard widget. See [#837](https://github.com/DataDog/datadog-api-client-go/pull/837).
* [Added] Generate intake endpoints. See [#834](https://github.com/DataDog/datadog-api-client-go/pull/834).
* [Added] Add endpoint for listing all downtimes for the specified monitor. See [#828](https://github.com/DataDog/datadog-api-client-go/pull/828).
* [Added] Add `modified_at` attribute to user response v2 schema. See [#817](https://github.com/DataDog/datadog-api-client-go/pull/817).
* [Added] Add default environment loading in clients. See [#816](https://github.com/DataDog/datadog-api-client-go/pull/816).
* [Added] Add `passed`, `noSavingResponseBody`, `noScreenshot`, and `disableCors` fields to Synthetics. See [#815](https://github.com/DataDog/datadog-api-client-go/pull/815).
* [Added] Add compliance usage endpoint and compliance host statistics. See [#814](https://github.com/DataDog/datadog-api-client-go/pull/814).
* [Added] Add tag filter options for `/api/v{1,2}/metrics`. See [#813](https://github.com/DataDog/datadog-api-client-go/pull/813).
* [Added] Add usage fields for Heroku and OpenTelemetry. See [#810](https://github.com/DataDog/datadog-api-client-go/pull/810).
* [Added] Add `global_time_target` field to SLO widget. See [#808](https://github.com/DataDog/datadog-api-client-go/pull/808).
* [Added] Add method to export an API test in Synthetics. See [#807](https://github.com/DataDog/datadog-api-client-go/pull/807).
* [Added] Add metadata to usage top average metrics response. See [#806](https://github.com/DataDog/datadog-api-client-go/pull/806).
* [Added] Add median as valid aggregator for formulas and functions. See [#800](https://github.com/DataDog/datadog-api-client-go/pull/800).
* [Fixed] Browser Test message required. See [#803](https://github.com/DataDog/datadog-api-client-go/pull/803).
* [Changed] Return correct object in `GetBrowserTest` endpoint. See [#827](https://github.com/DataDog/datadog-api-client-go/pull/827).
* [Changed] Add agent rules in security monitoring rules queries. See [#809](https://github.com/DataDog/datadog-api-client-go/pull/809).

## 1.0.0-beta.18 / 2021-03-22

* [Added] Add `legend_layout` and `legend_columns` to timeseries widget definition. See [#791](https://github.com/DataDog/datadog-api-client-go/pull/791).

## 1.0.0-beta.17 / 2021-03-15

* [Added] Add support for multistep tests in Synthetics. See [#775](https://github.com/DataDog/datadog-api-client-go/pull/775).
* [Added] Add core web vitals to synthetics browser test results. See [#771](https://github.com/DataDog/datadog-api-client-go/pull/771).
* [Added] Add v2 metric tags and metric volumes endpoints. See [#769](https://github.com/DataDog/datadog-api-client-go/pull/769).
* [Added] Add new endpoints for browser and API tests in Synthetics. See [#762](https://github.com/DataDog/datadog-api-client-go/pull/762).
* [Changed] Update response schema for service level objective operation `GetSLOHistory`. See [#784](https://github.com/DataDog/datadog-api-client-go/pull/784).
* [Changed] Make query name required in formulas and functions queries. See [#774](https://github.com/DataDog/datadog-api-client-go/pull/774).

## 1.0.0-beta.16 / 2021-03-02

* [Added] Add groupby_simple_monitor option to monitors. See [#758](https://github.com/DataDog/datadog-api-client-go/pull/758).
* [Added] Allow formula and functions in query value requests. See [#756](https://github.com/DataDog/datadog-api-client-go/pull/756).
* [Added] Allow formula and functions in toplist requests. See [#753](https://github.com/DataDog/datadog-api-client-go/pull/753).
* [Added] Add slack resource. See [#744](https://github.com/DataDog/datadog-api-client-go/pull/744).
* [Added] Add detectionMethod and newValueOptions fields to security monitoring rules. See [#739](https://github.com/DataDog/datadog-api-client-go/pull/739).
* [Added] Expose "event-v2 alert" monitor type. See [#738](https://github.com/DataDog/datadog-api-client-go/pull/738).
* [Added] Add new US3 region. See [#737](https://github.com/DataDog/datadog-api-client-go/pull/737).
* [Added] Add org_name field to usage attribution response. See [#736](https://github.com/DataDog/datadog-api-client-go/pull/736).
* [Added] Add profile_metrics_query properties to dashboard widget requests. See [#728](https://github.com/DataDog/datadog-api-client-go/pull/728).
* [Added] Add geomap widget to dashboards. See [#720](https://github.com/DataDog/datadog-api-client-go/pull/720).
* [Added] Add v2 API for metric tag configuration. See [#718](https://github.com/DataDog/datadog-api-client-go/pull/718).
* [Added] Add Lambda invocations usage to response. See [#716](https://github.com/DataDog/datadog-api-client-go/pull/716).
* [Added] Remove unstable flag for logs apis. See [#709](https://github.com/DataDog/datadog-api-client-go/pull/709).
* [Fixed] Add missing tlsVersion and minTlsVersion to Synthetics assertion types. See [#731](https://github.com/DataDog/datadog-api-client-go/pull/731).
* [Fixed] Change analyzed_spans to spans in dashboard. See [#711](https://github.com/DataDog/datadog-api-client-go/pull/711).
* [Changed] Rename objects throughout the code for consistency. See [#724](https://github.com/DataDog/datadog-api-client-go/pull/724).
* [Changed] Rename objects for formula and functions to be more generic. See [#747](https://github.com/DataDog/datadog-api-client-go/pull/747).

## v1.0.0-beta.15 / 2021-02-08

* [Added] Add restricted roles to monitor update. See [#691](https://github.com/DataDog/datadog-api-client-go/pull/691).
* [Added] Add endpoint for IoT billing usage. See [#684](https://github.com/DataDog/datadog-api-client-go/pull/684).
* [Added] Add query parameters for SLO search endpoint. See [#682](https://github.com/DataDog/datadog-api-client-go/pull/682).
* [Added] Add fields for formula and function query definition and widget formulas. See [#680](https://github.com/DataDog/datadog-api-client-go/pull/680).
* [Added] Add global_time to time_window SLO widget. See [#675](https://github.com/DataDog/datadog-api-client-go/pull/675).
* [Added] Update required fields in SLO correction create and update requests. See [#668](https://github.com/DataDog/datadog-api-client-go/pull/668).
* [Fixed] Fix AWS tag filter delete request. See [#701](https://github.com/DataDog/datadog-api-client-go/pull/701).
* [Fixed] Remove unnecessary field from TimeSeriesFormulaAndFunctionEventQuery. See [#700](https://github.com/DataDog/datadog-api-client-go/pull/700).
* [Fixed] Fix unit format in SLO history response. See [#695](https://github.com/DataDog/datadog-api-client-go/pull/695).
* [Fixed] Change group_by from object to list of objects. See [#694](https://github.com/DataDog/datadog-api-client-go/pull/694).
* [Fixed] Fix location of monitor restricted roles. See [#687](https://github.com/DataDog/datadog-api-client-go/pull/687).
* [Fixed] Fix paging parameter names for logs aggregate queries. See [#681](https://github.com/DataDog/datadog-api-client-go/pull/681).

## v1.0.0-beta.14 / 2021-01-19

* [Added] Add log index creation. See [#662](https://github.com/DataDog/datadog-api-client-go/pull/662).
* [Added] Add SLO Corrections. See [#654](https://github.com/DataDog/datadog-api-client-go/pull/654).
* [Added] Add new live and rehydrated logs breakdowns for Usage API. See [#652](https://github.com/DataDog/datadog-api-client-go/pull/652).
* [Added] Add support for Synthetics variables from test. See [#641](https://github.com/DataDog/datadog-api-client-go/pull/641).
* [Fixed] Add additionalProperties: false to synthetics target field. See [#657](https://github.com/DataDog/datadog-api-client-go/pull/657).
* [Fixed] Fix missing field for synthetics variables from test. See [#649](https://github.com/DataDog/datadog-api-client-go/pull/649).
* [Changed] Extract key sorting enum to a specific schema in key management endpoint. See [#646](https://github.com/DataDog/datadog-api-client-go/pull/646).
* [Changed] Extract enum to specific schema in incidents endpoint. See [#650](https://github.com/DataDog/datadog-api-client-go/pull/650).
* [Changed] Fix some integer/number formats in Logs and Synthetics endpoints. See [#658](https://github.com/DataDog/datadog-api-client-go/pull/658).

## 1.0.0-beta.13 / 2021-01-06

* [Added] Added filters to rule endpoints in security monitoring API. See [#632](https://github.com/DataDog/datadog-api-client-go/pull/632).
* [Added] Add Azure app services fields to usage v1 endpoints. See [#631](https://github.com/DataDog/datadog-api-client-go/pull/631).
* [Added] Add mobile RUM OS types usage fields. See [#629](https://github.com/DataDog/datadog-api-client-go/pull/629).
* [Added] Add config variables for synthetics API tests. See [#628](https://github.com/DataDog/datadog-api-client-go/pull/628).
* [Added] Add endpoints for the public API of Logs2Metrics. See [#626](https://github.com/DataDog/datadog-api-client-go/pull/626).
* [Added] Add endpoints for API Keys v2. See [#620](https://github.com/DataDog/datadog-api-client-go/pull/620).
* [Added] Add utils to validate and create valid enums. See [#617](https://github.com/DataDog/datadog-api-client-go/pull/617).
* [Added] Add javascript value to synthetics browser variable types. See [#616](https://github.com/DataDog/datadog-api-client-go/pull/616).
* [Added] Add synthetics assertion operator. See [#609](https://github.com/DataDog/datadog-api-client-go/pull/609).
* [Added] Application keys v2 API. See [#605](https://github.com/DataDog/datadog-api-client-go/pull/605).
* [Fixed] Redact auth methods from debug logs. See [#618](https://github.com/DataDog/datadog-api-client-go/pull/618).
* [Removed] Remove Synthetic resources property. See [#622](https://github.com/DataDog/datadog-api-client-go/pull/622).

## 1.0.0-beta.12 / 2020-12-07

* [Added] Mark Usage Attribution endpoint as public beta. See [#592](https://github.com/DataDog/datadog-api-client-go/pull/592).
* [Added] Add AWS filtering endpoints. See [#589](https://github.com/DataDog/datadog-api-client-go/pull/589).
* [Added] Add limit parameter for get usage top average metrics. See [#586](https://github.com/DataDog/datadog-api-client-go/pull/586).
* [Added] Add endpoint to fetch process summaries. See [#585](https://github.com/DataDog/datadog-api-client-go/pull/585).
* [Added] Add synthetics private location endpoints. See [#584](https://github.com/DataDog/datadog-api-client-go/pull/584).
* [Added] Add user_update, recommendation and snapshot as event alert types. See [#583](https://github.com/DataDog/datadog-api-client-go/pull/583).
* [Added] Add Usage Attribution endpoint. See [#582](https://github.com/DataDog/datadog-api-client-go/pull/582).
* [Added] Add new API for incident management usage. See [#578](https://github.com/DataDog/datadog-api-client-go/pull/578).
* [Added] Add the incident schema. See [#572](https://github.com/DataDog/datadog-api-client-go/pull/572).
* [Added] Add IP prefixes by location for synthetics endpoints. See [#565](https://github.com/DataDog/datadog-api-client-go/pull/565).
* [Added] Add filter parameter for listing teams and services. See [#564](https://github.com/DataDog/datadog-api-client-go/pull/564).
* [Added] Add restricted roles to monitor create and edit requests. See [#562](https://github.com/DataDog/datadog-api-client-go/pull/562).
* [Fixed] Quota & retention are now editable fields in log indexes. See [#568](https://github.com/DataDog/datadog-api-client-go/pull/568).
* [Changed] Mark request bodies as required or explicitly optional. See [#598](https://github.com/DataDog/datadog-api-client-go/pull/598).
* [Changed] Deprecate subscription and billing fields in create organization endpoint. See [#588](https://github.com/DataDog/datadog-api-client-go/pull/588).
* [Changed] Mark query field as optional when searching logs. See [#577](https://github.com/DataDog/datadog-api-client-go/pull/577).
* [Changed] Change event_query property to use log query definition in dashboard widgets. See [#573](https://github.com/DataDog/datadog-api-client-go/pull/573).
* [Changed] Rename tracing without limits and traces usage endpoints. See [#561](https://github.com/DataDog/datadog-api-client-go/pull/561).
* [Removed] Remove org_id parameter from Usage Attribution endpoint. See [#594](https://github.com/DataDog/datadog-api-client-go/pull/594).

## v1.0.0-beta.11 / 2020-11-06

* [Added] Add 3 new palettes to the conditional formatting options. See [#554](https://github.com/DataDog/datadog-api-client-go/pull/554).

## v1.0.0-beta.10 / 2020-11-02

* [Changed] Change teams and services objects names to be incident specific. See [#538](https://github.com/DataDog/datadog-api-client-go/pull/538).
* [Removed] Remove `require_full_window` client default value for monitors. See [#540](https://github.com/DataDog/datadog-api-client-go/pull/540).

## 1.0.0-beta.9 / 2020-10-27

* [Added] Add missing synthetics step types. See [#534](https://github.com/DataDog/datadog-api-client-go/pull/534).
* [Added] Add include_tags in logs archives. See [#530](https://github.com/DataDog/datadog-api-client-go/pull/530).
* [Added] Add dns server and client certificate support to synthetics tests. See [#523](https://github.com/DataDog/datadog-api-client-go/pull/523).
* [Added] Add rehydration_tags property to the logs archives. See [#513](https://github.com/DataDog/datadog-api-client-go/pull/513).
* [Added] Add endpoint to reorder Logs Archives. See [#505](https://github.com/DataDog/datadog-api-client-go/pull/505).
* [Added] Add has_search_bar and cell_display_mode properties to table widget definition. See [#502](https://github.com/DataDog/datadog-api-client-go/pull/502).
* [Added] Add target_format property to the Logs attribute remapper . See [#501](https://github.com/DataDog/datadog-api-client-go/pull/501).
* [Added] Add dual y-axis configuration to time-series widget in Dashboard. See [#498](https://github.com/DataDog/datadog-api-client-go/pull/498).
* [Added] Mark logs aggregate endpoint as stable. See [#496](https://github.com/DataDog/datadog-api-client-go/pull/496).
* [Added] Add endpoint to get a Synthetics global variable. See [#489](https://github.com/DataDog/datadog-api-client-go/pull/489).
* [Added] Add assertion types for DNS Synthetics tests. See [#486](https://github.com/DataDog/datadog-api-client-go/pull/486).
* [Added] Add DNS test type to Synthetics. See [#482](https://github.com/DataDog/datadog-api-client-go/pull/482).
* [Added] Add API endpoints for teams and services. See [#470](https://github.com/DataDog/datadog-api-client-go/pull/470).
* [Added] Add mobile_rum_session_count_sum property to usage responses. See [#469](https://github.com/DataDog/datadog-api-client-go/pull/469).
* [Fixed] Fix synthetics_check_id type in MonitorOptions. See [#526](https://github.com/DataDog/datadog-api-client-go/pull/526).
* [Fixed] Remove default for cell_display_mode in table widget. See [#519](https://github.com/DataDog/datadog-api-client-go/pull/519).
* [Fixed] Fix tags attribute type in event aggregation API. See [#463](https://github.com/DataDog/datadog-api-client-go/pull/463).
* [Changed] Change `columns` attribute type from string array to object array in APM stats query widget. See [#509](https://github.com/DataDog/datadog-api-client-go/pull/509).
* [Changed] Rename to ApmStats and add required properties. See [#490](https://github.com/DataDog/datadog-api-client-go/pull/490).
* [Changed] Remove unused `aggregation_key` and `related_event_id` properties from events responses. See [#480](https://github.com/DataDog/datadog-api-client-go/pull/480).
* [Changed] Define required fields for v2 requests. See [#475](https://github.com/DataDog/datadog-api-client-go/pull/475).
* [Changed] Mark required type fields in User and Roles API v2. See [#467](https://github.com/DataDog/datadog-api-client-go/pull/467).
* [Removed] Remove check_type parameter from ListTests endpoint. See [#465](https://github.com/DataDog/datadog-api-client-go/pull/465).

## v1.0.0-beta.8 / 2020-09-16

* [Added] Add `aggregation` and `metric` fields to `SecurityMonitoringRuleQuery`. See [#457](https://github.com/DataDog/datadog-api-client-go/pull/457).
* [Added] Add tracing without limits to usage API. See [#449](https://github.com/DataDog/datadog-api-client-go/pull/449).
* [Added] Add response codes for AWS API. See [#443](https://github.com/DataDog/datadog-api-client-go/pull/443).
* [Added] Add `custom_links` support for Dashboard widgets. See [#442](https://github.com/DataDog/datadog-api-client-go/pull/442).
* [Added] Add profiling to usage API. See [#436](https://github.com/DataDog/datadog-api-client-go/pull/436).
* [Added] Add synthetics CI endpoint. See [#429](https://github.com/DataDog/datadog-api-client-go/pull/429).
* [Added] Add APM resources data source to table widgets. See [#428](https://github.com/DataDog/datadog-api-client-go/pull/428).
* [Added] Add list API for security monitoring signals. See [#424](https://github.com/DataDog/datadog-api-client-go/pull/424).
* [Added] Add create, edit and delete endpoints for synthetics global variables. See [#421](https://github.com/DataDog/datadog-api-client-go/pull/421).
* [Added] Add monitor option `renotify_interval` to synthetics tests. See [#420](https://github.com/DataDog/datadog-api-client-go/pull/420).
* [Added] Add event aggregation v2 API. See [#419](https://github.com/DataDog/datadog-api-client-go/pull/419).
* [Added] Add Profiling Host to Usage endpoint. See [#417](https://github.com/DataDog/datadog-api-client-go/pull/417).
* [Added] Add `distinctFields` to `SecurityMonitoringRuleQuery`. See [#412](https://github.com/DataDog/datadog-api-client-go/pull/412).
* [Added] Add missing `security_query` on `QueryValueWidgetRequest`. See [#407](https://github.com/DataDog/datadog-api-client-go/pull/407).
* [Added] Enable security source for dashboards. See [#403](https://github.com/DataDog/datadog-api-client-go/pull/403).
* [Added] Add SLO alerts to monitor enum. See [#401](https://github.com/DataDog/datadog-api-client-go/pull/401).
* [Fixed] Add 200 response code to PATCH v2 users. See [#441](https://github.com/DataDog/datadog-api-client-go/pull/441).
* [Fixed] Fix hourly host usage descriptions. See [#438](https://github.com/DataDog/datadog-api-client-go/pull/438).
* [Fixed] Remove enum from `legend_size` widget attribute. See [#432](https://github.com/DataDog/datadog-api-client-go/pull/432).
* [Fixed] Fix content-type spelling errors. See [#423](https://github.com/DataDog/datadog-api-client-go/pull/423).
* [Fixed] Properly mark `status` and `query` field as required for creation of Security Monitoring rule. See [#422](https://github.com/DataDog/datadog-api-client-go/pull/422).
* [Fixed] Fix name of `isEnabled` parameter for Security Monitoring rule. See [#409](https://github.com/DataDog/datadog-api-client-go/pull/409).
* [Removed] Remove 204 response from PATCH v2 users. See [#446](https://github.com/DataDog/datadog-api-client-go/pull/446).

## v1.0.0-beta.7 / 2020-07-22

* [Added] Adding four usage attribution endpoints. See [#393](https://github.com/DataDog/datadog-api-client-go/pull/393).
* [Added] Fix documentation for `v1/hosts`. See [#383](https://github.com/DataDog/datadog-api-client-go/pull/383).
* [Changed] Update synthetics test to contain latest features. See [#375](https://github.com/DataDog/datadog-api-client-go/pull/375).
* [Added] Usage Billable Summary response. See [#368](https://github.com/DataDog/datadog-api-client-go/pull/368).
* [Added] Add Logs Search API v2. See [#365](https://github.com/DataDog/datadog-api-client-go/pull/365).
* [Fixed] RRULE property for Downtimes API. See [#364](https://github.com/DataDog/datadog-api-client-go/pull/364).
* [Deprecated] Dashboards List v1 has been deprecated. See [#363](https://github.com/DataDog/datadog-api-client-go/pull/363).

## v1.0.0-beta.6 / 2020-06-19

* [Fixed] Update enum of synthetics devices IDs to match API. See [#351](https://github.com/DataDog/datadog-api-client-go/pull/351).

## v1.0.0-beta.5 / 2020-06-19

* [Added] Update to the latest openapi-generator 5 snapshot. See [#338](https://github.com/DataDog/datadog-api-client-go/pull/338).
* [Added] Add synthetics location endpoint. See [#334](https://github.com/DataDog/datadog-api-client-go/pull/334).
* [Fixed] Widget legend size can also be "0". See [#336](https://github.com/DataDog/datadog-api-client-go/pull/336).
* [Fixed] Log Index as an optional parameter (default to "*") for List Queries. See [#335](https://github.com/DataDog/datadog-api-client-go/pull/335).
* [Changed] Rename payload objects to request for `users` v2 API. See [#346](https://github.com/DataDog/datadog-api-client-go/pull/346).
  * This change includes backwards incompatible changes when using the `users` v2 endpoint.
* [Changed] Split schema for roles API. See [#337](https://github.com/DataDog/datadog-api-client-go/pull/337).
  * This change includes backwards incompatible changes when using the `role` endpoint.

## v1.0.0-beta.4 / 2020-06-09

* [BREAKING] Add missing values to enums. See [#320](https://github.com/DataDog/datadog-api-client-go/pull/320).
    * This change includes backwards incompatible changes when using the `MonitorSummary` widget.
* [BREAKING] Split schemas from DashboardList v2. See [#318](https://github.com/DataDog/datadog-api-client-go/pull/318).
    * This change includes backwards incompatible changes when using corresponding endpoints methods.
* [BREAKING] Clean synthetics test CRUD endpoints. See [#317](https://github.com/DataDog/datadog-api-client-go/pull/317).
    * This change includes backwards incompatible changes when using corresponding endpoints methods.
* [Added] Add Logs Archives endpoints. See [#323](https://github.com/DataDog/datadog-api-client-go/pull/323).

## v1.0.0-beta.3 / 2020-05-21

* [BREAKING] Update to openapi-generator 5.0.0. See [#303](https://github.com/DataDog/datadog-api-client-go/pull/303).
    * This change includes backwards incompatible changes when using structs generated from `oneOf` schemas.
* [Added] Add SIEM and SNMP usage API. See [#309](https://github.com/DataDog/datadog-api-client-go/pull/309).
* [Added] Add security monitoring to clients. See [#304](https://github.com/DataDog/datadog-api-client-go/pull/304).
* [Added] Add /v1/validate endpoint. See [#290](https://github.com/DataDog/datadog-api-client-go/pull/290).
* [Added] Add generated_files file. See [#270](https://github.com/DataDog/datadog-api-client-go/pull/270).
* [Fixed] Add authentication to Go examples. See [#299](https://github.com/DataDog/datadog-api-client-go/pull/299).
* [Fixed] Add 422 error codes to users and roles v2 endpoints. See [#296](https://github.com/DataDog/datadog-api-client-go/pull/296).
* [Fixed] Update import in Go examples. See [#295](https://github.com/DataDog/datadog-api-client-go/pull/295).
* [Fixed] Check duplicate object definitions. See [#288](https://github.com/DataDog/datadog-api-client-go/pull/288).
* [Fixed] Mark unstable endpoints with beta note. See [#281](https://github.com/DataDog/datadog-api-client-go/pull/281).
* [Changed] Update ServiceLevelObjective schema names. See [#279](https://github.com/DataDog/datadog-api-client-go/pull/279).
* [Deprecated] Add deprecated fields `logset`, `count` and `start` to appropriate dashboard widgets. See [#285](https://github.com/DataDog/datadog-api-client-go/pull/285).

## v1.0.0-beta.2 / 2020-05-04

* [Added] Add RUM Monitor Type and update documentation. See [#273](https://github.com/DataDog/datadog-api-client-go/pull/273).
* [Added] Add Logs Pipeline Processor. See [#268](https://github.com/DataDog/datadog-api-client-go/pull/268).
* [Added] Add additional fields to synthetics test request. See [#262](https://github.com/DataDog/datadog-api-client-go/pull/262).
* [Added] Add Monitor Pagination. See [#253](https://github.com/DataDog/datadog-api-client-go/pull/253).
* [Fixed] Mark synthetics test request "method" and "url" as optional. See [#265](https://github.com/DataDog/datadog-api-client-go/pull/265).
* [Fixed] Update error responses for roles v2 endpoints. See [#248](https://github.com/DataDog/datadog-api-client-go/pull/248).
* [Fixed] Add missing ListSLO's 404 response. See [#245](https://github.com/DataDog/datadog-api-client-go/pull/245).
* [Removed] Remove Pagerduty endpoints from the client. See [#264](https://github.com/DataDog/datadog-api-client-go/pull/264).

## 1.0.0-beta.1 / 2020-04-22

* [Added] Initial beta release of the Datadog API Client
