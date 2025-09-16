import { AtlasDashboard } from './core/dashboard';

export interface BalanceData {
  evm: {
    network: string;
    native: string;
    usdc: string;
  } | null;
  solana: {
    network: string;
    native: string;
    usdc: string;
  } | null;
}

export class AtlasDashboard {
  private db: any;
  private rpcUrls: { base: string; solana: string };

  constructor(config: {
    dbConnectionString: string;
    rpcUrls: { base: string; solana: string };
  }) {
    this.rpcUrls = config.rpcUrls;
  }

  async getBalances(
    evmAddress?: string,
    solanaAddress?: string
  ): Promise<BalanceData> {
    const balances: BalanceData = {
      evm: null,
      solana: null,
    };

    if (evmAddress) {
      const [eth, usdc] = await Promise.all([
        this.fetchETHBalance(evmAddress),
        this.fetchBaseUSDCBalance(evmAddress),
      ]);

      balances.evm = {
        network: 'base',
        native: eth,
        usdc,
      };
    }

    return balances;
  }

  private async fetchETHBalance(address: string): Promise<string> {
    const response = await fetch(this.rpcUrls.base, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        jsonrpc: '2.0',
        method: 'eth_getBalance',
        params: [address, 'latest'],
        id: 1,
      }),
    });

    const result = await response.json();
    if (result.result && result.result !== '0x') {
      const balance = BigInt(result.result);
      return (Number(balance) / 1e18).toFixed(6);
    }
    return '0.0';
  }

  private async fetchBaseUSDCBalance(address: string): Promise<string> {
    return '0.0';
  }
}

export default AtlasDashboard;



