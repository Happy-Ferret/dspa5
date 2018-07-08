cache GC (limit to 1GB, LRU mtime) (tmpreaper?)
sd01 in broadcaster
client accept input from stdin (level still an arg)
error flag in fragment to return error status to client
somehow vendor go-bindata
getname method? (maybe a setname method too, or just load from env)
split up modules into more than one main file where appropriate
readme -- recommend that `dspa` resolves to broadcaster or single speaker.
auto detection of OS/arch for binary download
dspa-client -- default to "dspa" host

# Undecided
repeat message for critical?
say warning/critical/error on that level?
dspa-broadcaster -- via sd01 and/or file and eventually also to ioterminals
dspa-scan -- look for dspa servers, save to .dspa-manifest.txt -- then dspa-client looks in there as alternative to broadcaster or broadcast flag
speak.sh do sox effects pitch -290
speaker + broadcaster could be same executable. Broadcaster could be implicitly running on all speakers... (super simple depoyment!)

