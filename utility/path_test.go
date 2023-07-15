package utility

import (
	"testing"

)

func TestGetProjectAbsRootPath(t *testing.T) {
	tests := []struct {
		name         string
		wantRootPath string
	}{
        {
            name: "Get project root path",
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            gotPath := GetProjectAbsRootPath()
            t.Log(gotPath)
		})
	}
}
