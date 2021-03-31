package names

import (
	"fmt"
	"testing"
)

func Test_cache(t *testing.T) {
	SS := "OauthIDAPI"
	tmp0 := UnMarshal(SS)
	fmt.Println(tmp0)
	tmp1 := Marshal(tmp0)
	fmt.Println(tmp1)

	if SS != tmp1 {
		fmt.Println("false.")
	}
}
