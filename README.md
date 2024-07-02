# Tic Tac Toe Game

A simple Tic Tac Toe game with a Go backend and React frontend.

## Project Structure

tic-tac-toe/
├── backend/
│ ├── main.go
│ ├── go.mod
│ └── go.sum
└── frontend/
├── src/
│ ├── App.js
│ ├── App.css
│ ├── Board.js
│ ├── Board.css
│ ├── Square.js
│ └── Square.css
├── package.json
└── public/
└── index.html


## Prerequisites

- [Go](https://golang.org/doc/install)
- [Node.js](https://nodejs.org/)
- [npm](https://www.npmjs.com/get-npm) or [yarn](https://yarnpkg.com/)

## Setup

### Backend

1. Navigate to the `backend` directory:

    ```sh
    cd backend
    ```

2. Initialize the Go module (if not already done):

    ```sh
    go mod init tic-tac-toe
    ```

3. Install the necessary packages:

    ```sh
    go get -u github.com/gorilla/mux
    go get -u github.com/rs/cors
    ```

4. Start the backend server:

    ```sh
    go run main.go
    ```

    The backend server will start at `http://localhost:8080`.

### Frontend

1. Navigate to the `frontend` directory:

    ```sh
    cd frontend
    ```

2. Install the dependencies:

    ```sh
    npm install
    ```

3. Start the frontend development server:

    ```sh
    npm start
    ```

    The frontend server will start at `http://localhost:3000`.

## Running the Game

1. Ensure the backend server is running at `http://localhost:8080`.
2. Open your browser and navigate to `http://localhost:3000` to start playing the game.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
