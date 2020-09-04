package packagename

import "strings"

var badNames = map[string]interface{}{
	"util":      nil,
	"utils":     nil,
	"utility":   nil,
	"utilities": nil,
	"base":      nil,
	"bases":     nil,
	"common":    nil,
	"commons":   nil,
	"helper":    nil,
	"helpers":   nil,
}

// Ok returns a boolean that is the supplied name is good
// for a package or not
func Ok(name string) (bool, error) {
	_, ok := badNames[strings.ToLower(name)]
	return !ok, nil
}
