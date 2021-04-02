package jwt

import (
	"fmt"
	"testing"
)

func TestToke_GetToken(t *testing.T) {
	token := NewToken()
	data, err := token.GetToken(1999)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
