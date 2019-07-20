import sys
import os
import ctypes


# Figure out shared library extension
go_basename = 'libstn'
uname = os.uname().sysname
ext = '.so'
if uname == 'Darwin':
    ext = '.dylib'
if uname == 'Windows':
    ext = '.dll'

# Find our shared library and load it
dir_path = os.path.dirname(os.path.realpath(__file__))
lib = ctypes.cdll.LoadLibrary(os.path.join(dir_path, go_basename+ext))

# Setup our Go functions to be nicely wrapped
go_version = lib.version
go_version.restype = ctypes.c_char_p

# Create our Python 3 class for working with library
class Stn:
    def version(self):
        """Return the version string of stngo"""
        value = go_version()
        if not isinstance(value, bytes):
            value = value.encode("utf-8")
        return value.decode() 

