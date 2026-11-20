# Importing Packages & Aliases

Mengimpor package internal dan memanfaatkan alias import saat terjadi benturan nama atau untuk mempermudah pembacaan kode.

## Contoh

```go
//go:build ignore

package main

import (
	"fmt"
	crand "crypto/rand"
	mrand "math/rand"
)

func main() {
	fmt.Println("Alias import mencegah benturan nama antar package")
	_ = crand.Reader
	_ = mrand.Int()
}

```

## Penjelasan

1. **Import Path Lengkap**: Saat mengimpor subpackage proyek sendiri, gunakan path lengkap yang diawali dengan nama module:
   `import "learn-go/pkg/mathutil"`
2. **Import Alias**: Jika dua package memiliki nama akhir yang sama (misal `crypto/rand` dan `math/rand`), berikan alias di depan path import:
   `import crand "crypto/rand"`
3. **Blank Import `_`**: `import _ "package"` hanya menjalankan fungsi `init()` dari package tersebut tanpa mengimpor pengenal fungsinya (sering dipakai untuk driver database).

## >_ Exercise

Buat fungsi `ResolveNamespace(alias, originalPackage string) string` yang:
- Jika `alias` tidak kosong, mengembalikan `"using alias <alias> for <originalPackage>"`
- Jika `alias` kosong `""`, mengembalikan `"using default package <originalPackage>"`

### Expected output

```
using alias crand for crypto/rand
using default package strings
```

## Hints (0/3)

Ketik `/hint` untuk petunjuk pertama jika Anda memerlukan bantuan.

## Key Takeaway

- Berikan alias pada impor jika terdapat benturan nama package.
- Subpackage dalam module diimpor menggunakan awalan module path.
