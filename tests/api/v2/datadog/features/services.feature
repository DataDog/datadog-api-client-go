@endpoint(services)
Feature: Services
  Create, update, delete and retrieve your organizations services.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "Services" API

  @generated @skip
  Scenario: Get a list of all services returns "OK" response
    Given new "GetServices" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Create a new service returns "CREATED" response
    Given new "CreateService" request
    And body {}
    When the request is sent
    Then the response status is 201 CREATED

  @generated @skip
  Scenario: Delete an existing service returns "DELETED" response
    Given new "DeleteService" request
    And request contains "service_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 204 DELETED

  @generated @skip
  Scenario: Get details of a service returns "OK" response
    Given new "GetService" request
    And request contains "service_id" parameter from "<PATH>"
    When the request is sent
    Then the response status is 200 OK

  @generated @skip
  Scenario: Update an existing service returns "OK" response
    Given new "PatchService" request
    And request contains "service_id" parameter from "<PATH>"
    And body {}
    When the request is sent
    Then the response status is 200 OK
