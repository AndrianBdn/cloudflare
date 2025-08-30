# E2E Tests for Cloudflare Provider

Tests the Cloudflare libdns provider against real Cloudflare API using the [libdns e2e framework](https://github.com/AndrianBdn/libdns/tree/master/e2e).

## Setup

1. **Get API Token**: See main README for token setup instructions. Test will use single or dual token depending on env variables.

2. **Set Environment Variables**:
```bash
export CLOUDFLARE_API_TOKEN="your-token-here"
export CLOUDFLARE_TEST_ZONE="yourzone.com."  # Include trailing dot
```

Or copy `.env.example` to `.env` and fill in values.

3. **Run Tests**

```bash
export $(cat .env | xargs) && go test -v
```

## What Gets Tested

- ListZones, GetRecords, AppendRecords, SetRecords, DeleteRecords
- Complete record lifecycle (create → update → delete)
- Various DNS record types

**Warning**: Tests create/delete real DNS records prefixed with "test-". Use a dedicated test zone or ensure you have backups.
