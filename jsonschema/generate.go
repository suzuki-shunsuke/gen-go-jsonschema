package jsonschema

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/invopop/jsonschema"
)

func Write(src any, path string) error {
	s := jsonschema.Reflect(src)
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal schema as JSON: %w", err)
	}
	if err := os.WriteFile(path, []byte(strings.ReplaceAll(string(b), "http://json-schema.org", "https://json-schema.org")+"\n"), 0o644); err != nil { //nolint:gosec,mnd
		return fmt.Errorf("write JSON Schema to %s: %w", path, err)
	}
	return nil
}
