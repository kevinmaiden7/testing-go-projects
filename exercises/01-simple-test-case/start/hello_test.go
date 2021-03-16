package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

func TestHelloEnglish(t *testing.T) {
	// Arrange
	name := ""
	language := "english"
	expected := "Hello World"

	actual, _ := Hello(name, language)	// Act

	assert.Equal(t, expected, actual) // Assert
}

func TestHelloSpanish(t *testing.T) {
	name := ""
	language := "spanish"
	expected := "Hola World"

	actual, _ := Hello(name, language)
	
	assert.Equal(t, expected, actual)
}

func TestHelloEnglishName(t *testing.T) {
	name := "Kevin"
	language := "english"
	expected := "Hello Kevin"

	actual, _ := Hello(name, language)
	
	assert.Equal(t, expected, actual)
}

func TestHelloGermanName(t *testing.T) {
	name := "Kevin"
	language := "german"
	expected := "Hallo Kevin"

	actual, _ := Hello(name, language)
	
	assert.Equal(t, expected, actual)
}

func TestHelloUnsupportedLanguage(t *testing.T) {
	name := ""
	language := ""
	expected := errors.New("need to provide a supported language")

	_, actual := Hello(name, language)
	
	assert.Equal(t, expected, actual)
}
