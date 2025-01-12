package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// CompressHandler wraps an HTTP handler and adds gzip compression for clients that support it.
func CompressHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the client supports gzip
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// If not, call the next handler without compression
			next.ServeHTTP(w, r)
			return
		}

		// Set the Content-Encoding header to gzip
		w.Header().Set("Content-Encoding", "gzip")

		// Create a gzip writer
		gz := gzip.NewWriter(w)
		defer gz.Close()

		// Wrap the ResponseWriter with gzipResponseWriter
		gzResponseWriter := gzipResponseWriter{ResponseWriter: w, Writer: gz}
		next.ServeHTTP(gzResponseWriter, r)
	})
}

// gzipResponseWriter is a wrapper for http.ResponseWriter that supports gzip compression.
type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
