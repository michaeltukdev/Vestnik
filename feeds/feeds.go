package feeds

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	ParsedDate  time.Time
	Source      string
	Category    string
}

type RSSFeed struct {
	URL      string `json:"url"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func FetchRSSFeed(url string) (*RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching RSS feed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var rss RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return nil, fmt.Errorf("error parsing RSS feed: %v", err)
	}

	for i, item := range rss.Channel.Items {
		parsedDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			parsedDate, err = time.Parse(time.RFC1123, item.PubDate)
			if err != nil {
				log.Printf("Failed to parse date for item '%s': %v", item.Title, err)
				continue
			}
		}
		rss.Channel.Items[i].ParsedDate = parsedDate
	}

	return &rss, nil
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

func FetchAndCombineFeeds() ([]Item, error) {
	feeds := GetFeeds()
	var allItems []Item

	for _, feed := range feeds {
		rssData, err := FetchRSSFeed(feed.URL)
		if err != nil {
			log.Printf("Error fetching feed '%s': %v", feed.Name, err)
			continue
		}
		for _, item := range rssData.Channel.Items {
			item.Source = feed.Name
			item.Category = feed.Category
			allItems = append(allItems, item)
		}
	}

	sort.Slice(allItems, func(i, j int) bool {
		return allItems[i].ParsedDate.After(allItems[j].ParsedDate)
	})

	return allItems, nil
}
