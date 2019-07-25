package totp

import (
    "testing"
)

func TestGetHOTPToken(t *testing.T) {
    out := getHOTPToken("lhe4kfhfqapxipzmohswb6i5adg2gauh", 0)
    if out != "449770" {
        t.Errorf("getHOTPToken test token expecting 449770, got %s",out)
    }
}
