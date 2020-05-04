package main

func main() {
	configuration := Configuration

	// Configuring port
	port := configuration.Server.Port
	if port == "" {
		port = "8080"
	}

	engine := NewEngine()
	_ = engine.Run(":" + port)
}
