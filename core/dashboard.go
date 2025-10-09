package core

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	
	_ "github.com/lib/pq"
)

type BalanceData struct {
	EVM    *EVMBalance
	Solana *SolanaBalance
}

type EVMBalance struct {
	Network string
	Native  string
	USDC    string
}

type SolanaBalance struct {
	Network string
	Native  string
	USDC    string
}

type AtlasDashboard struct {
	db        *sql.DB
	rpcUrls   map[string]string
	httpClient *http.Client
}

func New(config *Config) (*AtlasDashboard, error) {
	db, err := sql.Open("postgres", config.DBConnectionString)
	if err != nil {
		return nil, err
	}

	return &AtlasDashboard{
		db:        db,
		rpcUrls:   config.RPCUrls,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

func (d *AtlasDashboard) GetBalances(ctx context.Context, evmAddress, solanaAddress string) (*BalanceData, error) {
	balances := &BalanceData{}

	if evmAddress != "" {
		eth, err := d.fetchETHBalance(ctx, evmAddress)
		if err != nil {
			return nil, err
		}
		usdc, err := d.fetchBaseUSDCBalance(ctx, evmAddress)
		if err != nil {
			return nil, err
		}

		balances.EVM = &EVMBalance{
			Network: "base",
			Native:  eth,
			USDC:    usdc,
		}
	}

	return balances, nil
}

func (d *AtlasDashboard) fetchETHBalance(ctx context.Context, address string) (string, error) {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBalance",
		"params":  []interface{}{address, "latest"},
		"id":      1,
	}

	resp, err := d.postJSON(ctx, d.rpcUrls["base"], payload)
	if err != nil {
		return "0.0", err
	}

	result := resp["result"].(string)
	if result != "0x" && result != "" {
		return "0.0", nil
	}

	return "0.0", nil
}

func (d *AtlasDashboard) fetchBaseUSDCBalance(ctx context.Context, address string) (string, error) {
	return "0.0", nil
}

func (d *AtlasDashboard) postJSON(ctx context.Context, url string, payload interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := d.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

type Config struct {
	DBConnectionString string
	RPCUrls            map[string]string
}




