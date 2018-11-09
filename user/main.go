package user

import (
	"bufio"
	"fmt"
	"guess-the-word/postgres"
	"os"
	"strings"

	"github.com/caarlos0/env"
	"github.com/jackc/pgx"
)

type config struct {
	DbHost       string `env:"DB_HOST" envDefault:"localhost"`
	DbName       string `env:"DB_NAME" envDefault:"sample"`
	DbUser       string `env:"DB_USER" envDefault:"postgres"`
	DbPass       string `env:"DB_PASS" envDefault:"admin"`
	DbSchema     string `env:"DB_SCHEMA" envDefault:"sample"`
	InternalPort int    `env:"INTERNAL_PORT" envDefault:"8585"`
	Debug        bool   `env:"DEBUG" envDefault:"false"`
	DebugSpans   bool   `env:"DEBUG_SPANS" envDefault:"false"`
	GRPCServer   string `env:"GRPC_SERVER" envDefault:"localhost:50051"`
}

var cfg config
var db *pgx.ConnPool

func Setup() {
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	db, err = postgres.OpenDB(cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbHost)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func LoginWithData(username string, password string) (uid string) {
	query := fmt.Sprintf(`
		SELECT uid 
			FROM
				sample.users
			WHERE
				username='%s' and password='%s';
	`, username, password)

	db.QueryRow(query).Scan(&uid)

	if uid == "" {
		return "Failed"
	}

	return uid
}

func Login() (uid string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your username: ")
	// take input
	username, _ := reader.ReadString('\n')

	//clean input => in go there is two extra characters taken with human input
	username = strings.TrimSuffix(username, "\n")
	username = strings.TrimSuffix(username, string(13))

	fmt.Print("Enter your password: ")
	// take input
	password, _ := reader.ReadString('\n')

	//clean input => in go there is two extra characters taken with human input
	password = strings.TrimSuffix(password, "\n")
	password = strings.TrimSuffix(password, string(13))

	return LoginWithData(username, password)
}

func SignUpWithData(username string, password string) (uid string) {
	query := fmt.Sprintf(`
		INSERT INTO
			sample.users
				(uid, username, password)
			VALUES
				(uuid_generate_v4(), '%s', '%s');
	`, username, password)

	_, err := db.Exec(query)

	if err == nil {
		return LoginWithData(username, password)
	}

	return "Failed"
}

func SignUp() (uid string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your username: ")
	// take input
	username, _ := reader.ReadString('\n')

	//clean input => in go there is two extra characters taken with human input
	username = strings.TrimSuffix(username, "\n")
	username = strings.TrimSuffix(username, string(13))

	fmt.Print("Enter your password: ")
	// take input
	password, _ := reader.ReadString('\n')

	//clean input => in go there is two extra characters taken with human input
	password = strings.TrimSuffix(password, "\n")
	password = strings.TrimSuffix(password, string(13))

	return SignUpWithData(username, password)
}

func AddScore(uid string, score float64) (isSuccessful bool) {
	query := fmt.Sprintf(`
		INSERT INTO
			sample.scores
				(id, uid, score)
			VALUES
				(uuid_generate_v4(), '%s', '%f');
	`, uid, score)

	_, err := db.Exec(query)

	if err == nil {
		return true
	}

	return false
}

func GetTotalScore(uid string) (score float64) {
	query := fmt.Sprintf(`
		SELECT SUM(score) 
			FROM
				sample.scores
			WHERE
				uid='%s';
	`, uid)

	db.QueryRow(query).Scan(&score)

	return score
}
