#!/usr/bin/env python3

from stn import Stn
import os
import sys
import json

def test_version(t, stn_obj):
    with open("../codemeta.json", mode = "r", encoding = "utf-8") as f:
        src = f.read()
    codemeta = json.loads(src)
    expected = "v" + codemeta["version"]
    result = stn_obj.version()
    if result != expected:
        t.error(f"expected {expected}, got {result}")


#
# Test harness
#
class ATest:
    def __init__(self, test_name, verbose = False):
        self._test_name = test_name
        self._error_count = 0
        self._verbose = False

    def test_name(self):
        return self._test_name

    def is_verbose(self):
        return self._verbose

    def verbose_on(self):
        self._verbose = True

    def verbose_off(self):
        self.verbose = False

    def print(self, *msg):
        if self._verbose == True:
            print(*msg)

    def error(self, *msg):
        fn_name = self._test_name
        self._error_count += 1
        print(f"\t{fn_name}", *msg)

    def error_count(self):
        return self._error_count

class TestRunner:
    def __init__(self, set_name, verbose = False):
        self._set_name = set_name
        self._tests = []
        self._error_count = 0
        self._verbose = verbose

    def add(self, fn, params = []):
        self._tests.append((fn, params))

    def run(self):
        #print(f"Starting {self._set_name}")
        for test in self._tests:
            fn_name = test[0].__name__
            t = ATest(fn_name, self._verbose)
            fn, params = test[0], test[1]
            fn(t, *params)
            error_count = t.error_count()
            if error_count > 0:
                print(f"\t\t{fn_name} failed, {error_count} errors found")
            self._error_count += error_count
        error_count = self._error_count
        set_name = self._set_name
        if error_count > 0:
            print(f"Failed {set_name}, {error_count} total errors found")
            sys.exit(1)
        print("PASS")
        print("Ok", __file__)
        sys.exit(0)


if __name__ == '__main__':
    s = Stn()
    test_runner = TestRunner(os.path.basename(__file__))
    test_runner.add(test_version, [s])
    test_runner.run()


