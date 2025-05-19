package songgenerator

import "fmt"

type note struct {
	duration     float64
	volume       float64
	name         string
	frequency    float64
	frequencyEnd *float64 // So that it defaults to nil
}

func GetSong(sampleRate float64) []int16 {
	var samples []int16
	samples = append(samples, chordProgression(sampleRate, 4, 4)...)
	background := chordProgression(sampleRate, 5, 3)
	mainLoop := mainLoops(sampleRate)
	for i := range background {
		if i >= len(mainLoop) {
			samples = append(samples, background[i])
			continue
		}
		samples = append(samples, mainLoop[i]+background[i])
	}
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
func chordProgression(sampleRate float64, voiceCount float64, octave int) []int16 {
	initialVolume := 0.5

	n := func(name string) string {
		return fmt.Sprintf("%s%d", name, octave)
	}

	voice1 := []note{
		{duration: 4, volume: initialVolume, name: n("C")},
		lerpNote(3, initialVolume, n("C"), n("Db")),
		{duration: 4, volume: initialVolume, name: n("Db")},
		lerpNote(3, initialVolume, n("Db"), n("D")),
		{duration: 4, volume: initialVolume, name: n("D")},
		lerpNote(3, initialVolume, n("D"), n("C#")),
		{duration: 4, volume: initialVolume, name: n("C#")},
	}

	voice2 := []note{
		{duration: 4, volume: initialVolume, name: n("E")},
		lerpNote(3, initialVolume, n("E"), n("F")),
		{duration: 4, volume: initialVolume, name: n("F")},
		lerpNote(3, initialVolume, n("F"), n("F")),
		{duration: 4, volume: initialVolume, name: n("F")},
		lerpNote(3, initialVolume, n("F"), n("F#")),
		{duration: 4, volume: initialVolume, name: n("F#")},
	}

	voice3 := []note{
		{duration: 4, volume: initialVolume, name: n("G")},
		lerpNote(3, initialVolume, n("G"), n("Gb")),
		{duration: 4, volume: initialVolume, name: n("Gb")},
		lerpNote(3, initialVolume, n("Gb"), n("F")),
		{duration: 4, volume: initialVolume, name: n("F")},
		lerpNote(3, initialVolume, n("F"), n("F#")),
		{duration: 4, volume: initialVolume, name: n("F#")},
	}

	voice4 := []note{
		{duration: 4, volume: initialVolume, name: n("B")},
		lerpNote(3, initialVolume, n("B"), n("Bb")),
		{duration: 4, volume: initialVolume, name: n("Bb")},
		lerpNote(3, initialVolume, n("Bb"), n("B")),
		{duration: 4, volume: initialVolume, name: n("B")},
		lerpNote(3, initialVolume, n("B"), n("A#")),
		{duration: 4, volume: initialVolume, name: n("A#")},
	}
	// Intro
	if octave == 4 {
		voice1 = append(voice1, lerpNote(1, initialVolume, "C#4", "C#3"))
		voice2 = append(voice2, lerpNote(1, initialVolume, "F#4", "F#3"))
		voice3 = append(voice3, lerpNote(1, initialVolume, "F#4", "F#3"))
		voice4 = append(voice4, lerpNote(1, initialVolume, "A#4", "A#3"))
	}

	samples1 := getSamples(voice1, sampleRate, voiceCount)
	samples2 := getSamples(voice2, sampleRate, voiceCount)
	samples3 := getSamples(voice3, sampleRate, voiceCount)
	samples4 := getSamples(voice4, sampleRate, voiceCount)

	var samples []int16
	for i := range samples1 {
		samples = append(samples, samples1[i]+samples2[i]+samples3[i]+samples4[i])
	}
	return samples
}
func mainLoops(sampleRate float64) []int16 {
	volume := 0.5
	var notes []note
	noteNames := []string{"C4", "E4", "G4", "B4", "C5", "B4", "G4", "E4"}
	for i := range 4 * len(noteNames) {
		notes = append(notes, note{
			volume:   volume,
			duration: 1.0 / 8.0,
			name:     noteNames[i%len(noteNames)],
		})
	}
	samples := getSamples(notes, sampleRate, 5)
	return samples
}
