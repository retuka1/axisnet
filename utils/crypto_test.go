package utils

import (
	"fmt"
	"testing"
)

const PRIVATE_KEY = "ae432353a4eb4eb9"

type TokenResult struct {
	Token string `json:"token"`
}

func TestDecript(t *testing.T) {

	text := "x_jKZv-CaM7J_SGo-SNQXpTNcpaIRFgDIkm8wvM1lr_NcctXqY6t5wFgKHx5wIkklEGaWVq5lLUc5L4YxQ4x9PLXWzM0mykbQ2-rZORAnvPrGTuRb67KfPeIFofStBkbvc8sso_YttnDlFH8_P6phQ=="
	result := Decrypt(text, PRIVATE_KEY)

	fmt.Println(result)
}
