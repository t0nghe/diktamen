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

export type userTrySent = {
  sentId: number;
  userInputWords: {
    indexInSent: number;
    inputText: string;
  }[];
};

export type inputValuesType = { [key: string]: string };

export type inputFieldsType = { [key: string]: HTMLInputElement };

export type examineSent = {
  sentId: number;
  indexInArticle: number;
  tryText?: string;
  sentWords?: sentWord[];
};
