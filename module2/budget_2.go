package module2

import (
	"errors"
	"time"
)

// START Initial code

// Budget stores information for a budget
type Budget struct {
	Max   float32
	Items []Item
}

// Item stores information for an item
type Item struct {
	Description string
	Price       float32
}

var report map[time.Month]*Budget

// InitializeReport is used to initialize/reset the report map
func InitializeReport() {
	report = make(map[time.Month]*Budget)
}

func init() {
	InitializeReport()
}

// CurrentCost returns how much we have added
// to our current budget.
func (b Budget) CurrentCost() float32 {
	var sum float32
	for _, item := range b.Items {
		sum += item.Price
	}
	return sum
}

// errDoesNotFitBudget when item does not fit budget
var errDoesNotFitBudget = errors.New("Item does not fit the budget")

// errReportIsFull when an attemp is made to add a new budget
// to an already full report
var errReportIsFull = errors.New("Report is full")

// errDuplicateEntry when an attemp is made to add a budget
// to an existing month
var errDuplicateEntry = errors.New("Cannot add duplicate entry")

// END Initial code

// START Project code

// AddItem adds a new item to the budget
func (b *Budget) AddItem(description string, price float32) error {
	if (b.CurrentCost() + price) > b.Max {
		return errDoesNotFitBudget
	}
	newItem := Item{Description: description, Price: price}
	b.Items = append(b.Items, newItem)
	return nil
}

// RemoveItem removes an item matching the description
func (b *Budget) RemoveItem(description string) {
	for i := range b.Items {
		if b.Items[i].Description == description {
			b.Items = append(b.Items[:i], b.Items[i+1:]...)
			break
		}
	}
}

// CreateBudget creates a new budget for a given time and
// with a set max
func CreateBudget(month time.Month, max float32) (*Budget, error) {
	var newBudget *Budget
	if len(report) >= 12 {
		return nil, errReportIsFull
	}
	if _, hasEntry := report[month]; hasEntry {
		return nil, errDuplicateEntry
	}
	newBudget = &Budget{Max: max}
	report[month] = newBudget
	return newBudget, nil
}

// GetBudget returns the budget for a given month
func GetBudget(month time.Month) *Budget {
	if budget, ok := report[month]; ok {
		return budget
	}
	return nil
}

// END Project code
