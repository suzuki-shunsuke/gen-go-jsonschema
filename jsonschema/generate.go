package jsonschema

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/invopop/jsonschema"
	"golang.org/x/mod/modfile"
)

type Options struct {
	ModFile string
}

func Write(src any, path string, opts *Options) error {
	r := &jsonschema.Reflector{}
	if opts != nil && opts.ModFile != "" {
		m, err := os.ReadFile(opts.ModFile)
		if err != nil {
			return fmt.Errorf("read mod file: %w", err)
		}
		if err := r.AddGoComments(modfile.ModulePath(m), "."); err != nil {
			return err
		}
	}
	s := r.Reflect(src)
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal schema as JSON: %w", err)
	}
	if err := os.WriteFile(path, []byte(strings.ReplaceAll(string(b), "http://json-schema.org", "https://json-schema.org")+"\n"), 0o644); err != nil { //nolint:gosec,mnd
		return fmt.Errorf("write JSON Schema to %s: %w", path, err)
	}
	return nil
}
