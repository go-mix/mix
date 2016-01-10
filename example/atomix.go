/** Author: Charney Kaye */

package main

// typedef unsigned char Uint16;
// void AudioCallback(void *userdata, Uint16 *stream, int len);
import "C"
import (
	log "github.com/Sirupsen/logrus"
	"time"

	"github.com/outrightmental/go-atomix"
	"github.com/veandco/go-sdl2/sdl"
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
		step = 150 * time.Millisecond
		loops = 1
	)

	var (
		p808  = "assets/sounds/percussion/808/"
		kick1 = p808 + "kick1.wav"
		// kick2 = p808 + "kick2.wav"
		// snare = p808 + "snare.wav"
		// marac = p808 + "maracas.wav"
	)

	atomix.Debug(true)
	atomix.Configure(sdl.AudioSpec{
			Freq:     sampleHz,
			Format:   sdl.AUDIO_S16,
			Channels: 1,
			Samples:  numSamples,
		})
	atomix.StartAt(time.Now().Add(1 * time.Second))

	t := 1 * time.Second // padding before music
	for n := 0; n < loops; n++ {
        atomix.SetFire(kick1, t,               16 *step, 1.0, 0)
        // atomix.SetFire(marac, t + 1 *step,     1 *step,  0.5, 0)
        // atomix.SetFire(snare, t + 4 *step,     4 *step,  0.8, 0)
        // atomix.SetFire(marac, t + 6 *step,     1 *step,  0.5, 0)
        // atomix.SetFire(kick2, t + 7 *step,     4 *step,  0.9, 0)
        // atomix.SetFire(marac, t + 10 *step,    1 *step,  0.5, 0)
        // atomix.SetFire(kick2, t + 10 *step,    4 *step,  0.9, 0)
        // atomix.SetFire(snare, t + 12 *step,    4 *step,  0.8, 0)
        // atomix.SetFire(marac, t + 14 *step,    1 *step,  0.5, 0)
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
