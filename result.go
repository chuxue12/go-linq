package linq

import (
	"math"
	"reflect"
)

// All determines whether all elements of a collection satisfy a condition.
func (q Query) All(predicate func(interface{}) bool) bool {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if !predicate(item) {
			return false
		}
	}

	return true
}

// AllT is the typed version of All.
//
// NOTE: All method has better performance than AllT
//
// predicateFn is of a type "func(TSource) bool"
func (q Query) AllT(predicateFn interface{}) bool {

	predicateGenericFunc, err := newGenericFunc(
		"AllT", "predicateFn", predicateFn,
		simpleParamValidator(newElemTypeSlice(new(genericType)), newElemTypeSlice(new(bool))),
	)
	if err != nil {
		panic(err)
	}
	predicateFunc := func(item interface{}) bool {
		return predicateGenericFunc.Call(item).(bool)
	}

	return q.All(predicateFunc)
}

// Any determines whether any element of a collection exists.
func (q Query) Any() bool {
	_, ok := q.Iterate()()
	return ok
}

// AnyWith determines whether any element of a collection satisfies a condition.
func (q Query) AnyWith(predicate func(interface{}) bool) bool {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			return true
		}
	}

	return false
}

// AnyWithT is the typed version of AnyWith.
//
// NOTE: AnyWith method has better performance than AnyWithT
//
// predicateFn is of a type "func(TSource) bool"
func (q Query) AnyWithT(predicateFn interface{}) bool {

	predicateGenericFunc, err := newGenericFunc(
		"AnyWithT", "predicateFn", predicateFn,
		simpleParamValidator(newElemTypeSlice(new(genericType)), newElemTypeSlice(new(bool))),
	)
	if err != nil {
		panic(err)
	}

	predicateFunc := func(item interface{}) bool {
		return predicateGenericFunc.Call(item).(bool)
	}

	return q.AnyWith(predicateFunc)
}

// Average computes the average of a collection of numeric values.
func (q Query) Average() (r float64) {
	next := q.Iterate()
	item, ok := next()
	if !ok {
		return math.NaN()
	}

	n := 1
	switch item.(type) {
	case int, int8, int16, int32, int64:
		conv := getIntConverter(item)
		sum := conv(item)

		for item, ok = next(); ok; item, ok = next() {
			sum += conv(item)
			n++
		}

		r = float64(sum)
	case uint, uint8, uint16, uint32, uint64:
		conv := getUIntConverter(item)
		sum := conv(item)

		for item, ok = next(); ok; item, ok = next() {
			sum += conv(item)
			n++
		}

		r = float64(sum)
	default:
		conv := getFloatConverter(item)
		r = conv(item)

		for item, ok = next(); ok; item, ok = next() {
			r += conv(item)
			n++
		}
	}

	return r / float64(n)
}

// Contains determines whether a collection contains a specified element.
func (q Query) Contains(value interface{}) bool {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if item == value {
			return true
		}
	}

	return false
}

// Count returns the number of elements in a collection.
func (q Query) Count() (r int) {
	next := q.Iterate()

	for _, ok := next(); ok; _, ok = next() {
		r++
	}

	return
}

// CountWith returns a number that represents how many elements
// in the specified collection satisfy a condition.
func (q Query) CountWith(predicate func(interface{}) bool) (r int) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			r++
		}
	}

	return
}

// CountWithT is the typed version of CountWith.
//
// NOTE: CountWith method has better performance than CountWithT
//
// predicateFn is of a type "func(TSource) bool"
func (q Query) CountWithT(predicateFn interface{}) int {

	predicateGenericFunc, err := newGenericFunc(
		"CountWithT", "predicateFn", predicateFn,
		simpleParamValidator(newElemTypeSlice(new(genericType)), newElemTypeSlice(new(bool))),
	)
	if err != nil {
		panic(err)
	}

	predicateFunc := func(item interface{}) bool {
		return predicateGenericFunc.Call(item).(bool)
	}

	return q.CountWith(predicateFunc)
}

// First returns the first element of a collection.
func (q Query) First() interface{} {
	item, _ := q.Iterate()()
	return item
}

// FirstWith returns the first element of a collection that satisfies
// a specified condition.
func (q Query) FirstWith(predicate func(interface{}) bool) interface{} {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			return item
		}
	}

	return nil
}

// FirstWithT is the typed version of FirstWith.
//
// NOTE: FirstWith method has better performance than FirstWithT
//
// predicateFn is of a type "func(TSource) bool"
func (q Query) FirstWithT(predicateFn interface{}) interface{} {

	predicateGenericFunc, err := newGenericFunc(
		"FirstWithT", "predicateFn", predicateFn,
		simpleParamValidator(newElemTypeSlice(new(genericType)), newElemTypeSlice(new(bool))),
	)
	if err != nil {
		panic(err)
	}

	predicateFunc := func(item interface{}) bool {
		return predicateGenericFunc.Call(item).(bool)
	}

	return q.FirstWith(predicateFunc)
}

// Last returns the last element of a collection.
func (q Query) Last() (r interface{}) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		r = item
	}

	return
}

// LastWith returns the last element of a collection that satisfies
// a specified condition.
func (q Query) LastWith(predicate func(interface{}) bool) (r interface{}) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			r = item
		}
	}

	return
}

// LastWithT is the typed version of LastWith.
//
// NOTE: LastWith method has better performance than LastWithT
//
// predicateFn is of a type "func(TSource) bool"
func (q Query) LastWithT(predicateFn interface{}) interface{} {

	predicateGenericFunc, err := newGenericFunc(
		"LastWithT", "predicateFn", predicateFn,
		simpleParamValidator(newElemTypeSlice(new(genericType)), newElemTypeSlice(new(bool))),
	)
	if err != nil {
		panic(err)
	}

	predicateFunc := func(item interface{}) bool {
		return predicateGenericFunc.Call(item).(bool)
	}

	return q.LastWith(predicateFunc)
}

// Max returns the maximum value in a collection of values.
func (q Query) Max() (r interface{}) {
	next := q.Iterate()
	item, ok := next()
	if !ok {
		return nil
	}

	compare := getComparer(item)
	r = item

	for item, ok := next(); ok; item, ok = next() {
		if compare(item, r) > 0 {
			r = item
		}
	}

	return
}

// Min returns the minimum value in a collection of values.
func (q Query) Min() (r interface{}) {
	next := q.Iterate()
	item, ok := next()
	if !ok {
		return nil
	}

	compare := getComparer(item)
	r = item

	for item, ok := next(); ok; item, ok = next() {
		if compare(item, r) < 0 {
			r = item
		}
	}

	return
}

// Results iterates over a collection and returnes slice of interfaces
func (q Query) Results() (r []interface{}) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		r = append(r, item)
	}

	return
}

// SequenceEqual determines whether two collections are equal.
func (q Query) SequenceEqual(q2 Query) bool {
	next := q.Iterate()
	next2 := q2.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		item2, ok2 := next2()
		if !ok2 || item != item2 {
			return false
		}
	}

	_, ok2 := next2()
	return !ok2
}

// Single returns the only element of a collection, and nil
// if there is not exactly one element in the collection.
func (q Query) Single() interface{} {
	next := q.Iterate()
	item, ok := next()
	if !ok {
		return nil
	}

	_, ok = next()
	if ok {
		return nil
	}

	return item
}

// SingleWith returns the only element of a collection that satisfies
// a specified condition, and nil if more than one such element exists.
func (q Query) SingleWith(predicate func(interface{}) bool) (r interface{}) {
	next := q.Iterate()
	found := false

	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			if found {
				return nil
			}

			found = true
			r = item
		}
	}

	return
}

// SingleWithT is the typed version of SingleWith.
//
// NOTE: SingleWith method has better performance than SingleWithT
//
// predicateFn is of a type "func(TSource) bool"
func (q Query) SingleWithT(predicateFn interface{}) interface{} {
	predicateGenericFunc, err := newGenericFunc(
		"SingleWithT", "predicateFn", predicateFn,
		simpleParamValidator(newElemTypeSlice(new(genericType)), newElemTypeSlice(new(bool))),
	)
	if err != nil {
		panic(err)
	}

	predicateFunc := func(item interface{}) bool {
		return predicateGenericFunc.Call(item).(bool)
	}

	return q.SingleWith(predicateFunc)
}

// SumInts computes the sum of a collection of numeric values.
//
// Values can be of any integer type: int, int8, int16, int32, int64.
// The result is int64. Method returns zero if collection contains no elements.
func (q Query) SumInts() (r int64) {
	next := q.Iterate()
	item, ok := next()
	if !ok {
		return 0
	}

	conv := getIntConverter(item)
	r = conv(item)

	for item, ok = next(); ok; item, ok = next() {
		r += conv(item)
	}

	return
}

// SumUInts computes the sum of a collection of numeric values.
//
// Values can be of any unsigned integer type: uint, uint8, uint16, uint32, uint64.
// The result is uint64. Method returns zero if collection contains no elements.
func (q Query) SumUInts() (r uint64) {
	next := q.Iterate()
	item, ok := next()
	if !ok {
		return 0
	}

	conv := getUIntConverter(item)
	r = conv(item)

	for item, ok = next(); ok; item, ok = next() {
		r += conv(item)
	}

	return
}

// SumFloats computes the sum of a collection of numeric values.
//
// Values can be of any float type: float32 or float64. The result is float64.
// Method returns zero if collection contains no elements.
func (q Query) SumFloats() (r float64) {
	next := q.Iterate()
	item, ok := next()
	if !ok {
		return 0
	}

	conv := getFloatConverter(item)
	r = conv(item)

	for item, ok = next(); ok; item, ok = next() {
		r += conv(item)
	}

	return
}

// ToChannel iterates over a collection and outputs each element
// to a channel, then closes it.
func (q Query) ToChannel(result chan<- interface{}) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		result <- item
	}

	close(result)
}

// ToMap iterates over a collection and populates result map with elements.
// Collection elements have to be of KeyValue type to use this method.
// To populate a map with elements of different type use ToMapBy method.
// ToMap doesn't empty the result map before populating it.
func (q Query) ToMap(result interface{}) {
	q.ToMapBy(
		result,
		func(i interface{}) interface{} {
			return i.(KeyValue).Key
		},
		func(i interface{}) interface{} {
			return i.(KeyValue).Value
		})
}

// ToMapBy iterates over a collection and populates the result map with elements.
// Functions keySelector and valueSelector are executed for each element of the collection
// to generate key and value for the map. Generated key and value types must be assignable
// to the map's key and value types.
// ToMapBy doesn't empty the result map before populating it.
func (q Query) ToMapBy(
	result interface{},
	keySelector func(interface{}) interface{},
	valueSelector func(interface{}) interface{},
) {
	res := reflect.ValueOf(result)
	m := reflect.Indirect(res)
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		key := reflect.ValueOf(keySelector(item))
		value := reflect.ValueOf(valueSelector(item))

		m.SetMapIndex(key, value)
	}

	res.Elem().Set(m)
}

// ToMapByT is the typed version of ToMapBy.
//
// NOTE: ToMapBy method has better performance than ToMapByT
//
// keySelectorFn is of a type "func(TSource)TKey"
//
// valueSelectorFn is of a type "func(TSource)TValue"
func (q Query) ToMapByT(result interface{}, keySelectorFn interface{}, valueSelectorFn interface{}) {
	keySelectorGenericFunc, err := newGenericFunc(
		"ToMapByT", "keySelectorFn", keySelectorFn,
		simpleParamValidator(newElemTypeSlice(new(genericType)), newElemTypeSlice(new(genericType))),
	)
	if err != nil {
		panic(err)
	}

	keySelectorFunc := func(item interface{}) interface{} {
		return keySelectorGenericFunc.Call(item)
	}

	valueSelectorGenericFunc, err := newGenericFunc(
		"ToMapByT", "valueSelectorFn", valueSelectorFn,
		simpleParamValidator(newElemTypeSlice(new(genericType)), newElemTypeSlice(new(genericType))),
	)
	if err != nil {
		panic(err)
	}

	valueSelectorFunc := func(item interface{}) interface{} {
		return valueSelectorGenericFunc.Call(item)
	}

	q.ToMapBy(result, keySelectorFunc, valueSelectorFunc)
}

// ToSlice iterates over a collection, copy to result slice the collection elements,
// appends items if initial capacity is not enough and returns subslice if initial capacity is excessive.
// Collection elements must be assignable to the slice's element type.
func (q Query) ToSlice(result interface{}) {
	res := reflect.ValueOf(result)
	slice := reflect.Indirect(res)
	next := q.Iterate()
	index := 0
	for item, ok := next(); ok; item, ok = next() {
		if index >= slice.Len() {
			slice = grow(slice, 1)
		}
		slice.Index(index).Set(reflect.ValueOf(item))
		index++
	}

	res.Elem().Set(slice.Slice(0, index))
}

// grow grows the slice s so that it can hold extra more values, allocating
// more capacity if needed.
func grow(s reflect.Value, extra uint) reflect.Value {
	len := s.Len()
	newLen := len + int(extra)
	cap := s.Cap()
	if newLen <= cap {
		return s.Slice(0, newLen)
	}
	if cap == 0 {
		cap = int(extra)
	} else {
		for cap < newLen {
			cap *= 2
		}
	}
	newSlice := reflect.MakeSlice(s.Type(), newLen, cap)
	reflect.Copy(newSlice, s)
	return newSlice
}
