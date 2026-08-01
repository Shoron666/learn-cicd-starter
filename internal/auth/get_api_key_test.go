package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "no authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "malformed authorization header",
			headers: http.Header{"Authorization": []string{"Hi"}},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name:    "valid api key",
			headers: http.Header{"Authorization": []string{"ApiKey 123"}},
			want:    "123",
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.headers)
		if !reflect.DeepEqual(err, tc.wantErr) {
			t.Errorf("%s: expected error: %v, got: %v", tc.name, tc.wantErr, err)
		}
		if got != tc.want {
			t.Errorf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
