package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	const apiKey = "api_key_123"

	noPrefixHeader := http.Header{}
	noPrefixHeader.Set("Authorization", apiKey)

	badPrefixHeader := http.Header{}
	badPrefixHeader.Set("Authorization", "Bearer "+apiKey)

	goodHeader := http.Header{}
	goodHeader.Set("Authorization", "ApiKey "+apiKey)

	tests := []struct {
		name    string
		input   http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "no auth header",
			input:   http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "no prefix",
			input:   noPrefixHeader,
			want:    "",
			wantErr: true,
		},
		{
			name:    "bad prefix",
			input:   badPrefixHeader,
			want:    "",
			wantErr: true,
		},
		{
			name:    "success",
			input:   goodHeader,
			want:    apiKey,
			wantErr: false,
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if tc.wantErr && err == nil {
			t.Fatalf("%s: expected error, got: %v", tc.name, err)
		}

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
