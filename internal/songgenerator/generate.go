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
	startingFreq, _ := getFrequency("C4")
	endingFreq, _ := getFrequency("Db4")
	voice1 = append(voice1,
		note{
			duration:     3,
			volume:       initialVolume,
			frequency:    startingFreq,
			frequencyEnd: &endingFreq,
		},
	)

	samples := getSamples(voice1, float64(sampleRate))
	return samples
}
