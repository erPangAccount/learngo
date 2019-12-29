package model

type SearchResult struct {
	Keyword  string
	PageInfo struct {
		Total      int
		PrevPage   int
		TargetPage int
		NextPage   int
		PageSize   int
	}
	Items []interface{}
}

type Item struct {
	Url       string
	Avatar    string
	NickName  string
	Sex       string
	Age       int
	Weight    int
	Height    int
	OtherInfo string
	ObjInfo   string
}
