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
If you already have a funded account, simply pass the address of that account using `--address`

If you want to import and use private keys:
e.g. for testnet -> `nano keys/testnet/private_keys.txt` - and save your list of private keys to this file

If you want to import and use keystore files:
e.g. for testnet -> `cp -r keystore-folder keys/testnet`

Harmony TF will automatically identify keyfiles no matter what you call the folders or what the files are called - as long as they reside under `keys/testnet` (or whatever network you're using) they'll be identified.

### Running tests
To run all test cases:
`./tests --network NETWORK`

To run all test cases using an already funded address:
`./tests --network NETWORK --address YOUR_FUNDED_ADDRESS`

To connect to custom defined RPC nodes:
`./tests --network NETWORK --mode custom --address YOUR_FUNDED_ADDRESS`

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

## Writing test cases
Test cases are defined as YAML files and are placed in testcases/ - see this folder for existing test cases and how to impelement test cases.

There are some requirements for these files - this README will be updated with a list of these eventually :)
