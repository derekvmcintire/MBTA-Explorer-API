# MBTA Explorer API Documentation

## Overview
This is a work in progress.

The MBTA Explorer API provides real-time and static data related to the MBTA subway system. It is powered by the [MBTA V3 API](https://www.mbta.com/developers/v3-api) and offers live streaming of subway vehicle positions, caching for performance, and polyline decoding for route mapping.

## Features
- **Integration with MBTA V3 API**: Fetches and streams MBTA subway data directly from the official API.
- **Live Streaming**: Streams live vehicle positions from a single connection to the MBTA API, forwarding data to multiple clients.
- **Secure API Key Management**: Handles MBTA API keys securely.
- **Polyline Decoding**: Decodes polyline data for accurate route visualization.
- **Caching**: Utilizes Memcached for performance improvement.
- **CORS Configuration**: Configured for a frontend origin at `http://localhost:5173` by default.

## Installation

### Prerequisites
1. **Go**: Install [Go](https://go.dev/) (version 1.23.3 or higher).
2. **Memcached**: Install and run [Memcached](https://memcached.org/), or set up Memcached in Docker.
3. **MBTA API Key**: Obtain an API key from the [MBTA Developer Portal](https://www.mbta.com/developers/v3-api).

### Environment Variables
Create a `.env` file in the root directory:

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

#### Fetch Routes
- **`GET /api/routes`** Fetches MBTA route shapes and stops. It makes two separate requests to the MBTA V3 API. First to the `/stops` endpoint and secondly to the `/shapes` endpoint. It then combines the data and returns it in a single request.

- **Compression**: `GET /api/routes` returns a compressed response using `gzip`. Most modern browsers will handle this automatically, but be sure your client is setting the appropriate header:

  ```typescript
  headers: {
    'Accept-Encoding': 'gzip'
  }
  ```
  
- **Example Request**:
  ```bash
  curl --location 'http://localhost:8080/api/routes?route_ids=Mattapan'
  ```
  
- **Example Response**:
  ```json
  [
    {
      "id": "Mattapan",
      "coordinates": [
        [
          [
            42.2841,
            -71.0633
          ],
          [
            42.2839,
            -71.06318
          ],
          [
            42.28379,
            -71.0631
          ]
        ],
        [
          [
            42.26752,
            -71.09199
          ],
          [
            42.26752,
            -71.09199
          ],
          [
            42.267649999999996,
            -71.09116
          ]
        ]
      ],
      "stops": [
        {
          "id": "place-asmnl",
          "attributes": {
            "address": "Dorchester Ave and Ashmont St, Boston, MA 02124",
            "at_street": "",
            "description": null,
            "latitude": 42.28452,
            "longitude": -71.063777,
            "municipality": "Boston",
            "name": "Ashmont",
            "on_street": "",
            "platform_code": null,
            "platform_name": null,
            "vehicle_type": 0,
            "wheelchair_boarding": 1
          }
        },
        {
          "id": "place-cedgr",
          "attributes": {
            "address": "Fellsway St and Milton St, Dorchester, MA 02124",
            "at_street": "",
            "description": null,
            "latitude": 42.279629,
            "longitude": -71.060394,
            "municipality": "Boston",
            "name": "Cedar Grove",
            "on_street": "",
            "platform_code": null,
            "platform_name": null,
            "vehicle_type": 0,
            "wheelchair_boarding": 1
          }
        }
      ]
    }
  ]
  ```

  # LEFT OFF HERE, NEED TO UPDATE BELOW THIS LINE



  ---
  ---
  ---

#### Fetch Route Shapes
- **URL**: `GET /api/shapes?route_id={route_id}`
- **Description**: Fetches the shape of a route for mapping.
- **Example Request**:
  ```bash
  curl -X GET http://localhost:8080/api/shapes?route_id=Red
  ```
- **Example Response**:
  ```json
  {
    "route_id": "Red",
    "shapes": [
      {
        "id": "shape_7001",
        "polyline": "_p~iF~ps|U_ulLnnqC_mqNvxq`@"
      }
    ]
  }
  ```

### Streaming Endpoints

#### Stream Vehicles
- **URL**: `GET /stream/vehicles`
- **Description**: Streams live subway vehicle positions. Currently limited to subway data.
- **Example Request**:
  ```bash
  curl -N http://localhost:8080/stream/vehicles
  ```
- **Example Response (streamed)**:
  ```json
  {
    "id": "vehicle_1234",
    "route_id": "Red",
    "current_status": "IN_TRANSIT_TO",
    "latitude": 42.320685,
    "longitude": -71.052391,
    "direction_id": 0
  }
  {
    "id": "vehicle_5678",
    "route_id": "Orange",
    "current_status": "STOPPED_AT",
    "latitude": 42.363021,
    "longitude": -71.05829,
    "direction_id": 1
  }
  ```

## Configuration

### CORS Middleware
The application is configured to allow requests from `http://localhost:5173`. Update `cors.go` in the `internal/infrastructure/middleware` package to adjust origins.

### Memcached
Ensure Memcached is running. For Docker:
```bash
docker run --name memcached -d -p 11211:11211 memcached
```

## Development

### Code Formatting
```bash
make fmt
```

### Linting
```bash
make lint
```

### Testing
```bash
make test
```

## Future Enhancements
- **Dynamic Streaming**: Support for additional transit types such as buses and commuter rail.
- **Frontend Integration**: Sample frontend to visualize live data.
- **Docker Support**: Include Docker setup for easier deployment.

## Contributing
Contributions are welcome! Please open a pull request with a detailed description of your changes.

## License
This project is licensed under the [MIT License](LICENSE).

## Contact
For questions or support, open an issue in the repository.

