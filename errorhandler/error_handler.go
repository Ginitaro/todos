package errorhandler

import (
	"fmt"
	"github.com/juju/errors"
	"log"
	"todos/env"
)

const None = 0
const Default = 1
const Strict = 2

func CatchError(err error, msg string) {

	params := env.GetEnvironmentVariable()

	ErrorMode := params.ErrorMode

	// Check if there was an error
	if err != nil {

		switch ErrorMode {
		// Ignore every error
		case None:
			break
		default:

			// Show error message
			if ErrorMode >= Default {

				errorMsg := params.ErrorMessage
				if len(msg) != 0 {
					errorMsg = msg
				}

				fmt.Println(errorMsg)
			}

			// Kill process and show error
			if ErrorMode == Strict {
				log.Fatal(errors.Trace(err))
			}
			break
		}
	}
}
