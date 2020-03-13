// Code generated by 'gofp'. DO NOT EDIT.
package info
import "sync" 

func Map(f func(Info) Info, list []Info) []Info {
	if f == nil {
		return []Info{}
	}
	newList := make([]Info, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func Filter(f func(Info) bool, list []Info) []Info {
	if f == nil {
		return []Info{}
	}
	var newList []Info
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterPtr(f func(*Info) bool, list []*Info) []*Info {
	if f == nil {
		return []*Info{}
	}
	var newList []*Info
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func Remove(f func(Info) bool, list []Info) []Info {
	if f == nil {
		return []Info{}
	}
	var newList []Info
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func Some(f func(Info) bool, list []Info) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func Every(f func(Info) bool, list []Info) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func DropWhile(f func(Info) bool, list []Info) []Info {
	if f == nil {
		return []Info{}
	}
	var newList []Info
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]Info, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func TakeWhile(f func(Info) bool, list []Info) []Info {
	if f == nil {
		return []Info{}
	}
	var newList []Info
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

func PMap(f func(Info) Info, list []Info) []Info {
	if f == nil {
		return []Info{}
	}

	ch := make(chan map[int]Info)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Info, i int, v Info) {
			defer wg.Done()
			ch <- map[int]Info{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Info, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

func FilterMap(fFilter func(Info) bool, fMap func(Info) Info, list []Info) []Info {
	if fFilter == nil || fMap == nil {
		return []Info{}
	}
	var newList []Info
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func Rest(l []Info) []Info {
	if l == nil {
		return []Info{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []Info{}
	}

	newList := make([]Info, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

func Reduce(f func(Info, Info) Info, list []Info, initializer ...Info) Info {
	var init Info 
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}
	
	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return Reduce(f, list[1:], r)
}

// DropLast drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLast(list []Info) []Info {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []Info{}
	}

	newList := make([]Info, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}


// MapInfoStr takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapInfoStr(f func(Info) string, list []Info) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapInfoStr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: Info output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapInfoStr(f func(Info) string, list []Info) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v Info) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapInfoStr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - Info and returns true/false.
//	2. Function: takes Info as argument and returns string
// 	3. List of type Info
//
// Returns:
//	New List of type string
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapInfoStr(fFilter func(Info) bool, fMap func(Info) string, list []Info) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapStrInfo takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapStrInfo(f func(string) Info, list []string) []Info {
	if f == nil {
		return []Info{}
	}
	newList := make([]Info, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapStrInfo applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: Info
//	2. List
//
// Returns
//	New List of type Info
//	Empty list if all arguments are nil or either one is nil
func PMapStrInfo(f func(string) Info, list []string) []Info {
	if f == nil {
		return []Info{}
	}

	ch := make(chan map[int]Info)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Info, i int, v string) {
			defer wg.Done()
			ch <- map[int]Info{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Info, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapStrInfo filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - string and returns true/false.
//	2. Function: takes string as argument and returns Info
// 	3. List of type string
//
// Returns:
//	New List of type Info
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapStrInfo(fFilter func(string) bool, fMap func(string) Info, list []string) []Info {
	if fFilter == nil || fMap == nil {
		return []Info{}
	}
	var newList []Info
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

