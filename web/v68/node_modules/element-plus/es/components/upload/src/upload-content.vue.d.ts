import { Awaitable, Mutable } from "../../../utils/typescript.js";
import "../../../utils/index.js";
import { UploadAjaxError } from "./ajax.js";
import { ListType, UploadData, UploadFile, UploadHooks, UploadProgressEvent, UploadRawFile, UploadRequestHandler, UploadUserFile } from "./upload.js";
import { UploadContentProps } from "./upload-content.js";
import * as vue from "vue";

//#region ../../packages/components/upload/src/upload-content.vue.d.ts
declare var __VLS_9: {}, __VLS_11: {};
type __VLS_Slots = {} & {
  default?: (props: typeof __VLS_9) => any;
} & {
  default?: (props: typeof __VLS_11) => any;
};
declare const __VLS_base: vue.DefineComponent<__VLS_WithDefaults<__VLS_TypePropsToOption<UploadContentProps>, {
  readonly beforeUpload: () => void;
  readonly onRemove: () => void;
  readonly onStart: () => void;
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
  abort: (file?: UploadFile) => void;
  upload: (rawFile: UploadRawFile) => Promise<void>;
}, unknown, {}, {}, vue.ComponentOptionsMixin, vue.ComponentOptionsMixin, {}, string, vue.PublicProps, Readonly<vue.ExtractPropTypes<__VLS_WithDefaults<__VLS_TypePropsToOption<UploadContentProps>, {
  readonly beforeUpload: () => void;
  readonly onRemove: () => void;
  readonly onStart: () => void;
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
  disabled: boolean;
  name: string;
  onError: (err: UploadAjaxError, rawFile: UploadRawFile) => void;
  onProgress: (evt: UploadProgressEvent, rawFile: UploadRawFile) => void;
  data: Awaitable<UploadData> | ((rawFile: UploadRawFile) => Awaitable<UploadData>);
  beforeUpload: UploadHooks["beforeUpload"];
  onRemove: (file: UploadFile | UploadRawFile) => void;
  onSuccess: (response: any, rawFile: UploadRawFile) => unknown;
  onExceed: UploadHooks["onExceed"];
  action: string;
  method: string;
  showFileList: boolean;
  accept: string;
  fileList: UploadUserFile[];
  autoUpload: boolean;
  listType: ListType;
  httpRequest: UploadRequestHandler;
  onStart: (rawFile: UploadRawFile) => void;
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