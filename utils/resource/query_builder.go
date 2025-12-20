package resource

func CreateQueryQuestionMarks(args []string) string {
	str := ""
	for i := range args {
		str += "?"
		isLast := i+1 == len(args)
		if !isLast {
			str += ", "
		}
	}
	return str
}
