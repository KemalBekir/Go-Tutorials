package main

const N = 100000

type Node struct {
	suff int
	l    int
	c    [26]int
	cnt  int
}

var a [N + 1]byte
var b [N + 2]Node

func getSuff(i int, x int) int {
	for i-1-b[x].l < 0 || a[i-1-b[x].l] != a[i] {
		x = b[x].suff
	}
	return x
}

func palindromicBorder(s string) int32 {
	for i := 0; i < len(s); i++ {
		a[i] = s[i]
	}
	a[len(s)] = 0

	b[0].suff = 1
	b[0].l = 0
	b[1].suff = 1
	b[1].l = -1

	x := 1
	y := 2

	for i := 0; a[i] != 0; i++ {
		x = getSuff(i, x)
		if b[x].c[a[i]-'a'] == 0 {
			b[y].l = b[x].l + 2
			b[y].suff = b[getSuff(i, b[x].suff)].c[a[i]-'a']
			b[y].cnt = 0
			b[x].c[a[i]-'a'] = y
			y++
		}
		x = b[x].c[a[i]-'a']
		b[x].cnt++
	}

	for i := y - 1; i >= 0; i-- {
		b[b[i].suff].cnt += b[i].cnt
	}

	var ans int64 = 0
	for i := 2; i < y; i++ {
		c := b[i].cnt
		ans += int64(c-1) * int64(c) / 2
	}

	return int32(ans % 1000000007)
}