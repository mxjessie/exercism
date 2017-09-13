// Package twofer has a package comment that summarizes what it's about.
package twofer

// ShareWith has a comment documenting it.
func ShareWith(instr string) string {
	if len(instr) > 0 {
		return "One for " + instr + ", one for me."
	} else {
		return "One for you, one for me."
	}
}
