# Audio Pitch Shifting

Mix now supports pitch shifting audio without manual resampling, allowing you to change the pitch of audio sources dynamically.

## Usage

Use `SetFireWithPitch()` instead of `SetFire()` to enable pitch shifting:

```go
import (
    "time"
    "github.com/go-mix/mix"
)

// Normal playback (with ADSR envelope disabled)
mix.SetFire("sound.wav", time.Duration(0), 0, 1.0, 0, 0, 0, 1.0, 0)

// Pitch shifted up one octave (2.0x pitch), ADSR envelope disabled
mix.SetFireWithPitch("sound.wav", time.Duration(0), 0, 1.0, 0, 0, 0, 1.0, 0, 2.0, 1.0)

// Pitch shifted down one octave (0.5x pitch), ADSR envelope disabled
mix.SetFireWithPitch("sound.wav", time.Duration(0), 0, 1.0, 0, 0, 0, 1.0, 0, 0.5, 1.0)

// Pitch shifted up a perfect fifth (~1.5x pitch), ADSR envelope disabled
mix.SetFireWithPitch("sound.wav", time.Duration(0), 0, 1.0, 0, 0, 0, 1.0, 0, 1.5, 1.0)
```

## API

```go
func SetFireWithPitch(
    source string,          // path to audio file
    begin time.Duration,    // start time
    sustain time.Duration,  // sustain duration (0 = play full file)
    volume float64,         // volume (0 to 1)
    pan float64,            // pan (-1 to +1)
    attack time.Duration,   // ADSR: Attack time (0 = disabled)
    decay time.Duration,    // ADSR: Decay time (0 = disabled)
    sustainLevel float64,   // ADSR: Sustain level (0 to 1, use 1.0 to disable)
    release time.Duration,  // ADSR: Release time (0 = disabled)
    pitch float64,          // pitch multiplier (1.0 = no change)
    timeStretch float64     // time stretch multiplier (currently unused)
) *fire.Fire
```

## Pitch Multiplier

The `pitch` parameter is a multiplier:
- `1.0` = original pitch (no change)
- `2.0` = up one octave (twice the frequency)
- `0.5` = down one octave (half the frequency)
- `1.5` = up a perfect fifth
- `0.75` = down a perfect fourth

## How It Works

The implementation uses sample-rate modification with linear interpolation:
- Higher pitch values increase playback speed (shorter duration)
- Lower pitch values decrease playback speed (longer duration)
- Linear interpolation between samples ensures smooth playback without artifacts

## Note on Time Stretching

Currently, changing pitch also changes playback speed proportionally. True independent time-stretching (changing speed without affecting pitch) would require more advanced algorithms like phase vocoding or time-domain harmonic scaling, which may be added in future versions.

## Examples

### Pitch Shifting a Drum Loop

```go
// Play a drum loop at different pitches for variation
// ADSR parameters: attack=0, decay=0, sustainLevel=1.0, release=0 (disabled)
mix.SetFireWithPitch("drums.wav", 0*time.Second, 0, 1.0, 0, 0, 0, 1.0, 0, 1.0, 1.0)  // original
mix.SetFireWithPitch("drums.wav", 4*time.Second, 0, 1.0, 0, 0, 0, 1.0, 0, 0.9, 1.0)  // slightly lower
mix.SetFireWithPitch("drums.wav", 8*time.Second, 0, 1.0, 0, 0, 0, 1.0, 0, 1.1, 1.0)  // slightly higher
```

### Creating Harmonies

```go
// Play the same sample at different pitches to create harmony
baseTime := 2 * time.Second
mix.SetFireWithPitch("note.wav", baseTime, 0, 0.8, -0.5, 0, 0, 1.0, 0, 1.0, 1.0)   // root (left)
mix.SetFireWithPitch("note.wav", baseTime, 0, 0.8, 0, 0, 0, 1.0, 0, 1.25, 1.0)     // major third (center)
mix.SetFireWithPitch("note.wav", baseTime, 0, 0.8, 0.5, 0, 0, 1.0, 0, 1.5, 1.0)    // perfect fifth (right)
```

### Bass Drop Effect

```go
// Create a bass drop by pitch shifting down over time
for i := 0; i < 10; i++ {
    pitch := 1.0 - float64(i)*0.1  // gradually lower pitch
    mix.SetFireWithPitch("bass.wav", time.Duration(i)*200*time.Millisecond, 
                        200*time.Millisecond, 1.0, 0, 0, 0, 1.0, 0, pitch, 1.0)
}
```
