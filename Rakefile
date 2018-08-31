task :build do
  sh "env GOOS=linux go build -o bin/eb-env eb-env/*.go"
  sh "env GOOS=linux go build -o bin/yaml-check yaml-check/*.go"
  sh "env GOOS=linux go build -o bin/elb-sg elb-sg/*.go"
end

task :build_zip => [:build] do
  sh "zip -r aws-utilities bin/"
end

task :default => [:build_zip]
