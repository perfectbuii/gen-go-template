Feature: retrieve hub

    Background: basic background
        Given a signed in "admin"
        And a background

	# authenticate
    Scenario Outline: authenticate when retrieve hub
        Given a signed in "<role>"
        When user retrieve hub
        Then returns "<status code>" status code

        Examples:
            | role           | status code |
            | admin          |             |

	# retrieve hub
    Scenario: retrieve hub
        When user retrieve hub
        Then returns "OK" status code
        And our system must return result correctly
	
	# retrieve invalid hub
    Scenario: retrieve invalid hub
        Given hub is deleted
        When user retrieve hub
        Then <%[2]s>returns "NotFound" status code
