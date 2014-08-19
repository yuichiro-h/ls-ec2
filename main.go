package main

import (
	"fmt"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/ec2"
	"log"
)

func main() {
	env, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	ec2 := ec2.New(env, aws.APNortheast)
	res, err := ec2.DescribeInstances([]string{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range res.Reservations {
		for _, i := range r.Instances {
			for _, t := range i.Tags {
				if t.Key == "Name" {
					fmt.Printf("%s %s\n", t.Value, i.IPAddress)
				}
			}
		}
	}
}
