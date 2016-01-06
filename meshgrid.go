package meshgrid

// Single generates an n-dimensional grid using a single set of values.
// dim specifies the number of dimensions, the entries in x specify the gridded values.
func Single(dim int, x []float64) [][]float64 {
	dims := make([]int, dim)
	for i := range dims {
		dims[i] = len(x)
	}
	sz := size(dims)
	pts := make([][]float64, sz)
	sub := make([]int, dim)
	for i := 0; i < sz; i++ {
		SubFor(sub, i, dims)
		pt := make([]float64, dim)
		for j := range pt {
			pt[j] = x[sub[j]]
		}
		pts[i] = pt
	}
	return pts
}

// Multiple generates an n-dimensional grid using a specified set of locations
// in each dimension.
func Multiple(x [][]float64) [][]float64 {
	dim := len(x)
	dims := make([]int, dim)
	for i := range x {
		dims[i] = len(x[i])
	}
	sz := size(dims)
	pts := make([][]float64, sz)
	sub := make([]int, dim)
	for i := 0; i < sz; i++ {
		SubFor(sub, i, dims)
		pt := make([]float64, dim)
		for j := range pt {
			pt[j] = x[j][sub[j]]
		}
		pts[i] = pt
	}
	return pts
}

func size(dims []int) int {
	n := 1
	for _, v := range dims {
		n *= v
	}
	return n
}

// IdxFor gives a linear index to a multi-dimensional subscript. Dims specifies
// the maximum size in each dimension. IdxFor is the converse of IdxFor.
func IdxFor(sub, dims []int) int {
	// This is a higher-dimensional version of matrix indexing.
	// mat index is i*ld + j
	// This is i*lda*ldb*ldc + j*ldb*ldc + k*ldc + l
	for _, v := range dims {
		if v <= 0 {
			panic("bad dims")
		}
	}
	var idx int
	stride := 1
	for i := len(dims) - 1; i >= 0; i-- {
		v := sub[i]
		d := dims[i]
		if v < 0 || v >= d {
			panic("bad subscript")
		}
		idx += v * stride
		stride *= d
	}
	return idx
}

// SubFor constructs the multi-dimensional subscript for the input linear index.
// Dims specifies the maximum size in each dimension. SubFor is the converse of
// IdxFor.
//
// If sub is non-nil the result is stored in-place into sub. If it is nil a new
// slice of the appropriate length is allocated.
func SubFor(sub []int, idx int, dims []int) []int {
	for _, v := range dims {
		if v <= 0 {
			panic("bad dims")
		}
	}
	if sub == nil {
		sub = make([]int, len(dims))
	}
	if len(sub) != len(dims) {
		panic("size mismatch")
	}
	if idx < 0 {
		panic("bad index")
	}
	stride := 1
	for i := len(dims) - 1; i >= 1; i-- {
		stride *= dims[i]
	}
	for i := 0; i < len(dims)-1; i++ {
		v := idx / stride
		if v >= dims[i] {
			panic("bad index")
		}
		sub[i] = v
		idx -= v * stride
		stride /= dims[i+1]
	}
	if idx > dims[len(sub)-1] {
		panic("bad dims")
	}
	sub[len(sub)-1] = idx
	return sub
}
