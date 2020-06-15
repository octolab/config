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
			Features{{"featureA", true}, {"featureB", false}},
			"featureA=true, featureB=false",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.features.String())
		})
	}
}
