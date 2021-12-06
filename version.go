package gocom

/*
import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	major int
	minor int
	patch string
}

func (v *Version) VersionToString() string {
	return fmt.Sprintf("v%d.%d.%s", v.major, v.minor, v.patch)
}

func StringToVersion(v string) (Version, error) {
	version_data := strings.Split(v, ".")
	if len(version_data) != 3 {
		return nil, errors.New("Error")
	}

	major, ma_err := strconv.Atoi(version_data[0])
	minor, mi_err := strconv.Atoi(version_data[1])

	if ma_err || mi_err {
		return nil, errors.New("Error")
	}

	return Version{
		major: major,
		minor: minor,
		patch: version_data[2],
	}, nil
}

func IsSupported(currentVersion Version, supportedVersions []Version) bool {
	return false
}
*/
