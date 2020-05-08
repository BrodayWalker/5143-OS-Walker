#!/usr/bin/env python3
"""
ServerClass.py
Description:
    This file is the "controller" talking to "Message.py" (where most the work happens).
    It starts the socket connection and then runs the "loop" in which the server listens
    until killed.
Requires:
    Message.py :: ServerMessage
    
    This line:  `message = ServerMessage(self.sel, conn, addr)` is what gives this file message
    handling capability. 
"""
import config
import sys
import selectors
import json
import io
import struct
import socket
import traceback
import random

from Message import ServerMessage

class Server:
    # Broday changed:
    # No db argument anymore
    def __init__(self, host=None,port=None):
        self.host = host
        self.port = int(port)

        self.test_semaphore = 1

        if not self.host:
            self.host = config.host
        
        if not self.port:
            self.port = int(config.port)

        # Broday
        # The server needs to hold a random number for clients to guess
        # use modulus for now to have a reasonably-size number
        self.target = random.randint(0, 1000)
        print(f"TESTING\n random number: {self.target}")

        # Broday
        # DefaulSelector is an alias to the most efficient implementation available on the
        # current platform
        # Ref: https://docs.python.org/3/library/selectors.html
        self.sel = selectors.DefaultSelector()


    '''
    Broday

    '''
    def make_new_target(self):
        self.target = random.randint(0, 1000) 

    def accept_wrapper(self, sock):
        conn, addr = sock.accept()  # Should be ready to read
        print("accepted connection from", addr)
        
        conn.setblocking(False)

        # Broday
        # Reference Message.py to see the ServerMessage class implementation
        # Get the message
        message = ServerMessage(self.sel, conn, addr, self.target)

        # Broday
        # Register a file object for selection, monitoring it for I/O events
        # This results in a new SelectorKey instance
        self.sel.register(conn, selectors.EVENT_READ, data=message)


    '''
    Broday
    From https://realpython.com/python-sockets/

    The basic order of API calls for the server is as follows:
    1. socket()
    2. bind()
    3. listen()
    4. accept()
    '''
    def run_server(self):
        '''
        Broday

        Steps 1-3 are accomplished no matter what. If there are no incoming requests
        from a client, we will never call accept(). When a client request is received,
        it is accepted by accept_wrapper().
        '''
        #host, port = self.host, int(self.port)

        # Broday
        # 1. Get socket
        # A socket that accepts IPv4 connections using TCP
        lsock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        # Avoid bind() exception: OSError: [Errno 48] Address already in use
        lsock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)

        # 2. Bind host and port
        lsock.bind((self.host, self.port))

        # 3. Listen for incoming connections
        lsock.listen()
        print("listening on", (self.host, self.port))

        lsock.setblocking(False)

        self.sel.register(lsock, selectors.EVENT_READ, data=None)

        # Listen until Ctrl + C
        try:
            # Broday
            # The event loop 
            while True:
                # Broday
                # an event that must be waited for on this file object
                # select returns a list of (key, events) tuples, one
                # for each ready file object

                # Blocks until there are sockets ready for I/O
                events = self.sel.select(timeout=None)

                # events is a list of (key, event) tuples where key is a SelectorKey
                # that contains a fileobj attribute, which is the socket object
                # mask is an event mask of the operations that are ready
                for key, mask in events:
                    # Broday
                    # fileobj is the file object registered
                    if key.data is None:
                        # Broday
                        # Accept the listening connection and register the socket with the selector
                        self.accept_wrapper(key.fileobj)
                    else:
                        # Broday
                        # If key.data is not None, then we are dealing with a client socket
                        # that has already been accepted, so we receive the data and do something
                        # with it

                        '''
                        Broday
                        Thoughts: is this where I want to put the mutex?
                        '''
                        message = key.data

                        '''
                        Broday
                        More thoughts: here the object could be trying to obtain the lock. If it does,
                        it plays the guessing game.
                        '''
                        try:
                            # Broday
                            # Keep in mind that message is an instance of ServerMessage (see accept_wrapper())
                            # process_events(mask) is implemented in the base class Message, which ServerMessage
                            # inherits from. This determines whether read() or write() is invoked.
                            message.process_events(mask)
                        except Exception:
                            print(
                                "main: error: exception for",
                                f"{message.addr}:\n{traceback.format_exc()}",
                            )
                            message.close()
                                              
        except KeyboardInterrupt:
            print("caught keyboard interrupt, exiting")
        finally:
            self.sel.close()
