# MBTA Train Tracker API

A Go-based API for tracking live vehicle locations and fetching route information from the MBTA v3 API. This application provides a streaming endpoint for live subway vehicle data, along with endpoints for fetching routes, stops, and shapes.

## Features

- **Live Streaming**: Stream live vehicle positions for the MBTA subway system.
- **Static Data Fetching**:
  - Fetch MBTA routes.
  - Fetch stops for a specific route.
  - Fetch route shapes for mapping.
- **Caching**: Utilizes Memcached for caching data to improve performance.
- **CORS Configuration**: Configured for development with a default `localhost:5173` frontend origin.

## Installation

### Prerequisites

1. **Go**: Ensure [Go](https://go.dev/) is installed (version 1.23.3 or higher).
2. **Memcached**: Install and run [Memcached](https://memcached.org/). Alternatively, set up Memcached in Docker (Docker setup not included yet).
3. **MBTA API Key**: Obtain an API key from the [MBTA Developer Portal](https://www.mbta.com/developers/v3-api).

### Environment Variables

Create a `.env` file in the root directory and add the following:

```dotenv
MBTA_API_KEY=your_mbta_api_key_here
```

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo-name/explorer.git
   cd explorer
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Build the project:
   ```bash
   make build
   ```
4. Run the project:
   ```bash
   make run
   ```

## Endpoints

### Static Data Endpoints

- **Fetch Routes**: `GET /api/routes`
- **Fetch Route Stops**: `GET /api/stops`
- **Fetch Route Shapes**: `GET /api/shapes`
- **Live Vehicle Positions**: `GET /api/live`

### Streaming Endpoint

- **Stream Vehicles**: `GET /stream/vehicles`  
  Streams live subway vehicle positions. (Currently limited to subway data.)

## Configuration

### CORS Middleware

The application is configured to allow requests from `http://localhost:5173`. Update the middleware configuration in `internal/infrastructure/middleware` if a different origin is required.

### Memcached

Ensure Memcached is running before starting the API. For Docker:

```bash
docker run --name memcached -d -p 11211:11211 memcached
```

## Development

### Code Formatting

Run the formatter:

```bash
make fmt
```

### Linting

Run the linter:

```bash
make lint
```

### Testing

Run tests:

```bash
make test
```

## Future Enhancements

- **Docker Support**: Add Docker setup for easy deployment.
- **Dynamic Streaming**: Allow clients to specify the transit type (e.g., bus, commuter rail).
- **Frontend Integration**: Provide a sample frontend to visualize live tracking.

## Project Structure

cmd/
└── api/
└── main.go // Entry point of the application

internal/
├── adapters/
│ ├── data/
│ │ ├── fetch_data.go // Generic fetch helper for HTTP requests
│ │ ├── mbta_api.go // MBTA client for API data retrieval
│ │ └── memcached.go // Memcached client for caching responses
│ ├── distribute/
│ │ └── distributor.go // Stream distribution manager
│ └── http/
│ └── routes.go // API route definitions
│ └── mbta/
│ ├── api/
│ │ ├── handlers/
│ │ │ ├── live_position_handler.go // Handles /api/live for initial live data
│ │ │ ├── route_shapes_handler.go // Handles fetching route shapes
│ │ │ ├── route_stops_handler.go // Handles fetching route stops
│ │ │ ├── routes_handler.go // Handles fetching routes
│ │ │ └── stream_vehicles_handler.go // Handles /stream/vehicles for live updates
│ │ └── response/
│ │ └── get_subway_response.go // Defines GetRouteResponse struct
│ ├── stream/
│ ├── fetch.go // Helper for fetching stream data
│ ├── process.go // Processes incoming stream
│ ├── scanner.go // Scans incoming stream
│ └── source.go // Stream source manager
├── constants/
│ └── stream.go // URL constants for live vehicle streaming
├── core/
│ ├── domain/
│ │ ├── models/
│ │ │ ├── route.go // Domain model for a route
│ │ │ ├── shape.go // Domain model for a shape
│ │ │ ├── stop.go // Domain model for a stop
│ │ │ └── vehicle.go // Domain model for a vehicle
│ ├── usecases/
│ │ ├── fetch_from_mbta.go // Logic for fetching MBTA data
│ │ ├── mbta_helper.go // Helper for MBTA API data
│ │ ├── stream_manager.go // Combines stream manager and distributor
│ │ └── stream_from_mbta.go // Logic for streaming MBTA data
│ └── stream_manager.go // Stream manager logic
├── infrastructure/
│ ├── config/
│ │ ├── mbta_api_config.go // MBTA API configuration (e.g., GetAPIKey)
│ │ └── memcached_config.go // Memcached configuration
│ ├── middleware/
│ └── cors.go // CORS middleware
├── pkg/
│ └── decode_polyline.go // Utility for decoding polyline data
├── ports/
│ ├── data/
│ │ └── api.go // Defines MBTAClient interface
│ ├── streaming/
│ └── stream.go // Interfaces for StreamManager and related managers

Root Files:
├── .env // Environment variables
├── .gitignore // Git ignore rules
├── go.mod // Go module definition
├── go.sum // Go module dependency checksums
├── Makefile // Build automation commands
└── README.md // Project documentation

## Contributing

Contributions are welcome! Please submit a pull request with a detailed description of your changes.

## License

This project is licensed under the [MIT License](LICENSE).

## Contact

For questions or support, please open an issue in the repository.
