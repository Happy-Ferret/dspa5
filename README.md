# Features

* Pipelined and caching, fast even on a Pi
* Broadcasting (whole network, house, room)
* Auto discovery of speakers


# Implementation
* Go
* Flite arctic slt (via ansible, not included -- ser binary via env)
* GRPC
* SD01
* https://github.com/faiface/beep for playback (instead of sox play) if possible


# Environment variables
```
DSPA_DATA_DIR=/var/cache/dspa
DSPA_TTS_CMD="synth.sh -o"
DSPA_FILE_EXT=wav
DSPA_PLAY_CMD=play
```

Equivalent calls using bash:

```
echo "Hello, world!" | synth.sh -o output.wav
play output.wav

```
Given filename is guaranteed not to exist and will be temporary. It will be
copied atomically to the cache and played from there.


# Possible further work

* Chrome extension like before
* Server methods to subscribe to WAV (or MP3) files to announce in browser (DSPA_NO_SPEAK)
* Controller to broadcast to discovered nodes, or a broadcast flag to broadcast to subnet


# vendoring

```
go get github.com/tools/godep
godep save ./...
```

magically ends up in vendor/
