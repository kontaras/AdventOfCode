package main

func FindPacketStart(signal string) int {
	var i = 3
	var one byte
	var two, three, four = signal[0], signal[1], signal[2]

	for ; i < len(signal); i++ {
		one, two, three, four = two, three, four, signal[i]

		if one != two && one != three && one != four &&
			two != three && two != four && three != four {
			break
		}
	}

	return i + 1
}

func FindMessageStart(signal string) int {
	const buffLen = 14

	var i = 0
	for ; i < len(signal); i++ {
		set := make(map[byte]byte)
		for j := 0; j < buffLen; j++ {
			set[signal[i+j]] = 0
		}

		if len(set) == buffLen {
			break
		}
	}

	return i + buffLen
}
