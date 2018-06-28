# Requirements

Fast, even on a Pi:
* Pipelined (chime/synth/play)
* cached


# Ideas

# Implementation
* Go
* Flite arctic slt (via ansible, not included -- ser binary via env)
* GRPC
* DS01
* https://github.com/faiface/beep for playback (instead of sox play) if possible


# Environment variables
Path to a script that accepts text to synthesise on STDIN and saves to wav
file given as argument. Arguments can be added, space delimited.
```
DSPA_TTS_CMD="synth.sh -o"
```

Equivalent call using bash:

```
echo "Hello, world!" | synth.sh -o output.wav

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
