# GOpenDOAR
Simple (Golang) harvesting client for OpenDOAR.

[OpenDOAR](https://v2.sherpa.ac.uk/opendoar/) is a directory of open-access repositories, maintained by Jisc. This client harvests metadata records from OpenDOAR in JSON format and writes them to a local folder as JSON files. 

### Instructions

1. Build the Go sources into a binary in the usual way.
2. All configuration is in the `config/config.yaml` file. The `apiKeyENVKey` property expects the name of a local ENV variable containing your API key for OpenDOAR. You can get an API key from [this page](https://v2.sherpa.ac.uk/api/). 
3. Assuming that you built the sources to an executable called `gopendoar`, then run with the following command line: `gopendoar harvest --config=./config/config.yaml --debug=true`

This is entirely **unsupported** software.