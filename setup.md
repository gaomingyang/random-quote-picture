# random-quote-picture setup

## Run

### run web server
go to the Project Directory, and run the following command in terminal to start the development server in your local machine.
``` 
go run main.go
```
Then open your browser and visit `http://127.0.0.1:8080`.

### run cli
Open your terminal app, go to the Project Directory, and run the following command below. After run one line command ,you'll see the output result.
```
go run cli/cli.go -category quote
go run cli/cli.go -category quote -key 123
go run cli/cli.go -category picture
go run cli/cli.go -category=picture -grayscale=1
```


## Test

### 1.test by web server page
open your browser, visit `http://127.0.0.1:8080`, check the content and functions on the web page.  

### 2.test by command cli  
go to the Project Directory, and run the following command in terminal  

function test  
```
go test ./tests 
or
go test ./tests -v
```

benchmark test  
```
go test ./tests -bench=.
```

### 3.The case when quote provider (forismatic) is unvailable.

To test this case, you can do follow this way:

(1) start web server if it is not running
```
go run main.go 
```

(2) visit the webpage http://127.0.0.1:8080 and click the button "Change Quote" several times to get and cache several quotes.

(3) modify the conf/app.yaml file，change the value of quoteServerUrl from 'http://api.forismatic.com/api/1.0/' to a wrong value, such as 'http://api123.forismatic.com/api/1.0/'

(4) continue to click the button "Change Quote" on the webpage, you will see it still works well. But when you look at the terminal, there will be some error logs show that "get quote from forismatic error"


### 4.The case when picture privider (Picsum) is unvailable.

the test step is similar with the case above. 

(1) start web server if it is not running
```
go run main.go 
```

(2) visit the webpage http://127.0.0.1:8080 and click the button "Change Picture", and currently the function is right.

(3) modify the conf/app.yaml file，change the value of pictureServerUrl from 'https://picsum.photos' to a wrong value, such as 'https://picsum123.photos'. So that the webpage will get a wrong image url which can't be opened.

(4) continue to click the button "Change Picture" on the webpage, you will see the backup pictures. 