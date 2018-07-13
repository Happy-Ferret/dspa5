cache GC (limit to 1GB, LRU mtime) (tmpreaper?)
client accept input from stdin (level still an arg)
somehow vendor go-bindata (add to gopkg required)
getname method? (maybe a setname method too, or just load from env)
split up modules into more than one main file where appropriate
ansible -- make broadcaster optional (or disabled/endabled) via vars

# Undecided
broadcaster -- hard defined hosts file? arg to enable or disable sd01?
repeat message for critical?
say warning/critical/error on that level?
dspa-scan -- look for dspa servers, save to .dspamanifest -- then dspa-broadcaster looks in there as alternative to broadcaster or broadcast flag (showing names)
speak.sh do sox effects pitch -290
speaker + broadcaster could be same executable. Broadcaster could be implicitly running on all speakers... (super simple depoyment!) (or not for centralised deployment)
make rule to compile TTS engine and embed into exe? (or recommend engines)
centralised synthesis? (would make deployments less demanding)



# centralised concept

Centralisation advantages:

1. Only server has to have working TTS engine
1. Easy to integrate with other speaker systems (sonos, web extension)


To retain convenient service discovery model (speakers are endpoints) it would
be necessary to push the necessary audio fragments.

Potential (concurrent) protocol:

1. Server synthesises one fragment at a time like current.
1. Server attempts to stream each fragment (by CAS reference) to each speaker
   as available, ahead of time. Speaker can abort stream if fragment is known.
1. Server maintains concurrent stream to each speaker sending synchronised
   fragment references to play -- each fragment when the server knows all
   speakers have said fragment.

