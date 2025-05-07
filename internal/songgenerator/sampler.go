package songgenerator

import (
	"log"
	"math"
)

func getSamples(notes []note, sampleRate float64) []int16 {
	var frequencies []float64
	for _, note := range notes {
		frequency, err := getFrequency(note.name)
		if err != nil {
			// Assume a typo or smth equivalent
			// Still, I'm printing to let the user know
			log.Println("Error getting frequency:", err)
			continue
		}
		for i := range int(note.duration * sampleRate) {
			_ = i
			frequencies = append(frequencies, frequency)
		}
	}

	var samples []int16
	for i, frequency := range frequencies {
		// HACK - 2 is the number of voices
		// It probably shouldn't be hardcoded
		sample := int16(math.Sin(2.0*math.Pi*frequency*float64(i)/sampleRate) * (32767 / 2))

		samples = append(samples, sample)
	}
	return samples
}
