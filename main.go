package main

import (
	"fmt"
	"math"
	"time"
)

func root3(x int64) int64 {
	return int64(math.Cbrt(float64(x))) + 1
}

func root2(x int64) int64 {
	return int64(math.Sqrt(float64(x))) + 1
}

func eratosthenes(x32, x3 int64) *[]bool {
	table := make([]bool, x32+1)
	table[0] = true
	table[1] = true

	for i := int64(0); i <= x32; i++ {
		if !table[i] && i <= x3 {
			for j := i * i; j <= x32; j += i {
				table[j] = true
			}
		}
	}
	return &table
}

func calc_pi_x32(prime_x32 *[]bool, x32 int64) *[]int32 {
	pi_x32 := make([]int32, x32+1)
	count := int64(0)
	for i := int64(0); i <= x32; i++ {
		if !(*prime_x32)[i] {
			count++
		}
		pi_x32[i] = int32(count)
	}
	return &pi_x32
}

func bit_add(bit *[]int32, i, x, size int64) {
	for ; i <= size; i += i & (-i) {
		(*bit)[i] += int32(x)
	}
}

func bit_sum(bit *[]int32, i int64) int64 {
	s := int64(0)
	for ; i > 0; i -= i & (-i) {
		s += int64((*bit)[i])
	}
	return s
}

func calc_phi(prime_x32 *[]bool, x, x3, x32, pi_x3 int64) int64 {
	mobius := make([]int8, x3+1)
	min_prime := make([]int32, x3+1)
	prime := make([]int32, 0, x3+1)
	prime = append(prime, 2)

	for i := int64(0); i <= x3; i++ {
		if i&1 == 1 {
			mobius[i] = 1
			min_prime[i] = int32(x3)
		}
	}
	phi := int64(0)
	for i := int64(1); i <= x3; i += 2 {
		if !(*prime_x32)[i] {

			prime = append(prime, int32(i))

			for j := i; j <= x3; j += i {
				mobius[j] *= -1
				if min_prime[j] == int32(x3) {
					min_prime[j] = int32(len(prime)) - 1
				}
			}
			i2 := i * i
			for j := i2; j <= x3; j += i2 {
				mobius[j] = 0
			}
		}
		phi += int64(mobius[i]) * ((x/i + 1) / 2)
	}

	bit := make([]int32, x32+1)
	is_bit_1 := make([]bool, x32+1)

	for i := int64(1); i <= x32; i++ {
		if (i & 1) == 1 {
			bit_add(&bit, i, 1, x32)
			is_bit_1[i] = true
		}
	}

	for b := int64(0); b+1 < int64(len(prime)); b++ {
		for m := int64(3); m <= x3; m += 2 {
			if b+1 < int64(min_prime[m]) && int64(prime[b+1])*m > x3 {
				phi -= int64(mobius[m]) * bit_sum(&bit, x/(m*int64(prime[b+1])))
			}
		}

		for i := int64(0); i <= x32; i += int64(prime[b+1]) {
			if is_bit_1[i] {
				bit_add(&bit, i, -1, x32)
				is_bit_1[i] = false
			}
		}
	}
	return phi
}

func main() {
	var x int64
	fmt.Scan(&x)
	start := time.Now()

	x2 := root2(x)
	x3 := root3(x)
	x32 := x3 * x3

	prime_x32 := eratosthenes(x32, x3)

	fmt.Println("the sieve was finished")

	pi_x32 := calc_pi_x32(prime_x32, x32)

	phi2 := int64((*pi_x32)[x3])*(int64((*pi_x32)[x3])-1)/2 - int64((*pi_x32)[x2])*(int64((*pi_x32)[x2])-1)/2
	for p := x3 + 1; p <= x2; p++ {
		if !(*prime_x32)[p] {
			phi2 += int64((*pi_x32)[x/p])
		}
	}

	phi := calc_phi(prime_x32, x, x3, x32, int64((*pi_x32)[x3]))

	pi := int64((*pi_x32)[x3]) - phi2 + phi - 1
	end := time.Now()

	fmt.Println(pi)
	fmt.Println(end.Sub(start))
}
