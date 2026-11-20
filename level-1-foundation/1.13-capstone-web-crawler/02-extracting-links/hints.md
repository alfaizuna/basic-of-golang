# Hints

### Hint 1
Parse base: `parsedBase, err := url.Parse(baseStr)`.

### Hint 2
Token tag name: `token := tokenizer.Token()`; cek `token.Data == "a"`.

### Hint 3
Loop `token.Attr` untuk mencari `attr.Key == "href"` lalu lakukan `parsedBase.ResolveReference`.
