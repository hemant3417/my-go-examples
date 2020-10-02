//Package acdc provides a dum function
package acdc

//Sum Adds slice of integers
func Sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
