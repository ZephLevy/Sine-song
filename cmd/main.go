package main

import (
	"encoding/binary"
	"math"
	"os"
)

const (
	sampleRate     = 44100
	duration       = 1.0
	numChannels    = 1 // Stereo
	bitsPerSample  = 16
	bytesPerSample = bitsPerSample / 8
	numSamples     = int(sampleRate * duration)
	byteRate       = sampleRate * numChannels * bytesPerSample
	blockAlign     = numChannels * bytesPerSample
	dataSize       = numSamples * numChannels * bytesPerSample
	chunkSize      = 36 + dataSize
)

func main() {
	frequency := 20.0

	file, err := os.Create("song.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writeHeaders(file)

	// --- Audio Samples ---
	for i := range numSamples {
		sample := int16(math.Sin(2*math.Pi*frequency*float64(i)/sampleRate) * 32767)
		binary.Write(file, binary.LittleEndian, sample)
		frequency = (float64(i) / float64(numSamples)) * (20000.0 / 20.0)
	}
}

func writeHeaders(file *os.File) {
	file.Write([]byte("RIFF"))
	binary.Write(file, binary.LittleEndian, uint32(chunkSize))
	file.Write([]byte("WAVE"))

	// --- fmt subchunk ---
	file.Write([]byte("fmt "))
	binary.Write(file, binary.LittleEndian, uint32(16)) // Subchunk1Size
	binary.Write(file, binary.LittleEndian, uint16(1))  // AudioFormat (1 = PCM)
	binary.Write(file, binary.LittleEndian, uint16(numChannels))
	binary.Write(file, binary.LittleEndian, uint32(sampleRate))
	binary.Write(file, binary.LittleEndian, uint32(byteRate))
	binary.Write(file, binary.LittleEndian, uint16(blockAlign))
	binary.Write(file, binary.LittleEndian, uint16(bitsPerSample))

	// --- data subchunk ---
	file.Write([]byte("data"))
	binary.Write(file, binary.LittleEndian, uint32(dataSize))
}
