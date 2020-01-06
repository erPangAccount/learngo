package workerRpc

import "crawler/engine/zhenai"

type WorkService struct{}

func (WorkService) Process(args Request, result *RequestResult) error {
	enginRequest, err := DeserializeRequest(args)
	if err != nil {
		return err
	}

	requestResult, err := zhenai.Worker(enginRequest)
	if err != nil {
		return err
	}

	*result = SerializeRequestResult(requestResult)

	return nil
}
