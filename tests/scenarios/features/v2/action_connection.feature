@endpoint(action-connection) @endpoint(action-connection-v2)
Feature: Action Connection
  Action connections extend your installed integrations and allow you to
  take action in your third-party systems (e.g. AWS, GitLab, and Statuspage)
  with Datadog’s Workflow Automation and App Builder products.  Datadog’s
  Integrations automatically provide authentication for Slack, Microsoft
  Teams, PagerDuty, Opsgenie, JIRA, GitHub, and Statuspage. You do not need
  additional connections in order to access these tools within Workflow
  Automation and App Builder.  We offer granular access control for editing
  and resolving connections.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "ActionConnection" API

  @team:DataDog/workflow-automation-dev
  Scenario: Create a new Action Connection returns "Bad Request" response
    Given new "CreateActionConnection" request
    And body with value {"data":{"type":"action_connection","attributes":{"name":"Cassette Connection","integration":{"type":"AWS","credentials":{"type":"AWSAssumeRole","role":"MyRoleUpdated","account_id":"1"}}}}}
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/workflow-automation-dev
  Scenario: Create a new Action Connection returns "Successfully created Action Connection" response
    Given new "CreateActionConnection" request
    And body with value {"data":{"type":"action_connection","attributes":{"name":"Cassette Connection DELETE_ME","integration":{"type":"AWS","credentials":{"type":"AWSAssumeRole","role":"MyRoleUpdated","account_id":"123456789123"}}}}}
    When the request is sent
    Then the response status is 201 Successfully created Action Connection

  @team:DataDog/workflow-automation-dev
  Scenario: Delete an existing Action Connection returns "Not Found" response
    Given new "DeleteActionConnection" request
    And request contains "connection_id" parameter with value "aaa11111-aa11-aa11-aaaa-aaaaaa111111"
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/workflow-automation-dev
  Scenario: Delete an existing Action Connection returns "Successfully deleted Action Connection" response
    Given there is a valid "action_connection" in the system
    And new "DeleteActionConnection" request
    And request contains "connection_id" parameter from "action_connection.data.id"
    When the request is sent
    Then the response status is 204 The resource was deleted successfully.

  @generated @skip @team:DataDog/workflow-automation-dev
  Scenario: Delete an existing Action Connection returns "The resource was deleted successfully." response
    Given new "DeleteActionConnection" request
    And request contains "connection_id" parameter from "REPLACE.ME"
    When the request is sent
    Then the response status is 204 The resource was deleted successfully.

  @team:DataDog/workflow-automation-dev
  Scenario: Get an existing Action Connection returns "Bad Request" response
    Given new "GetActionConnection" request
    And request contains "connection_id" parameter with value "bad-format"
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/workflow-automation-dev
  Scenario: Get an existing Action Connection returns "Not Found" response
    Given new "GetActionConnection" request
    And request contains "connection_id" parameter with value "aaa11111-aa11-aa11-aaaa-aaaaaa111111"
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/workflow-automation-dev
  Scenario: Get an existing Action Connection returns "Successfully get Action Connection" response
    Given new "GetActionConnection" request
    And request contains "connection_id" parameter with value "cb460d51-3c88-4e87-adac-d47131d0423d"
    When the request is sent
    Then the response status is 200 Successfully get Action Connection

  @team:DataDog/workflow-automation-dev
  Scenario: Update an existing Action Connection returns "Bad Request" response
    Given new "UpdateActionConnection" request
    And request contains "connection_id" parameter with value "cb460d51-3c88-4e87-adac-d47131d0423d"
    And body with value {"data":{"type":"action_connection","attributes":{"name":"Cassette Connection","integration":{"type":"AWS","credentials":{"type":"AWSAssumeRole","role":"MyRoleUpdated","account_id":"1"}}}}}
    When the request is sent
    Then the response status is 400 Bad Request

  @team:DataDog/workflow-automation-dev
  Scenario: Update an existing Action Connection returns "Not Found" response
    Given new "UpdateActionConnection" request
    And request contains "connection_id" parameter with value "aaa11111-aa11-aa11-aaaa-aaaaaa111111"
    And body with value {"data":{"type":"action_connection","attributes":{"name":"Cassette Connection","integration":{"type":"AWS","credentials":{"type":"AWSAssumeRole","role":"MyRoleUpdated","account_id":"123456789123"}}}}}
    When the request is sent
    Then the response status is 404 Not Found

  @team:DataDog/workflow-automation-dev
  Scenario: Update an existing Action Connection returns "Successfully updated Action Connection" response
    Given new "UpdateActionConnection" request
    And request contains "connection_id" parameter with value "cb460d51-3c88-4e87-adac-d47131d0423d"
    And body with value {"data":{"type":"action_connection","attributes":{"name":"Cassette Connection","integration":{"type":"AWS","credentials":{"type":"AWSAssumeRole","role":"MyRoleUpdated","account_id":"123456789123"}}}}}
    When the request is sent
    Then the response status is 200 Successfully updated Action Connection
