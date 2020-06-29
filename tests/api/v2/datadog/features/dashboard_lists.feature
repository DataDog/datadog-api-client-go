@v2 @@endpoint(dashboard-lists)
Feature: Dashboard Lists
  Interact with your dashboard lists through the API to make it easier to
  organize, find, and share all of your dashboards with your team and
  organization.

  Background:
    Given a valid "apiKeyAuth" key
    And a valid "appKeyAuth" key
    And an instance of "DashboardLists" API

  @todo
  Scenario: Delete items from a dashboard list leading to OK
    Given new "DeleteDashboardListItems" request
    And body {}
    When the request is sent
    Then the status is 200 OK

  @todo
  Scenario: Get a Dashboard List leading to OK
    Given new "GetDashboardListItems" request
    When the request is sent
    Then the status is 200 OK

  @todo
  Scenario: Add Items to a Dashboard List leading to OK
    Given new "CreateDashboardListItems" request
    And body {}
    When the request is sent
    Then the status is 200 OK

  @todo
  Scenario: Update items of a dashboard list leading to OK
    Given new "UpdateDashboardListItems" request
    And body {}
    When the request is sent
    Then the status is 200 OK
