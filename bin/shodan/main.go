package main

import (
	"github.com/acuteaura/shodan"

	// module imports
	_ "github.com/acuteaura/shodan/modules/core"
	_ "github.com/acuteaura/shodan/modules/gmtools"
	_ "github.com/acuteaura/shodan/modules/greet"
	_ "github.com/acuteaura/shodan/modules/log"
	_ "github.com/acuteaura/shodan/modules/oauth"
	_ "github.com/acuteaura/shodan/modules/webui"
)

func main() {
	shodan.Run()
}
