@endpoint(dashboards) @endpoint(dashboards-v1)
Feature: Dashboards
  Interact with your dashboard lists through the API to make it easier to
  organize, find, and share all of your dashboards with your team and
  organization.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "Dashboards" API

  @generated @skip
  Scenario: Create a new dashboard returns "Bad Request" response
    Given new "CreateDashboard" request
    And body with value {"description": null, "is_read_only": false, "layout_type": "ordered", "notify_list": [null], "reflow_type": "auto", "restricted_roles": [null], "template_variable_presets": [{"name": null, "template_variables": [{"name": null, "value": null}]}], "template_variables": [{"default": "my-host", "name": "host1", "prefix": "host"}], "title": "", "widgets": [{"definition": {"requests": {"fill": {"q": "system.cpu.user"}}, "type": "hostmap"}}]}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Create a new dashboard returns "OK" response
    Given new "CreateDashboard" request
    And body with value {"description": null, "is_read_only": false, "layout_type": "ordered", "notify_list": [null], "reflow_type": "auto", "restricted_roles": [null], "template_variable_presets": [{"name": null, "template_variables": [{"name": null, "value": null}]}], "template_variables": [{"default": "my-host", "name": "host1", "prefix": "host"}], "title": "", "widgets": [{"definition": {"requests": {"fill": {"q": "system.cpu.user"}}, "type": "hostmap"}}]}
    When the request is sent
    Then the response status is 200 OK

  Scenario: Create a new dashboard with a profile metric query
    Given new "CreateDashboard" request
    And body from file "dashboard_payload.json"
    When the request is sent
    Then the response status is 200 OK

  Scenario: Create a new dashboard with an audit logs query
    Given new "CreateDashboard" request
    And body with value {"layout_type": "ordered", "title": "{{ unique }} with Audit Logs Query", "widgets": [{"definition": {"type": "timeseries","requests": [{"response_format": "timeseries","queries": [{"search": {"query": ""},"data_source": "audit","compute": {"aggregation": "count"},"name": "query1","indexes": ["*"],"group_by": []}]}]},"layout": {"x": 2,"y": 0,"width": 4,"height": 2}}]}
    When the request is sent
    Then the response status is 200 OK

  Scenario: Create a new dashboard with timeseries widget containing style attributes
    Given new "CreateDashboard" request
    And body with value {"layout_type": "ordered", "title": "{{ unique }} with timeseries widget","widgets": [{"definition": {"type": "timeseries","requests": [{"q": "sum:trace.test.errors{env:prod,service:datadog-api-spec} by {resource_name}.as_count()","on_right_yaxis": false,"style": {"palette": "warm","line_type": "solid","line_width": "normal"},"display_type": "bars"}]}}]}
    When the request is sent
    Then the response status is 200 OK
    And the response "widgets[0].definition.requests[0].on_right_yaxis" is false
    And the response "widgets[0].definition.requests[0].style" is equal to {"palette": "warm","line_type": "solid","line_width": "normal"}
    And the response "widgets[0].definition.requests[0].display_type" is equal to "bars"
    And the response "widgets[0].definition.requests[0].q" is equal to "sum:trace.test.errors{env:prod,service:datadog-api-spec} by {resource_name}.as_count()"

  @generated @skip
  Scenario: Delete a dashboard returns "Dashboards Not Found" response
    Given new "DeleteDashboard" request
    And request contains "dashboard_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 404 Dashboards Not Found

  @generated @skip
  Scenario: Delete a dashboard returns "OK" response
    Given new "DeleteDashboard" request
    And request contains "dashboard_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Get a dashboard returns "Item Not Found" response
    Given new "GetDashboard" request
    And request contains "dashboard_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 404 Item Not Found

  @generated @skip
  Scenario: Get a dashboard returns "OK" response
    Given new "GetDashboard" request
    And request contains "dashboard_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  Scenario: Get all dashboards returns "OK" response
    Given new "ListDashboards" request
    And there is a valid "dashboard" in the system
    And request contains "filter[shared]" parameter with value false
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update a dashboard returns "Bad Request" response
    Given new "UpdateDashboard" request
    And request contains "dashboard_id" parameter from "<PATH>"
    And body with value {"description": null, "is_read_only": false, "layout_type": "ordered", "notify_list": [null], "reflow_type": "auto", "restricted_roles": [null], "template_variable_presets": [{"name": null, "template_variables": [{"name": null, "value": null}]}], "template_variables": [{"default": "my-host", "name": "host1", "prefix": "host"}], "title": "", "widgets": [{"definition": {"requests": {"fill": {"q": "system.cpu.user"}}, "type": "hostmap"}}]}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Update a dashboard returns "Item Not Found" response
    Given new "UpdateDashboard" request
    And request contains "dashboard_id" parameter from "<PATH>"
    And body with value {"description": null, "is_read_only": false, "layout_type": "ordered", "notify_list": [null], "reflow_type": "auto", "restricted_roles": [null], "template_variable_presets": [{"name": null, "template_variables": [{"name": null, "value": null}]}], "template_variables": [{"default": "my-host", "name": "host1", "prefix": "host"}], "title": "", "widgets": [{"definition": {"requests": {"fill": {"q": "system.cpu.user"}}, "type": "hostmap"}}]}
    When the request is sent
    Then the response status is 404 Item Not Found

  @generated @skip
  Scenario: Update a dashboard returns "OK" response
    Given new "UpdateDashboard" request
    And request contains "dashboard_id" parameter from "<PATH>"
    And body with value {"description": null, "is_read_only": false, "layout_type": "ordered", "notify_list": [null], "reflow_type": "auto", "restricted_roles": [null], "template_variable_presets": [{"name": null, "template_variables": [{"name": null, "value": null}]}], "template_variables": [{"default": "my-host", "name": "host1", "prefix": "host"}], "title": "", "widgets": [{"definition": {"requests": {"fill": {"q": "system.cpu.user"}}, "type": "hostmap"}}]}
    When the request is sent
    Then the response status is 200 OK
