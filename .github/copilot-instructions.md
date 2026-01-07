# Copilot Instructions for go-mix/mix

## Project Overview

Mix is a sequence-based Go-native audio mixer for music applications. It's designed for precise playback timing when audio files and their playback timing is known in advance (e.g., sequencers), unlike game audio mixers which have +/- 2ms accuracy.

Key features:
- Stores and mixes audio in native Go `[]float64`
- Implements Paul Vögler's "Loudness Normalization by Logarithmic Dynamic Range Compression"
- Provides sub-millisecond accuracy for music sequence playback
- Time is specified as `time.Duration` since epoch (when `mix.Start()` was called)

## Tech Stack

- **Language**: Go 1.11+ (currently testing on Go 1.20, 1.21, 1.22)
- **Testing**: `github.com/stretchr/testify`
- **Audio Processing**: 
  - `github.com/krig/go-sox` for audio file processing
  - `github.com/youpy/go-riff` for RIFF/WAV handling
- **System Dependencies**: libsox-dev (on Ubuntu/Debian systems)

## Build and Test Commands

```bash
# Format code
make fmt
# or: go fmt ./...

# Run tests
make test
# or: go get -v ./... && go test ./...

# Run demo (no audio playback)
make demo

# Export demo to WAV file
make demo.wav

# Generate test coverage
make cover
```

## Code Style and Conventions

### General Go Style
- Follow standard Go conventions and idioms
- Use `gofmt` for formatting (enforced via EditorConfig)
- Tabs for indentation in `.go` files (indent_size = 4)
- Add package-level comments for all packages
- Keep functions focused and single-purpose

### Naming Conventions
- Use camelCase for unexported names
- Use PascalCase for exported names
- Package names should be lowercase, single-word

### Documentation
- Add godoc-style comments for all exported functions, types, and constants
- Package comments should describe the purpose and usage
- See existing code for examples (e.g., `lib/mix/mix.go`)

### Testing
- Use testify/assert for assertions
- Test files should be named `*_test.go`
- Test function names should start with `Test`
- Include both positive and negative test cases
- Test panic conditions using `defer func() { recover() }()` pattern

### Error Handling
- Prefer panics for configuration errors that should fail fast
- Use explicit error returns for runtime errors
- Provide clear, descriptive error messages

## Project Structure

```
/
├── mix.go              # Main API entry point
├── mix_test.go         # API tests
├── lib/                # Internal library code
│   ├── fire/           # Audio firing/scheduling
│   ├── mix/            # Core mixing logic
│   └── source/         # Audio source management
├── bind/               # Audio bindings and specifications
│   ├── api.go          # Binding API
│   ├── debug/          # Debug utilities
│   ├── hardware/       # Hardware interface
│   ├── opt/            # Options
│   ├── sample/         # Sample handling
│   ├── sox/            # SoX integration
│   ├── spec/           # Audio specifications
│   └── wav/            # WAV file handling
└── demo/               # Demo application
    └── demo.go         # Example usage
```

## Key Concepts

### Time Handling
- API accepts `time.Duration` since epoch
- Internally tracked as samples-since-epoch at playback frequency (e.g., 48000 Hz)
- Epoch is when `mix.Start()` or `mix.StartAt()` is called

### Audio Specifications
- Defined in `bind/spec` package
- Key parameters: Freq (sample rate), Format (audio format), Channels (mono/stereo)
- Common spec: `{ Freq: 48000, Format: AudioF32, Channels: 2 }`

### Mixing Algorithm
- Based on logarithmic dynamic range compression
- Implemented natively in Go using `[]float64`
- Provides automatic loudness normalization

## Development Workflow

1. Make changes to relevant `.go` files
2. Run `make fmt` to format code
3. Run `make test` to ensure tests pass
4. For audio-related changes, test with `make demo` or `make demo.wav`
5. Add tests for new functionality
6. Ensure all tests pass before committing

## Testing Requirements

- All exported functions should have corresponding tests
- Test coverage should be maintained or improved
- Use `make cover` to view test coverage reports
- Tests should not require external audio files unless in demo/

## Common Patterns

### Setting Up Mix
```go
spec := bind.AudioSpec{
    Freq:     48000,
    Format:   bind.AudioF32,
    Channels: 2,
}
mix.Configure(spec)
defer mix.Teardown()
```

### Scheduling Audio
```go
mix.SetFire(source, begin, sustain, volume, pan)
// source: audio file name (string)
// begin: start time as time.Duration
// sustain: duration to play as time.Duration  
// volume: volume level (float64)
// pan: stereo pan position (float64)
```

### Audio Source Management
- Sources are loaded from a configurable path
- Use `mix.SetSoundsPath()` to set the base directory
- Audio is pre-converted to the main output frequency

## Important Notes

- This is a music sequencing library, not a game audio library
- Timing accuracy is critical - changes should not compromise precision
- The mixing algorithm should not be modified without understanding the theory
- Always clean up with `Teardown()` to prevent resource leaks
- System dependencies (libsox-dev) are required for development
