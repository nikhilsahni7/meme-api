Meme API

This is a Go application that provides a RESTful API for fetching memes from the imgflip API and creating new memes. The application stores the created memes in an in-memory data structure.
Prerequisites

    Go (version 1.16 or later)

Getting Started

    Clone the repository:

git clone https://github.com/nikhilsahni7/meme-api.git

Change to the project directory:

cd meme-api

Build and run the application:

    go run main.go or download air and then run the application using air command

    The server will start running on http://localhost:8080.

Environment Variables

The application uses the following environment variable:

    PORT: The port number on which the server will listen for incoming requests. If not set, the default port 8080 will be used.

You can set this environment variable in a .env file in the project root directory, or export it in your shell.
Endpoints
Health Check

    URL: /check
    Method: GET
    Description: Checks if the server is running and responds with a simple "OK" message.

Get Memes

    URL: /memes
    Method: GET
    Description: Fetches memes from the imgflip API and returns a JSON array containing the names and URLs of the memes.

Create Meme

    URL: /memes
    Method: POST
    Description: Creates a new meme. The meme data should be sent in the request body as JSON.

Get Created Memes

    URL: /created-memes
    Method: GET
    Description: Retrieves the list of memes created through the POST /memes endpoint.

Request Body Format

The POST /memes endpoint expects the request body to be in the following JSON format:

json

{
"id": "string",
"name": "string",
"url": "string",
"width": number,
"height": number,
"box_count": number,
"captions": number
}

Dependencies

This project uses the following third-party dependencies:

    Gin - A HTTP web framework for Go.

These dependencies are automatically installed when you build the project.

This project is licensed under the MIT License.
