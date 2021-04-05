# Token Deploy


## Download ERC20 contract
1. Download @openzeppelin from https://github.com/OpenZeppelin/openzeppelin-contracts/tree/solc-0.6

2. copy & rename contracts/ to @openzeppelin/

## Compile

```
sh compile_token.sh
```

## Usage

```
go run main.go config.go  \
    -config=config_rinkeby.json \
    -name="MYToken" \
    -symbol=MYT \
    -supply=1000000000
```

or parent folder is token_deploy

```
go build
```

```
./token_deploys \
    -config=config_rinkeby.json \
    -name="MYToken" \
    -symbol=MYT \
    -supply=1000000000
```