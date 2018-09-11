package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	sg = flag.String("sg", "", "Security group to add tcp 80 open rule.")

	cloudflareIpv4 = []string{
		"103.21.244.0/22",
		"103.22.200.0/22",
		"103.31.4.0/22",
		"104.16.0.0/12",
		"108.162.192.0/18",
		"131.0.72.0/22",
		"141.101.64.0/18",
		"162.158.0.0/15",
		"172.64.0.0/13",
		"173.245.48.0/20",
		"188.114.96.0/20",
		"190.93.240.0/20",
		"197.234.240.0/22",
		"198.41.128.0/17",
	}

	cloudflareIpv6 = []string{
		"2400:cb00::/32",
		"2405:8100::/32",
		"2405:b500::/32",
		"2606:4700::/32",
		"2803:f800::/32",
		"2c0f:f248::/32",
		"2a06:98c0::/29",
	}
)

func main() {
	flag.Parse()

	session.Must(session.NewSession())
	svc := ec2.New(session.New())

	ipRanges, ipv6Ranges := []*ec2.IpRange{}, []*ec2.Ipv6Range{}

	for _, ipAddress := range cloudflareIpv4 {
		ipRanges = append(ipRanges, &ec2.IpRange{CidrIp: aws.String(ipAddress)})
	}

	for _, ipAddress := range cloudflareIpv6 {
		ipv6Ranges = append(ipv6Ranges, &ec2.Ipv6Range{CidrIpv6: aws.String(ipAddress)})
	}

	input := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: sg,
		IpPermissions: []*ec2.IpPermission{
			{
				FromPort:   aws.Int64(80),
				IpProtocol: aws.String("tcp"),
				IpRanges:   ipRanges,
				Ipv6Ranges: ipv6Ranges,
				ToPort:     aws.Int64(80),
			},
			{
				FromPort:   aws.Int64(443),
				IpProtocol: aws.String("tcp"),
				IpRanges:   ipRanges,
				Ipv6Ranges: ipv6Ranges,
				ToPort:     aws.Int64(443),
			},
		},
	}

	result, err := svc.AuthorizeSecurityGroupIngress(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
