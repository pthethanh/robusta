package types

// Policy any
const (
	PolicyObjectAny  = "*"
	PolicyActionAny  = "*"
	PolicySubjectAny = "*"
)

// Policy effects
const (
	PolicyEffectAllow = "allow"
	PolicyEffectDeny  = "deny"
)

// Policy of article
const (
	PolicyObjectArticle       = "article"
	PolicyActionArticleCreate = "article:create"
	PolicyActionArticleDelete = "article:delete"
	PolicyActionArticleUpdate = "article:update"
	PolicyActionArticleRead   = "article:read"
)

// Policy of challenge
const (
	PolicyObjectChallenge       = "challenge"
	PolicyActionChallengeCreate = "challenge:create"
	PolicyActionChallengeDelete = "challenge:delete"
	PolicyActionChallengeUpdate = "challenge:update"
	PolicyActionChallengeRead   = "challenge:read"
)

// Policy of comment
const (
	PolicyObjectComment       = "comment"
	PolicyActionCommentCreate = "comment:create"
	PolicyActionCommentDelete = "comment:delete"
	PolicyActionCommentUpdate = "comment:update"
	PolicyActionCommentRead   = "comment:read"
)

// Policy of folder
const (
	PolicyObjectFolder         = "folder"
	PolicyActionFolderCreate   = "folder:create"
	PolicyActionFolderDelete   = "folder:delete"
	PolicyActionFolderUpdate   = "folder:update"
	PolicyActionFolderRead     = "folder:read"
	PolicyActionFolderReadList = "folder:read-list"
)

// Policy of solution
const (
	PolicyObjectSolution               = "solution"
	PolicyActionSolutionRead           = "solution:read"
	PolicyActionSolutionReadListDetail = "solution:read-list-detail"
)

// Policy of user
const (
	PolicyObjectUser         = "user"
	PolicyActionUserReadList = "solution:read-list"
)
