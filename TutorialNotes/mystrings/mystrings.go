package mystrings

/*
Reverse reverses a string from left to right
We need to capitalize the first letter of the function otherise we wont be able to acees this function outside mystrings package
*/

func Reverse(s string) string {
	result := ""

	for _, v := range s {
		result = string(v) + result
	}
	return result
}
