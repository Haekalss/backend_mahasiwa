package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://example.com",
	"http://127.0.0.1:8088",
	"https://frontend-mahasiwa-47l1.vercel.app/",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
