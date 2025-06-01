package main

import (
	"bufio"
	"encoding/binary"
	"os"

	"github.com/ZephLevy/The-Simplest-Song/internal/songgenerator"
)

const (
	sampleRate     = 44100
	numChannels    = 1 // Mono
	bitsPerSample  = 16
	bytesPerSample = bitsPerSample / 8
	byteRate       = sampleRate * numChannels * bytesPerSample
	blockAlign     = numChannels * bytesPerSample
)

type countingBuffer struct {
	totalSize int
	writer    *bufio.Writer
}

func (buffer *countingBuffer) Write(samples []byte) (int, error) {
	n, err := buffer.writer.Write(samples)
	buffer.totalSize += n
	return n, err
}

func main() {
	file, err := os.Create("song.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// This reserves enough space for the headers
	file.Write(make([]byte, 44))

	fileBuffer := bufio.NewWriter(file)
	writer := &countingBuffer{writer: fileBuffer}

	for range 2 * sampleRate {
		binary.Write(writer, binary.LittleEndian, int16(0))
	}

	samples := songgenerator.GetSong(sampleRate)
	for _, sample := range samples {
		binary.Write(writer, binary.LittleEndian, sample)
	}
	writer.writer.Flush()

	// Go to start of file and write headers
	file.Seek(0, 0)
	writeHeaders(file, writer.totalSize)
}

func writeHeaders(file *os.File, dataSize int) {
	file.Write([]byte("RIFF"))
	binary.Write(file, binary.LittleEndian, uint32(dataSize+36))
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
