package gocom

import "fmt"

type Version struct {
	major int
	minor int
	patch string
}

func FormatVersion(major, minor int, patch string) string {
	return fmt.Sprintf("v%d.%d.%s", major, minor, patch)
}

func ParseVersion(version string) *Version {
	return nil
}

func IsSupported() bool {
	return false
}