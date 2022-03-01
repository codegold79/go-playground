package main

import (
	"errors"
	"fmt"

	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

var errOne = errors.New("the first problem")
var errTwo = errors.New("the second problem")
var errThree = errors.New("the third problem")

func main() {
	var errs []error

	errs = append(errs, errOne)
	errs = append(errs, errTwo)

	kerr := kerrors.NewAggregate(errs)

	fmt.Println("does errOne exist?", errors.Is(kerr, errOne))
	fmt.Println("does errTwo exist?", errors.Is(kerr, errTwo))
	fmt.Println("does errThree exist?", errors.Is(kerr, errThree))
}
