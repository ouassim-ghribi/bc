# Merkel-Tree

### Setup

1. First clone the repository:

   ```bash
   git clone git@github.com:ouassim-ghribi/bc.git

   cd bc
   ```

2. This implementation uses the new Bun javascript runtime so in order to get up and running you need first to install bun on your PC (I didn't have time to dockerize the whole thing :/)

   ```bash
   curl -fsSL https://bun.sh/install | bash
   ```

   To verify installation:

   ```bash
   bun -v
   ```

3. Then

### Folder structure

    ├── src                 // Root folder
        ├── api.json        // api response from bitcoin blockchain
        ├── api.ts          // retrieve transactions from bitcoin blockchain
        ├── constants.ts    // constants file
        ├── index.ts        // main file
        ├── merkle.ts       // implementation logic for merkle tree
        ├── types.ts        // types used
        └── utils.ts        // utility functions
    ├── .gitignore
    ├── package.json
    └── tsconfig.json

4. To run just do the following:

   ```bash
   bun run src/index.ts
   ```

### References:

- [Merkle Trees Wikipedia](https://en.wikipedia.org/wiki/Merkle_tree)
- [Merkle Proofs Explained](https://medium.com/crypto-0-nite/merkle-proofs-explained-6dd429623dc5)
- [Merkle Tree: A Beginners Guide](https://kba.ai/merkle-tree-a-beginners-guide/)

- [Merkle Proof](https://computersciencewiki.org/index.php/Merkle_proof)
