@endpoint(monitors) @endpoint(monitors-v1)
Feature: Monitors
  [Monitors](https://docs.datadoghq.com/monitors) allow you to watch a
  metric or check that you care about and notifies your team when a defined
  threshold has exceeded.  For more information, see [Creating
  Monitors](https://docs.datadoghq.com/monitors/create/types/).

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "Monitors" API

  @generated @skip @team:DataDog/monitor-app
  Scenario: Check if a monitor can be deleted returns "Bad Request" response
    Given new "CheckCanDeleteMonitor" request
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/monitor-app
  Scenario: Check if a monitor can be deleted returns "Deletion conflict error" response
    Given new "CheckCanDeleteMonitor" request
    When the request is sent
    Then the response status is 409 Deletion conflict error

  @team:DataDog/monitor-app
  Scenario: Check if a monitor can be deleted returns "OK" response
    Given there is a valid "monitor" in the system
    And new "CheckCanDeleteMonitor" request
    And request contains "monitor_ids" parameter with value [{{monitor.id}}]
    When the request is sent
    Then the response status is 200 OK

  @team:DataDog/monitor-app
  Scenario: Create a RUM formula and functions monitor returns "OK" response
    Given new "CreateMonitor" request
    And body with value {"name": "{{ unique }}","type": "rum alert","query": "formula(\"query2 / query1 * 100\").last(\"15m\") >= 0.8","message": "some message Notify: @hipchat-channel", "tags": ["test:{{ unique_lower_alnum }}", "env:ci"],"priority": 3,"options":{"thresholds":{"critical":0.8},"variables":[{"data_source": "rum","name": "query2","search": {"query": ""},"indexes": ["*"],"compute": {"aggregation": "count"},"group_by": []}, {"data_source": "rum","name": "query1","search": {"query": "status:error"},"indexes": ["*"],"compute": {"aggregation": "count"},"group_by": []}]}}
    When the request is sent
    Then the response status is 200 OK

  @team:DataDog/monitor-app
  Scenario: Create a ci-pipelines monitor returns "OK" response
    Given new "CreateMonitor" request
    And body with value {"name": "{{ unique }}","type": "ci-pipelines alert","query": "ci-pipelines(\"ci_level:pipeline @git.branch:staging* @ci.status:error\").rollup(\"count\").by(\"@git.branch,@ci.pipeline.name\").last(\"5m\") >= 1","message": "some message Notify: @hipchat-channel",	"tags": ["test:{{ unique_lower_alnum }}", "env:ci"],"priority": 3,"options":{"thresholds":{"critical":1}}}
    When the request is sent
    Then the response status is 200 OK

  @team:DataDog/monitor-app
  Scenario: Create a monitor returns "Bad Request" response
    Given new "CreateMonitor" request
    And body with value {"type": "log alert", "query": "query"}
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/monitor-app
  Scenario: Create a monitor returns "OK" response
    Given there is a valid "role" in the system
    And new "CreateMonitor" request
    And body with value {"name": "{{ unique }}", "type": "log alert", "query": "logs(\"service:foo AND type:error\").index(\"main\").rollup(\"count\").by(\"source\").last(\"5m\") > 2", "message": "some message Notify: @hipchat-channel", "tags": ["test:{{ unique_lower_alnum }}", "env:ci"], "priority": 3, "restricted_roles": ["{{ role.data.id }}"]}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/monitor-app
  Scenario: Delete a monitor returns "Bad Request" response
    Given new "DeleteMonitor" request
    And request contains "monitor_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/monitor-app
  Scenario: Delete a monitor returns "Item not found error" response
    Given new "DeleteMonitor" request
    And request contains "monitor_id" parameter with value 0
    When the request is sent
    Then the response status is 404 Item not found error

  @team:DataDog/monitor-app
  Scenario: Delete a monitor returns "OK" response
    Given there is a valid "monitor" in the system
    And new "DeleteMonitor" request
    And request contains "monitor_id" parameter from "monitor.id"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/monitor-app
  Scenario: Edit a monitor returns "Bad Request" response
    Given new "UpdateMonitor" request
    And request contains "monitor_id" parameter from "REPLACE.ME"
    And body with value {"message": null, "name": null, "options": {"enable_logs_sample": null, "escalation_message": "none", "evaluation_delay": null, "groupby_simple_monitor": null, "include_tags": true, "locked": null, "min_failure_duration": 0, "min_location_failed": 1, "new_group_delay": null, "new_host_delay": 300, "no_data_timeframe": null, "notify_audit": false, "notify_no_data": false, "renotify_interval": null, "renotify_occurrences": null, "renotify_statuses": ["alert"], "require_full_window": null, "silenced": null, "synthetics_check_id": null, "threshold_windows": {"recovery_window": null, "trigger_window": null}, "thresholds": {"critical": null, "critical_recovery": null, "ok": null, "unknown": null, "warning": null, "warning_recovery": null}, "timeout_h": null, "variables": [{"compute": {"aggregation": "avg", "interval": 60000, "metric": "@duration"}, "data_source": "rum", "group_by": [{"facet": "status", "limit": 10, "sort": {"aggregation": "avg", "metric": null, "order": "desc"}}], "indexes": ["days-3", "days-7"], "name": "query_errors", "search": {"query": "service:query"}}]}, "priority": null, "query": null, "restricted_roles": [null], "tags": [null], "type": "query alert"}
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/monitor-app
  Scenario: Edit a monitor returns "Monitor Not Found error" response
    Given new "UpdateMonitor" request
    And request contains "monitor_id" parameter with value 0
    And body with value {"name": "updated", "options": {"evaluation_delay": null, "new_group_delay": 600, "new_host_delay":null, "renotify_interval":null, "thresholds": {"critical":2, "warning": null}, "timeout_h": null}}
    When the request is sent
    Then the response status is 404 Monitor Not Found error

  @team:DataDog/monitor-app
  Scenario: Edit a monitor returns "OK" response
    Given there is a valid "monitor" in the system
    And new "UpdateMonitor" request
    And request contains "monitor_id" parameter from "monitor.id"
    And body with value {"name": "{{ monitor.name }}-updated", "options": {"evaluation_delay": null, "new_group_delay": 600, "new_host_delay":null, "renotify_interval":null, "thresholds": {"critical":2, "warning": null}, "timeout_h": null}}
    When the request is sent
    Then the response status is 200 OK
    And the response "name" is equal to "{{ monitor.name }}-updated"

  @generated @skip @team:DataDog/monitor-app
  Scenario: Get a monitor's details returns "Bad Request" response
    Given new "GetMonitor" request
    And request contains "monitor_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/monitor-app
  Scenario: Get a monitor's details returns "Monitor Not Found error" response
    Given new "GetMonitor" request
    And request contains "monitor_id" parameter with value 12345
    When the request is sent
    Then the response status is 404 Monitor Not Found error

  @team:DataDog/monitor-app
  Scenario: Get a monitor's details returns "OK" response
    Given there is a valid "monitor" in the system
    And new "GetMonitor" request
    And request contains "monitor_id" parameter from "monitor.id"
    When the request is sent
    Then the response status is 200 OK
    And the response "id" has the same value as "monitor.id"

  @team:DataDog/monitor-app
  Scenario: Get a synthetics monitor's details
    Given there is a valid "synthetics_api_test" in the system
    And new "GetMonitor" request
    And request contains "monitor_id" parameter from "synthetics_api_test.monitor_id"
    When the request is sent
    Then the response status is 200 OK
    And the response "options.synthetics_check_id" has the same value as "synthetics_api_test.public_id"

  @team:DataDog/monitor-app
  Scenario: Get all monitor details returns "Bad Request" response
    Given new "ListMonitors" request
    And request contains "group_states" parameter with value "notagroupstate"
    When the request is sent
    Then the response status is 400 Bad Request

  @integration-only @team:DataDog/monitor-app
  Scenario: Get all monitor details returns "OK" response
    Given new "ListMonitors" request
    When the request is sent
    Then the response status is 200 OK

  @skip @team:DataDog/monitor-app
  Scenario: Get all monitor details with tags
    Given there is a valid "monitor" in the system
    And new "ListMonitors" request
    And request contains "tags" parameter with value "test:{{ unique_lower_alnum }}"
    And request contains "page_size" parameter with value 1
    When the request is sent
    Then the response status is 200 OK
    And the response "[0].id" has the same value as "monitor.id"

  @team:DataDog/monitor-app
  Scenario: Monitors group search returns "Bad Request" response
    Given new "SearchMonitorGroups" request
    And request contains "query" parameter with value "status:notastatus"
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/monitor-app
  Scenario: Monitors group search returns "OK" response
    Given new "SearchMonitorGroups" request
    When the request is sent
    Then the response status is 200 OK

  @team:DataDog/monitor-app
  Scenario: Monitors search returns "Bad Request" response
    Given new "SearchMonitors" request
    And request contains "query" parameter with value "status:notastatus"
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/monitor-app
  Scenario: Monitors search returns "OK" response
    Given new "SearchMonitors" request
    When the request is sent
    Then the response status is 200 OK

  @team:DataDog/monitor-app
  Scenario: Validate a monitor returns "Invalid JSON" response
    Given new "ValidateMonitor" request
    And body with value {"type": "log alert", "query": "query"}
    When the request is sent
    Then the response status is 400 Invalid JSON

  @team:DataDog/monitor-app
  Scenario: Validate a monitor returns "OK" response
    Given new "ValidateMonitor" request
    And body from file "monitor_payload.json"
    When the request is sent
    Then the response status is 200 OK
