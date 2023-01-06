Feature: delete hub

    Background: basic background
        Given a signed in "admin"
        And a background

	# authenticate
    Scenario Outline: authenticate when delete hub
        Given a signed in "<role>"
        When user delete hub
        Then returns "<status code>" status code

        Examples:
            | role           | status code |
            | admin          |             |

	# delete hub
    Scenario: delete hub
        When user delete hub
        Then <%[2]s>returns "OK" status code
        And hub have been deleted correctly

	# delete hub again
    Scenario: delete hub again
        Given user delete hub
        When user delete hub again
        Then <%[2]s>returns "NotFound" status code
