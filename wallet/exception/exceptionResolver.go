package exception

import "fmt"

func ResolveWithoutPanic(tryCb func() (any, error), finallyCb func() any) any {
	tryResult, err := tryCb()
	if err != nil {
		fmt.Print(err)
	}
	finallyCb()
	return tryResult
}
