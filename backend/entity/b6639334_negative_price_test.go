package entity

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestPriceBook(t *testing.T) {
	g := NewGomegaWithT(t)
	p := Book{
		Title: "kkk",
		Price: 49.0,
		Code:  "BK123456",
	}
	ok, err := ValidateBook(&p)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err).Error().To(Equal("Price must be between 50 and 5000"))

}
func TestPriceBook2(t *testing.T) {
	g := NewGomegaWithT(t)
	p := Book{
		Title: "kkk",
		Price: 5000.01,
		Code:  "BK123456",
	}
	ok, err := ValidateBook(&p)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err).Error().To(Equal("Price must be between 50 and 5000"))

}
