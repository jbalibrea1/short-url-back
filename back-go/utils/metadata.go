package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Metadata contiene la información de los metadatos de la URL
type Metadata struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Favicon     string `json:"favicon"`
}

// GetMetadata obtiene los metadatos de una URL dada
func GetMetadata(urlStr string) (*Metadata, error) {
	meta, err := fetchMetadata(urlStr)
	if err != nil {
		// Si obtenemos un error 403 con http, intentar con https
		if strings.HasPrefix(urlStr, "http://") && err.Error() == "unexpected status code: 403" {
			urlStr = "https://" + strings.TrimPrefix(urlStr, "http://")
			meta, err = fetchMetadata(urlStr)
		}
	}
	return meta, err
}

func fetchMetadata(urlStr string) (*Metadata, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return fmt.Errorf("stopped after 10 redirects")
			}
			req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
			req.Header.Add("Accept-Language", "en-US,en;q=0.9")
			return nil
		},
	}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	finalURL := resp.Request.URL.String()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	meta := &Metadata{}

	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		meta.Title = strings.TrimSpace(s.Text())
	})

	doc.Find("meta[name=description]").Each(func(i int, s *goquery.Selection) {
		desc, _ := s.Attr("content")
		meta.Description = strings.TrimSpace(desc)
		maxLen := 50
		if len(meta.Description) > maxLen {
			meta.Description = meta.Description[:maxLen] + "..."
		}
	})

	var faviconURL string
	doc.Find("link[rel='icon'], link[rel='shortcut icon']").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && !strings.HasPrefix(href, "data:") {
			faviconURL = href
		}
	})

	if faviconURL == "" {
		parsedURL, err := url.Parse(finalURL)
		if err == nil {
			faviconURL = parsedURL.Scheme + "://" + parsedURL.Host + "/favicon.ico"
		} else {
			faviconURL = finalURL + "/favicon.ico"
		}
	}

	if !strings.HasPrefix(faviconURL, "http") {
		baseURL, _ := url.Parse(finalURL)
		faviconURL = baseURL.ResolveReference(&url.URL{Path: path.Join(baseURL.Path, faviconURL)}).String()
	}
	meta.Favicon = faviconURL

	return meta, nil
}
