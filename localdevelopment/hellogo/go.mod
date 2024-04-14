module github.com/anishakd4/hellogo

go 1.22.2

//Note: ../mystrings means look in the parent directory of hellogo for the mystrings sibling directory.
replace github.com/anishakd4/mystrings v0.0.0 => ../mystrings
//We are telling go how to find this package in my system

require (
	github.com/anishakd4/mystrings v0.0.0
)
