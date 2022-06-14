package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		upperCaseLine := strings.ToUpper(scanner.Text())
		fmt.Println(upperCaseLine)
	}

	if err := scanner.Err(); err != nil {
		log.Panicf("error: %v\n", err)
	}
}

/*
Example of use:

./out/64-line-filters <<- EOF
what'sup
doc
howyoudoin?
EOF

output:

WHAT'SUP
DOC
HOWYOUDOIN?

*/
