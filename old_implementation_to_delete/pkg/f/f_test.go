package f

import "testing"

func TestSanitize(t *testing.T) {

	in := "/home/Users//user////my/folder"
	expected := "/home/Users/user/my/folder"
	out,err := Sanitize(in)
	if err != nil && out != expected {
		t.Error("Double slash sanityze wrong, expecting: "+expected+" got: ", out  )
	}

	in = "home/Users//user////my/folder"
	expected = "home/Users/user/my/folder"
	out,err = Sanitize(in)
	if err != nil && out != expected {
		t.Error("Double slash sanityze wrong, expecting: "+expected+" got: ", out  )
	}

	in = "/home/Users/user/../user2/./folder/../folder/../"
	expected = "/home/Users/user2"
	out,err = Sanitize(in)
	if err != nil && out != expected {
		t.Error("Path changing not correct, expecting: "+expected+" got: ", out   )
	}

	in = "/home/Users/../../../.././../"
	out,err = Sanitize(in)
	if err == nil {
		t.Error("trying to go up the path, expecting: "+expected+" got: ", out   )
	}

	in = "/home///Users/////user/.././user2/.////folder///../folder/..////"
	expected = "/home/Users/user2"
	out,err = Sanitize(in)
	if err != nil && out != expected {
		t.Error("Complex test: "+expected+" got: ", out   )
	}

	in = "/home///Users/////user/.././user2/.////folder///../folder/..////"
	expected = "/home/Users/user2"
	out,err = Sanitize(in)
	if err != nil && out != expected {
		t.Error("Complex test with relative path: "+expected+" got: ", out   )
	}
}

func TestFileTypeFromExtension(t *testing.T)  {

	out := FileTypeFromExtension("jpg")
	expected :=FileTypeExtInfo{
		MimeType: "image/jpeg",
		Type : "image",
	}
	if out != expected {
		t.Error("FileTypeExtInfo does not match to expected: " )
		t.Error(expected)
		t.Error(" but got: ")
		t.Error(out)
	}
}