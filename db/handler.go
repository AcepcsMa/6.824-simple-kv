package db

import "../network"

type KVHandler struct {
	Instance *KVDB
}

func (handler *KVHandler) Get(arg *network.GetArg, response *network.GetResponse) error {
	value, err := handler.Instance.Get(arg.Key)
	if err != nil {
		response.StatusCode = network.FAIL
	}
	response.Value = value
	response.StatusCode = network.SUCCESS
	return nil
}

func (handler *KVHandler) Put(arg *network.PutArg, response *network.PutResponse) error {
	err := handler.Instance.Put(arg.Key, arg.Value)
	if err != nil {
		response.StatusCode = network.FAIL
	}
	response.StatusCode = network.SUCCESS
	return nil
}