@endpoint(logs-archives)
Feature: Logs Archives
  Archives forward all the logs ingested to a cloud storage system.  See the
  [Archives Page](https://app.datadoghq.com/logs/pipelines/archives) for a
  list of the archives currently configured in our UI.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "LogsArchives" API

  @generated @skip
  Scenario: Get all archives returns "OK" response
    Given new "ListLogsArchives" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Create an archive returns "OK" response
    Given new "CreateLogsArchive" request
    And body {}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Delete an archive returns "OK" response
    Given new "DeleteLogsArchive" request
    When the request is sent
    Then the response status is 204 OK

  @generated @skip
  Scenario: Get an archive returns "OK" response
    Given new "GetLogsArchive" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update an archive returns "OK" response
    Given new "UpdateLogsArchive" request
    And body {}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Revoke role from an archive returns "OK" response
    Given new "RemoveRoleFromArchive" request
    And body {}
    When the request is sent
    Then the response status is 204 OK

  @generated @skip
  Scenario: List read roles for an archive returns "OK" response
    Given new "ListArchiveReadRoles" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Grant role to an archive returns "OK" response
    Given new "AddReadRoleToArchive" request
    And body {}
    When the request is sent
    Then the response status is 204 OK
