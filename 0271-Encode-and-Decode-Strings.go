func encode(strs []string) string {
	var encoded strings.Builder

	for _, str := range strs {
		length := strconv.Itoa(len(str))
		encoded.WriteString(length)
		encoded.WriteString("#")
		encoded.WriteString(str)
	}

	return encoded.String()
}

func decode(s string) []string {
	var decoded []string
	i := 0

	for i < len(s) {
		j := strings.Index(s[i:], "#")
		if j == -1 {
			break
		}
		j += i 
		length, err := strconv.Atoi(s[i:j])
		if err != nil {
			break
		}
		start := j + 1
		end := start + length
		decoded = append(decoded, s[start:end])
		i = end
	}

	return decoded
}
