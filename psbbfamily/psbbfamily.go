package psbbfamily

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func PsbbFamily() {
	reader := bufio.NewReader(os.Stdin)
	for input := 1; input <= 3; input++ {
		fmt.Printf("\nInput #%d\n", input)
		fmt.Print("Input the number of families : ")
		nStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read input:", err)
			return
		}
		nStr = strings.TrimSpace(nStr)
		n, err := strconv.Atoi(nStr)
		if err != nil || n <= 0 {
			fmt.Println("Number of families must be a positive integer.")
			return
		}

		fmt.Print("Input the number of members in the family (separated by a space): ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read input:", err)
			return
		}
		line = strings.TrimSpace(line)
		membersStr := strings.Fields(line)

		if len(membersStr) != n {
			fmt.Println("Input must be equal with count of family")
			return
		}

		members := make([]int, 0, n)
		for i, s := range membersStr {
			num, err := strconv.Atoi(s)
			if err != nil || num <= 0 || num > 4 {
				fmt.Printf("Invalid member count at family %d: '%s' (must be 1-4)\n", i+1, s)
				return
			}
			members = append(members, num)
		}

		sort.Ints(members)

		bus := 0
		i, j := 0, n-1
		for i <= j {
			if i == j {
				bus++
				break
			}
			if members[i]+members[j] <= 4 {
				bus++
				i++
				j--
			} else {
				bus++
				j--
			}
		}

		fmt.Printf("Minimum bus required for input #%d is : %d\n", input, bus)
	}
}
