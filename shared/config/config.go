package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func ReadConfig(AppConfig string) *Config {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	port, _ := strconv.Atoi(os.Getenv(AppConfig))

	// mysql
	mysqlHost := os.Getenv("DB_MYSQL_HOST")
	mysqlPort, _ := strconv.Atoi(os.Getenv("DB_MYSQL_PORT"))
	mysqlUsername := os.Getenv("DB_MYSQL_USER")
	mysqlPassword := os.Getenv("DB_MYSQL_PASSWORD")
	mysqlDbName := os.Getenv("DB_MYSQL_NAME")

	mysqldebug, _ := strconv.ParseBool(os.Getenv("DB_MYSQL_DEBUG"))
	mysqlmaxopenconns, _ := strconv.Atoi(os.Getenv("DB_MYSQL_POOL_MAXOPENCONNS"))
	mysqlmaxidleconns, _ := strconv.Atoi(os.Getenv("DB_MYSQL_POOL_MAXIDLECONS"))
	mysqlmaxlifetime, _ := strconv.Atoi(os.Getenv("DB_MYSQL_POOL_MAXLIFETIME"))
	jwtUserSecret := os.Getenv("JWT_SECRET_KEY")
	jwtUserSecretExpire, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	rmqUrl := os.Getenv("RMQ_URL")
	rmqUsername := os.Getenv("RMQ_USERNAME")
	rmqPass := os.Getenv("RMQ_PASSWORD")

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisUsername := os.Getenv("REDIS_USERNAME")
	redisPass := os.Getenv("REDIS_PASSWORD")
	redisTls, _ := strconv.ParseBool(os.Getenv("REDIS_TLS"))
	redisDB := os.Getenv("REDIS_DB")

	config := Config{
		Port: port,
		Database: Database{
			Mysql: Mysql{
				Host:        mysqlHost,
				Port:        mysqlPort,
				Username:    mysqlUsername,
				Password:    mysqlPassword,
				DBName:      mysqlDbName,
				Debug:       mysqldebug,
				MaxOpenConn: mysqlmaxopenconns,
				MaxIdleConn: mysqlmaxidleconns,
				MaxLifetime: mysqlmaxlifetime,
			},
		},
		Jwt: Jwt{
			User: JwtConfig{
				Secret:       jwtUserSecret,
				SecretExpire: jwtUserSecretExpire,
			},
		},
		Rabbitmq: Rabbitmq{
			Host: rmqUrl,
			User: rmqUsername,
			Pass: rmqPass,
		}, Redis: Redis{
			Host:     redisHost,
			Port:     redisPort,
			Username: redisUsername,
			Password: redisPass,
			Tls:      redisTls,
			Db:       redisDB,
		},
	}
	return &config
}
