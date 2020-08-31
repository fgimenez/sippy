package internal

import bugsv1 "github.com/openshift/sippy/pkg/apis/bugs/v1"

// These types are used to decode information from bugzilla, but we don't want to expose these for anyone else.

type Search struct {
	Results Results `json:"results"`
}

// search string is the key
type Results map[string]Result

type Result struct {
	Matches []Match `json:"matches"`
}

type Match struct {
	Bug bugsv1.Bug `json:"bugInfo"`
}
