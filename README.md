<p align="center">
  <h3 align="center">ShortyResty</h3>

  <p align="center">
    This is a REST API that accepts URL links (in JSON) and shortens them into smaller URL with random IDs.
  </p>
</p>

## About The Project

We've all been there, trying to type out needlessly long URL in order to get to where we want. This project was meant to solve this issue by shortening links and allowing visitors of shortened links to be redirected to the original URL.

### Built/Tested With
* [Go](https://golang.org/)
* [Insomnia](https://insomnia.rest/)

## Getting Started

### Prerequisites

Please install the latest version of go before running this code.

Download Go using Go installer: https://golang.org/doc/install

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/RichardLin7/ShortyResty.git
   ```
2. Run main.go 

3. Send POST request to http://127.0.0.1:8080/ in JSON using this format:
    ```sh
    {"url":  "https://www.exampleurl.com/somethingthatyoudon'twanttotype"}
    ```

4. Use the recieved short_url in sent back on your chosen browser (or send GET request to short_url).
