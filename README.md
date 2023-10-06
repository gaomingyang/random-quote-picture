# random-quote-picture

the [setup.md](setup.md) shows instructions on how to run and test this application

本代码库用Golang和gin框架实现,由于功能比较简单只有一个简单的页面，就不做前后端分离了，直接后端全部实现。


## Tasks

- Create a small web app that displays an HTML page containing a random quote, and random picture from the two APIs below.
- Create a back-end CLI tool to do the same
- How can you return an image on cli?
- Add an option to both the CLI, and web app so a category (“key” in the API) can optionally be specified for the quote.  (sane/working default if not specified)
- Add an option to both the CLI, and web app to allow specifying a grayscale image
- Make it resilient to one or both providers being unavailable
- Provide unit tests
- Provide any documentation you think might be helpful
- Feel free to add anythng that you think might be useful


## APIs to be Used
- http://forismatic.com/en/api/
- https://picsum.photos/