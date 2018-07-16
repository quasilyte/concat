// +build !amd64

package concat

// Strings returns x+y concatenation result.
// Works faster than Go "+" operator if neither of strings is empty.
func Strings(x, y string) string {
	length := len(x) + len(y)
	if length == 0 {
		return ""
	}
	b := make([]byte, length)
	copy(b, x)
	copy(b[len(x):], y)
	return goString(&b[0], length)
}
