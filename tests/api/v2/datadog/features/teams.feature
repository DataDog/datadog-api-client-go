@endpoint(teams)
Feature: Teams
  Create, update, delete and retrieve your organizations Teams.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "Teams" API

  @generated @skip
  Scenario: Get a list of all teams returns "OK" response
    Given new "GetTeams" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Create a new team returns "CREATED" response
    Given new "CreateTeam" request
    And body {}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Delete an existing team returns "DELETED" response
    Given new "DeleteTeam" request
    And request contains "team_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Get details of a team returns "OK" response
    Given new "GetTeam" request
    And request contains "team_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update an existing team returns "OK" response
    Given new "PatchTeam" request
    And request contains "team_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 200 OK
