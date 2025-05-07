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
	samples := getSamples(voice1, sampleRate)
	return samples
}
