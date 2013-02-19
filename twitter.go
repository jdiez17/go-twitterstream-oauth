package twitterstream

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mrjones/oauth"
)

type User struct {
	Id          uint64
	Screen_name string
	Name        string
}

type Tweet struct {
	Id   uint64
	User User
	Text string
}

const k_userstream = "https://userstream.twitter.com/2/user.json"

//const k_userstream = "https://stream.twitter.com/1.1/statuses/sample.json"

func DoUserStream(oauth Oauth, accessToken *oauth.AccessToken) (<-chan *Tweet, error) {
	tweets := make(chan *Tweet)
	response, err := oauth.Consumer.Get(k_userstream, map[string]string{}, accessToken)

	if err != nil {
		return nil, err
	}
	go func() {
		defer close(tweets)
		defer response.Body.Close()
		reader := bufio.NewReader(response.Body)

		for {
			buf, err := reader.ReadBytes('\n')
			if err != nil {
				panic(err)
			}

			tweet := new(Tweet)

			err = json.Unmarshal(buf, tweet)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			if tweet.Id != 0 {
				tweets <- tweet
			}
		}
	}()

	return tweets, nil
}
