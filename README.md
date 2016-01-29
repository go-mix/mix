# Atomix 

[![Build Status](https://travis-ci.org/outrightmental/go-atomix.svg?branch=master)](https://travis-ci.org/outrightmental/go-atomix)

#### Sequence-based Go-native audio mixer for Music apps

    package main
    
    // typedef unsigned char Uint16;
    // void AudioCallback(void *userdata, Uint16 *stream, int len);
    import "C"
    import (
      "fmt"
      "github.com/outrightmental/go-atomix"
      "github.com/veandco/go-sdl2/sdl"
      "os"
      "time"
    )
    
    var (
      sampleHz   = int32(44100)
      numSamples = uint16(4096)
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
      if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
        fmt.Printf("Cannot init SDL. Error: %v\n", err)
        return
      }
      defer func() {
        if r := recover(); r != nil {
          fmt.Printf("Player Recovered: %v\n", r)
        }
        sdl.PauseAudio(true)
        atomix.Teardown()
        sdl.Quit()
      }()
    
      atomix.Debug(true)
      atomix.Configure(sdl.AudioSpec{
        Freq:     sampleHz,
        Format:   sdl.AUDIO_F32,
        Channels: 1,
        Samples:  numSamples,
      })
      atomix.SetSoundsPath(prefix)
      atomix.StartAt(time.Now().Add(1 * time.Second))
    
      t := 1 * time.Second // padding before music
      for n := 0; n < loops; n++ {
        for s := 0; s < len(pattern); s++ {
          atomix.SetFire(pattern[s], t+time.Duration(s)*step, 0, 1.0, 0)
        }
        t += time.Duration(len(pattern)) * step
      }
    
      spec := atomix.Spec()
      sdl.OpenAudio(spec, nil)
      sdl.PauseAudio(false)
    
      fmt.Printf("SDL OpenAudio > Atomix, pid:%v, spec:%v\n", os.Getpid(), spec)
      time.Sleep(t + 1*time.Second) // padding after music
    }


### What?

Game audio mixers are designed to play audio spontaneously, but when the timing is known in advance (e.g. sequence-based music apps) there is a demand for much greater accuracy in playback timing.

Read the API documentation at [godoc.org/github.com/outrightmental/go-atomix](https://godoc.org/github.com/outrightmental/go-atomix)

**Atomix** seeks to solve the problem of audio mixing on top of bare SDL, specifically for the purpose of the playback of sequences where audio files and their playback timing is known in advance.
 
Though it is called via C bindings by the SDL audio callback, atomix stores and mixes audio in native Go `[]float64` and natively implements Paul Vögler's "Loudness Normalization by Logarithmic Dynamic Range Compression" (details below)

Built on **[go-sdl2](https://github.com/veandco/go-sdl2)** - Go bindings for the C++ library "Simple DirectMedia Layer" **[SDL 2.0](https://www.libsdl.org/)**

Author: [Charney Kaye](http://w.charney.io)

***NOTICE: THIS PROJECT IS IN ALPHA STAGE, AND THE API MAY BE SUBJECT TO CHANGE.***

Best efforts will be made to preserve each API version in a release tag that can be parsed, e.g. [http://gopkg.in](http://gopkg.in) 

### Why?

Even after selecting a hardware interface library such as [C++ SDL 2.0](https://www.libsdl.org/) via [go-sdl2](https://github.com/veandco/go-sdl2), there remains a critical design problem to be solved.

This design is a **music application mixer**. Most available options are geared towards Game development, including the proprietary [SDL_mixer](https://www.libsdl.org/projects/SDL_mixer/) project for which the go-sdl2 team [has also implemented bindings](https://github.com/veandco/go-sdl2/blob/master/sdl_mixer/sdl_mixer.go).

Game audio mixers offer playback timing accuracy +/- 2 milliseconds. But that's totally unacceptable for music, specifically sequence-based sample playback.

The design pattern particular to Game design is that the timing of the audio is not know in advance- the timing that really matterns is that which is assembled in near-real-time in response to user interaction.

In the field of Music development, often the timing is known in advance, e.g. a ***sequencer**, the composition of music by specifying exactly how, when and which audio files will be played relative to the beginning of playback.

Ergo, **atomix** seeks to solve the problem of audio mixing on top of bare SDL, specifically for the purpose of the playback of sequences where audio files and their playback timing is known in advance. It seeks to do this with the absolute minimal logical overhead on top of SDL.

Atomix takes maximum advantage of Go by storing and mixing audio in native Go `[]float64` and natively implementing Paul Vögler's "Loudness Normalization by Logarithmic Dynamic Range Compression"

### Time

To the Atomix API, time is specified as a time.Duration-since-epoch, where the epoch is the moment that atomix.Start() was called.

Internally, time is tracked as samples-since-epoch at the master output playback frequency (e.g. 48000 Hz). This is most efficient because source audio is pre-converted to the master output playback frequency, and all audio maths are performed in terms of samples.

### The Mixing Algorithm

Insipired by the theory paper "Mixing two digital audio streams with on the fly Loudness Normalization by Logarithmic Dynamic Range Compression" by Paul Vögler, 2012-04-20. A .PDF has been included [here](docs/LogarithmicDynamicRangeCompression-PaulVogler.pdf), from the paper originally published [here](http://www.voegler.eu/pub/audio/digital-audio-mixing-and-normalization.html).

### Usage

There's an example implementation of **go-sdl2** + **go-atomix** included in the `example/` folder in this repository.
