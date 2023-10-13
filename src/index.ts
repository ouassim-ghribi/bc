import {
  generateMerkleProof,
  generateMerkleRoot,
  generateMerkleTree,
  getMerkleRootFromMerkleProof,
} from "./merkle";

import txs from "./api.json";

const startTime = Date.now();

const merkleRoot = generateMerkleRoot(txs.tx);

const generatedMerkleProof = generateMerkleProof(txs.tx[9], txs.tx);

const merkleTree = generateMerkleTree(txs.tx);

const merkleRootFromMerkleProof =
  getMerkleRootFromMerkleProof(generatedMerkleProof);

const endTime = Date.now();
console.log("merkleRoot: ", merkleRoot);
console.log("generatedMerkleProof: ", generatedMerkleProof);
console.log("merkleTree: ", merkleTree);

console.log("merkleRootFromMerkleProof: ", merkleRootFromMerkleProof);
console.log(
  "merkleRootFromMerkleProof === merkleRoot: ",
  merkleRootFromMerkleProof === merkleRoot
);

const executionTime = endTime - startTime;
console.log(`Execution time: ${executionTime} milliseconds`);

// Fetch latest block transaction data and write to api.json file
// fetchBlockTransactions();
