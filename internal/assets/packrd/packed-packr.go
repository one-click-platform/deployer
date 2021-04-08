// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "d8c00c24c4e277b2f88d5107ea934201"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"dd1ee035574cfa7beaeca289d9042b01": "1f8b08000000000000ff7c504b72c2300cddeb146f994ce1046c7b85ae19c51141ad63bbb23d34b7ef04333464d19dad27bddff188b75927e322f84844ce647d161ebc809d8b35948c8e0040470c3a6531658f643ab32df892e570476566f528f253106241a8dea306fdaed2f0c439dfa28de72be72b86a5083f17a93feda493c54f714fe9c7f7fcbf85c0b3bc3a68f3478e97eb865ca2894e61a540f7b7d6c3e42226c149deb4a063dfce5aae6ed53b6cd8fb7b0ea26da9eff11688468b6957ea897e030000ffff8e7a23367c010000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("migrations", "./migrations")
		b.SetResolver("01_initial.sql", packr.Pointer{ForwardBox: gk, ForwardPath: "dd1ee035574cfa7beaeca289d9042b01"})
	}()
	return nil
}()
