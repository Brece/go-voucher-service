# go-voucher-service

This service generates a pdf voucher with the given booking data.

### Setup

-   Run

```bash
go run .
```

-   Generate voucher for testing

```bash
curl -X POST http://localhost:8080/voucher -o voucher.pdf
```
