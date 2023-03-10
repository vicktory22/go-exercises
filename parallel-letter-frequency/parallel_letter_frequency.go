package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}

	for _, r := range s {
		m[r]++
	}

	return m
}

func toChannel(s string, fCh chan FreqMap) {
	fCh <- Frequency(s)
}

func ConcurrentFrequency(l []string) FreqMap {
	freqCh := make(chan FreqMap, len(l))

	totals := FreqMap{}

	for _, s := range l {
		go toChannel(s, freqCh)
	}

	for range l {
		for k, v := range <-freqCh {
			totals[k] += v
		}
	}

	return totals
}
