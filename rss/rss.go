package rss

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// type RSS struct {
// 	XMLName xml.Name `xml:"rss"`
// 	Channel Channel  `xml:"channel"`
// }

// type Channel struct {
// 	Title       string `xml:"title"`
// 	Link        string `xml:"link"`
// 	Description string `xml:"description"`
// 	Items       []Item `xml:"item"`
// }

// type Item struct {
// 	Title       string `xml:"title"`
// 	Link        string `xml:"link"`
// 	Description string `xml:"description"`
// 	PubDate     string `xml:"pubDate"`
// }

// func FetchRSSFeed(url string) (*RSS, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, fmt.Errorf("error fetching RSS feed: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading response body: %v", err)
// 	}

// 	var rss RSS
// 	err = xml.Unmarshal(body, &rss)
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing RSS feed: %v", err)
// 	}

// 	return &rss, nil
// }

type RSSFeed struct {
	URL      string `json:"url"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func GetFeeds() []*RSSFeed {
	feedFile, err := os.Open("./feeds.json")
	if err != nil {
		log.Fatalf("Failed to open feeds file: %s", err)
	}
	defer feedFile.Close()

	byteValue, err := io.ReadAll(feedFile)
	if err != nil {
		log.Fatalf("Failed to read feeds file: %s", err)
	}

	var feeds []*RSSFeed
	if err := json.Unmarshal(byteValue, &feeds); err != nil {
		log.Fatalf("Failed to unmarshal feeds file: %s", err)
	}

	return feeds
}
