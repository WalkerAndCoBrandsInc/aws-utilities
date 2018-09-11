# AWS Utilities

Small utilities programs for use with AWS ElasticBeanstalk.

## eb-env

eb-env fetches given environment variable from the environment or defaults to the ElasticBeanstalk [file ](/opt/elasticbeanstalk/deploy/configuration/containerconfiguration) that contains configuration options, along with environment variable. If environment variable is not found, it returns an empty string.

NOTE, this needs sudo.

    sudo go run eb-env/main.go -name=RACK_ENV

## yaml-check

yaml-check accepts a directory and checks if the files in there pass YAML syntax. You can optionally pass the extension of file to look for (since EB .ebextensions require .config suffix). If any of the file is not valid syntax, it returns a 1 exit code.

    ./yaml-check -dir=. -ext=CONFIG

## elb-sg

elb-sg adds open port tcp 80 of the AWS instance running the script to the given security group. This is used by `go-proxy-*` to open up ingress rules in appropriate `bevel-*` ELB.

    ./elb-sg -sg=<sg-insert-id-here>

## cloudflare-sg

cloudflare-sg opens tcp port 80, 443 of all known cloudflare ips as of (Sept 11, 2018) as specified in https://www.cloudflare.com/ips to given security group.

    ./cloudflare-sg -sg=<sg-insert-id-here>

This script is meant to be on once for each of `bevel-*` load balancer security group, hence not part of aws-utilities.zip.

## Build (IMPORTANT)

To build and zip the libraries:

    rake

Commit the changes including aws-utilities.zip and then open a PR to develop. Once merged, EB instances will clone the repo and unzip the binaries from aws-utilities.zip. This is important, without it your changes won't be accessible!
