package largestTimeForGivenDigits

import "testing"

func largestTimeFromDigits(A []int) string {
	qSort(A, 0, len(A)-1)
	ret := [5]byte{0,0,':',0,0}
	for i:= len(A)-1; i>=0; i-- {
		if ret[0] == 0 && A[i]<=2 {
			ret[0] = byte('0'+A[i])
		} else if ret[1] == 0 &&A[i]<=3 {
			ret[1] = byte('0'+A[i])
		}else if ret[3] == 0 &&A[i] <= 5 {
			ret[3] = byte('0'+A[i])
		}else if A[i]<= 9 {
			ret[4] = byte('0'+A[i])
		}
	}
	for _, v := range ret {
		if v == 0 {
			return ""
		}
	}
	return string(ret[:])
}

func qSort(A []int, left, right int) {
	if left >= right {
		return
	}
	p := partition(A, left, right)
	qSort(A, left, p-1)
	qSort(A, p+1, right)
}

func partition(A []int, left, right int) int{
	i, val:= left, A[right]
	for j:= i; j<= right; j++ {
		if A[j] < val {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	A[i], A[right] = A[right], A[i]
	return i
}

func TestLargestTimeFromDigits(t *testing.T) {
	A := []int{1,2,3,4}
	t.Log(largestTimeFromDigits(A))
}