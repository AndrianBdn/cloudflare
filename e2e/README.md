# Tests for Cloudflare Provider

Tests the Cloudflare libdns provider against real Cloudflare API using the official [libdnstest framework](https://github.com/libdns/libdns/tree/master/libdnstest).

## How To Run

1. **Get API Token and setup zone**: See main README for token setup instructions. Test will use single or dual token depending on env variables. Setup some test Cloudflare zone.

2. **Set Environment Variables**:
```bash
export CLOUDFLARE_API_TOKEN="your-token-here"
export CLOUDFLARE_TEST_ZONE="example.org."  # Include trailing dot
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
