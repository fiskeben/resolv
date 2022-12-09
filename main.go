package resolv

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Resolver contains the data from resolv.conf
type Resolver struct {
	Domains     []string
	Nameservers []string
	Search      []string
	Sortlist    []string
}

// Config reads /etc/resolv.conf and returns it as a Resolver
func Config() (Resolver, error) {
	f, err := os.Open("/etc/resolv.conf")
	if err != nil {
		return Resolver{}, err
	}
	defer f.Close()
	return parse(f)
}

func parse(f io.Reader) (Resolver, error) {
	domains := make([]string, 0)
	nameservers := make([]string, 0)
	search := make([]string, 0)
	sortlist := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		kind := parts[0]
		rest := parts[1:]

		switch kind {
		case "domain":
			for _, d := range rest {
				d := strings.TrimSpace(d)
				if d != "" {
					domains = append(domains, d)
				}
			}
		case "nameserver":
			n := strings.Join(rest, "")
			n = strings.TrimSpace(n)
			nameservers = append(nameservers, n)
		case "search":
			for _, s := range rest {
				s := strings.TrimSpace(s)
				if s != "" {
					search = append(search, s)
				}
			}
		case "sortlist":
			for _, s := range rest {
				s := strings.TrimSpace(s)
				if s != "" {
					sortlist = append(sortlist, s)
				}
			}
		}
	}

	return Resolver{
		Domains:     domains,
		Nameservers: nameservers,
		Search:      search,
		Sortlist:    sortlist,
	}, nil
}
