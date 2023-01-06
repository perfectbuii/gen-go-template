Feature: list hub

    Background: basic background
        Given a signed in "admin"
        And a background

	# authenticate
    Scenario Outline: authenticate when list hub
        Given a signed in "<role>"
        When user list hub
        Then returns "<status code>" status code

        Examples:
            | role           | status code |
            | admin          |             |

	# list hub
    Scenario: list hub
        When user list hub
        Then returns "OK" status code
        And our system must return results correctly
