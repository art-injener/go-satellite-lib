package generator

import (
	"bytes"
	"encoding/binary"
	"math"
	"math/cmplx"
)

// LFMParams содержит параметры для функции генерации сигнала LFM.
type LFMParams struct {
	SampleRate float64 // Частота дискретизации
	Duration   float64 // Длительность в секундах
	StartFreq  float64
	EndFreq    float64
}

// LFMSignal генерирует сигнал LFM с заданными параметрами.
// Пример использования : LFMParams{ SampleRate: 48000.0, Duration: 0.1, StartFreq: 1000.0, EndFreq: 2000.
func LFMSignal(params LFMParams) []complex128 {
	numSamples := int(params.SampleRate * params.Duration)
	signal := make([]complex128, numSamples)

	timeStep := 1.0 / params.SampleRate
	chirpRate := (params.EndFreq - params.StartFreq) / params.Duration

	for i := 0; i < numSamples; i++ {
		time := float64(i) * timeStep
		phase := math.Pi * chirpRate * time * time
		signal[i] = cmplx.Rect(1,
			2*math.Pi*(params.StartFreq*time+phase))
	}

	return signal

}

// MutableSignalParams содержит параметры для функции генерации сигнала с переменной частотой и амплитудой.
type MutableSignalParams struct {
	SampleRate    float64               // Частота дискретизации
	Duration      float64               // Длительность в секундах
	FrequencyFunc func(float64) float64 // Функция, определяющая частоту в зависимости от времени
	AmplitudeFunc func(float64) float64 // Функция, определяющая амплитуду в зависимости от времени
}

// MutableFreqAmplSignal генерирует сигнал с переменной частотой и амплитудой с заданными параметрами.
// params := MutableSignalParams{
// SampleRate: 48000.0,
// Duration: 0.1,
// FrequencyFunc: func(time float64) float64 { return 1000.0 + 500.0*time} // Линейно изменяющаяся частота от 1000 до 1500 Гц },
// AmplitudeFunc: func(time float64) float64 { return 1.0 - 0.5*time } // Линейно уменьшающаяся амплитуда от 1 до 0.5
func MutableFreqAmplSignal(params MutableSignalParams) []complex128 {
	numSamples := int(params.SampleRate * params.Duration)
	signal := make([]complex128, numSamples)

	timeStep := 1.0 / params.SampleRate

	for i := 0; i < numSamples; i++ {
		time := float64(i) * timeStep
		frequency := params.FrequencyFunc(time)
		amplitude := params.AmplitudeFunc(time)
		phase := 2 * math.Pi * frequency * time
		signal[i] = cmplx.Rect(amplitude, phase)
	}

	return signal

}

func CmplxToBytes(signal []complex128) ([]byte, error) {
	byteBuffer := new(bytes.Buffer)
	for _, value := range signal {
		err := binary.Write(byteBuffer, binary.LittleEndian, real(value))
		if err != nil {
			return nil, err
		}

		err = binary.Write(byteBuffer, binary.LittleEndian, imag(value))
		if err != nil {
			return nil, err
		}
	}

	return byteBuffer.Bytes(), nil
}

// SimpleMutableSignal Генерация радиосигнала с перестройкой частоты и амплитуды
// Пример использования
// frequency := 1000.0   // Начальная частота сигнала в герцах
// amplitude := 0.5      // Начальная амплитуда сигнала в попугаях
func SimpleMutableSignal(frequency float64, amplitude float64) []byte {
	sampleRate := 2 * 100 // Частота дискретизации
	duration := 1.0       // Продолжительность сигнала в секундах
	numSamples := int(duration * float64(sampleRate))
	signal := make([]byte, numSamples) // 1 байт на каждый отсчет (byte)

	// Генерация сигнала
	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate) // Время в секундах
		// Перестройка частоты и амплитуды
		currentFrequency := frequency + 1000.0*math.Sin(2*math.Pi*5.0*t)
		currentAmplitude := amplitude + 0.2*math.Sin(2*math.Pi*2.0*t)
		// Генерация синусоидального сигнала
		value := currentAmplitude * math.Sin(2*math.Pi*currentFrequency*t)
		// Преобразование значения в byte и запись в байтовый срез
		sample := byte((value + 1) * 0.5 * 255)
		signal[i] = sample
	}

	return signal
}

// SimpleSignal - Генерация синусоидального сигнала
func SimpleSignal(frequency float64, amplitude float64) []byte {
	sampleRate := 2 * 1000 * 1000 // Частота дискретизации
	duration := 0.5               // Продолжительность сигнала в секундах
	numSamples := int(duration * float64(sampleRate))
	signal := make([]byte, numSamples) // 1 байт на каждый отсчет (byte)

	// Генерация сигнала
	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate) // Время в секундах
		value := amplitude * math.Sin(2*math.Pi*frequency*t)
		// Преобразование значения в byte и запись в байтовый срез
		sample := byte((value + 1) * 0.5 * 255)
		signal[i] = sample
	}

	return signal
}
