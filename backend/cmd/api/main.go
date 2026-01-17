package main

import (
"encoding/json"
"log"
"net/http"
"os"
"time"
)

type healthResponse struct {
Status string `json:"status"`
Time   string `json:"time"`
}

func main() {
port := os.Getenv("PORT")
if port == "" {
port = "8080"
}

mux := http.NewServeMux()

mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
_ = json.NewEncoder(w).Encode(healthResponse{
Status: "ok",
Time:   time.Now().UTC().Format(time.RFC3339),
})
})

mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
w.WriteHeader(http.StatusOK)
_, _ = w.Write([]byte("StartTech API running\n"))
})

server := &http.Server{
Addr:              ":" + port,
Handler:           loggingMiddleware(mux),
ReadHeaderTimeout: 5 * time.Second,
}

log.Printf("listening on :%s", port)
log.Fatal(server.ListenAndServe())
}

func loggingMiddleware(next http.Handler) http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
start := time.Now()
next.ServeHTTP(w, r)
log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
})
}
