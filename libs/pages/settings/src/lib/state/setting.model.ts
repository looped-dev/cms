import { SiteSettings } from '@looped-cms/graphql';

export type Setting = SiteSettings;

export function createSetting(params: Partial<Setting>) {
  return { ...params } as Setting;
}
