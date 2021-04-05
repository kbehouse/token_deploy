# ContractPath=/Users/kartik/arbitrage/tx-manager/
# ContractName=tx
docker run -v ${PWD}:/c \
        -v ${PWD}/bin_abi:/output \
        ethereum/solc:0.6.4  \
        @openzeppelin/=/c/@openzeppelin/ \
        -o /output --abi --bin /c/solc/token.sol 

# abiben
[ ! -d "token/" ] && mkdir token
abigen --bin=bin_abi/Token.bin --abi=bin_abi/Token.abi --pkg=token --out=token/token.go