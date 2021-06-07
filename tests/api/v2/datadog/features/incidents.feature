@endpoint(incidents) @endpoint(incidents-v2)
Feature: Incidents
  Manage incident response.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "Incidents" API

  @generated @skip
  Scenario: Create an incident returns "Bad Request" response
    Given operation "CreateIncident" enabled
    And new "CreateIncident" request
    And body with value {"data": {"attributes": {"customer_impacted": false, "fields": {"severity": {"type": "dropdown", "value": "SEV-5"}}, "initial_timeline_cells": [{"cell_type": "markdown", "content": {"content": "An example timeline cell message."}, "important": false}], "notification_handles": ["@test.user@test.com"], "title": "A test incident title"}, "relationships": {"commander": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}}, "type": "incidents"}}
    When the request is sent
    Then the response status is 400 Bad Request

  Scenario: Create an incident returns "CREATED" response
    Given there is a valid "user" in the system
    And operation "CreateIncident" enabled
    And new "CreateIncident" request
    And body with value {"data": {"type": "incidents", "attributes": {"title": "{{unique}}", "customer_impacted": false, "fields": {"state": {"type": "dropdown", "value": "resolved"}}}, "relationships": {"commander": {"data": {"type": "{{ user.data.type }}", "id": "{{ user.data.id }}"}}}}}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Create an incident returns "Not Found" response
    Given operation "CreateIncident" enabled
    And new "CreateIncident" request
    And body with value {"data": {"attributes": {"customer_impacted": false, "fields": {"severity": {"type": "dropdown", "value": "SEV-5"}}, "initial_timeline_cells": [{"cell_type": "markdown", "content": {"content": "An example timeline cell message."}, "important": false}], "notification_handles": ["@test.user@test.com"], "title": "A test incident title"}, "relationships": {"commander": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}}, "type": "incidents"}}
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip
  Scenario: Delete an existing incident returns "Bad Request" response
    Given operation "DeleteIncident" enabled
    And new "DeleteIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Delete an existing incident returns "Not Found" response
    Given operation "DeleteIncident" enabled
    And new "DeleteIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 404 Not Found

  Scenario: Delete an existing incident returns "OK" response
    Given operation "DeleteIncident" enabled
    And there is a valid "incident" in the system
    And new "DeleteIncident" request
    And request contains "incident_id" parameter from "incident.data.id"
    When the request is sent
    Then the response status is 204 OK

  @generated @skip
  Scenario: Get a list of incidents returns "Bad Request" response
    Given operation "ListIncidents" enabled
    And new "ListIncidents" request
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Get a list of incidents returns "Not Found" response
    Given operation "ListIncidents" enabled
    And new "ListIncidents" request
    When the request is sent
    Then the response status is 404 Not Found

  Scenario: Get a list of incidents returns "OK" response
    Given operation "ListIncidents" enabled
    And there is a valid "incident" in the system
    And new "ListIncidents" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Get the details of an incident returns "Bad Request" response
    Given operation "GetIncident" enabled
    And new "GetIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Get the details of an incident returns "Not Found" response
    Given operation "GetIncident" enabled
    And new "GetIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 404 Not Found

  Scenario: Get the details of an incident returns "OK" response
    Given operation "GetIncident" enabled
    And there is a valid "incident" in the system
    And new "GetIncident" request
    And request contains "incident_id" parameter from "incident.data.id"
    When the request is sent
    Then the response status is 200 OK
    And the response "data.attributes.title" has the same value as "incident.data.attributes.title"

  @generated @skip
  Scenario: Update an existing incident returns "Bad Request" response
    Given operation "UpdateIncident" enabled
    And new "UpdateIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    And body with value {"data": {"attributes": {"customer_impact_end": null, "customer_impact_scope": "Example customer impact scope", "customer_impact_start": null, "customer_impacted": false, "detected": null, "fields": {"severity": {"type": "dropdown", "value": "SEV-5"}}, "notification_handles": ["@test.user@test.com"], "resolved": null, "title": "A test incident title"}, "id": "00000000-0000-0000-0000-000000000000", "relationships": {"commander_user": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}, "created_by_user": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}, "integrations": {"data": [{"id": "00000000-0000-0000-0000-000000000000", "type": "incident_integrations"}, {"id": "00000000-0000-0000-0000-000000000000", "type": "incident_integrations"}]}, "last_modified_by_user": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}, "postmortem": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "incident_postmortems"}}}, "type": "incidents"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip
  Scenario: Update an existing incident returns "Not Found" response
    Given operation "UpdateIncident" enabled
    And new "UpdateIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    And body with value {"data": {"attributes": {"customer_impact_end": null, "customer_impact_scope": "Example customer impact scope", "customer_impact_start": null, "customer_impacted": false, "detected": null, "fields": {"severity": {"type": "dropdown", "value": "SEV-5"}}, "notification_handles": ["@test.user@test.com"], "resolved": null, "title": "A test incident title"}, "id": "00000000-0000-0000-0000-000000000000", "relationships": {"commander_user": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}, "created_by_user": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}, "integrations": {"data": [{"id": "00000000-0000-0000-0000-000000000000", "type": "incident_integrations"}, {"id": "00000000-0000-0000-0000-000000000000", "type": "incident_integrations"}]}, "last_modified_by_user": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}, "postmortem": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "incident_postmortems"}}}, "type": "incidents"}}
    When the request is sent
    Then the response status is 404 Not Found

  Scenario: Update an existing incident returns "OK" response
    Given operation "UpdateIncident" enabled
    And there is a valid "incident" in the system
    And new "UpdateIncident" request
    And request contains "incident_id" parameter from "incident.data.id"
    And body with value {"data": {"id": "{{incident.data.id}}", "type": "incidents", "attributes": {"fields": {"state": {"type": "dropdown", "value":"resolved"}}, "title": "{{ incident.data.attributes.title }}-updated"}}}
    When the request is sent
    Then the response status is 200 OK
    And the response "data.attributes.title" is equal to "{{ incident.data.attributes.title }}-updated"
