import * as vue from "vue";
import { Component } from "vue";
import { Loading } from "@element-plus/icons-vue";

//#region ../../packages/utils/vue/icon.d.ts
type IconPropType = string | Component;
declare const iconPropType: vue.PropType<string | Component>;
declare const CloseComponents: {
  Close: any;
};
declare const TypeComponents: {
  Close: any;
  SuccessFilled: any;
  InfoFilled: any;
  WarningFilled: any;
  CircleCloseFilled: any;
};
declare const TypeComponentsMap: {
  primary: any;
  success: any;
  warning: any;
  error: any;
  info: any;
};
declare const ValidateComponentsMap: {
  validating: any;
  success: any;
  error: any;
};
type IconComponent = typeof Loading;
//#endregion
export { CloseComponents, IconComponent, IconPropType, TypeComponents, TypeComponentsMap, ValidateComponentsMap, iconPropType };