#!/usr/bin/env python3
""" Client.py
Look at above shebang and make sure your python is at same location.
You can test by typing `which python3` at your console.

Description:
    Sends a message to a server at indicated ip address and port.
Requires:
    ClientClass.py
Usage:
    Pass in values to the client using key value pairs 
    ./Client.py host=10.0.61.34 port=6000
"""
import config
import sys
from ClientClass import Client
from ClientClass import GuessClient
from ClientClass import Request
from ClientClass import StayConnected
from helpers import myArgParse
import json

def Usage():
    print("Usage: <host> <port> <action>")
    print(f"Example: {sys.argv[0]} host=10.0.61.34 port=6000 action=guess")
    sys.exit()


if __name__ == "__main__":
    """ 
    Main client driver. 
    """
    kwargs, args = myArgParse(sys.argv)
    #print(kwargs,args)

    # get items from command line OR load them from config file
    host = kwargs.get("host", config.host)          # MANDATORY 
    port = int(kwargs.get("port", config.port))     # MANDATORY Port to connect to (XXXXX, e,g, 6000) 

    # Broday
    # Remove collection argument
    # Replace this request in the GuessClient class. GuessClient only makes number guesses
    # so it's request class handling is implemented specifically for that purpose.
    # The request is now built inside the GuestClient class.
    # Old call: request = request.createRequest(action=action, key=key, data=data, value=value, params=params)

    client = GuessClient(host, port)

    # Broday
    # GuessClient starts it's connection in-class so it can keep guessing numbers
    # until it is correct
    client.start_guessing()

    # Testing the client that stays connected
    #client = StayConnected(host, port)
    #client.connect()
    