package workerpool

type URl struct {
	Url string
	Err error
}

type Pool struct {
	urlStruct []*URl
}

func СreatePool(arr []string) *Pool {
	p := &Pool{}

	for i := 0; i < len(arr); i++ {
		newUrl := &URl{
			Url: arr[i],
		}
		p.urlStruct = append(p.urlStruct, newUrl)
	}

	return p
}
