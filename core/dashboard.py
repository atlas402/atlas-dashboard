from typing import Optional, Dict, Any
from dataclasses import dataclass
import asyncpg
import aiohttp

@dataclass
class BalanceData:
    evm: Optional[Dict[str, str]] = None
    solana: Optional[Dict[str, str]] = None

class AtlasDashboard:
    def __init__(
        self,
        db_connection_string: str,
        rpc_urls: Dict[str, str]
    ):
        self.db_connection_string = db_connection_string
        self.rpc_urls = rpc_urls
        
    async def get_balances(
        self,
        evm_address: Optional[str] = None,
        solana_address: Optional[str] = None
    ) -> BalanceData:
        balances = BalanceData()
        
        if evm_address:
            eth_balance = await self._fetch_eth_balance(evm_address)
            usdc_balance = await self._fetch_base_usdc_balance(evm_address)
            balances.evm = {
                'network': 'base',
                'native': eth_balance,
                'usdc': usdc_balance,
            }
        
        return balances
    
    async def _fetch_eth_balance(self, address: str) -> str:
        async with aiohttp.ClientSession() as session:
            async with session.post(
                self.rpc_urls['base'],
                json={
                    'jsonrpc': '2.0',
                    'method': 'eth_getBalance',
                    'params': [address, 'latest'],
                    'id': 1,
                }
            ) as response:
                result = await response.json()
                if result.get('result') and result['result'] != '0x':
                    balance = int(result['result'], 16)
                    return f"{balance / 1e18:.6f}"
        return '0.0'
    
    async def _fetch_base_usdc_balance(self, address: str) -> str:
        return '0.0'


