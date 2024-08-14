#!/bin/bash
curl -XPOST --url 'http://localhost:8332' --user 'abcrpcuser:password' --data-binary '{"jsonrpc": "1.0", "id":"1", "method": "stop", "params": [] }'