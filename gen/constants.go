package gen

type key int

const (
	KeyPrincipalID         key    = iota
	KeyLoaders             key    = iota
	KeyExecutableSchema    key    = iota
	KeyJWTClaims           key    = iota
	KeyMutationTransaction key    = iota
	KeyMutationEvents      key    = iota
	SchemaSDL              string = `scalar Time

type Query {
  file(id: ID, q: String, filter: FileFilterType): File
  files(offset: Int, limit: Int = 30, q: String, sort: [FileSortType!], filter: FileFilterType): FileResultType!
}

type Mutation {
  createFile(input: FileCreateInput!): File!
  updateFile(id: ID!, input: FileUpdateInput!): File!
  deleteFile(id: ID!): File!
  deleteAllFiles: Boolean!
}

enum ObjectSortType {
  ASC
  DESC
}

type File @key(fields: "id") {
  id: ID!
  email: String
  UID: ID
  Size: Int
  ContentType: String
  URL: String
  Name: String
  Reference: String
  ReferenceID: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
}

input FileCreateInput {
  id: ID
  email: String
  UID: ID
  Size: Int
  ContentType: String
  URL: String
  Name: String
  Reference: String
  ReferenceID: ID
}

input FileUpdateInput {
  email: String
  UID: ID
  Size: Int
  ContentType: String
  URL: String
  Name: String
  Reference: String
  ReferenceID: ID
}

input FileSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  email: ObjectSortType
  emailMin: ObjectSortType
  emailMax: ObjectSortType
  UID: ObjectSortType
  UIDMin: ObjectSortType
  UIDMax: ObjectSortType
  Size: ObjectSortType
  SizeMin: ObjectSortType
  SizeMax: ObjectSortType
  SizeAvg: ObjectSortType
  ContentType: ObjectSortType
  ContentTypeMin: ObjectSortType
  ContentTypeMax: ObjectSortType
  URL: ObjectSortType
  URLMin: ObjectSortType
  URLMax: ObjectSortType
  Name: ObjectSortType
  NameMin: ObjectSortType
  NameMax: ObjectSortType
  Reference: ObjectSortType
  ReferenceMin: ObjectSortType
  ReferenceMax: ObjectSortType
  ReferenceID: ObjectSortType
  ReferenceIDMin: ObjectSortType
  ReferenceIDMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
}

input FileFilterType {
  AND: [FileFilterType!]
  OR: [FileFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_null: Boolean
  email: String
  emailMin: String
  emailMax: String
  email_ne: String
  emailMin_ne: String
  emailMax_ne: String
  email_gt: String
  emailMin_gt: String
  emailMax_gt: String
  email_lt: String
  emailMin_lt: String
  emailMax_lt: String
  email_gte: String
  emailMin_gte: String
  emailMax_gte: String
  email_lte: String
  emailMin_lte: String
  emailMax_lte: String
  email_in: [String!]
  emailMin_in: [String!]
  emailMax_in: [String!]
  email_like: String
  emailMin_like: String
  emailMax_like: String
  email_prefix: String
  emailMin_prefix: String
  emailMax_prefix: String
  email_suffix: String
  emailMin_suffix: String
  emailMax_suffix: String
  email_null: Boolean
  UID: ID
  UIDMin: ID
  UIDMax: ID
  UID_ne: ID
  UIDMin_ne: ID
  UIDMax_ne: ID
  UID_gt: ID
  UIDMin_gt: ID
  UIDMax_gt: ID
  UID_lt: ID
  UIDMin_lt: ID
  UIDMax_lt: ID
  UID_gte: ID
  UIDMin_gte: ID
  UIDMax_gte: ID
  UID_lte: ID
  UIDMin_lte: ID
  UIDMax_lte: ID
  UID_in: [ID!]
  UIDMin_in: [ID!]
  UIDMax_in: [ID!]
  UID_null: Boolean
  Size: Int
  SizeMin: Int
  SizeMax: Int
  SizeAvg: Int
  Size_ne: Int
  SizeMin_ne: Int
  SizeMax_ne: Int
  SizeAvg_ne: Int
  Size_gt: Int
  SizeMin_gt: Int
  SizeMax_gt: Int
  SizeAvg_gt: Int
  Size_lt: Int
  SizeMin_lt: Int
  SizeMax_lt: Int
  SizeAvg_lt: Int
  Size_gte: Int
  SizeMin_gte: Int
  SizeMax_gte: Int
  SizeAvg_gte: Int
  Size_lte: Int
  SizeMin_lte: Int
  SizeMax_lte: Int
  SizeAvg_lte: Int
  Size_in: [Int!]
  SizeMin_in: [Int!]
  SizeMax_in: [Int!]
  SizeAvg_in: [Int!]
  Size_null: Boolean
  ContentType: String
  ContentTypeMin: String
  ContentTypeMax: String
  ContentType_ne: String
  ContentTypeMin_ne: String
  ContentTypeMax_ne: String
  ContentType_gt: String
  ContentTypeMin_gt: String
  ContentTypeMax_gt: String
  ContentType_lt: String
  ContentTypeMin_lt: String
  ContentTypeMax_lt: String
  ContentType_gte: String
  ContentTypeMin_gte: String
  ContentTypeMax_gte: String
  ContentType_lte: String
  ContentTypeMin_lte: String
  ContentTypeMax_lte: String
  ContentType_in: [String!]
  ContentTypeMin_in: [String!]
  ContentTypeMax_in: [String!]
  ContentType_like: String
  ContentTypeMin_like: String
  ContentTypeMax_like: String
  ContentType_prefix: String
  ContentTypeMin_prefix: String
  ContentTypeMax_prefix: String
  ContentType_suffix: String
  ContentTypeMin_suffix: String
  ContentTypeMax_suffix: String
  ContentType_null: Boolean
  URL: String
  URLMin: String
  URLMax: String
  URL_ne: String
  URLMin_ne: String
  URLMax_ne: String
  URL_gt: String
  URLMin_gt: String
  URLMax_gt: String
  URL_lt: String
  URLMin_lt: String
  URLMax_lt: String
  URL_gte: String
  URLMin_gte: String
  URLMax_gte: String
  URL_lte: String
  URLMin_lte: String
  URLMax_lte: String
  URL_in: [String!]
  URLMin_in: [String!]
  URLMax_in: [String!]
  URL_like: String
  URLMin_like: String
  URLMax_like: String
  URL_prefix: String
  URLMin_prefix: String
  URLMax_prefix: String
  URL_suffix: String
  URLMin_suffix: String
  URLMax_suffix: String
  URL_null: Boolean
  Name: String
  NameMin: String
  NameMax: String
  Name_ne: String
  NameMin_ne: String
  NameMax_ne: String
  Name_gt: String
  NameMin_gt: String
  NameMax_gt: String
  Name_lt: String
  NameMin_lt: String
  NameMax_lt: String
  Name_gte: String
  NameMin_gte: String
  NameMax_gte: String
  Name_lte: String
  NameMin_lte: String
  NameMax_lte: String
  Name_in: [String!]
  NameMin_in: [String!]
  NameMax_in: [String!]
  Name_like: String
  NameMin_like: String
  NameMax_like: String
  Name_prefix: String
  NameMin_prefix: String
  NameMax_prefix: String
  Name_suffix: String
  NameMin_suffix: String
  NameMax_suffix: String
  Name_null: Boolean
  Reference: String
  ReferenceMin: String
  ReferenceMax: String
  Reference_ne: String
  ReferenceMin_ne: String
  ReferenceMax_ne: String
  Reference_gt: String
  ReferenceMin_gt: String
  ReferenceMax_gt: String
  Reference_lt: String
  ReferenceMin_lt: String
  ReferenceMax_lt: String
  Reference_gte: String
  ReferenceMin_gte: String
  ReferenceMax_gte: String
  Reference_lte: String
  ReferenceMin_lte: String
  ReferenceMax_lte: String
  Reference_in: [String!]
  ReferenceMin_in: [String!]
  ReferenceMax_in: [String!]
  Reference_like: String
  ReferenceMin_like: String
  ReferenceMax_like: String
  Reference_prefix: String
  ReferenceMin_prefix: String
  ReferenceMax_prefix: String
  Reference_suffix: String
  ReferenceMin_suffix: String
  ReferenceMax_suffix: String
  Reference_null: Boolean
  ReferenceID: ID
  ReferenceIDMin: ID
  ReferenceIDMax: ID
  ReferenceID_ne: ID
  ReferenceIDMin_ne: ID
  ReferenceIDMax_ne: ID
  ReferenceID_gt: ID
  ReferenceIDMin_gt: ID
  ReferenceIDMax_gt: ID
  ReferenceID_lt: ID
  ReferenceIDMin_lt: ID
  ReferenceIDMax_lt: ID
  ReferenceID_gte: ID
  ReferenceIDMin_gte: ID
  ReferenceIDMax_gte: ID
  ReferenceID_lte: ID
  ReferenceIDMin_lte: ID
  ReferenceIDMax_lte: ID
  ReferenceID_in: [ID!]
  ReferenceIDMin_in: [ID!]
  ReferenceIDMax_in: [ID!]
  ReferenceID_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
  createdBy_null: Boolean
}

type FileResultType {
  items: [File!]!
  count: Int!
}`
)