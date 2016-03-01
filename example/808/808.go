/** Author: Charney Kaye */

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"gopkg.in/pkg/profile.v1"

	"gopkg.in/ontomix.v0"
	"gopkg.in/ontomix.v0/bind"
)

var (
	out         bool
	playback    string
	profileMode string
	sampleHz    = float64(48000)
	spec        = bind.AudioSpec{
		Freq:     sampleHz,
		Format:   bind.AudioF32,
		Channels: 2,
	}
	bpm     = 120
	step    = time.Minute / time.Duration(bpm*4)
	loops   = 16
	prefix  = "../_sound/808/"
	kick1   = "kick1.wav"
	kick2   = "kick2.wav"
	marac   = "maracas.wav"
	snare   = "snare.wav"
	hitom   = "hightom.wav"
	clhat   = "cl_hihat.wav"
	pattern = []string{
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
	// command-line arguments
	flag.BoolVar(&out, "out", false, "stdout to byte stream (e.g. >file or |aplay)")
	flag.StringVar(&playback, "playback", "sdl", "out playback binding [sdl, portaudio, null]")
	flag.StringVar(&profileMode, "profile", "", "enable profiling [cpu, mem, block]")
	flag.Parse()

	// CPU/Memory/Block profiling
	if len(profileMode) > 0 {
		playback = "null"
		switch profileMode {
		case "cpu":
			defer profile.Start(profile.CPUProfile).Stop()
		case "mem":
			defer profile.Start(profile.MemProfile, profile.MemProfileRate(4096)).Stop()
		case "block":
			defer profile.Start(profile.BlockProfile).Stop()
		default:
			// do nothing
		}
	}

	// configure ontomix
	bind.UseOutputString(playback)
	defer ontomix.Teardown()
	ontomix.Configure(spec)
	ontomix.SetSoundsPath(prefix)

	// setup the music
	t := 1 * time.Second // buffer before music
	for n := 0; n < loops; n++ {
		for s := 0; s < len(pattern); s++ {
			ontomix.SetFire(pattern[s], t+time.Duration(s)*step, 0, 1.0, rand.Float64()*2-1)
		}
		t += time.Duration(len(pattern)) * step
	}
	t += 2 * time.Second // buffer after music

	//
	if out {
		ontomix.OutputBegin()
		ontomix.OutputContinueTo(t)
		ontomix.OutputClose()
	} else {
		ontomix.Debug(true)
		ontomix.StartAt(time.Now().Add(1 * time.Second))
		fmt.Printf("Ontomix: 808 Example - pid:%v playback:%v spec:%v\n", os.Getpid(), playback, spec)
		for ontomix.FireCount() > 0 {
			time.Sleep(1 * time.Second)
		}
	}

}
