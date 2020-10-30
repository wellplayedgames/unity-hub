package editor

import (
	"strings"
)

const (
	versionSplitChars = ".abf"
)

// CompareVersions compares two Unity editor versions.
func CompareVersions(a, b string) int {
	for {
		idxA := strings.IndexAny(a, versionSplitChars)
		idxB := strings.IndexAny(b, versionSplitChars)

		if (idxA > 0) != (idxB > 0) {
			if idxA > 0 {
				return 1
			}
			return -1
		}

		if idxA < 0 {
			break
		}

		aPart := a[:idxA]
		bPart := b[:idxB]
		aSep := int(a[idxA])
		bSep := int(b[idxB])
		a = a[idxA+1:]
		b = b[idxB+1:]

		if c := strings.Compare(aPart, bPart); c != 0 {
			return c
		}

		if c := aSep - bSep; c != 0 {
			return c
		}
	}

	return strings.Compare(a, b)
}
