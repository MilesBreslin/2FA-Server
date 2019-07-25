// help from
// https://blog.gojekengineering.com/a-diy-two-factor-authenticator-in-golang-32e5641f6ec5
package hmac

import (
  "bytes"
  "crypto/hmac"
  "crypto/sha1"
  "encoding/base32"
  "encoding/binary"
  "fmt"
  "strings"
  "time"
)

func getHOTPToken(secret string, interval int64) string {

  //Converts secret to base32 Encoding. Base32 encoding desires a 32-character
  //subset of the twenty-six letters A–Z and ten digits 0–9
  key, err := base32.StdEncoding.DecodeString(strings.ToUpper(secret))

  if err != nil {
    fmt.Println("error:", err)
    return "first fail"
  }

  bs := make([]byte, 8)
  binary.BigEndian.PutUint64(bs, uint64(interval))

  //Signing the value using HMAC-SHA1 Algorithm
  hash := hmac.New(sha1.New, key)
  hash.Write(bs)
  h := hash.Sum(nil)

  // We're going to use a subset of the generated hash.
  // Using the last nibble (half-byte) to choose the index to start from.
  // This number is always appropriate as it's maximum decimal 15, the hash will
  // have the maximum index 19 (20 bytes of SHA1) and we need 4 bytes.
  o := (h[19] & 15)

  var header uint32
  //Get 32 bit chunk from hash starting at the o
  r := bytes.NewReader(h[o : o+4])

  err = binary.Read(r, binary.BigEndian, &header)

  if err != nil {
    fmt.Println("error:", err)
    return "fail"
  }

  //Ignore most significant bits as per RFC 4226.
  //Takes division from one million to generate a remainder less than < 7 digits
  h12 := (int(header) & 0x7fffffff) % 1000000

  //Converts number as a string

  return fmt.Sprintf("%06d",h12)
}

func getTOTPToken(secret string) string {
  //The TOTP token is just a HOTP token seeded with every 30 seconds.
  interval := time.Now().Unix() / 30
  return getHOTPToken(secret, interval)
}

