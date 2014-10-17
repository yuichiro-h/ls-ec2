package main

import (
	"fmt"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/ec2"
	"log"
	"sort"
)

type Instance struct {
	Name      string
	PrivateIp string
	Ip        string
}

type Instances []Instance

func (in Instances) Len() int {
	return len(in)
}

func (in Instances) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func (in Instances) Less(i, j int) bool {
	return in[i].Name < in[j].Name
}

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
	var instances Instances
	for _, r := range res.Reservations {
		for _, i := range r.Instances {
			name := "None"
			for _, t := range i.Tags {
				switch t.Key {
				case "Name":
					name = t.Value
				}
			}
			instances = append(instances, Instance{Name: name, PrivateIp: i.PrivateIPAddress, Ip: i.IPAddress})
		}
	}

	sort.Sort(instances)

	for _, i := range instances {
		fmt.Printf("%-13s - %-13s - %s\n", i.PrivateIp, i.Ip, i.Name)
	}
}
