# 🚂 Movie Engine

Movie engine is a toy project to search movies semantically built in Go, templ, tailwind and Postgres.

The app is live [here](https://movie-engine-zjnztlkezq-uc.a.run.app/)

It works by generating embedding vectors from an extensive selection of movie plots. Each time the user enters a query, their input is converted into a vector. This vector is then compared with those of the movies, retrieving the closest matches to their search.

Open AI [Embeddings API](https://platform.openai.com/docs/guides/embeddings) was used to create the movie plots and user queries embeddings. `text-embedding-3-small` model is very very cheap and produces good results.

[OMDB Api](www.omdbapi.com) is the primary source of movies. I plan to extend this a lot.

Each embedding vector has multiple dimensions and look like this:
```
[
  -0.04135531,
  0.03265167,
  -0.04588954,
  0.026488766,
  -0.027335677,
  0.0027182582,
  0.0069316397,
  0.0003070052,
  0.0059967805,
  ...
]
```
When having for both plot and user's query, we can calculate the distance between this vectors and see if they're talking about similar things.

For vector math [pgvector](https://github.com/pgvector/pgvector) did all the work.
Embedding SQL queries were as simple as:
```sql
SELECT id, title, year, plot, runtime, director, poster, ratings FROM omdb_movies ORDER BY embedding <-> '[-0.04135531, 0.03265167,-0.04588954,0.026488766,-0.027335677, 0.0027182582, ...]' LIMIT 5
```
This would return the 5 closest records based on the distance.

### Running the project

1. Install [air](https://github.com/cosmtrek/air) `go install github.com/cosmtrek/air@latest`
2. Install [templ](https://github.com/a-h/templ) `go install github.com/a-h/templ/cmd/templ@latest`
3. Install [node](https://nodejs.org/) (required for tailwind)
4. Create a .env file following env.example template and add your DB and Open AI API key
5. `go mod tidy`
6. `npm i`
7. `air`

### Roadmap
- [ ] Database schema and migrations
- [ ] Movie scraper
- [ ] Save user queries (movie rank)
- [ ] Scheduled scraper for new movies
- [ ] 2000+ movies
- [ ] Language options
