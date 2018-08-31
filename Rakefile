task :build do
  sh "go build -o bin/eb-env eb-env/*.go"
  sh "go build -o bin/yaml-check yaml-check/*.go"
end

task :build_zip => [:build] do
  sh "zip -r aws-utilities bin/"
end

task :upload => [:build_zip] do
  sh " aws s3 cp aws-utilities.zip s3://form-cookbooks/aws-utilities.zip"
end
