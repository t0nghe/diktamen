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
