/*
Package twofer ...
*/
package twofer

// ShareWith returns the string ...
func ShareWith(name string) string {
	ret := "One for you, one for me."
	if name != "" {
		ret = "One for " + name + ", one for me."
	}
	return ret
}
