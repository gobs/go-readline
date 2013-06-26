// test program for the readline package
package readline

import "fmt"

func ExampleReadLine() {
	prompt := "by your command> ";

	//loop until ReadLine returns nil (signalling EOF)
	L: for {
		switch result := ReadLine(&prompt); true {
		case result == nil: break L //exit loop

		case *result != "": //ignore blank lines
			fmt.Println(*result);
			AddHistory(*result); //allow user to recall this line
		}
	}
}
