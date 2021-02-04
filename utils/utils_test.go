package utils

import (
	"fmt"
	"testing"
)

func TestFilterEmoji(t *testing.T) {
	fmt.Println(FilterEmoji("Thats a nice joke 😆😆😆 😛"))
}
