# Gator CLI
Welcome to Gator, a RSS feed blog aggregator!

## Features 
- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

## Requirements
1. PostgreSQL
    * Gator requires a PostgreSQL database for storing and managing its data.
2. Go Programming Language
    * Gator is written in Go, so you'll need to have Go installed to build and run the program.

## Installation 
1. To install the Gator CLI, use the following command:
```
   go install github.com/MechamJonathan/blog_aggregator@latest
```

## Configuration
Gator CLI requires a configuration file to run. Here is how to set it up:
1. In your root directory, create the configuration file and name it '.gatorconfig.json'.
2. The JSON file should have this structure:
```
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```
3. Run database migrations. For example, using Goose:
```
goose postgres <connection_string> up
```
4. Generate SQL code. For example, using SQLC, run the following at your project's root directory:
```
sqlc generate
```
5. In your project directory build the application using:
```
go build
```

## Run Application
In the Gator root directory, run the following:
```
go run . <command> 
```
# Commands
- register: register as a user
  ```
  go run . register <name>
  ```
- login: login as a registered user
  ```
  go run . login <name>
  ```
- reset: reset the program and databases
  ```
  go run . reset
  ```
- users: view all registered users
  ```
  go run . users
  ```
- agg: start the blog aggregator
  ```
  go run . agg <time_between_reqs>
  ```
    * time_between_reqs examples: 1m, 1s, 1h
- addfeed: add feed
  ```
  go run . addfeed <name> <url>
  ```
- feeds: view all feeds
   ```
  go run . feeds
  ```
- follow: follow a feed
   ```
  go run . follow <name> <url>
  ```
- following: see all feeds current user is following
   ```
  go run . following
  ```
- unfollow: unfollow a feed
   ```
  go run . unfollow <name> <url>
  ```
- browse: browse all feeds current user is following
   ```
  go run . browse
  ```

## Example Workflow
1. register as a user
```
> blog_aggregator % go run . register jonathan

  User created successfully:
   * ID:      9b9b0710-1274-4202-8b85-7e6f8483d26e
   * Name:    jonathan
```

2. Login as user
```
> blog_aggregator % go run . login jonathan

  User switched successfully!
```

3. Add feed
```
> blog_aggregator % go run . addfeed HackerNews https://news.ycombinator.com/rss

  Feed added successfully:* ID:            f93a2b07-f25a-4266-b7a8-562efa8d633e
  * Created:       2025-01-20 19:09:17.086092 +0000 +0000
  * Updated:       2025-01-20 19:09:17.086092 +0000 +0000
  * Name:          HackerNews
  * URL:           https://news.ycombinator.com/rss
  * UserID:        9b9b0710-1274-4202-8b85-7e6f8483d26e
  User jonathan is now following 'HackerNews'
  =====================================
```

4. Start the aggregator (run in separate terminal window)
```
> blog_aggregator % go run . agg 1m

  2025/01/20 12:14:12 Collecting feeds every 1m0s...
  2025/01/20 12:14:12 Found a feed to fetch!
  2025/01/20 12:14:13 Feed HackerNews collected, 30 posts found
```

5. Browse feeds
```
> blog_aggregator % go run . browse

  Found 2 posts for user jonathan:
  Mon Jan 20 from HackerNews
  --- I am (not) a failure: Lessons learned from six failed startup attempts ---
      <a href="https://news.ycombinator.com/item?id=42771676">Comments</a>
  Link: http://blog.rongarret.info/2025/01/i-am-not-failure-lessons-learned-from.html
  =====================================
  Mon Jan 20 from HackerNews
  --- Seeking an IPL-V Interpreter ---
      <a href="https://news.ycombinator.com/item?id=42771413">Comments</a>
  Link: https://news.ycombinator.com/item?id=42771413
```
