# Aurelia

## About Aurelia

Aurelia is a function to hash password with HMAC (Hash-based message authentication code), PBKDF2 (Password-Based Key Derivation Function 2) and Random Bytes. The main reason to create this library is to make good hash for security reason and the hash must be:

- Has a random salt every hash generation
- Has a user unique hash so it will generate differently if the key and credential is different using HMAC
- Has a unique signature for authentication using PBKDF2
- Has a different hash value every generation but authenticable
- Simple to authentication

This library is inspired by a post in **crackstation** about how to make secure hash [here](https://crackstation.net/hashing-security.htm).

## How to use

Here is a simple step to generate a hash:

```go
import "github.com/codefluence-x/aurelia"

func main() {
  credential := "credential"
  key := "key"

  fmt.Printf("Hashed: %s\n", aurelia.Hash(credential, key))
}
```

The hash result will be random every generation and the one of the result is:

```text
AURELIA_4dd9b11ae216e0c231a330447adb8181f34ed8cd.1cace0ea1bca8bf2252b2b2a84148b630567772a.UC_46896b6a18f393cc6cbdf2664e0ba85ce7ba75db5d0e712eaedf690c337bfee2
```

Authentication is simply like this:

```go
import "github.com/codefluence-x/aurelia"

func main() {
  credential := "credential"
  key := "key"
  hashed := "AURELIA_BABLABLABLABLA"

  if aurelia.Authenticate(credential, key, hased) {
    fmt.Println("Authentication success!")
  } else {
    fmt.Println("Authentication failed :(")
  }
}
```

That's it, have fun to use it! :)

## Other Version

- [PHP](https://github.com/insomnius/Aurphm)
