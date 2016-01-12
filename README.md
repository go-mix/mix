# Atomix 

[![Build Status](https://travis-ci.org/outrightmental/go-atomix.svg?branch=master)](https://travis-ci.org/outrightmental/go-atomix)

#### Sequence-based mixer for Music apps, built on go-sdl2.

Read the API documentation at [godoc.org/github.com/outrightmental/go-atomix](https://godoc.org/github.com/outrightmental/go-atomix)

**Atomix** seeks to solve the problem of audio mixing on top of bare SDL, specifically for the purpose of the playback of sequences where audio files and their playback timing is known in advance.
 
Though it is called via C bindings by the SDL audio callback, atomix stores and mixes audio in native Go `[]float64`

Built on **[go-sdl2](https://github.com/veandco/go-sdl2)** - Go bindings for the C++ library "Simple DirectMedia Layer" **[SDL 2.0](https://www.libsdl.org/)**

Author: [Charney Kaye](http://w.charney.io)

Copyright 2015 Outright Mental, Inc.

### Why?

For sequence mixing in music application development.

Following principles of modularity and reusability according to [The Unix Philosophy](http://en.wikipedia.org/wiki/Unix_philosophy) and 

[go-sdl2](https://github.com/veandco/go-sdl2) is excellent, yet still leaves a design problem to be solved, that was equally inherent when using the original [C++ SDL 2.0](https://www.libsdl.org/) library.

This design problem is a **music application mixer**. Most available options are geared towards Game development, including the proprietary [SDL_mixer](https://www.libsdl.org/projects/SDL_mixer/) project for which the go-sdl2 team [has also implemented bindings](https://github.com/veandco/go-sdl2/blob/master/sdl_mixer/sdl_mixer.go). The design pattern particular to Game design is that the timing of the audio is not know in advance- the timing that really matterns is that which is assembled in near-real-time in response to user interaction.

In the field of Music development, often the timing is known in advance, e.g. a ***sequencer**, the composition of music by specifying exactly how, when and which audio files will be played relative to the beginning of playback.

Ergo, **atomix** seeks to solve the problem of audio mixing on top of bare SDL, specifically for the purpose of the playback of sequences where audio files and their playback timing is known in advance. It seeks to do this with the absolute minimal logical overhead on top of SDL.

Though it is called via C bindings by the SDL audio callback, atomix stores and mixes audio in native Go `[]float64`

### Time

To the Atomix API, time is specified as a time.Duration-since-epoch, where the epoch is the moment that atomix.Start() was called.

Internally, time is tracked as samples-since-epoch at the master output playback frequency (e.g. 48000 Hz). This is most efficient because source audio is pre-converted to the master output playback frequency, and all audio maths are performed in terms of samples.

### The Mixing Algorithm

Insipired by the theory paper "Mixing two digital audio streams with on the fly Loudness Normalization by Logarithmic Dynamic Range Compression" by Paul Vögler, 2012-04-20. A .PDF has been included [here](docs/LogarithmicDynamicRangeCompression-PaulVogler.pdf), from the paper originally published [here](http://www.voegler.eu/pub/audio/digital-audio-mixing-and-normalization.html).

### Usage

There's an example implementation of **go-sdl2** + **go-atomix** included in the `example/` folder in this repository.
