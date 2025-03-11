# Movie API

## Description

The Movie API is a simple RESTful API built with Golang that allows users to manage a collection of movies. It provides functionality to create, read, update, and delete movie records. The API uses the Gorilla Mux package for routing and JSON for data exchange. This project is ideal for understanding the basics of RESTful API development in Go.

## Functionalities

- Retrieve all movies
- Retrieve a specific movie by ID
- Create a new movie
- Update an existing movie
- Delete a movie

## Technologies Used

- Golang
- Gorilla Mux
- JSON for data exchange
- HTTP server

## How to Run the Project

### Prerequisites

Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/dl/).

### Steps to Run the Project

1. Clone the repository:
   ```sh
   git clone <repository_url>
   cd movie-api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the server:
   ```sh
   go run main.go
   ```
4. The server will start on `http://localhost:8000`

## What This Code Does

- **Defines Movie and Director structures**: The `Movie` struct represents a movie with attributes such as ID, ISBN, Title, and Director details.
- **Preloads sample movie data**: The application initializes with two predefined movies.
- **Implements API endpoints**:
  - `GET /movies` - Fetches all movies.
  - `GET /movies/{id}` - Retrieves a specific movie by ID.
  - `POST /movies` - Creates a new movie with a randomly generated ID.
  - `PUT /movies/{id}` - Updates an existing movie while preserving its ID.
  - `DELETE /movies/{id}` - Deletes a specific movie by ID.
- **Uses Gorilla Mux for routing**: The API endpoints are mapped to handler functions using the Mux router.
- **Handles JSON encoding and decoding**: The API efficiently parses and responds with JSON-formatted data.

## API Endpoints

| Method | Endpoint       | Description              |
| ------ | -------------- | ------------------------ |
| GET    | `/movies`      | Get all movies           |
| GET    | `/movies/{id}` | Get a specific movie     |
| POST   | `/movies`      | Create a new movie       |
| PUT    | `/movies/{id}` | Update an existing movie |
| DELETE | `/movies/{id}` | Delete a movie           |

## Example Request/Response

### Get All Movies

**Request:**

```sh
GET /movies
```

**Response:**

```json
[
  {
    "id": "1",
    "isbn": "438227",
    "title": "Movie One",
    "director": {
      "firstname": "Muhammad",
      "lastname": "Aman"
    }
  },
  {
    "id": "2",
    "isbn": "45455",
    "title": "Movie Two",
    "director": {
      "firstname": "Muhammad",
      "lastname": "Zain"
    }
  }
]
```

## Contributing

If you wish to contribute, feel free to fork the repository and submit a pull request.

## License

This project is licensed under the MIT License.

