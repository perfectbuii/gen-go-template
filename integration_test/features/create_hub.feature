Feature: Create hub

    # Background: basic background
    #     Given a signed in "admin"
    #     And a background

	# # authenticate
    # Scenario Outline: authenticate when create hub
    #     Given a signed in "<role>"
    #     When user create hub
    #     Then returns "<status code>" status code

    #     Examples:
    #         | role           | status code |
    #         | admin          |             |

	# create hub
    Scenario: create hub
        When user create hub
        Then returns "OK" status code
        And hub must be created
