package vms

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/azure-sdk-for-go/services/preview/commerce/mgmt/2015-06-01-preview/commerce"
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

	// Authenticate with Azure
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatalf(err.Error())
	}

	client := commerce.NewRateCardClient(sid)
	client.Authorizer = authorizer
	filter := "OfferDurableId eq 'MS-AZR-0121p' and Currency eq 'USD' and Locale eq 'en-US' and RegionInfo eq 'US'"
	result, err := client.Get(ctx, filter)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, item := range *result.Meters {
		if *item.MeterCategory == "Virtual Machines" {
			fmt.Printf("Category: %s, \t\t  Sub Category: %s, \t\t Rate: %f \n", *item.MeterCategory, *item.MeterSubCategory, *item.MeterRates["0"])
		}
	}

	log.Println("Currency: ", *result.Currency)

}
