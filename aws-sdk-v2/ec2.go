package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func getAWSConfig(ctx context.Context, region string) (aws.Config, error) {
	return config.LoadDefaultConfig(ctx, config.WithRegion(region))
}

func main() {
	// get instance id from cli
	instance_id := flag.String("instance_id", "", "Provide EC2 instance id")
	region := flag.String("region", "ap-northeast-1", "AWS region to use")
	flag.Parse()

	// context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Validate required parameters
	if *instance_id == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Get AWS Config
	cfg, err := getAWSConfig(ctx, *region)

	if err != nil {
		log.Fatalf("Failed to load AWS Config: %v", err)
	}

	// Create EC2 client
	ec2Client := ec2.NewFromConfig(cfg)

	// Set Input for EC2 API
	input := &ec2.DescribeInstanceStatusInput{
		InstanceIds:         []string{*instance_id},
		IncludeAllInstances: aws.Bool(true),
	}

	// Call AWS API
	result, err := ec2Client.DescribeInstanceStatus(context.TODO(), input)

	if err != nil {
		log.Fatalf("Error Getting Instance Status: %v", err)
	}

	// Converting Result to Json
	jsonData, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(jsonData))

	if len(result.InstanceStatuses) == 0 {
		log.Println("No status information found")
		return
	}
	fmt.Println("Instance Status Information:")
	for _, status := range result.InstanceStatuses {
		log.Println("Status", status.InstanceStatus.Status)
		log.Println("State", status.InstanceState.Name)
		log.Println("InstanceId", *status.InstanceId)
	}

}
