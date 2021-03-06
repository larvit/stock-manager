#!/usr/bin/env python3

import os
import requests
import sys

def addStockItems(amount):
	res = requests.post(
		os.environ["API_URL"] + '/add-stock-items',
		json=[{'balance': amount, 'id': 'cli-product', 'name': 'CLI product'}]
	)

	if res.status_code != 200:
		print("๐จ Could not add stock items, API responded with " + res.status_code)
		exit(1)

	print("โ Added " + str(amount) + " items")

def checkCLIParams():
	if len(sys.argv) == 1:
		print('')
		print('')
		print('Usage:  sm [CMD][Amount]')
		print('')
		print('CMD:')
		print('  I  "Inleverera"')
		print('  S  "Sรคlj"')
		print('  L  "Lager"')
		print('')
		print('Amount must be an integer.')
		print('Amount is ignored if CMD is L.')
		print('')
		exit(0)

	if len(sys.argv) != 2:
		print("๐จ Exactly one command line argument is required")
		exit(1)

	cmd = sys.argv[1][0]
	amount = sys.argv[1][1:]

	if cmd != "S" and cmd != "I" and cmd != "L":
		print("๐จ CMD must be S, I or L")
		exit(1)

	if (amount.isdigit() == False or amount == "0") and cmd != "L":
		print("๐จ Amount must be a positive integer")
		exit(1)

def checkEnv():
	if "API_URL" not in os.environ:
		print("๐จ API_URL environment variable is required, but is not set", file = sys.stderr)
		exit(1)

def showStockAmount():
	res = requests.get(os.environ["API_URL"] + '/stock-balances?ids[]=cli-product')
	if res.status_code != 200:
		print("๐จ Could not get stock items, API responded with " + res.status_code)
		exit(1)

	stockItems = res.json()
	if "cli-product" in stockItems:
		print(stockItems["cli-product"]["balance"])
	else:
		print(0)

def subtractStockItems(amount):
	res = requests.post(
		os.environ["API_URL"] + '/subtract-stock-items',
		json=[{'amount': amount, 'id': 'cli-product'}]
	)

	if res.status_code != 200:
		print("๐จ Could not subtract stock items, API responded with " + res.status_code)
		exit(1)

	print("โ Subtracted " + str(amount) + " items")
