package gen

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/yb172/experiments/kube/internets/proto/wordgen"
	"golang.org/x/net/html"
)

const addr = "https://randomword.com"

// Server provides generation services
type Server struct {
}

// GetWord generates word
func (s *Server) GetWord(ctx context.Context, in *wordgen.GetWordReq) (*wordgen.GetWordResp, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return nil, fmt.Errorf("error whlie requesting word: %v", err)
	}
	b := resp.Body
	defer b.Close() // close Body when the function returns

	word, err := extractWord(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while parsing page with word")
	}

	return &wordgen.GetWordResp{Word: word}, nil
}

func checkWord(t html.Token) (string, bool) {
	for _, a := range t.Attr {
		if a.Key == "id" && a.Val == "random_word" {
			return t.Data, true
		}
	}
	return "", false
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

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			if word, hasWord := checkWord(t); hasWord {
				return word, nil
			}
		}
	}
}
