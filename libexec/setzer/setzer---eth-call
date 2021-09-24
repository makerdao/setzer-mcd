#!/usr/bin/env bash
set -eo pipefail

# all in hex format
contract_addr=${1#0x}
method_id=${2#0x}
args=${3#0x} # ABI-encoded

: ${ETH_RPC_URL:=http://127.0.0.1:8545}
: ${ETH_BLOCK:="latest"}

hex_0x=$(curl \
  -sS \
  -X POST \
  -H 'Content-Type: application/json' \
  --data '{"jsonrpc":"2.0","method":"eth_call","params":[{"to":"'"0x$contract_addr"'","data":"'"0x$method_id$args"'"},"'"$ETH_BLOCK"'"],"id":1}' \
  "$ETH_RPC_URL" | jshon -e result -u)

if ! [[ $hex_0x == 0x* ]]; then
  echo >&2 "Error: invalid JSON-RPC result received: $hex_0x"
  exit 1
fi

# output in an uppercase hex format, e.g. AABB
hex=${hex_0x#0x}
printf '%s' "${hex^^}"
