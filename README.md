# harmony-tests
harmony-tests executes a set of test cases for testing that regular and staking transactions work properly on Harmony's blockchain.

It uses the harmony-tf testing framework under the hood.

## Installation

```
rm -rf harmony-tests && mkdir -p harmony-tests && cd harmony-tests
bash <(curl -s -S -L https://raw.githubusercontent.com/harmony-one/harmony-tests/master/scripts/install.sh)
```

## Build/deploy

If you need to build/compile the tool from scratch:

Regular build:
`make all`

Static build:
`make static`

Static build, package all testcases and upload everything to Amazon S3:
`make upload-linux`

(Static builds have only been tested on Linux)

## Usage

### Funding

#### Existing address
If you already have a funded account, simply pass the address of that account using `--address`:

- `./tests --network NETWORK --address YOUR_FUNDED_ADDRESS`

#### Private keys
If you want to import and use private keys:

- `mkdir -p keys/NETWORK/`
- `nano keys/NETWORK/private_keys.txt`
- paste your private keys and save the file

#### Keystore files
If you want to import and use keystore files:

- `mkdir -p keys/NETWORK/`
- `cp -r PATH/TO/YOUR/KEYSTORE/FOLDER keys/NETWORK`

Harmony TF will automatically identify keyfiles no matter what you call the folders or what the files are called - as long as they reside under `keys/NETWORK` they'll be identified.

### Running tests
To run all test cases:
`./tests --network NETWORK`

To run all test cases using an already funded address:
`./tests --network NETWORK --address YOUR_FUNDED_ADDRESS`

To connect to custom defined RPC nodes configured using config.yml:
`./tests --network NETWORK --mode custom --address YOUR_FUNDED_ADDRESS`

To connect to custom nodes using command line args:
`./tests --network NETWORK --mode custom --nodes http://SHARD0NODEIP:9500,http://SHARD1NODEIP:9500 --address YOUR_FUNDED_ADDRESS`

### Specific test cases

#### Staking

- Only run staking tests: `./tests --network NETWORK --test staking`
- Only run staking -> create validators: `./tests --network NETWORK --test staking/validators/create`
- Only run staking -> edit validators: `./tests --network NETWORK --test staking/validators/edit`
- Only run staking -> delegation: `./tests --network NETWORK --test staking/delegation/delegate`
- Only run staking -> undelegation: `./tests --network NETWORK --test staking/delegation/undelegate`

#### Transactions
- Only run tx tests: `./tests --network NETWORK --test transactions`
- Only run tx -> cross app shard tests: `./tests --network NETWORK --test transactions/cross_app_shard`
- Only run tx -> cross beacon shard tests: `./tests --network NETWORK --test transactions/cross_beacon_shard`
- Only run tx -> same app shard tests: `./tests --network NETWORK --test transactions/same_app_shard`
- Only run tx -> same beacon shard tests: `./tests --network NETWORK --test transactions/same_beacon_shard`

### Exporting test data
Harmony-tests currently supports exporting the test suite results as a CSV report:

`./tests --network NETWORK --address YOUR_FUNDED_ADDRESS --export csv`

JSON export is also planned, but not yet implemented.

## Writing test cases
Test cases are defined as YAML files and are placed in testcases/ - see this folder for existing test cases and how to impelement test cases.

There are some requirements for these files - this README will be updated with a list of these eventually :)
