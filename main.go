package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type supply struct {
	fs http.Handler
}

func (s *supply) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Requested: %s", r.URL.Path)
	w.Header().Add("Content-Type", "text/javascript")
	s.fs.ServeHTTP(w, r)
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fs := http.FileServer(http.Dir("./plugins"))
	supp := &supply{fs}
	http.ListenAndServeTLS(viper.GetString("address"), viper.GetString("certFile"), viper.GetString("keyFile"), supp)
}
