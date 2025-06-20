import { GraphQLClient, RequestOptions } from 'graphql-request';
import gql from 'graphql-tag';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
type GraphQLClientRequestHeaders = RequestOptions['requestHeaders'];
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Time: { input: any; output: any; }
  UUID: { input: any; output: any; }
};

export type AddRevisionInput = {
  id: Scalars['UUID']['input'];
  parent: Scalars['UUID']['input'];
  revision: CreateRecipeRevisionInput;
  slug: Scalars['String']['input'];
};

export type CreateRecipeInput = {
  forkdFrom?: InputMaybe<Scalars['UUID']['input']>;
  private: Scalars['Boolean']['input'];
  revision: CreateRecipeRevisionInput;
  slug: Scalars['String']['input'];
};

export type CreateRecipeRevisionIngredient = {
  comment?: InputMaybe<Scalars['String']['input']>;
  ingredient: Scalars['String']['input'];
  quantity: Scalars['Float']['input'];
  unit: Scalars['String']['input'];
};

export type CreateRecipeRevisionInput = {
  changeComment?: InputMaybe<Scalars['String']['input']>;
  description?: InputMaybe<Scalars['String']['input']>;
  ingredients: Array<CreateRecipeRevisionIngredient>;
  photo?: InputMaybe<Scalars['String']['input']>;
  steps: Array<CreateRecipeRevisionStep>;
  tags: Array<Scalars['String']['input']>;
  title: Scalars['String']['input'];
};

export type CreateRecipeRevisionStep = {
  instruction: Scalars['String']['input'];
  photo?: InputMaybe<Scalars['String']['input']>;
  step: Scalars['Int']['input'];
};

export type Ingredient = {
  __typename?: 'Ingredient';
  description?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type ListRecipeInput = {
  authorId?: InputMaybe<Scalars['UUID']['input']>;
  limit?: InputMaybe<Scalars['Int']['input']>;
  nextCursor?: InputMaybe<Scalars['String']['input']>;
  publishEnd?: InputMaybe<Scalars['Time']['input']>;
  publishStart?: InputMaybe<Scalars['Time']['input']>;
  sortCol?: InputMaybe<ListRecipeSortCol>;
  sortDir?: InputMaybe<SortDir>;
};

export enum ListRecipeSortCol {
  PublishDate = 'PUBLISH_DATE',
  Slug = 'SLUG'
}

export type ListRevisionsInput = {
  limit?: InputMaybe<Scalars['Int']['input']>;
  nextCursor?: InputMaybe<Scalars['String']['input']>;
  parentId?: InputMaybe<Scalars['UUID']['input']>;
  publishEnd?: InputMaybe<Scalars['Time']['input']>;
  publishStart?: InputMaybe<Scalars['Time']['input']>;
  recipeId?: InputMaybe<Scalars['UUID']['input']>;
  sortCol?: InputMaybe<ListRecipeSortCol>;
  sortDir?: InputMaybe<SortDir>;
};

export type LoginResponse = {
  __typename?: 'LoginResponse';
  token: Scalars['String']['output'];
  user: User;
};

export type MeasurementUnit = {
  __typename?: 'MeasurementUnit';
  description?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  recipe?: Maybe<RecipeMutation>;
  user?: Maybe<UserMutation>;
};

export type PaginatedRecipeRevisions = PaginatedResult & {
  __typename?: 'PaginatedRecipeRevisions';
  items: Array<RecipeRevision>;
  pagination: PaginationInfo;
};

export type PaginatedRecipes = PaginatedResult & {
  __typename?: 'PaginatedRecipes';
  items: Array<Recipe>;
  pagination: PaginationInfo;
};

export type PaginatedResult = {
  pagination: PaginationInfo;
};

export type PaginationInfo = {
  __typename?: 'PaginationInfo';
  count: Scalars['Int']['output'];
  nextCursor?: Maybe<Scalars['String']['output']>;
};

export type Query = {
  __typename?: 'Query';
  recipe?: Maybe<RecipeQuery>;
  user?: Maybe<UserQuery>;
};

export type Recipe = {
  __typename?: 'Recipe';
  author: User;
  featuredRevision?: Maybe<RecipeRevision>;
  forkedFrom?: Maybe<RecipeRevision>;
  id: Scalars['UUID']['output'];
  initialPublishDate: Scalars['Time']['output'];
  private: Scalars['Boolean']['output'];
  revisions: PaginatedRecipeRevisions;
  slug: Scalars['String']['output'];
};


export type RecipeRevisionsArgs = {
  input?: InputMaybe<ListRevisionsInput>;
};

export type RecipeIngredient = {
  __typename?: 'RecipeIngredient';
  comment?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  ingredient: Ingredient;
  quantity: Scalars['Float']['output'];
  revision: RecipeRevision;
  unit: MeasurementUnit;
};

export type RecipeMutation = {
  __typename?: 'RecipeMutation';
  addRevision: RecipeRevision;
  create: Recipe;
};


export type RecipeMutationAddRevisionArgs = {
  input: AddRevisionInput;
};


export type RecipeMutationCreateArgs = {
  input: CreateRecipeInput;
};

export type RecipeQuery = {
  __typename?: 'RecipeQuery';
  byId?: Maybe<Recipe>;
  bySlug?: Maybe<Recipe>;
  list: PaginatedRecipes;
};


export type RecipeQueryByIdArgs = {
  id: Scalars['UUID']['input'];
};


export type RecipeQueryBySlugArgs = {
  authorDisplayName: Scalars['String']['input'];
  slug: Scalars['String']['input'];
};


export type RecipeQueryListArgs = {
  input?: InputMaybe<ListRecipeInput>;
};

export type RecipeRevision = {
  __typename?: 'RecipeRevision';
  changeComment?: Maybe<Scalars['String']['output']>;
  id: Scalars['UUID']['output'];
  ingredients: Array<RecipeIngredient>;
  parent?: Maybe<RecipeRevision>;
  photo?: Maybe<Scalars['String']['output']>;
  publishDate: Scalars['Time']['output'];
  rating?: Maybe<Scalars['Float']['output']>;
  recipe: Recipe;
  recipeDescription?: Maybe<Scalars['String']['output']>;
  steps: Array<RecipeStep>;
  title: Scalars['String']['output'];
};

export type RecipeStep = {
  __typename?: 'RecipeStep';
  content: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  index: Scalars['Int']['output'];
  photo?: Maybe<Scalars['String']['output']>;
  revision: RecipeRevision;
};

export enum SortDir {
  Asc = 'ASC',
  Desc = 'DESC'
}

export type Tag = {
  __typename?: 'Tag';
  description?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  userGenerated: Scalars['Boolean']['output'];
};

export type User = {
  __typename?: 'User';
  displayName: Scalars['String']['output'];
  email: Scalars['String']['output'];
  id: Scalars['UUID']['output'];
  joinDate: Scalars['Time']['output'];
  photo?: Maybe<Scalars['String']['output']>;
  recipes: PaginatedRecipes;
  updatedAt: Scalars['Time']['output'];
};


export type UserRecipesArgs = {
  input?: InputMaybe<ListRecipeInput>;
};

export type UserMutation = {
  __typename?: 'UserMutation';
  login: LoginResponse;
  logout: Scalars['Boolean']['output'];
  requestMagicLink?: Maybe<Scalars['String']['output']>;
  signup?: Maybe<Scalars['String']['output']>;
  update: User;
};


export type UserMutationLoginArgs = {
  code: Scalars['String']['input'];
  token: Scalars['String']['input'];
};


export type UserMutationRequestMagicLinkArgs = {
  email: Scalars['String']['input'];
};


export type UserMutationSignupArgs = {
  displayName: Scalars['String']['input'];
  email: Scalars['String']['input'];
};


export type UserMutationUpdateArgs = {
  input: UserUpdateInput;
};

export type UserQuery = {
  __typename?: 'UserQuery';
  byDisplayName?: Maybe<User>;
  byEmail?: Maybe<User>;
  byId?: Maybe<User>;
  current?: Maybe<User>;
};


export type UserQueryByDisplayNameArgs = {
  displayName: Scalars['String']['input'];
};


export type UserQueryByEmailArgs = {
  email: Scalars['String']['input'];
};


export type UserQueryByIdArgs = {
  id: Scalars['UUID']['input'];
};

export type UserUpdateInput = {
  displayName?: InputMaybe<Scalars['String']['input']>;
  photo?: InputMaybe<Scalars['String']['input']>;
};

export type CheckUserSignupQueryVariables = Exact<{
  email: Scalars['String']['input'];
  displayName: Scalars['String']['input'];
}>;


export type CheckUserSignupQuery = { __typename?: 'Query', user?: { __typename?: 'UserQuery', byEmail?: { __typename?: 'User', email: string } | null, byDisplayName?: { __typename?: 'User', displayName: string } | null } | null };

export type ListRecipesQueryVariables = Exact<{ [key: string]: never; }>;


export type ListRecipesQuery = { __typename?: 'Query', recipe?: { __typename?: 'RecipeQuery', list: { __typename?: 'PaginatedRecipes', items: Array<{ __typename?: 'Recipe', slug: string, id: any, author: { __typename?: 'User', displayName: string }, featuredRevision?: { __typename?: 'RecipeRevision', recipeDescription?: string | null, publishDate: any, rating?: number | null, title: string } | null }> } } | null };

export type LoginMutationVariables = Exact<{
  token: Scalars['String']['input'];
  code: Scalars['String']['input'];
}>;


export type LoginMutation = { __typename?: 'Mutation', user?: { __typename?: 'UserMutation', login: { __typename?: 'LoginResponse', token: string } } | null };

export type LogoutMutationVariables = Exact<{ [key: string]: never; }>;


export type LogoutMutation = { __typename?: 'Mutation', user?: { __typename?: 'UserMutation', logout: boolean } | null };

export type RequestMagicLinkMutationVariables = Exact<{
  email: Scalars['String']['input'];
}>;


export type RequestMagicLinkMutation = { __typename?: 'Mutation', user?: { __typename?: 'UserMutation', requestMagicLink?: string | null } | null };

export type SignupMutationVariables = Exact<{
  email: Scalars['String']['input'];
  displayName: Scalars['String']['input'];
}>;


export type SignupMutation = { __typename?: 'Mutation', user?: { __typename?: 'UserMutation', signup?: string | null } | null };

export type CurrentUserQueryVariables = Exact<{ [key: string]: never; }>;


export type CurrentUserQuery = { __typename?: 'Query', user?: { __typename?: 'UserQuery', current?: { __typename?: 'User', id: any, email: string, photo?: string | null, joinDate: any, displayName: string } | null } | null };

export type RecipeBySlugQueryVariables = Exact<{
  slug: Scalars['String']['input'];
  authorDisplayName: Scalars['String']['input'];
}>;


export type RecipeBySlugQuery = { __typename?: 'Query', recipe?: { __typename?: 'RecipeQuery', bySlug?: { __typename?: 'Recipe', id: any, initialPublishDate: any, author: { __typename?: 'User', displayName: string }, forkedFrom?: { __typename?: 'RecipeRevision', title: string, id: any, recipe: { __typename?: 'Recipe', author: { __typename?: 'User', displayName: string } } } | null, featuredRevision?: { __typename?: 'RecipeRevision', id: any, publishDate: any, photo?: string | null, title: string, rating?: number | null, recipeDescription?: string | null, ingredients: Array<{ __typename?: 'RecipeIngredient', id: string, quantity: number, comment?: string | null, unit: { __typename?: 'MeasurementUnit', name: string }, ingredient: { __typename?: 'Ingredient', name: string } }>, steps: Array<{ __typename?: 'RecipeStep', id: string, content: string, index: number, photo?: string | null }> } | null } | null } | null };


export const CheckUserSignupDocument = gql`
    query CheckUserSignup($email: String!, $displayName: String!) {
  user {
    byEmail(email: $email) {
      email
    }
    byDisplayName(displayName: $displayName) {
      displayName
    }
  }
}
    `;
export const ListRecipesDocument = gql`
    query ListRecipes {
  recipe {
    list {
      items {
        slug
        id
        author {
          displayName
        }
        featuredRevision {
          recipeDescription
          publishDate
          rating
          title
        }
      }
    }
  }
}
    `;
export const LoginDocument = gql`
    mutation Login($token: String!, $code: String!) {
  user {
    login(token: $token, code: $code) {
      token
    }
  }
}
    `;
export const LogoutDocument = gql`
    mutation Logout {
  user {
    logout
  }
}
    `;
export const RequestMagicLinkDocument = gql`
    mutation RequestMagicLink($email: String!) {
  user {
    requestMagicLink(email: $email)
  }
}
    `;
export const SignupDocument = gql`
    mutation Signup($email: String!, $displayName: String!) {
  user {
    signup(email: $email, displayName: $displayName)
  }
}
    `;
export const CurrentUserDocument = gql`
    query CurrentUser {
  user {
    current {
      id
      email
      photo
      joinDate
      displayName
    }
  }
}
    `;
export const RecipeBySlugDocument = gql`
    query RecipeBySlug($slug: String!, $authorDisplayName: String!) {
  recipe {
    bySlug(slug: $slug, authorDisplayName: $authorDisplayName) {
      id
      initialPublishDate
      author {
        displayName
      }
      forkedFrom {
        title
        id
        recipe {
          author {
            displayName
          }
        }
      }
      featuredRevision {
        id
        publishDate
        photo
        title
        rating
        recipeDescription
        ingredients {
          id
          unit {
            name
          }
          ingredient {
            name
          }
          quantity
          comment
        }
        steps {
          id
          content
          index
          photo
        }
      }
    }
  }
}
    `;

export type SdkFunctionWrapper = <T>(action: (requestHeaders?:Record<string, string>) => Promise<T>, operationName: string, operationType?: string, variables?: any) => Promise<T>;


const defaultWrapper: SdkFunctionWrapper = (action, _operationName, _operationType, _variables) => action();

export function getSdk(client: GraphQLClient, withWrapper: SdkFunctionWrapper = defaultWrapper) {
  return {
    CheckUserSignup(variables: CheckUserSignupQueryVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<CheckUserSignupQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<CheckUserSignupQuery>(CheckUserSignupDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'CheckUserSignup', 'query', variables);
    },
    ListRecipes(variables?: ListRecipesQueryVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<ListRecipesQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<ListRecipesQuery>(ListRecipesDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'ListRecipes', 'query', variables);
    },
    Login(variables: LoginMutationVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<LoginMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<LoginMutation>(LoginDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'Login', 'mutation', variables);
    },
    Logout(variables?: LogoutMutationVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<LogoutMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<LogoutMutation>(LogoutDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'Logout', 'mutation', variables);
    },
    RequestMagicLink(variables: RequestMagicLinkMutationVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<RequestMagicLinkMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<RequestMagicLinkMutation>(RequestMagicLinkDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'RequestMagicLink', 'mutation', variables);
    },
    Signup(variables: SignupMutationVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<SignupMutation> {
      return withWrapper((wrappedRequestHeaders) => client.request<SignupMutation>(SignupDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'Signup', 'mutation', variables);
    },
    CurrentUser(variables?: CurrentUserQueryVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<CurrentUserQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<CurrentUserQuery>(CurrentUserDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'CurrentUser', 'query', variables);
    },
    RecipeBySlug(variables: RecipeBySlugQueryVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<RecipeBySlugQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<RecipeBySlugQuery>(RecipeBySlugDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'RecipeBySlug', 'query', variables);
    }
  };
}
export type Sdk = ReturnType<typeof getSdk>;