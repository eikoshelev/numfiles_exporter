# numfiles_exporter

Monitoring the number of files in the specified directories

### Two options for work:
* working with one target directory (the `-directory` flag) —determines all subdirectories and reads the number of files in each of them:
```
./numfiles_exporter -directory="/opt/test"
```
* working with several separately taken directories - defined in the targets.yaml file (the `-targets` flag):
```
targetDirectories:  # yaml array
- "/etc"
- "/var"
- "/opt/test"
```

Metrics will be available by `localhost:9095/metrics`:

```
# HELP number_of_files Number of files in the target directory
# TYPE number_of_files gauge
number_of_files{directory="/etc"} 87.0
number_of_files{directory="/var"} 24.0
number_of_files{directory="/opt/test"} 3.0
```

### Flags used:
```
./numfiles_exporter -h
Usage of ./numfiles_exporter:
  -directory string
    	directory with target subdirectories
  -handler string
    	handler for which metrics will be available (default "/metrics")
  -port string
    	port of return of metrics (default ":9095")
  -targets string
    	file with separate target directories
  -timeout int
    	interval (seconds) of checking the number of files in the target directory (default 10)
```

### Usage:
```
git clone https://github.com/eikoshelev/numfiles_exporter.git
```
```
cd numfiles_exporter
```
```
go build
```

#### For docker container (for alpine base image):
```
GOOS=linux GOARCH=amd64 go build
```
```
docker build -t numfiles_exporter:0.1.0 .
```
