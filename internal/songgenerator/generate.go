package songgenerator

type note struct {
	duration float64
	name     string
}

func GetSong(sampleRateInt int) []int16 {
	sampleRate := float64(sampleRateInt)
	voice1 := []note{
		{duration: 1, name: "D5"},
		{duration: 1, name: "B4"},
		{duration: 1, name: "G4"},
		{duration: 1, name: "D4"},
		{duration: 1 / 3.0, name: "E4"},
		{duration: 1 / 3.0, name: "F#4"},
		{duration: 1 / 3.0, name: "G4"},
		{duration: 2 / 3.0, name: "E4"},
		{duration: 1 / 3.0, name: "G4"},
		{duration: 2, name: "D4"},
	}
	voice2 := []note{
		{duration: 1 / 3.0, name: "G3"},
		{duration: 1 / 3.0, name: "B3"},
		{duration: 1 / 3.0, name: "B3"},
		{duration: 1 / 3.0, name: "G3"},
		{duration: 1 / 3.0, name: "B3"},
		{duration: 1 / 3.0, name: "B3"},
		{duration: 1 / 3.0, name: "G3"},
		{duration: 1 / 3.0, name: "B3"},
		{duration: 1 / 3.0, name: "B3"},
		{duration: 1 / 3.0, name: "G3"},
		{duration: 1 / 3.0, name: "B3"},
		{duration: 1 / 3.0, name: "B3"},
		{duration: 1 / 3.0, name: "C3"},
		{duration: 1 / 3.0, name: "E3"},
		{duration: 1 / 3.0, name: "E3"},
		{duration: 1 / 3.0, name: "C3"},
		{duration: 1 / 3.0, name: "E3"},
		{duration: 1 / 3.0, name: "E3"},
		{duration: 1 / 3.0, name: "G3"},
		{duration: 1 / 3.0, name: "B3"},
		{duration: 1 / 3.0, name: "B3"},
		{duration: 1 / 3.0, name: "G3"},
		{duration: 1 / 3.0, name: "B3"},
		{duration: 1 / 3.0, name: "B3"},
	}
	samples1 := getSamples(voice1, sampleRate)
	samples2 := getSamples(voice2, sampleRate)
	var samples []int16
	for i := range samples1 {
		samples = append(samples, samples1[i]+samples2[i])
	}
	return samples
}
