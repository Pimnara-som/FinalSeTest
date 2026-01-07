package entity

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCodeBook(t *testing.T) {
	g := NewGomegaWithT(t)
	p := Book{
		Title: "kkk",
		Price: 49.0,
		Code:  "d123456",
	}
	ok, err := ValidateBook(&p)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err).Error().To(Equal("Code must start with BK followed by 6 digits (0-9)"))

}
