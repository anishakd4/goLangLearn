package main

//Interesting thing about main package is that it will always have a main function, it serves as entrypoint to the program.
//Only main package will have a main function which runs when the program starts

//standard library is made up of library packages
import (
	"fmt"

	"github.com/anishakd4/mystrings"
)

/*
PACKAGES
Every Go program is made up of packages.

You have probably noticed the package main at the top of all the programs you have been writing.

A package named "main" has an entrypoint at the main() function. A main package is compiled into an executable program.

A package by any other name is a "library package". Libraries have no entry point. Libraries simply export functionality
that can be used by other packages.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

This program is an executable. It is a "main" package and imports from the fmt and math/rand library packages.

*/

/*
PACKAGE NAMING

NAMING CONVENTION
By convention, a package's name is the same as the last element of its import path. For instance, the
math/rand package comprises files that begin with:

package rand

That said, package names aren't required to match their import path. For example, I could write a new package
with the path github.com/mailio/rand and name the package random:

package random

While the above is possible, it is discouraged for the sake of consistency.

ONE PACKAGE / DIRECTORY
A directory of Go code can have at most one package. All .go files in a single directory must all belong to the
same package. If they don't an error will be thrown by the compiler. This is true for main and library packages alike.
*/

/*
MODULES

Go programs are organized into packages. A package is a directory of Go code that's all compiled together. Functions,
types, variables, and constants defined in one source file are visible to all other source files within
the same package (directory).

A repository contains one or more modules. A module is a collection of Go packages that are released together.

A GO REPOSITORY TYPICALLY CONTAINS ONLY ONE MODULE, LOCATED AT THE ROOT OF THE REPOSITORY.
A file named go.mod at the root of a project declares the module. It contains:

1. The module path
2. The version of the Go language your project requires
3. Optionally, any external package dependencies your project has

The module path is just the import path prefix for all packages within the module. Here's an example of a go.mod file:

module github.com/bootdotdev/exampleproject
go 1.22.1
require github.com/google/examplepackage v1.3.0

Each module's path not only serves as an import path prefix for the packages within but also indicates where the go
command should look to download it.

An "import path" is a string used to import a package. A package's import path is its module path joined with its
subdirectory within the module. For example, the module github.com/google/go-cmp contains a package in the directory
cmp/. That package's import path is github.com/google/go-cmp/cmp. Packages in the standard library do not have a
module path prefix.

*/

/*
The go run command is used to quickly compile and run a Go package. The compiled binary is not saved in your working
directory. Use go build instead to compile production executables.

I rarely use go run other than to quickly do some testing or debugging.

Conventionally, the file in the main package that contains the main() function is called main.go.
*/

/*
GO BUILD

go build compiles go code into an executable program

BUILD AN EXECUTABLE
Ensure you are in your hellogo repo, then run:

go build

Run the new program:

./hellogo

We can run this executable in a new machine and run it without even needing to install to go tool chain in that machine, this
is a compiled machine code

*/

/*
GO INSTALL

BUILD AN EXECUTABLE
Ensure you are in your hellogo repo, then run:

go install

Navigate out of your project directory:

cd ../

Go has compiled and installed the hellogo program globally. Run it with:

hellogo

TIP ABOUT "NOT FOUND"
If you get an error regarding "hellogo not found" it means you probably don't have your Go environment setup properly.
Specifically, go install is adding your binary to your GOBIN directory, but that may not be in your PATH.

*/
func main() {
	//fmt.Println("hello world")
	fmt.Println(mystrings.Reverse("hello world"))
}
