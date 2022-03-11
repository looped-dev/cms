import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Email: any;
  Map: any;
  MongoID: any;
  MongoTime: any;
  Time: any;
};

export type FacebookCard = {
  description?: Maybe<Scalars['String']>;
  image?: Maybe<Scalars['String']>;
  title?: Maybe<Scalars['String']>;
  type?: Maybe<Scalars['String']>;
  url?: Maybe<Scalars['String']>;
};

export type FacebookCardInput = {
  description?: InputMaybe<Scalars['String']>;
  image?: InputMaybe<Scalars['String']>;
  title?: InputMaybe<Scalars['String']>;
  type?: InputMaybe<Scalars['String']>;
  url?: InputMaybe<Scalars['String']>;
};

export type Image = {
  alt?: Maybe<Scalars['String']>;
  caption?: Maybe<Scalars['String']>;
  description?: Maybe<Scalars['String']>;
  id: Scalars['ID'];
  sizes?: Maybe<Sizes>;
  slug: Scalars['String'];
  url: Scalars['String'];
};

export type InitialSetupInput = {
  email: Scalars['Email'];
  name: Scalars['String'];
  password: Scalars['String'];
  siteName: Scalars['String'];
};

export type InitialSetupResponse = {
  accessToken: Scalars['String'];
  refreshToken: Scalars['String'];
  staff: Staff;
};

export type Member = {
  createdAt: Scalars['Time'];
  email: Scalars['Email'];
  id: Scalars['ID'];
  /** Email address verification is vital for sending subscription */
  isEmailVerified: Scalars['Boolean'];
  name: Scalars['String'];
  /** Password is optional as members might not need to login */
  password?: Maybe<Scalars['String']>;
  subscription: Array<MemberSubscription>;
  updatedAt: Scalars['Time'];
};

export type MemberSubscription = {
  createdAt: Scalars['Time'];
  description: Scalars['String'];
  id: Scalars['ID'];
  /** For free subscriptions, this is set to 0 */
  price?: Maybe<Scalars['String']>;
  title: Scalars['String'];
  updatedAt: Scalars['Time'];
};

export type Mutation = {
  /**
   * Create initial Staff for the site and Site Name and login the user before
   * redirecting them to the dashboard
   */
  initialSetup: InitialSetupResponse;
  staffAcceptInvite: Staff;
  staffChangePassword: Staff;
  staffDelete: Staff;
  staffForgotPassword: Staff;
  staffInvite: Staff;
  staffLogin: StaffLoginResponse;
  staffLogout: Scalars['Boolean'];
  staffRefreshToken: StaffLoginResponse;
  staffResetPassword: Staff;
  staffUpdate: Staff;
  updatePage?: Maybe<Page>;
  updatePageStatus?: Maybe<Page>;
  updatePost?: Maybe<Post>;
  updatePostStatus?: Maybe<Post>;
  updateSiteSettings: SiteSettings;
};

export type MutationInitialSetupArgs = {
  input: InitialSetupInput;
};

export type MutationStaffAcceptInviteArgs = {
  input: StaffAcceptInviteInput;
};

export type MutationStaffChangePasswordArgs = {
  input: StaffChangePasswordInput;
};

export type MutationStaffDeleteArgs = {
  input: StaffDeleteInput;
};

export type MutationStaffForgotPasswordArgs = {
  input: StaffForgotPasswordInput;
};

export type MutationStaffInviteArgs = {
  input: StaffInviteInput;
};

export type MutationStaffLoginArgs = {
  input: StaffLoginInput;
};

export type MutationStaffRefreshTokenArgs = {
  input: StaffRefreshTokenInput;
};

export type MutationStaffResetPasswordArgs = {
  input: StaffResetPasswordInput;
};

export type MutationStaffUpdateArgs = {
  input: StaffUpdateInput;
};

export type MutationUpdatePageArgs = {
  input: UpdatePageInput;
};

export type MutationUpdatePageStatusArgs = {
  input: UpdatePageStatusInput;
};

export type MutationUpdatePostArgs = {
  input: UpdatePostInput;
};

export type MutationUpdatePostStatusArgs = {
  input: UpdatePostStatusInput;
};

export type MutationUpdateSiteSettingsArgs = {
  input: SiteSettingsInput;
};

export type Page = {
  content: Scalars['String'];
  createdAt: Scalars['Time'];
  excerpt?: Maybe<Scalars['String']>;
  featuredImage?: Maybe<Image>;
  id: Scalars['ID'];
  publishedAt: Scalars['Time'];
  /** SEO metadata details for the page */
  seo?: Maybe<Seo>;
  slug: Scalars['String'];
  title: Scalars['String'];
  updatedAt: Scalars['Time'];
};

export type Post = {
  content: Scalars['String'];
  createdAt: Scalars['Time'];
  excerpt?: Maybe<Scalars['String']>;
  featuredImage?: Maybe<Image>;
  id: Scalars['ID'];
  isFeatured: Scalars['Boolean'];
  /**
   * Members who have access to this post - this is determined by subscription groups
   * they are part of.
   */
  postAccess?: Maybe<Array<MemberSubscription>>;
  publishedAt: Scalars['Time'];
  /** SEO metadata details for the post or page */
  seo?: Maybe<Seo>;
  slug: Scalars['String'];
  title: Scalars['String'];
  updatedAt: Scalars['Time'];
};

export enum PostOrPageStatus {
  Draft = 'DRAFT',
  Pending = 'PENDING',
  Published = 'PUBLISHED',
  Scheduled = 'SCHEDULED',
  Trashed = 'TRASHED',
}

export type Query = {
  getPage?: Maybe<Page>;
  getPageByID?: Maybe<Page>;
  getPost?: Maybe<Post>;
  getPostByID?: Maybe<Post>;
  getPosts?: Maybe<Array<Post>>;
  isSiteSetup: Scalars['Boolean'];
  settings: SiteSettings;
};

export type QueryGetPageArgs = {
  slug: Scalars['String'];
};

export type QueryGetPageByIdArgs = {
  id: Scalars['String'];
};

export type QueryGetPostArgs = {
  slug: Scalars['String'];
};

export type QueryGetPostByIdArgs = {
  id: Scalars['String'];
};

export type QueryGetPostsArgs = {
  page?: InputMaybe<Scalars['Int']>;
  perPage?: InputMaybe<Scalars['Int']>;
};

export type Seo = {
  description?: Maybe<Scalars['String']>;
  facebook?: Maybe<FacebookCard>;
  image?: Maybe<Scalars['String']>;
  title?: Maybe<Scalars['String']>;
  twitter?: Maybe<TwitterCard>;
};

export type SeoInput = {
  description?: InputMaybe<Scalars['String']>;
  facebook?: InputMaybe<FacebookCardInput>;
  image?: InputMaybe<Scalars['String']>;
  title?: InputMaybe<Scalars['String']>;
  twitter?: InputMaybe<TwitterCardInput>;
};

export type SiteSettings = {
  baseURL: Scalars['String'];
  seo: Seo;
  siteName: Scalars['String'];
};

export type SiteSettingsInput = {
  baseURL: Scalars['String'];
  seo: SeoInput;
  siteName: Scalars['String'];
};

export type Size = {
  height: Scalars['Int'];
  url: Scalars['String'];
  width: Scalars['Int'];
};

export type Sizes = {
  full?: Maybe<Size>;
  large?: Maybe<Size>;
  medium?: Maybe<Size>;
  medium_large?: Maybe<Size>;
  thumbnail?: Maybe<Size>;
};

export type Staff = {
  createdAt: Scalars['MongoTime'];
  email: Scalars['Email'];
  emailVerified: Scalars['Boolean'];
  id: Scalars['MongoID'];
  name: Scalars['String'];
  role: StaffRole;
  updatedAt: Scalars['MongoTime'];
};

export type StaffAcceptInviteInput = {
  code: Scalars['String'];
  confirmPassword: Scalars['String'];
  email: Scalars['String'];
  name: Scalars['String'];
  password: Scalars['String'];
};

export type StaffChangePasswordInput = {
  confirmPassword: Scalars['String'];
  id: Scalars['ID'];
  password: Scalars['String'];
};

export type StaffDeleteInput = {
  id: Scalars['ID'];
};

export type StaffForgotPasswordInput = {
  email: Scalars['Email'];
};

export type StaffInviteInput = {
  email: Scalars['Email'];
  role: StaffRole;
};

export type StaffLoginInput = {
  email: Scalars['Email'];
  password: Scalars['String'];
};

export type StaffLoginResponse = {
  accessToken: Scalars['String'];
  refreshToken: Scalars['String'];
  staff: Staff;
};

export type StaffRefreshTokenInput = {
  accessToken: Scalars['String'];
  refreshToken: Scalars['String'];
};

export type StaffRefreshTokenResponse = {
  accessToken: Scalars['String'];
  refreshToken: Scalars['String'];
};

export type StaffRegisterInput = {
  email: Scalars['Email'];
  name: Scalars['String'];
  password: Scalars['String'];
  role: StaffRole;
};

export type StaffResetPasswordInput = {
  code: Scalars['String'];
  confirmPassword: Scalars['String'];
  email: Scalars['Email'];
  password: Scalars['String'];
};

export enum StaffRole {
  /**
   * Trusted staff user who should be able to manage all content and users, as well
   * as site settings and options.
   */
  Administrator = 'ADMINISTRATOR',
  /**
   * A trusted user who can create, edit and publish their own posts, but can’t
   * modify others.
   */
  Author = 'AUTHOR',
  /**
   * Can invite and manage other Authors and Contributors, as well as edit and
   * publish any posts on the site.
   */
  Editor = 'EDITOR',
  /** Has full site access and is the owner of the site and cannot be deleted. */
  Owner = 'OWNER',
}

export type StaffUpdateInput = {
  email: Scalars['Email'];
  name: Scalars['String'];
};

export type Tag = {
  createdAt: Scalars['Time'];
  description?: Maybe<Scalars['String']>;
  id: Scalars['ID'];
  image?: Maybe<Image>;
  slug: Scalars['String'];
  title: Scalars['String'];
  updatedAt: Scalars['Time'];
};

export type TwitterCard = {
  card?: Maybe<Scalars['String']>;
  creator?: Maybe<Scalars['String']>;
  description?: Maybe<Scalars['String']>;
  image?: Maybe<Scalars['String']>;
  site?: Maybe<Scalars['String']>;
  title?: Maybe<Scalars['String']>;
};

export type TwitterCardInput = {
  card?: InputMaybe<Scalars['String']>;
  creator?: InputMaybe<Scalars['String']>;
  description?: InputMaybe<Scalars['String']>;
  image?: InputMaybe<Scalars['String']>;
  site?: InputMaybe<Scalars['String']>;
  title?: InputMaybe<Scalars['String']>;
};

export type UpdatePageInput = {
  content?: InputMaybe<Scalars['String']>;
  excerpt?: InputMaybe<Scalars['String']>;
  featuredImage?: InputMaybe<Scalars['String']>;
  /** If a post is featured, default to false. */
  isFeatured?: InputMaybe<Scalars['Boolean']>;
  /** List of subscription groups with access to the the post */
  postAccess?: InputMaybe<Array<Scalars['ID']>>;
  seo?: InputMaybe<SeoInput>;
  title?: InputMaybe<Scalars['String']>;
};

export type UpdatePageStatusInput = {
  id: Scalars['String'];
  status: PostOrPageStatus;
};

export type UpdatePostInput = {
  content?: InputMaybe<Scalars['String']>;
  excerpt?: InputMaybe<Scalars['String']>;
  featuredImage?: InputMaybe<Scalars['String']>;
  /** If a post is featured, default to false. */
  isFeatured?: InputMaybe<Scalars['Boolean']>;
  /** List of subscription groups with access to the the post */
  postAccess?: InputMaybe<Array<Scalars['ID']>>;
  seo?: InputMaybe<SeoInput>;
  title?: InputMaybe<Scalars['String']>;
};

export type UpdatePostStatusInput = {
  id: Scalars['String'];
  status: PostOrPageStatus;
};

export type StaffLoginMutationVariables = Exact<{
  input: StaffLoginInput;
}>;

export type StaffLoginMutation = {
  staffLogin: {
    accessToken: string;
    refreshToken: string;
    staff: {
      id: any;
      name: string;
      email: any;
      role: StaffRole;
      createdAt: any;
      updatedAt: any;
    };
  };
};

export type SetupSiteMutationVariables = Exact<{
  input: InitialSetupInput;
}>;

export type SetupSiteMutation = {
  initialSetup: {
    refreshToken: string;
    accessToken: string;
    staff: { id: any };
  };
};

export type IsSiteSetupQueryVariables = Exact<{ [key: string]: never }>;

export type IsSiteSetupQuery = { isSiteSetup: boolean };

export const StaffLoginDocument = gql`
  mutation StaffLogin($input: StaffLoginInput!) {
    staffLogin(input: $input) {
      accessToken
      refreshToken
      staff {
        id
        name
        email
        role
        createdAt
        updatedAt
      }
    }
  }
`;

@Injectable({
  providedIn: 'root',
})
export class StaffLoginGQL extends Apollo.Mutation<
  StaffLoginMutation,
  StaffLoginMutationVariables
> {
  document = StaffLoginDocument;

  constructor(apollo: Apollo.Apollo) {
    super(apollo);
  }
}
export const SetupSiteDocument = gql`
  mutation SetupSite($input: InitialSetupInput!) {
    initialSetup(input: $input) {
      staff {
        id
      }
      refreshToken
      accessToken
    }
  }
`;

@Injectable({
  providedIn: 'root',
})
export class SetupSiteGQL extends Apollo.Mutation<
  SetupSiteMutation,
  SetupSiteMutationVariables
> {
  document = SetupSiteDocument;

  constructor(apollo: Apollo.Apollo) {
    super(apollo);
  }
}
export const IsSiteSetupDocument = gql`
  query isSiteSetup {
    isSiteSetup
  }
`;

@Injectable({
  providedIn: 'root',
})
export class IsSiteSetupGQL extends Apollo.Query<
  IsSiteSetupQuery,
  IsSiteSetupQueryVariables
> {
  document = IsSiteSetupDocument;

  constructor(apollo: Apollo.Apollo) {
    super(apollo);
  }
}
