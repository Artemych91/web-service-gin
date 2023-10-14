# Web Service with Gin - Album API

This is a simple web service implemented in Go using the Gin web framework, providing an API to manage music albums.

## Project Structure

The project is structured as follows:

- `main.go`: The main entry point of the application, setting up the HTTP server and defining the API endpoints.
- `database/db.go`: Contains functions to initialize and manage the database connection.
- `handlers/album_handlers.go`: Defines HTTP request handlers for album-related operations.
- `models/album.go`: Contains the definition of the `album` struct.

## Getting Started

### Prerequisites

- Go installed on your machine ([Install Go](https://golang.org/doc/install))

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Artemych91/web-service-gin.git
   cd web-service-gin

2. Run the application:

   ```bash
   go run main.go

3. The application will start and listen on localhost:8080.

## API Endpoints

### Get All Albums

- **URL:** `/albums`
- **Method:** `GET`
- **Response:** List of all albums in JSON format.

### Get Album by ID

- **URL:** `/albums/:id`
- **Method:** `GET`
- **Parameters:**
  - `id`: ID of the album to retrieve.
- **Response:** Details of the specified album in JSON format.

### Create a New Album

- **URL:** `/albums`
- **Method:** `POST`
- **Request Body:** JSON payload containing album details (title, artist, price).
- **Response:** Details of the newly created album in JSON format.

## Contributing

Feel free to contribute by opening issues or submitting pull requests. Any contributions you make are greatly appreciated!

## License

This project is licensed under the [MIT License](LICENSE).
