import { merkleProof } from "./types";
import { Direction, sha256 } from "./utils";

export const ensureEven = (hashes: string[]) => {
  if (hashes.length % 2 !== 0) {
    hashes.push(hashes[hashes.length - 1]);
  }
};

export const getLeafNodeDirectionInMerkleTree = (
  hash: string,
  merkleTree: string[][]
): Direction => {
  const hashIndex = merkleTree[0].findIndex((h) => h === hash);
  return hashIndex % 2 === 0 ? Direction.LEFT : Direction.RIGHT;
};

export const generateMerkleRoot = (hashes: string[]): string => {
  if (!hashes || hashes.length == 0) {
    return "";
  }
  ensureEven(hashes);
  const combinedHashes = [];
  for (let i = 0; i < hashes.length; i += 2) {
    const hashPairConcatenated = hashes[i] + hashes[i + 1];
    const hash = sha256(hashPairConcatenated);
    combinedHashes.push(hash);
  }
  // If the combinedHashes length is 1, it means that we have the merkle root already
  // and we can return
  if (combinedHashes.length === 1) {
    return combinedHashes.join("");
  }
  return generateMerkleRoot(combinedHashes);
};

export const generate = (hashes: string[], tree: string[][]): string[] => {
  if (hashes.length === 1) {
    return hashes;
  }
  ensureEven(hashes);
  const combinedHashes = [];
  for (let i = 0; i < hashes.length; i += 2) {
    const hashesConcatenated = hashes[i] + hashes[i + 1];
    const hash = sha256(hashesConcatenated);
    combinedHashes.push(hash);
  }
  tree.push(combinedHashes);
  return generate(combinedHashes, tree);
};

export const generateMerkleTree = (hashes: string[]): string[][] => {
  if (!hashes || hashes.length === 0) {
    return [];
  }
  const tree = [hashes];
  generate(hashes, tree);
  return tree;
};

export const generateMerkleProof = (
  hash: string,
  hashes: string[]
): merkleProof[] | null => {
  if (!hash || !hashes || hashes.length === 0) {
    return null;
  }
  const tree = generateMerkleTree(hashes);
  const merkleProof = [
    {
      hash,
      direction: getLeafNodeDirectionInMerkleTree(hash, tree),
    },
  ];
  let hashIndex = tree[0].findIndex((h) => h === hash);
  for (let level = 0; level < tree.length - 1; level++) {
    const isLeftChild = hashIndex % 2 === 0;
    const siblingDirection = isLeftChild ? Direction.RIGHT : Direction.LEFT;
    const siblingIndex = isLeftChild ? hashIndex + 1 : hashIndex - 1;
    const siblingNode = {
      hash: tree[level][siblingIndex],
      direction: siblingDirection,
    };
    merkleProof.push(siblingNode);
    hashIndex = Math.floor(hashIndex / 2);
  }
  return merkleProof;
};

export const getMerkleRootFromMerkleProof = (
  merkleProof:
    | {
        hash: string;
        direction: Direction;
      }[]
    | null
): string => {
  if (!merkleProof || merkleProof.length === 0) {
    return "";
  }
  const merkleRootFromProof = merkleProof.reduce((hashProof1, hashProof2) => {
    if (hashProof2.direction === Direction.RIGHT) {
      const hash = sha256(hashProof1.hash + hashProof2.hash);
      return { hash, direction: Direction.RIGHT };
    }
    const hash = sha256(hashProof2.hash + hashProof1.hash);
    return { hash, direction: Direction.LEFT };
  });
  return merkleRootFromProof.hash;
};
