package mediachannel

// Status ...
var Status = make(chan string)

// Progress ...
var Progress = make(chan float32)

// PlayList ...
var PlayList = make(chan string)
