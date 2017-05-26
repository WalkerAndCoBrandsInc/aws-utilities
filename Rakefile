VERSION = 0.1

task :build do
  sh "go build -o bin/assign-to-layer-#{VERSION} assign-to-layer/*.go"
  sh "go build -o bin/eb-env-#{VERSION} eb-env/*.go"
end

task :build_zip => [:build] do
  sh "zip -r aws-utilities-#{VERSION}.zip bin/"
end
