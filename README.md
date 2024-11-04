# EVM-Custom-Node-Precompile

## Project Goals

- Develop a Custom Precompile for EVM Nodes:
  - Extend the Ethereum Virtual Machine (EVM) by adding a new precompile to the Geth (Go Ethereum) client.
  - Enhance the capabilities of the EVM to support specialized operations that improve performance.
- Optimize Rollup Performance:
  - Increase the throughput of Layer 2 solutions (rollups) by reducing computational overhead.
  - Enable faster transaction processing and lower gas csts for rollup operations.
- Enhance On-Chain Computation:
  - Provide more efficient execution of complex computations directly on-chain.
  - Improve the scalability and efficiency of smart contracts require intensive processing.

## Custom Precompile's Core Operations

- Batch Processing of Transactions
  - Process a batch of Layer 2 transactions in a single call to reduce the number of seperate transactions, saving both gas and processing time.
- Compression of Transaction Data:
  - Implement a function that compresses data before storing it on-chain, reducing the cost of storage and helping optimize rollup efficiency.
- Merkle Tree Proofs of Data Integity:
  - Add support for constructing or verifying Merkle trees directly in the precompile. This could improve rollup verification by quickly validating state changes.
- Matrix Multiplication or Cryptographic Operations:
  - Enable support for complex mathematical functions or cryptographic operations like elliptic curve pairing, reducing the need for such computations in regular EVM contracts.

## How to use the project

Note: This README assumes that the initial required dependencies like the geth, go and go compiler, solidity and solidity compiler are already installed on your system.

In order to use the project, first clone the repo onto your computer:

```sh
git clone https://github.com/Cass402/EVM-Custom-Node-Precompile.git
```

then head into the cloned repository and the clone the go-ethereum fork into the EVM-Custom-Node-Precompile folder using the following command in the terminal:

```sh
git clone https://github.com/Cass402/go-ethereum.git
```

Navigate to the cloned go-ethereum folder using:

```sh
cd go-ethereum
```

and finally make the geth which compiles all the required code and the precompile contracts:

```sh
make geth
```
