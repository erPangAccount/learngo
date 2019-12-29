package controller

import (
	"crawler/engine"
	"crawler/frontend/common"
	"crawler/frontend/model"
	"github.com/olivere/elastic"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   common.SearchResultTempalteView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) (SearchResultHandler, error) {
	var result SearchResultHandler
	result.view = common.CreateSearchResultTemplateView(template)
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(engine.ElasticHost))
	if err != nil {
		return result, err
	}
	result.client = client

	return result, nil
}

func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	targetPage, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		targetPage = 0
	}

	var searchResult model.SearchResult
	searchResult, err = s.getSearchResult(q, targetPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = s.view.Render(w, searchResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s SearchResultHandler) getSearchResult(q string, targetPage int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Keyword = q
	result.PageInfo.TargetPage = targetPage
	result.PageInfo.PageSize = 10

	query := elastic.NewQueryStringQuery(q)

	searchResult, err := s.client.Search("test").Query(query).Do(context.Background())
	log.Printf("%s, %v", searchResult, err)

	if err != nil {
		return result, err
	}

	result.PageInfo.Total = int(searchResult.TotalHits())
	if targetPage > 0 {
		result.PageInfo.PrevPage = targetPage - 1
	}

	if targetPage*result.PageInfo.PageSize < result.PageInfo.Total {
		result.PageInfo.NextPage = targetPage + 1
	}
	result.Items = searchResult.Each(reflect.TypeOf(engine.Item{}))

	return result, nil
}
