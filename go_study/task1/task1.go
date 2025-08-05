package main

func main() {

}

func change(nums []int) {

	nums = []int{1, 2, 3, 4, 5}
}

// 只出现一次的数字
func onceNum(nums []int) int {
	s := 0
	for _, num := range nums {
		s = s ^ num
	}
	return s
}

// 回文数
func isReverseNum(num int) bool {
	//121   13331
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}

	reverse := 0
	for reverse < num {
		reverse *= 10
		reverse += num % 10
		num /= 10
	}
	return reverse == num || reverse/10 == num
}

// 有效的括号
func okString(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	strs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}

	// ([])
	for i := 0; i < len(s); i++ {

		if p, ok := strs[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != p {
				return false
			}
			//和栈顶存在配对 栈顶弹出
			stack = stack[:len(stack)-1]
		} else {
			//入栈
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		pre = compare(pre, strs[i])
	}
	return pre
}

func compare(str string, str1 string) string {
	pre := ""
	for i := 0; i < len(str) && i < len(str1); i++ {
		if str[i] != str1[i] {
			break
		}
		pre = str[:i+1]
	}
	return pre
}

// 加一
func plusOne(digits []int) []int {
	return doPlusOne(len(digits)-1, digits)
}

func doPlusOne(i int, digits []int) []int {

	if i >= 0 {
		//不用动态扩展一位
		if digits[i]+1 < 10 {
			digits[i] = digits[i] + 1
		} else {
			digits[i] = 0
			digits = doPlusOne(i-1, digits)

		}
	} else {
		//需要动态扩展一位
		digitsPlus := []int{}
		digitsPlus = append(digitsPlus, 0)
		for j := 0; j < len(digits); j++ {
			digitsPlus = append(digitsPlus, digits[j])
		}
		digits = doPlusOne(0, digitsPlus)
	}

	return digits
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	index := 1
	base := nums[0]
	for i := 1; i < len(nums); i++ {
		if base == nums[i] {
			continue
		}
		nums[index] = nums[i]
		base = nums[i]
		index++
	}
	return index
}

// 两数之和
func findIndex(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		t, ok := m[target-num]
		if ok {
			return []int{i, t}
		}
		m[num] = i
	}
	return nil
}
