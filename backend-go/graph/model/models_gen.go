// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type FullArticle struct {
	ArticleID        string      `json:"articleId"`
	ArticleTitle     string      `json:"articleTitle"`
	ArticleSentCount int         `json:"articleSentCount"`
	UserNextUpIndex  string      `json:"userNextUpIndex"`
	Sentences        []*Sentence `json:"sentences"`
}

type InputMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type LenWordTuple struct {
	Index    int    `json:"index"`
	Letters  int    `json:"letters"`
	WordForm string `json:"wordForm"`
}

type LoginToken struct {
	Success bool    `json:"success"`
	Token   *string `json:"token"`
}

type Message struct {
	Success bool    `json:"success"`
	Message *string `json:"message"`
}

type SentDetails struct {
	SentID    string          `json:"sentId"`
	MediaURI  string          `json:"mediaUri"`
	WordForms []*LenWordTuple `json:"wordForms"`
}

type SentScore struct {
	SentID         string `json:"sentId"`
	MediaURI       string `json:"mediaUri"`
	OriginalText   string `json:"originalText"`
	Comparison     string `json:"comparison"`
	WordCount      int    `json:"wordCount"`
	IncorrectCount int    `json:"incorrectCount"`
}

type Sentence struct {
	SentID             string    `json:"sentId"`
	SentIndexInArticle int       `json:"sentIndexInArticle"`
	UserTried          bool      `json:"userTried"`
	UserAttempt        []*string `json:"userAttempt"`
}

type UserArticle struct {
	ArticleID        string `json:"articleId"`
	ArticleTitle     string `json:"articleTitle"`
	ArticleSentCount int    `json:"articleSentCount"`
	UserNextUpIndex  string `json:"userNextUpIndex"`
}

type UserAttempt struct {
	SentID                   string `json:"sentId"`
	InputWordFormsJSONString string `json:"inputWordFormsJsonString"`
}

type UserCredentials struct {
	Username     string  `json:"username"`
	PasswordHash string  `json:"passwordHash"`
	Email        *string `json:"email"`
}