# NIK Validator

This package provides a simple NIK (Nomor Induk Kependudukan) validator for validating Indonesian national identification numbers.

### Usage
```go
package main

import (
	"fmt"
	"github.com/klepon46/nik-validator/pkg"
)

func main() {
	validator := pkg.NewValidator()

	nik := "your_nik_here"

	err := validator.Validate(nik)
	if err != nil {
		fmt.Println("NIK validation failed:", err)
		return
	}

	fmt.Println("NIK is valid")
}
```

### Installation
```bash
go get github.com/klepon46/nik-validator
```

### License
This package is licensed under the MIT license. See the [LICENSE](LICENSE) file for more information.

### Contributing
Contributions are welcome! Please feel free to submit a pull request.
