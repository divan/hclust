package hclust

import "testing"

func TestMean(t *testing.T) {
	a := []float64{2.0, 10.0, 18.0}
	expected := 10.0
	got := mean(a)
	if got != expected {
		t.Errorf("expected to get %v, but got %v", expected, got)
	}
}

func TestResiduals(t *testing.T) {
	a := []float64{2.0, 10.0, 18.0}
	expected := []float64{-8.0, 0.0, 8.0}
	got := residuals(a)
	for i := 0; i < len(a); i++ {
		if got[i] != expected[i] {
			t.Errorf("expected to get %v, but got %v", expected[i], got[i])
		}
	}
}

func TestVariance(t *testing.T) {
	a := []float64{28.0, 29.0, 30.0, 31.0, 32.0}
	expected := 2.0
	got := variance(a)
	if got != expected {
		t.Errorf("expected to get %v, but got %v", expected, got)
	}

	a = []float64{800.0, 720.0, 655.0, 655.0, 625.0, 600.0, 590.0, 529.0, 513.0, 502.0, 502.0, 502.0}
	expected = 8549.40972222222
	got = variance(a)
	if got != expected {
		t.Errorf("expected to get %v, but got %v", expected, got)
	}
}
