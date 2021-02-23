package collatzconjecture

func CollatzConjecture(n int) (int,int){
	i :=  0
	for ;n != 1;{
		if n%2 == 0 {
			n = n/2
		} else {
			n = 3*n + 1
		}
		i += 1
	}
	return i,1
}