package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joeshaw/envdecode"
)

type config struct {
	Username string `env:"TWITTER_USERNAME,required"`

	Twitter struct {
		ConsumerKey    string `env:"TWITTER_CONSUMER_KEY,required"`
		ConsumerSecret string `env:"TWITTER_CONSUMER_SECRET,required"`
		AccessToken    string `env:"TWITTER_ACCESS_TOKEN,required"`
		AccessSecret   string `env:"TWITTER_ACCESS_SECRET,required"`
		TweetLimit     int    `env:"TWITTER_TWEET_LIMIT,default=3200"`
	}

	AWS struct {
		Region      string `env:"AWS_REGION,required"`
		AccessKeyID string `env:"AWS_ACCESS_KEY_ID,required"`
		Bucket      string `env:"S3_BUCKET,default=thingsleilasays"`
		ObjectName  string `env:"S3_OBJECT_NAME,default=tweets.json"`
	}
}

func newTwitterClient(consumerKey, consumerSecret, accessToken, accessSecret string) *twitter.Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return twitter.NewClient(httpClient)
}

func main() {
	var cfg config
	if err := envdecode.Decode(&cfg); err != nil {
		log.Fatal(err)
	}

	client := newTwitterClient(cfg.Twitter.ConsumerKey, cfg.Twitter.ConsumerSecret, cfg.Twitter.AccessToken, cfg.Twitter.AccessSecret)
	params := &twitter.UserTimelineParams{
		ScreenName: cfg.Username,
		Count:      cfg.Twitter.TweetLimit,
	}
	tweets, resp, err := client.Timelines.UserTimeline(params)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := json.Marshal(tweets)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
