# 

# **PACKAGE NAMING**

## **NAMING CONVENTION**

By *convention*, a package's name is the same as the last element of its import path. For instance, the `math/rand` package comprises files that begin with:

`package rand`

That said, package names aren't *required* to match their import path. For example, I could write a new package with the path `github.com/mailio/rand` and name the package `random`:

`package random`

While the above is possible, it is discouraged for the sake of consistency.

## **ONE PACKAGE / DIRECTORY**

A directory of Go code can have **at most** one package. All `.go` files in a single directory must all belong to the same package. If they don't an error will be thrown by the compiler. This is true for main and library packages alike.



# 

# **MODULES**

Go programs are organized into *packages*. A package is a directory of Go code that's all compiled together. Functions, types, variables, and constants defined in one source file are visible to **all other source files within the same package (directory)**.

A *repository* contains one or more *modules*. A module is a collection of Go packages that are released together.

## **A GO REPOSITORY TYPICALLY CONTAINS ONLY ONE MODULE, LOCATED AT THE ROOT OF THE REPOSITORY.**

A file named `go.mod` at the root of a project declares the module. It contains:

- The module path
- The version of the Go language your project requires
- Optionally, any external package dependencies your project has

The module path is just the import path prefix for all packages within the module. Here's an example of a `go.mod` file:

`module github.com/bootdotdev/exampleproject

go 1.20

require github.com/google/examplepackage v1.3.0`

Each module's path not only serves as an import path prefix for the packages within but *also indicates where the go command should look to download it*. For example, to download the module `golang.org/x/tools`, the go command would consult the repository located at https://golang.org/x/tools.

> An "import path" is a string used to import a package. A package's import path is its module path joined with its subdirectory within the module. For example, the module github.com/google/go-cmp contains a package in the directory cmp/. That package's import path is github.com/google/go-cmp/cmp. Packages in the standard library do not have a module path prefix.
> 
- Paraphrased from Golang.org's [code organization](https://golang.org/doc/code#Organization)

## **DO I NEED TO PUT MY PACKAGE ON GITHUB?**

You don't *need* to publish your code to a remote repository before you can build it. A module can be defined locally without belonging to a repository. However, it's a good habit to keep a copy of all your projects on a remote server, like GitHub.



#

# BUILD

To build a production ready executable simply run go build, if the main package is in another folder run go build .\cmd\go-learning

If you want to build a executable for a machine with different os run the following command 
```
GOOS=target-OS GOARCH=target-architecture go build package-import-path
```



# 

# **GO INSTALL**

## **BUILD AN EXECUTABLE**

Ensure you are in your `hellogo` repo, then run:

`go install`

Navigate out of your project directory:

`cd ../`

Go has installed the `hellogo` program globally. Run it with:

`hellogo`

## **TIP ABOUT "NOT FOUND"**

If you get an error regarding "hellogo not found" it means you probably don't have your Go environment setup properly. Specifically, `go install` is adding your binary to your `GOBIN` directory, but that may not be in your `PATH`.

You can read more about that here in the [go install docs](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies).



# **CUSTOM PACKAGE**

Let's write a package to import and use in `hellogo`.

Create a sibling directory at the same level as the `hellogo` directory:

`mkdir mystrings
cd mystrings`

Initialize a module:

`go mod init {REMOTE}/{USERNAME}/mystrings`

Then create a new file `mystrings.go` in that directory and paste the following code:

`// by convention, we name our package the same as the directory
package mystrings

// Reverse reverses a string left to right
// Notice that we need to capitalize the first letter of the function
// If we don't then we won't be able to access this function outside of the
// mystrings package
func Reverse(s string) string {
  result := ""
  for _, v := range s {
    result = string(v) + result
  }
  return result
}`

Note that there is no `main.go` or `func main()` in this package.

`go build` won't build an executable from a library package. However, `go build` will still compile the package and save it to our local build cache. It's useful for checking for compile errors.

Run:

`go build`

# **GO INSTALL**

## **BUILD AN EXECUTABLE**

Ensure you are in your `hellogo` repo, then run:

`go install`

Navigate out of your project directory:

`cd ../`

Go has installed the `hellogo` program globally. Run it with:

`hellogo`

## **TIP ABOUT "NOT FOUND"**

If you get an error regarding "hellogo not found" it means you probably don't have your Go environment setup properly. Specifically, `go install` is adding your binary to your `GOBIN` directory, but that may not be in your `PATH`.

You can read more about that here in the [go install docs](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies).



# **CUSTOM PACKAGE CONTINUED**

Let's use our new `mystrings` package in `hellogo`

Modify hellogo's `main.go` file:

```
package main

import (
	"fmt"

	"{REMOTE}/{USERNAME}/mystrings"
)

func main() {
	fmt.Println(mystrings.Reverse("hello world"))
}
```

Don't forget to replace {REMOTE} and {USERNAME} with the values you used before. Then edit hellogo's `go.mod` file to contain the following:

```
module example.com/username/hellogo

go 1.20

replace example.com/username/mystrings v0.0.0 => ../mystrings

require (
	example.com/username/mystrings v0.0.0
)
```

Now build and run the new program:

`go build
./hellogo`

## Replace
The replace command is used to tell the dependecy manager that the dependecy is not on github but instead on my machine.



# **CLEAN PACKAGES**

I’ve often seen, and have been responsible for, throwing code into packages without much thought. I’ve quickly drawn a line in the sand and started putting code into different folders (which in Go are different packages by definition) just for the sake of findability. Learning to properly build small and reusable packages can take your Go career to the next level.

## **RULES OF THUMB**

### **1. HIDE INTERNAL LOGIC**

If you're familiar with the pillars of OOP, this is a practice in *encapsulation*.

Oftentimes an application will have complex logic that requires a lot of code. In almost every case the logic that the application cares about can be exposed via an API, and most of the dirty work can be kept within a package. For example, imagine we are building an application that needs to classify images. We could build a package:

`package classifier

// ClassifyImage classifies images as "hotdog" or "not hotdog"
func ClassifyImage(image []byte) (imageType string) {
	if hasHotdogColors(image) && hasHotdogShape(image) {
		return "hotdog"
	} else {
		return "not hotdog"
	}
}

func hasHotdogShape(image []byte) bool {
	// internal logic that the application doesn't need to know about
	return true
}

func hasHotdogColors(image []byte) bool {
	// internal logic that the application doesn't need to know about
	return true
}`

We create an API by only exposing the function(s) that the application-level needs to know about. All other logic is unexported to keep a clean separation of concerns. The application doesn’t need to know how to classify an image, just the result of the classification.

### **2. DON’T CHANGE APIS**

The unexported functions within a package can and should change often for testing, refactoring, and bug fixing.

A well-designed library will have a stable API so that users aren’t receiving breaking changes each time they update the package version. In Go, this means not changing exported function’s signatures.

### **3. DON’T EXPORT FUNCTIONS FROM THE MAIN PACKAGE**

A `main` package isn't a library, there's no need to export functions from it.

### **4. PACKAGES SHOULDN'T KNOW ABOUT DEPENDENTS**

Perhaps one of the most important and most broken rules is that a package shouldn’t know anything about its dependents. In other words, a package should never have specific knowledge about a particular application that uses it.

## **FURTHER READING**

You can optionally [read more here](https://blog.boot.dev/golang/how-to-separate-library-packages-in-go/) if you're interested.