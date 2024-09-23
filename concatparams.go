package piscine

func ConcatParams(args []string) string {
	arr := ""
	for i := 0; i < len(args); i++ {
		arr += args[i]
		if i != len(args)-1 {
			arr += "\n"
		}
	}
	return arr
}
