package main

import (
	"path/filepath"
	"testing"
)

func Test_parse(t *testing.T) {
	rdir := cdir()
	replaceMap := make(map[string]string, 0)
	replaceMap["https://github.com/knative/docs/blob/master"] = rdir

	pc := &parseContext{replaceMap, make(map[string]int, 0), make(map[string]int, 0)}

	tests := []struct {
		name string
		path []string
	}{
		{name: "parse", path: []string{filepath.Join("testsample", "README.md")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc.parse(tt.path)
		})
	}
}
