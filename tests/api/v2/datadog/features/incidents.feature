@endpoint(incidents)
Feature: Incidents
  Interact with incidents api endpoints.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "Incidents" API

  @generated @skip
  Scenario: Get a list of incidents returns "OK" response
    Given new "GetIncidents" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Create a new incident returns "CREATED" response
    Given new "CreateIncident" request
    And body {}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Get a list of all configuration fields returns "OK" response
    Given new "GetUserDefinedFields" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Deletes an existing field returns "DELETED" response
    Given new "DeleteUserDefinedField" request
    And request contains "field_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Get the details of a field returns "OK" response
    Given new "GetUserDefinedField" request
    And request contains "field_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Get all choices for the given field returns "OK" response
    Given new "GetUserDefinedFieldChoices" request
    And request contains "field_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Create a new choice returns "CREATED" response
    Given new "CreateUserDefinedFieldChoice" request
    And request contains "field_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Delete an existing incident returns "DELETED" response
    Given new "DeleteIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Get the details of an incident returns "OK" response
    Given new "GetIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update an existing incident returns "OK" response
    Given new "PatchIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Get a list of all Postmortems returns "OK" response
    Given new "GetIncidentPostmortems" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Create a new postmortem returns "CREATED" response
    Given new "CreateIncidentPostmortem" request
    And body {}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Update an existing postmortem returns "DELETED" response
    Given new "DeleteIncidentPostmortem" request
    And request contains "postmortem_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Get details of a postmortem returns "OK" response
    Given new "GetIncidentPostmortem" request
    And request contains "postmortem_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update an existing postmortem returns "OK" response
    Given new "PatchIncidentPostmortem" request
    And request contains "postmortem_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Get all selections for the given incident returns "OK" response
    Given new "GetUserDefinedFieldSelections" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Create a new selection returns "CREATED" response
    Given new "CreateUserDefinedFieldSelection" request
    And request contains "incident_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Deletes an existing selection returns "DELETED" response
    Given new "DeleteUserDefinedFieldSelection" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "selection_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Update an existing selection returns "OK" response
    Given new "PatchUserDefinedFieldSelection" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "selection_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Get all services for the given incident returns "OK" response
    Given new "GetServicesForIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Add a new service to an incident returns "CREATED" response
    Given new "AddServiceToIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Remove a service from an incident returns "DELETED" response
    Given new "RemoveServiceFromIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "service_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Get all teams for the given incident returns "OK" response
    Given new "GetTeamsForIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Add a new team to an incident returns "CREATED" response
    Given new "AddTeamToIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Remove a team from an incident returns "DELETED" response
    Given new "RemoveTeamFromIncident" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "team_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Get all timeline cells for the given incident returns "OK" response
    Given new "GetIncidentTimelineCells" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Create a new timeline cell returns "CREATED" response
    Given new "CreateIncidentTimelineCell" request
    And request contains "incident_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Deletes an existing timeline cell returns "DELETED" response
    Given new "DeleteIncidentTimelineCell" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "cell_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Get the details of an incident timeline cell returns "OK" response
    Given new "GetIncidentTimelineCell" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "cell_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update an existing timeline cell returns "OK" response
    Given new "PatchIncidentTimelineCell" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "cell_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Get all to-dos for the given incident returns "OK" response
    Given new "GetIncidentTodos" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Create a new to-do returns "CREATED" response
    Given new "CreateIncidentTodo" request
    And request contains "incident_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Deletes an existing to-do for the given incident returns "DELETED" response
    Given new "DeleteIncidentTodo" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "todo_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Get the details of an incident to-do returns "OK" response
    Given new "GetIncidentTodo" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "todo_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update an existing to-do returns "OK" response
    Given new "PatchIncidentTodo" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "todo_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Creates the initial user-defined fields returns "OK" response
    Given new "InitializedUserDefinedFields" request
    When the request is sent
    Then the response status is 201 OK

  @generated @skip
  Scenario: Delete an existing choice for a user-defined field returns "DELETED" response
    Given new "DeleteUserDefinedFieldChoice" request
    And request contains "field_id" parameter from "<PATH>"
    And request contains "choice_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Get the details of a user-defined field choice returns "OK" response
    Given new "GetUserDefinedFieldChoice" request
    And request contains "field_id" parameter from "<PATH>"
    And request contains "choice_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update an existing choice for a user-defined field returns "OK" response
    Given new "PatchUserDefinedFieldChoice" request
    And request contains "field_id" parameter from "<PATH>"
    And request contains "choice_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Search incidents returns "OK" response
    Given new "SearchIncidents" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Get all metadata items for an incident returns "OK" response
    Given new "GetIncidentIntegrationMetadatas" request
    And request contains "incident_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Create a new integration metadata returns "CREATED" response
    Given new "CreateIncidentIntegrationMetadata" request
    And request contains "incident_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Deletes a metadata for the given incident returns "DELETED" response
    Given new "DeleteIncidentIntegrationMetadata" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "metadata_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Get details of a metadata returns "OK" response
    Given new "GetIncidentIntegrationMetadata" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "metadata_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update an existing integration metadata returns "OK" response
    Given new "PatchIncidentIntegrationMetadata" request
    And request contains "incident_id" parameter from "<PATH>"
    And request contains "metadata_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Search tag keys for use in user-defined fields returns "OK" response
    Given new "GetTagKeys" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Search tag values for a key to select on incidents returns "OK" response
    Given new "GetTagValues" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update an existing user-defined field returns "OK" response
    Given new "PatchUserDefinedField" request
    And request contains "field_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Create a new user-defined field returns "Created" response
    Given new "CreateUserDefinedField" request
    And body {}
    When the request is sent
    Then the response status is 201 Created
