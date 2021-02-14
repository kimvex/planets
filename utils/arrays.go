package utils

//FindArray for find value of day into array
func FindArray(toFind int, PeakDays []int) bool {
	find := false
	for i := 0; i < len(PeakDays); i++ {
		if toFind == PeakDays[i] {
			find = true
		}
	}

	return find
}
