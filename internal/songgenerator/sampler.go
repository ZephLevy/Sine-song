package songgenerator

import (
	"log"
	"math"
)

func getSamples(notes []note, sampleRate float64) []int16 {
	var samples []int16
	var phase float64

	for _, note := range notes {
		frequency, err := getFrequency(note.name)
		if err != nil {
			// Assume a typo or smth equivalent
			// Still, I'm printing to let the user know
			log.Println("Error getting frequency:", err)
			continue
		}

		// What happens when we just use i is that
		// when the note changes, it made a popping sound.
		// This was because i got reset for each note,
		// so a peak in one note could abruptly change to a
		// trough.
		// Here, we use a phase variable to keep track of it between notes

		numSamples := int(note.duration * sampleRate)
		phaseIncrement := 2.0 * math.Pi * frequency / sampleRate

		// This helps fix popping when volume changes
		// It does not work when the volume change is too large.
		// Solution? Don't make large volume shifts.
		multiplier := 1.0
		for range numSamples {
			if phase == 0 {
				multiplier = note.volume
			}
			sample := int16(math.Sin(phase) * (32767 / 3) * multiplier)
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
