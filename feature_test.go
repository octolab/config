package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/config"
)

func TestFeature_String(t *testing.T) {
	tests := map[string]struct {
		feature  Feature
		expected string
	}{
		"enabled":  {Feature{Name: "feature", Enabled: true}, "feature=true"},
		"disabled": {Feature{Name: "feature", Enabled: false}, "feature=false"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.feature.String())
		})
	}
}

func TestFeatures_String(t *testing.T) {
	tests := map[string]struct {
		features Features
		expected string
	}{
		"nil":   {nil, "-"},
		"empty": {make(Features, 0), "-"},
		"with data": {
			Features{{Name: "featureA", Enabled: true}, {Name: "featureB"}},
			"featureA=true, featureB=false",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.features.String())
		})
	}
}

func TestFeatures_FindByID(t *testing.T) {
	var id = [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	tests := map[string]struct {
		features Features
		id       [16]byte
		expected Feature
	}{
		"nil":   {nil, id, Feature{}},
		"empty": {make(Features, 0), id, Feature{}},
		"with data": {
			Features{
				{ID: id, Name: "featureA", Enabled: true},
				{Name: "featureB"},
			},
			id,
			Feature{ID: id, Name: "featureA", Enabled: true},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.features.FindByID(test.id))
		})
	}
}

func TestFeatures_FindByName(t *testing.T) {
	tests := map[string]struct {
		features Features
		name     string
		expected Feature
	}{
		"nil":   {nil, "feature", Feature{}},
		"empty": {make(Features, 0), "feature", Feature{}},
		"with data": {
			Features{
				{Name: "featureA", Enabled: true},
				{Name: "featureB"},
			},
			"featureA",
			Feature{Name: "featureA", Enabled: true},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.features.FindByName(test.name))
		})
	}
}
