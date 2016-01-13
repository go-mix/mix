/** Author: Charney Kaye */

package main

// typedef unsigned char Uint16;
// void AudioCallback(void *userdata, Uint16 *stream, int len);
import "C"
import (
	log "github.com/Sirupsen/logrus"
	"github.com/outrightmental/go-atomix"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	sampleHz   = 44100
	numSamples = 4096
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
		bpm   = 127
		step  = time.Minute / time.Duration(bpm*4)
		loops = 4
	)

	var (
		prefix    = "assets/sounds/percussion/808/"
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
			marac,
			marac,
		}
	)

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
	log.WithFields(log.Fields{
		"spec": spec,
	}).Info("SDL OpenAudio > Atomix")

	time.Sleep(t + 1*time.Second) // padding after music
}
