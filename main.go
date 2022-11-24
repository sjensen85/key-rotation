package main

import (
	"fmt"

	keyvault "github.com/sjensen85/key-rotation/pkg/azure"
)

func main() {
	secret := keyvault.GetSecret("adls-access-key")
	fmt.Printf("Secret value %s", secret)
}
