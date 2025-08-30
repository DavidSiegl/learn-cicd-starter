package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectedErr string
	}{
		{
			name:        "no authorization header",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: "no authorization header included",
		},
		{
			name:        "empty authorization header",
			headers:     http.Header{"Authorization": {""}},
			expectedKey: "",
			expectedErr: "no authorization header included",
		},
		{
			name:        "malformed header - wrong prefix",
			headers:     http.Header{"Authorization": {"Bearer abc123"}},
			expectedKey: "",
			expectedErr: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if tt.expectedErr != "" {
				if err == nil {
					t.Errorf("expected error '%s', got nil", tt.expectedErr)
					return
				}
				if err.Error() != tt.expectedErr {
					t.Errorf("expected error '%s', got '%s'", tt.expectedErr, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got '%s'", err.Error())
					return
				}
			}

			if key != tt.expectedKey {
				t.Errorf("expected key '%s', got '%s'", tt.expectedKey, key)
			}
		})
	}
}

func TestGetAPIKey_ErrorTypes(t *testing.T) {
	// Test that the specific error variable is returned
	headers := http.Header{}
	_, err := GetAPIKey(headers)

	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}
