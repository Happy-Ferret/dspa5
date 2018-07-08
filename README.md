# Features

* Pipelined and caching, fast even on a Pi
* Broadcasting (whole network, house, room)
* Chimes based on severity
* Auto discovery of speakers
* Synchronisation of speech on nearby displays (soon)


# Implementation
* Go
* Flite arctic slt (via ansible, not included -- ser binary via env)
* GRPC for protocol
* [sd01](https://github.com/naggie/sd01) for service discovery
* Go dep for vendoring

# Environment variables

## dspa-speaker
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

* Chrome extension like the original DSPA
* Architecture to subscribe to WAV (or MP3)
* Broadcaster to broadcast to discovered/known nodes

# Deployment

See included ansible role. It is recommended that the local DNS server has
`dspa` resolve to the main broadcaster or single speaker.

