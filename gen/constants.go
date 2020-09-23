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
  uid: ID
  name: String
  size: Int
  contentType: String
  url: String
  reference: String
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
}

input FileCreateInput {
  id: ID
  uid: ID
  name: String
  size: Int
  contentType: String
  url: String
  reference: String
}

input FileUpdateInput {
  uid: ID
  name: String
  size: Int
  contentType: String
  url: String
  reference: String
}

input FileSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  uid: ObjectSortType
  uidMin: ObjectSortType
  uidMax: ObjectSortType
  name: ObjectSortType
  nameMin: ObjectSortType
  nameMax: ObjectSortType
  size: ObjectSortType
  sizeMin: ObjectSortType
  sizeMax: ObjectSortType
  sizeAvg: ObjectSortType
  contentType: ObjectSortType
  contentTypeMin: ObjectSortType
  contentTypeMax: ObjectSortType
  url: ObjectSortType
  urlMin: ObjectSortType
  urlMax: ObjectSortType
  reference: ObjectSortType
  referenceMin: ObjectSortType
  referenceMax: ObjectSortType
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
  uid: ID
  uidMin: ID
  uidMax: ID
  uid_ne: ID
  uidMin_ne: ID
  uidMax_ne: ID
  uid_gt: ID
  uidMin_gt: ID
  uidMax_gt: ID
  uid_lt: ID
  uidMin_lt: ID
  uidMax_lt: ID
  uid_gte: ID
  uidMin_gte: ID
  uidMax_gte: ID
  uid_lte: ID
  uidMin_lte: ID
  uidMax_lte: ID
  uid_in: [ID!]
  uidMin_in: [ID!]
  uidMax_in: [ID!]
  uid_null: Boolean
  name: String
  nameMin: String
  nameMax: String
  name_ne: String
  nameMin_ne: String
  nameMax_ne: String
  name_gt: String
  nameMin_gt: String
  nameMax_gt: String
  name_lt: String
  nameMin_lt: String
  nameMax_lt: String
  name_gte: String
  nameMin_gte: String
  nameMax_gte: String
  name_lte: String
  nameMin_lte: String
  nameMax_lte: String
  name_in: [String!]
  nameMin_in: [String!]
  nameMax_in: [String!]
  name_like: String
  nameMin_like: String
  nameMax_like: String
  name_prefix: String
  nameMin_prefix: String
  nameMax_prefix: String
  name_suffix: String
  nameMin_suffix: String
  nameMax_suffix: String
  name_null: Boolean
  size: Int
  sizeMin: Int
  sizeMax: Int
  sizeAvg: Int
  size_ne: Int
  sizeMin_ne: Int
  sizeMax_ne: Int
  sizeAvg_ne: Int
  size_gt: Int
  sizeMin_gt: Int
  sizeMax_gt: Int
  sizeAvg_gt: Int
  size_lt: Int
  sizeMin_lt: Int
  sizeMax_lt: Int
  sizeAvg_lt: Int
  size_gte: Int
  sizeMin_gte: Int
  sizeMax_gte: Int
  sizeAvg_gte: Int
  size_lte: Int
  sizeMin_lte: Int
  sizeMax_lte: Int
  sizeAvg_lte: Int
  size_in: [Int!]
  sizeMin_in: [Int!]
  sizeMax_in: [Int!]
  sizeAvg_in: [Int!]
  size_null: Boolean
  contentType: String
  contentTypeMin: String
  contentTypeMax: String
  contentType_ne: String
  contentTypeMin_ne: String
  contentTypeMax_ne: String
  contentType_gt: String
  contentTypeMin_gt: String
  contentTypeMax_gt: String
  contentType_lt: String
  contentTypeMin_lt: String
  contentTypeMax_lt: String
  contentType_gte: String
  contentTypeMin_gte: String
  contentTypeMax_gte: String
  contentType_lte: String
  contentTypeMin_lte: String
  contentTypeMax_lte: String
  contentType_in: [String!]
  contentTypeMin_in: [String!]
  contentTypeMax_in: [String!]
  contentType_like: String
  contentTypeMin_like: String
  contentTypeMax_like: String
  contentType_prefix: String
  contentTypeMin_prefix: String
  contentTypeMax_prefix: String
  contentType_suffix: String
  contentTypeMin_suffix: String
  contentTypeMax_suffix: String
  contentType_null: Boolean
  url: String
  urlMin: String
  urlMax: String
  url_ne: String
  urlMin_ne: String
  urlMax_ne: String
  url_gt: String
  urlMin_gt: String
  urlMax_gt: String
  url_lt: String
  urlMin_lt: String
  urlMax_lt: String
  url_gte: String
  urlMin_gte: String
  urlMax_gte: String
  url_lte: String
  urlMin_lte: String
  urlMax_lte: String
  url_in: [String!]
  urlMin_in: [String!]
  urlMax_in: [String!]
  url_like: String
  urlMin_like: String
  urlMax_like: String
  url_prefix: String
  urlMin_prefix: String
  urlMax_prefix: String
  url_suffix: String
  urlMin_suffix: String
  urlMax_suffix: String
  url_null: Boolean
  reference: String
  referenceMin: String
  referenceMax: String
  reference_ne: String
  referenceMin_ne: String
  referenceMax_ne: String
  reference_gt: String
  referenceMin_gt: String
  referenceMax_gt: String
  reference_lt: String
  referenceMin_lt: String
  referenceMax_lt: String
  reference_gte: String
  referenceMin_gte: String
  referenceMax_gte: String
  reference_lte: String
  referenceMin_lte: String
  referenceMax_lte: String
  reference_in: [String!]
  referenceMin_in: [String!]
  referenceMax_in: [String!]
  reference_like: String
  referenceMin_like: String
  referenceMax_like: String
  reference_prefix: String
  referenceMin_prefix: String
  referenceMax_prefix: String
  reference_suffix: String
  referenceMin_suffix: String
  referenceMax_suffix: String
  reference_null: Boolean
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
