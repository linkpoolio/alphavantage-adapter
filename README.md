# Alpha Vantage External Adaptor ![Travis-CI](https://travis-ci.org/linkpoolio/alpha-vantage-cl-ea.svg?branch=master) [![codecov](https://codecov.io/gh/linkpoolio/alpha-vantage-cl-ea/branch/master/graph/badge.svg)](https://codecov.io/gh/linkpoolio/alpha-vantage-cl-ea)
External Adaptor for Chainlink to allow access to Alpha Vantages' APIs.

To give Alpha Vantages own description:
> Composed of a tight-knit community of researchers, engineers, and business professionals, Alpha Vantage Inc. is a leading provider of free APIs for realtime and historical data on stocks, forex (FX), and digital/crypto currencies. Our success is driven by rigorous research, cutting edge technology, and a disciplined focus on democratizing access to data.

### Preconditions
Retrieve an API key from Alpha Vantage for free:

https://www.alphavantage.co/support/#api-key

### Setup Instructions
#### Local Install
Make sure [Golang](https://golang.org/pkg/) is installed.

Build:
```
make build
```

Then run the adaptor:
```
./alpha-vantage-cl-ea -p <port> -apiKey <API_KEY>
```

##### Arguments

| Char   | Default  | Usage |
| ------ |:--------:| ----- |
| p      | 8080     | Port number to serve |
| apiKey      | nil     | Your API Key for Alpha Vantage |

#### Docker
To run the container:
```
docker run -it -p 8080:8080 linkpoolio/alpha-vantage-cl-ea -apiKey=yourkey
```

Container also supports passing in CLI arguments.

### Usage

To call the API, you need to send a POST request to `http://localhost:<port>/query` with the request body being of the ChainLink `RunResult` type.

The `data` passed in should match the GET parameter options on Alpha Vantage's API docs:

https://www.alphavantage.co/documentation/

For example:
```
curl -X POST -H 'Content-Type: application/json' -d '{ "jobRunId": "1234", "data": {"function": "TIME_SERIES_MONTHLY_ADJUSTED", "symbol": "MSFT"}}' http://localhost:8080/query
```
Should return something similar to:
```json
{
    "jobRunId": "1234",
    "status": "",
    "error": null,
    "pending": false,
    "data": {
        "Meta Data": {
            "1. Information": "Monthly Adjusted Prices and Volumes",
            "2. Symbol": "MSFT",
            "3. Last Refreshed": "2018-08-07 11:15:38",
            "4. Time Zone": "US/Eastern"
        },
        ....
    }
}
```

### ChainLink Node Setup

To integrate this adaptor with your node, use the following commands:

**Add Bridge Type**
```
curl -X POST -H 'Content-Type: application/json' -d '{"name":"alphavantage","url":"http://localhost:8080/query"}' http://localhost:6688/v2/bridge_types
```

**Create Spec**
```
curl -X POST -H 'Content-Type: application/json' -d '{"initiators":[{"type":"web"}],"tasks":[{"type":"alphavantage"},{"type":"noop"}]}' http://localhost:6688/v2/specs
```

**New Spec Run**

Notice the parameters `function` and `symbol`. These are passed into the external adaptor by the node, they map up to the GET parameters on the documentation.
```
curl -X POST -H 'Content-Type: application/json' -d '{"function": "TIME_SERIES_MONTHLY_ADJUSTED", "symbol": "MSFT"}' http://localhost:6688/v2/specs/<specId>/runs
```

### Contract Usage
To use this adaptor within a Solidity contract, add the following keys to your ChainLink run:
```
ChainlinkLib.Run memory run = newRun(specId, this, "fulfill(bytes32,bytes32)");
run.add("function", "TIME_SERIES_MONTHLY_ADJUSTED");
run.add("symbol", "MSFT");
```

You can add and modify the keys to match what's specified in the API documentation.
