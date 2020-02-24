An RSS reader made with the Go.

Main features:
- Parse the RSS items of one or several RSS feeds
- Sort and display RSS items per channel
- Save JSON file from each output

### Installation

1. In case of Go not yet installed on your machine, follow the instructions here:  
https://golang.org/doc/install

2. Get the repo, either as git or using `go get`:
    ```
    $ go get -u github.com/batyanko/rssreader/
    ```
3. Run `make setup` to download linter and unit test dependencies

### Usage

1. Open a console and change to the `cmd/reader_app` dir in the project

2. Build the binary and run it (following example for unix-style terminal):
    ```
    $ go build
    $ ./reader_app 
    ```

Example output:
```
-----------------------
RSS items in channel Liftoff News:
-----------------------

Title: Star City
Source: Liftoff News
Source URL: http://liftoff.msfc.nasa.gov/
Link: http://liftoff.msfc.nasa.gov/news/2003/news-starcity.asp
Publish Date: 2003-06-03 09:39:21 +0000 GMT
Description: How do Americans get ready to work with Russians aboard the International Space Station? They take a crash course in culture, language and protocol at Russia's <a href="http://howe.iki.rssi.ru/GCTC/gctc_e.htm">Star City</a>.

Title: 
Source: Liftoff News
Source URL: http://liftoff.msfc.nasa.gov/
Link: 
Publish Date: 2003-05-30 11:06:42 +0000 GMT
Description: Sky watchers in Europe, Asia, and parts of Alaska and Canada will experience a <a href="http://science.nasa.gov/headlines/y2003/30may_solareclipse.htm">partial eclipse of the Sun</a> on Saturday, May 31st.

...
```

With each run of the application, a JSON file is created, following a `rss_items_<timestamp>.json` format.  


### License

This project is licensed under the Apache Version 2.0 license.  
http://www.apache.org/licenses/LICENSE-2.0