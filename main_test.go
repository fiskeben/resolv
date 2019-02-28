package resolv

import (
	"io"
	"os"
	"testing"

	"github.com/go-test/deep"
)

func Test_parse(t *testing.T) {
	tests := []struct {
		name    string
		arg     io.Reader
		want    Resolver
		wantErr bool
	}{
		{
			name:    "parses a regular resolv.conf",
			arg:     open("testdata/regular.resolv.conf"),
			want:    Resolver{Domains: []string{"my.local"}, Nameservers: []string{"192.168.1.1", "8.8.8.8", "8.8.4.4"}, Search: []string{}, Sortlist: []string{}},
			wantErr: false,
		},
		{
			name:    "parses a tabbed resolv.conf",
			arg:     open("testdata/tabbed.resolv.conf"),
			want:    Resolver{Domains: []string{"my.local"}, Nameservers: []string{"192.168.1.1", "8.8.8.8", "8.8.4.4"}, Search: []string{}, Sortlist: []string{}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		r, err := parse(tt.arg)
		if err != nil && !tt.wantErr {
			t.Errorf("%s got unexpected error: %v", tt.name, err)
			continue
		}
		if err == nil && tt.wantErr {
			t.Errorf("%s expected an error", tt.name)
			continue
		}
		if diff := deep.Equal(r, tt.want); diff != nil {
			t.Errorf("%s diff: %v", tt.name, diff)
		}
	}
}

func open(path string) io.Reader {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}
