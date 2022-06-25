# Read me 

## Work flow

All tasks related to this project, including documenting, coding and testing are listed as issues. All contributions will be commits to this repository. Upon commit and push, the related issue is closed. 

Issues (tasks) are organized into week-based milestones. These are my equivalent of sprints.

Decisions over and thinking behind features and tech choices are documented in `docs` folder as if they are code. Writings in this folder should be dry and concise. Make sure they are ready for a third party to read at all times.

## Main features

Take audiobooks and podcasts with transcripts. Segment text according to sentence boundaries and align recording segments with sentences.

Present recordings of sentences to the user. Prompt the user to type out what they hear. Depending on the proficiency of the user, skip words that are too easy [^1].

[^1]: Hello 

### User stories

TBD

## Tech stack

### Backend

GraphQL service running on Flask written in Python. The reason for using Python is there will be NLP tasks.

Ariadne is used as GraphQL server. 

### Frontend

The frontend web interface will be implemented twice. First in Vue.js then in React.js. The purpose of this is to keep myself familiar with both frameworks. 

### Deployment

Static files: AWS S3

Others TBD
