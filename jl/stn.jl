#=
stn.jl is a wrapper around the C-shared library, libstn, of the Go package stn.
=#
module stn
export version

function version()
    value = ccall((:version, "./libstn"), Cstring, (),)
    convert(UTF8String, bytestring(value))
end

end
