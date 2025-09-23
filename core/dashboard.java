package com.atlas402.dashboard.core;

import java.util.concurrent.CompletableFuture;

public class AtlasDashboard {
    private final Config config;
    
    public AtlasDashboard(Config config) {
        this.config = config;
    }
    
    public CompletableFuture<BalanceData> getBalances(String evmAddress, String solanaAddress) {
        return CompletableFuture.supplyAsync(() -> {
            return new BalanceData();
        });
    }
}



