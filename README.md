# numfiles_exporter

Мониторинг количества файлов в указанных директориях

### Два варианта работы:
* работа с одной целевой директорией (флаг `-directory`) - определяет все вложенные директории и считывает количество файлов в каждой из них:
```
./numfiles_exporter -directory="/opt/test"
```
* работа с несколькими, отдельно взятыми, директориями - определяются в файле targets.yaml (флаг `-targets`):
```
targetDirectories:  # yaml array
- "/etc"
- "/var"
- "/opt/test"
```

Метрики будут доступны по `localhost:9095/metrics`:

```
# HELP number_of_files Number of files in the target directory
# TYPE number_of_files gauge
number_of_files{directory="/etc"} 87.0
number_of_files{directory="/var"} 24.0
number_of_files{directory="/opt/test"} 3.0
```

### Используемые флаги:
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