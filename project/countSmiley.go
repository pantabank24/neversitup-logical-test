package project

func CountSmileys(arr []string) int {
	eyes := map[string]bool{":": true, ";": true}
	nose := map[string]bool{"-": true, "~": true}
	mouth := map[string]bool{")": true, "D": true}

	count := 0

	for _, face := range arr {
		if eyes[face[0:1]] && mouth[face[len(face)-1:]] {
			if len(face) == 3 && nose[face[1:2]] {
				count++
			} else if len(face) == 2 {
				count++
			}
		}
	}

	return count
}
