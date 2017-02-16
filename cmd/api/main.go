package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joeshaw/envdecode"
)

type config struct {
	Port int `env:"PORT,default=5000"`

	AWS struct {
		Region          string `env:"AWS_REGION,required"`
		AccessKeyID     string `env:"AWS_ACCESS_KEY_ID,required"`
		SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY,required"`
		Bucket          string `env:"S3_BUCKET,default=thingsleilasays"`
		ObjectName      string `env:"S3_OBJECT_NAME,default=tweets.json"`
	}
}

func main() {
	var cfg config
	if err := envdecode.Decode(&cfg); err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
}
