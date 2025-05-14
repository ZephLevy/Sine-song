package songgenerator

type note struct {
	duration     float64
	volume       float64
	name         string
	frequency    float64
	frequencyEnd *float64 // So that it defaults to nil
}

func GetSong(sampleRate float64) []int16 {
	var voice1 []note
	initialVolume := 0.5
	voice1 = append(voice1, note{duration: 5, volume: initialVolume, name: "C4"})
	voice1 = append(voice1, lerpNote(2, initialVolume, "C4", "Db4"))

	samples := getSamples(voice1, float64(sampleRate))
	return samples
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
