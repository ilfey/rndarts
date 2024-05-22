package nekos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Kind = string

const (
	NEKO     Kind = "neko"
	KITSU    Kind = "kitsune"
	WAIFU    Kind = "waifu"
	HUSBANDO Kind = "husbando"
)

func GetRandomImage(kind Kind) (*Response, error) {
	resp, err := http.Get("https://nekos.best/api/v2/" + kind)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(string(responseBody))
	}

	res := new(Response)

	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type Response struct {
	Results []struct {
		ArtistHref *string `json:"artist_href"`
		ArtistName *string `json:"artist_name"`
		SourceURL  *string `json:"source_url"`
		AnimeName  *string `json:"anime_name"`
		URL        *string `json:"url"`
	} `json:"results"`
}
