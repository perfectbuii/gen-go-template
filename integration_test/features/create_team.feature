Feature: Create team

    Background: basic background
        Given a signed in "admin"
        And a background

	# authenticate
    Scenario Outline: authenticate when create team
        Given a signed in "<role>"
        When user create team
        Then returns "<status code>" status code

        Examples:
            | role           | status code |
            | admin          |             |

	# create team
    Scenario: create team
        When user create team
        Then returns "OK" status code
        And team must be created
