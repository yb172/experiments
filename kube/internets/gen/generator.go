package gen

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/yb172/experiments/kube/internets/proto/wordgen"
	"golang.org/x/net/html"
)

const addr = "http://randomword.com"

// Service provides generation services
type Service struct {
}

// GetWord generates word
func (s *Service) GetWord(ctx context.Context, in *wordgen.GetWordReq) (*wordgen.GetWordResp, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return nil, fmt.Errorf("error whlie requesting word: %v", err)
	}
	b := resp.Body
	defer b.Close() // close Body when the function returns

	word, err := extractWord(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing page with word: %v", err)
	}

	return &wordgen.GetWordResp{Word: word}, nil
}

func checkWord(t html.Token) bool {
	// Element should be div
	if t.Data != "div" {
		return false
	}
	for _, a := range t.Attr {
		// And it should has id attribute with value "random_word"
		if a.Key == "id" && a.Val == "random_word" {
			return true
		}
	}
	return false
}

func extractWord(resp io.Reader) (string, error) {
	z := html.NewTokenizer(resp)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document
			return "", fmt.Errorf("word is not found")
		case tt == html.StartTagToken:
			t := z.Token()
			if isWord := checkWord(t); isWord {
				z.Next()
				return string(z.Text()), nil
			}
		}
	}
}
