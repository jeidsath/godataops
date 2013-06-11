package dataops

type Interface interface {
	Len() int
	Swap(i, j int)
}

func Reverse(data Interface) {
	dataLength := data.Len()
	for ii, jj := 0, dataLength-1; ii < dataLength/2; ii, jj = ii+1, jj-1 {
		data.Swap(ii, jj)
	}
}

func Rotate(data Interface, rotate int) {
        dataLength := data.Len()
        modRotate = rotate % dataLength
        for ii, jj := dataLength - modRotate, dataLength-1; ii >= 0; ii, jj = ii-1, jj-1 {
                data.Swap(ii, jj)
        }
}
