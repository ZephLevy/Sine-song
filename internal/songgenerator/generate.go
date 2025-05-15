package songgenerator

type note struct {
	duration     float64
	volume       float64
	name         string
	frequency    float64
	frequencyEnd *float64 // So that it defaults to nil
}

func GetSong(sampleRate float64) []int16 {
	return intro(sampleRate)
}

func lerpNote(lerpDuration float64, volume float64, noteStart string, noteEnd string) note {
	startingFreq, err := getFrequency(noteStart)
	endFreq, err2 := getFrequency(noteEnd)
	if err != nil || err2 != nil {
		// Better to just panic and let the user know
		panic("Error getting frequencies. Check note naming.")
	}
	return note{
		duration:     lerpDuration,
		volume:       volume,
		frequency:    startingFreq,
		frequencyEnd: &endFreq,
	}

}
func intro(sampleRate float64) []int16 {
	initialVolume := 0.5
	voice1 := []note{
		{duration: 4, volume: initialVolume, name: "C4"},
		lerpNote(3, initialVolume, "C4", "Db4"),
		{duration: 4, volume: initialVolume, name: "Db4"},
		lerpNote(3, initialVolume, "Db4", "D4"),
		{duration: 4, volume: initialVolume, name: "D4"},
		lerpNote(3, initialVolume, "D4", "C#4"),
		{duration: 4, volume: initialVolume, name: "C#4"},
	}

	voice2 := []note{
		{duration: 4, volume: initialVolume, name: "E4"},
		lerpNote(3, initialVolume, "E4", "F4"),
		{duration: 4, volume: initialVolume, name: "F4"},
		lerpNote(3, initialVolume, "F4", "F4"),
		{duration: 4, volume: initialVolume, name: "F4"},
		lerpNote(3, initialVolume, "F4", "F#4"),
		{duration: 4, volume: initialVolume, name: "F#4"},
	}

	voice3 := []note{
		{duration: 4, volume: initialVolume, name: "G4"},
		lerpNote(3, initialVolume, "G4", "Gb4"),
		{duration: 4, volume: initialVolume, name: "Gb4"},
		lerpNote(3, initialVolume, "Gb4", "F4"),
		{duration: 4, volume: initialVolume, name: "F4"},
		lerpNote(3, initialVolume, "F4", "F#4"),
		{duration: 4, volume: initialVolume, name: "F#4"},
	}

	voice4 := []note{
		{duration: 4, volume: initialVolume, name: "B4"},
		lerpNote(3, initialVolume, "B4", "Bb4"),
		{duration: 4, volume: initialVolume, name: "Bb4"},
		lerpNote(3, initialVolume, "Bb4", "B4"),
		{duration: 4, volume: initialVolume, name: "B4"},
		lerpNote(3, initialVolume, "B4", "A#4"),
		{duration: 4, volume: initialVolume, name: "A#4"},
	}

	samples1 := getSamples(voice1, sampleRate)
	samples2 := getSamples(voice2, sampleRate)
	samples3 := getSamples(voice3, sampleRate)
	samples4 := getSamples(voice4, sampleRate)

	var samples []int16
	for i := range samples1 {
		// This is a bad idea if the voices are of different length
		samples = append(samples, samples1[i]+samples2[i]+samples3[i]+samples4[i])
	}
	return samples
}
