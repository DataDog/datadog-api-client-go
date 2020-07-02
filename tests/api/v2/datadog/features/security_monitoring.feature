@v2 @security_monitoring @todo
Feature: Security Monitoring
  Detection Rules for generating signals

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "SecurityMonitoring" API

  Scenario: List rules returns "OK" response
    Given new "ListSecurityMonitoringRules" request
    When the request is sent
    Then the response status is 200 OK

  Scenario: Create a detection rule returns "OK" response
    Given new "CreateSecurityMonitoringRule" request
    And body {}
    When the request is sent
    Then the response status is 200 OK

  Scenario: Delete an existing rule returns "OK" response
    Given new "DeleteSecurityMonitoringRule" request
    When the request is sent
    Then the response status is 204 OK

  Scenario: Get a rule's details returns "OK" response
    Given new "GetSecurityMonitoringRule" request
    When the request is sent
    Then the response status is 200 OK

  Scenario: Update an existing rule returns "OK" response
    Given new "UpdateSecurityMonitoringRule" request
    And body {}
    When the request is sent
    Then the response status is 200 OK
