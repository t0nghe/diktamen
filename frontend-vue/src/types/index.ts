export type seenArticleType = {
  articleId: number;
  articleTitle: string;
  articleSentCount: number;
  userFinishedIndex: number;
};

export type sentWord = {
  length: number;
  isCloze: boolean;
  wordform: string;
  indexInSent: number;
  lastInputText?: string;
  lastInputScore?: number;
};
