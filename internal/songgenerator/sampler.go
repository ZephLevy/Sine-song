package songgenerator

import (
	"log"
	"math"
)

type waveformFunc func(phase float64) float64

func sineWave(phase float64) float64 {
	return math.Sin(phase)
}

func sawtoothWave(phase float64) float64 {
	return 2*(phase/(2*math.Pi)) - 1
}

func getSamples(notes []note, sampleRate float64, wave waveformFunc) []int {
	var samples []int

	// What happens when we just use i and reset it for each note is that
	// when the note changes, it made a popping sound.
	// This was because i got reset for each note,
	// so a peak in one note could abruptly change to a
	// trough.
	// Here, we use a phase variable to keep track of it between notes
	var phase float64
	for _, note := range notes {
		var frequency float64
		if note.name == "" {
			frequency = note.frequency
		} else {
			var err error
			frequency, err = getFrequency(note.name)
			if err != nil {
				// Assume a typo or smth equivalent
				// Still, I'm printing to let the user know
				log.Println("Error getting frequency:", err)
				continue
			}
		}

		numSamples := int(note.duration * sampleRate)

		for i := range numSamples {
			t := float64(i) / float64(numSamples)
			var currentFreq float64

			if note.frequencyEnd != nil {
				currentFreq = (1.0-t)*frequency + t*(*note.frequencyEnd)
			} else {
				currentFreq = frequency
			}

			phaseIncrement := 2.0 * math.Pi * currentFreq / sampleRate
			// sample := int16(math.Sin(phase) * (32767 / voiceCount) * note.volume)
			value := wave(phase)
			sample := int(value * 32767 * note.volume)
			samples = append(samples, sample)
			phase += phaseIncrement

			// I don't know how I managed to have the foresight
			// to think of this, but it seems like a good idea
			if phase > 2*math.Pi {
				phase -= 2 * math.Pi
			}
		}
	}
	return samples
}

func normalize(samples [][]int) []int16 {
	var result []int16
	mixed := make([]int, len(samples[0]))

	for i := range mixed {
		for _, track := range samples {
			mixed[i] += track[i]
		}
	}

	// Find max absolute sample
	var maxAbs int
	for _, s := range mixed {
		if abs := int(math.Abs(float64(s))); abs > maxAbs {
			maxAbs = abs
		}
	}

	// Normalize to int16 range
	for _, s := range mixed {
		normalized := float64(s) / float64(maxAbs) * 32767
		result = append(result, int16(normalized))
	}

	return result
}
