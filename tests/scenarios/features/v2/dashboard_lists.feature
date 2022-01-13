@endpoint(dashboard-lists) @endpoint(dashboard-lists-v2)
Feature: Dashboard Lists
  Interact with your dashboard lists through the API to organize, find, and
  share all of your dashboards with your team and organization.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "DashboardLists" API
    And request contains "dashboard_list_id" parameter from "REPLACE.ME"

  @generated @skip @team:DataDog/dashboards
  Scenario: Add Items to a Dashboard List returns "Bad Request" response
    Given new "CreateDashboardListItems" request
    And body with value {"dashboards": [{"id": "q5j-nti-fv6", "type": "host_timeboard"}]}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/dashboards
  Scenario: Add Items to a Dashboard List returns "Not Found" response
    Given new "CreateDashboardListItems" request
    And body with value {"dashboards": [{"id": "q5j-nti-fv6", "type": "host_timeboard"}]}
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/dashboards
  Scenario: Add Items to a Dashboard List returns "OK" response
    Given new "CreateDashboardListItems" request
    And body with value {"dashboards": [{"id": "q5j-nti-fv6", "type": "host_timeboard"}]}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/dashboards
  Scenario: Delete items from a dashboard list returns "Bad Request" response
    Given new "DeleteDashboardListItems" request
    And body with value {"dashboards": [{"id": "q5j-nti-fv6", "type": "host_timeboard"}]}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/dashboards
  Scenario: Delete items from a dashboard list returns "Not Found" response
    Given new "DeleteDashboardListItems" request
    And body with value {"dashboards": [{"id": "q5j-nti-fv6", "type": "host_timeboard"}]}
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/dashboards
  Scenario: Delete items from a dashboard list returns "OK" response
    Given new "DeleteDashboardListItems" request
    And body with value {"dashboards": [{"id": "q5j-nti-fv6", "type": "host_timeboard"}]}
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/dashboards
  Scenario: Get items of a Dashboard List returns "Not Found" response
    Given new "GetDashboardListItems" request
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/dashboards
  Scenario: Get items of a Dashboard List returns "OK" response
    Given new "GetDashboardListItems" request
    When the request is sent
    Then the response status is 200 OK

  @generated @skip @team:DataDog/dashboards
  Scenario: Update items of a dashboard list returns "Bad Request" response
    Given new "UpdateDashboardListItems" request
    And body with value {"dashboards": [{"id": "q5j-nti-fv6", "type": "host_timeboard"}]}
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/dashboards
  Scenario: Update items of a dashboard list returns "Not Found" response
    Given new "UpdateDashboardListItems" request
    And body with value {"dashboards": [{"id": "q5j-nti-fv6", "type": "host_timeboard"}]}
    When the request is sent
    Then the response status is 404 Not Found

  @generated @skip @team:DataDog/dashboards
  Scenario: Update items of a dashboard list returns "OK" response
    Given new "UpdateDashboardListItems" request
    And body with value {"dashboards": [{"id": "q5j-nti-fv6", "type": "host_timeboard"}]}
    When the request is sent
    Then the response status is 200 OK
