package broker

import (
	v1 "github.com/imneov/modelmesh/api/modelfulx/v1alpha"
	"io"
)

type Queue struct {
}

func (q Queue) Predict(stream v1.MFService_PredictServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		//@TODO
		_ = in //// User Logic
		if err := stream.Send(nil); err != nil {
			return err
		}

	}
}
