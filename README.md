# AWS Utilities

Bundle of small utilities programs for use with AWS ElasticBeanstalk.

## assign-to-layer

assign-to-layer assigns instances that are already registered to the given stack to the given layer.

    env AWS_REGION=us-east-1 STACK_ID= LAYER_ID= go run assign-to-layer/main.go

## eb-env

eb-env fetches given environment variable from the environment or defaults to the ElasticBeanstalk [file ](/opt/elasticbeanstalk/deploy/configuration/containerconfiguration) that contains configuration options, along with environment variable. If environment variable is not found, it returns an empty string.

    sudo go run eb-env/main.go -name=RACK_ENV

## yaml-check

yaml-check accepts a directory and checks if the files in there pass YAML syntax. You can optionally pass the extension of file to look for (since EB .ebextensions require .config suffix). If any of the file is not valid syntax, it returns a 1 exit code.

    go run yaml-check/main.go -dir=. -ext=CONFIG

## Build

To build binaries:

    env GOOS=linux rake build

To build binaries and zip them:

    env GOOS=linux rake build_zip
