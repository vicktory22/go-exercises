package hamming

type LengthMismatchError struct{}

func (e *LengthMismatchError) Error() string {
	return "Unable to find distance.  String lengths are different"
}

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, &LengthMismatchError{}
	}

	hammingDistance := 0

	for idx := range a {
		if a[idx] == b[idx] {
			continue
		}

		hammingDistance++
	}

	return hammingDistance, nil
}
