# All Actions

This document list all the actions a user can do. 

## A01 - register

- Related story: US01
- Related mutation: userSignUp
    - input: username, email, passwordHash
    - response: success message, failure message (username taken, email taken)

* Frontend needs to hash password.

## A02 - login

- Related story: US02
- Related mutation: userLogIn
    - input: username, passwordHash
    - response: If success, auth token. If failure, failure message.

## A03 - list articles

- Related story: US04
- Related query: listUserArticles
    - input: username, auth header
    - response: [ { articleId, articleTitle, articleSentCount, userNextUpIndex }]

If `articleUserSentCount` is 0, we know this user hasn't started listening to this article. Otherwise, we display a progress bar.

## A04 - get the full article

When the user clicks on an article, we get: 1) ordered array of sentence ID's in this article; 2) ordered array contained all the user input for sentences they've attempted; 3) next sentence to attempt.

- Related story: US05
- Related query: getUserFullArticle
    - input: username, auth header, article ID
    - response: [
        articleId, 
        articleTitle,
        articleSentCount,
        userNextUpIndex,
        sentences:  {
            sentIndexInArticle int, sentId string, userTried bool, userAttempt []string
            }
        ]

## A05 - get the audio for a sentence, so that the user can play

We will display a waveform (hopefully), a play button, and a row of input fields. Each input field is for one word. Length of this input field varies according to the word length.

- Related query: fetchSentAudio
    - input: username, auth header, sentId
    - response: [
        sentId
        mediaUri string, 
        wordForms []{int, string} // How many letters are there in each word in this sentence. We call them “word forms” because we want to say they are not lemmas.
        // If int is `-1`, it means we don't need the user to input this word.
    ]

## A06 - user attempts supply the text of a sentence

- related mutation: userAttempt

## A07 - compare user attempts to correct text

When the user finishes dictating one article, there will be a “compare with original text” button. When clicked, the frontend shows comparison results. 

- related query: scoreArticle // This is indeed yet another query.
- input: username, auth header, articleId
- response: [
        { mediaUri string, userAttempt string, originalText string, comparison string, acc float }
    ]

Hopefully we'll compare user input with original text as it is submitted. HTML tags `<span class="err">` will be added. Then the result of comparison is inserted to DB. (TBD Is this doable?) At the same time, acc is calculated based on edit distance. (TBD Haven't thought through how this can be done.)

## A08 - add mistaken sentences to review deck

TBD... How do we add mistaken sentences to review deck? Do we do it automatically or let the user decide? 

- input: username, auth header, sentId

## A09 - review mistaken sentences

TBD... 