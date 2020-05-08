#!/usr/bin/env python3
import sys
import selectors
import json
import io
import struct
import socket
import traceback
import random
import time

from Message import ClientMessage

"""
 ______                           _   
 | ___ \                         | |  
 | |_/ /___  __ _ _   _  ___  ___| |_ 
 |    // _ \/ _` | | | |/ _ \/ __| __|
 | |\ \  __/ (_| | |_| |  __/\__ \ |_ 
 \_| \_\___|\__, |\__,_|\___||___/\__|
               | |                    
               |_|                    
"""
class Request:
    """
    This class builds an appropriate json request for the client to send. 
    It makes sure all requests are formatted the same.  
    """
    def __init__(self):
        self.request = {}
        self.request["type"] = "text/json"
        self.request["encoding"] = "utf-8"
        self.request["content"] = {}

    def createRequest(self, **kwargs):
        """ Loops through kwargs and pulls out all key value pairs creating
        a python dictionary to be sent to the server as json
        """
        for k,v in kwargs.items():
            self.request["content"][k] = v
            
        return self.request

"""
  _____ _ _            _   
 /  __ \ (_)          | |  
 | /  \/ |_  ___ _ __ | |_ 
 | |   | | |/ _ \ '_ \| __|
 | \__/\ | |  __/ | | | |_ 
  \____/_|_|\___|_| |_|\__|
                           
"""
class Client:
    def __init__(self, host=None, port=None, debug=False):
        #self.sel = selectors.DefaultSelector()
        self.host = host
        self.port = port
        self.debug = debug
        self.response = None

        if not self.host:
            self.host = config.host
        
        if not self.port:
            self.port = int(config.port)

    def start_connection(self, request):
        sel = selectors.DefaultSelector()

        addr = (self.host, self.port)

        if self.debug:
            print("starting connection to", addr)

        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.setblocking(False)
        # Broday
        # must use connect_ex to avoid BlockingIOError exception
        sock.connect_ex(addr)

        events = selectors.EVENT_READ | selectors.EVENT_WRITE

        message = ClientMessage(sel, sock, addr, request)
    
        sel.register(sock, events, data=message)

        try:
            while True:
                events = sel.select(timeout=1)

                for key, mask in events:
                    message = key.data

                    try:
                        message.process_events(mask)

                    except Exception:
                        print(
                            "main: error: exception for",
                            f"{message.addr}:\n{traceback.format_exc()}",
                        )

                        message.close()

                # Check for a socket being monitored to continue.
                if not sel.get_map():
                    break
        except KeyboardInterrupt:
            print("caught keyboard interrupt, exiting")
        finally:
            sel.close()
            self.response = message.response
            

    def get_response(self):
        return self.response



'''
 $$$$$$\                                           $$$$$$\  $$\ $$\                      $$\     
$$  __$$\                                         $$  __$$\ $$ |\__|                     $$ |    
$$ /  \__|$$\   $$\  $$$$$$\   $$$$$$$\  $$$$$$$\ $$ /  \__|$$ |$$\  $$$$$$\  $$$$$$$\ $$$$$$\   
$$ |$$$$\ $$ |  $$ |$$  __$$\ $$  _____|$$  _____|$$ |      $$ |$$ |$$  __$$\ $$  __$$\\_$$  _|  
$$ |\_$$ |$$ |  $$ |$$$$$$$$ |\$$$$$$\  \$$$$$$\  $$ |      $$ |$$ |$$$$$$$$ |$$ |  $$ | $$ |    
$$ |  $$ |$$ |  $$ |$$   ____| \____$$\  \____$$\ $$ |  $$\ $$ |$$ |$$   ____|$$ |  $$ | $$ |$$\ 
\$$$$$$  |\$$$$$$  |\$$$$$$$\ $$$$$$$  |$$$$$$$  |\$$$$$$  |$$ |$$ |\$$$$$$$\ $$ |  $$ | \$$$$  |
 \______/  \______/  \_______|\_______/ \_______/  \______/ \__|\__| \_______|\__|  \__|  \____/                                                                                                                                                                                                                                                                                              
'''
class GuessClient(Client):
    '''
    Broday

    The GameClient class extends the Client class to give game-specific behavior. 
    The class specifically implements a client that connects to a server, passes an integer
    number guess to the server, receives a response from the server (-1, 0, or 1) for too low,
    correct, and too high respectively. 
    '''
    def __init__(self, host=None, port=None, debug=False):
        super().__init__(host, port, debug)
        # Change later to full size
        # self.guess = random.randint(0, sys.maxsize)
        # self.guess = random.randint(0, 1000)
        self.guess = 500
        self.num_guesses = 0
        self.last_result = -1

        if self.debug == True:
            print(self.host)
            print(self.port)
            print(self.debug)
            print(self.guess)

    '''
    Broday

    Creates an instance of the Request class and returns it. The request will be properly formatted
    and includes an integer number guess.
    '''
    def build_guess(self, last_result=-1, num_guesses=None):
        found = False

        # Adjust guess
        if last_result != 0:
            # If the last guess was low, make it larger
            if last_result == -1:
                self.guess += 1
            # If the last guess was high, make it smaller
            elif last_result == 1:
                self.guess -= 1
            else:
                self.guess = random.randint(0, 1000)
        elif last_result == 0:
            request = None
            found = True

        self.num_guesses += 1 


        request = Request()
        request = request.createRequest(value=self.guess)

        return request, found

    '''
    Broday
    '''
    def start_guessing(self):

        found = False

        while found == False:
            # Build the guess
            guess_request, found = self.build_guess(self.last_result)

            # Start the connection
            if guess_request != None and found != True:
                self.start_connection(guess_request)

                # This will need to be changed for multiple clients
                self.last_result = self.get_response()

                print(f"Last result: {self.last_result}")

                time.sleep(.1)

        
        print('Done guessing')



'''
Broday
'''
class StayConnected(Client):
    def __init__(self, host=None, port=None, debug=False):
        super().__init__(host, port, debug)
        self.guess = 80
        self.num_guesses = 0
        self.last_result = -1

    def build_guess(self, last_result=-1, num_guesses=None):
        found = False

        # Adjust guess
        if last_result != 0:
            # If the last guess was low, make it larger
            if last_result == -1:
                self.guess += 1
            # If the last guess was high, make it smaller
            elif last_result == 1:
                self.guess -= 1
            else:
                self.guess = random.randint(0, 1000)
        elif last_result == 0:
            request = None
            found = True

        self.num_guesses += 1 


        request = Request()
        request = request.createRequest(value=self.guess)

        return request, found


    def connect(self):
        sel = selectors.DefaultSelector()

        addr = (self.host, self.port)

        if self.debug:
            print("starting connection to", addr)

        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.setblocking(False)
        # Broday
        # must use connect_ex to avoid BlockingIOError exception
        sock.connect_ex(addr)

        events = selectors.EVENT_READ | selectors.EVENT_WRITE


        ####
        guess_request, found = self.build_guess(self.last_result)

        message = ClientMessage(sel, sock, addr, guess_request, close_on_response=False)
        ####

    
        sel.register(sock, events, data=message)

        try:
            while True:
                events = sel.select(timeout=1)

                for key, mask in events:
                    message = key.data

                    try:
                        message.process_events(mask)

                    except Exception:
                        print(
                            "main: error: exception for",
                            f"{message.addr}:\n{traceback.format_exc()}",
                        )

                        message.close()

                # Check for a socket being monitored to continue.
                if not sel.get_map():
                    break
        except KeyboardInterrupt:
            print("caught keyboard interrupt, exiting")
        finally:
            sel.close()

    def run_game():
        self.connect()
        
    


    

    

            




        