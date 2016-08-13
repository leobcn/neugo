package neugo

import (
	"errors"

	"github.com/jinseokYeom/neugo/matrix"
)

var (
	// wrong number of inputs
	ErrInputLen = errors.New("invalid input length")
)

type NeuralNet struct {
	conf    *Config          // configuration
	weights []*matrix.Matrix // list of weights
}

func NewNeuralNet(conf *Config) *NeuralNet {
	return &NeuralNet{
		conf: conf,
		weights: func() []*matrix.Matrix {
			weights := make([]*matrix.Matrix, 2)
			// input layer -> hidden layer
			weights[0] = matrix.NewNorm(conf.NumInput,
				conf.NumHidden, 0.0, 6.0)
			// hidden layer[t] -> hidden layer[t+1]
			for i := 1; i < conf.NumLayer-1; i++ {
				weights[i] = matrix.NewNorm(conf.NumHidden,
					conf.NumHidden, 0.0, 6.0)
			}
			// hidden layer -> output layer
			weights[conf.NumLayer-1] = matrix.NewNorm(conf.NumHidden,
				conf.NumOutput, 0.0, 6.0)
			return weights
		}(),
	}
}

// Get the neural network's weights.
func (n *NeuralNet) Weights() []*matrix.Matrix {
	return weights
}

// Activate the neural network.
func (n *NeuralNet) Activate(inputs []float64) []float64 {
	if len(inputs) != n.conf.NumInput {
		return nil, ErrInputLen
	}
}
