@endpoint(network-device-monitoring) @endpoint(network-device-monitoring-v2)
Feature: Network Device Monitoring
  The Network Device Monitoring API allows you to fetch devices and
  interfaces and their attributes. See the [Network Device Monitoring
  page](https://docs.datadoghq.com/network_monitoring/) for more
  information.

  Background:
    Given a valid "apiKeyAuth" key in the system
    And a valid "appKeyAuth" key in the system
    And an instance of "NetworkDeviceMonitoring" API
    And new "ListDevices" request

  @generated @skip @team:DataDog/network-device-monitoring
  Scenario: Get the list of devices returns "Bad Request" response
    When the request is sent
    Then the response status is 400 Bad Request

  @generated @skip @team:DataDog/network-device-monitoring
  Scenario: Get the list of devices returns "OK" response
    When the request is sent
    Then the response status is 200 OK
