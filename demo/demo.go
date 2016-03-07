/** Author: Charney Kaye */

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"gopkg.in/pkg/profile.v1"

	"github.com/go-ontomix/ontomix"
	"github.com/go-ontomix/ontomix/bind"
	"github.com/go-ontomix/ontomix/bind/spec"
)

var (
	out         string
	profileMode string
	sampleHz    = float64(48000)
	specs       = spec.AudioSpec{
		Freq:     sampleHz,
		Format:   spec.AudioF32,
		Channels: 2,
	}
	bpm     = 120
	step    = time.Minute / time.Duration(bpm*4)
	loops   = 8
	prefix  = "sound/808/"
	kick1   = "kick1.wav"
	kick2   = "kick2.wav"
	marac   = "maracas.wav"
	snare   = "snare.wav"
	hitom   = "hightom.wav"
	lotom   = "tom1.wav"
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
	flag.StringVar(&out, "out", "sdl", "playback binding [sdl, portaudio, null] _OR_ [wav] for direct stdout (e.g. >file or |aplay)")
	flag.StringVar(&profileMode, "profile", "", "enable profiling [cpu, mem, block]")
	flag.Parse()

	// CPU/Memory/Block profiling
	if len(profileMode) > 0 {
		out = "null" // TODO: evaluate whether profiling is actually working
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
	bind.UseOutputString(out)
	defer ontomix.Teardown()
	ontomix.Configure(specs)
	ontomix.SetSoundsPath(prefix)

	// setup the music
	t := 1 * time.Second // buffer before music
	for n := 0; n < loops; n++ {
		for s := 0; s < len(pattern); s++ {
			ontomix.SetFire(
				pattern[s], t+time.Duration(s)*step, 0, 1.0, rand.Float64()*2-1)
		}
		t += time.Duration(len(pattern)) * step
	}
	t += 5 * time.Second // buffer after music

	//
	if bind.IsDirectOutput() {
		ontomix.Debug(true)
		for p := time.Duration(0); p <= t; p += t / 4 {
			ontomix.OutputContinueTo(p)
		}
		ontomix.OutputClose()
	} else {
		ontomix.Debug(true)
		ontomix.StartAt(time.Now().Add(1 * time.Second))
		fmt.Printf("Ontomix: 808 Example - pid:%v playback:%v spec:%v\n", os.Getpid(), out, specs)
		for ontomix.FireCount() > 0 {
			time.Sleep(1 * time.Second)
		}
	}

}
