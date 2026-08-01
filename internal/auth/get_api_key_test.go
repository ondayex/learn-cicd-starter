package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    // Define a struct for our test cases
    tests := []struct {
        name          string
        headers       http.Header
        expectedKey   string
        expectError   bool
    }{
        {
            name: "Valid Authorization Header",
            headers: http.Header{
                "Authorization": []string{"ApiKey my-secret-token-123"},
            },
            expectedKey: "my-secret-token-123",
            expectError: false,
        },
        {
            name:          "Missing Authorization Header",
            headers:       http.Header{},
            expectedKey:   "",
            expectError:   true,
        },
        {
            name: "Malformed Header - Missing Prefix",
            headers: http.Header{
                "Authorization": []string{"my-secret-token-123"},
            },
            expectedKey: "",
            expectError: true,
        },
        {
            name: "Malformed Header - Wrong Prefix",
            headers: http.Header{
                "Authorization": []string{"Bearer my-secret-token-123"},
            },
            expectedKey: "",
            expectError: true,
        },
        {
            name: "Malformed Header - Only Prefix Provided",
            headers: http.Header{
                "Authorization": []string{"ApiKey"},
            },
            expectedKey: "",
            expectError: true,
        },
    }

    // Iterate through all test cases
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            actualKey, err := GetAPIKey(tc.headers)

            // Check error expectation
            if (err != nil) != tc.expectError {
                t.Fatalf("expected error = %v, got error = %v", tc.expectError, err)
            }

            // Check returned key expectation
            if actualKey != tc.expectedKey {
                t.Errorf("expected key = %q, got key = %q", tc.expectedKey, actualKey)
            }
        })
    }
}
