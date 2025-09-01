# Atlas Dashboard

> User analytics and activity tracking SDK for x402 ecosystem

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![x402](https://img.shields.io/badge/x402-Compatible-green)](https://x402.org)

Atlas Dashboard provides comprehensive analytics for users interacting with the x402 ecosystem.

## Installation

### TypeScript/JavaScript

```bash
npm install @atlas402/dashboard
```

### Python

```bash
pip install atlas-dashboard
```

### Go

```bash
go get github.com/atlas402/dashboard
```

### Java

```xml
<dependency>
  <groupId>com.atlas402</groupId>
  <artifactId>dashboard</artifactId>
  <version>1.0.0</version>
</dependency>
```

## Quick Start

### TypeScript

```typescript
import { AtlasDashboard } from '@atlas402/dashboard';

const dashboard = new AtlasDashboard({
  dbConnectionString: process.env.DATABASE_URL,
  rpcUrls: {
    base: 'https://mainnet.base.org',
    solana: 'https://api.mainnet-beta.solana.com',
  },
});

const balances = await dashboard.getBalances('0x...', 'SolanaAddress...');
```

## License

Apache 2.0
