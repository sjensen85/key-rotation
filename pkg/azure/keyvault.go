package keyvault

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func PutSecret() bool {
	return true
}

func GetSecret(secretName string) string {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		// TODO: handle error
	}

	client, err := azsecrets.NewClient("https://fb-dataplatform-kv-dev.vault.azure.net", cred, nil)
	if err != nil {
		// TODO: handle error
	}

	version := ""
	resp, err := client.GetSecret(context.TODO(), secretName, version, nil)
	if err != nil {
		// TODO: handle error
	}

	fmt.Printf("Secret Name: %s\tSecret Value: %s", resp.ID.Name(), *resp.Value)
	return *resp.Value
}