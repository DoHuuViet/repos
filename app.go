package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	apiEndpointFormat = "https://api.github.com/users/%s/repos"
)

type Repository struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	StargazersCount int64 `json:"stargazers_count"`
	Owner RepositoryOwner `json:"owner"`
}

type RepositoryOwner struct {
	Id int64 `json:"id"`
	Login string `json:"login"`
}

type ErrorResp struct {
	Code int `json:"status_code"`
	Message string `json:"error_description"`
}

func GetRepositories(c *gin.Context)  {
	userId := c.Param("userId")
	name := c.Param("name")
	repositories, code := GetRepositoriesByUserId(userId, name)
	if code == http.StatusInternalServerError {
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorResp{500, "internal system error"})
		return
	}

	if code == http.StatusNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorResp{404, "not found"})
		return
	}
	c.JSON(http.StatusOK, repositories)
}

func main() {
	router := gin.Default()
	router.GET("/:userId/repositories", GetRepositories)
	router.GET("/:userId/repositories/:name", GetRepositories)
	router.Run(":5000")
}

func GetRepositoriesByUserId(userId string, name string) ([]Repository, int) {
	url := fmt.Sprintf(apiEndpointFormat, userId)
	httpClient := &http.Client{Timeout: 10 * time.Second}
	resp, err := httpClient.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Cannot get user's repositories: %s", err)
		return nil, 500
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, 404
	}

	respBody, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		log.Fatal(er)
		return nil, 500
	}

	var repositories []Repository
	e := json.Unmarshal(respBody, &repositories)

	if e != nil {
		log.Fatal(e)
		return nil, 500
	}

	if len(name) > 0 {
		repos := []Repository{}
		for _, v := range repositories {
			if strings.Contains(v.Name, name) {
				repos = append(repos, v)
			}
		}
		return repos, 200
	} else {
		sortByStars(repositories)
		//sortByName(repositories)
		return repositories, 200
	}
}

func sortByName(repositories []Repository) {
	sort.SliceStable(repositories, func(i, j int) bool {
		return repositories[i].Name < repositories[j].Name
	})
}

func sortByStars(repositories []Repository) {
	sort.SliceStable(repositories, func(i, j int) bool {
		return repositories[i].StargazersCount < repositories[j].StargazersCount
	})
}




