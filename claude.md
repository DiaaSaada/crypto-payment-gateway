You are architecting a USDT stablecoin payment gateway written in GoLang. Before writing any code, I need you to produce a comprehensive high-level plan and save it as `PLAN.md` in the project root.

## Product Overview

A self-hosted, non-custodial USDT payment gateway that:
- Generates a unique deposit address per purchase order (HD wallet / BIP-44 derivation)
- Monitors the blockchain for incoming USDT deposits on the selected network
- Once the expected amount is confirmed, automatically sweeps the funds to the merchant's configured cold wallet address
- Exposes a REST API for merchant integration (create order, check status, webhooks)
- Supports multiple USDT networks: Tron (TRC-20), Ethereum (ERC-20), BSC (BEP-20), and Polygon — start with Tron TRC-20 as primary since it has the lowest fees

## What the plan must cover

Structure the PLAN.md with these sections:

### 1. System Architecture
- High-level component diagram (described in text/mermaid)
- Service boundaries: API server, blockchain watchers, sweep engine, webhook dispatcher
- Data flow: order creation → address generation → deposit detection → confirmation → sweep → merchant notification

### 2. Tech Stack & Dependencies
- Go libraries for each blockchain (e.g., go-ethereum, gotron-sdk, etc.)
- Database choice (PostgreSQL recommended) and schema outline (orders, wallets, transactions, merchants tables)
- Message queue / job system for sweep jobs and webhook retries (e.g., Redis + asynq or NATS)
- Configuration management and secrets handling (HD master seed, RPC endpoints)

### 3. HD Wallet & Key Management
- BIP-44 derivation strategy: how derivation paths map to merchants and orders
- Master seed storage (encrypted at rest, loaded into memory only)
- Address generation flow: derive child key → compute address per network
- Security considerations: seed never exposed via API, key material in memory only

### 4. Blockchain Monitoring Service
- Per-chain watcher design: polling vs websocket vs hybrid
- How to detect incoming USDT (TRC-20 TriggerSmartContract events, ERC-20 Transfer events, etc.)
- Confirmation thresholds per chain (e.g., 19 blocks for Tron, 12 for Ethereum, 15 for BSC)
- Handling edge cases: underpayment, overpayment, multiple partial payments, expired orders
- RPC node strategy: multiple providers for redundancy, fallback logic

### 5. Sweep Engine
- Two-step sweep for ERC-20 tokens: fund gas → transfer tokens (not needed for Tron)
- Gas estimation and fee management per chain
- Batching strategy: sweep immediately vs batch on schedule
- Retry logic for failed sweeps
- Dust threshold: minimum amount worth sweeping

### 6. Merchant REST API Design
- Endpoints: POST /orders, GET /orders/:id, GET /orders/:id/status, POST /merchants (admin)
- Webhook payload structure and HMAC signing for verification
- Webhook retry policy (exponential backoff, max retries, dead letter)
- Authentication: API key + secret per merchant
- Rate limiting strategy

### 7. Security Plan
- Threat model: compromised RPC, stolen DB, leaked API keys, insider threat
- Input validation and address verification
- Rate limiting on address generation to prevent resource exhaustion
- Audit logging for all sensitive operations
- Network-level: only cold wallet address is configurable, no withdrawal API

### 8. Project Structure
- Propose a clean Go project layout following standard Go project conventions
- cmd/, internal/, pkg/ structure
- Clear separation: domain logic, blockchain adapters, API handlers, background workers

### 9. MVP Scope vs Future
- What's in MVP: single merchant, Tron TRC-20 only, basic webhook, PostgreSQL, single binary
- Phase 2: multi-chain (ERC-20, BEP-20), multi-merchant, admin dashboard
- Phase 3: smart contract deposit addresses (CREATE2), gas optimization, analytics

### 10. Deployment & Ops
- Dockerized single binary with docker-compose (app + postgres + redis)
- Environment variable configuration
- Health check endpoints
- Logging strategy (structured JSON logs)
- Metrics: orders created, deposits detected, sweeps completed, sweep failures

Be thorough, opinionated, and practical. Prioritize security throughout. Where there are trade-offs, state them clearly with your recommendation and reasoning. Include rough estimates of complexity where helpful.

Save everything to PLAN.md in the project root.