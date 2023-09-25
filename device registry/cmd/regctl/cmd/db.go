package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	hostname string
	dbuser   string
	dbpass   string
	dbname   string
	sslmode  string
	dbport   int
)

func NewDBCmd() *cobra.Command {

	//hostname = "localhost"
	//	port     = "5432"
	//	user     = "postgres"
	//	password = "postgres"
	//	dbname   = "postgres"
	//	sslmode  = "disable"

	dbCmd := &cobra.Command{
		Use:   "db",
		Short: "registry database management",
		Long:  "command for registry database management",
		Run: func(cmd *cobra.Command, args []string) {
			logUsage(cmd.UsageString())
		},
	}

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "regctl db init",
		Long:  "initialize databases for regsvc",
		Run: func(cmd *cobra.Command, args []string) {

			wd, err := os.Getwd()
			if err != nil {
				log.Println(err)
			}

			cfg := dbConfig{
				Hostname: hostname,
				Username: dbuser,
				Password: dbpass,
				DBName:   dbname,
				SSLMode:  sslmode,
				Port:     dbport,
			}

			db, err := dbConnect(cfg)
			if err != nil {
				logError(err)
				return
			}

			path := filepath.Join(wd, "data.sql")

			c, ioErr := ioutil.ReadFile(path)

			if ioErr != nil {
				logError(ioErr)
				return
			}

			requests := strings.Split(string(c), ";")

			for _, request := range requests {
				_, err := db.Exec(request)
				if err != nil {
					logError(err)
					os.Exit(1)
					return
				}
			}

		},
	}

	testCmd := &cobra.Command{
		Use:   "test",
		Short: "regctl db test ",
		Long:  "test if all databases are well set",
		Run: func(cmd *cobra.Command, args []string) {
			logUsage(cmd.UsageString())
		},
	}

	allCmd := &cobra.Command{
		Use:   "all",
		Short: "test all registry database tables",
		Long:  "test if all databases are well set",
		Run: func(cmd *cobra.Command, args []string) {
			logUsage(cmd.UsageString())
		},
	}

	usersCmd := &cobra.Command{
		Use:   "users",
		Short: "test users database table",
		Long:  "test if all databases are well set",
		Run: func(cmd *cobra.Command, args []string) {
			logUsage(cmd.UsageString())
		},
	}

	nodesCmd := &cobra.Command{
		Use:   "nodes",
		Short: "test nodes database table",
		Long:  "test if all databases are well set",
		Run: func(cmd *cobra.Command, args []string) {
			logUsage(cmd.UsageString())
		},
	}

	regionsCmd := &cobra.Command{
		Use:   "regions",
		Short: "test regions database table",
		Long:  "test the regions database",
		Run: func(cmd *cobra.Command, args []string) {
			logUsage(cmd.UsageString())
		},
	}

	testCmd.AddCommand(allCmd, nodesCmd, usersCmd, regionsCmd)

	pingCmd := &cobra.Command{
		Use:   "ping",
		Short: "regctl db ping",
		Long:  "test the regions database",
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Printf("hostname :%s, port : %d, user : %s, password : %s, dbname: %s , sslmode: %s\n",
			//	hostname,dbport,dbuser,dbpass,dbname,sslmode)
			cfg := dbConfig{
				Hostname: hostname,
				Username: dbuser,
				Password: dbpass,
				DBName:   dbname,
				SSLMode:  sslmode,
				Port:     dbport,
			}

			_, err := dbConnect(cfg)

			if err != nil {
				logError(err)
				return
			}
			logOK()
		},
	}

	dbCmd.PersistentFlags().StringVar(&hostname, "hostname", "localhost", "registry db hostname")
	dbCmd.PersistentFlags().StringVar(&dbuser, "dbuser", "postgres", "database user")
	dbCmd.PersistentFlags().StringVar(&dbpass, "dbpass", "postgres", "database password")
	dbCmd.PersistentFlags().StringVar(&dbname, "dbname", "postgres", "database name")
	dbCmd.PersistentFlags().StringVar(&sslmode, "sslmode", "disable", "ssl mode")
	dbCmd.PersistentFlags().IntVar(&dbport, "dbport", 5432, "database port")

	dbCmd.AddCommand(initCmd, testCmd, pingCmd)

	return dbCmd
}

type dbConfig struct {
	Hostname string
	Username string
	Password string
	DBName   string
	SSLMode  string
	Port     int
}

func dbConnect(config dbConfig) (db *sql.DB, err error) {
	url := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.Hostname, config.Port, config.Username, config.DBName, config.Password, config.SSLMode)

	db, err = sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
