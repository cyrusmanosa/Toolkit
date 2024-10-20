package main

import (
	ginLoginTest "appTest/Gin/LoginTest"
	ginUploadTest "appTest/Gin/UploadTest"
	muxLoginTest "appTest/Mux/LoginTest"
	muxUploadTest "appTest/Mux/UploadTest"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	GinMode()
}

func MuxMode() {
	// MuxLoginRoutes()
	MuxUploadRoutes()
}

func MuxLoginRoutes() {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./LoginTest"))))

	mux.HandleFunc("/api/login", muxLoginTest.Login)
	mux.HandleFunc("/api/logout", muxLoginTest.Logout)

	log.Println("Starting application on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func MuxUploadRoutes() {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./UploadTest"))))

	mux.HandleFunc("/upload", muxUploadTest.UploadFiles)
	mux.HandleFunc("/upload-one", muxUploadTest.UploadOneFiles)

	log.Println("Starting application on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

//// ------------------------------ GIN ----------------------------------

func GinMode() {
	// ginLoginRouters()
	ginUploadRouters()
}

func ginLoginRouters() {
	r := gin.Default()
	r.Use(cors.Default())

	r.LoadHTMLFiles("/Users/cyrusman/Desktop/ProgrammingLearning/Udemy/Toolkit/appTest/Gin/LoginTest/index.html")

	// Serve the HTML file at the root URL
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) // Serve the loaded HTML file
	})

	r.POST("/api/login", ginLoginTest.Login)
	r.POST("/api/logout", ginLoginTest.Logout)

	log.Println("Starting application on port 8080")
	r.Run(":8080")
}

func ginUploadRouters() {
	r := gin.Default()
	r.Use(cors.Default())

	r.LoadHTMLFiles("/Users/cyrusman/Desktop/ProgrammingLearning/Udemy/Toolkit/appTest/Gin/UploadTest/index.html")

	// Serve the HTML file at the root URL
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) // Serve the loaded HTML file
	})

	r.POST("/upload", ginUploadTest.UploadFiles)
	r.POST("/upload-one", ginUploadTest.UploadOneFiles)

	log.Println("Starting application on port 8080")
	r.Run(":8080")
}
