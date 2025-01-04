# MBTA-Explorer-API

# Project Folder Structure

```plaintext
cmd/
└── api/
    └── main.go            // Entry point of the application

internal/
├── adapters/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── live_position_handler.go    // Handles /api/live for initial live data
│   │   │   ├── route_shapes_handler.go     // Handles fetching route shapes
│   │   │   ├── route_stops_handler.go      // Handles fetching route stops
│   │   │   ├── routes_handler.go           // Handles fetching routes
│   │   │   └── stream_vehicles_handler.go  // Handles /stream/vehicles for live data updates
│   │   └── response/
│   │       └── get_subway_response.go      // Defines GetRouteResponse struct
│   ├── data/
│   │   ├── fetch_data.go                   // Generic fetch helper for HTTP requests
│   │   ├── mbta_api.go                     // MBTA client for API data retrieval
│   │   └── memcached.go                    // Memcached client for caching responses
│   └── http/
│       └── routes.go                       // API route definitions
├── constants/
│   └── stream.go                           // URL constants for live vehicle streaming
├── core/
│   ├── domain/
│   │   ├── models/
│   │   │   ├── route.go                    // Domain model for a route
│   │   │   ├── shape.go                    // Domain model for a shape
│   │   │   ├── stop.go                     // Domain model for a stop
│   │   │   └── vehicle.go                  // Domain model for a vehicle
│   ├── fetch_from_mbta.go                  // Core logic for fetching MBTA data
│   └── stream_manager.go                   // Stream manager logic
├── infrastructure/
│   ├── config/
│   │   ├── mbta_api_config.go              // MBTA API configuration (e.g., GetAPIKey)
│   │   └── memcached_config.go             // Memcached configuration
│   ├── middleware/
│   │   └── cors.go                         // CORS middleware
│   └── stream/
│       ├── broadcast.go                    // Broadcasting stream data to clients
│       ├── factory.go                      // Stream manager factory
│       ├── fetch.go                        // Fetching stream data from MBTA
│       ├── initialize.go                   // Stream manager initialization
│       ├── manager.go                      // Core StreamManager implementation
│       ├── scanner.go                      // Processes and scans stream data
│       ├── sse.go                          // Server-Sent Events (SSE) parsing
│       └── utils.go                        // Utility functions for stream management
├── pkg/
│   └── decode_polyline.go                  // Utility for decoding polyline data
├── ports/
│   └── data/
│       └── api.go                          // Defines MBTAClient interface
└── usecases/
    ├── fetch_from_mbta.go                  // Logic for fetching MBTA data
    └── stream_from_mbta.go                 // (WIP) Logic for streaming MBTA data

Root Files
├── .env                                   // Environment variables
├── .gitignore                             // Git ignore rules
├── go.mod                                 // Go module definition
├── go.sum                                 // Go module dependency checksums
├── Makefile                               // Build automation commands
└── README.md                              // Project documentation
```
