VERSION = 0.1

task :build do
  # This script is not longer required.
  # sh "go build -o bin/assign-to-layer-#{VERSION} assign-to-layer/*.go"
  sh "go build -o bin/eb-env-#{VERSION} eb-env/*.go"
end

task :build_zip => [:build] do
  sh "zip -r aws-utilities-#{VERSION}.zip bin/"
end

task :upload => [:build_zip] do
  sh " aws s3 cp aws-utilities-#{VERSION}.zip s3://form-cookbooks/aws-utilities-#{VERSION}.zip"
end
