package goparse_test

import (
	"reflect"
	"testing"

	"github.com/afmahmuda/goparse"
)

type testType struct {
	RootKey rootKey `yaml:"root_key" json:"root_key,omitempty"`
}

type rootKey struct {
	ChildKey1 string    `yaml:"child_key_1" json:"child_key_1,omitempty"`
	ChildKey2 childKey2 `yaml:"child_key_2" json:"child_key_2,omitempty"`
}

type childKey2 struct {
	GrandChildKey1 string   `yaml:"grand_child_key_1" json:"grand_child_key_1,omitempty"`
	GrandChildKey2 string   `yaml:"grand_child_key_2" json:"grand_child_key_2,omitempty"`
	GrandChildKey3 []string `yaml:"grand_child_key_3" json:"grand_child_key_3,omitempty"`
}

func TestParseFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    testType
		wantErr bool
	}{
		{"negative_file_not_found_yaml", args{"./true_happiness.json"}, testType{}, true},
		{"negative_file_not_found_json", args{"./true_happiness.yaml"}, testType{}, true},
		{"negative_extension_empty", args{"./true_happiness"}, testType{}, true},
		{"negative_extension_not_supported", args{"./true_happiness.jsong"}, testType{}, true},
		{"negative_bad_yaml", args{"./testfiles/example_bad.yaml"}, testType{}, true},
		{"negative_bad_json", args{"./testfiles/example_bad.json"}, testType{}, true},
		{"positive_yaml", args{"./testfiles/example_good.yaml"}, testType{RootKey: rootKey{ChildKey1: "child value 1", ChildKey2: childKey2{GrandChildKey1: "grand child value 1", GrandChildKey2: "grand child value 2", GrandChildKey3: []string{"grand child value 3.1", "grand child value 3.2"}}}}, false},
		{"positive_json", args{"./testfiles/example_good.json"}, testType{RootKey: rootKey{ChildKey1: "child value 1", ChildKey2: childKey2{GrandChildKey1: "grand child value 1", GrandChildKey2: "grand child value 2", GrandChildKey3: []string{"grand child value 3.1", "grand child value 3.2"}}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := goparse.ParseFile[testType](tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
