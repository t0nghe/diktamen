# Diktamen

Hej! 

Thanks for actually checking out this repo.

This is the codebase for [Diktamen](https://d.tonghe.xyz/), a web app that helps learners learn Swedish through dictation.

In this app, the learner chooses an article, then types out what they hear sentence by sentence. When there are mistakes, they will be reminded about those mistakes in a few hours or a few days, based on how well they did last time.

This web app runs on Go on the backend and Vue.js on the frontend. Audio files are hosted on AWS S3. User data and language data are stored in a MySQL cluster on AWS RDS.

The frontend and backend exchange data via GraphQL. Resolvers that handle queries and mutations are implemented using [gqlgen](https://github.com/99designs/gqlgen). On the frontend, [Apollo Client](https://www.apollographql.com/) and [vue-apollo](https://apollo.vuejs.org/) are used. 

The markdown files in `docs` folder are outdated. But they might reflect some of my thought process. 

The development of this project took a little less than seven weeks. The code might feel hasty here and there. I would come back and refactor when I think of better ways to do things.

## Backend

`gqlgen` is pretty handy. After defining queries and mutations in `schema.graphqls`, you can use it to generate the underlying Go code to handle them (`/graph/generated`) and to define data types for your Go code (`/graph/model`).

You only need to write code for the business logic and database queries. Such code is found under [`/internal`](https://github.com/t0nghe/diktamen/tree/master/backend-go/internal).

[`/internal/queries`](https://github.com/t0nghe/diktamen/tree/master/backend-go/internal/queries) contains all the SQL queries and data transformation. [`/internal/auth`](https://github.com/t0nghe/diktamen/tree/master/backend-go/internal/auth) contains code for signing up and logging in and for handling token passed in with HTTP headers. [`/internal/pkg`](https://github.com/t0nghe/diktamen/tree/master/backend-go/internal/pkg) contains code for database connection, scoring of user input using Levenshtein distance, and processing of tokens using JWT.

## Frontend

[`/src/views`](https://github.com/t0nghe/diktamen/tree/master/frontend-vue/src/views) contains views that routes point to. [`src/components`](https://github.com/t0nghe/diktamen/tree/master/frontend-vue/src/components) contains components that make up these views.

Notably, on the learning page ([`/src/views/LearnView.vue`](https://github.com/t0nghe/diktamen/blob/master/frontend-vue/src/views/LearnView.vue)), when the data for the sentence that the user is going to listen to ([`sentence-trying`](https://github.com/t0nghe/diktamen/blob/master/frontend-vue/src/components/Sentence/SentenceTrying.vue)) is loaded, the first empty input box gets cursor focus. As the user types, when an input box contains enough letters, the focus is moved to the next empty input box.

At this point, the user can control play and pause of the audio either by 1) pressing spacebar or 2) clicking on the play-pause button, which is displayed underneath the text. This is done with the [`play-pause`](https://github.com/t0nghe/diktamen/blob/master/frontend-vue/src/components/Interaction/PlayPause.vue) component, where two divs are used: one underneath text to display the button and the other transparent, above text, thus clickable.

## NLP-Python

Under `nlp-python` folder, there are a few scripts that I used to populate database.
