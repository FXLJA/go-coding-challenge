package challenge_02

func Grade(score int) string {
	if score < 0 {
		return "Input salah"
	} else {
		if score < 55 {
			return "TIDAK LULUS"
		} else if score < 65 {
			return "D"
		} else if score < 75 {
			return "C"
		} else if score < 85 {
			return "B"
		} else {
			return "A"
		}
	}
}
