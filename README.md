# Alpha Vantage External Adapter 
External Adapter for Chainlink to query Alpha Vantages' APIs.

Built with [Bridges](https://github.com/linkpoolio/bridges).

https://www.alphavantage.co/documentation/

To give Alpha Vantages own description:
> Composed of a tight-knit community of researchers, engineers, and business professionals, Alpha Vantage Inc. is a leading provider of free APIs for realtime and historical data on stocks, forex (FX), and digital/crypto currencies. Our success is driven by rigorous research, cutting edge technology, and a disciplined focus on democratizing access to data.

### Contract Usage
To use this adapter on-chain, find a node that supports this adapter and build your request like so:
```
Chainlink.Request memory req = buildChainlinkRequest(jobId, this, this.fulfill.selector);
run.add("function", "GLOBAL_QUOTE");
run.add("symbol", "MSFT");
string[] memory copyPath = new string[](2);
copyPath[0] = "Global Quote";
copyPath[1] = "05. price";
```

### Setup Instructions
#### Local Install
Make sure [Golang](https://golang.org/pkg/) is installed.

Build:
```
make build
```

Then run the adapter:
```
API_KEY=apikey ./alphavantage-adapter
```

#### Docker
To run the container:
```
docker run -it -e API_KEY=apikey -p 8080:8080 linkpool/alphavantage-adapter
```

Container also supports passing in CLI arguments.

You can add and modify the keys to match what's specified in the API documentation.

### Usage

```
curl -X POST -H 'Content-Type: application/json' \
-d @- << EOF
{
	"jobRunId": "1234",
	"data": {
		"function": "GLOBAL_QUOTE",
		"symbol": "MSFT"
	}
}
EOF
```
Response:
```json
{
    "jobRunId": "1234",
    "status": "completed",
    "error": null,
    "pending": false,
    "data": {
        "Global Quote": {
            "01. symbol": "MSFT",
            "02. open": "133.7900",
            "03. high": "135.6500",
            "04. low": "131.8284",
            "05. price": "135.2800",
            "06. volume": "26682074",
            "07. latest trading day": "2019-08-07",
            "08. previous close": "134.6900",
            "09. change": "0.5900",
            "10. change percent": "0.4380%"
        },
        "function": "GLOBAL_QUOTE",
        "symbol": "MSFT"
    }
}
```
