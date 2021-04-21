package checker

// WebsiteChecker takes a url and returns true if the response is good
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites checks the status of a list of URLs.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultsChannel <- result{u, wc(u)}
			// results[u] = wc(u)
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	return results
}
