package sf

import (
	"fmt"
	"testing"
)

func TestGetID(t *testing.T) {
	fmt.Println(GetBase64Id())
	fmt.Println(GetInt64Id())
	fmt.Println(GetStringId())
}
