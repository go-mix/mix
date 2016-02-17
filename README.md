# Ontomix 

[![Build Status](https://travis-ci.org/outrightmental/ontomix.svg?branch=master)](https://travis-ci.org/outrightmental/ontomix)

http://gopkg.in/outrightmental/ontomix.v0

#### Go-native audio mixer for Music apps

See `example/808.go`:

    package main
    
    import (
      "fmt"
      "os"
      "math/rand"
      "time"
      
      "github.com/outrightmental/ontomix"
      "github.com/outrightmental/ontomix/bind"
    )
    
    var (
      sampleHz   = float64(48000)
      spec = bind.AudioSpec{
        Freq:     sampleHz,
        Format:   bind.AudioF32,
        Channels: 2,
        }
      bpm        = 120
      step       = time.Minute / time.Duration(bpm*4)
      loops      = 16
      prefix     = "assets/sounds/percussion/808/"
      kick1      = "kick1.wav"
      kick2      = "kick2.wav"
      marac      = "maracas.wav"
      snare      = "snare.wav"
      hitom      = "hightom.wav"
      clhat      = "cl_hihat.wav"
      pattern    = []string{
        kick2,
        marac,
        clhat,
        marac,
        snare,
        marac,
        clhat,
        kick2,
        marac,
        marac,
        hitom,
        marac,
        snare,
        kick1,
        clhat,
        marac,
      }
    )
    
    func main() {
      defer atomix.Teardown()    
      atomix.Debug(true)
      atomix.Configure(spec)
      atomix.SetSoundsPath(prefix)
      atomix.StartAt(time.Now().Add(1 * time.Second))
    
      t := 1 * time.Second // padding before music
      for n := 0; n < loops; n++ {
        for s := 0; s < len(pattern); s++ {
          atomix.SetFire(pattern[s], t+time.Duration(s)*step, 0, 1.0, rand.Float64() * 2 - 1)
        }
        t += time.Duration(len(pattern)) * step
      }
    
      atomix.OpenAudio()
    
      fmt.Printf("Ontomix, pid:%v, spec:%v\n", os.Getpid(), spec)
      for atomix.FireCount() > 0 {
        time.Sleep(1 * time.Second)
      }
    }

Run the above from the root of the project with:

    make example

### What?

Game audio mixers are designed to play audio spontaneously, but when the timing is known in advance (e.g. sequence-based music apps) there is a demand for much greater accuracy in playback timing.

Read the API documentation at [godoc.org/github.com/outrightmental/ontomix](https://godoc.org/github.com/outrightmental/ontomix)

**Ontomix** seeks to solve the problem of audio mixing for the purpose of the playback of sequences where audio files and their playback timing is known in advance.
 
Ontomix stores and mixes audio in native Go `[]float64` and natively implements Paul Vögler's "Loudness Normalization by Logarithmic Dynamic Range Compression" (details below)

Author: [Charney Kaye](http://w.charney.io)

#### NOTICE: THIS PROJECT IS IN ALPHA STAGE, AND THE API MAY BE SUBJECT TO CHANGE.

Best efforts will be made to preserve each API version in a release tag that can be parsed, e.g. **[gopkg.in/outrightmental/ontomix.v0](http://gopkg.in/outrightmental/ontomix.v0)** 

### Why?

Even after selecting a hardware interface library such as [PortAudio](http://www.portaudio.com/) or [C++ SDL 2.0](https://www.libsdl.org/), there remains a critical design problem to be solved.

This design is a **music application mixer**. Most available options are geared towards Game development.

Game audio mixers offer playback timing accuracy +/- 2 milliseconds. But that's totally unacceptable for music, specifically sequence-based sample playback.

The design pattern particular to Game design is that the timing of the audio is not know in advance- the timing that really matterns is that which is assembled in near-real-time in response to user interaction.

In the field of Music development, often the timing is known in advance, e.g. a ***sequencer**, the composition of music by specifying exactly how, when and which audio files will be played relative to the beginning of playback.

Ergo, **atomix** seeks to solve the problem of audio mixing for the purpose of the playback of sequences where audio files and their playback timing is known in advance. It seeks to do this with the absolute minimal logical overhead on top of the audio interface.

Ontomix takes maximum advantage of Go by storing and mixing audio in native Go `[]float64` and natively implementing Paul Vögler's "Loudness Normalization by Logarithmic Dynamic Range Compression"

### Time

To the Ontomix API, time is specified as a time.Duration-since-epoch, where the epoch is the moment that atomix.Start() was called.

Internally, time is tracked as samples-since-epoch at the master out playback frequency (e.g. 48000 Hz). This is most efficient because source audio is pre-converted to the master out playback frequency, and all audio maths are performed in terms of samples.

### The Mixing Algorithm

Insipired by the theory paper "Mixing two digital audio streams with on the fly Loudness Normalization by Logarithmic Dynamic Range Compression" by Paul Vögler, 2012-04-20. A .PDF has been included [here](docs/LogarithmicDynamicRangeCompression-PaulVogler.pdf), from the paper originally published [here](http://www.voegler.eu/pub/audio/digital-audio-mixing-and-normalization.html).

### Usage

There's an example implementation of **ontomix** included in the `example/` folder in this repository. Run it using the defaults:

    go run 808.go
    
Or specify options, e.g. using SDL for playback

    go run 808.go --playback sdl

To show the help screen:

    go run 808.go --help

### Dependencies

#### SDL2

Ubuntu:

    sudo apt-get install libsdl2-dev

#### Portaudio

Ubuntu:

    sudo apt-get install portaudio19-dev
