package main

import (
	"Go-Tutorials/go-rest-guide/pkg/recipes"
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readTestData(t *testing.T, name string) []byte {
	t.Helper()
	content, err := os.ReadFile("../../testdata/" + name)
	if err != nil {
		t.Errorf("Coud not read %v", name)
	}

	return content
}

func TestRecipeHandlerCRUD_Integration(t *testing.T) {

	// Create a MemStore and Recipe Handler
	store := recipes.NewMemStore()
	recipesHandler := NewRecipesHandler(store)

	// Test data
	hamAndCheese := readTestData(t, "ham_and_cheese_recipe.json")
	hamAndCheeseReader := bytes.NewReader(hamAndCheese)

	hamAndCheeseWithButter := readTestData(t, "ham_and_cheese_with_butter_recipe.json")
	hamAndCheeseWithButterReader := bytes.NewReader(hamAndCheeseWithButter)

	// CREATE - add a new recipe
	req := httptest.NewRequest(http.MethodPost, "/recipes", hamAndCheeseReader)
	w := httptest.NewRecorder()
	recipesHandler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode)

	saved, _ := store.List()
	assert.Len(t, saved, 1)

	// GET - find the record you just added
	req = httptest.NewRequest(http.MethodGet, "/recipes/ham-and-cheese-toasties", nil)
	w = httptest.NewRecorder()
	recipesHandler.ServeHTTP(w, req)

	res = w.Result()
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	assert.JSONEq(t, string(hamAndCheese), string(data))

	// UPDATE - add butter to ham and cheese recipe
	req = httptest.NewRequest(http.MethodPut, "/recipes/ham-and-cheese-toasties", hamAndCheeseWithButterReader)
	w = httptest.NewRecorder()
	recipesHandler.ServeHTTP(w, req)

	res = w.Result()
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode)

	updateHamAndCheese, err := store.Get("ham-and-cheese-toasties")
	assert.NoError(t, err)

	assert.Contains(t, updateHamAndCheese.Ingredients, recipes.Ingredient{Name: "butter"})

	// DELETE - remove the ham and cheese recipe
	req = httptest.NewRequest(http.MethodDelete, "/recipes/ham-and-cheese-toasties", nil)
	w = httptest.NewRecorder()
	recipesHandler.ServeHTTP(w, req)

	res = w.Result()
	defer res.Body.Close()
	assert.Equal(t, 200, res.StatusCode)

	saved, _ = store.List()
	assert.Len(t, saved, 0)

}
