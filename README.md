# random-quote-picture

the [setup.md](setup.md) shows instructions on how to run and test this application

This project is implemented using Golang and the Gin framework. Due to its simplicity with only a single basic page, we did not opt for frontend-backend separate to different repository. Both the page and the API are provided through the backend web service.

## Tasks

- Create a small web app that displays an HTML page containing a random quote, and random picture from the two APIs below.
- Create a back-end CLI tool to do the same
- How can you return an image on cli?
- Add an option to both the CLI, and web app so a category (“key” in the API) can optionally be specified for the quote.  (sane/working default if not specified)
- Add an option to both the CLI, and web app to allow specifying a grayscale image
- Make it resilient to one or both providers being unavailable
- Provide unit tests
- Provide any documentation you think might be helpful
- Feel free to add anything that you think might be useful


## APIs to be Used
- http://forismatic.com/en/api/
- https://picsum.photos/