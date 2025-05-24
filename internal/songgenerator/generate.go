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
	var frequencies []int
	func() {
		freqs := chordProgression(sampleRate, 4)
		frequencies = append(frequencies, freqs...)

	}()
	func() {
		background := chordProgression(sampleRate, 3)
		mainLoop := mainLoops(sampleRate)
		for i := range mainLoop {
			frequencies = append(frequencies, background[i]+mainLoop[i])
		}

	}()
	func() {
		background := chordProgression(sampleRate, 3)
		mainLoop := mainLoops(sampleRate)
		bass := bass(sampleRate)
		for i := range mainLoop {
			frequencies = append(frequencies, background[i]+mainLoop[i]+bass[i])
		}
	}()

	func() {
		frequencies = append(frequencies, endingProgression(sampleRate, 4)...)
	}()

	endingSilence := make([]int, int(sampleRate)*4)
	frequencies = append(frequencies, endingSilence...)

	return normalize([][]int{frequencies})
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
func chordProgression(sampleRate float64, octave int) []int {
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

	samples1 := getSamples(voice1, sampleRate, sineWave)
	samples2 := getSamples(voice2, sampleRate, sineWave)
	samples3 := getSamples(voice3, sampleRate, sineWave)
	samples4 := getSamples(voice4, sampleRate, sineWave)

	var samples []int
	for i := range samples1 {
		sample := samples1[i] + samples2[i] + samples3[i] + samples4[i]
		samples = append(samples, sample)
	}
	return samples
}
func mainLoops(sampleRate float64) []int {
	volume := 0.5
	var notes []note

	// C7 arpeggio
	func() {
		noteNames := []string{"C4", "E4", "G4", "B4", "C5", "B4", "G4", "E4"}
		for i := range 4 * len(noteNames) {
			notes = append(notes, note{
				volume:   volume,
				duration: 1.0 / 8.0,
				name:     noteNames[i%len(noteNames)],
			})
		}
		notes = append(notes, lerpNote(3.0, volume, noteNames[0], "D4"))
	}()

	// Db7♭5 arpeggio
	func() {
		noteNames := []string{"Db4", "F4", "Gb4", "Bb4", "Db5", "Bb4", "Gb4", "F4"}
		for i := range 4 * len(noteNames) {
			notes = append(notes, note{
				volume:   volume + 0.15,
				duration: 1.0 / 8.0,
				name:     noteNames[i%len(noteNames)],
			})
		}
		notes = append(notes, lerpNote(3.0, volume, noteNames[0], "D#4"))
	}()

	// Bdim arpeggio
	func() {
		noteNames := []string{"D4", "F4", "B4", "D5", "F5", "D5", "B4", "F4"}
		for i := range 4 * len(noteNames) {
			notes = append(notes, note{
				volume:   volume + 0.3,
				duration: 1.0 / 8.0,
				name:     noteNames[i%len(noteNames)],
			})
		}
		notes = append(notes, lerpNote(3.0, volume, noteNames[0], "C4"))
	}()

	func() {
		noteNames := []string{"C#4", "F#4", "A#4", "C#5", "F#5", "C#5", "A#4", "F#4"}
		for i := range 4 * len(noteNames) {
			notes = append(notes, note{
				volume:   volume + 0.3,
				duration: 1.0 / 8.0,
				name:     noteNames[i%len(noteNames)],
			})
		}
	}()
	samples := getSamples(notes, sampleRate, sineWave)
	return samples
}

func bass(sampleRate float64) []int {
	volume := 0.15
	var notes []note

	roots := []string{"C3", "Db3", "D3", "C#3"}
	altNotes := [][]string{
		{"G2", "C4", "Bb2"},  // for C
		{"Ab2", "Db4", "B2"}, // for Db
		{"A2", "D4", "C3"},   // for D
		{"G#2", "F#4", "B2"}, // for C#
	}

	for i, root := range roots {
		alts := altNotes[i]

		notes = append(notes,
			note{duration: 0.5, volume: volume, name: root},
			note{duration: 0.5, volume: volume * 0.9, name: alts[0]},
			note{duration: 0.5, volume: volume * 0.85, name: alts[1]},
			note{duration: 0.5, volume: volume * 0.8, name: root},
			note{duration: 1.0, volume: volume * 0.75, name: alts[2]},
			note{duration: 1.0, volume: volume, name: root},
		)

		// 3s lerp except final segment
		if i < len(roots)-1 {
			notes = append(notes,
				lerpNote(3.0, volume*0.7, root, roots[i+1]),
			)
		}
	}

	return getSamples(notes, sampleRate, sawtoothWave)
}

func endingProgression(sampleRate float64, octave int) []int {
	initialVolume := 0.5

	n := func(name string) string {
		return fmt.Sprintf("%s%d", name, octave)
	}

	// In the intro, each note lerps to the one closest to it.
	// That's not the case here which creates a cool howling effect
	voice1 := []note{
		lerpNote(3, initialVolume, n("C#"), n("D")),
		{duration: 4, volume: initialVolume, name: n("D")},
		lerpNote(3, initialVolume, n("D"), n("G")),
		{duration: 4, volume: initialVolume, name: n("G")},
		lerpNote(3, initialVolume, n("G"), n("C")),
		{duration: 4, volume: initialVolume, name: n("C")},
	}

	voice2 := []note{
		lerpNote(3, initialVolume, n("F#"), n("F#")),
		{duration: 4, volume: initialVolume, name: n("F#")},
		lerpNote(3, initialVolume, n("F#"), n("B")),
		{duration: 4, volume: initialVolume, name: n("B")},
		lerpNote(3, initialVolume, n("B"), n("E")),
		{duration: 4, volume: initialVolume, name: n("E")},
	}

	voice3 := []note{
		lerpNote(3, initialVolume, n("F#"), n("A")),
		{duration: 4, volume: initialVolume, name: n("A")},
		lerpNote(3, initialVolume, n("A"), n("D")),
		{duration: 4, volume: initialVolume, name: n("D")},
		lerpNote(3, initialVolume, n("D"), n("G")),
		{duration: 4, volume: initialVolume, name: n("G")},
	}

	voice4 := []note{
		lerpNote(3, initialVolume, n("A#"), n("C")),
		{duration: 4, volume: initialVolume, name: n("C")},
		lerpNote(3, initialVolume, n("C"), n("F")),
		{duration: 4, volume: initialVolume, name: n("F")},
		lerpNote(3, initialVolume, n("F"), n("B")),
		{duration: 4, volume: initialVolume, name: n("B")},
	}

	samples1 := getSamples(voice1, sampleRate, sineWave)
	samples2 := getSamples(voice2, sampleRate, sineWave)
	samples3 := getSamples(voice3, sampleRate, sineWave)
	samples4 := getSamples(voice4, sampleRate, sineWave)

	var samples []int
	for i := range samples1 {
		sample := samples1[i] + samples2[i] + samples3[i] + samples4[i]
		samples = append(samples, sample)
	}
	return samples
}
