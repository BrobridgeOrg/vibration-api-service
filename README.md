# Vibration API Service

API service expose APIs to control timer mechanism.

## Update proto definition

Rebuild to apply `proto` changes, just run commands below:

```shell
cd pb
protoc --go_out=plugins=grpc:. *.proto
```

## License

Licensed under the MIT License

## Authors

Copyright(c) 2020 Fred Chien <<fred@brobridge.com>>
