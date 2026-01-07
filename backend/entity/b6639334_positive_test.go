package entity

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestPositiveBook(t *testing.T) {
	g := NewGomegaWithT(t)
	p := Book{
		Title: "kkk",
		Price: 60.5,
		Code:  "BK123456",
	}
	ok, err := ValidateBook(&p)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
	//case1

}
