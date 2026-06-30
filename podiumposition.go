package piscine

func PodiumPosition(podium [][]string) [][]string {
	length := len(podium) / 2
	for a := 0; a < length; a++ {
		b := len(podium) - 1 - a
		podium[a], podium[b] = podium[b], podium[a]
	}
	return podium
}
