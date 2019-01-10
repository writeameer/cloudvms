package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/pricing"
)

var (
	svc  *pricing.Pricing
	sess *session.Session
)

func main() {

	// Create session info from aws creds file
	sess, _ = session.NewSession(&aws.Config{Region: aws.String("ap-south-1")})

	// Create pricing client
	svc = pricing.New(sess)

	stuff()
}

func services() {
	input := &pricing.DescribeServicesInput{
		MaxResults:  aws.Int64(1),
		ServiceCode: aws.String("AmazonEC2"),
	}

	result, _ := svc.DescribeServices(input)

	//log.Println(result.Services[0])
	for _, service := range result.Services {
		for _, attribute := range service.AttributeNames {
			attributeValues(*attribute)
		}
	}

	//fmt.Println(result)
}
func attributeValues(attributeName string) {
	input := &pricing.GetAttributeValuesInput{
		AttributeName: aws.String(attributeName),
		MaxResults:    aws.Int64(2),
		ServiceCode:   aws.String("AmazonEC2"),
	}

	result, _ := svc.GetAttributeValues(input)

	fmt.Println(result)
}

func stuff() {
	input := &pricing.GetProductsInput{
		Filters: []*pricing.Filter{
			{
				Field: aws.String("serviceCode"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("AmazonEC2"),
			},
		},
		FormatVersion: aws.String("aws_v1"),
		MaxResults:    aws.Int64(1),
		ServiceCode:   aws.String("AmazonEC2"),
	}

	result, err := svc.GetProducts(input)

	log.Println(err)
	fmt.Println(result)
}
