export async function fetchBlockTransactions() {
  const block = await fetch(
    "https://blockchain.info/q/latesthash?cors=true"
  ).then(async (r) => await r.text());

  const blockData = await fetch(
    `https://blockchain.info/rawblock/${block}?cors=true`
  ).then((res) => res.json());

  const tx = blockData.tx.map((tx: { hash: string }) => tx.hash);

  Bun.write("api.json", JSON.stringify({ tx }, null, 2));
}
