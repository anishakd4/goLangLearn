module github.com/anishakd4/datetest

go 1.22.2

require github.com/wagslane/go-tinytime v0.0.2 // indirect
//so go toolchain actually downloaded this code from github.com and added it to go.mod as a dependency and 
//create this new file called go.sum that kind of contains any transient dependencies or dependencies used by 
//the tinytime package that we imported

//go mod init github.com/anishakd4/datetest created this dataset package
// go get github.com/wagslane/go-tinytime pulled this module to this package