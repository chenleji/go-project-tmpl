package consul

import (
	"github.com/hashicorp/consul/api"
)

// eg: xxx/yyy/
func (client *ConsulClient) ListKV(path string, query *api.QueryOptions) (kvPairs api.KVPairs, queryMeta *api.QueryMeta, err error) {
	kvPairs, queryMeta, err = client.c.KV().List(path, query)
	return
}

func (client *ConsulClient) PutKV(kvPair *api.KVPair, options *api.WriteOptions) (writeMeta *api.WriteMeta, err error) {
	writeMeta, err = client.c.KV().Put(kvPair, options)
	return
}

func (client *ConsulClient) GetKV(key string, query *api.QueryOptions) (pair *api.KVPair, queryMeta *api.QueryMeta, err error) {
	pair, queryMeta, err = client.c.KV().Get(key, query)
	return
}

func (client *ConsulClient) DeleteKV(key string, options *api.WriteOptions) (meta *api.WriteMeta, err error) {
	meta, err = client.c.KV().Delete(key, options)
	return
}

func (client *ConsulClient) DeleteTree(path string, options *api.WriteOptions) (writeMeta *api.WriteMeta, err error) {
	writeMeta, err = client.c.KV().DeleteTree(path, options)
	return
}
