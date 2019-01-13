package vms

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/preview/commerce/mgmt/2015-06-01-preview/commerce"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var (
	ctx = context.Background()
)

// ListAzureGroups Sample func to list azure arm groups
func ListAzureGroups() {
	sid := os.Getenv("AZURE_SUBSCRIPTION_ID")

	// Authenticate with Azure
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatalf(err.Error())
	}

	client := resources.NewGroupsClient(sid)
	client.Authorizer = authorizer

	groups, err := client.List(ctx, "", nil)

	if err != nil {
		log.Printf(err.Error())
	}

	for _, item := range groups.Values() {
		log.Println(*item.Name)
	}
}

// ListRateCard Lists some Azure rate card
func ListRateCard() {
	sid := os.Getenv("AZURE_SUBSCRIPTION_ID")

	client := commerce.NewRateCardClient(sid)
	client.Authorizer = getAuthorizer()
	filter := "OfferDurableId eq 'MS-AZR-0121p' and Currency eq 'AUD' and Locale eq 'en-US' and RegionInfo eq 'AU'"
	result, err := client.Get(ctx, filter)
	if err != nil {
		log.Fatalf(err.Error())
	}

	count := 0
	for _, item := range *result.Meters {
		if *item.MeterCategory == "Virtual Machines" {
			count++
			fmt.Printf("%d,%s,%s,%f,%s\n", count, *item.MeterCategory, *item.MeterName, *item.MeterRates["0"], *item.MeterRegion)
		}
	}

}

// ListAzureVMS Lists Azure VM Types
func ListAzureVMS() {
	sid := os.Getenv("AZURE_SUBSCRIPTION_ID")
	client := compute.NewVirtualMachinesClient(sid)
	client.Authorizer = getAuthorizer()

	results, err := client.ListAll(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for i, item := range results.Values() {
		log.Println(*item.Name)
		log.Printf("%d %s", i, item.VirtualMachineProperties.HardwareProfile.VMSize)
	}
}

func getAuthorizer() autorest.Authorizer {
	// Authenticate with Azure
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return authorizer
}
