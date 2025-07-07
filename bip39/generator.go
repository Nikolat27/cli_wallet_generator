package bip39

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"log/slog"
	"os"
	"strings"
)

const (
	entropyBytes int    = 16
	whitespace   string = " "
)

type Mnemonic struct {
	words []string
}

func NewMnemonic() (*Mnemonic, error) {
	entropy, err := generateEntropy()
	if err != nil {
		return nil, err
	}

	checksumEntropy := appendChecksum(entropy)

	chunks := splitBits(checksumEntropy)

	wordsIndex := bitsToWordsIndex(chunks)

	words, err := getWords(wordsIndex)
	if err != nil {
		return nil, err
	}

	w := &Mnemonic{
		words: words,
	}

	return w, nil
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

func bitsToWordsIndex(chunks [][]bool) []int {
	var wordsIndex []int
	for _, bitSlice := range chunks {
		index := 0
		for _, b := range bitSlice {
			index = index << 1
			if b {
				index += 1
			}
		}
		wordsIndex = append(wordsIndex, index)
	}
	return wordsIndex
}

func getWords(wordsIndex []int) ([]string, error) {
	wordsList, err := readWordsFile()
	if err != nil {
		return nil, err
	}

	chosenWords := make([]string, 0, len(wordsIndex))

	for _, index := range wordsIndex {
		word := string(wordsList[index])
		chosenWords = append(chosenWords, word)
	}

	return chosenWords, nil
}

func readWordsFile() ([][]byte, error) {
	file, err := os.Open("bip39/words.txt")
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			slog.Error("closing file", "error", err)
		}
	}()

	words, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	wordsList := bytes.Split(words, []byte("\n"))
	return wordsList, nil
}

func (m *Mnemonic) String() string {
	return strings.Join(m.words, whitespace)
}
