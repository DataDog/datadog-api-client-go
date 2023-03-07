@endpoint(incidents) @endpoint(incidents-v2)
Feature: Incidents
  Manage incident response.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "Incidents" API

  @team:DataDog/incident-app
  Scenario: Add commander to an incident returns "OK" response
    Given operation "UpdateIncident" enabled
    And there is a valid "user" in the system
    And there is a valid "incident" in the system
    And new "UpdateIncident" request
    And request contains "incident_id" parameter from "incident.data.id"
    And body with value {"data": {"id": "{{incident.data.id}}", "type": "incidents", "relationships": {"commander_user": {"data": {"id": "{{user.data.id}}", "type": "users"}}}}}
    When the request is sent
    Then the response status is 200 OK

  @team:DataDog/incident-app
  Scenario: Create an incident attachment returns "OK" response
    Given operation "UpdateIncidentAttachments" enabled
    And there is a valid "incident" in the system
    And new "UpdateIncidentAttachments" request
    And request contains "incident_id" parameter from "incident.data.id"
    And body with value {"data": [{"type": "incident_attachments", "attributes": {"attachment_type": "link", "attachment": {"documentUrl": "https://www.example.com/doc", "title": "{{unique}}"}}}]}
    When the request is sent
    Then the response status is 200 OK
    And the response "data" has length 1
    And the response "data[0].type" is equal to "incident_attachments"
    And the response "data[0].attributes.attachment_type" is equal to "link"
    And the response "data[0].attributes.attachment.documentUrl" is equal to "https://www.example.com/doc"

  @generated @skip @team:DataDog/incident-app
  Scenario: Create an incident integration metadata returns "Bad Request" response
    Given operation "CreateIncidentIntegration" enabled
    And new "CreateIncidentIntegration" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"incident_id": "00000000-aaaa-0000-0000-000000000000", "integration_type": 1, "metadata": {"channels": []}}, "type": "incident_integrations"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/incident-app
  Scenario: Create an incident integration metadata returns "CREATED" response
    Given operation "CreateIncidentIntegration" enabled
    And new "CreateIncidentIntegration" request
    And there is a valid "incident" in the system
    And request contains "incident_id" parameter from "incident.data.id"
    And body with value {"data": {"attributes": {"incident_id": "{{ incident.data.id }}", "integration_type": 1, "metadata": {"channels": [{"channel_id": "C0123456789", "channel_name": "#new-channel", "team_id": "T01234567", "redirect_url": "https://slack.com/app_redirect?channel=C0123456789&team=T01234567"}]}}, "type": "incident_integrations"}}
    When the request is sent
    Then the response status is 201 CREATED
    And the response "data.type" is equal to "incident_integrations"
    And the response "data.attributes.metadata.channels" has length 1
    And the response "data.attributes.metadata.channels[0].channel_name" is equal to "#new-channel"

  @generated @skip @team:DataDog/incident-app
  Scenario: Create an incident integration metadata returns "Not Found" response
    Given operation "CreateIncidentIntegration" enabled
    And new "CreateIncidentIntegration" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"incident_id": "00000000-aaaa-0000-0000-000000000000", "integration_type": 1, "metadata": {"channels": []}}, "type": "incident_integrations"}}
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/incident-app
  Scenario: Create an incident returns "Bad Request" response
    Given operation "CreateIncident" enabled
    And new "CreateIncident" request
    And body with value {"data": {"attributes": {"customer_impacted": false, "fields": {"severity": {"type": "dropdown", "value": "SEV-5"}}, "initial_cells": [{"cell_type": "markdown", "content": {"content": "An example timeline cell message."}, "important": false}], "notification_handles": [{"display_name": "Jane Doe", "handle": "@test.user@test.com"}], "title": "A test incident title"}, "relationships": {"commander_user": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}}, "type": "incidents"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/incident-app
  Scenario: Create an incident returns "CREATED" response
    Given there is a valid "user" in the system
    And operation "CreateIncident" enabled
    And new "CreateIncident" request
    And body with value {"data": {"type": "incidents", "attributes": {"title": "{{unique}}", "customer_impacted": false, "fields": {"state": {"type": "dropdown", "value": "resolved"}}}, "relationships": {"commander_user": {"data": {"type": "{{ user.data.type }}", "id": "{{ user.data.id }}"}}}}}
    When the request is sent
    Then the response status is 201 CREATED
    And the response "data.relationships.commander_user.data.id" has the same value as "user.data.id"
    And the response "data.attributes.title" has the same value as "unique"

  @generated @skip @team:DataDog/incident-app
  Scenario: Create an incident returns "Not Found" response
    Given operation "CreateIncident" enabled
    And new "CreateIncident" request
    And body with value {"data": {"attributes": {"customer_impacted": false, "fields": {"severity": {"type": "dropdown", "value": "SEV-5"}}, "initial_cells": [{"cell_type": "markdown", "content": {"content": "An example timeline cell message."}, "important": false}], "notification_handles": [{"display_name": "Jane Doe", "handle": "@test.user@test.com"}], "title": "A test incident title"}, "relationships": {"commander_user": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}}, "type": "incidents"}}
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/incident-app
  Scenario: Create, update, and delete incident attachments returns "Bad Request" response
    Given operation "UpdateIncidentAttachments" enabled
    And new "UpdateIncidentAttachments" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And body with value {"data": [{"attributes": {"attachment": {"documentUrl": "https://app.datadoghq.com/notebook/123", "title": "Postmortem IR-123"}, "attachment_type": "postmortem"}, "id": "00000000-abcd-0002-0000-000000000000", "type": "incident_attachments"}, {"attributes": {"attachment": {"documentUrl": "https://www.example.com/webstore-failure-runbook", "title": "Runbook for webstore service failures"}, "attachment_type": "link"}, "type": "incident_attachments"}, {"id": "00000000-abcd-0003-0000-000000000000", "type": "incident_attachments"}]}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/incident-app
  Scenario: Create, update, and delete incident attachments returns "Not Found" response
    Given operation "UpdateIncidentAttachments" enabled
    And new "UpdateIncidentAttachments" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And body with value {"data": [{"attributes": {"attachment": {"documentUrl": "https://app.datadoghq.com/notebook/123", "title": "Postmortem IR-123"}, "attachment_type": "postmortem"}, "id": "00000000-abcd-0002-0000-000000000000", "type": "incident_attachments"}, {"attributes": {"attachment": {"documentUrl": "https://www.example.com/webstore-failure-runbook", "title": "Runbook for webstore service failures"}, "attachment_type": "link"}, "type": "incident_attachments"}, {"id": "00000000-abcd-0003-0000-000000000000", "type": "incident_attachments"}]}
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/incident-app
  Scenario: Create, update, and delete incident attachments returns "OK" response
    Given operation "UpdateIncidentAttachments" enabled
    And new "UpdateIncidentAttachments" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And body with value {"data": [{"attributes": {"attachment": {"documentUrl": "https://app.datadoghq.com/notebook/123", "title": "Postmortem IR-123"}, "attachment_type": "postmortem"}, "id": "00000000-abcd-0002-0000-000000000000", "type": "incident_attachments"}, {"attributes": {"attachment": {"documentUrl": "https://www.example.com/webstore-failure-runbook", "title": "Runbook for webstore service failures"}, "attachment_type": "link"}, "type": "incident_attachments"}, {"id": "00000000-abcd-0003-0000-000000000000", "type": "incident_attachments"}]}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/incident-app
  Scenario: Delete an existing incident returns "Bad Request" response
    Given operation "DeleteIncident" enabled
    And new "DeleteIncident" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/incident-app
  Scenario: Delete an existing incident returns "Not Found" response
    Given operation "DeleteIncident" enabled
    And new "DeleteIncident" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/incident-app
  Scenario: Delete an existing incident returns "OK" response
    Given operation "DeleteIncident" enabled
    And there is a valid "incident" in the system
    And new "DeleteIncident" request
    And request contains "incident_id" parameter from "incident.data.id"
    When the request is sent
    Then the response status is 204 OK

  @generated @skip @team:DataDog/incident-app
  Scenario: Delete an incident integration metadata returns "Bad Request" response
    Given operation "DeleteIncidentIntegration" enabled
    And new "DeleteIncidentIntegration" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And request contains "integration_metadata_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/incident-app
  Scenario: Delete an incident integration metadata returns "Not Found" response
    Given operation "DeleteIncidentIntegration" enabled
    And new "DeleteIncidentIntegration" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And request contains "integration_metadata_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/incident-app
  Scenario: Delete an incident integration metadata returns "OK" response
    Given operation "DeleteIncidentIntegration" enabled
    And new "DeleteIncidentIntegration" request
    And there is a valid "incident" in the system
    And request contains "incident_id" parameter from "incident.data.id"
    And the "incident" has an "incident_integration_metadata"
    And request contains "integration_metadata_id" parameter from "incident_integration_metadata.data.id"
    When the request is sent
    Then the response status is 204 OK

  @generated @skip @team:DataDog/incident-app
  Scenario: Get a list of an incident's integration metadata returns "Bad Request" response
    Given operation "ListIncidentIntegrations" enabled
    And new "ListIncidentIntegrations" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/incident-app
  Scenario: Get a list of an incident's integration metadata returns "Not Found" response
    Given operation "ListIncidentIntegrations" enabled
    And new "ListIncidentIntegrations" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/incident-app
  Scenario: Get a list of an incident's integration metadata returns "OK" response
    Given operation "ListIncidentIntegrations" enabled
    And new "ListIncidentIntegrations" request
    And there is a valid "incident" in the system
    And request contains "incident_id" parameter from "incident.data.id"
    And the "incident" has an "incident_integration_metadata"
    When the request is sent
    Then the response status is 200 OK
    And the response "data[0].attributes.metadata.channels[0].channel_name" is equal to "#example-channel-name"

  @generated @skip @team:DataDog/incident-app
  Scenario: Get a list of attachments returns "Bad Request" response
    Given operation "ListIncidentAttachments" enabled
    And new "ListIncidentAttachments" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/incident-app
  Scenario: Get a list of attachments returns "Not Found" response
    Given operation "ListIncidentAttachments" enabled
    And new "ListIncidentAttachments" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/incident-app
  Scenario: Get a list of attachments returns "OK" response
    Given operation "ListIncidentAttachments" enabled
    And new "ListIncidentAttachments" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/incident-app
  Scenario: Get a list of incidents returns "Bad Request" response
    Given operation "ListIncidents" enabled
    And new "ListIncidents" request
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/incident-app
  Scenario: Get a list of incidents returns "Not Found" response
    Given operation "ListIncidents" enabled
    And new "ListIncidents" request
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/incident-app
  Scenario: Get a list of incidents returns "OK" response
    Given operation "ListIncidents" enabled
    And there is a valid "incident" in the system
    And new "ListIncidents" request
    When the request is sent
    Then the response status is 200 OK
    And the response "data[0].type" is equal to "incidents"

  @replay-only @team:DataDog/incident-app @with-pagination
  Scenario: Get a list of incidents returns "OK" response with pagination
    Given operation "ListIncidents" enabled
    And new "ListIncidents" request
    And request contains "page[size]" parameter with value 2
    When the request with pagination is sent
    Then the response status is 200 OK
    And the response has 3 items

  @team:DataDog/incident-app
  Scenario: Get incident attachments returns "OK" response
    Given operation "ListIncidentAttachments" enabled
    And there is a valid "incident" in the system
    And the "incident" has an "incident_attachment"
    And new "ListIncidentAttachments" request
    And request contains "incident_id" parameter from "incident.data.id"
    When the request is sent
    Then the response status is 200 OK
    And the response "data" has length 1
    And the response "data[0].type" is equal to "incident_attachments"
    And the response "data[0].attributes.attachment_type" is equal to "link"
    And the response "data[0].attributes.attachment.documentUrl" is equal to "https://www.example.com/doc"

  @generated @skip @team:DataDog/incident-app
  Scenario: Get incident integration metadata details returns "Bad Request" response
    Given operation "GetIncidentIntegration" enabled
    And new "GetIncidentIntegration" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And request contains "integration_metadata_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/incident-app
  Scenario: Get incident integration metadata details returns "Not Found" response
    Given operation "GetIncidentIntegration" enabled
    And new "GetIncidentIntegration" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And request contains "integration_metadata_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/incident-app
  Scenario: Get incident integration metadata details returns "OK" response
    Given operation "GetIncidentIntegration" enabled
    And new "GetIncidentIntegration" request
    And there is a valid "incident" in the system
    And request contains "incident_id" parameter from "incident.data.id"
    And the "incident" has an "incident_integration_metadata"
    And request contains "integration_metadata_id" parameter from "incident_integration_metadata.data.id"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/incident-app
  Scenario: Get the details of an incident returns "Bad Request" response
    Given operation "GetIncident" enabled
    And new "GetIncident" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/incident-app
  Scenario: Get the details of an incident returns "Not Found" response
    Given operation "GetIncident" enabled
    And new "GetIncident" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/incident-app
  Scenario: Get the details of an incident returns "OK" response
    Given operation "GetIncident" enabled
    And there is a valid "incident" in the system
    And new "GetIncident" request
    And request contains "incident_id" parameter from "incident.data.id"
    When the request is sent
    Then the response status is 200 OK
    And the response "data.attributes.title" has the same value as "incident.data.attributes.title"

  @team:DataDog/incident-app
  Scenario: Remove commander from an incident returns "OK" response
    Given operation "UpdateIncident" enabled
    And there is a valid "incident" in the system
    And new "UpdateIncident" request
    And request contains "incident_id" parameter from "incident.data.id"
    And body with value {"data": {"id": "{{incident.data.id}}", "type": "incidents", "relationships": {"commander_user": {"data": null}}}}
    When the request is sent
    Then the response status is 200 OK
    And the response "data.relationships.commander_user.data" is equal to null

  @generated @skip @team:DataDog/incident-app
  Scenario: Search for incidents returns "Bad Request" response
    Given operation "SearchIncidents" enabled
    And new "SearchIncidents" request
    And request contains "query" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/incident-app
  Scenario: Search for incidents returns "Not Found" response
    Given operation "SearchIncidents" enabled
    And new "SearchIncidents" request
    And request contains "query" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/incident-app
  Scenario: Search for incidents returns "OK" response
    Given operation "SearchIncidents" enabled
    And there is a valid "incident" in the system
    And new "SearchIncidents" request
    And request contains "query" parameter with value "state:(active OR stable OR resolved)"
    When the request is sent
    Then the response status is 200 OK
    And the response "data.type" is equal to "incidents_search_results"
    And the response "data.attributes.incidents[0].data.type" is equal to "incidents"

  @generated @skip @team:DataDog/incident-app
  Scenario: Update an existing incident integration metadata returns "Bad Request" response
    Given operation "UpdateIncidentIntegration" enabled
    And new "UpdateIncidentIntegration" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And request contains "integration_metadata_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"incident_id": "00000000-aaaa-0000-0000-000000000000", "integration_type": 1, "metadata": {"channels": []}}, "type": "incident_integrations"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/incident-app
  Scenario: Update an existing incident integration metadata returns "Not Found" response
    Given operation "UpdateIncidentIntegration" enabled
    And new "UpdateIncidentIntegration" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And request contains "integration_metadata_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"incident_id": "00000000-aaaa-0000-0000-000000000000", "integration_type": 1, "metadata": {"channels": []}}, "type": "incident_integrations"}}
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/incident-app
  Scenario: Update an existing incident integration metadata returns "OK" response
    Given operation "UpdateIncidentIntegration" enabled
    And new "UpdateIncidentIntegration" request
    And there is a valid "incident" in the system
    And request contains "incident_id" parameter from "incident.data.id"
    And the "incident" has an "incident_integration_metadata"
    And request contains "integration_metadata_id" parameter from "incident_integration_metadata.data.id"
    And body with value {"data": {"attributes": {"incident_id": "{{ incident.data.id }}", "integration_type": 1, "metadata": {"channels": [{"channel_id": "C0123456789", "channel_name": "#updated-channel-name", "team_id": "T01234567", "redirect_url": "https://slack.com/app_redirect?channel=C0123456789&team=T01234567"}]}}, "type": "incident_integrations"}}
    When the request is sent
    Then the response status is 200 OK
    And the response "data.attributes.metadata.channels[0].channel_name" is equal to "#updated-channel-name"

  @generated @skip @team:DataDog/incident-app
  Scenario: Update an existing incident returns "Bad Request" response
    Given operation "UpdateIncident" enabled
    And new "UpdateIncident" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"customer_impact_end": null, "customer_impact_scope": "Example customer impact scope", "customer_impact_start": null, "customer_impacted": false, "detected": null, "fields": {"severity": {"type": "dropdown", "value": "SEV-5"}}, "notification_handles": [{"display_name": "Jane Doe", "handle": "@test.user@test.com"}], "title": "A test incident title"}, "id": "00000000-0000-0000-4567-000000000000", "relationships": {"commander_user": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}, "integrations": {"data": [{"id": "00000000-abcd-0005-0000-000000000000", "type": "incident_integrations"}, {"id": "00000000-abcd-0006-0000-000000000000", "type": "incident_integrations"}]}, "postmortem": {"data": {"id": "00000000-0000-abcd-3000-000000000000", "type": "incident_postmortems"}}}, "type": "incidents"}}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/incident-app
  Scenario: Update an existing incident returns "Not Found" response
    Given operation "UpdateIncident" enabled
    And new "UpdateIncident" request
    And request contains "incident_id" parameter from "REPLACE.ME"
    And body with value {"data": {"attributes": {"customer_impact_end": null, "customer_impact_scope": "Example customer impact scope", "customer_impact_start": null, "customer_impacted": false, "detected": null, "fields": {"severity": {"type": "dropdown", "value": "SEV-5"}}, "notification_handles": [{"display_name": "Jane Doe", "handle": "@test.user@test.com"}], "title": "A test incident title"}, "id": "00000000-0000-0000-4567-000000000000", "relationships": {"commander_user": {"data": {"id": "00000000-0000-0000-0000-000000000000", "type": "users"}}, "integrations": {"data": [{"id": "00000000-abcd-0005-0000-000000000000", "type": "incident_integrations"}, {"id": "00000000-abcd-0006-0000-000000000000", "type": "incident_integrations"}]}, "postmortem": {"data": {"id": "00000000-0000-abcd-3000-000000000000", "type": "incident_postmortems"}}}, "type": "incidents"}}
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/incident-app
  Scenario: Update an existing incident returns "OK" response
    Given operation "UpdateIncident" enabled
    And there is a valid "incident" in the system
    And new "UpdateIncident" request
    And request contains "incident_id" parameter from "incident.data.id"
    And body with value {"data": {"id": "{{incident.data.id}}", "type": "incidents", "attributes": {"fields": {"state": {"type": "dropdown", "value":"resolved"}}, "title": "{{ incident.data.attributes.title }}-updated"}}}
    When the request is sent
    Then the response status is 200 OK
    And the response "data.attributes.title" is equal to "{{ incident.data.attributes.title }}-updated"
