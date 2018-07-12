cache GC (limit to 1GB, LRU mtime) (tmpreaper?)
client accept input from stdin (level still an arg)
somehow vendor go-bindata (add to gopkg required)
getname method? (maybe a setname method too, or just load from env)
split up modules into more than one main file where appropriate
ansible -- make broadcaster optional (or disabled/endabled) via vars

# Undecided
broadcaster -- hard defined host? arg to enable or disable sd01?
repeat message for critical?
say warning/critical/error on that level?
dspa-scan -- look for dspa servers, save to .dspa-manifest.txt -- then dspa-broadcaster looks in there as alternative to broadcaster or broadcast flag (showing names)
speak.sh do sox effects pitch -290
speaker + broadcaster could be same executable. Broadcaster could be implicitly running on all speakers... (super simple depoyment!)
