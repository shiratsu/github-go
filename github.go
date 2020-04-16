package github

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const searchURL = "https://api.github.com/search/repositories"

type GithubService interface {
	SearchRepos(terms string) *RepositorySearchResult
}

type GithubClient struct {
	RequestUrl string
}

func DefaultClient() *GithubClient {
	return &GithubClient{searchURL}
}

// (4) 本番URLにリクエストするAPIクライアントを返すパッケージ関数
func DefaultClient() *GithubClient {
    return &GithubClient{searchURL}
}


type StatusCode int

func Success() int {
    return 1
}

func ErrorNetwork() int {
    return -1
}

func ErrorHttp() int {
    return -2
}
