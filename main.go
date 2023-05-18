package main

import (
    "fmt"
    "html/template"
    "log"
    "net/smtp"
    "os"
    "sentiment"
)

func main() {
    // create  template
    tmpl := template.Must(template.ParseFiles("responder.html"))

    // create a new SMTP client
    client, err := smtp.Dial("smtp.gmail.com:587")
    if err != nil {
        log.Fatal(err)
    }

    // Login
    err = client.Login("MyEmail@gmail.com", "MyPassword")
    if err != nil {
        log.Fatal(err)
    }

    // read from stdin
    msg, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
package main


func main() {
    // Get the email content from the command line.
    emailContent, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    // Create a sentiment analysis object.
    sentimentAnalyzer := sentiment.New()

    // Analyze the email content.
    sentiment, score := sentimentAnalyzer.Analyze(emailContent)

    // Print the sentiment and score.
    fmt.Printf("Sentiment: %s\nScore: %f\n", sentiment, score)
}
    }

    
    // create a new email
    e := email.NewEmail()
    e.From = "MyEmail@gmail.com"
    e.To = "recipient_email@gmail.com"
    e.Subject = "Automatic Response"
    e.Text = "This is an automatic response."

    // Render the template
    err = tmpl.Execute(os.Stdout, e)
    if err != nil {
        log.Fatal(err)
    }

    // Send the email
    err = client.SendMail("smtp.gmail.com:587", e.From, e.To, []byte(e.Text))
    if err != nil {
        log.Fatal(err)
    }

    // Close the client
    client.Close()
}
