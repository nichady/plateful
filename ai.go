package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func init() {
	customsearchService, err := customsearch.NewService(context.Background(), option.WithAPIKey(os.Getenv("GOOGLE_SEARCH_API_KEY")))
	if err != nil {
		panic(err)
	}

	search = customsearchService
}

var (
	search   *customsearch.Service
	searchID = os.Getenv("GOOGLE_SEARCH_ENDPOINT")
	client   = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
)

func generateRecipe(g Generate) (Recipe, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: `You are an assistant to come up with food recipes. You should respond in JSON format, like this: {"name": "The name of the recipe", "body": "<everything else, including ingredients, instructions, etc...>"}`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: makePrompt(g),
				},
			},
		},
	)
	if err != nil {
		return Recipe{}, err
	}

	var recipe Recipe
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &recipe)
	if err != nil {
		return Recipe{}, err
	}

	result, err := search.Cse.List().Cx(searchID).Q(recipe.Name).SearchType("image").ImgType("photo").Do()
	if err != nil {
		return Recipe{}, err
	}

	recipe.Image = result.Items[0].Link

	return recipe, nil
}

func makePrompt(g Generate) string {
	str := ""

	if len(g.InludeIngredients) > 0 {
		str += fmt.Sprintf("I have these ingredients: %s.\n", g.InludeIngredients)
	}

	if len(g.ExcludeIngredients) > 0 {
		str += fmt.Sprintf("I don't have these ingredients: %s.\n", g.ExcludeIngredients)
	}

	str += fmt.Sprintf("Come up with a recipe in the style of one of these places: %s.\n", strings.Join(g.Locations, ", "))
	str += fmt.Sprintf("It should include these flavors: %s.\n", strings.Join(g.Flavors, ", "))
	str += fmt.Sprintf("It can be one of these: %s.\n", strings.Join(g.Times, ", "))
	str += "You can include other ingredients."

	return str
}
