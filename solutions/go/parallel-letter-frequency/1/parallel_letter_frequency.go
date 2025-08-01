package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

func run(text string, c chan FreqMap) {
    fm := Frequency(text)
    c <- fm
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {

    c := make(chan FreqMap, len(texts))
	for _, text := range texts {
        go run(text, c)
    }

    frequencies := FreqMap{}
    for range texts {
        for r, freq := range <-c {
            frequencies[r] += freq
        }
    }    
    return frequencies
}
