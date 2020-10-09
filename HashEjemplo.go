package main

import (
 "bytes"
 "crypto/sha256"
 "encoding"
 "fmt"
 "log"
 "hash"
)

//main
func main() {
 const (
	//Texto a incriptar
    example1 = "Profe estoy en zapatoca viaje para ver a mis viejos SALU2"
    //Utiliza este texto o cadena para convertir la clave final 
    example2 = "XXCodigo_de_incriptacionXX"
    )
    var firstHash hash.Hash
    firstHash = sha256.New()
    firstHash.Write([]byte(example1))
    var marshaler encoding.BinaryMarshaler
    var ok bool
    marshaler, ok = firstHash.(encoding.BinaryMarshaler)
    if !ok {
        log.Fatal("Si no se genra el hash")
    }
    var data []byte
    var err error
    data, err = marshaler.MarshalBinary()
    if err != nil {
        log.Fatal("Falla al crear el hash: ", err)
    }
    var secondHash hash.Hash
    secondHash = sha256.New()
    var unmarshaler encoding.BinaryUnmarshaler
    unmarshaler, ok = secondHash.(encoding.BinaryUnmarshaler)
    if !ok {
        log.Fatal("Volvio a falar el hash")
    }
    if err := unmarshaler.UnmarshalBinary(data); err != nil {
        log.Fatal("Falla al crear el hash: ", err)
    }
    firstHash.Write([]byte(example2))
    secondHash.Write([]byte(example2))
    fmt.Printf("%x\n", firstHash.Sum(nil))
    fmt.Println(bytes.Equal(firstHash.Sum(nil), secondHash.Sum(nil)))
}