#!/usr/bin/env python3
import os
from helpers import GetIp

if __name__=='__main__':
    print("Getting ip using `GetHostname`")
    ip1 = GetIp.by_socket()
    print(ip1)

    print("Getting ip using `ifcnfg`")
    ip2 = GetIp.by_ifcfg() 
    print(ip2)

