package main
import "github.com/aws/aws-lambda-go/lambda"

func handleRequest() (string, error) {
  return "Vishnu Pedireddi Hello from Go!", nil
}

func main() {
  lambda.Start(handleRequest)
}
