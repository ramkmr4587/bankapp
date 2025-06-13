package server

func Start() {
	r := setupRouter()
	r.Run(":8080")
}
