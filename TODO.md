cache GC (limit to 1GB, LRU mtime) (tmpreaper?)
client accept input from stdin (level still an arg)
error flag in fragment to return error status to client
somehow vendor go-bindata (add to gopkg required)
getname method? (maybe a setname method too, or just load from env)
split up modules into more than one main file where appropriate
dspa-client -- default to "dspa" host, warn that DSPA_SPEAKER_ADDR should be set if fails
dspa-broadcaster -- via sd01 and/or file and eventually also to ioterminals
client should say when there's an error'd fragment

# Undecided
repeat message for critical?
say warning/critical/error on that level?
dspa-scan -- look for dspa servers, save to .dspa-manifest.txt -- then dspa-broadcaster looks in there as alternative to broadcaster or broadcast flag
speak.sh do sox effects pitch -290
speaker + broadcaster could be same executable. Broadcaster could be implicitly running on all speakers... (super simple depoyment!)


# sd01
UDPAddr -- irrelevant port in struct
