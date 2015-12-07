package gostressed


import (
    "fmt"
    "time"
    "os"
    "io"
    "io/ioutil"
    "crypto/sha256"
	"encoding/hex"
	
	"net/http"
	"math/rand"
	"strings"
	"bytes"
)

type Notifier interface {
	Notify(a string)
}

func someRoutine(a Notifier) {
	fullText := "Hello people! Let's do some cool programming stuff"
	s := strings.Split(fullText, " ")
	for _, word := range s {
		time.Sleep(time.Duration(2)*time.Second)
		a.Notify(word)
	}
}

func RunGoRoutine(a Notifier) {
	go someRoutine(a)
}

func HTTPGetCall() string {
	fmt.Println("Making GET HTTP calls")
	resp, err := http.Get("http://www.theverge.com/")
	if err != nil {
		fmt.Println("Error loading site:")
		return ""	// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) 
	if err != nil {
		fmt.Println("Cannot read response body")
		return ""	// handle error
	}

	return string(body)
}

func HTTPPostCall() string {
	fmt.Println("Making Post HTTP calls")
    var reqBody = []byte(`name=Jane+Doe&address=123+Main+St`)
    req, err := http.NewRequest("POST", "https://www.apple.com", bytes.NewBuffer(reqBody))
    //req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
    	fmt.Println("Error processing POST request")
		return ""
    }
    defer resp.Body.Close()

    body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
    	fmt.Println("Error reading response from POST request")
		return ""
    }

    return string(body)
}

func SortFixedList() bool {
	fmt.Println("Sorting List")
	numeros := []int{32,56,2,94,46,67,32,71,29,78,78,78,12,33,64,90,97,1,2,27,68,88,83,5,71,46,81,7,37,59,25,93,21,26,31,20,90,7,92,36,100,66,93,12,76,24,17,46,15,9,63,37,18,32,43,80,44,70,77,45,82,66,32,11,85,10,62,17,100,43,34,7,73,38,90,45,23,3,68,45,67,48,47,35,14,72,87,74,10,82,34,59,92,15,2,87,73,80,4,43}
	
	qsort(numeros)
	return true
}

func GenerateAndSort() bool {
	resp := make([]int, 1000000)
	for i := range resp {
		resp[i] = rand.Intn(1000)
	}

	qsort(resp)
	return true
}

func qsort(a []int) []int {
  if len(a) < 2 { return a }

  left, right := 0, len(a) - 1

  // Pick a pivot
  pivotIndex := len(a)/2

  // Move the pivot to the right
  a[pivotIndex], a[right] = a[right], a[pivotIndex]

  // Pile elements smaller than the pivot on the left
  for i := range a {
    if a[i] < a[right] {
      a[i], a[left] = a[left], a[i]
      left++
    }
  }

  // Place the pivot after the last smaller element
  a[left], a[right] = a[right], a[left]

  // Go down the rabbit hole
  qsort(a[:left])
  qsort(a[left + 1:])

  return a
}

func WriteToFile(pathToFile string) bool {
	fmt.Println("Writing to file:",pathToFile)

	f, err := os.Create(pathToFile)
	if err != nil {
		fmt.Println("Error creating file at:",pathToFile)
		return false
	}
    defer f.Close()

    var buffer bytes.Buffer
    for i := 0; i < 1000000; i++ {
        buffer.WriteString(fmt.Sprintf("%c",i%26 + 65))
        if (i%26 == 0) {
			buffer.WriteString("\n")
        }
    }

    _ , err = f.WriteString(buffer.String())
    if err != nil {
		fmt.Println("Error writing string to file")
		return false
	}
    f.Sync()

	return true
}

func ReadFromFile(pathToFile string) string {
	//path = strings.Replace(path, "randomFile", "BlabCake", 1)
	// fmt.Println("Reading from file at:",path)
	// f, err := os.Open(path)
	// if err != nil {
	// 	fmt.Println("Error reading file at:",path)
	// 	return ""
	// }
	// defer f.Close()

	dat, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		fmt.Println("Error reading file at:",pathToFile)
		return ""
	}
	return string(dat)

 // 	r := bufio.NewReader(f)
 // 	for {
	// 	line, _, err := r.ReadLine()
	// 	if err != nil {
	// 		fmt.Println("Done reading file")
	// 		break
	// 	}
	// 	fmt.Println("Read: ",string(line))
	// }
}

func HashFile(pathToFile string) string {
	hasher := sha256.New()
	f, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println("Error reading file at:",pathToFile)
		return ""
	}
	defer f.Close()
	if _, err := io.Copy(hasher, f); err != nil {
	    fmt.Println("Error hashing file:",err)
	}

	resp := hex.EncodeToString(hasher.Sum(nil))
	return resp
}
