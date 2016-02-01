/** Author: Charney Kaye */

package main

import (
	"fmt"
	"github.com/outrightmental/go-atomix"
	"github.com/outrightmental/go-atomix/bind"
	"os"
	"time"
)

var (
	sampleHz   = float64(48000)
	spec = bind.AudioSpec{
		Freq:     sampleHz,
		Format:   bind.AudioF32,
		Channels: 1,
		}
	bpm        = 120
	step       = time.Minute / time.Duration(bpm*4)
	loops      = 1
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
	defer func() {
		atomix.Teardown()
	}()

	atomix.Debug(true)
	atomix.Configure(spec)
	atomix.SetSoundsPath(prefix)
	atomix.StartAt(time.Now().Add(1 * time.Second))

	t := 1 * time.Second // padding before music
	for n := 0; n < loops; n++ {
		for s := 0; s < len(pattern); s++ {
			atomix.SetFire(pattern[s], t+time.Duration(s)*step, 0, 1.0, 0)
		}
		t += time.Duration(len(pattern)) * step
	}

	atomix.OpenAudio()

	fmt.Printf("Atomix, pid:%v, spec:%v\n", os.Getpid(), spec)
	time.Sleep(t + 1*time.Second) // wait until music + 1 second
}
