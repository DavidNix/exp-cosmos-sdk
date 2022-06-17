#!/usr/bin/env sh

rm -rf $$HOME/.simapp
./build/regen keys add validator --keyring-backend test
./build/regen keys add validator --keyring-backend test
VALIDATOR_ADDRESS="$(./build/regen keys show validator -a --keyring-backend test)"
./build/regen init test1 --chain-id regen-1
./build/regen add-genesis-account "$VALIDATOR_ADDRESS" 100000000000stake
./build/regen gentx validator 100000000stake --chain-id regen-1 --keyring-backend test
./build/regen collect-gentxs
echo "Validator account:"
echo "$VALIDATOR_ADDRESS"
sleep 5
./build/regen start
