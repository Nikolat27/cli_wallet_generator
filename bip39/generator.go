package bip39

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"log/slog"
	"os"
)

const (
	entropyBytes int = 16
)

func GenerateSeeds() ([]string, error) {
	entropy, err := generateEntropy()
	if err != nil {
		return nil, err
	}

	checksumEntropy := appendChecksum(entropy)

	chunks := splitBits(checksumEntropy)

	ints := bitsToInts(chunks)

	words, err := getWords(ints)
	if err != nil {
		return nil, err
	}

	return words, nil
}

func generateEntropy() ([]byte, error) {
	entropy := make([]byte, entropyBytes)
	_, err := rand.Read(entropy)
	if err != nil {
		return nil, err
	}

	return entropy, nil
}

func appendChecksum(entropy []byte) []bool {
	hash := sha256.Sum256(entropy)
	var bits []bool

	for _, b := range entropy {
		for i := 7; i >= 0; i-- {
			bits = append(bits, (b&(1<<i)) != 0)
		}
	}
	// Append only 4 bits checksum
	for i := 7; i >= 4; i-- {
		bits = append(bits, (hash[0]&(1<<i)) != 0)
	}
	return bits
}

func splitBits(bits []bool) [][]bool {
	var chunks [][]bool
	for i := 0; i < len(bits); i += 11 {
		chunks = append(chunks, bits[i:i+11])
	}
	return chunks
}

func bitsToInts(chunks [][]bool) []int {
	var ints []int
	for _, bitSlice := range chunks {
		val := 0
		for _, b := range bitSlice {
			val = val << 1
			if b {
				val += 1
			}
		}
		ints = append(ints, val)
	}
	return ints
}

func getWords(ints []int) ([]string, error) {
	separatedWords, err := readWordsFile()
	if err != nil {
		return nil, err
	}

	chosenWords := make([]string, 0, len(ints))

	for _, val := range ints {
		chosenWords = append(chosenWords, string(separatedWords[val]))
	}

	return chosenWords, nil
}

func readWordsFile() ([][]byte, error) {
	file, err := os.Open("bip39/words.txt")
	if err != nil {
		return nil, err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			slog.Error("closing file", "error", err)
		}
	}()

	words, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	separatedWords := bytes.Split(words, []byte("\n"))
	return separatedWords, nil
}
