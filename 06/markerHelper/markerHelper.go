package markerHelper

func IsMarker(input string) bool {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}
	return true
}

func GetFirstMarker(length int, message string) int {
	firstMarker := length
	marker := message[0:length]
	for firstMarker < len(message) {
		if IsMarker(marker) {
			break
		}
		marker = marker[1:length] + message[firstMarker:firstMarker+1]
		firstMarker++
	}
	return firstMarker
}
