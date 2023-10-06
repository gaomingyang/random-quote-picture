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
go run cli/cli.go -category=picture -grayscale=1
```


## Test

### 1.test by web server page
open your browser, visit `http://127.0.0.1:8080`, check the data on the web page.  

### 2.test by command cli  
go to the Project Directory, and run the following command in terminal  
```

# function test
go test ./tests -v

# benchmark test
go test -bench=.
```

### 3.The case when quote server (forismatic) is unvailable.

As this is just a demo project. To reduce the dependencies of this project, I stored the cached quote data directly in program memory. In a production environment, we can store the data in better storage solutions like Redis, MongoDB, or even MySQL to enhance availability.

This program will cache 100 quotes which were got from forismatic. When forismatic service is unvailable, it will get quote from cache.
To test this case, you can do follow this way:

(1) start web server if it is not running
```
go run main.go 
```

(2) visit the webpage http://127.0.0.1:8080 and click the button "Change Quote" several times to get and cache quotes.

(3) modify the conf/app.yaml file，change the value of quoteServerUrl from 'http://api.forismatic.com/api/1.0/' to other value, such as 'http://api123.forismatic.com/api/1.0/'

(4) continue to click the button "Change Quote" on the webpage, you will see it still works,but you can look at the terminal. There will be some error logs show that "get quote from forismatic error"

