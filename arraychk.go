package main
import "fmt"


func arraylist() {
	
	var i [25] int
	fmt.Println("Before :", i)

	fmt.Println("Arrays\n\n")

	for j := 0; j < 25; j++ {
		i[j] = j * 10
	}

	fmt.Println("After :", i)
	
	arr_len := len(i)
	for j := 0; j < arr_len; j++ {
		if (i[j]%2 == 0) {
			fmt.Println("Even ~> ", i[j])
		} else {
			fmt.Println("Odd  ~> ", i[j])
		}
	}
}

func arr2slice() {	
	fmt.Println("\n\nConverting array to slice\n\n")

	var i [25] int
	s := make([]int, 12)
	fmt.Println(copy(s, i[10:]))
	fmt.Println(s)

	s = append(s, 88, 99, 999)
	fmt.Println(s)

	fmt.Println("\n\nMaps\n\n")

	mps := make(map[int]string)
	mps[11] = "Sun"
	mps[12] = "Mon"
	mps[13] = "Tue"
	mps[14] = "Wed"
	mps[15] = "Thu"
	mps[16] = "Fri"
	mps[17] = "Sat"

	fmt.Println("Lenth of Maps :", len(mps))
	fmt.Println(mps)
	if (mps[17] == "Sat") {
		fmt.Println("Its Weekend :", mps[17])
	} 
	
	delete (mps, 14)
	fmt.Println(mps)
}

func ranging() int {
	var nums [10] int

	for x :=0; x < 10; x++ {
		nums[x] = x * 3
	}
		fmt.Println("\n\n", nums)
	
	total := 0
	for _, n := range nums {
		fmt.Println(n, total)
		total += n
	}
	
	fmt.Println("Totol : ", total)
	
return total
}

func main() {
	
	arraylist()
        arr2slice()
        totalval := ranging()
	fmt.Println("Tota returned : ", totalval)
}

