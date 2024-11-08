# Verity Calculator

The Verity Calculator is a tool designed to assist Destiny 2 players in solving the Verity puzzle encounter. This app simulates the encounter and generates step-by-step instructions for players. Built with Go and Go Fiber for a high-performance, multithreaded backend, the app handles concurrent requests efficiently, offering enhanced gameplay support.

## Features
- **Puzzle Solver**: Simulates the Verity encounter to provide a series of steps for players.
- **Multithreaded Backend**: Utilizes Go Fiber to handle multiple requests concurrently, enhancing performance.
- **Performance Comparison**: Benchmarked against a Node.js Express version, showing a 1.5x performance improvement in handling concurrent requests.

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/iambyt3z/verity-calculator.git
    cd verity-calculator
    ```
2. Install dependencies:
    ```bash
    go mod download
    ```
3. Run the application:
    ```bash
    go run main.go
    ```

## Usage
- Access the app in your browser or via HTTP clients like `curl` or Postman.
- Input necessary parameters related to the puzzle encounter to receive calculated steps.

## Performance
- The Go Fiber backend can handle approximately **800 concurrent requests** compared to **500 with Express**.
- Average latency for Go Fiber remains around **2 seconds** under peak load but performs similarly to Express at lower request levels.

## Technologies
- **Backend**: Go, Go Fiber
- **Load Testing**: Conducted using benchmarking tools to compare performance under high concurrency.

## License
This project is licensed under the MIT License.

---

Enjoy your gameplay! For contributions or issues, please open a pull request or issue.
