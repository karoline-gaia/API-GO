package configs

var cfg *conf

type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	JWTScecret    string
	JwtExperesIn  int
}

func LoadConfig(path string) (*conf, error)
