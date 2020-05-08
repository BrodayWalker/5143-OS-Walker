#!/usr/bin/env python3
"""Api.py

"""
from pymongo import MongoClient
import pprint
import json
import datetime


def logg(message):
    """ Simple log function to log errors to a file.
    """
    f = open("logfile.log","a")
    f.write(message+"\n")
    f.close()

class Api(object):
    def __init__(self,db,request):
        self.request = request
        self.action = self.request["action"]
        self.mongo = MongoHelper(db)
        if self.request.get("collection") != None:
            self.mongo.setCollection(self.request.get("collection"))

    def processRequest(self):
        """ This method takes the "request" dictionary (a constructor parameter) and determines
            what to do with the data inside. Here are some request examples:

            0) test - requires the following: 

            Example Dictionary:
            request = {
                "action": "test"
            }

            1) SearchKey - requires the following:
                action      : This is a search using a key, so "searchkey" would work ;)
                collection  : The collection to look in
                key         : The key "column name" to search
                value       : The value the key should equal 

                Example Dictionary:
                request = {
                    "action": "searchkey",
                    "collection":"stockdata",
                    "key":"stock", 
                    "value":"GOOG"
                }

                From Client in Terminal:

                    ./Client.py action=searchkey collection=info key=Symbol value=GOOG

            2) Search - requires the following:
                action      : This is a search using a mongo object, so "search" would work ;)
                collection  : The collection to look in
                params      : The mongo object using to search with. 

                Example Dictionary:
                request = {
                    "action": "searchkey",
                    "collection":"stockdata",
                    "params":{
                        "Symbol":"GOOG",
                        "Year":2018
                    }
                }

                From Client in Terminal:

                    ./Client.py action=search collection=stockdata params='{"Symbol":"GOOG","Year":2018}'

            3) Insert - requires the following:
                action      : This is an insert, so "insert" would work ;)
                collection  : The collection to add data to
                data        : The data to insert

                Example Dictionary:
                request = {
                    "action": "insert",
                    "collection":"stockdata",
                    "data": {
                        "stock":"GOOG",
                        "price":1000.88,
                        "date":"13 Jan 2018
                    }
                }

                From Client Terminal:
                    ./Client.py action=insert collection=temporary data='{"stock":"GOOG","price":1000.88,"date":"13 Jan 2018"}'
        """
        if self.action == "test":
            return {"results":{"Success":"Your client is communicating with the server."}}

        collection = self.request.get("collection",None)
        if collection == None:
            return {"results":{"Error":"Inserting into mongo needs a specified 'collection'."}}
        
        self.mongo.setCollection(collection)

        if self.action == "insert":
            data = self.request.get("data")
            if data == None:
                return {"results":{"Error":"Inserting into mongo needs 'data'."}}
            result = self.mongo.insert(data)
            content = {"results": result}
        elif self.action == "searchkey":
            key = self.request.get("key")
            value = self.request.get("value")
            if key == None or value == None:
                return {"results":{"Error":"Searching mongo needs a key and value."}}
            result = self.mongo.searchkey(key,value)
            content = {"results": result}
        elif self.action == "search":
            params = self.request.get("params",None)
            if params == None:
                return {"results":{"Error":"Searching mongo needs a params object."}}
            params = json.loads(params)
            result = self.mongo.search(params)
            content = {"results": result}
        else:
            content = {"result": f'Error: invalid action "{action}".'}

        return content


"""
 ___  ___                        _   _      _                 
 |  \/  |                       | | | |    | |                
 | .  . | ___  _ __   __ _  ___ | |_| | ___| |_ __   ___ _ __ 
 | |\/| |/ _ \| '_ \ / _` |/ _ \|  _  |/ _ \ | '_ \ / _ \ '__|
 | |  | | (_) | | | | (_| | (_) | | | |  __/ | |_) |  __/ |   
 \_|  |_/\___/|_| |_|\__, |\___/\_| |_/\___|_| .__/ \___|_|   
                      __/ |                  | |              
                     |___/                   |_|              
"""
class MongoHelper(object):
    """
    constructor
    """
    def __init__(self,db,collection=None):
        self.client = MongoClient('mongodb://localhost:27017')
        self.db = self.client[db]
        self.collection = collection

    def setCollection(self,name):
        self.collection = name

    def insert(self,data=None,collection=None):
        
        uid = None

        if collection != None:
            self.collection = collection

        mongodata = json.loads(data.encode("utf-8"))

        mongodata["last_modified"] = datetime.datetime.utcnow()

        result = self.db[self.collection].insert_one(mongodata)

        uid = str(result.inserted_id)
        if uid != None:
            response = {"success": True,"result_id":uid,"message":f"Inserted 1 item into {self.collection}"}
        else:
            response = {"success": False,"message":f"Failed to instert into {self.collection}"}

        return response

    def searchkey(self,key,value,collection=None):
        result_list = []

        if collection != None:
            self.collection = collection

        if key and value:
            params = {key:value}

        return self.search(params)

    def search(self,params,collection=None):
        result_list = []

        if collection != None:
            self.collection = collection

        try:
            result = self.db[self.collection].find(params)
        except: 
            return {"success": False,"collection":self.collection,"message":f"No results with {params} "}

        
        for row in result:
            row['_id'] = str(row['_id'])
            result_list.append(row)

        if len(result_list) > 0:
            response = {"success": True,"count":len(result_list),"data":result_list,}
        else:
            response = {"success": False,"collection":self.collection,"message":f"No results with params {params} "}

        return response 


    def delete(self,uid):
        pass

"""
 ______ ___   _   __ _____ 
 |  ___/ _ \ | | / /|  ___|
 | |_ / /_\ \| |/ / | |__  
 |  _||  _  ||    \ |  __| 
 | |  | | | || |\  \| |___ 
 \_|  \_| |_/\_| \_/\____/                      
"""
class FakeDataBase(object):
    """
    constructor
    """
    def __init__(self,request):
        self.request = request
        self.action = self.request["action"]

        self.fakeDB = {
            "morpheus": "Follow the white rabbit. \U0001f430",
            "ring": "In the caves beneath the Misty Mountains. \U0001f48d",
            "dogface": "\U0001f43e Playing ball! \U0001f3d0",
            "gorilla":"\U0001F98D",
            "dog":"\U0001F415",
            "camel":"\U0001F42A",
            "rhino":"\U0001F98F"
        }

    def processRequest(self):
        if self.action == "search":
            query = self.request.get("key")
            answer = self.search(query)
            content = {"result": answer}
        elif action == "insert":
            collection = self.request.get("collection")
            data = self.request.get("data")
            answer = self.insert(collection,data)
            content = {"result": answer}
        else:
            content = {"result": f'Error: invalid action "{action}".'}

    def search(self,key=None):
        if not key:
            logg("key = none")
            return {"error":"key is none"}
        if key in self.fakeDB:
            value = self.fakeDB[key]
            return {"response":value}
        else:
            return {"response":{"error":f"Key: {key} not found"}}

    def insert(self,collection=None,data=None):

        if not collection:
            response = {"Error": "collection = None"}
        elif not data:
            response = {"Error": "data = None"}

        response = {"success": {"message": {"insert":f"inserting into {collection}","data":data}}}


