package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestGetRepositories(t *testing.T) {
	data := `[{"id":274146961,"name":"aerostat","description":"","stargazers_count":0,"owner":{"id":58124,"login":"Chau"}},{"id":169416390,"name":"blog_app","description":"","stargazers_count":0,"owner":{"id":58124,"login":"Chau"}},{"id":157074489,"name":"brackets_validation","description":"","stargazers_count":0,"owner":{"id":58124,"login":"Chau"}},{"id":157812554,"name":"celery","description":"Distributed Task Queue (development branch)","stargazers_count":0,"owner":{"id":58124,"login":"Chau"}},{"id":153463439,"name":"currency_converse","description":"","stargazers_count":0,"owner":{"id":58124,"login":"Chau"}}]`
	textBytes := []byte(data)
	var want []Repository
	json.Unmarshal(textBytes, &want)
	got, _ := GetRepositoriesByUserId("chau", "")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want %v but got %v", want, got)
	}
}

func TestGetRepositories2(t *testing.T) {
	data := `[{"id":274146961,"name":"aerostat","description":"","stargazers_count":0,"owner":{"id":58124,"login":"Chau"}}]`
	textBytes := []byte(data)
	var want []Repository
	json.Unmarshal(textBytes, &want)
	got, _ := GetRepositoriesByUserId("chau", "aerostat")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want %v but got %v", want, got)
	}
}

// test case userId not found
func TestGetRepositories3(t *testing.T) {
	_, code := GetRepositoriesByUserId("xxxxx1231", "")
	if code != http.StatusNotFound {
		t.Errorf("Want status code %v but got %v", http.StatusNotFound, code)
	}
}
