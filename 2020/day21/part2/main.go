package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

/*
 *   --- Day 21: Allergen Assessment ---
 *            --- Part One ---
 *
 *   https://adventofcode.com/2020/day/21#part2
 */

type food struct {
	ingredients []string
	allergens   []string
}

type ingredient struct {
	name     string
	allergen string
}

func intersection(list1, list2 []string) []string {
	intersection := make([]string, 0, 0)
	m := make(map[string]int)
	for i, e := range list1 {
		m[e] = i
	}
	for _, e := range list2 {
		i, exists := m[e]
		if exists {
			intersection = append(intersection, list1[i])
		}
	}
	return intersection
}

func contains(s string, list []string) bool {
	for _, e := range list {
		if e == s {
			return true
		}
	}
	return false
}

func removeIngredient(name string, foods []food) {
	for j, food := range foods {
		for i, ingredient := range food.ingredients {
			if name == ingredient {
				foods[j].ingredients = append(food.ingredients[:i], food.ingredients[i+1:]...)
				break
			}
		}
	}
}

func removeAllergen(name string, foods []food) {
	for j, food := range foods {
		for i, allergen := range food.allergens {
			if name == allergen {
				foods[j].allergens = append(food.allergens[:i], food.allergens[i+1:]...)
				break
			}
		}
	}
}

func main() {
	// Puzzle input
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	foods := make([]food, 0, 0)
	for scanner.Scan() {
		f := food{}

		parts := strings.Split(scanner.Text(), " (contains ")
		ingredientsStr := parts[0]
		f.ingredients = strings.Split(ingredientsStr, " ")

		if len(parts) > 1 {
			allergensStr := parts[1][:len(parts[1])-1]
			f.allergens = strings.Split(allergensStr, ", ")
		}

		foods = append(foods, f)
	}

	allAllergens := make(map[string]struct{})
	for _, f := range foods {
		for _, i := range f.allergens {
			allAllergens[i] = struct{}{}
		}
	}

	allergensMap := make(map[string]string)
	for len(allergensMap) < len(allAllergens) {
		for allergen := range allAllergens {
			var inter []string
			for i := 0; i < len(foods); i++ {
				if contains(allergen, foods[i].allergens) {
					if len(inter) == 0 {
						inter = foods[i].ingredients
					} else {
						inter = intersection(foods[i].ingredients, inter)
					}
				}
			}
			if len(inter) == 1 {
				ingredient := inter[0]
				allergensMap[ingredient] = allergen
				removeIngredient(ingredient, foods)
				removeAllergen(allergen, foods)
			}
		}
	}
	for ingredient, allergen := range allergensMap {
		log.Println(ingredient, "contains", allergen)
	}

	canonicalDangerousIngredientList := ""
	{
		dangerousIngredients := make([]ingredient, 0, 0)
		for i, a := range allergensMap {
			dangerousIngredients = append(dangerousIngredients, ingredient{i, a})
		}
		sort.Slice(dangerousIngredients, func(i, j int) bool {
			return strings.Compare(dangerousIngredients[i].allergen, dangerousIngredients[j].allergen) < 0
		})

		for _, ingredient := range dangerousIngredients {
			canonicalDangerousIngredientList += ingredient.name + ","
		}
		canonicalDangerousIngredientList = canonicalDangerousIngredientList[:len(canonicalDangerousIngredientList)-1]
	}
	log.Println()
	log.Println("the canonical dangerous ingredient list is", canonicalDangerousIngredientList)
}
