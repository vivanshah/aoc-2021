package day

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Day21 struct {
	Dishes              []dish
	AllergenIngredients map[string]map[string]int
	AllergenOccurences  map[string]int
	AllIngredients      map[string]int
}

type dish struct {
	Ingredients []string
	Allergens   []string
}

// ReadFile reads a file
func (d *Day21) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day21.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	d.AllergenOccurences = map[string]int{}
	d.Dishes = []dish{}
	d.AllergenIngredients = map[string]map[string]int{}
	d.AllIngredients = map[string]int{}
	for _, l := range lines {
		l = strings.TrimSpace(l)
		p := strings.Split(l, "(contains ")
		iT := strings.TrimSpace(p[0])
		i := strings.Fields(iT)
		aP := strings.TrimRight(p[1], ")")
		a := strings.Split(aP, ", ")
		di := dish{Ingredients: []string{}, Allergens: []string{}}
		for _, ingredient := range i {
			di.Ingredients = append(di.Ingredients, ingredient)
			d.AllIngredients[ingredient] = d.AllIngredients[ingredient] + 1
		}
		for _, allergen := range a {
			di.Allergens = append(di.Allergens, allergen)
			d.AllergenOccurences[allergen] = d.AllergenOccurences[allergen] + 1
			if d.AllergenIngredients[allergen] == nil {
				d.AllergenIngredients[allergen] = map[string]int{}
			}
			for _, ingredient := range di.Ingredients {
				d.AllergenIngredients[allergen][ingredient] = d.AllergenIngredients[allergen][ingredient] + 1
			}
		}
		d.Dishes = append(d.Dishes, di)

	}

	fmt.Println("Read ", len(d.Dishes), " Dishes")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day21) Part1() {
	fmt.Println("Day 21 Part 1")
	potentials := map[string]bool{}
	for allergen, count := range d.AllergenOccurences {
		for ingredient, iC := range d.AllergenIngredients[allergen] {
			if iC == count {
				potentials[ingredient] = true
			}
		}
	}
	result := 0
	for i, o := range d.AllIngredients {
		if !potentials[i] {
			result = result + o
		}
	}
	fmt.Println(result)

}

// Part2 executes part 2 of of this day's puzzle
func (d *Day21) Part2() {
	fmt.Println("Day 21 Part 2")
	potentials := map[string]bool{}
	for allergen, count := range d.AllergenOccurences {
		for ingredient, iC := range d.AllergenIngredients[allergen] {
			if iC == count {
				potentials[ingredient] = true
			}
		}
	}
	dangerMap := map[string]string{}
	inert := map[string]bool{}
	for i := range d.AllIngredients {
		if !potentials[i] {
			inert[i] = true
		}
	}

	for allergen, count := range d.AllergenOccurences {
		for i := range d.AllIngredients {
			if inert[i] {
				continue
			}
			if count == d.AllergenIngredients[allergen][i] {
				dangerMap[allergen] = i
				inert[i] = true
				break
			}
		}
	}
	danger := kvList{}
	for k, v := range dangerMap {
		danger = append(danger, kv{Key: k, Value: v})
	}
	sort.Sort(danger)
	fmt.Print("\n")
	for _, dangerous := range danger {
		fmt.Print(dangerous.Key)
		fmt.Print(",")
	}
	fmt.Print("\n")
	for _, dangerous := range danger {
		fmt.Print(dangerous.Value)
		fmt.Print(",")
	}
	fmt.Print("\n")
}

type kv struct {
	Key   string
	Value string
}

type kvList []kv

func (p kvList) Len() int           { return len(p) }
func (p kvList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p kvList) Less(i, j int) bool { return p[i].Key < p[j].Key }
