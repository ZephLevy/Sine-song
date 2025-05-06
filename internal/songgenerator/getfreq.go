package songgenerator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var noteToSemitone = map[string]int{
	"C":  0,
	"C#": 1, "Db": 1,
	"D":  2,
	"D#": 3, "Eb": 3,
	"E":  4,
	"F":  5,
	"F#": 6, "Gb": 6,
	"G":  7,
	"G#": 8, "Ab": 8,
	"A":  9,
	"A#": 10, "Bb": 10,
	"B": 11,
}

func getFrequency(note string) (float64, error) {
	note = strings.TrimSpace(strings.ToUpper(note))

	var name string
	var octaveStr string

	// Handle sharp/flat notes
	if len(note) >= 3 && (note[1] == '#' || note[1] == 'B') {
		name = note[:2]
		octaveStr = note[2:]
	} else if len(note) >= 2 {
		name = note[:1]
		octaveStr = note[1:]
	} else {
		return 0, fmt.Errorf("invalid note format: %s", note)
	}

	semitone, ok := noteToSemitone[name]
	if !ok {
		return 0, fmt.Errorf("unknown note name: %s", name)
	}

	octave, err := strconv.Atoi(octaveStr)
	if err != nil {
		return 0, fmt.Errorf("invalid octave: %s", octaveStr)
	}

	// Midi to frequency algorithm
	midi := (octave+1)*12 + semitone
	freq := 440.0 * math.Pow(2, float64(midi-69)/12.0)
	return freq, nil
}
