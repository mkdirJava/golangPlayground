
# Project demoing usage of Golang

## Aim
* To take a look at Goalng and see how it does certain principles
    1. Object Orientation
    2. Functional
    3. Tests
* To look at its module dependecy management
* To look at debugging an application 
---
## Pleasant Surprises
GOLANG SUPPORTS GENERICS !!!

---

## Getting started

Requirements
1. Have Golang 1.18 installed
2. Working with VsCode


---
### Opening the Project in VsCode
1. Double click on 
        
        <project root>/projectStructure.code-workspace
---
### Building a Library

1. Go to ./ObjectOrientated
    * Run 

            go build ./... 
    * This should build the binaries and put it on the Go directory ready for the Importing project/ Module (directory) to use
---
### Run tests

    1. Go to ./ObjectOrientated
        * Run
            go test ./... -v  
---
### Importing another local project 

1. Go to ./Importing 
    * Run
        
            go build  
    * There should be an executeable that is build and able to run ( nativley )

    * Do the following Curl request you should get somthing back

            curl localhost:8080
---
### Debugging in vscode

1. Install delve
    * Run
        ''' go install github.com/go-delve/delve/cmd/dlv@latest '''
        * This will bring the binary into local 
        * Should see somthing when running

                dlv version
    * Build the artifact in Importing ( it has the main )
        1. Run in the Importing directory

                go build -gcflags=all="-N -l"

        * The flags are to prevent optimisations and inlining, if left out it make debugging difficult 

    * Start the application in debug mode     
        
        1. Run the application with delve 
                
                 dlv --listen=:2345 --headless=true --api-version=2 exec ./Importing/importing.exe 
        2. Back to VsCode and click on debugging, configurations are already setup in 
                
                <projectRoot>/.vscode/

            * Add a break point at 
                    
                    <project root>/Importing/main.go
            * run 
                    
                    Connect to server (root)
            * Do some curl requests and see it debug

---
## Noticeable Qerks

* Vscode needs to recognise the modules as workspaces in their own right hence there is projectStructure.code-workspace file at the project root

* Seems that in Golang, 
    * Panics are unchecked execptions
        * Looks like defering a function and checking the status is a way to recover from an unchecked exception 
    * The checked exceptions you pass back in a tuple for the developer to handle


## Rough State Machine
There is a rough state machine localed in 
        
        <project root>/SearchForTruth

Tests are located in 

        <project root>/SearchForTruth/models/Model_test.go

There are tests demonstrating and asserting the behaviour