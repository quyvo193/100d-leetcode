// medium problem
package backtracking

func GenerateParenthesis(n int) []string {
	res := []string{}

	var dfs func(path string, openN, closeN int)
	dfs = func(path string, openN, closeN int) {
		if len(path) == n*2 {
			res = append(res, path)
			return
		}

		if openN < n {
			dfs(path+"(", openN+1, closeN)
		}

		if closeN < openN {
			dfs(path+")", openN, closeN+1)
		}
	}

	dfs("", 0, 0)

	return res
}
