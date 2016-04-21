package main

import (
  "fmt"
  "crypto/md5"
  "crypto/sha256"
  "crypto/hmac"
  "crypto/aes"
  "crypto/rsa"
  "crypto/cipher"
  "crypto/rand"
  "crypto/x509"
  "encoding/hex"
  "encoding/base64"
  "encoding/pem"
  "io"
  "io/ioutil"
)

func demoHash_md5() {
  fmt.Println("--------- Demo encoding hash using md5 -------")
  message := "Hello world, go!"
  fmt.Println("plaintext:")
  fmt.Println(message)

  h := md5.New()
  h.Write([]byte(message))
  hash_message := hex.EncodeToString(h.Sum(nil))
  fmt.Println("hashing message:")
  fmt.Println(hash_message)
  fmt.Println("--------- -------\n")
}

func demoHash_sha256() {
  fmt.Println("--------- Demo encoding hash using sha256 -------")
  message := "Hello world, go!"
  fmt.Println("plaintext:")
  fmt.Println(message)

  h := sha256.New()
  h.Write([]byte(message))
  hash_message := hex.EncodeToString(h.Sum(nil))
  fmt.Println("hashing message:")
  fmt.Println(hash_message)
  fmt.Println("--------- -------\n")
}

func demoHash_key(key, message string) {
  fmt.Println("--------- Demo encoding hash using key : HMAC and sha256-------")
  fmt.Println("key:")
  fmt.Println(key)
  fmt.Println("plaintext:")
  fmt.Println(message)

  hmacKey := []byte(key)
  h := hmac.New(sha256.New, hmacKey)
  h.Write([]byte(message))
  hash_message := hex.EncodeToString(h.Sum(nil))
  fmt.Println("hashing message:")
  fmt.Println(hash_message)
  fmt.Println("--------- -------\n")
}

func encrypt_symmetric_crypto(key, message string) string {
  fmt.Println("--------- Demo encrypt encrypt_symmetric_crpyto ------")
  if len(key)!=16 && len(key)!=24 && len(key)!=32 {
    fmt.Println("key must 16,24,32 byte length")
    return ""
  }

  bc, err := aes.NewCipher([]byte(key))

  if err != nil {
    panic(err)
  }

  text := []byte(message)
  ciphertext := make([]byte, aes.BlockSize+len(text))
  iv := ciphertext[:aes.BlockSize]

  if _, err := io.ReadFull(rand.Reader, iv); err != nil {
    panic(err)
  }

  cfb := cipher.NewCFBEncrypter(bc, iv)
  cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)

  return base64.StdEncoding.EncodeToString(ciphertext)
}

func decrypt_symmetric_crypto(key, message string) string {
  fmt.Println("--------- Demo decrypt decrypt_symmetric_crpyto ------")
  if len(key)!=16 && len(key)!=24 && len(key)!=32 {
    fmt.Println("key must 16,24,32 byte length")
    return ""
  }
  encrypted, _ := base64.StdEncoding.DecodeString(message)

  bc, err := aes.NewCipher([]byte(key))
  if err != nil {
    panic(err)
  }
  if len(encrypted) < aes.BlockSize{
    panic("ciphertext too short")
  }
  iv := encrypted[:aes.BlockSize]
  encrypted = encrypted[aes.BlockSize:]
  cfb := cipher.NewCFBDecrypter(bc, iv)
  cfb.XORKeyStream(encrypted, encrypted)

  return string(encrypted)
}

func generateRSAkeys(){
  fmt.Println("Generating RSA keys....")

  // change files and their paths
  privKeyFile := "./private.rsa.key"
  pubKeyFile := "./public.rsa.key"
  // generate RSA keys
  privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
  if err != nil {
    panic(err)
  }
  //extract private and public keys from RSA keys
  privASN1 := x509.MarshalPKCS1PrivateKey(privateKey)
  pubASN1, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
  if err != nil {
    panic(err)
  }
  // store private and public keys into files
  privBytes := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privASN1,})
  pubBytes := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: pubASN1,})

  ioutil.WriteFile(privKeyFile, privBytes, 0644)
  ioutil.WriteFile(pubKeyFile, pubBytes, 0644)

  fmt.Println("Done")
}

func encrypt_asymmetric_crypto(message string) string {
  fmt.Println("--- Demo encrypt encrypt_asymmetric_crypto....")
  // public key file
  pubKeyFile := "./public.rsa.key"
  //read public key from file
  pubBytes, err := ioutil.ReadFile(pubKeyFile)
  if err != nil {
    panic(err)
  }
  pubBlock, _ := pem.Decode(pubBytes)
  if pubBlock == nil {
    fmt.Println("Failed to load public keu file")
    return ""
  }
  // decode the RSA public key
  publicKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
  if err != nil{
    fmt.Printf("bad puclic key: %s", err)
    return ""
  }
  // encrypt message
  msg := []byte(message)
  encryptedmsg, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), msg)
  if err != nil {
    panic(err)
  }
  return base64.StdEncoding.EncodeToString(encryptedmsg)
}

func decrypt_asymmetric_crypto(message string) string {
  fmt.Println("--- Demo decrypt decrypt_asymmetric_crypto....")
  //private key file
  privKeyFile := "./private.rsa.key"
  //read private keu from file
  privBytes, err := ioutil.ReadFile(privKeyFile)
  if err != nil {
    panic(err)
  }
  privBlock, _ := pem.Decode(privBytes)
  if privBlock == nil {
    fmt.Println("Failed to load private key file")
    return ""
  }
  // devoce the RSA private key
  privateKey, err := x509.ParsePKCS1PrivateKey(privBlock.Bytes)
  if err != nil {
    fmt.Printf("bad public key: %s", err)
    return ""
  }
  // encrypt message
  encrypted, _ := base64.StdEncoding.DecodeString(message)
  decrypeddmsg, err := rsa.DecryptPKCS1v15( rand.Reader, privateKey, encrypted)

  if err != nil {
    panic(err)
  }
  return string(decrypeddmsg)
}


func main() {
  // asymetric crypto
  generateRSAkeys()
  plainText := "Hello world, go!"
  fmt.Println("plainText:")
  fmt.Println(plainText)

  rsa_encrypted := encrypt_asymmetric_crypto(plainText)
  fmt.Println("encrypted:")
  fmt.Println(rsa_encrypted)

  rsa_decrypted := decrypt_asymmetric_crypto(rsa_encrypted)
  fmt.Println("decrypted:")
  fmt.Println(rsa_decrypted)

  //symetric crypto
  key := "this is key 1234"
  message := "Hello world, go!"
  encrypted := encrypt_symmetric_crypto(key, message)
  fmt.Println("message:")
  fmt.Println(message)
  fmt.Println("key:")
  fmt.Println(key)
  fmt.Println("encrypted:")
  fmt.Println(encrypted)

  decrypted := decrypt_symmetric_crypto(key, encrypted)
  fmt.Println("key:")
  fmt.Println(key)
  fmt.Println("decrypted:")
  fmt.Println(decrypted)


  demoHash_md5()
  demoHash_sha256()
  demoHash_key("mykey", "Hello World, go!")
}
