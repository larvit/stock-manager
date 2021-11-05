*** Settings ***

Library  OperatingSystem
Library  RequestsLibrary

Suite Setup  Create Session  storageApi  %{API_URL}

*** Variables ***


*** Test Cases ***

Check Environment
	Environment Variable Should Be Set  API_URL

Add Stock Items
	&{item_1} =  Create dictionary  balance=${7}  id=test-333242  name=Green football
	&{item_2} =  Create dictionary  balance=${2}  id=test-art-fem  name=Truck door, a bit rusty

	@{body} =  Set Variable  ${item_1}  ${item_2}

	POST On Session  storageApi  /add-stock-items  json=@{body}  expected_status=200

Get Stock Items
	Environment Variable Should Be Set  API_URL

	${params} =  Create Dictionary  ids[]=test-333242

	${res} =  GET On Session  storageApi  /stock-balances  params=${params}  expected_status=200

	Should Be Equal  ${res.json()}[test-333242][balance]   ${7}
	Should Be Equal  ${res.json()}[test-333242][id]        test-333242
	Should Be Equal  ${res.json()}[test-333242][name]      Green football

Subtract Stock Items
	&{item_1} =  Create dictionary  amount=${7}  id=test-333242
	&{item_2} =  Create dictionary  amount=${2}  id=test-art-fem

	@{body} =  Set Variable  ${item_1}  ${item_2}

	POST On Session  storageApi  /subtract-stock-items  json=@{body}  expected_status=200