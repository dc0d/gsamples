package main

import (
	"encoding/json"
	"fmt"

	"github.com/dc0d/gsamples/node"
)

func main() {
	n := node.New()
	sampleTree(n, 0)

	js, _ := json.MarshalIndent(&n, "", "  ")
	fmt.Println(string(js))

	/*
		Output:

		{
		  "Name": "",
		  "Children": {
		    "0-1": {
		      "Name": "",
		      "Children": {
		        "1-1": {
		          "Name": "",
		          "Children": {}
		        },
		        "1-2": {
		          "Name": "",
		          "Children": {}
		        },
		        "1-3": {
		          "Name": "",
		          "Children": {}
		        }
		      }
		    },
		    "0-2": {
		      "Name": "",
		      "Children": {
		        "1-1": {
		          "Name": "",
		          "Children": {}
		        },
		        "1-2": {
		          "Name": "",
		          "Children": {}
		        },
		        "1-3": {
		          "Name": "",
		          "Children": {}
		        }
		      }
		    },
		    "0-3": {
		      "Name": "",
		      "Children": {
		        "1-1": {
		          "Name": "",
		          "Children": {}
		        },
		        "1-2": {
		          "Name": "",
		          "Children": {}
		        },
		        "1-3": {
		          "Name": "",
		          "Children": {}
		        }
		      }
		    }
		  }
		}
	*/
}

func sampleTree(n node.Node, d int) {
	if d >= 2 {
		return
	}
	for i := 0; i < 3; i++ {
		c := node.New()
		id := fmt.Sprintf("%v-%v", d, i+1)
		n.Children[id] = c
	}
	d++
	for _, vc := range n.Children {
		sampleTree(vc, d)
	}
}
