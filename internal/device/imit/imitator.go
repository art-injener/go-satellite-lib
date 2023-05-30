package imit

import dsp "github.com/art-injener/go-satellite-lib/internal/dsp/generator"

type DeviceImpl struct {
}

func NewDeviceImpl() *DeviceImpl {
	return &DeviceImpl{}
}

func (dev *DeviceImpl) Find() {

}

func (dev *DeviceImpl) Open() error {
	return nil
}

func (dev *DeviceImpl) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (dev *DeviceImpl) Close() error {
	return nil
}

func (dev *DeviceImpl) Generate() ([]byte, error) {
	return dev.Signal()
}

func (dev *DeviceImpl) Signal() ([]byte, error) {
	params := dsp.MutableSignalParams{
		SampleRate:    48000.0,
		Duration:      0.1,
		FrequencyFunc: func(time float64) float64 { return 1000.0 + 500.0*time },
		AmplitudeFunc: func(time float64) float64 { return 1.0 - 0.5*time },
	}

	return dsp.CmplxToBytes(dsp.MutableFreqAmplSignal(params))
}
