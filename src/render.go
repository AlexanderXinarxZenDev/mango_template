// mango/render.go
package mango

import (
	"fmt"
	"os"
	"strings"
)

// Render reads an HTML file, replaces {{ key }} placeholders with values from data,
// and returns the rendered string.
func Render(filePath string, data map[string]any) (string, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	template := string(b)
	return parse(template, data), nil
}

// parse replaces all {{ key }} occurrences in the template using the data map.
// Undefined keys become empty strings.
func parse(template string, data map[string]any) string {
	var buf strings.Builder
	i := 0
	n := len(template)

	for i < n {
		// Look for opening "{{"
		if i+1 < n && template[i] == '{' && template[i+1] == '{' {
			start := i
			i += 2

			// Skip whitespace after {{
			for i < n && template[i] == ' ' {
				i++
			}

			// Extract key until }}
			keyStart := i
			for i < n && !(template[i] == '}' && i+1 < n && template[i+1] == '}') {
				i++
			}
			key := strings.TrimSpace(template[keyStart:i])

			// Consume the closing "}}"
			if i+1 < n && template[i] == '}' && template[i+1] == '}' {
				i += 2
			} else {
				// Malformed: copy original {{...}}
				buf.WriteString(template[start:i])
				continue
			}

			// Lookup and write value
			val := ""
			if v, ok := data[key]; ok {
				val = fmt.Sprint(v)
			}
			buf.WriteString(val)
		} else {
			buf.WriteByte(template[i])
			i++
		}
	}
	return buf.String()
}