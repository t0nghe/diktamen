import gql from "graphql-tag";

export const incorrSents = gql`
  query examineIncorrectSents($articleId: Int!) {
    examineIncorrectSents(articleId: $articleId) {
      sentId
      indexInArticle
      sentWords {
        length
        isCloze
        wordform
        indexInSent
        lastInputText
        lastInputScore
      }
    }
  }
`;

export const corrSents = gql`
  query examineCorrectSents($articleId: Int!) {
    examineCorrectSents(articleId: $articleId) {
      sentId
      indexInArticle
      tryText
    }
  }
`;

export const seenArticles = gql`
  query listUserArticles {
    listUserArticles {
      articleId
      articleTitle
      articleSentCount
      articleDescription
      userFinishedIndex
    }
  }
`;

export const unseenArticles = gql`
  query unseenArticles {
    listUserUnseenArticles {
      articleId
      articleTitle
      articleSentCount
      articleDescription
      userFinishedIndex
    }
  }
`;

export const unseenSents = gql`
  query displayUnseenSents($articleId: Int!) {
    displayUnseenSents(articleId: $articleId) {
      sentId
      indexInArticle
      mediaUri
      sentWords {
        length
        isCloze
        wordform
        indexInSent
      }
    }
  }
`;

export const seenSents = gql`
  query displaySeenSents($articleId: Int!) {
    displaySeenSents(articleId: $articleId) {
      sentId
      indexInArticle
      tryText
    }
  }
`;

export const mutationTrySent = gql`
  mutation trySent($sentId: Int!, $userInputJson: String!) {
    trySent(input: { sentId: $sentId, userInputJson: $userInputJson }) {
      sentId
      indexInArticle
      tryText
    }
  }
`;

export const singleArticle = gql`
  query getUserArticle($articleId: Int!) {
    getUserArticle(articleId: $articleId) {
      articleId
      articleTitle
      articleDescription
      articleSentCount
      userFinishedIndex
    }
  }
`;

export const dueSents = gql`
  query displayDueSents($days: Int) {
    displayDueSents(daysAhead: $days) {
      sentId
      mediaUri
      sentWords {
        length
        isCloze
        wordform
        indexInSent
        lastInputText
        lastInputScore
      }
    }
  }
`;
