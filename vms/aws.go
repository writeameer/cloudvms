package vms

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pricing"
)

// ListAws Lists VMS on AWS and their details
func ListAws() {

	// Create session info from aws creds file
	sess, _ := session.NewSession(&aws.Config{Region: aws.String("ap-south-1")})

	// Create pricing client
	svc := pricing.New(sess)

	input := &pricing.GetProductsInput{
		Filters: []*pricing.Filter{
			{
				Field: aws.String("servicecode"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("AmazonEC2"),
			},
		},
		FormatVersion: aws.String("aws_v1"),
		ServiceCode:   aws.String("AmazonEC2"),
	}

	result, err := svc.GetProducts(input)

	for i, item := range result.PriceList {
		product := item["product"].(map[string]interface{})
		attributes := product["attributes"].(map[string]interface{})

		instanceType := attributes["instanceType"].(string)
		memory := attributes["memory"].(string)
		vcpu := attributes["vcpu"].(string)
		family := attributes["instanceFamily"].(string)
		fmt.Printf("%d. Instance Type: %s, \t\t Family: %s, \t\t Memory: %s, \t\t vCPU: %s \n", i+1, instanceType, family, memory, vcpu)
	}

	if err != nil {
		log.Printf("Error: %v", err)
	}

}
