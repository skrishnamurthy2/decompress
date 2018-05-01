package decompress

import (
  "strconv"
)

func Decompress(compressed []byte, finishEarly bool) ([]byte, int) {

  var (
	dataCompressed = make([]byte, 0)
	timesCompressed = make([]byte, 0)
	insideCompression = false
	index = 0
	uncompressedResult = make([]byte, 0)
  )

  for index < len(compressed) {

    if compressed[index] == '[' {

	  insideCompression = true
	  index++

	} else if(compressed[index] == ']') {

	  insideCompression = false
	  index++

	  repeat, _ := strconv.Atoi(string(timesCompressed))

	  uncompressedResult = slideAppend(uncompressedResult, sliceRepeatCopy(dataCompressed, repeat))

	  if finishEarly {
	    return uncompressedResult, index
	  }

	  timesCompressed = make([]byte, 0)
	  dataCompressed = make([]byte, 0)

	} else {

	  if compressed[index] >= '0' && compressed[index] <= '9' {

		if insideCompression {

		  result, indexProcessed := decompress(compressed[index:], true)
		  index += indexProcessed - 1

		  dataCompressed = slideAppend(dataCompressed, result)

		} else {

		  timesCompressed = append(timesCompressed, compressed[index])
		}

	  } else {

	    if(insideCompression) {

		  dataCompressed = append(dataCompressed, compressed[index])
		} else {

		  uncompressedResult = append(uncompressedResult, compressed[index])

		}
	  }

	  index++;
	}

  }

  return uncompressedResult, index

}

func sliceRepeatCopy(src [] byte, repeat int) []byte {

  destination := make([]byte, 0)

  for i := 0; i < repeat; i++ {

	for _, sliceValue := range src {
	  destination = append(destination, sliceValue)
	}

  }
  return destination
}

func slideAppend(dest [] byte, src [] byte) []byte {

  for _, uncompByte := range src {
	dest = append(dest, uncompByte)
  }
  return dest
}
