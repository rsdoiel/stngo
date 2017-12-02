#!/usr/bin/env python3

import sys
import os
import ctypes

# NOTE: The shared library has different names on different platform
if sys.platform == 'darwin':
    libstn = ctypes.CDLL("lib/libstn.dynlib")
elif sys.platform == 'linux':
    libstn = ctypes.CDLL("lib/libstn.so")
elif sys.platform == 'windows':
    libstn = ctypes.CDLL("lib/libstn.dll")
else:
    libstn = ctypes.CDLL("lib/libstn.so")

# bring in my Go exported C functions
libstn.Version.restype = ctypes.c_char_p

class stn:
    def Version(self):
        '''Return the version string of stngo'''
        return libstn.Version()

if __name__ == '__main__':
    s = stn()
    print(s.Version())


