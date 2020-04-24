#!/usr/bin/env bash

echo "Installing Harmony Tests"
rm -rf hmy && curl -LO https://harmony.one/hmycli && mv hmycli hmy && chmod u+x hmy
rm -rf tests && curl -LOs http://tools.harmony.one.s3.amazonaws.com/release/linux-x86_64/harmony-tests/tests && chmod u+x tests
rm -rf testcases && curl -LOs http://tools.harmony.one.s3.amazonaws.com/release/linux-x86_64/harmony-tests/testcases.tar.gz && tar -xzf testcases.tar.gz && rm -rf testcases.tar.gz
rm -rf config.yml && curl -LOs https://raw.githubusercontent.com/harmony-one/harmony-tests/master/config.yml
mkdir -p keys/stressnet keys/pangaea
echo "Harmony Tests have now been downloaded!"
echo ""
echo "Make sure to either add keystore files to keys/testnet or create the file keys/testnet/private_keys.txt and add testnet private keys to it"
echo "The tests can be run on multiple networks - testnet is just used as an example. To run the tests on other networks, e.g. localnet, use --network localnet"
echo ""
echo "When you've added keyfiles or private keys, invoke the tests using ./tests"
echo ""
echo "To see all available configuration options, run ./tests --help or check out the configuration specified in config.yml"
