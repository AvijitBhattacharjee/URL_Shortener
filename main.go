package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

const serverAddress = ":8080" // Define server address as a constant

// In-memory store for URL domains and their counts
var domainCounts = make(map[string]int)
var shortenedURLs = make(map[string]string)
var urlIDCounter int
var mu sync.Mutex

// URLShortenRequest represents the request payload for shortening a URL
type URLShortenRequest struct {
	URL string `json:"url"`
}

// DomainCount represents the domain and its count
type DomainCount struct {
	Domain string `json:"domain"`
	Count  int    `json:"count"`
}

func main() {
	r := mux.NewRouter()

	// Route to shorten a URL
	r.HandleFunc("/shorten", shortenURL).Methods("POST")

	// Route to get top 3 domains
	r.HandleFunc("/metrics/top-domains", getTopDomains).Methods("GET")

	// Route to redirect based on shortened URL ID
	r.HandleFunc("/{id}", redirectURL).Methods("GET")

	log.Printf("Server started on %s", serverAddress)

	err := http.ListenAndServe(serverAddress, r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	var req URLShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	domain := extractDomain(req.URL)
	if domain == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	mu.Lock()
	domainCounts[domain]++
	urlIDCounter++
	id := generateID(urlIDCounter)
	shortenedURLs[id] = req.URL
	log.Printf("Domain %s count incremented to %d", domain, domainCounts[domain])
	log.Printf("URL %s shortened to ID %s", req.URL, id)
	mu.Unlock()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "URL shortened successfully", "id": id, "domain": domain})
}

func getTopDomains(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if len(domainCounts) == 0 {
		log.Println("No domains have been shortened yet")
		json.NewEncoder(w).Encode([]DomainCount{})
		return
	}

	domainList := make([]DomainCount, 0, len(domainCounts))
	for domain, count := range domainCounts {
		domainList = append(domainList, DomainCount{Domain: domain, Count: count})
	}

	sort.Slice(domainList, func(i, j int) bool {
		return domainList[i].Count > domainList[j].Count
	})

	topDomains := domainList
	if len(domainList) > 3 {
		topDomains = domainList[:3]
	}

	log.Printf("Top domains: %+v", topDomains)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(topDomains)
}

func redirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	mu.Lock()
	url, exists := shortenedURLs[id]
	mu.Unlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func extractDomain(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	parsedURL := strings.Split(url, "/")
	if len(parsedURL) > 2 {
		return parsedURL[2]
	}
	return ""
}

func generateID(counter int) string {
	return strconv.Itoa(counter)
}
