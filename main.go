package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"gopkg.in/ini.v1"
)

var regions = []string{
	"us-east-1",
	"eu-west-1",
	"eu-west-2", // added region
}

func listEC2Instances(region, profile string) error {
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)
	if err != nil {
		return fmt.Errorf("[%s/%s] config error: %w", profile, region, err)
	}

	client := ec2.NewFromConfig(cfg)
	result, err := client.DescribeInstances(context.Background(), &ec2.DescribeInstancesInput{})
	if err != nil {
		return fmt.Errorf("[%s/%s] describe error: %w", profile, region, err)
	}

	count := 0
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			if instance.State.Name == types.InstanceStateNameRunning {
				count++
			}
		}
	}
	fmt.Printf("[%s/%s] Running instances: %d\n", profile, region, count)
	return nil
}

func getProfiles() ([]string, error) {
	awsConfigPath := os.Getenv("AWS_CONFIG_FILE")
	if awsConfigPath == "" {
		awsConfigPath = os.ExpandEnv("$HOME/.aws/config")
	}
	cfgFile, err := ini.Load(awsConfigPath)
	if err != nil {
		return nil, err
	}

	var profiles []string
	for _, section := range cfgFile.Sections() {
		if strings.HasPrefix(section.Name(), "profile ") {
			profiles = append(profiles, strings.TrimPrefix(section.Name(), "profile "))
		}
	}
	return profiles, nil
}

func main() {

	profiles, err := getProfiles()
	if err != nil {
		log.Fatalf("Failed to get profiles: %v", err)
	}
	start := time.Now()
	// Synchronous version (commented out )

	for _, profile := range profiles {
		for _, region := range regions {
			if err := listEC2Instances(region, profile); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
	fmt.Printf("\nDone in %s\n", time.Since(start))
	start2 := time.Now()

	// Concurrent version with WaitGroup
	var wg sync.WaitGroup
	for _, profile := range profiles {
		for _, region := range regions {
			wg.Add(1)
			go func(p, r string) {
				defer wg.Done()
				if err := listEC2Instances(r, p); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}(profile, region)
		}
	}

	wg.Wait()

	fmt.Printf("\nDone in %s\n", time.Since(start2))
}
