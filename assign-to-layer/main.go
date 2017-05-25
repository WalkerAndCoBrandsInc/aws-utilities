package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/opsworks"
)

var (
	AWS_REGION = os.Getenv("AWS_REGION")
	STACK_ID   = os.Getenv("STACK_ID")
	LAYER_ID   = os.Getenv("LAYER_ID")
)

func main() {
	log.Printf("ENV: AWS_REGION=%s STACK_ID=%s LAYER_ID=%s", AWS_REGION, STACK_ID, LAYER_ID)

	sess := session.Must(session.NewSession())
	svc := opsworks.New(sess)

	instancesOutput, err := getInstances(svc)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Fetched %v instances in stack\n", len(instancesOutput.Instances))

	for _, instance := range instancesOutput.Instances {
		log.Printf("Instance: %s has status: %s\n", *instance.InstanceId, *instance.Status)

		if *instance.Status != "registered" {
			continue
		}

		if err := assignToLayer(svc, *instance.InstanceId); err != nil {
			log.Printf("ERROR assigning instance: %s; ERR: ", *instance.InstanceId, err)
			continue
		}

		log.Printf("Assigned: %s\n", *instance.InstanceId)
	}
}

func getInstances(svc *opsworks.OpsWorks) (*opsworks.DescribeInstancesOutput, error) {
	params := &opsworks.DescribeInstancesInput{StackId: aws.String(STACK_ID)}
	return svc.DescribeInstances(params)
}

func assignToLayer(svc *opsworks.OpsWorks, instanceId string) error {
	params := &opsworks.AssignInstanceInput{
		InstanceId: aws.String(instanceId),
		LayerIds:   []*string{aws.String(LAYER_ID)},
	}

	_, err := svc.AssignInstance(params)
	return err
}
