# Read me 

## Work flow

All tasks related to this project, including documenting, coding and testing are listed as issues. All contributions will be commits to this repository. Upon commit and push, the related issue is closed. 

Issues (tasks) are organized into week-based milestones. These are my equivalent of sprints.

Decisions over and thinking behind features and tech choices are documented in `docs` folder as if they are code. Writings in this folder should be dry and concise. Make sure they are ready for a third party to read at all times.

Use commitlint format in commit messages:

- `type(scope): detail`

Types include: `feat`, `bugfix`, `doc`, `refactor`, `test`.

## Main features

Take audiobooks and podcasts with transcripts. Segment text according to sentence boundaries and align recording segments with sentences.

Present recordings of sentences to the user. Prompt the user to type out what they hear. Depending on the proficiency of the user [^1], skip words that are too easy [^2].

[^1]: TBD - How to know the proficiency of a user is another question.

[^2]: TBD - How to determine which words are easy. Because some simple words tend to be difficult to understand, for example, you might miss one or two prepositions.

When the user finishes *n* sentences, we reveal the correct text and highlight the words they got wrong. We ask how familiar the user is with those words:

- unknown: “I have never seen this before”. 
- difficult to hear: “I know the meaning when I see it. I don't know the pronunciation.” 
- difficult to type: “I know the meaning and pronunciation. The spelling is difficult.”

For *unknown* words, we add a double sided flash card for it [^3]. For *DTH* and *DTT* words, we add “play and type” cards. 

[^3]: This flashcard takes too much effort to make.

The user will then be reminded of old cards on a schedule looser than the Anki algorithm.

### User stories

- When the user registers, they should be given a brief vocabulary test and a brief listening test. 
- When the user logs in, they should be shown a list of available articles. On each article, a progress bar indicates how much they've progressed in that article.
- When the user clicks an article, they are brought to their last unsolved sentence in that article.
- When the user is working on an article and is shown an unsolved sentence, they are shown a play button and 1) an empty input field; or 2) a sentence with hard to understand words hidden.

More details to come.

### Design

#### Double sided flash cards

- Front: `Äpple / Jag åt ett äpple.` -> `Apple / I ate an apple.`
- Back: `Apple / ä__ / (play sound >)` -> `äpple` [^3]

#### Play and type cards

- Front: `(Play sound >)`
- Back: `Det här är en mening som innehåller ett ___ ord.`

The user only needs to supply the word that's missing. 

## Tech stack

### Backend

GraphQL service running on Go. The reason for using Go is it's a good practice to learn it.

NLP features will be implemented as a standalone service in Python.

Communication between services are done using ZeroMQ, it seems to be a very simple alternative to other message brokers. 

### Frontend

The frontend web interface will be implemented twice. First in Vue.js then in React.js. The purpose of this is to keep myself familiar with both frameworks.

- Audio player: Wavesurfer.js

### Deployment

Static files: AWS S3

Others TBD
