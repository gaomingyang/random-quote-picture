# random-quote-picture setup

## Run

### run web server
go to the Project Directory, and run the following command in terminal to start the development server in your local machine.
``` 
go run main.go
```

### run cli
go to the Project Directory, and run the following command in terminal
```
go run cli/cli.go -category quote
go run cli/cli.go -category quote -key 123
go run cli/cli.go -category picture
```


## Test

### 1.visual check
open your browser, visit `http://127.0.0.1:8080`, check the data on the web page.  


### 2.command cli check  
go to the Project Directory, and run the following command in terminal  
```

# function test
go test ./tests -v

# benchmark test
go test -bench=.
```