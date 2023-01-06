Feature: update hub

    Background: basic background
        Given a signed in "admin"
		And a background

	# authenticate
    Scenario Outline: authenticate when update hub
        Given a signed in "<role>"
        When user update hub
        Then returns "<status code>" status code

        Examples:
            | role           | status code |
            | admin          |             |

	# update hub
    Scenario: update hub
        When user update hub
        Then returns "OK" status code
        And updated hub set as expected

	# update invalid hub
    Scenario: update invalid hub
        Given hub is deleted
        When user update hub
        Then returns "NotFound" status code
