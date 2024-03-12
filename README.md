## About this repository
This is a basic discord bot written in Golang with basic message and interaction commands.

## Installation
1. Go to the [Discord developer portal](https://discord.com/developers)
2. Create a new application
3. Ensure that you allow the bot to receive messages by going to the bot tab and toggling on `message content intent`.
4. Add bot to your server by going to the OAuth2 tab and generating a URL to add it to your server.  Under the settings please tick scopes `bot` and permissions `administrator` (this is only for testing purposes, for production bots please ensure you select specific scopes).  Copy generated URL into any browser and add it to your server.  
5. Get the token for the bot by going to the bot tab in the [Discord developer portal](https://discord.com/developers) (BOT_TOKEN for env variable).  If this is your first time, click the `reset token` button to generate your token - please save this to a secure place for future use.
6. Clone the repository

   ```sh
   git clone https://github.com/artukam/go-test-discord-bot.git
   ```

7. Install dependencies

   ```sh
   cd go-test-discord-bot
   go mod download all
   ```

8. Create your local `.env` variables file by copying the `.env.example`
    ```sh
    cp .env.example .env
    ```

9. Populate your local `.env` with the following variables from your discord dashboard:

    ```dotenv
    BOT_TOKEN=Your discord bot token
    APP_ID=Your discord bot application id
    ```
    * BOT_TOKEN: Get the token for the bot by going to the bot tab in the [Discord developer portal](https://discord.com/developers) (BOT_TOKEN for env variable).  If this is your first time, click the `reset token` button to generate your token - please save this to a secure place for future use.
    * APP_ID: Get the app id from the General information tab in the [Discord developer portal](https://discord.com/developers)

10. Start your discord bot
    ```sh
    go run main.go  
    ```