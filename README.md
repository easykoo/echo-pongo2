# echo-pongo2

## Usage

~~~ go

import p "github.com/easykoo/echo-pongo2"

	t := pongo2.PrepareTemplates(pongo2.Options{
		Directory:  "public/views/",
		Extensions: []string{".html"},
	})
	e.SetRenderer(t)

	mData := make(map[string]interface{})
	return c.Render(http.StatusOK, "share", mData)

~~~

##License

This code is under an Apache v2 License.


## Author

* [Steven](https://github.com/easykoo)