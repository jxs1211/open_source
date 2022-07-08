package config

import (
	"reflect"
	"testing"
)

func TestLoadConf(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		// TODO: Add test cases.
		{"base", args{"../../../examples/config.yaml"}, &Config{
			[]Tool{Tool{
				Name:    "githubactions",
				Version: "0.0.1",
				Options: map[string]interface{}{
					"repo": "https://github.com/ironcore864/go-hello-http",
					"language": map[string]interface{}{
						"name":    "go",
						"version": 1.17,
					},
					"branch": "master",
				}}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadConf(tt.args.fname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
