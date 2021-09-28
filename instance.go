package gmiddleware

import (
	"log"
	"os"
)

var mlog = log.New(os.Stdout, "[ZWWX]", log.LstdFlags)
