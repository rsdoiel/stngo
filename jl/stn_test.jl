#=
This is a wrapper for Julia around the Go based C-shared library libstn.go.
=#

import JSON
include("stn.jl")

function test_version() 
    """Read in version information from ../codemeta.json and makesure stn.version() returns appropriate string"""
    src = '{"version":""}'
    open("../codemeta.json") do f
        src = readstring(f)
    end
    codemeta = JSON.parse(src)
    expected = "v" * codemeta["version"]
    result = stn.version()
    if expected != result
        println("expected ", expected, " got ", result)
        return 1
    end
    return 0
end

#
# Main processing
#

# Pre-test check
error_count = 0
ok = true

error_count += test_version()

println("Tests completed")

# Summarize our test results
if error_count > 0
    println("Total errors ", error_count)
    exit(1)
end
println("PASS")
println("OK, stn_test.jl")

