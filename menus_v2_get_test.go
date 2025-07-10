package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMenuV2Get(t *testing.T) {
	req := client.NewMenuV2GetRequest()
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	lines := 0
	for _, line := range splitLines(string(b)) {
		fmt.Println(line)
		lines++
		if lines >= 10 {
			break
		}
	}

}

// splitLines splits a string into lines.
func splitLines(s string) []string {
	var lines []string
	start := 0
	for i, c := range s {
		if c == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}
