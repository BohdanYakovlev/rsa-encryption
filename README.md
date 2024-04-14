# RSA Encryption

This Go program implements RSA encryption and decryption using randomly generated keys. RSA (Rivest-Shamir-Adleman) is an asymmetric cryptographic algorithm widely used for secure data transmission. It involves generating a public-private key pair, where the public key is used for encryption and the private key is used for decryption.

## How to Run the Code

To run this console app, follow these steps:

1. Clone this repository to your local machine:

```
git clone <repository_url>
```

2. Navigate to the directory containing the Go code:

```
cd rsa-encryption
```

3. Ensure you have Go installed on your system. If not, you can download it from [here](https://golang.org/dl/).

4. Execute the Go code by running the following command:

```
go run main.go
```
5.The program will generate a public-private key pair, encrypt a sample integer value, and then decrypt it.

## Algorithm
### Key Generation

1. **Generate Public Key**: 
   - Generate two large prime numbers, p and q.
   - Calculate the modulus, n = p * q.
   - Choose an integer e such that 1 < e < φ(n) and gcd(e, φ(n)) = 1, where φ(n) is Euler's totient function.
   - The public key consists of the pair (e, n).

2. **Generate Private Key**: 
   - Compute d as the modular multiplicative inverse of e modulo φ(n), i.e., d ≡ e^(-1) (mod φ(n)).
   - The private key consists of the pair (d, n).

### Encryption

To encrypt a message M using the public key (e, n):
   - Compute the ciphertext C = M^e (mod n).

### Decryption

To decrypt the ciphertext C using the private key (d, n):
   - Compute the plaintext M = C^d (mod n).

## Conclusion

This program demonstrates the RSA encryption and decryption process, which is essential for secure data transmission in various applications. By generating a public-private key pair and utilizing modular arithmetic, RSA ensures confidentiality and integrity of the transmitted data. Despite its security, RSA's efficiency depends on the size of the keys, making it suitable for small data sizes in practical implementations.

Feel free to reach out if you have any questions or encounter any issues while running the code.
