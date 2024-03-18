package handler

import (
	"github.com/gofrs/uuid/v5"
	"net/http"
	"strings"
)

func getUuidFromRequest(req *http.Request) (uuid.UUID, error) {
	path := req.URL.Path
	parts := strings.Split(path, "/")

	idString := parts[len(parts)-1]
	id, err := uuid.FromString(idString)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func getTwoUuidFromRequest(req *http.Request) (uuid.UUID, uuid.UUID, error) {
	path := req.URL.Path
	parts := strings.Split(path, "/")

	idString := parts[len(parts)-1]
	id, err := uuid.FromString(idString)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	id2String := parts[len(parts)-2]
	id2, err := uuid.FromString(id2String)
	if err != nil {
		return uuid.Nil, uuid.Nil, err
	}

	return id, id2, nil
}

func getUuidStringFromRequest(req *http.Request) (uuid.UUID, string, error) {
	path := req.URL.Path
	parts := strings.Split(path, "/")

	idString := parts[len(parts)-2]
	id, err := uuid.FromString(idString)
	if err != nil {
		return uuid.Nil, "", err
	}
	str := parts[len(parts)-1]

	return id, str, nil
}
