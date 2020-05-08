#!/usr/bin/env python3
import config
import sys
import selectors
import json
import io
import struct
import socket
import traceback
import time

"""
 ___  ___                               
 |  \/  |                               
 | .  . | ___  ___ ___  __ _  __ _  ___ 
 | |\/| |/ _ \/ __/ __|/ _` |/ _` |/ _ \
 | |  | |  __/\__ \__ \ (_| | (_| |  __/
 \_|  |_/\___||___/___/\__,_|\__, |\___|
                              __/ |     
                             |___/      
"""
class Message:
    """ This is the base class for ClientMessage and ServerMessage. It has all the duplicate 
        functionality from the code we grabbed from realpython.com's tutorial. The tutorial had
        large message classes for client and server, so I created a base message class so that 
        the client and server classes would be much smaller.  
    """
    def __init__(self, selector, sock, addr):
        self.selector = selector
        self.sock = sock
        self.addr = addr
        self._recv_buffer = b""
        self._send_buffer = b""
        self._jsonheader_len = None
        self.jsonheader = None
        

    def _set_selector_events_mask(self, mode):
        """Set selector to listen for events: mode is 'r', 'w', or 'rw'."""
        if mode == "r":
            events = selectors.EVENT_READ
        elif mode == "w":
            events = selectors.EVENT_WRITE
        elif mode == "rw":
            events = selectors.EVENT_READ | selectors.EVENT_WRITE
        else:
            raise ValueError(f"Invalid events mask mode {repr(mode)}.")
        self.selector.modify(self.sock, events, data=self)

    def _read(self):
        try:
            # Should be ready to read
            data = self.sock.recv(4096)
        except BlockingIOError:
            # Resource temporarily unavailable (errno EWOULDBLOCK)
            pass
        else:
            if data:
                self._recv_buffer += data
            else:
                raise RuntimeError("Peer closed.")

    def _write(self):
        if self._send_buffer:
            if config.debug:
                print("sending", repr(self._send_buffer), "to", self.addr)
            try:
                # Should be ready to write

                # Broday
                # Actual sending of information
                sent = self.sock.send(self._send_buffer)
            except BlockingIOError:
                # Resource temporarily unavailable (errno EWOULDBLOCK)
                pass
            else:
                # Broday
                # Overwrite sent message to clear it from buffer
                self._send_buffer = self._send_buffer[sent:]

    def _json_encode(self, obj, encoding):
        return json.dumps(obj, ensure_ascii=False).encode(encoding)

    def _json_decode(self, json_bytes, encoding):
        tiow = io.TextIOWrapper(io.BytesIO(json_bytes), encoding=encoding, newline="")
        obj = json.load(tiow)
        tiow.close()

        return obj

    def _create_message(self, *, content_bytes, content_type, content_encoding):
        jsonheader = {
            "byteorder": sys.byteorder,
            "content-type": content_type,
            "content-encoding": content_encoding,
            "content-length": len(content_bytes),
        }

        jsonheader_bytes = self._json_encode(jsonheader, "utf-8")

        message_hdr = struct.pack(">H", len(jsonheader_bytes))

        message = message_hdr + jsonheader_bytes + content_bytes

        return message

    def process_events(self, mask):
        if mask & selectors.EVENT_READ:
            self.read()
        if mask & selectors.EVENT_WRITE:
            self.write()

    def read(self):
        self._read()

        if self._jsonheader_len is None:
            self.process_protoheader()

        if self._jsonheader_len is not None:
            if self.jsonheader is None:
                self.process_jsonheader()

        self.spec_read()

    def spec_read(self):
        """ Specific Read: This is an abstract method that each client and server must
            implement, since they do slightly different things for each read. 
        """
        pass

    def write(self):
        """ Each client and server do different things for a write(), so this again
            is an abstract method that gets implemented by the child.
        """
        pass

    def close(self):
        if config.debug:
            print("closing connection to", self.addr)
        try:
            self.selector.unregister(self.sock)
        except Exception as e:
            print(
                f"error: selector.unregister() exception for",
                f"{self.addr}: {repr(e)}",
            )

        try:
            self.sock.close()
        except OSError as e:
            print(
                f"error: socket.close() exception for", f"{self.addr}: {repr(e)}",
            )
        finally:
            # Delete reference to socket object for garbage collection
            self.sock = None

    def process_protoheader(self):
        hdrlen = 2

        if len(self._recv_buffer) >= hdrlen:
            self._jsonheader_len = struct.unpack(">H", self._recv_buffer[:hdrlen])[0]
            self._recv_buffer = self._recv_buffer[hdrlen:]

    def process_jsonheader(self):
        hdrlen = self._jsonheader_len
        
        if len(self._recv_buffer) >= hdrlen:
            self.jsonheader = self._json_decode(self._recv_buffer[:hdrlen], "utf-8")

            self._recv_buffer = self._recv_buffer[hdrlen:]

            for reqhdr in (
                "byteorder",
                "content-length",
                "content-type",
                "content-encoding",
            ):
                if reqhdr not in self.jsonheader:
                    raise ValueError(f'Missing required header "{reqhdr}".')


    ##############################################################################################
    def process_server_request(self):
        content_len = self.jsonheader["content-length"]

        if not len(self._recv_buffer) >= content_len:
            return

        data = self._recv_buffer[:content_len]

        self._recv_buffer = self._recv_buffer[content_len:]

        if self.jsonheader["content-type"] == "text/json":
            encoding = self.jsonheader["content-encoding"]

            self.request = self._json_decode(data, encoding)

            print("received request", repr(self.request), "from", self.addr)
        else:
            # Binary or unknown content-type
            self.request = data
            print(
                f'received {self.jsonheader["content-type"]} request from',
                self.addr,
            )
        # Set selector to listen for write events, we're done reading.
        self._set_selector_events_mask("w")



""" 
  _____                         ___  ___                               
 /  ___|                        |  \/  |                               
 \ `--.  ___ _ ____   _____ _ __| .  . | ___  ___ ___  __ _  __ _  ___ 
  `--. \/ _ \ '__\ \ / / _ \ '__| |\/| |/ _ \/ __/ __|/ _` |/ _` |/ _ \
 /\__/ /  __/ |   \ V /  __/ |  | |  | |  __/\__ \__ \ (_| | (_| |  __/
 \____/ \___|_|    \_/ \___|_|  \_|  |_/\___||___/___/\__,_|\__, |\___|
                                                             __/ |     
                                                            |___/      
"""
class ServerMessage(Message):
    """ServerMessage:
    Extends: Message
    Description: Adds necessary server message functionality. In our case, packaging a response
                 and interacting with mongo db. 
    """

    # Broday
    # Add integer response to client guess to list of args. 
    # 1 signifies a guess that is too high
    # 0 signified a guess that is correct
    # -1 signifies a guess that is too low
    def __init__(self, selector, sock, addr, target, guess_response=None):
        super().__init__(selector, sock, addr)  # call parent constructor
        self.guess_response = guess_response    
        self.request = None
        self.response_created = False

        # Broday added these
        self.target = target
        self.result = None

    def write(self):
        if self.request:
            if not self.response_created:
                self.create_response()
        self._write()
    
    def spec_read(self):
        if self.jsonheader:
            if self.request is None:
                # Broday
                # Changed from original process_request() method
                # to a class-specific implementation of process_server_request()

    #############################################################################################            
                self.process_server_request()

    #############################################################################################


    '''
    Broday wrote this
    Override the process_server_request method in the base class to implement some of
    the integer guessing game functionality. process_server_request now serves as the
    de facto comparison method that checks to see if the client guessed the target
    number held by the server.

    '''
    def process_server_request(self):
        content_len = self.jsonheader["content-length"]

        if not len(self._recv_buffer) >= content_len:
            return

        # Broday
        # read in the data from the receive buffer
        data = self._recv_buffer[:content_len]

        # Broday
        # clear the data out of the receive buffer now that we've read it
        # in and stored it
        self._recv_buffer = self._recv_buffer[content_len:]

        if self.jsonheader["content-type"] == "text/json":
            encoding = self.jsonheader["content-encoding"]

            # The request passed by the client
            self.request = self._json_decode(data, encoding)

            # Broday
            # Check if value is in the client's request
            if 'value' in self.request:
                print(self.request['value'])

                # Convert from string to int 
                guess = int(self.request['value'])

                # Check the client's guess vs the server's target number
                # Guess is correct
                if guess == self.target:
                    self.result = 0
                # Guess is low
                elif guess < self.target:
                    self.result = -1
                # Guess is high
                else:
                    self.result = 1
            else:
                print('Invalid value argument. Pass an integer value.')

            print("received request", repr(self.request), "from", self.addr)

        else:
            # Binary or unknown content-type
            self.request = data
            print(
                f'received {self.jsonheader["content-type"]} request from',
                self.addr,
            )
        # Set selector to listen for write events, we're done reading.
        self._set_selector_events_mask("w")

    def create_response(self):
        
        if self.jsonheader['content-type'] == 'text/json':
            # Building response to send to client
            content_encoding = 'utf-8'
            response = {
                # Broday
                # This is where we encode the result
                'content_bytes': self._json_encode(self.result, content_encoding),
                'content_type': 'text/json',
                'content_encoding': content_encoding,
            }
        else:
            # Binary or unknown content-type
            response = {
                'content_bytes': b"First 10 bytes of request: "
                + self.request[:10],
                'content_type': 'binary/custom-server-binary-type',
                'content_encoding': 'binary'
            }

        # Message created and ready to send
        message = self._create_message(**response)
        self.response_created = True
        self._send_buffer += message

"""
  _____ _ _            _  ___  ___                               
 /  __ \ (_)          | | |  \/  |                               
 | /  \/ |_  ___ _ __ | |_| .  . | ___  ___ ___  __ _  __ _  ___ 
 | |   | | |/ _ \ '_ \| __| |\/| |/ _ \/ __/ __|/ _` |/ _` |/ _ \
 | \__/\ | |  __/ | | | |_| |  | |  __/\__ \__ \ (_| | (_| |  __/
  \____/_|_|\___|_| |_|\__\_|  |_/\___||___/___/\__,_|\__, |\___|
                                                       __/ |     
                                                      |___/      
"""
class ClientMessage(Message):
    def __init__(self, selector, sock, addr, request, close_on_response=True):
        super().__init__(selector, sock, addr)
        self._request_queued = False
        self.request = request
        self.response = None

        # Broday
        self.close_on_response = close_on_response

        if self.jsonheader:
            if self.response is None:
                self.process_response()

    def spec_read(self):
        if self.jsonheader:
            if self.response is None:
                self.process_response()

    def write(self):
        if not self._request_queued:
            self.queue_request()

        self._write()

        if self._request_queued:
            if not self._send_buffer:
                # Set selector to listen for read events, we're done writing.
                self._set_selector_events_mask("r")

    def queue_request(self):
        content = self.request["content"]
        content_type = self.request["type"]
        content_encoding = self.request["encoding"]

        if content_type == "text/json":
            req = {
                "content_bytes": self._json_encode(content, content_encoding),
                "content_type": content_type,
                "content_encoding": content_encoding,
            }
        else:
            req = {
                "content_bytes": content,
                "content_type": content_type,
                "content_encoding": content_encoding,
            }

        message = self._create_message(**req)

        self._send_buffer += message
        self._request_queued = True

    def process_response(self):
        content_len = self.jsonheader["content-length"]

        if not len(self._recv_buffer) >= content_len:
            return

        data = self._recv_buffer[:content_len]

        self._recv_buffer = self._recv_buffer[content_len:]

        if self.jsonheader["content-type"] == "text/json":
            encoding = self.jsonheader["content-encoding"]
            self.response = self._json_decode(data, encoding)

            if config.debug:
                print("received response", repr(self.response), "from", self.addr)

            #self._process_response_json_content()
        else:
            # Binary or unknown content-type
            self.response = data
            print(f'Unknown content type: received {self.jsonheader["content-type"]} response from', self.addr,)
            #self._process_response_binary_content()
        
        if self.close_on_response:
            # Close when response has been processed
            self.close()

