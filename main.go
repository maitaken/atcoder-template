// Template: https://github.com/maitaken/atcoder-template
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	BUFSIZE = 10000000
	MOD     = 1000000007
	INT_INF = math.MaxInt32
)

var rdr *bufio.Reader

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	solve()
}

func solve() {
}

func readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, s2i(v))
	}
	return slice
}

func readint() int {
	return s2i(readline())
}

func readint2() (int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1])
}

func readint3() (int, int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1]), s2i(lines[2])
}

func readint4() (int, int, int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1]), s2i(lines[2]), s2i(lines[3])
}

// For int
func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func mod(x, y int) int {
	m := x % y
	if m < 0 {
		return m + y
	}
	return m
}

func modpow(x, y int) int {
	ret := 1
	for ; y != 0; y >>= 1 {
		if y&1 == 1 {
			ret = mod((ret * x), MOD)
		}
		x = mod(x*x, MOD)
	}
	return ret
}

func min(values ...int) int {
	ret := INT_INF
	for _, v := range values {
		if ret > v {
			ret = v
		}
	}
	return ret
}

func max(values ...int) int {
	ret := -INT_INF
	for _, v := range values {
		if ret < v {
			ret = v
		}
	}
	return ret
}

func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}

func i2s(i int) string {
	return strconv.Itoa(i)
}

func gcd(v1, v2 int) int {
	if v1 > v2 {
		v1, v2 = v2, v1
	}
	for v1 != 0 {
		v1, v2 = v2%v1, v1
	}
	return v2
}

func lcm(v1, v2 int) int {
	return v1 * v2 / gcd(v1, v2)
}

func extgcd(a, b, c int) (int, int, int) {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a, 1, 0
	}
	d, x, y := extgcd(b, mod(a, b), c)
	return d, y, x - int(a/b)*y
}

func bit(size int) [][]bool {
	bSize := int(math.Pow(2, float64(size)))
	b := make([][]bool, bSize)
	for i := 0; i < bSize; i++ {
		b[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			if i>>j&1 == 1 {
				b[i][j] = true
			} else {
				b[i][j] = false
			}
		}
	}
	return b
}

/* ------------------------------------------------ */
/* Data stracture                                   */
/* ------------------------------------------------ */
type IntHeap []int

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// BIT Tree
type FenwickTree []int

func NewFenwickTree(size int) *FenwickTree {
	var t FenwickTree
	t = make([]int, size+1)
	return &t
}

func (t FenwickTree) add(index, v int) {
	for i := index + 1; i < len(t); i += (i & -i) {
		t[i] += v
	}
}

func (t FenwickTree) sum(index int) int {
	total := 0
	for i := index + 1; i != 0; i -= (i & -i) {
		total += t[i]
	}
	return total
}

type Comb struct {
	length int
	fac    []int
	inv    []int
	finv   []int
}

func NewComb() *Comb {
	return &Comb{
		length: 2,
		fac:    []int{1, 1},
		inv:    []int{1, 1},
		finv:   []int{1, 1},
	}
}

func (c Comb) calc(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}

	if c.length <= n {
		for i := c.length; i <= n; i++ {
			c.fac = append(c.fac, mod(c.fac[i-1]*i, MOD))
			c.inv = append(c.inv, MOD-mod(c.inv[mod(MOD, i)]*(MOD/i), MOD))
			c.finv = append(c.finv, mod(c.finv[i-1]*c.inv[i], MOD))
		}
	}
	return mod(c.fac[n]*mod(c.finv[k]*c.finv[n-k], MOD), MOD)
}
