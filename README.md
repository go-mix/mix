# Atomix 

[![Build Status](https://travis-ci.org/outrightmental/go-atomix.svg?branch=master)](https://travis-ci.org/outrightmental/go-atomix)

#### Sequence-based mixer for Music apps, built on go-sdl2.

Read the API documentation at [godoc.org/github.com/outrightmental/go-atomix](https://godoc.org/github.com/outrightmental/go-atomix)

**Atomix** seeks to solve the problem of audio mixing on top of bare SDL, specifically for the purpose of the playback of sequences where audio files and their playback timing is known in advance. It seeks to do this with the absolute minimal logical overhead to SDL, and implement that logic in pure Go, though it is called entirely via C bindings by the SDL audio callback, for the most idiomatic Go approach to solving the **sequence mixing** problem.

Built on **[go-sdl2](https://github.com/veandco/go-sdl2)** - Go bindings for the C++ library "Simple DirectMedia Layer" **[SDL 2.0](https://www.libsdl.org/)**

Author: [Charney Kaye](http://w.charney.io)

Copyright 2015 Outright Mental, Inc.

### Why?

For sequence mixing in music application development.

Following principles of modularity and reusability according to [The Unix Philosophy](http://en.wikipedia.org/wiki/Unix_philosophy) and assuming the usefulness of [go-sdl2](https://github.com/veandco/go-sdl2) there is still a design problem to be solved, that was equally inherent when using the original [C++ SDL 2.0](https://www.libsdl.org/) library.

This design problem is an **audio mixer**. Most available options are geared towards Game development, including the proprietary [SDL_mixer](https://www.libsdl.org/projects/SDL_mixer/) project for which the go-sdl2 team [has also implemented bindings](https://github.com/veandco/go-sdl2/blob/master/sdl_mixer/sdl_mixer.go). The design pattern particular to Game design is that the timing of the audio is not know in advance- the timing that really matterns is that which is assembled in near-real-time in response to user interaction.

In the field of Music development, often the timing is known in advance, e.g. a ***sequencer**, the composition of music by specifying exactly how, when and which audio files will be played relative to the beginning of playback.

Ergo, **atomix** seeks to solve the problem of audio mixing on top of bare SDL, specifically for the purpose of the playback of sequences where audio files and their playback timing is known in advance. It seeks to do this with the absolute minimal logical overhead to SDL, and implement that logic in pure Go, though it is called entirely via C bindings by the SDL audio callback, for the most idiomatic Go approach to solving the **sequence mixing** problem.

### Usage

Here's an example implementation of **go-sdl2** + **go-atomix**:

    /** Author: Charney Kaye */
    
    package main
    
    // typedef unsigned char Uint16;
    // void AudioCallback(void *userdata, Uint16 *stream, int len);
    import "C"
    import (
      "time"
    
      "github.com/outrightmental/go-atomix"
      "github.com/veandco/go-sdl2/sdl"
    )
    
    const (
      sampleHz   = 44100
      numSamples = 22050
    )
    
    func main() {
      if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
        log.WithFields(log.Fields{
          "error": err,
        }).Fatal("Cannot init SDL")
        return
      }
        defer func() {
            if r := recover(); r != nil {
          log.WithFields(log.Fields{
            "recover": r,
          }).Warn("Player Recovered")
            }
        sdl.PauseAudio(true)
        atomix.Teardown()
        sdl.Quit()
        }()
    
      var (
        step = 125 * time.Millisecond
        loops = 4
      )
    
      var (
        p808  = "assets/sounds/percussion/808/"
        kick1 = p808 + "kick1.wav"
        kick2 = p808 + "kick2.wav"
        snare = p808 + "snare.wav"
        marac = p808 + "maracas.wav"
      )
    
      atomix.Debug(true)
      atomix.Configure(sdl.AudioSpec{
          Freq:     sampleHz,
          Format:   sdl.AUDIO_U16,
          Channels: 2,
          Samples:  numSamples,
        })
    
      t := 1 * time.Second // padding before music
      for n := 0; n < loops; n++ {
            atomix.Play(kick1, t,            4 *step,  1.0)
            atomix.Play(marac, t + 1 *step,  1 *step,  0.5)
            atomix.Play(snare, t + 4 *step,  4 *step,  0.8)
            atomix.Play(marac, t + 6 *step,  1 *step,  0.5)
            atomix.Play(kick2, t + 7 *step,  4 *step,  0.9)
            atomix.Play(marac, t + 10 *step, 1 *step,  0.5)
            atomix.Play(kick2, t + 10 *step, 4 *step,  0.9)
            atomix.Play(snare, t + 12 *step, 4 *step,  0.8)
            atomix.Play(marac, t + 14 *step, 1 *step,  0.5)
        t += 16 * step
      }
    
      spec := atomix.Spec()
      sdl.OpenAudio(spec, nil)
      sdl.PauseAudio(false)
      log.WithFields(log.Fields{
        "spec": spec,
      }).Info("SDL OpenAudio > Atomix")
    
      time.Sleep(t + 1 * time.Second) // padding after music
    }


### Development

Testing:

    go get github.com/stretchr/testify/assert
    go test

### Contributing

0. Find an issue that bugs you / open a new one.
1. Discuss.
2. Branch off, commit, test.
3. Make a pull request / attach the commits to the issue.
