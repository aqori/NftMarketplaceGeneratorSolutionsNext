// cmd/nftmarketplacegeneratorsolutionsnext/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplacegeneratorsolutionsnext/internal/nftmarketplacegeneratorsolutionsnext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplacegeneratorsolutionsnext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
