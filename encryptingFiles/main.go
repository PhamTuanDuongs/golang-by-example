package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

// EncryptFile encrypts a file using AES encryption
func EncryptFile(file []byte, inputFilePath string, outputFilePath string) error {
	// Open the input file
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}

	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}

	defer outputFile.Close()

	// Create AES block cipher
	block, err := aes.NewCipher(file)
	if err != nil {
		return err
	}

	// Generate a random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// Write the IV to the beginning of the output file
	if _, err := outputFile.Write(iv); err != nil {
		return err
	}

	// Create cipher block mode
	stream := cipher.NewCFBEncrypter(block, iv)
	//   Encrypt the file content

	writer := &cipher.StreamWriter{
		S: stream,
		W: outputFile}

	if _, err := io.Copy(writer, inputFile); err != nil {
		return err
	}

	return nil
}

// DecryptFile decrypts a file using AES decryption
func DecryptFile(key []byte, inputFilePath string, outputFilePath string) error {
	// Open the input file
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	// allocate space for ciphered data
	out := make([]byte, aes.BlockSize)
	// Read the IV from the beginning of the input file
	if _, err := io.ReadFull(inputFile, out); err != nil {
		return err
	}

	// Create AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create cipher block mode
	stream := cipher.NewCFBDecrypter(block, out)

	// Decrypt the file content
	reader := &cipher.StreamReader{S: stream, R: inputFile}
	if _, err := io.Copy(outputFile, reader); err != nil {
		return err
	}

	return nil
}

func main() {
	// Key for AES encryption (must be 16, 24, or 32 bytes long)
	key := []byte("vcg academy 1234")
	fmt.Println(len(key))

	// Input and output file paths
	inputFile := "Action.pdf"
	encryptedFile := "encrypted.pdf"
	decryptedFile := "decrypted.pdf"

	// Encrypt the file
	if err := EncryptFile(key, inputFile, encryptedFile); err != nil {
		log.Fatalf("Error encrypting file: %v", err)
	}
	fmt.Println("File encrypted successfully!")

	// Decrypt the file
	if err := DecryptFile(key, encryptedFile, decryptedFile); err != nil {
		log.Fatalf("Error decrypting file: %v", err)
	}
	fmt.Println("File decrypted successfully!")

	// // Delete the decrypted file
	// if err := os.Remove(decryptedFile); err != nil {
	// 	fmt.Println(err)
	// }
}
