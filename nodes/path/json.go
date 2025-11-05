package path

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const defaultSeparator = "/"

type Path struct {
	raw       string
	separator string
	segments  []string
	values    []interface{}
	recursive bool
}

func (p *Path) String() string {
	return p.raw
}

func (p *Path) Segments() []string {
	return p.segments
}

func (p *Path) Values() []interface{} {
	return p.values
}

func (p *Path) Depth() int {
	if len(p.Segments()) == 0 {
		return 0
	}
	return len(p.Segments()) - 1
}

func (p *Path) Size() int {
	return len(p.Segments())
}

func (p *Path) At(index int) string {
	if index >= p.Size() || index < 0 {
		return ""
	}
	return p.Segments()[index]
}

func (p *Path) IsRecursive() bool {
	return p.recursive
}

func ParsePath(path string, sep string) (*Path, error) {
	p := &Path{}
	p.raw = path
	if p.raw == "" {
		return nil, fmt.Errorf("path is empty")
	}
	p.separator = sep
	if p.separator == "" { // Set default separator
		p.separator = defaultSeparator
	}

	if !strings.HasPrefix(p.raw, fmt.Sprintf("%[1]s%[1]s", p.separator)) {
		p.recursive = true
	}
	p.segments = strings.Split(p.raw, p.separator)
	return p, nil
}

type Json struct {
	data []byte
	dec  *json.Decoder
}

func (j *Json) FindValues(paths ...*Path) error {
	depth := 0

	maxDepth := 0
	for _, p := range paths {
		if maxDepth < p.Depth() {
			maxDepth = p.Depth()
		}
	}
	cursor := make([]string, maxDepth+1)

	for {
		t, err := j.dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch v := t.(type) {
		case json.Delim:
			if v == '{' || (depth == 0 && v == '[') {
				depth++
			}
			if v == '}' {
				depth--
			}
			continue
		case string:
			if depth > maxDepth {
				continue
			}
			cursor[depth] = v
			for _, p := range paths {
				if depth == p.Depth() &&
					v == p.At(depth) &&
					strings.Join(p.Segments()[:depth], p.separator) == strings.Join(cursor[:depth], p.separator) {
					t, _ = j.dec.Token() // Value
					p.values = append(p.values, t)
				}
			}
			//dec.More()
		}
	}

	return nil
}

func JsonParse(b []byte) (*Json, error) {
	if !json.Valid(b) {
		return nil, fmt.Errorf("invallid json")
	}
	j := &Json{}
	j.data = b
	j.dec = json.NewDecoder(bytes.NewReader(j.data))
	return j, nil
}
