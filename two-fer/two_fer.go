//Package twofer concatenates a string into a sentence.
package twofer

//ShareWith takes a string and concatenates it.
func ShareWith(name string) string {
	//  Go variables passed by value
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
