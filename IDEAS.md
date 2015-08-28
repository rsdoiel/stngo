
While *shorthand* should remain exceptionally simple I have throught of two additions that might fix the project in spirit.

+ reading in a file, processing it with shorthand before assignment to a label
    + LABEL :{ some_included_file_with_shorthand.txt
    + The inner shorthands should not mutate the outer files shorthand, it should function like a closure over the included file.
+ running an external shell command and assigning the commands standard out to the label
    + LABEL :! some_shell_commands_exec_from_shorthand

A questionable idea is to allow Golang regex match or replace.  While straight forward to implemented it starts pushing shorthand towards a programming language which is not the purpose. Likewise looping and conditionals seem like it should be avoided. Supporting shell evaluations seems like the edge and if you did need regexp processing you could shell out to *sed* at that point.

Should *shorthand* breakout into its own project? Is it more appropriate in my *ws* project?


