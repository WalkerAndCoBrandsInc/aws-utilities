package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	flag.Parse()

	sg := os.Getenv("ELB_SECURITY_GROUP")
	if sg == "" {
		fmt.Println("No security group found under env 'ELB_SECURITY_GROUP'")
		os.Exit(1)
	}

	resp, err := http.Get("http://169.254.169.254/latest/meta-data/public-ipv4")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	ipAddress := string(body)

	session.Must(session.NewSession())
	svc := ec2.New(session.New())
	input := &ec2.AuthorizeSecurityGroupEgressInput{
		DryRun:  aws.Bool(true),
		GroupId: aws.String(sg),
		IpPermissions: []*ec2.IpPermission{
			{
				FromPort:   aws.Int64(80),
				IpProtocol: aws.String("tcp"),
				IpRanges: []*ec2.IpRange{
					{CidrIp: aws.String(fmt.Sprintf("%s/32", ipAddress))},
				},
				ToPort: aws.Int64(80),
			},
		},
	}

	result, err := svc.AuthorizeSecurityGroupEgress(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
