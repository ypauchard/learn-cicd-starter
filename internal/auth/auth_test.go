package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		headerKey   string
		headerValue string
		expect      string
		expectedErr string
	}{
		{
			headerKey:   "Authorization",
			expectedErr: "no authorization header included",
		},
		{
			headerKey:   "Authorization",
			headerValue: "Bearer xxxxxxxx",
			expectedErr: "malformed authorization header",
		},
		{
			headerKey:   "Authorization",
			headerValue: "ApiKey xxxxxxxx",
			expect:      "xxxxxxxx",
			expectedErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetApiKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.headerKey, test.headerValue)
			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectedErr) {
					return
				}
				t.Errorf("Unexpected: TestGetApiKey:%v\n", err)
				return
			}
			if output != test.expect {
				t.Errorf("Unexpected: TestGetApiKey:%s\n", output)
				return
			}

		})
	}
}
