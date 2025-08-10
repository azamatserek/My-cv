package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Define structs to match your database schema
type Professor struct {
	Name      string `json:"name"`
	Title     string `json:"title"`
	Bio       string `json:"bio"`
	Contact   string `json:"contact"`
	Email     string `json:"email"`
	Image     string `json:"image"`
	Initials  string `json:"initials"`
}

type Publication struct {
	Title   string `json:"title"`
	Journal string `json:"journal"`
	Year    int    `json:"year"`
}

type Experience struct {
	Title      string `json:"title"`
	Institution string `json:"institution"`
	Period     string `json:"period"`
}

type Education struct {
	Degree      string `json:"degree"`
	Institution string `json:"institution"`
	Year        int    `json:"year"`
}

type Skills struct {
	Skills []string `json:"skills"`
}

// Full data structure to be sent to the frontend
type CVData struct {
	Professor   Professor     `json:"professor"`
	Experience  []Experience  `json:"experience"`
	Education   []Education   `json:"education"`
	Publications []Publication `json:"publications"`
	Skills      []string      `json:"skills"`
}

var db *sql.DB

func main() {
	// Database connection string. This should be configured via environment variables
	// in a production environment.
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "user", "password", "mydatabase")

	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Could not connect to database. Using mock data.")
	} else {
		log.Println("Successfully connected to the database!")
	}

	// Setup API endpoints
	http.HandleFunc("/api/data", getDataHandler)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// getDataHandler retrieves CV data from the database or returns mock data
func getDataHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the frontend to access the API
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// For simplicity, we'll use mock data. You can replace this with database queries.
	data := getMockData()

	// Set content type and encode the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getMockData provides a placeholder CV data structure
func getMockData() CVData {
	return CVData{
		Professor: Professor{
			Name: "Dr. Alex Johnson",
			Title: "Professor of Computer Science",
			Bio: "Dr. Johnson specializes in distributed systems and parallel computing. His research focuses on optimizing data-intensive applications for high-performance clusters.",
			Contact: "(555) 123-4567 | 123 University Ave, Anytown, USA",
			Email: "alex.johnson@example.edu",
			Image: "https://placehold.co/150x150/000000/FFFFFF?text=AJ",
			Initials: "AJ",
		},
		Experience: []Experience{
			{Title: "Professor", Institution: "State University", Period: "2015 - Present"},
			{Title: "Associate Professor", Institution: "State University", Period: "2010 - 2015"},
			{Title: "Assistant Professor", Institution: "State University", Period: "2005 - 2010"},
		},
		Education: []Education{
			{Degree: "Ph.D. in Computer Science", Institution: "University of Tech", Year: 2005},
			{Degree: "M.S. in Computer Science", Institution: "University of Tech", Year: 2002},
		},
		Publications: []Publication{
			{Title: "A Novel Approach to Distributed Hash Tables", Journal: "Journal of Distributed Systems", Year: 2022},
			{Title: "Scalable Concurrency Control for Cloud Databases", Journal: "ACM Transactions on Computing", Year: 2021},
		},
		Skills: []string{"Go", "Python", "Docker", "PostgreSQL", "React", "Cloud Computing", "Distributed Systems"},
	}
}