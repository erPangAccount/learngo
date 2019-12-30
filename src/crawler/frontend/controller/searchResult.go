package controller

import (
	"crawler/engine"
	"crawler/frontend/common"
	"crawler/frontend/model"
	model2 "crawler/model"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v6"
	"net/http"
	"reflect"
	"regexp"
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
		targetPage = 1
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
	q = replaceQueryString(q)
	from := (targetPage - 1) * result.PageInfo.PageSize

	query := elastic.NewQueryStringQuery(q)

	searchResult, err := s.client.Search("test").Query(query).From(from).Do(context.Background())
	if err != nil {
		return result, err
	}

	result.PageInfo.Total = searchResult.TotalHits()
	if targetPage > 1 {
		result.PageInfo.PrevPage = targetPage - 1
	}

	if int64(targetPage*result.PageInfo.PageSize) < result.PageInfo.Total {
		result.PageInfo.NextPage = targetPage + 1
	} else {
		result.PageInfo.NextPage = 0
	}

	for _, val := range searchResult.Each(reflect.TypeOf(engine.Item{})) {
		temp := val.(engine.Item)
		temp.DoType, _ = model2.FromJsonObj(temp.DoType)

		result.Items = append(result.Items, temp)
	}

	return result, nil
}

func replaceQueryString(q string) string {
	var queryStringRe = regexp.MustCompile(`([a-z]+):`)
	return queryStringRe.ReplaceAllString(q, "DoType.$1:")
}
