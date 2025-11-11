package i18n

import (
	_ "embed"
	"encoding/json"
)

type Dict map[string]string

//go:embed en.json
var enRaw []byte

//go:embed de.json
var deRaw []byte

var cache = map[string]Dict{}

func Translate(lang string) Dict {
	switch lang {
	case "de":
		return load("de", deRaw)
	default:
		return load("en", enRaw)
	}
}

func load(code string, raw []byte) Dict {
	if d, ok := cache[code]; ok {
		return d
	}
	var d Dict
	_ = json.Unmarshal(raw, &d)
	cache[code] = d
	return d
}
