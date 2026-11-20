# Hints

### Hint 1
Di method Touch(key): `l.clients[key] = &ClientInfo{lastSeen: l.now()}`.

### Hint 2
Di method Cleanup(maxIdle): loop `for key, info := range l.clients { if cur.Sub(info.lastSeen) > maxIdle { delete(l.clients, key); count++ } }`.
