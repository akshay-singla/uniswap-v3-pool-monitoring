# Uniswap V3 Pool Monitoring

This project is a monitoring service for Uniswap V3 pools. It continuously tracks and logs essential data points of configured pools, stores them in a persistent datastore, and provides access to the data through a REST API.


## Prerequisites

Before running the monitoring service, make sure you have the following:

- An Infura project URL: You can generate the Infura project URL by creating an account on [Infura](https://infura.io/) and creating a new project. Take note of the project ID, as it will be used in the service.

- Set this in the ENV using "INFURA_URL" env variable.

- Pool address: Choose the Uniswap V3 pool you want to monitor and obtain its address. You can find the pool address on blockchain explorers like [Etherscan](https://etherscan.io/).

- Set this in the main.go in poolConfig or set this in line number 14.


## Functionality

The service fetches and calculates the following data points from the configured pool contracts every 12 blocks and saves them into a persistent datastore:

```
 {
    "token0Balance": "balance of token0 in the pool",
    "token1Balance": "balance of token1 in the pool",
    "tick": "price tick at block",
    "blockNumber": "block at which this data was read"
}
```


## The service exposes this data through the following REST endpoints:

#### Get the balances saved in the data store with a block query filter:

* GET /v1/api/pool/:pool_id?block={'latest','6969696'}

Response:
```
{
  "token0Balance": "balance of token0 in the pool at the blockNumber nearest to the given `block`",
  "token1Balance": "balance of token1 in the pool at the blockNumber nearest to the given `block`",
  "tick": "price tick nearest to the latest block",
  "blockNumber": "block at which this data was read"
}
```

### Get the historical values of the balances in the data store, along with the token deltas which represent the change since the previous blockNumber:

* GET /v1/api/pool/:pool_id/historic

Response:
```
[
  {
    "token0Balance": "balance of token0 in the pool at the blockNumber nearest to the given `block`",
    "token1Balance": "balance of token1 in the pool at the blockNumber nearest to the given `block`",
    "tick": "price tick nearest to the latest block",
    "blockNumber": "block at which this data was read"
  },
  ...
]
```

## Setup and Usage

1. Clone the repository:
    ```
        git clone https://github.com/your-username/uniswap-v3-pool-monitoring.git
    ```

2. Configure the project by updating the necessary values in the configuration file.

3. Build and run the application:
    ```
        go build
        ./uniswap-v3-pool-monitoring
    ```

4. The monitoring service will start running on localhost:8080 by default.
