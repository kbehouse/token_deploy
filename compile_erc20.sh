# ContractPath=/Users/kartik/arbitrage/tx-manager/
# ContractName=tx
docker run -v ${PWD}:/c \
        -v ${PWD}/erc20_bin_abi:/output \
        ethereum/solc:0.6.4  \
        @openzeppelin/=/c/@openzeppelin/ \
        -o /output --abi --bin /c/@openzeppelin/token/ERC20/ERC20.sol

# abiben
mkdir erc20
abigen --bin=erc20_bin_abi/ERC20.bin --abi=erc20_bin_abi/ERC20.abi --pkg=erc20 --out=erc20/erc20.go