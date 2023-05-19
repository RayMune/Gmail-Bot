# Gmail-Bot
# Email Sentiment Analysis and Autoresponder

This Go program allows you to analyze the sentiment of an email and send an automatic response using the Gmail SMTP server. It utilizes the "sentiment" package for sentiment analysis and the "net/smtp" package for email communication.

## Prerequisites

Before running the program, ensure that you have the following:

- Go programming language installed on your system.
- A Gmail account for sending the automatic response.
- The email content stored in a file.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/RayMune/Gmail-Bot.git
   ```

2. Build the executable:

   ```bash
   go build
   ```

## Usage

1. Run the program:

   ```bash
   ./email-sentiment-analysis email_content.txt
   ```

   Replace `email_content.txt` with the path to your email content file.

2. The program will analyze the sentiment of the email content and print the sentiment and score to the console.

3. An automatic response will be sent to the recipient's email address specified in the code. Ensure you have provided valid Gmail credentials and valid recipient email address.

## Customization

To customize the program, you can modify the following:

- Gmail credentials: Replace `MyEmail@gmail.com` with your Gmail email address and `MyPassword` with your account password.
- Automatic response: Modify the subject and content of the email in the `e.Subject` and `e.Text` fields, respectively.
- Sentiment analysis: If needed, you can customize the sentiment analysis logic in the code by using a different sentiment analysis package or implementing your own sentiment analysis algorithm.

## Limitations

- The program currently only supports Gmail's SMTP server. If you want to use a different email provider, you will need to modify the SMTP settings accordingly.
- The sentiment analysis algorithm used in the "sentiment" package may have its own limitations and accuracy constraints.

## Contributing

Contributions to the project are welcome. If you have any improvements or suggestions, please submit a pull request or open an issue on the repository.
