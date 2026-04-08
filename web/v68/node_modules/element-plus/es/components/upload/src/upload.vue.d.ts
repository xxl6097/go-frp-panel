import { Awaitable, Mutable } from "../../../utils/typescript.js";
import "../../../utils/index.js";
import { ListType, UploadData, UploadFile, UploadHooks, UploadProps, UploadRawFile, UploadRequestHandler, UploadStatus, UploadUserFile } from "./upload.js";
import * as vue from "vue";

//#region ../../packages/components/upload/src/upload.vue.d.ts
declare var __VLS_10: {
    file: UploadFile;
    index: number;
  }, __VLS_21: {}, __VLS_23: {}, __VLS_33: {}, __VLS_35: {}, __VLS_37: {}, __VLS_39: {}, __VLS_50: {
    file: UploadFile;
    index: number;
  };
type __VLS_Slots = {} & {
  file?: (props: typeof __VLS_10) => any;
} & {
  trigger?: (props: typeof __VLS_21) => any;
} & {
  default?: (props: typeof __VLS_23) => any;
} & {
  trigger?: (props: typeof __VLS_33) => any;
} & {
  default?: (props: typeof __VLS_35) => any;
} & {
  default?: (props: typeof __VLS_37) => any;
} & {
  tip?: (props: typeof __VLS_39) => any;
} & {
  file?: (props: typeof __VLS_50) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<UploadProps>, {
  readonly beforeUpload: () => void;
  readonly onRemove: () => void;
  readonly onChange: () => void;
  readonly onPreview: () => void;
  readonly onSuccess: () => void;
  readonly onProgress: () => void;
  readonly onError: () => void;
  readonly onExceed: () => void;
  readonly action: "#";
  readonly method: "post";
  readonly data: () => Mutable<{}>;
  readonly name: "file";
  readonly showFileList: true;
  readonly accept: "";
  readonly fileList: () => never[];
  readonly autoUpload: true;
  readonly listType: "text";
  readonly httpRequest: UploadRequestHandler;
  readonly disabled: undefined;
}>, {
  /** @description cancel upload request */abort: (file?: UploadFile) => void; /** @description upload the file list manually */
  submit: () => void; /** @description clear the file list  */
  clearFiles: (states?: UploadStatus[]) => void; /** @description select the file manually */
  handleStart: (rawFile: UploadRawFile) => void; /** @description remove the file manually */
  handleRemove: (file: UploadFile | UploadRawFile) => void;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<UploadProps>, {
  readonly beforeUpload: () => void;
  readonly onRemove: () => void;
  readonly onChange: () => void;
  readonly onPreview: () => void;
  readonly onSuccess: () => void;
  readonly onProgress: () => void;
  readonly onError: () => void;
  readonly onExceed: () => void;
  readonly action: "#";
  readonly method: "post";
  readonly data: () => Mutable<{}>;
  readonly name: "file";
  readonly showFileList: true;
  readonly accept: "";
  readonly fileList: () => never[];
  readonly autoUpload: true;
  readonly listType: "text";
  readonly httpRequest: UploadRequestHandler;
  readonly disabled: undefined;
}>>>, {
  onChange: UploadHooks["onChange"];
  disabled: boolean;
  name: string;
  onError: UploadHooks["onError"];
  onProgress: UploadHooks["onProgress"];
  data: Awaitable<UploadData> | ((rawFile: UploadRawFile) => Awaitable<UploadData>);
  beforeUpload: UploadHooks["beforeUpload"];
  onRemove: UploadHooks["onRemove"];
  onPreview: UploadHooks["onPreview"];
  onSuccess: UploadHooks["onSuccess"];
  onExceed: UploadHooks["onExceed"];
  action: string;
  method: string;
  showFileList: boolean;
  accept: string;
  fileList: UploadUserFile[];
  autoUpload: boolean;
  listType: ListType;
  httpRequest: UploadRequestHandler;
}, {}>;
declare const __VLS_export: __VLS_WithSlots<typeof __VLS_base, __VLS_Slots>;
declare const _default: typeof __VLS_export;
type __VLS_TypePropsToOption<T> = { [K in keyof T]-?: {} extends Pick<T, K> ? {
  type: vue.PropType<Required<T>[K]>;
} : {
  type: vue.PropType<T[K]>;
  required: true;
} };
type __VLS_WithDefaults<P, D> = { [K in keyof Pick<P, keyof P>]: K extends keyof D ? __VLS_PrettifyLocal<P[K] & {
  default: D[K];
}> : P[K] };
type __VLS_WithSlots<T, S> = T & {
  new (): {
    $slots: S;
  };
};
type __VLS_PrettifyLocal<T> = (T extends any ? { [K in keyof T]: T[K] } : { [K in keyof T as K]: T[K] }) & {};
//#endregion
export { _default };