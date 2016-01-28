/** Author: Charney Kaye */

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
