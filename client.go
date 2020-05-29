package main

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/eventgrid/mgmt/2019-06-01/eventgrid"

	"github.com/Azure/go-autorest/autorest/azure/auth"
	//"github.com/Azure/go-autorest/autorest/to"
)

func main() {

	eventSubClient := eventgrid.NewEventSubscriptionsClient("<subscriptionID>")

	// create an authorizer from env vars or Azure Managed Service Idenity
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err == nil {
		eventSubClient.Authorizer = authorizer
	}

	// eventSubscriptionInfo :=
	// What is EventSubscription struct?
	eventSubClient.CreateOrUpdate(context.Background(), "<resourceGroupName>", "<eventSubName>", eventSubscriptionInfo)

}
