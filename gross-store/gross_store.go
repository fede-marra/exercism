package gross


// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	measurements := map[string]int{}
	measurements["quarter_of_a_dozen"] = 3
	measurements["half_of_a_dozen"] = 6
	measurements["dozen"] = 12
	measurements["small_gross"] = 120
	measurements["gross"] = 144
	measurements["great_gross"] = 1728
	return measurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	cant, ok := units[unit]
	if !ok {
		return false
	}
	cantItem, ok := bill[item]
	if ok {
		bill[item] = cantItem + cant
	} else {
		bill[item] = cant
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	cant, ok := units[unit]
	if !ok {
		return false
	}
	_, ok = bill[item]
	if !ok {
		return false
	}
	switch {
	case bill[item] - cant < 0:
		return false
	case bill[item] - cant == 0:
		delete(bill, item)
		return true
	default:
		bill[item] = bill[item] - cant
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	cant, ok := bill[item]
	return cant, ok
}
