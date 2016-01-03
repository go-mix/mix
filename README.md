# Atomix 

[![Build Status](https://travis-ci.org/outrightmental/go-atomix.svg?branch=master)](https://travis-ci.org/outrightmental/go-atomix)

Atomix is sequence mixing for Music development.

Read the API documentation at [godoc.org/github.com/outrightmental/go-atomix](https://godoc.org/github.com/outrightmental/go-atomix)

**Atomix** seeks to solve the problem of audio mixing on top of bare SDL, specifically for the purpose of the playback of sequences where audio files and their playback timing is known in advance. It seeks to do this with the absolute minimal logical overhead to SDL, and implement that logic in pure Go, though it is called entirely via C bindings by the SDL audio callback, for the most idiomatic Go approach to solving the **sequence mixing** problem.

Copyright 2015 Outright Mental, Inc.

### Why?

#### Atomix is sequence mixing for Music development.

Following principles of modularity and reusability according to [The Unix Philosophy](http://en.wikipedia.org/wiki/Unix_philosophy) and assuming the usefulness of [go-sdl2](https://github.com/veandco/go-sdl2) there is still a design problem to be solved, that was equally inherent when using the original [C++ SDL 2.0](https://www.libsdl.org/) library.

This design problem is an **audio mixer**. Most available options are geared towards Game development, including the proprietary [SDL_mixer](https://www.libsdl.org/projects/SDL_mixer/) project for which the go-sdl2 team [has also implemented bindings](https://github.com/veandco/go-sdl2/blob/master/sdl_mixer/sdl_mixer.go). The design pattern particular to Game design is that the timing of the audio is not know in advance- the timing that really matterns is that which is assembled in near-real-time in response to user interaction.

In the field of Music development, often the timing is known in advance, e.g. a ***sequencer**, the composition of music by specifying exactly how, when and which audio files will be played relative to the beginning of playback.

Ergo, **atomix** seeks to solve the problem of audio mixing on top of bare SDL, specifically for the purpose of the playback of sequences where audio files and their playback timing is known in advance. It seeks to do this with the absolute minimal logical overhead to SDL, and implement that logic in pure Go, though it is called entirely via C bindings by the SDL audio callback, for the most idiomatic Go approach to solving the **sequence mixing** problem.

### Usage

    T. B. D.

### Development

Testing:

    go get github.com/stretchr/testify/assert
    go test

### Contributing

0. Find an issue that bugs you / open a new one.
1. Discuss.
2. Branch off, commit, test.
3. Make a pull request / attach the commits to the issue.
