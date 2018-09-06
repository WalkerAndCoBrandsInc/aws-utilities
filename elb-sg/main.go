package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	sg = flag.String("sg", "", "Security group to add tcp 80 open rule.")
)

func main() {
	flag.Parse()

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
	input := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: sg,
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

	result, err := svc.AuthorizeSecurityGroupIngress(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
