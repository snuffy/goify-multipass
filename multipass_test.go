package multipass_test

import (
	"fmt"
	"testing"

	multipass "github.com/claimh-solais/goify-multipass"
)

const secret = "e2c1dc490fd04ad8bbcc426316196dba"

func TestTokenGenerate(t *testing.T) {
	var m = multipass.New(secret)
	customerInfo := new(multipass.CustomerInfo{
		Email: "xmhscratch@gmail.com",
	})
	token, err := m.Encode(customerInfo)
	if err != nil {
		t.Error("token generate failed: %v", err)
		return
	}
	fmt.Println(token)
}
