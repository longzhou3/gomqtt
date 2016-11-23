package service

type Cache struct {
	As   *Accounts    // 用户列表
	Sas  *StreamAddrs // stream地址缓存列表
	Cids *conIDs      // 用户链接缓存，映射account
	bts  *btCache     // 广播主题缓存,保存用户的acc
}

func NewCache() *Cache {
	cache := &Cache{
		As:   NewAccounts(),
		Sas:  NewStreamAddrs(),
		Cids: newconIDs(),
		bts:  newbtCache(),
	}
	return cache
}

func (cache *Cache) Init() {

}

func (cache *Cache) Start() {

}

func (cache *Cache) Close() error {
	return nil
}
