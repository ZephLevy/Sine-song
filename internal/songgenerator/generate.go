package songgenerator

type note struct {
	duration float64
	name     string
}

func GetSong(sampleRateInt int) []int16 {
	sampleRate := float64(sampleRateInt)

	// This could be made a lot simpler with loops and functions.
	// However, this is only to test the capabilities as a "demo" of sorts
	// It will be a lot more refined when I make my own song for it.
	voice1 := []note{
		{duration: 4, name: "REST"},
		{duration: 1, name: "D5"},
		{duration: 1, name: "B4"},
		{duration: 1, name: "G4"},
		{duration: 1, name: "D4"},
		{duration: (1 / 3.0), name: "E4"},
		{duration: (1 / 3.0), name: "F#4"},
		{duration: (1 / 3.0), name: "G4"},
		{duration: 2 / 3.0, name: "E4"},
		{duration: (1 / 3.0), name: "G4"},
		{duration: 2, name: "D4"},
	}
	voice2 := []note{
		// ---Intro---
		{duration: (1 / 3.0), name: "G3"},
		{duration: (1 / 3.0) / 2, name: "B3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "B3"},

		{duration: (1 / 3.0), name: "D3"},
		{duration: (1 / 3.0) / 2, name: "B3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "B3"},

		{duration: (1 / 3.0), name: "G3"},
		{duration: (1 / 3.0) / 2, name: "B3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "B3"},

		{duration: (1 / 3.0), name: "D3"},
		{duration: (1 / 3.0), name: "E3"},
		{duration: (1 / 3.0), name: "F#3"},

		// ---Actual song---
		{duration: (1 / 3.0), name: "G3"},
		{duration: (1 / 3.0) / 2, name: "B3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "B3"},

		{duration: (1 / 3.0), name: "G3"},
		{duration: (1 / 3.0) / 2, name: "B3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "B3"},

		{duration: (1 / 3.0), name: "G3"},
		{duration: (1 / 3.0) / 2, name: "B3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "B3"},

		{duration: (1 / 3.0), name: "G3"},
		{duration: (1 / 3.0) / 2, name: "B3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "B3"},

		{duration: (1 / 3.0), name: "C3"},
		{duration: (1 / 3.0) / 2, name: "E3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "E3"},

		{duration: (1 / 3.0), name: "C3"},
		{duration: (1 / 3.0) / 2, name: "E3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "E3"},

		{duration: (1 / 3.0), name: "G3"},
		{duration: (1 / 3.0) / 2, name: "B3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "B3"},

		{duration: (1 / 3.0), name: "G3"},
		{duration: (1 / 3.0) / 2, name: "B3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "B3"},
	}
	voice3 := []note{
		// ---Intro---
		{duration: (1 / 3.0), name: "REST"},
		{duration: (1 / 3.0) / 2, name: "D4"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "D4"},

		{duration: (1 / 3.0), name: "REST"},
		{duration: (1 / 3.0) / 2, name: "D4"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "D4"},

		{duration: (1 / 3.0), name: "REST"},
		{duration: (1 / 3.0) / 2, name: "D4"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "D4"},

		{duration: 1, name: "REST"},

		// ---Actual song---
		{duration: (1 / 3.0), name: "REST"},
		{duration: (1 / 3.0) / 2, name: "D4"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "D4"},

		{duration: (1 / 3.0), name: "REST"},
		{duration: (1 / 3.0) / 2, name: "D4"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "D4"},

		{duration: (1 / 3.0), name: "REST"},
		{duration: (1 / 3.0) / 2, name: "D4"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "D4"},

		{duration: (1 / 3.0), name: "REST"},
		{duration: (1 / 3.0) / 2, name: "D4"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "D4"},

		{duration: (1 / 3.0), name: "REST"},
		{duration: (1 / 3.0) / 2, name: "G3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "G3"},

		{duration: (1 / 3.0), name: "REST"},
		{duration: (1 / 3.0) / 2, name: "G3"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "G3"},

		{duration: (1 / 3.0), name: "REST"},
		{duration: (1 / 3.0) / 2, name: "D4"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "D4"},

		{duration: (1 / 3.0), name: "REST"},
		{duration: (1 / 3.0) / 2, name: "D4"},
		{duration: (1 / 3.0) / 2, name: "REST"},
		{duration: (1 / 3.0), name: "D4"},
	}
	samples1 := getSamples(voice1, sampleRate)
	samples2 := getSamples(voice2, sampleRate)
	samples3 := getSamples(voice3, sampleRate)
	var samples []int16
	// This is a bad idea if the slices are of different length
	for i := range samples1 {
		samples = append(samples, samples1[i]+samples2[i]+samples3[i])
	}
	return samples
}
